package rpcclient

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"

	m "github.com/linkit360/go-utils/metrics"
)

var cli *Client

type Client struct {
	connection *rpc.Client
	conf       ClientConfig
	m          *Metrics
}
type ClientConfig struct {
	Enabled bool   `yaml:"enabled"`
	DSN     string `default:":50307" yaml:"dsn"`
	Timeout int    `default:"10" yaml:"timeout"`
}

type Metrics struct {
	RPCConnectError m.Gauge
	RPCSuccess      m.Gauge
	Connected       prometheus.Gauge
}

func init() {
	log.SetLevel(log.DebugLevel)
}

func initMetrics() *Metrics {
	metrics := &Metrics{
		RPCConnectError: m.NewGauge("rpc", "acceptor", "errors", "RPC call errors"),
		RPCSuccess:      m.NewGauge("rpc", "acceptor", "success", "RPC call success"),
		Connected:       m.PrometheusGauge("rpc", "acceptor", "connected", "rpc connected"),
	}
	go func() {
		for range time.Tick(time.Minute) {
			metrics.RPCConnectError.Update()
			metrics.RPCSuccess.Update()
		}
	}()
	return metrics
}

func Init(clientConf ClientConfig) error {
	if cli != nil {
		return nil
	}
	var err error
	cli = &Client{
		conf: clientConf,
		m:    initMetrics(),
	}
	if !cli.conf.Enabled {
		return nil
	}
	if err = cli.dial(); err != nil {
		err = fmt.Errorf("cli.dial: %s", err.Error())
		log.WithField("error", err.Error()).Error("acceptor rpc client unavialable")
		return err
	}
	log.WithField("conf", fmt.Sprintf("%#v", clientConf)).Info("acceptor rpc client init done")

	return nil
}

func (c *Client) dial() error {
	if !c.conf.Enabled {
		return nil
	}
	if c.connection != nil {
		c.connection.Close()
		c.connection = nil
	}
	conn, err := net.DialTimeout(
		"tcp",
		c.conf.DSN,
		time.Duration(c.conf.Timeout)*time.Second,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"dsn":   c.conf.DSN,
			"error": err.Error(),
		}).Error("dialing acceptor")
		cli.m.Connected.Set(0)
		return err
	}

	c.connection = jsonrpc.NewClient(conn)
	log.WithFields(log.Fields{
		"dsn": c.conf.DSN,
	}).Debug("dialing acceptor")
	cli.m.Connected.Set(1)
	return nil
}

func call(funcName string, req interface{}, res interface{}) error {
	if cli == nil {
		return fmt.Errorf("Acceptor: %s", "disabled")
	}
	if !cli.conf.Enabled {
		return fmt.Errorf("Acceptor: %s", "disabled")
	}

	begin := time.Now()

	retryCount := 0
retry:
	if err := cli.connection.Call(funcName, req, &res); err != nil {
		cli.m.RPCConnectError.Inc()

		if err == rpc.ErrShutdown {
			if retryCount < 2 {
				retryCount = retryCount + 1
				cli.connection.Close()
				cli.dial()

				log.WithFields(log.Fields{
					"retry": retryCount,
					"error": err.Error(),
				}).Debug("retrying..")
				goto retry
			}
		}

		log.WithFields(log.Fields{
			"func":  funcName,
			"error": err.Error(),
		}).Error("call")
		return err
	}
	log.WithFields(log.Fields{
		"func": funcName,
		"took": time.Since(begin),
	}).Debug("rpccall")
	cli.m.RPCSuccess.Inc()
	return nil
}

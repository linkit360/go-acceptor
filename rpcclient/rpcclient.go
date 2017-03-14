package rpcclient

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/vostrok/acceptor/server/src/handlers"
	m "github.com/vostrok/utils/metrics"
)

var cli *Client

type Client struct {
	connection *rpc.Client
	conf       RPCClientConfig
	m          *Metrics
}
type RPCClientConfig struct {
	DSN     string `default:":50307" yaml:"dsn"`
	Timeout int    `default:"10" yaml:"timeout"`
}

type Metrics struct {
	RPCConnectError m.Gauge
	RPCSuccess      m.Gauge
	NotFound        m.Gauge
}

func initMetrics() *Metrics {
	m := &Metrics{
		RPCConnectError: m.NewGauge("rpc", "acceptor", "errors", "RPC call errors"),
		RPCSuccess:      m.NewGauge("rpc", "acceptor", "success", "RPC call success"),
		NotFound:        m.NewGauge("rpc", "acceptor", "404_errors", "RPC 404 errors"),
	}
	go func() {
		for range time.Tick(time.Minute) {
			m.RPCConnectError.Update()
			m.RPCSuccess.Update()
			m.NotFound.Update()
		}
	}()
	return m
}
func Init(clientConf RPCClientConfig) error {
	if cli != nil {
		return nil
	}
	var err error
	cli = &Client{
		conf: clientConf,
		m:    initMetrics(),
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
	if c.connection != nil {
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
		return err
	}
	c.connection = jsonrpc.NewClient(conn)
	log.WithFields(log.Fields{
		"dsn": c.conf.DSN,
	}).Debug("dialing acceptor")
	return nil
}

func call(funcName string, req interface{}, res interface{}) error {
	begin := time.Now()
	if cli.connection == nil {
		cli.dial()
	}
	if err := cli.connection.Call(funcName, req, &res); err != nil {
		cli.m.RPCConnectError.Inc()
		if err == rpc.ErrShutdown {
			log.WithFields(log.Fields{
				"func":  funcName,
				"error": err.Error(),
			}).Fatal("call")
		}
		log.WithFields(log.Fields{
			"func":  funcName,
			"error": err.Error(),
			"type":  fmt.Sprintf("%T", err),
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

func SendAggregatedData(data []handlers.Aggregate) error {
	var res handlers.Response
	err := call(
		"Aggregate.Receive",
		handlers.AggregateData{Aggregated: data},
		&res,
	)
	return err
}

func GetRandomAggregate() handlers.Aggregate {
	return handlers.Aggregate{
		ReportDate:           time.Now().UTC().Unix(),
		CampaignId:           777,
		TotalLPHits:          rand.Int63(),
		TotalLPMsisdnHits:    rand.Int63(),
		TotalMO:              rand.Int63(),
		TotalMOUniq:          rand.Int63(),
		TotalMOSuccessCharge: rand.Int63(),
		TotalPixelsSent:      rand.Int63(),
	}
}

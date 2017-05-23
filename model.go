package rpcclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

var client *http.Client
var config ClientConfig

/*
var cli *Client
type Client struct {
	connection *rpc.Client
	conf       ClientConfig
}
*/

type ClientConfig struct {
	Enabled    bool   `yaml:"enabled"`
	DSN        string `default:":50307" yaml:"dsn"`
	Timeout    int    `default:"10" yaml:"timeout"`
	InstanceId string `default:"" yaml:"instance_id"`
}

func Init(clientConf ClientConfig) error {
	log.SetLevel(log.DebugLevel)

	config = clientConf
	client = &http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
	}

	/*
		cli = &Client{
			conf: clientConf,
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
	*/

	return nil
}

func Call(funcName string, res interface{}, req ...interface{}) error {
	if !config.Enabled {
		return fmt.Errorf("Acceptor Client Disabled")
	}

	var url string = "http://" + config.DSN + "/" + funcName + "?instance_id=" + config.InstanceId
	var err error

	// GET by default
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if len(req) > 0 {
		// POST
		jsonValue, err := json.Marshal(req)
		if err != nil {
			return err
		}

		request, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		if err != nil {
			return err
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	return nil
}

/*
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
		return err
	}

	c.connection = jsonrpc.NewClient(conn)

	log.WithFields(log.Fields{
		"dsn": c.conf.DSN,
	}).Debug("dialing acceptor")

	return nil
}

func call(funcName string, req interface{}, res interface{}) error {
	if cli == nil {
		return fmt.Errorf("Acceptor Client Offline")
	}

	if !cli.conf.Enabled {
		return fmt.Errorf("Acceptor Client Disabled")
	}

	begin := time.Now()
	retryCount := 0
	for {
		if err := cli.connection.Call(funcName, req, &res); err != nil {
			if err == rpc.ErrShutdown {
				if retryCount < 2 {
					retryCount = retryCount + 1
					cli.connection.Close()
					cli.dial()

					log.WithFields(log.Fields{
						"retry": retryCount,
						"error": err.Error(),
					}).Debug("retrying..")

					continue
				}
			}

			log.WithFields(log.Fields{
				"func":  funcName,
				"error": err.Error(),
			}).Error("call")

			return err
		}
		break
	}

	log.WithFields(log.Fields{
		"func": funcName,
		"took": time.Since(begin),
	}).Debug("rpccall")

	return nil
}
*/

package src

import (
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/vostrok/aggregate/rpcclient"
	"github.com/vostrok/aggregate/server/src/config"
	"github.com/vostrok/aggregate/server/src/handlers"
	m "github.com/vostrok/utils/metrics"
)

var appConfig config.AppConfig

func Run() {

	appConfig = config.LoadConfig()

	handlers.InitMetrics()

	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	log.WithField("CPUCount", nuCPU)

	go runGin(appConfig)

	runRPC(appConfig)
}

func runGin(appConfig config.AppConfig) {
	r := gin.New()
	m.AddHandler(r)
	r.GET("/test", TestSend)

	r.Run(":" + appConfig.Server.HttpPort)

	log.WithField("port", appConfig.Server.HttpPort).Info("service port")

}

func runRPC(appConfig config.AppConfig) {

	l, err := net.Listen("tcp", "127.0.0.1:"+appConfig.Server.RPCPort)
	if err != nil {
		log.Fatal("netListen ", err.Error())
	} else {
		log.WithField("port", appConfig.Server.RPCPort).Info("rpc port")
	}

	server := rpc.NewServer()
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	server.RegisterName("Aggregate", &handlers.Aggregate{})

	for {
		if conn, err := l.Accept(); err == nil {
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		} else {
			log.WithField("error", err.Error()).Error("accept")
		}
	}
}

func TestSend(c *gin.Context) {
	begin := time.Now()
	var data = []handlers.Aggregate{}
	data = append(data, rpcclient.GetRandomAggregate())
	data = append(data, rpcclient.GetRandomAggregate())

	if err := rpcclient.Init(appConfig.Client); err != nil {
		log.WithField("error", err.Error()).Fatal("cannot init client")
	}
	if err := rpcclient.SendAggregatedData(data); err != nil {
		log.WithField("error", err.Error()).Error("accept")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"since": time.Since(begin)})
}

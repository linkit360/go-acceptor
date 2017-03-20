package handlers

import (
	log "github.com/Sirupsen/logrus"

	"github.com/linkit360/go-acceptor/server/src/base"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

type Response struct{}

type Aggregate struct {
}

type AggregateData struct {
	Aggregated []base.Aggregate `json:"aggregated,omitempty"`
}

func (rpc *Aggregate) Receive(req AggregateData, res *Response) error {
	return base.SaveRows(req.Aggregated)
}

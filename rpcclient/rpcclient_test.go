package rpcclient

import (
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/go-acceptor/server/src/base"
	"github.com/stretchr/testify/assert"
)

func init() {
	c := ClientConfig{
		DSN:     "localhost:50313",
		Timeout: 10,
	}

	if err := Init(c); err != nil {
		log.WithField("error", err.Error()).Fatal("cannot init client")
	}
}

func TestGetAllDestinations(t *testing.T) {
	data := []base.Aggregate{GetRandomAggregate(), GetRandomAggregate()}
	err := SendAggregatedData(data)
	assert.NoError(t, err, "No error while send the aggregate data")
}

package rpcclient

import (
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/vostrok/acceptor/server/src/handlers"
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
	data := []handlers.Aggregate{GetRandomAggregate(), GetRandomAggregate()}
	err := SendAggregatedData(data)
	assert.NoError(t, err, "No error while send the aggregate data")
}

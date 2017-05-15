package rpcclient

import (
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	acceptor "github.com/linkit360/go-acceptor-structs"
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
	data := []acceptor.Aggregate{GetRandomAggregate(), GetRandomAggregate()}
	_, err := SendAggregatedData(data)
	assert.NoError(t, err, "No error while send the aggregate data")
}

func TestGetBlackList(t *testing.T) {
	msisdns, err := GetBlackListed("cheese")
	assert.NoError(t, err, "Cheese blacklist")
	assert.Equal(t, 0, len(msisdns), "Count of blacklisted on cheese")
}

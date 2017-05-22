package rpcclient

import (
	"math/rand"
	"time"

	acceptor "github.com/linkit360/go-acceptor-structs"
)

func SendAggregatedData(data []acceptor.Aggregate) (acceptor.Response, error) {
	var res acceptor.Response
	err := call(
		"Aggregate.Receive",
		acceptor.AggregateData{Aggregated: data},
		&res,
	)
	return res, err
}

func GetRandomAggregate() acceptor.Aggregate {
	return acceptor.Aggregate{
		ReportAt:             time.Now().UTC().Unix(),
		CampaignCode:         "290",
		ProviderName:         "cheese",
		OperatorCode:         52000,
		LpHits:               rand.Int63n(200),
		LpMsisdnHits:         rand.Int63n(150),
		MoTotal:              rand.Int63n(200),
		MoChargeSuccess:      rand.Int63n(200),
		MoChargeSum:          1000.,
		MoChargeFailed:       rand.Int63n(200),
		MoRejected:           rand.Int63n(200),
		RenewalTotal:         rand.Int63n(200),
		RenewalChargeSuccess: rand.Int63n(200),
		RenewalChargeSum:     12312.,
		RenewalFailed:        rand.Int63n(200),
		Pixels:               rand.Int63n(200),
	}
}

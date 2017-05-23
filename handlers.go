package rpcclient

import (
	"math/rand"
	"time"

	acceptorStructs "github.com/linkit360/go-acceptor-structs"
)

/*
func SendAggregatedData(data []acceptorStructs.Aggregate) (acceptorStructs.AggregateResponse, error) {
	var res acceptorStructs.AggregateResponse
	err := call(
		"Aggregate.Receive",
		acceptorStructs.AggregateRequest{Aggregated: data},
		&res,
	)
	return res, err
}
*/

func GetRandomAggregate() acceptorStructs.Aggregate {
	return acceptorStructs.Aggregate{
		ReportAt:             time.Now().UTC().Unix(),
		CampaignCode:         "290",
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

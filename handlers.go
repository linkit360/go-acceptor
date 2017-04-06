package rpcclient

import (
	"math/rand"
	"time"

	acceptor "github.com/linkit360/go-acceptor-structs"
)

func SendAggregatedData(data []acceptor.Aggregate) error {
	var res acceptor.Response
	err := call(
		"Aggregate.Receive",
		acceptor.AggregateData{Aggregated: data},
		&res,
	)
	return err
}

func BlackListGet(providerName string) ([]string, error) {
	var res acceptor.BlackListResponse
	err := call(
		"BlackList.Get",
		acceptor.BlackListGetParams{ProviderName: providerName},
		&res,
	)
	if err != nil {
		return []string{}, err
	}
	return res.Msisdns, nil
}

func BlackListGetNew(providerName string, time string) ([]string, error) {
	var res acceptor.BlackListResponse
	err := call(
		"BlackList.GetNew",
		acceptor.BlackListGetParams{ProviderName: providerName, Time: time},
		&res,
	)
	if err != nil {
		return []string{}, err
	}
	return res.Msisdns, nil
}

func GetRandomAggregate() acceptor.Aggregate {
	return acceptor.Aggregate{
		ReportAt:     time.Now().UTC().Unix(),
		CampaignId:   rand.Int63n(9),
		ProviderName: "cheese",
		OperatorCode: 52000,
		LpHits:       rand.Int63n(200),
		LpMsisdnHits: rand.Int63n(150),
		Mo:           rand.Int63n(200),
		MoUniq:       rand.Int63n(200),
		MoSuccess:    rand.Int63n(150),
		RetrySuccess: rand.Int63n(150),
		Pixels:       rand.Int63n(200),
	}
}

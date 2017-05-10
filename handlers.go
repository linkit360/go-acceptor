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

func GetContents(providerName string) (map[int64]acceptor.Content, error) {
	var res acceptor.GetContentsResponse
	err := call(
		"Content.GetAll",
		acceptor.GetContentParams{ProviderName: providerName},
		&res,
	)
	if err != nil {
		return map[int64]acceptor.Content{}, err
	}
	return res.Contents, nil
}

func GetServices(providerName string) (map[int64]acceptor.Service, error) {
	var res acceptor.GetServicesResponse
	err := call(
		"Service.GetAll",
		acceptor.GetServicesParams{ProviderName: providerName},
		&res,
	)
	if err != nil {
		return map[int64]acceptor.Service{}, err
	}
	return res.Services, nil
}

func GetBlackListed(providerName string) ([]string, error) {
	var res acceptor.BlackListResponse
	err := call(
		"BlackList.GetAll",
		acceptor.BlackListGetParams{ProviderName: providerName},
		&res,
	)
	if err != nil {
		return []string{}, err
	}
	return res.Msisdns, nil
}

func GetNewBlackListed(providerName string, time string) ([]string, error) {
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

func CampaignsGet(provider string) ([]acceptor.CampaignsCampaign, error) {
	var res acceptor.CampaignsResponse
	err := call(
		"Campaigns.Get",
		acceptor.CampaignsGetParams{Provider: provider},
		&res,
	)
	if err != nil {
		return []acceptor.CampaignsCampaign{}, err
	}
	return res.Campaigns, nil
}

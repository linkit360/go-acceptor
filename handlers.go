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

// map by uuids
func GetServices(providerName string) (map[string]acceptor.Service, error) {
	var res acceptor.GetServicesResponse
	err := call(
		"Service.GetAll",
		acceptor.GetServicesParams{ProviderName: providerName},
		&res,
	)
	if err != nil {
		return map[string]acceptor.Service{}, err
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

func CampaignsGet(provider string) ([]acceptor.Campaign, error) {
	var res acceptor.CampaignsResponse
	err := call(
		"Campaign.GetAll",
		acceptor.CampaignsGetParams{Provider: provider},
		&res,
	)
	if err != nil {
		return []acceptor.Campaign{}, err
	}
	return res.Campaigns, nil
}

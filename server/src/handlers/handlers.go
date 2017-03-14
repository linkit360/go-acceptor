package handlers

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

type Aggregate struct {
	ReportDate           int64 `json:"report_date,omitempty"`
	CampaignId           int64 `json:"id_campaign,omitempty"`
	TotalLPHits          int64 `json:"total_lp_hits,omitempty"`
	TotalLPMsisdnHits    int64 `json:"total_lp_msisdn_hits,omitempty"`
	TotalMO              int64 `json:"total_mo,omitempty"`
	TotalMOUniq          int64 `json:"total_mo_uniq,omitempty"`
	TotalMOSuccessCharge int64 `json:"total_mo_success_charge,omitempty"`
	TotalPixelsSent      int64 `json:"total_pixels_sent,omitempty"`
}

type AggregateData struct {
	Aggregated []Aggregate `json:"aggregated,omitempty"`
}

type Response struct{}

func (rpc *Aggregate) Receive(req AggregateData, res *Response) error {

	data, _ := json.Marshal(req)
	log.Debugf("%s", string(data))
	*res = Response{}
	success.Inc()
	return nil
}

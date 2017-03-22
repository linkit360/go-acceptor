package base

import (
	"database/sql"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/linkit360/go-utils/db"
)

var pgsql *sql.DB
var config db.DataBaseConfig

type Aggregate struct {
	ReportAt     int64  `json:"report_at,omitempty"`
	CampaignId   int64  `json:"id_campaign,omitempty"`
	ProviderName string `json:"provider_name,omitempty"`
	OperatorCode int64  `json:"operator_code,omitempty"`
	LpHits       int64  `json:"lp_hits,omitempty"`
	LpMsisdnHits int64  `json:"lp_msisdn_hits,omitempty"`
	Mo           int64  `json:"mo,omitempty"`
	MoUniq       int64  `json:"mo_uniq,omitempty"`
	MoSuccess    int64  `json:"mo_success,omitempty"`
	RetrySuccess int64  `json:"retry_success,omitempty"`
	Pixels       int64  `json:"pixels,omitempty"`
}

func Init(dbConfig db.DataBaseConfig) {
	config = dbConfig
	pgsql = db.Init(config)
}

func SaveRows(rows []Aggregate) error {
	var query string = fmt.Sprintf(
		"INSERT INTO %sreports ("+

			"report_at, "+
			"id_campaign, "+
			"provider_name, "+
			"operator_code, "+
			"lp_hits, "+
			"lp_msisdn_hits, "+
			"mo, "+
			"mo_uniq, "+
			"mo_success, "+"retry_success, "+
			"pixels"+

			") VALUES ("+

			"to_timestamp($1), $2, $3, $4, $5, $6, $7, $8, $9, $10, $11"+

			");",
		config.TablePrefix)

	var saveCount = 0
	//TODO: make bulk request
	for _, row := range rows {
		if _, err := pgsql.Exec(
			query,
			row.ReportAt,
			row.CampaignId,
			row.ProviderName,
			row.OperatorCode,
			row.LpHits,
			row.LpMsisdnHits,
			row.Mo,
			row.MoUniq,
			row.MoSuccess,
			row.RetrySuccess,
			row.Pixels,
		); err != nil {
			fmt.Println(err.Error())
		}
		saveCount += 1
	}

	log.Info("Reports save: ", saveCount)
	return nil
}

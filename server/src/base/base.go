package base

import (
	"database/sql"
	"fmt"

	"github.com/vostrok/utils/db"
)

var pgsql *sql.DB
var config db.DataBaseConfig

type Aggregate struct {
	ReportDate   int64  `json:"report_date,omitempty"`
	CampaignId   int64  `json:"id_campaign,omitempty"`
	Provider     string `json:"id_provider,omitempty"`
	OperatorCode int64  `json:"operator_code,omitempty"`
	LPHits       int64  `json:"lp_hits,omitempty"`
	LPMsisdnHits int64  `json:"lp_msisdn_hits,omitempty"`
	Mo           int64  `json:"mo,omitempty"`
	MoUniq       int64  `json:"mo_uniq,omitempty"`
	MoSuccess    int64  `json:"mo_success,omitempty"`
	RetrySuccess int64  `json:"retry_success,omitempty"`
	Pixels       int64  `json:"pixels_sent,omitempty"`
}

func Init(dbConfig db.DataBaseConfig) {
	config = dbConfig
	pgsql = db.Init(config)
}

func SaveRows(rows []Aggregate) error {
	var query string = fmt.Sprintf(
		"INSERT INTO %sreports ("+

			"report_date, "+
			"id_campaign, "+
			"id_provider, "+
			"id_operator, "+
			"lp_hits, "+
			"lp_msisdn_hits, "+
			"mo, "+
			"mo_uniq, "+
			"mo_success, "+
			"retry_success, "+
			"pixels"+

			") VALUES ("+

			"to_timestamp($1), $2, $3, $4, $5, $6, $7, $8, $9, $10, $11"+

			");",
		config.TablePrefix)

	//TODO: make bulk request
	for _, row := range rows {
		if _, err := pgsql.Exec(query,
			row.ReportDate,
			row.CampaignId,
			row.Provider,
			row.OperatorCode,
			row.LPHits,
			row.LPMsisdnHits,
			row.Mo,
			row.MoUniq,
			row.MoSuccess,
			row.RetrySuccess,
			row.Pixels,
		); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

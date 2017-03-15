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
	Campaign     int32  `json:"id_campaign,omitempty"`
	Provider     string `json:"id_provider,omitempty"`
	Operator     int32  `json:"id_operator,omitempty"`
	LPHits       int32  `json:"total_lp_hits,omitempty"`
	LPMsisdnHits int32  `json:"total_lp_msisdn_hits,omitempty"`
	Mo           int32  `json:"total_mo,omitempty"`
	MoUniq       int32  `json:"total_mo_uniq,omitempty"`
	MoSuccess    int32  `json:"total_mo_success_charge,omitempty"`
	Pixels       int32  `json:"total_pixels_sent,omitempty"`
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
			"pixels"+

			") VALUES ("+

			"to_timestamp($1), $2, $3, $4, $5, $6, $7, $8, $9, $10"+

			");",
		config.TablePrefix)

	//TODO: make bulk request
	for _, row := range rows {
		if _, err := pgsql.Exec(query, row.ReportDate, row.Campaign, row.Provider, row.Operator, row.LPHits, row.LPMsisdnHits, row.Mo, row.MoUniq, row.MoSuccess, row.Pixels); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

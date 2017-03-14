package base

import (
	"database/sql"
	"fmt"
	"github.com/vostrok/utils/db"
)

var pgsql *sql.DB
var config db.DataBaseConfig

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

func Init(dbConfig db.DataBaseConfig) {
	config = dbConfig
	pgsql = db.Init(config)
}

func SaveRows(rows []Aggregate) error {
	var query string = fmt.Sprintf("INSERT INTO %stest (time, cnt) values ($1, $2);", config.TablePrefix)

	//TODO: make bulk request
	for _, row := range rows {
		if _, err := pgsql.Exec(query, row.ReportDate, 1); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

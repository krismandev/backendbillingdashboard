package datastruct

import (
	"backendbillingdashboard/core"
)

// LoginRequest is use for clients login
type ReconRequest struct {
	ListCompanyID []string `json:"list_company_id"`
	CompanyID     string   `json:"company_id"`
	CompanyName   string   `json:"company_name"`
	MonthUse      string   `json:"month_use"`

	Param core.DataTableParam `json:"param"`
}

type ReconDataStruct struct {
	CompanyID      string             `json:"company_id"`
	CompanyName    string             `json:"company_name"`
	MonthUse       string             `json:"month_use"`
	ListServerData []ServerDataStruct `json:"list_server_data"`
}

type ServerDataStruct struct {
	ServerID             string `json:"server_id"`
	ServerName           string `json:"server_name"`
	AccountID            string `json:"account_id"`
	AccountName          string `json:"account_name"`
	ExternalOperatorCode string `json:"external_operator_code"`
	ExternalRoute        string `json:"external_route"`
	SmsCount             string `json:"smscount"`
}

package datastruct

import (
	"backendbillingdashboard/core"
)

type ServerAccountRequest struct {
	ServerID          string                    `json:"server_id"`
	AccountID         string                    `json:"account_id"`
	ExternalAccountID string                    `json:"external_account_id"`
	ListAccountID     []string                  `json:"list_account_id"`
	ListServerAccount []ServerAccountDataStruct `json:"list_server_account"`
	ListServerID      []string                  `json:"list_server_id"`
	Param             core.DataTableParam       `json:"param"`
}

type ServerAccountDataStruct struct {
	ServerID           string `json:"server_id"`
	AccountID          string `json:"account_id"`
	ExternalAccountID  string `json:"external_account_id"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
}

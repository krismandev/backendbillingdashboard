package datastruct

import (
	"backendbillingdashboard/core"
)

// LoginRequest is use for clients login
type ServerRequest struct {
	ListServerID []string            `json:"list_server_id"`
	ServerID     string              `json:"server_id"`
	ServerName   string              `json:"server_name"`
	ServerUrl    string              `json:"server_url"`
	Param        core.DataTableParam `json:"param"`
}

type ServerDataStruct struct {
	ServerID   string `json:"server_id"`
	ServerName string `json:"server_name"`
	ServerUrl  string `json:"server_url"`
}

type ServerAccountRequest struct {
	ServerID      string   `json:"server_id"`
	AccountID     string   `json:"account_id"`
	ServerAccount string   `json:"server_account"`
	ListAccountID []string `json:"list_account_id"`
	ListServerID  []string `json:"list_server_id"`

	Param core.DataTableParam `json:"param"`
}

type ServerAccountDataStruct struct {
	ServerID           string   `json:"server_id"`
	AccountID          string   `json:"account_id"`
	ExternalAccountID  string   `json:"external_account_id"`
	ListAccountID      []string `json:"list_account_id"`
	LastUpdateUsername string   `json:"last_update_username"`
	LastUpdateDate     string   `json:"last_update_date"`
}

package datastruct

import (
	"backendbillingdashboard/core"
	dtCompany "backendbillingdashboard/modules/company/datastruct"
)

//LoginRequest is use for clients login
type ItemPriceRequest struct {
	ListItemPriceID    []string              `json:"list_item-price_id"`
	ListItemPrice      []ItemPriceDataStruct `json:"list_item_price"`
	ItemID             string                `json:"item_id"`
	CompanyID          string                `json:"company_id"`
	ServerID           string                `json:"server_id"`
	Price              string                `json:"price"`
	Category           string                `json:"category"`
	LastUpdateUsername string                `json:"last_update_username"`

	ListCompanyID []string            `json:"list_company_id"`
	ListServerID  []string            `json:"list_server_id"`
	ListItemID    []string            `json:"list_item_id"`
	Param         core.DataTableParam `json:"param"`
}

type ItemPriceDataStruct struct {
	ItemID             string `json:"item_id"`
	CompanyID          string `json:"company_id"`
	ServerID           string `json:"server_id"`
	Price              string `json:"price"`
	Category           string `json:"category"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`

	Company dtCompany.CompanyDataStruct `json:"company"`
	Server  ServerDataStruct            `json:"server"`
	Item    ItemDataStruct              `json:"item"`
}

type ServerDataStruct struct {
	ServerID   string `json:"server_id"`
	ServerName string `json:"server_name"`
	ServerUrl  string `json:"server_url"`
}

type ItemDataStruct struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
	Operator string `json:"operator"`
	Route    string `json:"route"`
	Category string `json:"category"`
	UOM      string `json:"uom"`
}

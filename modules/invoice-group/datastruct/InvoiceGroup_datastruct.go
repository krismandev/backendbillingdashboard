package datastruct

import (
	"backendbillingdashboard/core"
)

//LoginRequest is use for clients login
type InvoiceGroupRequest struct {
	ListGroupID        []string                       `json:"list_group_id"`
	GroupID            string                         `json:"group_id"`
	GroupName          string                         `json:"group_name"`
	CompanyID          string                         `json:"company_id"`
	InvoiceTypeID      string                         `json:"invoice_type_id"`
	LastUpdateUsername string                         `json:"last_update_username"`
	InvoiceGroupDetail []InvoiceGroupDetailDataStruct `json:"invoice_group_detail"`
	Param              core.DataTableParam            `json:"param"`
}

type InvoiceGroupDataStruct struct {
	GroupID            string                         `json:"group_id"`
	GroupName          string                         `json:"group_name"`
	CompanyID          string                         `json:"company_id"`
	InvoiceTypeID      string                         `json:"invoice_type_id"`
	InvoiceGroupDetail []InvoiceGroupDetailDataStruct `json:"invoice_group_detail"`
}

type InvoiceGroupDetailDataStruct struct {
	GroupID  string `json:"group_id"`
	Identity string `json:"identity"`
	Type     string `json:"type"`
}

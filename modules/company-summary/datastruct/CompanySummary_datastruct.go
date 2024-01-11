package datastruct

import (
	"backendbillingdashboard/core"
	dtCompany "backendbillingdashboard/modules/company/datastruct"
	dtInvoice "backendbillingdashboard/modules/invoice/datastruct"
)

// LoginRequest is use for clients login
type CompanySummaryRequest struct {
	MonthUse  string              `json:"month_use"`
	CompanyID string              `json:"company_id"`
	Param     core.DataTableParam `json:"param"`
}

type CompanySummaryDataStruct struct {
	Company     dtCompany.CompanyDataStruct   `json:"company"`
	ListInvoice []dtInvoice.InvoiceDataStruct `json:"list_invoice"`
}

type CompanyDataStruct struct {
	dtCompany.CompanyDataStruct
	ListInvoice []InvoiceDataStruct `json:"list_invoice"`
}

type InvoiceDataStruct struct {
	dtInvoice.InvoiceDataStruct
}

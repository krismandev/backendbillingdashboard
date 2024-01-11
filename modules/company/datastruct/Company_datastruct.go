package datastruct

import (
	"backendbillingdashboard/core"
)

// LoginRequest is use for clients login
type CompanyRequest struct {
	ListCompanyID        []string            `json:"list_company_id"`
	CompanyID            string              `json:"company_id"`
	Name                 string              `json:"name"`
	Status               string              `json:"status"`
	Address1             string              `json:"address1"`
	Address2             string              `json:"address2"`
	City                 string              `json:"city"`
	Country              string              `json:"country"`
	ContactPerson        string              `json:"contact_person"`
	ContactPersonPhone   string              `json:"contact_person_phone"`
	Phone                string              `json:"phone"`
	Fax                  string              `json:"fax"`
	Desc                 string              `json:"desc"`
	TermOfPayment        string              `json:"term_of_payment"`
	DefaultInvoiceTypeID string              `json:"default_invoice_type_id"`
	LastUpdateUsername   string              `json:"last_update_username"`
	Param                core.DataTableParam `json:"param"`
}

type CompanyDataStruct struct {
	CompanyID            string `json:"company_id"`
	Name                 string `json:"name"`
	Status               string `json:"status"`
	Address1             string `json:"address1"`
	Address2             string `json:"address2"`
	City                 string `json:"city"`
	Country              string `json:"country"`
	ContactPerson        string `json:"contact_person"`
	ContactPersonPhone   string `json:"contact_person_phone"`
	Phone                string `json:"phone"`
	Fax                  string `json:"fax"`
	Desc                 string `json:"desc"`
	TermOfPayment        string `json:"term_of_payment"`
	DefaultInvoiceTypeID string `json:"default_invoice_type_id"`
	LastUpdateUsername   string `json:"last_update_username"`
	LastUpdateDate       string `json:"last_update_date"`
}

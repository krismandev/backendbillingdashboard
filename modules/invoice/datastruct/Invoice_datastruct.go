package datastruct

import (
	"backendbillingdashboard/core"
	dtAccount "backendbillingdashboard/modules/account/datastruct"
	dtCompany "backendbillingdashboard/modules/company/datastruct"
	dtServer "backendbillingdashboard/modules/server/datastruct"
)

// LoginRequest is use for clients login
type InvoiceRequest struct {
	ListInvoiceID                  []string `json:"list_invoiceid"`
	InvoiceID                      string   `json:"invoice_id"`
	InvoiceNo                      string   `json:"invoice_no"`
	InvoiceDate                    string   `json:"invoice_date"`
	InvoiceStatus                  string   `json:"invoice_status"`
	CompanyID                      string   `json:"company_id"`
	MonthUse                       string   `json:"month_use"`
	InvoiceTypeID                  string   `json:"invoice_type_id"`
	PrintCounter                   string   `json:"print_counter"`
	Note                           string   `json:"note"`
	CancelDesc                     string   `json:"cancel_desc"`
	LastPrintUsername              string   `json:"last_print_username"`
	LastPrintDate                  string   `json:"last_print_date"`
	CreatedAt                      string   `json:"created_at"`
	CreatedBy                      string   `json:"created_by"`
	LastUpdateUsername             string   `json:"last_update_username"`
	LastUpdateDate                 string   `json:"last_update_date"`
	DiscountType                   string   `json:"discount_type"`
	Discount                       string   `json:"discount"`
	PPN                            string   `json:"ppn"`
	Paid                           string   `json:"paid"`
	PaymentMethod                  string   `json:"payment_method"`
	ExchangeRateDate               string   `json:"exchange_rate_date"`
	DueDate                        string   `json:"due_date"`
	ApproachingDueDateInterval     string   `json:"approaching_due_date_interval"`
	GrandTotal                     string   `json:"grand_total"`
	PPNAmount                      string   `json:"ppn_amount"`
	ListServerID                   []string `json:"list_server_id"`
	Attachment                     string   `json:"attachment"`
	AdjustmentNote                 string   `json:"adjustment_note"`
	InvoiceDateRange               []string `json:"invoice_date_range"`
	TaxInvoice                     string   `json:"tax_invoice"`
	AdditionalParamOperator        []string `json:"additional_param_operator"`
	AdditionalParamRoute           []string `json:"additional_param_route"`
	AdjustmentConfirmationUsername string   `json:"adjustment_confirmation_username"`
	AdjustmentConfirmationDate     string   `json:"adjustment_confirmation_date"`
	ReceivedDate                   string   `json:"received_date"`
	ReceiptLetterAttachment        string   `json:"receipt_letter_attachment"`

	ServerID     string `json:"server_id"`
	Sender       string `json:"sender"`
	BatchID      string `json:"batch_id"`
	CurrencyCode string `json:"currency_code"`
	//company name
	Name                          string   `json:"name"`
	ListExternalRootParentAccount []string `json:"list_external_rootparent_account"`
	ListCompanyID                 []string `json:"list_company_id"`
	ListAccountID                 []string `json:"list_account_id"`

	ListInvoiceDetail []InvoiceDetailStruct `json:"list_invoice_detail"`
	OldData           InvoiceDataStruct     `json:"old_data"`

	Param core.DataTableParam `json:"param"`
}

type InvoiceDataStruct struct {
	InvoiceID               string `json:"invoice_id"`
	InvoiceNo               string `json:"invoice_no"`
	InvoiceDate             string `json:"invoice_date"`
	InvoiceStatus           string `json:"invoice_status"`
	CompanyID               string `json:"company_id"`
	MonthUse                string `json:"month_use"`
	InvoiceTypeID           string `json:"invoice_type_id"`
	PrintCounter            string `json:"print_counter"`
	Note                    string `json:"note"`
	CancelDesc              string `json:"cancel_desc"`
	LastPrintUsername       string `json:"last_print_username"`
	LastPrintDate           string `json:"last_print_date"`
	CreatedAt               string `json:"created_at"`
	CreatedBy               string `json:"created_by"`
	LastUpdateUsername      string `json:"last_update_username"`
	LastUpdateDate          string `json:"last_update_date"`
	DiscountType            string `json:"discount_type"`
	Discount                string `json:"discount"`
	PPN                     string `json:"ppn"`
	Paid                    string `json:"paid"`
	PaymentMethod           string `json:"payment_method"`
	ExchangeRateDate        string `json:"exchange_rate_date"`
	DueDate                 string `json:"due_date"`
	GrandTotal              string `json:"grand_total"`
	PPNAmount               string `json:"ppn_amount"`
	Sender                  string `json:"sender"`
	BatchID                 string `json:"batch_id"`
	Attachment              string `json:"attachment"`
	AdjustmentNote          string `json:"adjustment_note"`
	TaxInvoice              string `json:"tax_invoice"`
	ReceivedDate            string `json:"received_date"`
	ReceiptLetterAttachment string `json:"receipt_letter_attachment"`

	InvoiceType InvoiceTypeDataStruct       `json:"invoice_type"`
	Company     dtCompany.CompanyDataStruct `json:"company"`

	ListInvoiceDetail []InvoiceDetailStruct `json:"list_invoice_detail"`
}

type InvoiceDetailStruct struct {
	InvoiceDetailID                string `json:"invoice_detail_id"`
	AccountID                      string `json:"account_id"`
	InvoiceID                      string `json:"invoice_id"`
	ItemID                         string `json:"item_id"`
	Qty                            string `json:"qty"`
	Adjustment                     string `json:"adjustment"`
	AdjustmentConfirmationUsername string `json:"adjustment_confirmation_username"`
	AdjustmentConfirmationDate     string `json:"adjustment_confirmation_date"`
	Uom                            string `json:"uom"`
	ItemPrice                      string `json:"item_price"`
	Note                           string `json:"note"`
	BalanceType                    string `json:"balance_type"`
	ServerID                       string `json:"server_id"`
	ExternalUserID                 string `json:"external_user_id"`
	ExternalUsername               string `json:"external_username"`
	ExternalSender                 string `json:"external_sender"`
	LastUpdateUsername             string `json:"last_update_username"`

	Item    ItemDataStruct              `json:"item"`
	Server  dtServer.ServerDataStruct   `json:"server"`
	Account dtAccount.AccountDataStruct `json:"account"`
}

type ItemDataStruct struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
	Operator string `json:"operator"`
	Route    string `json:"route"`
	Category string `json:"category"`
	UOM      string `json:"uom"`
}

type InvoiceTypeDataStruct struct {
	InvoiceTypeID      string `json:"invoice_type_id"`
	InvoiceTypeName    string `json:"invoice_type_name"`
	ServerID           string `json:"server_id"`
	Category           string `json:"category"`
	LoadFromServer     string `json:"load_from_server"`
	IsGroup            string `json:"is_group"`
	GroupType          string `json:"group_type"`
	LastUpdateUsername string `json:"last_update_username"`
	LastUpdateDate     string `json:"last_update_date"`
	CurrencyCode       string `json:"currency_code"`
}

type AccountDataStruct struct {
	AccountID          string `json:"account_id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	CompanyID          string `json:"company_id"`
	AccountType        string `json:"account_type"`
	BillingType        string `json:"billing_type"`
	Desc               string `json:"desc"`
	Address1           string `json:"address1"`
	Address2           string `json:"address2"`
	City               string `json:"city"`
	Phone              string `json:"phone"`
	ContactPerson      string `json:"contact_person"`
	ContactPersonPhone string `json:"contact_person_phone"`
	LastUpdateUsername string `json:"last_update_username"`
}

type GenerateInvoiceDataStruct struct {
	ID                 string `json:"id"`
	GenerateTime       string `json:"generate_time"`
	File               string `json:"file"`
	Progress           string `json:"progress"`
	LastUpdateUsername string `json:"last_update_username"`
}

type InquiryPaymentRequest struct {
	CompanyID   string              `json:"company_id"`
	CompanyName string              `json:"company_name"`
	MonthUse    string              `json:"month_use"`
	Param       core.DataTableParam `json:"param"`
}

type InquiryPaymentDataStruct struct {
	CompanyID             string                       `json:"company_id"`
	CompanyName           string                       `json:"company_name"`
	AccountID             string                       `json:"account_id"`
	AccountName           string                       `json:"account_name"`
	Amount                string                       `json:"amount"`
	InvoiceID             string                       `json:"invoice_id"`
	Invoicing             string                       `json:"invoicing"`
	ProformaInvoiceAmount string                       `json:"proforma_invoice_amount"`
	GrandTotal            string                       `json:"grand_total"`
	PPN                   string                       `json:"ppn"`
	OutStanding           string                       `json:"outstanding"`
	InquiryPaymentDetail  []InquiryPaymentDetailStruct `json:"inquiry_payment_detail"`
}

type InquiryPaymentDetailStruct struct {
	CompanyID   string `json:"company_id"`
	CompanyName string `json:"company_name"`
	AccountID   string `json:"account_id"`
	AccountName string `json:"account_name"`
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	ItemName    string `json:"item_name"`
	ItemID      string `json:"item_id"`
	Sender      string `json:"sender"`
	Amount      string `json:"amount"`
	InvoiceID   string `json:"invoice_id"`
	InvoiceNo   string `json:"invoice_no"`
}

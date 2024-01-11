package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/company-summary/datastruct"
	dtCompany "backendbillingdashboard/modules/company/datastruct"
	companyModel "backendbillingdashboard/modules/company/models"
	companyProcessor "backendbillingdashboard/modules/company/processors"
	dtInvoice "backendbillingdashboard/modules/invoice/datastruct"
	invoiceModel "backendbillingdashboard/modules/invoice/models"
	invoiceProcessor "backendbillingdashboard/modules/invoice/processors"
)

func GetListCompanySummary(conn *connections.Connections, req datastruct.CompanySummaryRequest) ([]datastruct.CompanySummaryDataStruct, error) {
	var output []datastruct.CompanySummaryDataStruct
	var err error

	// grab mapping data from model
	var comanyRequest dtCompany.CompanyRequest
	comanyRequest.Param = req.Param
	comanyRequest.CompanyID = req.CompanyID
	companyList, err := companyModel.GetCompanyFromRequest(conn, comanyRequest)
	if err != nil {
		return output, err
	}

	var listCompanyID []string
	for _, each := range companyList {
		listCompanyID = append(listCompanyID, each["company_id"])
	}

	var invoiceRequest dtInvoice.InvoiceRequest
	invoiceRequest.MonthUse = req.MonthUse
	invoiceRequest.ListCompanyID = listCompanyID
	invoiceRequest.Param.PerPage = 9999
	invoiceList, err := invoiceModel.GetListInvoice(conn, invoiceRequest)

	var listInvoiceID []string
	var listInvoice []dtInvoice.InvoiceDataStruct
	for _, each := range invoiceList {
		listInvoiceID = append(listInvoiceID, each["invoice_id"].(string))
	}

	invoiceDetailList, err := invoiceModel.GetListInvoiceDetail(conn, listInvoiceID)

	for _, each := range invoiceList {
		invoice := invoiceProcessor.CreateInvoiceStruct(each, invoiceDetailList)

		listInvoice = append(listInvoice, invoice)
	}

	for _, company := range companyList {
		// single := CreateSingleCompanySummaryStruct(companysummary)

		company := companyProcessor.CreateSingleCompanyStruct(company)

		var out datastruct.CompanySummaryDataStruct
		out.Company = company
		var companyInvoices []dtInvoice.InvoiceDataStruct
		for _, each := range listInvoice {
			if each.CompanyID == company.CompanyID {
				companyInvoices = append(companyInvoices, each)
			}
		}
		out.ListInvoice = companyInvoices
		output = append(output, out)
	}

	return output, err
}

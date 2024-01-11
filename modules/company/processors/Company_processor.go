package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/company/datastruct"
	"backendbillingdashboard/modules/company/models"
)

func GetListCompany(conn *connections.Connections, req datastruct.CompanyRequest) ([]datastruct.CompanyDataStruct, error) {
	var output []datastruct.CompanyDataStruct
	var err error

	// grab mapping data from model
	companyList, err := models.GetCompanyFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, company := range companyList {
		single := CreateSingleCompanyStruct(company)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleCompanyStruct(company map[string]string) datastruct.CompanyDataStruct {
	var single datastruct.CompanyDataStruct
	single.CompanyID = company["company_id"]
	single.Name = company["name"]
	single.Status = company["status"]
	single.Address1 = company["address1"]
	single.Address2 = company["address2"]
	single.City = company["city"]
	single.Country = company["country"]
	single.ContactPerson = company["contact_person"]
	single.ContactPersonPhone = company["contact_person_phone"]
	single.Phone = company["phone"]
	single.Fax = company["fax"]
	single.Desc = company["desc"]
	single.DefaultInvoiceTypeID = company["default_invoice_type_id"]
	single.LastUpdateUsername = company["last_update_username"]
	single.LastUpdateDate = company["last_update_date"]
	single.TermOfPayment = company["term_of_payment"]

	return single
}

func InsertCompany(conn *connections.Connections, req datastruct.CompanyRequest) (datastruct.CompanyDataStruct, error) {
	var output datastruct.CompanyDataStruct
	var err error

	err = models.InsertCompany(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSingleCompany(req.CompanyID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleCompanyStruct(single)
	return output, err
}

func UpdateCompany(conn *connections.Connections, req datastruct.CompanyRequest) (datastruct.CompanyDataStruct, error) {
	var output datastruct.CompanyDataStruct
	var err error

	err = models.UpdateCompany(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSingleCompany(req.CompanyID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleCompanyStruct(single)
	return output, err
}

func DeleteCompany(conn *connections.Connections, req datastruct.CompanyRequest) error {
	err := models.DeleteCompany(conn, req)
	return err
}

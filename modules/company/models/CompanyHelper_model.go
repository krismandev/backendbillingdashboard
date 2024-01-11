package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/company/datastruct"
	"errors"
	"strconv"
)

func GetSingleCompany(companyID string, conn *connections.Connections, req datastruct.CompanyRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	if len(companyID) == 0 {
		companyID = req.CompanyID
	}
	query := "SELECT company_id, name, status, address1, address2, city, country, contact_person, contact_person_phone, phone, fax, company.desc, last_update_username, last_update_date, term_of_payment FROM company WHERE company_id = ?"
	results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, companyID)
	if err != nil {
		return result, err
	}

	// convert from []map[string]string to single map[string]string
	for _, res := range results {
		result = res
		break
	}
	return result, err
}

func GetCompanies(conn *connections.Connections, listCompanyID []string) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	if len(listCompanyID) > 0 {
		var baseIn string
		for _, prid := range listCompanyID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "company_id IN ("+baseIn+")")
	}

	runQuery := "SELECT company_id, name, status, address1, address2, city, country, contact_person, contact_person_phone, phone, fax, company.desc, default_invoice_type_id,last_update_username, last_update_date, term_of_payment FROM company "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	// lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	// lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func CheckCompanyExists(stubID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, stubID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("Company ID is not exists")
	// }
	return nil
}

func CheckCompanyDuplicate(exceptID string, conn *connections.Connections, req datastruct.CompanyRequest) error {
	var param []interface{}
	qry := "SELECT COUNT(company_id) FROM company WHERE company_id = ?"
	param = append(param, req.CompanyID)
	if len(exceptID) > 0 {
		qry += " AND company_id <> ?"
		param = append(param, exceptID)
	}

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	datacount, _ := strconv.Atoi(cnt)
	if datacount > 0 {
		return errors.New("Company ID is already exists. Please use another Company ID")
	}
	return nil
}

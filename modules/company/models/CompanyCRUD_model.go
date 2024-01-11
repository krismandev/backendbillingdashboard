package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/company/datastruct"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetCompanyFromRequest(conn *connections.Connections, req datastruct.CompanyRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "company_id = ?", req.CompanyID)
	lib.AppendWhereLike(&baseWhere, &baseParam, "name like ? ", req.Name)
	lib.AppendWhere(&baseWhere, &baseParam, "status = ?", req.Status)
	lib.AppendWhere(&baseWhere, &baseParam, "city = ?", req.City)
	if len(req.ListCompanyID) > 0 {
		var baseIn string
		for _, prid := range req.ListCompanyID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "company_id IN ("+baseIn+")")
	}

	runQuery := "SELECT company_id, name, status, address1, address2, city, country, contact_person, contact_person_phone, phone, fax, company.desc, default_invoice_type_id,last_update_username, last_update_date, term_of_payment FROM company "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertCompany(conn *connections.Connections, req datastruct.CompanyRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT max(company_id) FROM company ")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	log.Info("HasilID", lastId)

	insertIdString := strconv.Itoa(insertId)
	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Name)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Status)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Address1)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Address2)
	lib.AppendComma(&baseIn, &baseParam, "?", req.City)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Country)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ContactPerson)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ContactPersonPhone)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Phone)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Fax)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Desc)
	lib.AppendComma(&baseIn, &baseParam, "?", req.TermOfPayment)
	lib.AppendComma(&baseIn, &baseParam, "?", req.LastUpdateUsername)

	qry := "INSERT INTO company (company_id, name, status, addr1, addr2, city, country, contactperson, contactpersonphone, phone, fax, company.desc,term_of_payment, last_update_username) VALUES (" + baseIn + ")"
	_, err = conn.DBAppConn.InsertGetLastID(qry, baseParam...)
	_, _, err = conn.DBAppConn.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "company")
	return err
}

func UpdateCompany(conn *connections.Connections, req datastruct.CompanyRequest) error {
	var err error

	var baseUp string
	var baseParam []interface{}

	lib.AppendComma(&baseUp, &baseParam, "name = ?", req.Name)
	lib.AppendComma(&baseUp, &baseParam, "status = ?", req.Status)
	lib.AppendComma(&baseUp, &baseParam, "address1 = ?", req.Address1)
	lib.AppendComma(&baseUp, &baseParam, "address2 = ?", req.Address2)
	lib.AppendComma(&baseUp, &baseParam, "city = ?", req.City)
	lib.AppendComma(&baseUp, &baseParam, "country = ?", req.Country)
	lib.AppendComma(&baseUp, &baseParam, "contact_person = ?", req.ContactPerson)
	lib.AppendComma(&baseUp, &baseParam, "contact_person_phone = ?", req.ContactPersonPhone)
	lib.AppendComma(&baseUp, &baseParam, "phone = ?", req.Phone)
	lib.AppendComma(&baseUp, &baseParam, "fax = ?", req.Fax)
	lib.AppendComma(&baseUp, &baseParam, "company.desc = ?", req.Desc)
	lib.AppendComma(&baseUp, &baseParam, "company.term_of_payment = ?", req.TermOfPayment)
	lib.AppendComma(&baseUp, &baseParam, "last_update_username = ?", req.LastUpdateUsername)
	qry := "UPDATE company SET " + baseUp + " WHERE company_id = ?"
	baseParam = append(baseParam, req.CompanyID)
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteCompany(conn *connections.Connections, req datastruct.CompanyRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "DELETE FROM company WHERE company_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.CompanyID)
	return err
}

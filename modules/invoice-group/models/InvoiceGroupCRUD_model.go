package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/invoice-group/datastruct"
	"strconv"
	"strings"
)

func GetInvoiceGroupFromRequest(conn *connections.Connections, req datastruct.InvoiceGroupRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "group_id = ?", req.GroupID)
	lib.AppendWhere(&baseWhere, &baseParam, "company_id = ?", req.CompanyID)
	lib.AppendWhere(&baseWhere, &baseParam, "group_name = ?", req.GroupName)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice_type_id = ?", req.InvoiceTypeID)
	if len(req.ListGroupID) > 0 {
		var baseIn string
		for _, prid := range req.ListGroupID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "group_id IN ("+baseIn+")")
	}

	runQuery := `SELECT company_invoice_group.group_id, group_name, company_id, invoice_type_id FROM company_invoice_group `
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	resultQry, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)

	for _, each := range resultQry {
		single := make(map[string]interface{})
		single["group_id"] = each["group_id"]
		single["group_name"] = each["group_name"]
		single["company_id"] = each["company_id"]
		single["invoice_type_id"] = each["invoice_type_id"]

		runQryGetDetail := `SELECT group_id, identity, type from company_invoice_group_detail WHERE group_id = ? `
		resultGetDetail, _, errGetDetail := conn.DBAppConn.SelectQueryByFieldNameSlice(runQryGetDetail, each["group_id"])
		if errGetDetail != nil {
			return nil, errGetDetail
		}
		var details []map[string]interface{}

		for _, dt := range resultGetDetail {
			detail := make(map[string]interface{})
			detail["group_id"] = dt["group_id"]
			detail["identity"] = dt["identity"]
			detail["type"] = dt["type"]
			details = append(details, detail)
		}
		single["company_invoice_group_detail"] = details

		result = append(result, single)
	}

	return result, err
}

func InsertInvoiceGroup(conn *connections.Connections, req datastruct.InvoiceGroupRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "company_invioce_group")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	lib.AppendComma(&baseIn, &baseParam, "?", req.GroupName)
	lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceTypeID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.CompanyID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.LastUpdateUsername)
	lib.AppendCommaRaw(&baseIn, "now()")

	qry := "INSERT INTO company_invoice_group (group_id,group_name,invoice_type_id,company_id, last_update_username, last_update_date) VALUES (" + baseIn + ")"
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	if err != nil {
		return err
	}

	var paramsInsertDetail []interface{}
	var stringGroup []string
	qryInsertDetails := "INSERT INTO company_invoice_group_detail (group_id,identity,type) VALUES"
	for _, each := range req.InvoiceGroupDetail {
		partquery := "(?, ?, ?)"
		paramsInsertDetail = append(paramsInsertDetail, insertIdString)
		paramsInsertDetail = append(paramsInsertDetail, each.Identity)
		paramsInsertDetail = append(paramsInsertDetail, each.Type)
		stringGroup = append(stringGroup, partquery)
	}

	qryInsertDetails = qryInsertDetails + strings.Join(stringGroup, ", ")
	_, _, err = conn.DBAppConn.Exec(qryInsertDetails, paramsInsertDetail...)
	if err != nil {
		return err
	}

	_, _, err = conn.DBAppConn.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "company_invioce_group")

	return err
}

func UpdateInvoiceGroup(conn *connections.Connections, req datastruct.InvoiceGroupRequest) error {
	var err error

	var baseUp string
	var baseParam []interface{}

	lib.AppendComma(&baseUp, &baseParam, "group_name = ?", req.GroupName)
	lib.AppendComma(&baseUp, &baseParam, "company_id = ?", req.CompanyID)
	lib.AppendComma(&baseUp, &baseParam, "invoice_type_id = ?", req.InvoiceTypeID)
	qry := "UPDATE company_invoice_group SET " + baseUp + " WHERE group_id = ?"
	baseParam = append(baseParam, req.GroupID)
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	if err != nil {
		return err
	}

	_, _, err = conn.DBAppConn.Exec("DELETE FROM company_invoice_group_detail WHERE group_id = ?", req.GroupID)

	var paramsInsertDetail []interface{}
	var stringGroup []string
	qryInsertDetails := "INSERT INTO company_invoice_group_detail (group_id,identity,type) VALUES"
	for _, each := range req.InvoiceGroupDetail {
		partquery := "(?, ?,?)"
		paramsInsertDetail = append(paramsInsertDetail, req.GroupID)
		paramsInsertDetail = append(paramsInsertDetail, each.Identity)
		paramsInsertDetail = append(paramsInsertDetail, each.Type)
		stringGroup = append(stringGroup, partquery)
	}
	qryInsertDetails = qryInsertDetails + strings.Join(stringGroup, ", ")
	_, _, err = conn.DBAppConn.Exec(qryInsertDetails, paramsInsertDetail...)
	return err
}

func DeleteInvoiceGroup(conn *connections.Connections, req datastruct.InvoiceGroupRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	_, _, err = conn.DBAppConn.Exec("DELETE FROM company_invoice_group WHERE group_id = ?", req.GroupID)

	_, _, err = conn.DBAppConn.Exec("DELETE FROM company_invoice_group_detail WHERE group_id = ? ", req.GroupID)

	return err
}

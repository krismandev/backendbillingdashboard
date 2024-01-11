package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/reconciliation/datastruct"
)

func GetReconFromRequest(conn *connections.Connections, req datastruct.ReconRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "c.company_id = ?", req.CompanyID)
	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(external_transdate, '%Y%m') = ?", req.MonthUse)
	if len(req.ListCompanyID) > 0 {
		var baseIn string
		for _, prid := range req.ListCompanyID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "c.company_id IN ("+baseIn+")")
	}

	runQuery := `SELECT a.account_id,b.name as account_name, external_operatorcode, external_route,c.company_id,a.server_id,d.server_name, sum(external_smscount) as smscount 
	FROM server_data as a 
	LEFT JOIN account b ON a.account_id = b.account_id
	LEFT JOIN company c ON b.company_id = c.company_id 
	LEFT JOIN server d ON d.server_id = a.server_id 
	`
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}

	runQuery += "GROUP BY 1,2,3,4,5,6,7"
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertRecon(conn *connections.Connections, req datastruct.ReconRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	// var baseIn string
	// var baseParam []interface{}

	// lib.AppendComma(&baseIn, &baseParam, "?", req.ReconID)
	// lib.AppendComma(&baseIn, &baseParam, "?", req.ReconName)

	// qry := "INSERT INTO reconciliation (reconciliationid, reconciliationname) VALUES (" + baseIn + ")"
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	return err
}

func UpdateRecon(conn *connections.Connections, req datastruct.ReconRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	// var baseUp string
	// var baseParam []interface{}

	// lib.AppendComma(&baseUp, &baseParam, "reconciliationname = ?", req.ReconName)
	// qry := "UPDATE reconciliation SET " + baseUp + " WHERE reconciliationid = ?"
	// baseParam = append(baseParam, req.ReconID)
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteRecon(conn *connections.Connections, req datastruct.ReconRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	// qry := "DELETE FROM reconciliation WHERE reconciliationid = ?"
	// _, _, err = conn.DBAppConn.Exec(qry, req.ReconID)
	return err
}

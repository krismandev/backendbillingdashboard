package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/invoice-group/datastruct"
)

func GetSingleInvoiceGroup(groupID string, conn *connections.Connections, req datastruct.InvoiceGroupRequest) (map[string]interface{}, error) {
	var result []map[string]interface{}
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC

	resultQry, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice("SELECT company_invoice_group.group_id, group_name, company_id, invoice_type_id FROM company_invoice_group WHERE group_id = ? ", groupID)

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

	return result[0], err
}

func CheckInvoiceGroupExists(stubID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, stubID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("InvoiceGroup ID is not exists")
	// }
	return nil
}

func CheckInvoiceGroupDuplicate(exceptID string, conn *connections.Connections, req datastruct.InvoiceGroupRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(stubid) FROM stub WHERE stubid = ?"
	// param = append(param, req.InvoiceGroupID)
	// if len(exceptID) > 0 {
	// 	qry += " AND stubid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("InvoiceGroup ID is already exists. Please use another InvoiceGroup ID")
	// }
	return nil
}

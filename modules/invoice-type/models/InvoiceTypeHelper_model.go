package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/invoice-type/datastruct"
)

func GetSingleInvoiceType(invoiceTypeID string, conn *connections.Connections) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "inv_type_id = ?", invoiceTypeID)

	runQuery := "SELECT inv_type_id, inv_type_name, server_id, category, load_from_server,is_group,group_type, last_update_username, last_update_date, currency_code,is_group FROM invoice_type "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}

	results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	for _, each := range results {
		result = each
	}
	return result, err
}

func CheckInvoiceTypeExists(invoiceTypeID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(invoice-typeid) FROM invoice-type WHERE invoice-typeid = ?"
	// param = append(param, invoice-typeID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("InvoiceType ID is not exists")
	// }
	return nil
}

func CheckInvoiceTypeDuplicate(exceptID string, conn *connections.Connections, req datastruct.InvoiceTypeRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(invoice-typeid) FROM invoice-type WHERE invoice-typeid = ?"
	// param = append(param, req.InvoiceTypeID)
	// if len(exceptID) > 0 {
	// 	qry += " AND invoice-typeid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("InvoiceType ID is already exists. Please use another InvoiceType ID")
	// }
	return nil
}

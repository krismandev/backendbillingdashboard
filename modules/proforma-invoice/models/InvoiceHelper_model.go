package models

import (
	"backendbillingdashboard/config"
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/invoice/datastruct"
	"errors"
	"strconv"
	"time"
)

func GetSingleInvoice(invoiceID string, conn *connections.Connections, req datastruct.InvoiceRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(invoiceID) == 0 {
	// 	invoiceID = req.InvoiceID
	// }
	// query := "SELECT invoiceid, invoicename FROM invoice WHERE invoiceid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, invoiceID)
	// if err != nil {
	// 	return result, err
	// }

	// // convert from []map[string]string to single map[string]string
	// for _, res := range results {
	// 	result = res
	// 	break
	// }
	return result, err
}

func GetListInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var resultQuery []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "invoice_id = ?", req.InvoiceID)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.invoice_no = ?", req.InvoiceNo)
	lib.AppendWhere(&baseWhere, &baseParam, "invoicestatus = ?", "A")
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.company_id = ?", req.CompanyID)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.paid = ?", req.Paid)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.invoice_date = ?", req.InvoiceDate)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.inv_type_id = ?", req.InvoiceTypeID)
	lib.AppendWhereLike(&baseWhere, &baseParam, "company.name like ? ", req.Name)
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.due_date > curdate() and invoice.due_date <= curdate() + interval ? day", req.ApproachingDueDateInterval)
	lib.AppendWhere(&baseWhere, &baseParam, "month_use = ?", req.MonthUse)
	if len(req.ListInvoiceID) > 0 {
		var baseIn string
		for _, prid := range req.ListInvoiceID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "invoice_id IN ("+baseIn+")")
	}

	if len(req.ListCompanyID) > 0 {
		var baseIn string
		for _, prid := range req.ListCompanyID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "invoice.company_id IN ("+baseIn+")")
	}

	if len(req.InvoiceDateRange) == 2 {
		if len(req.InvoiceDateRange[0]) > 0 {
			tinput := req.InvoiceDateRange[0]
			t, _ := time.Parse("2006-01-02", tinput)
			tgl := t.Format("2006-01-02")
			if len(baseWhere) > 0 {
				baseWhere += " AND "
			}
			baseWhere += " invoice_date >= '" + tgl + "'"
		}
		if len(req.InvoiceDateRange[1]) > 0 {
			tinput := req.InvoiceDateRange[1]
			t, _ := time.Parse("2006-01-02", tinput)
			tgl := t.Format("2006-01-02")
			if len(baseWhere) > 0 {
				baseWhere += " AND "
			}
			baseWhere += " invoice_date <= '" + tgl + "'"
		}
	}

	runQuery := `SELECT invoice.invoice_id, invoice.invoice_no, invoice.invoice_date, invoice.invoicestatus, invoice.company_id, invoice.month_use, 
	invoice.inv_type_id, invoice.printcounter, invoice.note, invoice.canceldesc, invoice.payment_method ,invoice.attachment,invoice.last_print_username, 
	invoice.last_print_date, invoice.created_at, invoice.created_by, invoice.last_update_username, invoice.exchange_rate_date, invoice.grand_total, invoice.sender,
	invoice.ppn_amount, invoice.last_update_date, invoice.discount_type, invoice.discount, invoice.ppn ,invoice.paid, invoice.due_date ,invoice.received_date,invoice.receipt_letter_attachment, 
	invoice_type.inv_type_id as tblinvoice_type_inv_type_id, invoice_type.inv_type_name, invoice_type.server_id as tblinvoice_type_server_id, invoice_type.category, 
	invoice_type.load_from_server,invoice_type.is_group,invoice_type.group_type, invoice_type.currency_code, company.company_id as tblcompany_company_id, company.name as tblcompany_name, company.address1, company.address2, 
	company.city, company.phone, company.contact_person, company.contact_person_phone FROM proforma_invoice as invoice  JOIN invoice_type ON 
	invoice.inv_type_id = invoice_type.inv_type_id JOIN company ON invoice.company_id = company.company_id `
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	resultQuery, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	for _, each := range resultQuery {
		single := make(map[string]interface{})
		single["invoice_id"] = each["invoice_id"]
		single["invoice_no"] = each["invoice_no"]
		single["invoice_date"] = each["invoice_date"]
		single["invoicestatus"] = each["invoicestatus"]
		single["company_id"] = each["company_id"]
		// first4 := each["month_use"]

		// last2 := each["month_use"][len(each["month_use"])-2:]

		single["month_use"] = each["month_use"]
		single["inv_type_id"] = each["inv_type_id"]
		single["printcounter"] = each["printcounter"]
		single["note"] = each["note"]
		single["canceldesc"] = each["canceldesc"]
		single["ppn"] = each["ppn"]
		single["last_print_username"] = each["last_print_username"]
		single["attachment"] = each["attachment"]
		single["last_print_date"] = each["last_print_date"]
		single["created_at"] = each["created_at"]
		single["created_by"] = each["created_by"]
		single["last_update_username"] = each["last_update_username"]
		single["last_update_date"] = each["last_update_date"]
		single["discount_type"] = each["discount_type"]
		single["discount"] = each["discount"]
		single["paid"] = each["paid"]
		single["payment_method"] = each["payment_method"]
		single["exchange_rate_date"] = each["exchange_rate_date"]
		single["due_date"] = each["due_date"]
		single["grand_total"] = each["grand_total"]
		single["ppn_amount"] = each["ppn_amount"]
		single["sender"] = each["sender"]
		single["received_date"] = each["received_date"]
		single["receipt_letter_attachment"] = each["receipt_letter_attachment"]

		invType := make(map[string]interface{})
		invType["inv_type_id"] = each["tblinvoice_type_inv_type_id"]
		invType["inv_type_name"] = each["inv_type_name"]
		invType["server_id"] = each["tblinvoice_type_server_id"]
		invType["category"] = each["category"]
		invType["load_from_server"] = each["load_from_server"]
		invType["is_group"] = each["is_group"]
		invType["group_type"] = each["group_type"]
		invType["currency_code"] = each["currency_code"]

		single["invoice_type"] = invType

		company := make(map[string]interface{})
		company["company_id"] = each["tblcompany_company_id"]
		company["name"] = each["tblcompany_name"]
		company["address1"] = each["address1"]
		company["address2"] = each["address2"]
		company["city"] = each["city"]
		company["phone"] = each["phone"]
		company["contact_person"] = each["contact_person"]
		company["contact_person_phone"] = each["contact_person_phone"]

		single["company"] = company

		// qryGetDetail := `SELECT invoice_detail.invoice_detail_id, invoice_detail.account_id, invoice_detail.invoice_id, invoice_detail.itemid,
		// invoice_detail.qty, invoice_detail.uom, invoice_detail.item_price, tiering, invoice_detail.note, invoice_detail.balance_type,
		// invoice_detail.server_id , invoice_detail.external_sender,
		// item.item_id as tblitem_item_id, item.item_name, item.operator, item.route, item.category,
		// item.uom as tblitem_uom, server.server_id as tblserver_server_id, server.server_name as tblserver_server_name,external_user_id, external_username,
		// account.name as account_name, account.account_type, account.company_id, account.status, account.non_taxable
		// FROM invoice_detail
		// JOIN item ON invoice_detail.itemid = item.item_id
		// JOIN server ON invoice_detail.server_id = server.server_id
		// LEFT JOIN account ON invoice_detail.account_id = account.account_id
		// WHERE invoice_id = ?`
		// resultGetDetail, _, errGetDetail := conn.DBAppConn.SelectQueryByFieldNameSlice(qryGetDetail, single["invoice_id"])
		// if errGetDetail != nil {
		// 	return nil, errGetDetail
		// }

		// var tampungDetail []map[string]interface{}
		// for _, detail := range resultGetDetail {
		// 	singleDetail := make(map[string]interface{})
		// 	singleDetail["invoice_detail_id"] = detail["invoice_detail_id"]
		// 	singleDetail["invoice_id"] = detail["invoice_id"]
		// 	singleDetail["itemid"] = detail["itemid"]
		// 	singleDetail["account_id"] = detail["account_id"]
		// 	singleDetail["qty"] = detail["qty"]
		// 	singleDetail["uom"] = detail["uom"]
		// 	singleDetail["item_price"] = detail["item_price"]
		// 	singleDetail["tiering"] = detail["tiering"]
		// 	singleDetail["note"] = detail["note"]
		// 	singleDetail["balance_type"] = detail["balance_type"]
		// 	singleDetail["server_id"] = detail["server_id"]
		// 	singleDetail["external_user_id"] = detail["external_user_id"]
		// 	singleDetail["external_username"] = detail["external_username"]
		// 	singleDetail["external_sender"] = detail["external_sender"]

		// 	item := make(map[string]interface{})
		// 	item["item_id"] = detail["tblitem_item_id"]
		// 	item["item_name"] = detail["item_name"]
		// 	item["operator"] = detail["operator"]
		// 	item["route"] = detail["route"]
		// 	item["category"] = detail["category"]
		// 	item["uom"] = detail["uom"]

		// 	server := make(map[string]interface{})
		// 	server["server_id"] = detail["tblserver_server_id"]
		// 	server["server_name"] = detail["tblserver_server_name"]

		// 	account := make(map[string]interface{})
		// 	account["account_id"] = detail["account_id"]
		// 	account["name"] = detail["account_name"]
		// 	account["account_type"] = detail["account_type"]
		// 	account["company_id"] = detail["company_id"]
		// 	account["status"] = detail["status"]
		// 	account["non_taxable"] = detail["non_taxable"]

		// 	singleDetail["item"] = item
		// 	singleDetail["server"] = server
		// 	singleDetail["account"] = account

		// 	tampungDetail = append(tampungDetail, singleDetail)

		// }
		// single["list_invoice_detail"] = tampungDetail

		// var category
		result = append(result, single)
	}
	return result, err
}

func GetListInvoiceDetail(conn *connections.Connections, listInvoiceID []string) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	if len(listInvoiceID) > 0 {
		var baseIn string
		for _, prid := range listInvoiceID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "invoice_id IN ("+baseIn+")")
	}

	runQuery := `SELECT invoice_detail.invoice_detail_id, invoice_detail.account_id, invoice_detail.invoice_id, invoice_detail.itemid, 
	invoice_detail.qty, invoice_detail.uom, invoice_detail.item_price, tiering, invoice_detail.note, invoice_detail.balance_type, 
	invoice_detail.server_id , invoice_detail.external_sender, 
	invoice_detail.adjustment,invoice_detail.adjustment_confirmation_username,invoice_detail.adjustment_confirmation_date,
	item.item_id as tblitem_item_id, item.item_name, item.operator, item.route, item.category, 
	item.uom as tblitem_uom, server.server_id as tblserver_server_id, server.server_name as tblserver_server_name,external_user_id, external_username,
	account.name as account_name, account.account_type, account.company_id, account.status, account.non_taxable
	FROM proforma_invoice_detail as invoice_detail 
	JOIN item ON invoice_detail.itemid = item.item_id  
	JOIN server ON invoice_detail.server_id = server.server_id 
	LEFT JOIN account ON invoice_detail.account_id = account.account_id `
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func CheckInvoiceExists(invoiceID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(invoiceid) FROM invoice WHERE invoiceid = ?"
	// param = append(param, invoiceID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("Invoice ID is not exists")
	// }
	return nil
}

func ValidateRevisionCounter(invoiceID string, conn *connections.Connections) (int, error) {
	var counter int = 0
	var err error
	res, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice("SELECT revision_counter FROM invoice WHERE invoice_id = ? ", invoiceID)

	if err != nil {
		return counter, err
	}
	if len(res) == 0 {
		return counter, err
	}

	counter, _ = strconv.Atoi(res[0]["revision_counter"])

	if counter >= config.Param.MaxInvoiceRevision {
		return counter, errors.New("Unable to revise invoice")
	}

	return counter, err

}

func CheckInvoiceNoDuplicate2(invoiceNo string, conn *connections.Connections) error {
	var param []interface{}
	param = append(param, invoiceNo)
	param = append(param, "A")
	qry := "SELECT COUNT(invoice_no) FROM invoice WHERE invoice_no = ? AND invoicestatus = ?"

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	datacount, _ := strconv.Atoi(cnt)
	if datacount > 0 {
		return errors.New("Invoice Number is already used. Please use another Invoice Number")
	}
	return nil
}

func CheckInvoiceNoDuplicate(invoiceNo string, conn *connections.Connections, req datastruct.InvoiceRequest) error {
	qryCheckInvoiceNo := "SELECT invoice_id, invoice_no, invoicestatus FROM invoice WHERE invoice_no = ? AND invoicestatus = ?"
	resCheck, countInvoice, errCheckInvoiceNo := conn.DBAppConn.SelectQueryByFieldName(qryCheckInvoiceNo, invoiceNo, "A")
	if errCheckInvoiceNo != nil {
		return errCheckInvoiceNo
	}

	// if countInvoice == 0 {
	// 	qryInserControlIdInvoiceDetail := "INSERT INTO control_id (control_id.key,control_id.period,control_id.last_id) VALUES (?,?,?)"
	// 	_, _, errInsert := conn.DBAppConn.Exec(qryInserControlIdInvoiceDetail, "invoice_detail", "0", "0")
	// 	if errInsert != nil {
	// 		return "", errInsert
	// 	}

	// 	return "0", nil
	// }

	invoice := resCheck[1]

	if countInvoice > 0 {
		if invoice["invoice_id"] == req.InvoiceID {
			return nil
		} else {
			return errors.New("The invoice number have already used. please use another invoice number instead.")
		}
	}

	return nil
}

func CheckInvoiceDuplicate(exceptID string, conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var param []interface{}
	qry := "SELECT COUNT(invoiceid) FROM invoice WHERE invoiceid = ?"
	param = append(param, req.InvoiceID)
	if len(exceptID) > 0 {
		qry += " AND invoiceid <> ?"
		param = append(param, exceptID)
	}

	cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	datacount, _ := strconv.Atoi(cnt)
	if datacount > 0 {
		return errors.New("Invoice ID is already exists. Please use another Invoice ID")
	}
	return nil
}

package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/invoice/datastruct"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func GetInvoiceFromRequest(conn *connections.Connections, req datastruct.InvoiceRequest) ([]map[string]interface{}, error) {
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
	lib.AppendWhere(&baseWhere, &baseParam, "invoice.inv_type_id = ?", req.InvoiceTypeID)
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
		lib.AppendWhereRaw(&baseWhere, "invoiceid IN ("+baseIn+")")
	}

	if len(req.ListCompanyID) > 0 {
		var baseIn string
		for _, prid := range req.ListCompanyID {
			lib.AppendComma(&baseIn, &baseParam, "?", prid)
		}
		lib.AppendWhereRaw(&baseWhere, "company_id IN ("+baseIn+")")
	}

	//TERAKHIR NAMBAH QUERY BUAT NAMPILIN ACCOUNT
	runQuery := `SELECT invoice.invoice_id, invoice.invoice_no, invoice.invoice_date, invoice.invoicestatus, invoice.company_id, invoice.month_use, 
	invoice.inv_type_id, invoice.printcounter, invoice.note, invoice.canceldesc, invoice.payment_method ,invoice.attachment,invoice.last_print_username, 
	invoice.last_print_date, invoice.created_at, invoice.created_by,invoice.tax_invoice, invoice.last_update_username, invoice.exchange_rate_date, invoice.grand_total, invoice.sender,
	invoice.ppn_amount, invoice.last_update_date, invoice.discount_type, invoice.discount, invoice.ppn ,invoice.paid, invoice.due_date , invoice.adjustment_note, invoice.received_date,invoice.receipt_letter_attachment,
	invoice_type.inv_type_id as tblinvoice_type_inv_type_id, invoice_type.inv_type_name, invoice_type.server_id as tblinvoice_type_server_id, invoice_type.category, 
	invoice_type.load_from_server,invoice_type.is_group,invoice_type.group_type, invoice_type.currency_code, company.company_id as tblcompany_company_id, company.name as tblcompany_name, company.address1, company.address2, 
	company.city, company.phone,company.desc, company.contact_person, company.contact_person_phone FROM proforma_invoice as invoice JOIN invoice_type ON 
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
		single["adjustment_note"] = each["adjustment_note"]
		single["grand_total"] = each["grand_total"]
		single["tax_invoice"] = each["tax_invoice"]
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
		company["desc"] = each["desc"]
		company["phone"] = each["phone"]
		company["contact_person"] = each["contact_person"]
		company["contact_person_phone"] = each["contact_person_phone"]

		single["company"] = company

		qryGetDetail := `SELECT invoice_detail.invoice_detail_id, invoice_detail.account_id, invoice_detail.invoice_id, invoice_detail.itemid, 
		invoice_detail.qty, invoice_detail.uom, invoice_detail.item_price, tiering, invoice_detail.note, invoice_detail.balance_type, 
		invoice_detail.server_id , invoice_detail.external_sender, invoice_detail.adjustment,invoice_detail.adjustment_confirmation_username,adjustment_confirmation_date,
		item.item_id as tblitem_item_id, item.item_name, item.operator, item.route, item.category, 
		item.uom as tblitem_uom, server.server_id as tblserver_server_id, server.server_name as tblserver_server_name,external_user_id, external_username,
		account.name as account_name, account.account_type, account.company_id, account.status, account.non_taxable
		FROM proforma_invoice_detail as invoice_detail  
		JOIN item ON invoice_detail.itemid = item.item_id  
		JOIN server ON invoice_detail.server_id = server.server_id 
		LEFT JOIN account ON invoice_detail.account_id = account.account_id 
		WHERE invoice_id = ?`
		resultGetDetail, _, errGetDetail := conn.DBAppConn.SelectQueryByFieldNameSlice(qryGetDetail, single["invoice_id"])
		if errGetDetail != nil {
			return nil, errGetDetail
		}

		var tampungDetail []map[string]interface{}
		for _, detail := range resultGetDetail {
			singleDetail := make(map[string]interface{})
			singleDetail["invoice_detail_id"] = detail["invoice_detail_id"]
			singleDetail["invoice_id"] = detail["invoice_id"]
			singleDetail["itemid"] = detail["itemid"]
			singleDetail["account_id"] = detail["account_id"]
			singleDetail["qty"] = detail["qty"]
			singleDetail["uom"] = detail["uom"]
			singleDetail["item_price"] = detail["item_price"]
			singleDetail["tiering"] = detail["tiering"]
			singleDetail["note"] = detail["note"]
			singleDetail["balance_type"] = detail["balance_type"]
			singleDetail["server_id"] = detail["server_id"]
			singleDetail["external_user_id"] = detail["external_user_id"]
			singleDetail["external_username"] = detail["external_username"]
			singleDetail["external_sender"] = detail["external_sender"]
			singleDetail["adjustment"] = detail["adjustment"]
			singleDetail["adjustment_confirmation_username"] = detail["adjustment_confirmation_username"]
			singleDetail["adjustment_confirmation_date"] = detail["adjustment_confirmation_date"]

			item := make(map[string]interface{})
			item["item_id"] = detail["tblitem_item_id"]
			item["item_name"] = detail["item_name"]
			item["operator"] = detail["operator"]
			item["route"] = detail["route"]
			item["category"] = detail["category"]
			item["uom"] = detail["uom"]

			server := make(map[string]interface{})
			server["server_id"] = detail["tblserver_server_id"]
			server["server_name"] = detail["tblserver_server_name"]

			account := make(map[string]interface{})
			account["account_id"] = detail["account_id"]
			account["name"] = detail["account_name"]
			account["account_type"] = detail["account_type"]
			account["company_id"] = detail["company_id"]
			account["status"] = detail["status"]
			account["non_taxable"] = detail["non_taxable"]

			singleDetail["item"] = item
			singleDetail["server"] = server
			singleDetail["account"] = account

			tampungDetail = append(tampungDetail, singleDetail)

		}
		single["list_invoice_detail"] = tampungDetail

		// var category
		result = append(result, single)
	}
	return result, err
}

func InsertInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "proforma_invoice")
	// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	//validate invoice number
	errCheckInvoiceNo := CheckInvoiceNoDuplicate(req.InvoiceNo, conn, req)
	if errCheckInvoiceNo != nil {
		return errCheckInvoiceNo
	}

	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	// lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceNo)
	lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceDate)
	lib.AppendComma(&baseIn, &baseParam, "?", "A")
	lib.AppendComma(&baseIn, &baseParam, "?", req.CompanyID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.MonthUse)
	lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceTypeID)
	lib.AppendCommaRaw(&baseIn, "now()")
	lib.AppendComma(&baseIn, &baseParam, "?", req.DiscountType)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Discount)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Note)
	lib.AppendCommaRaw(&baseIn, "now()")
	createdBy := req.LastUpdateUsername
	lib.AppendComma(&baseIn, &baseParam, "?", createdBy)
	lib.AppendComma(&baseIn, &baseParam, "?", "0")
	lib.AppendComma(&baseIn, &baseParam, "?", req.PPN)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentMethod)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ExchangeRateDate)
	lib.AppendComma(&baseIn, &baseParam, "?", req.DueDate)
	lib.AppendComma(&baseIn, &baseParam, "?", req.TaxInvoice)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Sender)
	lib.AppendComma(&baseIn, &baseParam, "?", req.BatchID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Attachment)
	lib.AppendComma(&baseIn, &baseParam, "?", req.AdjustmentNote)
	lib.AppendComma(&baseIn, &baseParam, "?", req.GrandTotal)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PPNAmount)

	monthUse := req.MonthUse
	// invoiceDateParsed, _ := time.Parse("2006-01-02", req.InvoiceDate)
	// prefixInvoiceNo := invoiceDateParsed.Format("200601")[2:6]

	if len(monthUse) == 0 {
		//tentukan month use dari invoice date
		parsedDate, _ := time.Parse("2006-01-02", req.InvoiceDate)
		monthUse = parsedDate.Format("200601")
	}

	tx, err := conn.DBAppConn.DB.Begin()
	if err != nil {
		return err
	}
	// prefixInvoiceNo := invoiceDateFormated
	// qryCheckControlIdPeriod := "SELECT control_id.key, control_id.period ,last_id FROM control_id WHERE control_id.key = ? AND period = ?"
	// resCheck, countControlIdPeriod, errCheckControlIdPeriod := conn.DBAppConn.SelectQueryByFieldName(qryCheckControlIdPeriod, "proforma_invoice_no", prefixInvoiceNo)
	// if errCheckControlIdPeriod != nil {
	// 	return errCheckControlIdPeriod
	// }

	// lastIdPeriod := resCheck[1]["last_id"]

	// if countControlIdPeriod == 0 {
	// 	errInsertControlIdPeriod := InsertControlIdPeriod(conn, prefixInvoiceNo)
	// 	if errInsertControlIdPeriod != nil {
	// 		return errInsertControlIdPeriod
	// 	}
	// }

	qry := "INSERT INTO proforma_invoice (invoice_id, invoice_date,invoicestatus, company_id, month_use ,inv_type_id, last_update_date, discount_type, discount,note, created_at,created_by, printcounter,ppn, payment_method, exchange_rate_date, due_date,tax_invoice, sender, batch_id,attachment,adjustment_note,grand_total,ppn_amount) VALUES (" + baseIn + ")"
	_, errInsert := tx.Exec(qry, baseParam...)
	if errInsert != nil {
		tx.Rollback()
		return errInsert
	}
	_, errUpdateId := tx.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "proforma_invoice")
	if errUpdateId != nil {
		tx.Rollback()
		return errUpdateId
	}

	// err = UpdateInvoiceNo(tx, countControlIdPeriod, prefixInvoiceNo, lastIdPeriod, insertIdString)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	if len(req.ListInvoiceDetail) > 0 {

		for _, each := range req.ListInvoiceDetail {
			var baseInDetail string
			var baseParamDetail []interface{}
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", insertIdString)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.AccountID)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ItemID)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.Qty)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.Adjustment)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.Uom)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ItemPrice)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.Note)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.BalanceType)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ServerID)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ExternalUserID)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ExternalUsername)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ExternalSender)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", createdBy)
			qryInsertDetail := "INSERT INTO proforma_invoice_detail (invoice_id,account_id ,itemid, qty,adjustment, uom, item_price,note, balance_type, server_id, external_user_id,external_username,external_sender ,last_update_username) VALUES (" + baseInDetail + ")"
			// var invDetailParam []interface{} = []interface{}{
			// 	insertIdString,
			// 	each.AccountID,
			// 	each.ItemID,
			// 	each.Qty,
			// 	each.Uom,
			// 	each.ItemPrice,
			// 	each.Note,
			// 	each.BalanceType,
			// 	each.ServerID,
			// 	each.ExternalUserID,
			// 	each.ExternalUsername,
			// 	createdBy,
			// }

			_, errInsertDetail := tx.Exec(qryInsertDetail, baseParamDetail...)
			if errInsertDetail != nil {
				tx.Rollback()
				return errInsertDetail
			}
		}

		var serverDataWhereClause string
		var serverDataUpdateParam []interface{}

		serverDataUpdateParam = append(serverDataUpdateParam, insertIdString)

		if req.InvoiceTypeID == "6" {
			var listUsername string
			for _, each := range req.ListInvoiceDetail {
				lib.AppendComma(&listUsername, &serverDataUpdateParam, "?", each.ExternalUsername)
			}
			lib.AppendWhereRaw(&serverDataWhereClause, "external_username IN ("+listUsername+")")

		} else if req.InvoiceTypeID == "7" {
			var listSender string
			for _, each := range req.ListInvoiceDetail {
				lib.AppendComma(&listSender, &serverDataUpdateParam, "?", each.ExternalSender)
			}
			lib.AppendWhereRaw(&serverDataWhereClause, "external_sender IN ("+listSender+")")
		}
		if len(req.ListExternalRootParentAccount) > 0 {
			if req.InvoiceTypeID != "6" {
				var rootParentAccountIn string
				for _, id := range req.ListExternalRootParentAccount {
					lib.AppendComma(&rootParentAccountIn, &serverDataUpdateParam, "?", id)
				}
				lib.AppendWhereRaw(&serverDataWhereClause, "external_rootparent_account IN ("+rootParentAccountIn+")")
			}
		}

		if len(req.ListServerID) > 0 {
			if len(req.ListServerID) == 1 && req.ListServerID[0] != "*" {
				lib.AppendWhere(&serverDataWhereClause, &serverDataUpdateParam, "server_data.server_id = ?", req.ListServerID[0])
			}
			if len(req.ListServerID) > 1 {
				var listServer string
				for _, each := range req.ListServerID {
					lib.AppendComma(&listServer, &serverDataUpdateParam, "?", each)
				}
				lib.AppendWhereRaw(&serverDataWhereClause, "server_data.server_id IN ("+listServer+")")
			}
		}
		if len(req.ListAccountID) > 0 {
			var accountIDIn string
			for _, id := range req.ListAccountID {
				lib.AppendComma(&accountIDIn, &serverDataUpdateParam, "?", id)
			}

			lib.AppendWhereRaw(&serverDataWhereClause, "server_data.account_id IN ("+accountIDIn+")")
		}
		if len(req.AdditionalParamOperator) > 0 {
			var listOperator string
			for _, each := range req.AdditionalParamOperator {
				lib.AppendComma(&listOperator, &serverDataUpdateParam, "?", each)
			}
			lib.AppendWhereRaw(&serverDataWhereClause, "server_data.external_operatorcode IN ("+listOperator+")")
		}

		if len(req.AdditionalParamRoute) > 0 {
			var listRoute string
			for _, each := range req.AdditionalParamRoute {
				lib.AppendComma(&listRoute, &serverDataUpdateParam, "?", each)
			}
			lib.AppendWhereRaw(&serverDataWhereClause, "server_data.external_route IN ("+listRoute+")")
		}
		lib.AppendWhere(&serverDataWhereClause, &serverDataUpdateParam, "DATE_FORMAT(server_data.external_transdate, '%Y%m') = ?", req.MonthUse)
		lib.AppendWhere(&serverDataWhereClause, &serverDataUpdateParam, "server_data.external_sender = ?", req.Sender)
		// lib.AppendWhere(&serverDataWhereClause, &serverDataUpdateParam, "server_data.external_operatorcode = ?", req.AdditionalParamOperator)
		// lib.AppendWhere(&serverDataWhereClause, &serverDataUpdateParam, "server_data.external_route = ?", req.AdditionalParamRoute)
		lib.AppendWhereRaw(&serverDataWhereClause, "server_data.invoice_id IS NULL")
		qryUpdateServerData := "UPDATE server_data set server_data.invoice_id = ? "

		if len(serverDataWhereClause) > 0 {
			qryUpdateServerData += "WHERE " + serverDataWhereClause
		}
		// var updateServerDataParam []interface{} = []interface{}{
		// 	insertIdString,
		// 	req.AccountID,
		// 	req.ServerID,
		// 	req.MonthUse,
		// }

		logrus.Info("Query Update Server Data-", qryUpdateServerData)
		logrus.Info("Update Server Data Where Clause-", serverDataWhereClause)
		logrus.Info("Update Server Data Param -", serverDataUpdateParam)
		// if len(req.ListServerID) > 0 {
		// 	_, errUpdateServerData := tx.Exec(qryUpdateServerData, serverDataUpdateParam...)
		// 	if errUpdateServerData != nil {
		// 		tx.Rollback()
		// 		logrus.Error("Error in SQL : ", errUpdateServerData)
		// 		return errUpdateServerData
		// 	}
		// } else {
		// 	logrus.Info("This invoice will not updating server data")
		// }

	}
	tx.Commit()

	return err
}

func InsertCustomInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	var baseIn string
	var baseParam []interface{}

	lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "invoice")
	// lastId, _ := conn.DBAppConn.GetFirstData("SELECT last_id FROM control_id where control_id.key=?", "account")

	intLastId, err := strconv.Atoi(lastId)
	insertId := intLastId + 1

	insertIdString := strconv.Itoa(insertId)

	//validate invoice number
	errCheckInvoiceNo := CheckInvoiceNoDuplicate(req.InvoiceNo, conn, req)
	if errCheckInvoiceNo != nil {
		return errCheckInvoiceNo
	}

	lib.AppendComma(&baseIn, &baseParam, "?", insertIdString)
	// lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceNo)
	lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceDate)
	lib.AppendComma(&baseIn, &baseParam, "?", "A")
	lib.AppendComma(&baseIn, &baseParam, "?", req.CompanyID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.MonthUse)
	lib.AppendComma(&baseIn, &baseParam, "?", req.InvoiceTypeID)
	lib.AppendCommaRaw(&baseIn, "now()")
	lib.AppendComma(&baseIn, &baseParam, "?", req.DiscountType)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Discount)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Note)
	lib.AppendCommaRaw(&baseIn, "now()")
	createdBy := req.LastUpdateUsername
	lib.AppendComma(&baseIn, &baseParam, "?", createdBy)
	lib.AppendComma(&baseIn, &baseParam, "?", "0")
	lib.AppendComma(&baseIn, &baseParam, "?", req.PPN)
	lib.AppendComma(&baseIn, &baseParam, "?", req.PaymentMethod)
	lib.AppendComma(&baseIn, &baseParam, "?", req.ExchangeRateDate)
	lib.AppendComma(&baseIn, &baseParam, "?", req.DueDate)
	lib.AppendComma(&baseIn, &baseParam, "?", req.TaxInvoice)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Sender)
	lib.AppendComma(&baseIn, &baseParam, "?", req.BatchID)
	lib.AppendComma(&baseIn, &baseParam, "?", req.Attachment)

	monthUse := req.MonthUse
	invoiceDateParsed, _ := time.Parse("2006-01-02", req.InvoiceDate)
	prefixInvoiceNo := invoiceDateParsed.Format("200601")[2:6]

	if len(monthUse) == 0 {
		//tentukan month use dari invoice date
		parsedDate, _ := time.Parse("2006-01-02", req.InvoiceDate)
		monthUse = parsedDate.Format("200601")
	}

	tx, err := conn.DBAppConn.DB.Begin()
	if err != nil {
		return err
	}
	// prefixInvoiceNo := invoiceDateFormated
	qryCheckControlIdPeriod := "SELECT control_id.key, control_id.period ,last_id FROM control_id WHERE control_id.key = ? AND period = ?"
	resCheck, countControlIdPeriod, errCheckControlIdPeriod := conn.DBAppConn.SelectQueryByFieldName(qryCheckControlIdPeriod, "invoice_no", prefixInvoiceNo)
	if errCheckControlIdPeriod != nil {
		return errCheckControlIdPeriod
	}

	lastIdPeriod := resCheck[1]["last_id"]

	if countControlIdPeriod == 0 {
		errInsertControlIdPeriod := InsertControlIdPeriod(conn, prefixInvoiceNo)
		if errInsertControlIdPeriod != nil {
			return errInsertControlIdPeriod
		}
	}

	qry := "INSERT INTO invoice (invoice_id, invoice_date,invoicestatus, company_id, month_use ,inv_type_id, last_update_date, discount_type, discount, invoice.note, created_at,created_by, printcounter,ppn, payment_method, exchange_rate_date, due_date,tax_invoice, sender, batch_id,attachment) VALUES (" + baseIn + ")"
	_, errInsert := tx.Exec(qry, baseParam...)
	if errInsert != nil {
		tx.Rollback()
		return errInsert
	}
	_, errUpdateId := tx.Exec("UPDATE control_id set last_id=? where control_id.key=?", insertIdString, "invoice")
	if errUpdateId != nil {
		tx.Rollback()
		return errUpdateId
	}

	err = UpdateInvoiceNo(tx, countControlIdPeriod, prefixInvoiceNo, lastIdPeriod, insertIdString)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(req.ListInvoiceDetail) > 0 {

		for _, each := range req.ListInvoiceDetail {
			var baseInDetail string
			var baseParamDetail []interface{}
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", insertIdString)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.AccountID)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ItemID)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.Qty)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.Uom)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ItemPrice)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.Note)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.BalanceType)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ServerID)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ExternalUserID)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ExternalUsername)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", each.ExternalSender)
			lib.AppendComma(&baseInDetail, &baseParamDetail, "?", createdBy)
			qryInsertDetail := "INSERT INTO invoice_detail (invoice_detail.invoice_id,invoice_detail.account_id ,invoice_detail.itemid, invoice_detail.qty, invoice_detail.uom, invoice_detail.item_price,invoice_detail.note, invoice_detail.balance_type, invoice_detail.server_id, invoice_detail.external_user_id,invoice_detail.external_username,external_sender ,last_update_username) VALUES (" + baseInDetail + ")"
			// var invDetailParam []interface{} = []interface{}{
			// 	insertIdString,
			// 	each.AccountID,
			// 	each.ItemID,
			// 	each.Qty,
			// 	each.Uom,
			// 	each.ItemPrice,
			// 	each.Note,
			// 	each.BalanceType,
			// 	each.ServerID,
			// 	each.ExternalUserID,
			// 	each.ExternalUsername,
			// 	createdBy,
			// }

			_, errInsertDetail := tx.Exec(qryInsertDetail, baseParamDetail...)
			if errInsertDetail != nil {
				tx.Rollback()
				return errInsertDetail
			}

			//update server data
			_, errUpdateServerData := tx.Exec("UPDATE server_data SET invoice_id = ? WHERE item_id = ? AND account_id = ? AND server_id = ? AND external_sender = ? AND external_username = ?", insertIdString, each.ItemID, each.AccountID, each.ServerID, each.ExternalSender, each.ExternalUsername)
			if errUpdateServerData != nil {
				tx.Rollback()
				return errUpdateServerData
			}
		}

		// var serverDataWhereClause string
		// var serverDataUpdateParam []interface{}

		// serverDataUpdateParam = append(serverDataUpdateParam, insertIdString)

		// if req.InvoiceTypeID == "6" {
		// 	var listUserID string
		// 	for _, each := range req.ListInvoiceDetail {
		// 		lib.AppendComma(&listUserID, &serverDataUpdateParam, "?", each.ExternalUserID)
		// 	}
		// 	lib.AppendWhereRaw(&serverDataWhereClause, "external_user_id IN ("+listUserID+")")

		// } else if req.InvoiceTypeID == "7" {
		// 	var listSender string
		// 	for _, each := range req.ListInvoiceDetail {
		// 		lib.AppendComma(&listSender, &serverDataUpdateParam, "?", each.ExternalSender)
		// 	}
		// 	lib.AppendWhereRaw(&serverDataWhereClause, "external_sender IN ("+listSender+")")
		// }
		// if len(req.ListExternalRootParentAccount) > 0 {
		// 	if req.InvoiceTypeID != "6" {
		// 		var rootParentAccountIn string
		// 		for _, id := range req.ListExternalRootParentAccount {
		// 			lib.AppendComma(&rootParentAccountIn, &serverDataUpdateParam, "?", id)
		// 		}
		// 		lib.AppendWhereRaw(&serverDataWhereClause, "external_rootparent_account IN ("+rootParentAccountIn+")")
		// 	}
		// }

		// if len(req.ListServerID) > 0 {
		// 	if len(req.ListServerID) == 1 && req.ListServerID[0] != "*" {
		// 		lib.AppendWhere(&serverDataWhereClause, &serverDataUpdateParam, "server_data.server_id = ?", req.ListServerID[0])
		// 	}
		// 	if len(req.ListServerID) > 1 {
		// 		var listServer string
		// 		for _, each := range req.ListServerID {
		// 			lib.AppendComma(&listServer, &serverDataUpdateParam, "?", each)
		// 		}
		// 		lib.AppendWhereRaw(&serverDataWhereClause, "server_data.server_id IN ("+listServer+")")
		// 	}
		// }
		// if len(req.ListAccountID) > 0 {
		// 	var accountIDIn string
		// 	for _, id := range req.ListAccountID {
		// 		lib.AppendComma(&accountIDIn, &serverDataUpdateParam, "?", id)
		// 	}

		// 	lib.AppendWhereRaw(&serverDataWhereClause, "server_data.account_id IN ("+accountIDIn+")")
		// }
		// if len(req.AdditionalParamOperator) > 0 {
		// 	var listOperator string
		// 	for _, each := range req.AdditionalParamOperator {
		// 		lib.AppendComma(&listOperator, &serverDataUpdateParam, "?", each)
		// 	}
		// 	lib.AppendWhereRaw(&serverDataWhereClause, "server_data.external_operatorcode IN ("+listOperator+")")
		// }

		// if len(req.AdditionalParamRoute) > 0 {
		// 	var listRoute string
		// 	for _, each := range req.AdditionalParamRoute {
		// 		lib.AppendComma(&listRoute, &serverDataUpdateParam, "?", each)
		// 	}
		// 	lib.AppendWhereRaw(&serverDataWhereClause, "server_data.external_route IN ("+listRoute+")")
		// }
		// lib.AppendWhere(&serverDataWhereClause, &serverDataUpdateParam, "DATE_FORMAT(server_data.external_transdate, '%Y%m') = ?", req.MonthUse)
		// lib.AppendWhere(&serverDataWhereClause, &serverDataUpdateParam, "server_data.external_sender = ?", req.Sender)
		// // lib.AppendWhere(&serverDataWhereClause, &serverDataUpdateParam, "server_data.external_operatorcode = ?", req.AdditionalParamOperator)
		// // lib.AppendWhere(&serverDataWhereClause, &serverDataUpdateParam, "server_data.external_route = ?", req.AdditionalParamRoute)
		// lib.AppendWhereRaw(&serverDataWhereClause, "server_data.invoice_id IS NULL")
		// qryUpdateServerData := "UPDATE server_data set server_data.invoice_id = ? "

		// if len(serverDataWhereClause) > 0 {
		// 	qryUpdateServerData += "WHERE " + serverDataWhereClause
		// }
		// // var updateServerDataParam []interface{} = []interface{}{
		// // 	insertIdString,
		// // 	req.AccountID,
		// // 	req.ServerID,
		// // 	req.MonthUse,
		// // }

		// logrus.Info("Query Update Server Data-", qryUpdateServerData)
		// logrus.Info("Update Server Data Where Clause-", serverDataWhereClause)
		// logrus.Info("Update Server Data Param -", serverDataUpdateParam)
		// if len(req.ListServerID) > 0 {
		// 	_, errUpdateServerData := tx.Exec(qryUpdateServerData, serverDataUpdateParam...)
		// 	if errUpdateServerData != nil {
		// 		tx.Rollback()
		// 		logrus.Error("Error in SQL : ", errUpdateServerData)
		// 		return errUpdateServerData
		// 	}
		// } else {
		// 	logrus.Info("This invoice will not updating server data")
		// }

	}
	tx.Commit()

	return err
}

func InsertControlIdPeriod(conn *connections.Connections, monthUse string) error {
	var err error
	qryInserControlIdPeriod := "INSERT INTO control_id (control_id.key,control_id.period,control_id.last_id) VALUES (?,?,?)"
	_, _, err = conn.DBAppConn.Exec(qryInserControlIdPeriod, "proforma_invoice_no", monthUse, "0")
	if err != nil {
		return err
	}

	return err
}

func CheckControlIdInvoiceDetail(conn *connections.Connections) (string, error) {
	qryCheckControlIdInvoiceDetail := "SELECT last_id FROM control_id WHERE control_id.key = ?"
	resCheck, countControlIdInvoiceDetail, errCheckControlIdInvoiceDetail := conn.DBAppConn.SelectQueryByFieldName(qryCheckControlIdInvoiceDetail, "invoice_detail")
	if errCheckControlIdInvoiceDetail != nil {
		return "", errCheckControlIdInvoiceDetail
	}

	if countControlIdInvoiceDetail == 0 {
		qryInserControlIdInvoiceDetail := "INSERT INTO control_id (control_id.key,control_id.period,control_id.last_id) VALUES (?,?,?)"
		_, _, errInsert := conn.DBAppConn.Exec(qryInserControlIdInvoiceDetail, "invoice_detail", "0", "0")
		if errInsert != nil {
			return "", errInsert
		}

		return "0", nil
	}

	lastIdInvoiceDetail := resCheck[1]["last_id"]
	return lastIdInvoiceDetail, nil
}

func UpdateControlId(conn *connections.Connections, newId string, key string) error {
	var err error
	_, _, err = conn.DBAppConn.Exec("UPDATE control_id set last_id = ? WHERE control_id.key=?", newId, key)
	return err
}

func UpdateInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	var baseUpInvoice string
	var baseParam []interface{}

	lib.AppendComma(&baseUpInvoice, &baseParam, "invoice_no = ?", req.InvoiceNo)
	lib.AppendComma(&baseUpInvoice, &baseParam, "invoice_date = ?", req.InvoiceDate)
	lib.AppendComma(&baseUpInvoice, &baseParam, "discount_type = ?", req.DiscountType)
	lib.AppendComma(&baseUpInvoice, &baseParam, "discount = ?", req.Discount)
	lib.AppendComma(&baseUpInvoice, &baseParam, "note = ?", req.Note)
	lib.AppendComma(&baseUpInvoice, &baseParam, "last_update_username = ?", req.LastUpdateUsername)
	lib.AppendComma(&baseUpInvoice, &baseParam, "ppn = ?", req.PPN)
	lib.AppendComma(&baseUpInvoice, &baseParam, "payment_method = ?", req.PaymentMethod)
	lib.AppendComma(&baseUpInvoice, &baseParam, "attachment = ?", req.Attachment)
	lib.AppendComma(&baseUpInvoice, &baseParam, "tax_invoice = ?", req.TaxInvoice)
	lib.AppendComma(&baseUpInvoice, &baseParam, "grand_total = ?", req.GrandTotal)
	lib.AppendComma(&baseUpInvoice, &baseParam, "ppn_amount = ?", req.PPNAmount)

	errCheckInvoiceNo := CheckInvoiceNoDuplicate(req.InvoiceNo, conn, req)
	if errCheckInvoiceNo != nil {
		return errCheckInvoiceNo
	}

	qry := "UPDATE proforma_invoice SET " + baseUpInvoice + " WHERE invoice_id = ?"
	baseParam = append(baseParam, req.InvoiceID)
	// baseParam = append(baseParam, req.InvoiceNo)
	// baseParam = append(baseParam, req.TransDate)
	// baseParam = append(baseParam, req.DiscountType)
	// baseParam = append(baseParam, req.Discount)
	// baseParam = append(baseParam, req.Note)
	// baseParam = append(baseParam, req.LastUpdateUsername)
	// baseParam = append(baseParam, req.PPN)
	// baseParam = append(baseParam, "now()")
	// baseParam = append(baseParam, req.InvoiceID)
	_, _, errUpdateInvoice := conn.DBAppConn.Exec(qry, baseParam...)
	if errUpdateInvoice != nil {
		return errUpdateInvoice
	}

	// qryGetInvoiceDetail := "SELECT invoice_detail_id, invoice_id, itemid, qty, uom, item_price FROM invoice_detail WHERE invoice_id = ? "
	// resultGetDetail, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(qryGetInvoiceDetail, req.InvoiceID)
	// for _, each := range resultGetDetail {
	// 	idx := slices.IndexFunc(req.ListInvoiceDetail, func(data datastruct.InvoiceDetailStruct) bool {
	// 		return data.ItemID == each["itemid"] && req.InvoiceID == each["invoice_id"] && data.Uom == each["uom"] && data.ItemPrice == each["item_price"]
	// 	})

	// 	//row ini ga dibawa dari request. delete
	// 	if idx == -1 {
	// 		_, _, err = conn.DBAppConn.Exec("DELETE FROM invoice_detail WHERE invoice_detail_id = ?", each["invoice_detail_id"])
	// 		if err != nil {
	// 			logrus.Errorf("Error in SQL : %v", err)
	// 			return err
	// 		}
	// 	}
	// }

	_, _, errUpdateId := conn.DBAppConn.Exec("DELETE FROM proforma_invoice_detail where invoice_id=?", req.InvoiceID)
	if errUpdateId != nil {
		return errUpdateId
	}

	if len(req.ListInvoiceDetail) > 0 {
		for _, each := range req.ListInvoiceDetail {

			var baseUpDetail string
			var baseParamDetail []interface{}
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", req.InvoiceID)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.AccountID)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.Qty)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.ItemID)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.Uom)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.ItemPrice)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.Note)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.BalanceType)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.ServerID)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.ExternalUserID)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.ExternalUsername)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", each.ExternalSender)
			lib.AppendComma(&baseUpDetail, &baseParamDetail, "?", req.LastUpdateUsername)
			// qryInsertDetail := "INSERT INTO invoice_detail (invoice_detail.invoice_id,invoice_detail.account_id ,invoice_detail.itemid, invoice_detail.qty, invoice_detail.uom, invoice_detail.item_price,invoice_detail.note, invoice_detail.balance_type, invoice_detail.server_id, invoice_detail.external_user_id,invoice_detail.external_username ,last_update_username) VALUES (" + baseInDetail + ")"
			qryUpdateDetail := "INSERT INTO proforma_invoice_detail (invoice_id, account_id,qty,itemid,uom,item_price,note,balance_type,server_id,external_user_id,external_username,external_sender,last_update_username) VALUES (" + baseUpDetail + ")"
			// var invDetailParam []interface{} = []interface{}{
			// 	req.InvoiceID,
			// 	each.AccountID,
			// 	each.Qty,
			// 	each.ItemID,
			// 	each.Uom,
			// 	each.ItemPrice,
			// 	each.Note,
			// 	each.BalanceType,
			// 	each.ServerID,
			// 	each.ExternalUserID,
			// 	each.ExternalUsername,
			// 	req.LastUpdateUsername,
			// }

			_, _, errUpdateDetail := conn.DBAppConn.Exec(qryUpdateDetail, baseParamDetail...)
			if errUpdateDetail != nil {
				logrus.Errorf("Error in SQL : %v", errUpdateDetail)
				return errUpdateDetail
			}
		}
	}

	return err
}

func DeleteInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	qry := "UPDATE proforma_invoice SET invoicestatus = ?, canceldesc = ? WHERE invoice_id = ?"
	// qryUpdateServerData := "UPDATE server_data SET invoice_id = NULL where server_data.invoice_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, "D", req.CancelDesc, req.InvoiceID)
	// _, _, err = conn.DBAppConn.Exec(qryUpdateServerData, req.InvoiceID)

	return err
}

func CancelInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	qry := "UPDATE proforma_invoice SET invoicestatus = ?, canceldesc = ? WHERE invoice_id = ?"
	// qryUpdateServerData := "UPDATE server_data SET invoice_id = NULL where server_data.invoice_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, "D", req.CancelDesc, req.InvoiceID)

	// _, _, err = conn.DBAppConn.Exec(qryUpdateServerData, req.InvoiceID)
	return err
}

func PrintInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	qry := "UPDATE proforma_invoice SET printcounter = printcounter + 1, last_print_username = ?, last_print_date = now(), grand_total = ?, ppn_amount = ?  WHERE invoice_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.LastUpdateUsername, req.GrandTotal, req.PPNAmount, req.InvoiceID)
	return err
}

func UpdateReceivedDate(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	qry := "UPDATE invoice SET received_date = ?, receipt_letter_attachment = ? WHERE invoice_id = ?"
	_, _, err = conn.DBAppConn.Exec(qry, req.ReceivedDate, req.ReceiptLetterAttachment, req.InvoiceID)
	return err
}

func InsertGenerateInvoiceHistory(conn *connections.Connections, req datastruct.GenerateInvoiceDataStruct) (int64, error) {
	var err error
	var baseIn string
	var baseParam []interface{}

	lib.AppendComma(&baseIn, &baseParam, "?", req.GenerateTime)
	lib.AppendComma(&baseIn, &baseParam, "?", req.LastUpdateUsername)

	qry := "INSERT INTO cms_batch_invoice (generate_time, last_update_username) VALUES (" + baseIn + ")"
	lastID, err := conn.DBFeConn.InsertGetLastID(qry, baseParam...)

	return lastID, err

}

func UpdateProgressGenerateInvoiceHistory(conn *connections.Connections, progress int, generateID string) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	var baseUp string
	var baseParam []interface{}

	progressStr := strconv.Itoa(progress)
	lib.AppendComma(&baseUp, &baseParam, "progress = ?", progressStr)
	qry := "UPDATE cms_batch_invoice SET " + baseUp + " WHERE id = ?"
	baseParam = append(baseParam, generateID)
	_, _, err = conn.DBFeConn.Exec(qry, baseParam...)
	return err
}

func UpdateRevisionCounter(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error
	invoice, err := GetInvoiceFromRequest(conn, req)

	if len(invoice) == 0 {
		return errors.New("Invoice Not Found. Can not update revision counter")
	}
	printCounter := invoice[0]["printcounter"].(string)

	printCounterInt, _ := strconv.Atoi(printCounter)

	// kalau belum dicetak, langsung return saja.
	if printCounterInt == 0 {
		return err
	}
	revisionCounter, _ := ValidateRevisionCounter(req.InvoiceID, conn)

	//langsung increment saja. udh divalidasi di service
	revisionCounter = revisionCounter + 1

	_, _, err = conn.DBAppConn.Exec("UPDATE proforma_invoice SET revision_counter = ? WHERE invoice_id = ? ", revisionCounter, req.InvoiceID)
	if err != nil {
		return err
	}

	return err

}

func UpdateInvoiceNo(tx *sql.Tx, countControlIdPeriod int, prefixInvoiceNo string, lastIdPeriod string, insertIdString string) error {
	var err error
	if countControlIdPeriod == 0 {
		seqNum := "0001"
		invoiceNo := prefixInvoiceNo + lib.StrPadLeft(seqNum, 4, "0")
		qryUpdateInvoice := "UPDATE proforma_invoice set invoice_no = ? WHERE invoice_id = ?"
		_, errUpdateInvoice := tx.Exec(qryUpdateInvoice, invoiceNo, insertIdString)
		if errUpdateInvoice != nil {
			tx.Rollback()
			return errUpdateInvoice
		}
		_, errUpdateLastIdPeriod := tx.Exec("UPDATE control_id set last_id = ? WHERE control_id.key=? AND control_id.period = ?", seqNum, "proforma_invoice_no", prefixInvoiceNo)
		if errUpdateLastIdPeriod != nil {
			tx.Rollback()
			return errUpdateLastIdPeriod
		}
	} else {
		//ambil 6 angka pertama ( atau bisa dibilang month use nya)
		// first6 := lastIdPeriod[0:6]
		//4 angka terakhir sebagai id increment nya
		last4 := lastIdPeriod
		last4Int, errConv := strconv.Atoi(last4)
		if errConv != nil {
			tx.Rollback()
			return errConv
		}
		last4Int += 1
		newLast4IdStr := strconv.Itoa(last4Int)
		//digabung kembali
		newInvoiceNo := prefixInvoiceNo + lib.StrPadLeft(newLast4IdStr, 4, "0")

		qryUpdateInvoice := "UPDATE proforma_invoice set invoice_no = ? WHERE invoice_id = ?"
		_, errUpdateInvoice := tx.Exec(qryUpdateInvoice, newInvoiceNo, insertIdString)
		if errUpdateInvoice != nil {
			tx.Rollback()
			return errUpdateInvoice
		}

		_, errUpdateLastIdPeriod := tx.Exec("UPDATE control_id set last_id = ? WHERE control_id.key=? AND control_id.period = ?", newLast4IdStr, "proforma_invoice_no", prefixInvoiceNo)
		if errUpdateLastIdPeriod != nil {
			tx.Rollback()
			return errUpdateLastIdPeriod
		}
	}

	return err
}

// summary
func GetInquiryPayment(conn *connections.Connections, req datastruct.InquiryPaymentRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "account.company_id = ?", req.CompanyID)
	lib.AppendWhereLike(&baseWhere, &baseParam, "company.name like ? ", req.CompanyName)
	// lib.AppendWhere(&baseWhere, &baseParam, "compan.company_id = ?", req.CompanyName)
	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(server_data.external_transdate, '%Y%m') = ?", req.MonthUse)

	qry := ` SELECT company.company_id,company.name as company_name,server_data.account_id, account.name as account_name,
	sum(item_price.price * external_smscount) as amount FROM dbbilling.server_data
	LEFT JOIN account ON account.account_id = server_data.account_id
	LEFT JOIN company on account.company_id = company.company_id
	LEFT JOIN item_price ON server_data.item_id=item_price.item_id AND server_data.server_id=item_price.server_id 
	AND item_price.company_id = account.company_id `

	if len(baseWhere) > 0 {
		qry += " WHERE " + baseWhere
	}

	qry += " GROUP BY 1,2,3,4"

	lib.AppendOrderBy(&qry, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&qry, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(qry, baseParam...)

	return result, err
}

func GetInquiryPaymentDetail(conn *connections.Connections, req datastruct.InquiryPaymentRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "account.company_id = ?", req.CompanyID)
	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(server_data.external_transdate, '%Y%m') = ?", req.MonthUse)

	qry := ` SELECT company.company_id,company.name as company_name,server_data.account_id, account.name as account_name,
	server_data.invoice_id, item.item_id, item.item_name,external_user_id,external_username, external_sender,invoice.invoice_no,
	sum(item_price.price * external_smscount) as amount FROM dbbilling.server_data
	LEFT JOIN invoice ON invoice.invoice_id = server_data.invoice_id
	LEFT JOIN account ON account.account_id = server_data.account_id
	LEFT JOIN company on account.company_id = company.company_id
	LEFT JOIN item ON item.item_id = server_data.item_id
	LEFT JOIN item_price ON server_data.item_id=item_price.item_id AND server_data.server_id=item_price.server_id 
	AND item_price.company_id = account.company_id `

	if len(baseWhere) > 0 {
		qry += " WHERE " + baseWhere
	}

	qry += " GROUP BY 1,2,3,4,5,6,7,8,9,10,11"

	lib.AppendOrderBy(&qry, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&qry, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(qry, baseParam...)

	return result, err
}

func AdjustmentConfirmation(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	var baseUp string
	var baseParam []interface{}

	lib.AppendComma(&baseUp, &baseParam, "adjustment_confirmation_username = ?", req.LastUpdateUsername)
	lib.AppendCommaRaw(&baseUp, "adjustment_confirmation_date = now()")

	qry := "UPDATE invoice_detail SET " + baseUp + " WHERE invoice_id = ? AND adjustment IS NOT NULL"
	baseParam = append(baseParam, req.InvoiceID)
	_, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	return err
}

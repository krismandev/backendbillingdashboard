package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	"backendbillingdashboard/modules/server-data/datastruct"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func GetServerDataFromRequest(conn *connections.Connections, req datastruct.ServerDataRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var resultQuery []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	baseParam = append(baseParam, req.Category)
	baseParam = append(baseParam, req.CurrencyCode)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data_id = ?", req.ServerDataID)
	// lib.AppendWhere(&baseWhere, &baseParam, "server_data.server_id = ?", req.ServerID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.external_user_id = ?", req.ExternalUserID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.external_username = ?", req.ExternalUsername)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.external_sender = ?", req.ExternalSender)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.item_id = ?", req.ItemID)
	lib.AppendWhere(&baseWhere, &baseParam, "server_data.external_rootparent_account = ?", req.ExternalRootParentAccount)
	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(server_data.external_transdate, '%Y%m') = ?", req.MonthUse)
	// lib.AppendWhere(&baseWhere, &baseParam, "server_data.external_operatorcode = ?", req.AdditionalParamOperator)
	// lib.AppendWhere(&baseWhere, &baseParam, "server_data.external_route = ?", req.AdditionalParamRoute)
	// if req.IgnoreInvoiceID == false {
	// 	lib.AppendWhereRaw(&baseWhere, "server_data.invoice_id IS NULL")
	// }

	var partQueryExternalAccontIDIn string
	if req.InvoiceTypeID == "5" {
		partQueryExternalAccontIDIn = "external_account_id"
	} else if len(req.InvoiceTypeID) > 0 {
		partQueryExternalAccontIDIn = "external_rootparent_account"
	}
	if len(req.ListExternalRootParentAccount) > 0 {
		var baseIn string
		for _, extrootaccountid := range req.ListExternalRootParentAccount {
			lib.AppendComma(&baseIn, &baseParam, "?", extrootaccountid)
		}
		lib.AppendWhereRaw(&baseWhere, partQueryExternalAccontIDIn+" IN ("+baseIn+")")
	}

	if len(req.ListUsername) > 0 {
		var baseIn string
		for _, username := range req.ListUsername {
			lib.AppendComma(&baseIn, &baseParam, "?", username)
		}
		lib.AppendWhereRaw(&baseWhere, "external_username IN ("+baseIn+")")
	}
	if len(req.ListSender) > 0 {
		var baseIn string
		for _, sender := range req.ListSender {
			lib.AppendComma(&baseIn, &baseParam, "?", sender)
		}
		lib.AppendWhereRaw(&baseWhere, "external_sender IN ("+baseIn+")")
	}

	if len(req.ListServerID) > 0 {
		if len(req.ListServerID) == 1 && req.ListServerID[0] != "*" {
			lib.AppendWhere(&baseWhere, &baseParam, "server_data.server_id = ?", req.ListServerID[0])
		}
		if len(req.ListServerID) > 1 {
			var listServer string
			for _, each := range req.ListServerID {
				lib.AppendComma(&listServer, &baseParam, "?", each)
			}
			lib.AppendWhereRaw(&baseWhere, "server_data.server_id IN ("+listServer+")")
		}
	}

	if len(req.AdditionalParamOperator) > 0 {
		var listOperator string
		for _, each := range req.AdditionalParamOperator {
			lib.AppendComma(&listOperator, &baseParam, "?", each)
		}
		lib.AppendWhereRaw(&baseWhere, "server_data.external_operatorcode IN ("+listOperator+")")
	}

	if len(req.AdditionalParamRoute) > 0 {
		var listOperator string
		for _, each := range req.AdditionalParamRoute {
			lib.AppendComma(&listOperator, &baseParam, "?", each)
		}
		lib.AppendWhereRaw(&baseWhere, "server_data.external_route IN ("+listOperator+")")
	}

	resultCurrencey, _, errGetCurrency := conn.DBOcsConn.SelectQueryByFieldNameSlice("SELECT balance_type, balance_name, exponent, balance_category FROM ocs.balance WHERE balance_category = ?", "C")
	if errGetCurrency != nil {
		return result, errGetCurrency
	}

	currencyIn := ""
	for i, each := range resultCurrencey {
		currencyIn = currencyIn + "'" + each["balance_type"] + "'"
		if i != len(resultCurrencey)-1 {
			currencyIn = currencyIn + ", "
		}
	}

	var partQryPrice string
	if req.UseBillingPrice == true {
		partQryPrice = `item_price.price as price, `
	} else {
		partQryPrice = `IF((server_data.external_invoice_price IS NOT NULL && server_data.external_invoice_price <> 0 && server_data.external_balance_type in (` + currencyIn + `) ),
		server_data.external_invoice_price, item_price.price) as price,`
	}

	// runQuery := "SELECT server_data_id, server_id, server_account, item_id, account_id, external_smscount,external_transdate, external_transcount, invoice_id FROM server_data "
	runQuery := `SELECT server_data.server_data_id, server_data.server_id, server_data.external_account_id, 
	server_data.external_rootparent_account, server_data.item_id, server_data.account_id, server_data.external_smscount,server_data.external_transdate, 
	server_data.external_transcount, server_data.external_balance_type, server_data.invoice_id,server_data.external_user_id, 
	server_data.external_sender, IFNULL(mapping.value,server_data.external_operatorcode) as external_operatorcode, server_data.external_route,
	server_data.external_username, 
	item.item_id as tblitem_item_id, item.item_name, item.uom, item.category, item_price.item_id as tblitem_price_item_id,
	` + partQryPrice + `
	item_price.server_id as tblitem_price_server_id, item_price.company_id as tblitem_price_company_id, 
	server.server_name as tblserver_server_name, IFNULL(mapping_operator.new_route,server_data.external_route) as new_route, 
	server_data.external_operatorcode as ori_external_operatorcode, server_data.external_route as ori_external_route
	FROM server_data 
	LEFT JOIN account ON account.account_id = server_data.account_id
	LEFT JOIN server ON server.server_id = server_data.server_id
	LEFT JOIN mapping ON server_data.external_operatorcode = mapping.key AND mapping.type="operator"
	LEFT JOIN mapping_operator ON IFNULL(mapping.value,server_data.external_operatorcode) = mapping_operator.operatorcode AND server_data.external_route = mapping_operator.route 
	LEFT JOIN item ON IFNULL(mapping.value,server_data.external_operatorcode)=item.operator AND IFNULL(mapping_operator.new_route,server_data.external_route) = item.route AND item.category = ?
	LEFT JOIN item_price ON item.item_id=item_price.item_id AND server_data.server_id=item_price.server_id 
	AND item_price.company_id = account.company_id AND item_price.currency_code = ?  
	 `
	if len(baseWhere) > 0 {
		runQuery += " WHERE " + baseWhere
	}
	log.Info("final query:", runQuery)
	log.Info("get server data query param: ", baseParam)
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	resultQuery, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)

	log.Info("Len Server Data - ", len(resultQuery))
	for _, each := range resultQuery {
		if len(each["tblitem_item_id"]) == 0 {
			logrus.Info("item id kosong - ", each)
		}
		single := make(map[string]interface{})
		single["server_data_id"] = each["server_data_id"]
		single["account_id"] = each["account_id"]
		single["server_id"] = each["server_id"]
		single["external_account_id"] = each["external_account_id"]
		single["external_rootparent_account"] = each["external_rootparent_account"]
		single["item_id"] = each["tblitem_item_id"]
		single["account_id"] = each["account_id"]
		single["external_transdate"] = each["external_transdate"]
		single["external_price"] = each["price"]
		single["external_balance_type"] = each["external_balance_type"]
		single["external_user_id"] = each["external_user_id"]
		single["external_username"] = each["external_username"]
		single["external_sender"] = each["external_sender"]
		single["external_operatorcode"] = each["external_operatorcode"]
		single["ori_external_operatorcode"] = each["ori_external_operatorcode"]
		single["ori_external_route"] = each["ori_external_route"]
		single["external_route"] = each["external_route"]
		single["external_smscount"] = each["external_smscount"]
		single["external_transcount"] = each["external_transcount"]
		single["invoice_id"] = each["invoice_id"]
		single["new_route"] = each["new_route"]
		// single["created_date"] = each["created_date"]

		// findItem, _, errFindItem := conn.DBAppConn.SelectQueryByFieldNameSlice("SELECT item_id, item_name, operator, route, category, uom FROM item WHERE item_id = ?", single["item_id"])
		// if errFindItem != nil {
		// 	return nil, errFindItem
		// }
		item := make(map[string]interface{})
		item["item_id"] = each["tblitem_item_id"]
		item["item_name"] = each["item_name"]
		item["category"] = each["	"]
		item["uom"] = each["uom"]

		server := make(map[string]interface{})
		server["server_id"] = each["server_id"]
		server["server_name"] = each["tblserver_server_name"]

		single["server"] = server

		// findItemPrice, _, errFindItemPrice := conn.DBAppConn.SelectQueryByFieldNameSlice("SELECT item_id,account_id,server_id,price FROM item_price WHERE item_id = ? AND account_id = ? AND server_id = ?", item["item_id"], req.AccountID, req.ServerID)
		// if errFindItemPrice != nil {
		// 	return nil, errFindItemPrice
		// }

		itemPrice := make(map[string]interface{})
		itemPrice["price"] = each["price"]
		itemPrice["item_id"] = each["tblitem_price_item_id"]
		itemPrice["server_id"] = each["tblitem_price_server_id"]
		itemPrice["company_id"] = each["tblitem_price_company_id"]
		item["item_price"] = itemPrice

		single["item"] = item
		// var category
		result = append(result, single)
	}

	return result, err
}

func GetSenderFromRequest(conn *connections.Connections, req datastruct.ServerDataRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
	lib.AppendWhere(&baseWhere, &baseParam, "external_rootparent_account = ?", req.ExternalRootParentAccount)
	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(external_transdate, '%Y%m') = ?", req.MonthUse)

	runQuery := "SELECT distinct(external_sender) FROM server_data "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)
	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func GetAccountFromServerData(conn *connections.Connections, req datastruct.ServerDataRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	// lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
	// lib.AppendWhere(&baseWhere, &baseParam, "external_rootparent_account = ?", req.ExternalRootParentAccount)
	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(external_transdate, '%Y%m') = ?", req.MonthUse)
	lib.AppendWhereRaw(&baseWhere, "server_data.invoice_id IS NULL")

	if len(req.ListAccountID) > 0 {
		var baseIn string
		for _, accountid := range req.ListAccountID {
			lib.AppendComma(&baseIn, &baseParam, "?", accountid)
		}
		lib.AppendWhereRaw(&baseWhere, "server_data.account_id IN ("+baseIn+")")
	}

	runQuery := "SELECT account.account_id, account.name, account.account_type, account.company_id, account.billing_type FROM server_data  LEFT JOIN account ON server_data.account_id = account.account_id "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	runQuery += " GROUP BY account_id "
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func GetUserFromServerData(conn *connections.Connections, req datastruct.ServerDataRequest) ([]map[string]string, error) {
	var result []map[string]string
	var err error

	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
	var baseWhere string
	var baseParam []interface{}

	// lib.AppendWhere(&baseWhere, &baseParam, "account_id = ?", req.AccountID)
	// lib.AppendWhere(&baseWhere, &baseParam, "external_rootparent_account = ?", req.ExternalRootParentAccount)
	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(external_transdate, '%Y%m') = ?", req.MonthUse)
	// lib.AppendWhereRaw(&baseWhere, "server_data.invoice_id IS NULL")

	if len(req.ListAccountID) > 0 {
		var baseIn string
		for _, accountid := range req.ListAccountID {
			lib.AppendComma(&baseIn, &baseParam, "?", accountid)
		}
		lib.AppendWhereRaw(&baseWhere, "server_data.account_id IN ("+baseIn+")")
	}

	runQuery := "SELECT account.account_id, account.name, server_data.external_user_id, server_data.external_username FROM server_data  LEFT JOIN account ON server_data.account_id = account.account_id "
	if len(baseWhere) > 0 {
		runQuery += "WHERE " + baseWhere
	}
	runQuery += " GROUP BY 1,2,3,4 "
	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

	result, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)
	return result, err
}

func InsertServerData(conn *connections.Connections, req datastruct.ServerDataRequest) error {
	var err error

	// -- THIS IS BASIC INSERT EXAMPLE
	// var baseIn string
	// var baseParam []interface{}

	// lib.AppendComma(&baseIn, &baseParam, "?", req.ServerDataID)
	// lib.AppendComma(&baseIn, &baseParam, "?", req.ServerDataName)

	// qry := "INSERT INTO server (serverid, servername) VALUES (" + baseIn + ")"
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)

	return err
}

func UpdateServerData(conn *connections.Connections, req datastruct.ServerDataRequest) error {
	var err error

	// -- THIS IS BASIC UPDATE EXAMPLE
	// var baseUp string
	// var baseParam []interface{}

	// lib.AppendComma(&baseUp, &baseParam, "servername = ?", req.ServerDataName)
	// qry := "UPDATE server SET " + baseUp + " WHERE serverid = ?"
	// baseParam = append(baseParam, req.ServerDataID)
	// _, _, err = conn.DBAppConn.Exec(qry, baseParam...)
	return err
}

func DeleteServerData(conn *connections.Connections, req datastruct.ServerDataRequest) error {
	var err error
	// -- THIS IS DELETE LOGIC EXAMPLE
	// qry := "DELETE FROM server WHERE serverid = ?"
	// _, _, err = conn.DBAppConn.Exec(qry, req.ServerDataID)
	return err
}

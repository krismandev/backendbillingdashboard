package models

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/server-data/datastruct"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetSingleServerData(serverDataID string, conn *connections.Connections, req datastruct.ServerDataRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(server-dataID) == 0 {
	// 	server-dataID = req.ServerDataID
	// }
	// query := "SELECT server-dataid, server-dataname FROM server-data WHERE server-dataid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, server-dataID)
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

// func GetServerData(conn *connections.Connections, req datastruct.ServerDataRequest) ([]map[string]interface{}, error) {
// 	var result []map[string]interface{}
// 	var resultQuery []map[string]string
// 	var err error

// 	// -- THIS IS BASIC GET REQUEST EXAMPLE LOGIC
// 	var baseWhere string
// 	var baseParam []interface{}

// 	baseParam = append(baseParam, req.CurrencyCode)
// 	lib.AppendWhere(&baseWhere, &baseParam, "server_data_id = ?", req.ServerDataID)
// 	lib.AppendWhere(&baseWhere, &baseParam, "server_data.server_id = ?", req.ServerID)
// 	lib.AppendWhere(&baseWhere, &baseParam, "server_data.account_id = ?", req.AccountID)
// 	lib.AppendWhere(&baseWhere, &baseParam, "server_data.external_sender = ?", req.ExternalSender)
// 	lib.AppendWhere(&baseWhere, &baseParam, "server_data.external_rootparent_account = ?", req.ExternalRootParentAccount)
// 	lib.AppendWhere(&baseWhere, &baseParam, "DATE_FORMAT(server_data.external_transdate, '%Y%m') = ?", req.MonthUse)
// 	lib.AppendWhereRaw(&baseWhere, "server_data.invoice_id IS NULL")
// 	// lib.AppendWhere(&baseWhere, &baseParam, "item_price.currency_code = ?", req.CurrencyCode)

// 	var partQueryIn string
// 	if req.InvoiceTypeID == "5" {
// 		partQueryIn = "external_account_id"
// 	} else {
// 		partQueryIn = "external_rootparent_account"
// 	}
// 	if len(req.ListExternalRootParentAccount) > 0 {
// 		var baseIn string
// 		for _, extrootaccountid := range req.ListExternalRootParentAccount {
// 			lib.AppendComma(&baseIn, &baseParam, "?", extrootaccountid)
// 		}
// 		lib.AppendWhereRaw(&baseWhere, partQueryIn+" IN ("+baseIn+")")
// 	}

// 	if len(req.ListUserID) > 0 {
// 		var baseIn string
// 		for _, userid := range req.ListUserID {
// 			lib.AppendComma(&baseIn, &baseParam, "?", userid)
// 		}
// 		lib.AppendWhereRaw(&baseWhere, "external_user_id IN ("+baseIn+")")
// 	}

// 	resultCurrencey, _, errGetCurrency := conn.DBOcsConn.SelectQueryByFieldNameSlice("SELECT balance_type, balance_name, exponent, balance_category FROM ocs.balance WHERE balance_category = ?", "C")
// 	if errGetCurrency != nil {
// 		return result, errGetCurrency
// 	}

// 	currencyIn := ""
// 	for i, each := range resultCurrencey {
// 		currencyIn = currencyIn + "'" + each["balance_type"] + "'"
// 		if i != len(resultCurrencey)-1 {
// 			currencyIn = currencyIn + ", "
// 		}
// 	}

// 	var partQryPrice string
// 	if req.UseBillingPrice == true {
// 		partQryPrice = `item_price.price, `
// 	} else {
// 		partQryPrice = `IF((server_data.external_price IS NOT NULL && server_data.external_price <> 0 && server_data.external_balance_type in (` + currencyIn + `) ),
// 		server_data.external_price, item_price.price) as price,`
// 	}

// 	log.Info("LihatCurrency-", currencyIn)
// 	// runQuery := "SELECT server_data_id, server_id, server_account, item_id, account_id, external_smscount,external_transdate, external_transcount, invoice_id FROM server_data "
// 	runQuery := `SELECT server_data.server_data_id, server_data.server_id, server_data.external_account_id,
// 	server_data.external_rootparent_account, server_data.item_id, server_data.account_id, server_data.external_smscount,server_data.external_transdate,
// 	server_data.external_transcount, server_data.external_balance_type, server_data.invoice_id,server_data.external_user_id,
// 	server_data.external_sender, IFNULL(mapping.value,server_data.external_operatorcode) as external_operatorcode, server_data.external_route,
// 	server_data.external_username,
// 	item.item_id as tblitem_item_id, item.item_name, item.uom, item.category, item_price.item_id as tblitem_price_item_id,
// 	` + partQryPrice + `
// 	item_price.server_id as tblitem_price_server_id, item_price.account_id as tblitem_price_account_id,
// 	server.server_name as tblserver_server_name, mapping_operator.new_route
// 	FROM server_data
// 	LEFT JOIN item ON server_data.item_id=item.item_id
// 	LEFT JOIN server ON server.server_id = server_data.server_id
// 	LEFT JOIN item_price ON item.item_id=item_price.item_id AND server_data.server_id=item_price.server_id
// 		AND server_data.account_id=item_price.account_id AND item_price.currency_code = ?
// 	LEFT JOIN mapping_operator ON server_data.external_operatorcode = mapping_operator.operatorcode AND server_data.external_route = mapping_operator.route
// 	LEFT JOIN mapping ON server_data.external_operatorcode = mapping.key AND mapping.type="operator" `
// 	if len(baseWhere) > 0 {
// 		runQuery += " WHERE " + baseWhere
// 	}
// 	log.Info("final query:", runQuery)
// 	lib.AppendOrderBy(&runQuery, req.Param.OrderBy, req.Param.OrderDir)
// 	lib.AppendLimit(&runQuery, req.Param.Page, req.Param.PerPage)

// 	resultQuery, _, err = conn.DBAppConn.SelectQueryByFieldNameSlice(runQuery, baseParam...)

// 	for _, each := range resultQuery {
// 		single := make(map[string]interface{})
// 		single["server_data_id"] = each["server_data_id"]
// 		single["account_id"] = each["account_id"]
// 		single["server_id"] = each["server_id"]
// 		single["external_account_id"] = each["external_account_id"]
// 		single["external_rootparent_account"] = each["external_rootparent_account"]
// 		single["item_id"] = each["item_id"]
// 		single["account_id"] = each["account_id"]
// 		single["external_transdate"] = each["external_transdate"]
// 		single["external_price"] = each["price"]
// 		single["external_balance_type"] = each["external_balance_type"]
// 		single["external_user_id"] = each["external_user_id"]
// 		single["external_username"] = each["external_username"]
// 		single["external_sender"] = each["external_sender"]
// 		single["external_operatorcode"] = each["external_operatorcode"]
// 		single["external_route"] = each["external_route"]
// 		single["external_smscount"] = each["external_smscount"]
// 		single["external_transcount"] = each["external_transcount"]
// 		single["invoice_id"] = each["invoice_id"]
// 		single["new_route"] = each["new_route"]
// 		// single["created_date"] = each["created_date"]

// 		// findItem, _, errFindItem := conn.DBAppConn.SelectQueryByFieldNameSlice("SELECT item_id, item_name, operator, route, category, uom FROM item WHERE item_id = ?", single["item_id"])
// 		// if errFindItem != nil {
// 		// 	return nil, errFindItem
// 		// }
// 		item := make(map[string]interface{})
// 		item["item_id"] = each["tblitem_item_id"]
// 		item["item_name"] = each["item_name"]
// 		item["category"] = each["category"]
// 		item["uom"] = each["uom"]

// 		server := make(map[string]interface{})
// 		server["server_id"] = each["server_id"]
// 		server["server_name"] = each["tblserver_server_name"]

// 		single["server"] = server

// 		// findItemPrice, _, errFindItemPrice := conn.DBAppConn.SelectQueryByFieldNameSlice("SELECT item_id,account_id,server_id,price FROM item_price WHERE item_id = ? AND account_id = ? AND server_id = ?", item["item_id"], req.AccountID, req.ServerID)
// 		// if errFindItemPrice != nil {
// 		// 	return nil, errFindItemPrice
// 		// }

// 		itemPrice := make(map[string]interface{})
// 		itemPrice["price"] = each["price"]
// 		itemPrice["item_id"] = each["tblitem_price_item_id"]
// 		itemPrice["server_id"] = each["tblitem_price_server_id"]
// 		itemPrice["account_id"] = each["tblitem_price_account_id"]
// 		item["item_price"] = itemPrice

// 		single["item"] = item
// 		// var category
// 		result = append(result, single)
// 	}

// 	return result, err
// }

func LoadServerData(conn *connections.Connections, req datastruct.ServerDataRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var resultQuery []map[string]string
	var err error

	// runQuery := "SELECT server_data_id, server_id, server_account, item_id, account_id, external_smscount,external_transdate, external_transcount, invoice_id FROM server_data "
	runQuery := "SELECT LPAD(FLOOR(RAND() * 999999.99), 6, '0') as data_id,'11' as server_id,b.accountid as server_account,d.item_id,b.accountid as server_account,a.transdate,b.user_id,a.sender,a.operatorcode,c.category,sum(a.smscount),sum(a.transcount),null as price,now() as created_date FROM `bulksms_log_daily` a left join smsgw.smsgw_user b on a.userid=b.user_id left join smsgw.user_to_application c on b.user_id=c.user_id  left join dbbilling.item d on d.operator=a.operatorcode and d.route=c.category  and c.application_id=1 and d.category='USAGE' left join dbbilling.server_account e on e.external_account_id=b.accountid and e.server_id=@serverid where ((resultcode IN (0,1,2,3,5,4002,4003,4004,4005,4006,4007,4008, 4009,4010,4068) and smscid not in (35,36,143,144)) or (resultcode IN (0,1,2,3,5) and smscid in (35,36,143,144))) and c.application_id=1 and d.category='USAGE' and date_format(a.transdate,'%Y%m') = '202204' group by b.accountid,d.item_id,e.account_id,a.transdate,b.user_name,a.sender,a.operatorcode,c.category"

	resultQuery, _, err = conn.DBDashbConn.SelectQueryByFieldNameSlice(runQuery)

	bulkInserQuery := "INSERT IGNORE INTO server_data(server_data_id,server_id,server_account,item_id,account_id,external_transdate,external_user_id,external_sender,external_operatorcode,external_route,external_smscount,external_transcount,external_price,created_date) VALUES "
	var paramsBulkInsert []interface{}
	var stringGroup []string

	for _, each := range resultQuery {
		partquery := "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
		paramsBulkInsert = append(paramsBulkInsert, each["data_id"])
		paramsBulkInsert = append(paramsBulkInsert, each["server_id"])
		paramsBulkInsert = append(paramsBulkInsert, each["server_account"])
		paramsBulkInsert = append(paramsBulkInsert, each["item_id"])
		paramsBulkInsert = append(paramsBulkInsert, each["server_account"])
		paramsBulkInsert = append(paramsBulkInsert, each["transdate"])
		paramsBulkInsert = append(paramsBulkInsert, each["user_id"])
		paramsBulkInsert = append(paramsBulkInsert, each["sender"])
		paramsBulkInsert = append(paramsBulkInsert, each["operatorcode"])
		paramsBulkInsert = append(paramsBulkInsert, each["category"])
		paramsBulkInsert = append(paramsBulkInsert, each["sum(a.smscount)"])
		paramsBulkInsert = append(paramsBulkInsert, each["sum(a.transcount)"])
		paramsBulkInsert = append(paramsBulkInsert, each["price"])
		paramsBulkInsert = append(paramsBulkInsert, each["created_date"])
		stringGroup = append(stringGroup, partquery)
	}

	final_query := bulkInserQuery + strings.Join(stringGroup, ", ")
	log.Info("FinalQuery", final_query)

	_, _, errInsert := conn.DBAppConn.Exec(final_query, paramsBulkInsert...)
	if errInsert != nil {
		return nil, err
	}

	return result, err
}

func CheckServerDataExists(serverDataID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(server-dataid) FROM server-data WHERE server-dataid = ?"
	// param = append(param, server-dataID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("ServerData ID is not exists")
	// }
	return nil
}

func CheckServerDataDuplicate(exceptID string, conn *connections.Connections, req datastruct.ServerDataRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(server-dataid) FROM server-data WHERE server-dataid = ?"
	// param = append(param, req.ServerDataID)
	// if len(exceptID) > 0 {
	// 	qry += " AND server-dataid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("ServerData ID is already exists. Please use another ServerData ID")
	// }
	return nil
}

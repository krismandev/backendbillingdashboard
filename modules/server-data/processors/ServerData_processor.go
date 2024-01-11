package processors

import (
	"backendbillingdashboard/connections"
	accountDt "backendbillingdashboard/modules/account/datastruct"
	"backendbillingdashboard/modules/server-data/datastruct"
	"backendbillingdashboard/modules/server-data/models"
	dtServer "backendbillingdashboard/modules/server/datastruct"
)

func GetListServerData(conn *connections.Connections, req datastruct.ServerDataRequest) ([]datastruct.ServerDataDataStruct, error) {
	var output []datastruct.ServerDataDataStruct
	var err error

	// grab mapping data from model
	serverDataList, err := models.GetServerDataFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, serverData := range serverDataList {
		single := CreateSingleServerDataStruct(serverData)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleServerDataStruct(serverData map[string]interface{}) datastruct.ServerDataDataStruct {
	var single datastruct.ServerDataDataStruct
	single.ServerDataID, _ = serverData["server_data_id"].(string)
	single.ServerID, _ = serverData["server_id"].(string)
	single.ExternalAccountID, _ = serverData["external_account_id"].(string)
	single.ItemID, _ = serverData["item_id"].(string)
	single.AccountID, _ = serverData["account_id"].(string)
	single.ExternalTransdate, _ = serverData["external_transdate"].(string)
	single.ExternalRootParentAccount, _ = serverData["external_rootparent_account"].(string)
	single.ExternalPrice, _ = serverData["external_price"].(string)
	single.ExternalUserID, _ = serverData["external_user_id"].(string)
	single.ExternalUsername, _ = serverData["external_username"].(string)
	single.ExternalSender, _ = serverData["external_sender"].(string)
	single.ExternalOperatorCode, _ = serverData["external_operatorcode"].(string)
	single.OriExternalOperatorCode, _ = serverData["ori_external_operatorcode"].(string)
	single.ExternalRoute, _ = serverData["external_route"].(string)
	single.OriExternalRoute, _ = serverData["ori_external_route"].(string)
	single.ExternalSMSCount, _ = serverData["external_smscount"].(string)
	single.ExternalTransCount, _ = serverData["external_transcount"].(string)
	single.ExternalBalanceType, _ = serverData["external_balance_type"].(string)
	single.InvoiceID, _ = serverData["invoice_id"].(string)
	single.NewRoute, _ = serverData["new_route"].(string)

	itemPrice := datastruct.ItemPriceDataStruct{
		ItemID:    serverData["item"].(map[string]interface{})["item_price"].(map[string]interface{})["item_id"].(string),
		CompanyID: serverData["item"].(map[string]interface{})["item_price"].(map[string]interface{})["company_id"].(string),
		ServerID:  serverData["item"].(map[string]interface{})["item_price"].(map[string]interface{})["server_id"].(string),
		Price:     serverData["item"].(map[string]interface{})["item_price"].(map[string]interface{})["price"].(string),
	}

	item := datastruct.ItemDataStruct{
		ItemID:    serverData["item"].(map[string]interface{})["item_id"].(string),
		ItemName:  serverData["item"].(map[string]interface{})["item_name"].(string),
		UOM:       serverData["item"].(map[string]interface{})["uom"].(string),
		Category:  serverData["item"].(map[string]interface{})["category"].(string),
		ItemPrice: itemPrice,
	}

	server := dtServer.ServerDataStruct{
		ServerID:   serverData["server"].(map[string]interface{})["server_id"].(string),
		ServerName: serverData["server"].(map[string]interface{})["server_name"].(string),
	}

	single.Item = item
	single.Server = server
	return single
}

func InsertServerData(conn *connections.Connections, req datastruct.ServerDataRequest) (datastruct.ServerDataDataStruct, error) {
	var output datastruct.ServerDataDataStruct
	var err error

	err = models.InsertServerData(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single server-data
	// single, err := models.GetSingleServerData(req.ServerDataID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleServerDataStruct(single)
	return output, err
}

func UpdateServerData(conn *connections.Connections, req datastruct.ServerDataRequest) (datastruct.ServerDataDataStruct, error) {
	var output datastruct.ServerDataDataStruct
	var err error

	err = models.UpdateServerData(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single server-data
	// single, err := models.GetSingleServerData(req.ServerDataID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleServerDataStruct(single)
	return output, err
}

func DeleteServerData(conn *connections.Connections, req datastruct.ServerDataRequest) error {
	err := models.DeleteServerData(conn, req)
	return err
}

func GetListSender(conn *connections.Connections, req datastruct.ServerDataRequest) ([]datastruct.SenderDataStruct, error) {
	var output []datastruct.SenderDataStruct
	var err error

	// grab mapping data from model
	senderList, err := models.GetSenderFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, sender := range senderList {
		single := CreateSingleSenderStruct(sender)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleSenderStruct(sender map[string]string) datastruct.SenderDataStruct {
	var single datastruct.SenderDataStruct
	single.Sender = sender["external_sender"]
	return single
}

func GetListAccount(conn *connections.Connections, req datastruct.ServerDataRequest) ([]accountDt.AccountDataStruct, error) {
	var output []accountDt.AccountDataStruct
	var err error

	// grab mapping data from model
	accountList, err := models.GetAccountFromServerData(conn, req)
	if err != nil {
		return output, err
	}

	for _, account := range accountList {
		single := CreateSingleAccountStruct(account)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleAccountStruct(sender map[string]string) accountDt.AccountDataStruct {
	var single accountDt.AccountDataStruct
	single.AccountID = sender["account_id"]
	single.Name = sender["name"]
	single.AccountType = sender["account_type"]
	single.BillingType = sender["billing_type"]
	single.CompanyID = sender["company_id"]
	return single
}

func GetListUser(conn *connections.Connections, req datastruct.ServerDataRequest) ([]datastruct.UserDataStruct, error) {
	var output []datastruct.UserDataStruct
	var err error

	// grab mapping data from model
	userList, err := models.GetUserFromServerData(conn, req)
	if err != nil {
		return output, err
	}

	for _, account := range userList {
		single := CreateSingleUserStruct(account)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleUserStruct(sender map[string]string) datastruct.UserDataStruct {
	var single datastruct.UserDataStruct
	single.AccountID = sender["account_id"]
	single.AccountName = sender["name"]
	single.ExternalUserID = sender["external_user_id"]
	single.ExternalUserName = sender["external_username"]
	return single
}

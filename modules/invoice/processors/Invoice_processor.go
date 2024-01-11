package processors

import (
	"backendbillingdashboard/config"
	"backendbillingdashboard/connections"
	"backendbillingdashboard/lib"
	dtAccount "backendbillingdashboard/modules/account/datastruct"
	dtCompany "backendbillingdashboard/modules/company/datastruct"
	companyModel "backendbillingdashboard/modules/company/models"
	dtInvoiceGroup "backendbillingdashboard/modules/invoice-group/datastruct"
	dtInvoiceType "backendbillingdashboard/modules/invoice-type/datastruct"
	invoiceTypeModel "backendbillingdashboard/modules/invoice-type/models"
	"backendbillingdashboard/modules/invoice/datastruct"
	"backendbillingdashboard/modules/invoice/models"
	proformaInvoiceModel "backendbillingdashboard/modules/proforma-invoice/models"
	dtServerAccount "backendbillingdashboard/modules/server-account/datastruct"
	serverAccountModel "backendbillingdashboard/modules/server-account/models"
	dtServerData "backendbillingdashboard/modules/server-data/datastruct"
	serverDataModel "backendbillingdashboard/modules/server-data/models"
	serverDataProcessor "backendbillingdashboard/modules/server-data/processors"
	dtServer "backendbillingdashboard/modules/server/datastruct"
	"encoding/base64"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func GetListInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) ([]datastruct.InvoiceDataStruct, error) {
	var output []datastruct.InvoiceDataStruct
	var err error

	// grab mapping data from model
	invoiceList, err := models.GetInvoiceFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, invoice := range invoiceList {
		single := CreateSingleInvoiceStruct(invoice)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleInvoiceStruct(invoice map[string]interface{}) datastruct.InvoiceDataStruct {
	var single datastruct.InvoiceDataStruct
	single.InvoiceID, _ = invoice["invoice_id"].(string)
	single.InvoiceNo, _ = invoice["invoice_no"].(string)
	single.InvoiceDate, _ = invoice["invoice_date"].(string)
	single.InvoiceStatus, _ = invoice["invoicestatus"].(string)
	single.CompanyID, _ = invoice["company_id"].(string)
	single.MonthUse, _ = invoice["month_use"].(string)
	single.InvoiceTypeID, _ = invoice["inv_type_id"].(string)
	single.PrintCounter, _ = invoice["printcounter"].(string)
	single.PPN, _ = invoice["ppn"].(string)
	single.Note, _ = invoice["note"].(string)
	single.CancelDesc, _ = invoice["canceldesc"].(string)
	single.LastPrintUsername, _ = invoice["last_print_username"].(string)
	single.LastPrintDate, _ = invoice["last_print_date"].(string)
	single.CreatedAt, _ = invoice["created_at"].(string)
	single.CreatedBy, _ = invoice["created_by"].(string)
	single.LastUpdateUsername, _ = invoice["last_update_username"].(string)
	single.Attachment, _ = invoice["attachment"].(string)
	single.LastUpdateDate, _ = invoice["last_update_date"].(string)
	single.DiscountType, _ = invoice["discount_type"].(string)
	single.Discount, _ = invoice["discount"].(string)
	single.Paid, _ = invoice["paid"].(string)
	single.PaymentMethod, _ = invoice["payment_method"].(string)
	single.ExchangeRateDate, _ = invoice["exchange_rate_date"].(string)
	single.DueDate, _ = invoice["due_date"].(string)
	single.GrandTotal, _ = invoice["grand_total"].(string)
	single.PPNAmount, _ = invoice["ppn_amount"].(string)
	single.Sender, _ = invoice["sender"].(string)
	single.TaxInvoice, _ = invoice["tax_invoice"].(string)
	single.AdjustmentNote, _ = invoice["adjustment_note"].(string)
	single.ReceivedDate, _ = invoice["received_date"].(string)
	single.ReceiptLetterAttachment, _ = invoice["receipt_letter_attachment"].(string)

	var invoiceType datastruct.InvoiceTypeDataStruct
	invoiceType.InvoiceTypeID = invoice["invoice_type"].(map[string]interface{})["inv_type_id"].(string)
	invoiceType.InvoiceTypeName = invoice["invoice_type"].(map[string]interface{})["inv_type_name"].(string)
	invoiceType.ServerID = invoice["invoice_type"].(map[string]interface{})["server_id"].(string)
	invoiceType.Category = invoice["invoice_type"].(map[string]interface{})["category"].(string)
	invoiceType.LoadFromServer = invoice["invoice_type"].(map[string]interface{})["load_from_server"].(string)
	invoiceType.IsGroup = invoice["invoice_type"].(map[string]interface{})["is_group"].(string)
	invoiceType.GroupType = invoice["invoice_type"].(map[string]interface{})["group_type"].(string)
	invoiceType.CurrencyCode = invoice["invoice_type"].(map[string]interface{})["currency_code"].(string)

	single.InvoiceType = invoiceType

	var company dtCompany.CompanyDataStruct
	company.CompanyID = invoice["company"].(map[string]interface{})["company_id"].(string)
	company.Name = invoice["company"].(map[string]interface{})["name"].(string)
	company.Address1 = invoice["company"].(map[string]interface{})["address1"].(string)
	company.Address2 = invoice["company"].(map[string]interface{})["address2"].(string)
	company.City = invoice["company"].(map[string]interface{})["city"].(string)
	company.Desc = invoice["company"].(map[string]interface{})["desc"].(string)
	company.ContactPerson = invoice["company"].(map[string]interface{})["contact_person"].(string)
	company.ContactPersonPhone = invoice["company"].(map[string]interface{})["contact_person_phone"].(string)

	single.Company = company

	var tampungDetail []datastruct.InvoiceDetailStruct
	for _, eachDetail := range invoice["list_invoice_detail"].([]map[string]interface{}) {
		var detail datastruct.InvoiceDetailStruct
		detail.InvoiceDetailID = eachDetail["invoice_detail_id"].(string)
		detail.InvoiceID = eachDetail["invoice_id"].(string)
		detail.ItemID = eachDetail["itemid"].(string)
		detail.ItemPrice = eachDetail["item_price"].(string)
		detail.AccountID = eachDetail["account_id"].(string)
		detail.Qty = eachDetail["qty"].(string)
		detail.Uom = eachDetail["uom"].(string)
		detail.Note = eachDetail["note"].(string)
		detail.BalanceType = eachDetail["balance_type"].(string)
		detail.ServerID = eachDetail["server_id"].(string)
		detail.ExternalUserID = eachDetail["external_user_id"].(string)
		detail.ExternalUsername = eachDetail["external_username"].(string)
		detail.ExternalSender = eachDetail["external_sender"].(string)
		detail.Adjustment = eachDetail["adjustment"].(string)
		detail.AdjustmentConfirmationUsername = eachDetail["adjustment_confirmation_username"].(string)
		detail.AdjustmentConfirmationDate = eachDetail["adjustment_confirmation_date"].(string)

		var item datastruct.ItemDataStruct
		item.ItemID = eachDetail["item"].(map[string]interface{})["item_id"].(string)
		item.ItemName = eachDetail["item"].(map[string]interface{})["item_name"].(string)
		item.Operator = eachDetail["item"].(map[string]interface{})["operator"].(string)
		item.Route = eachDetail["item"].(map[string]interface{})["route"].(string)
		item.Category = eachDetail["item"].(map[string]interface{})["category"].(string)
		item.UOM = eachDetail["item"].(map[string]interface{})["uom"].(string)

		var server dtServer.ServerDataStruct
		server.ServerID = eachDetail["server"].(map[string]interface{})["server_id"].(string)
		server.ServerName = eachDetail["server"].(map[string]interface{})["server_name"].(string)

		var account dtAccount.AccountDataStruct
		account.AccountID = eachDetail["account"].(map[string]interface{})["account_id"].(string)
		account.Name = eachDetail["account"].(map[string]interface{})["name"].(string)
		account.AccountType = eachDetail["account"].(map[string]interface{})["account_type"].(string)
		account.CompanyID = eachDetail["account"].(map[string]interface{})["company_id"].(string)
		account.NonTaxable = eachDetail["account"].(map[string]interface{})["non_taxable"].(string)

		detail.Account = account
		detail.Item = item
		detail.Server = server

		tampungDetail = append(tampungDetail, detail)
	}
	single.ListInvoiceDetail = tampungDetail

	return single
}

func InsertInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	if len(req.Attachment) > 0 {
		b64data := req.Attachment[strings.IndexByte(req.Attachment, ',')+1:]
		dec, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			// panic(err)
			return err
		}

		if _, err := os.Stat(config.Param.AttachmentFolder); os.IsNotExist(err) {
			errMkdir := os.MkdirAll(config.Param.AttachmentFolder, os.ModePerm)
			if errMkdir != nil {
				logrus.Error("ERROR-", err.Error())
				return errMkdir
			}
			// TODO: handle error
		}

		fileName := "attachment_invoice_" + req.InvoiceNo + ".jpg"
		path := config.Param.AttachmentFolder + "/" + fileName
		err = os.WriteFile(path, dec, 0600)
		// f, err := os.Create(config.Param.AttachmentFolder + "/attachment_invoice_" + req.InvoiceNo + ".jpg")
		if err != nil {
			return err
		}
		req.Attachment = fileName
		// defer f.Close()

	}
	err = models.InsertInvoice(conn, req)
	if err != nil {
		return err
	}

	// jika tidak ada error, return single instance of single invoice
	// single, err := models.GetSingleInvoice(req.InvoiceID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleInvoiceStruct(single)
	return err
}

func InsertCustomInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	if len(req.Attachment) > 0 {
		b64data := req.Attachment[strings.IndexByte(req.Attachment, ',')+1:]
		dec, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			// panic(err)
			return err
		}

		if _, err := os.Stat(config.Param.AttachmentFolder); os.IsNotExist(err) {
			errMkdir := os.MkdirAll(config.Param.AttachmentFolder, os.ModePerm)
			if errMkdir != nil {
				logrus.Error("ERROR-", err.Error())
				return errMkdir
			}
			// TODO: handle error
		}

		fileName := "attachment_invoice_" + req.InvoiceNo + ".jpg"
		path := config.Param.AttachmentFolder + "/" + fileName
		err = os.WriteFile(path, dec, 0600)
		// f, err := os.Create(config.Param.AttachmentFolder + "/attachment_invoice_" + req.InvoiceNo + ".jpg")
		if err != nil {
			return err
		}
		req.Attachment = fileName
		// defer f.Close()

	}
	err = models.InsertCustomInvoice(conn, req)
	if err != nil {
		return err
	}

	// jika tidak ada error, return single instance of single invoice
	// single, err := models.GetSingleInvoice(req.InvoiceID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleInvoiceStruct(single)
	return err
}

func UpdateInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	if len(req.Attachment) > 0 {
		b64data := req.Attachment[strings.IndexByte(req.Attachment, ',')+1:]
		dec, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			// panic(err)
			return err
		}

		if _, err := os.Stat(config.Param.AttachmentFolder); os.IsNotExist(err) {
			errMkdir := os.MkdirAll(config.Param.AttachmentFolder, os.ModePerm)
			if errMkdir != nil {
				logrus.Error("ERROR-", err.Error())
				return errMkdir
			}
			// TODO: handle error
		}

		fileName := "attachment_invoice_" + req.InvoiceNo + ".jpg"
		path := config.Param.AttachmentFolder + "/" + fileName
		err = os.WriteFile(path, dec, 0600)
		// f, err := os.Create(config.Param.AttachmentFolder + "/attachment_invoice_" + req.InvoiceNo + ".jpg")
		if err != nil {
			return err
		}

		req.Attachment = fileName
		// defer f.Close()

	} else {
		req.Attachment = req.OldData.Attachment
	}
	err = models.UpdateInvoice(conn, req)
	if err != nil {
		return err
	}

	err = models.UpdateRevisionCounter(conn, req)
	if err != nil {
		return err
	}
	// jika tidak ada error, return single instance of single invoice
	// single, err := models.GetSingleInvoice(req.InvoiceID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleInvoiceStruct(single)

	return err
}

func GenerateInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	var generateInvoiceReq datastruct.GenerateInvoiceDataStruct
	var now = time.Now()
	generateTime := now.Format("2006-01-02 15:04:05")
	generateInvoiceReq.GenerateTime = generateTime
	generateInvoiceReq.LastUpdateUsername = req.LastUpdateUsername

	batchID, err := models.InsertGenerateInvoiceHistory(conn, generateInvoiceReq)
	if err != nil {
		logrus.Error("Error inserting generate invoice log ", err)
		return err
	}

	batchIDStr := strconv.Itoa(int(batchID))

	go GenerateInvoiceWorker(conn, req, batchIDStr)
	return err
}

func GenerateInvoiceWorker(conn *connections.Connections, req datastruct.InvoiceRequest, batchID string) {
	var err error
	listCompany, err := companyModel.GetCompanies(conn, req.ListCompanyID)
	var invoiceGroupRequest dtInvoiceGroup.InvoiceGroupRequest
	invoiceGroupRequest.Param.PerPage = 9999
	// listGroup, err := invoiceGroupModel.GetInvoiceGroupFromRequest(conn, invoiceGroupRequest)
	var serverAccountRequest dtServerAccount.ServerAccountRequest
	serverAccountRequest.ListAccountID = req.ListCompanyID
	serverAccountRequest.Param.PerPage = 9999
	listServerAccount, err := serverAccountModel.GetServerAccountFromRequest(conn, serverAccountRequest)

	var invoiceTypeRequest dtInvoiceType.InvoiceTypeRequest
	invoiceTypeRequest.Param.PerPage = 999
	listInvoiceType, err := invoiceTypeModel.GetInvoiceTypeFromRequest(conn, invoiceTypeRequest)

	numOfCompany := len(req.ListCompanyID)
	percentPerCompany := 100 / numOfCompany
	var progress int = 0
	if err != nil {
		logrus.Error("Error : ", err)
	}

	for index, company := range listCompany {
		var invReq datastruct.InvoiceRequest
		invReq.InvoiceDate = req.InvoiceDate
		invReq.CompanyID = company["company_id"]
		invReq.MonthUse = req.MonthUse
		invReq.InvoiceTypeID = company["default_invoice_type_id"]
		invReq.DiscountType = "percent"
		invReq.Discount = "0.0"
		invReq.Note = "-"
		invReq.LastUpdateUsername = req.LastUpdateUsername
		invReq.PPN = "11.0"
		invReq.CreatedBy = req.LastUpdateUsername
		invReq.PaymentMethod = req.PaymentMethod
		invReq.ExchangeRateDate = req.ExchangeRateDate
		invReq.DueDate = req.DueDate
		invReq.Sender = req.Sender
		invReq.BatchID = batchID
		if len(company["default_invoice_type_id"]) == 0 {
			//hanya generate yg ada default nya
			continue
		}
		var invoiceType dtInvoiceType.InvoiceTypeDataStruct
		for _, each := range listInvoiceType {
			if each["inv_type_id"] == company["default_invoice_type_id"] {
				invoiceType.InvoiceTypeID = each["inv_type_id"]
				invoiceType.LoadFromServer = each["load_from_server"]
				invoiceType.IsGroup = each["is_group"]
				invoiceType.ServerID = each["server_id"]
				// inv_type_id, inv_type_name, server_id, category, load_from_server,is_group,group_type,
			}
		}
		if invoiceType.IsGroup == "1" {
			continue
		}

		// if company["default_invoice_type_id"] == "1" || company["default_invoice_type_id"] == "2" || company["default_invoice_type_id"] == "101" {
		// if invoiceType.ServerID != "*" {
		var serverDataReq dtServerData.ServerDataRequest
		var listExternalRootParentAccount []string
		var listServerID []string
		listServerID = strings.Split(invoiceType.ServerID, "|")
		for _, each := range listServerAccount {
			if invoiceType.ServerID != "*" {
				if each["account_id"] == company["company_id"] && lib.IsContain(each["server_id"], listServerID) {
					listExternalRootParentAccount = append(listExternalRootParentAccount, each["external_account_id"])
				}
			} else {
				//semua server
				if each["account_id"] == company["company_id"] {
					listExternalRootParentAccount = append(listExternalRootParentAccount, each["external_account_id"])
				}
			}
		}

		invReq.ListServerID = listServerID

		if len(listExternalRootParentAccount) == 0 {
			logrus.Warn("External Account ID is empty")
			continue
		}
		serverDataReq.ListExternalRootParentAccount = listExternalRootParentAccount
		serverDataReq.MonthUse = req.MonthUse
		// if company["default_invoice_type_id"] != "101" {
		// 	serverDataReq.ServerID = "11"
		// }
		serverDataReq.InvoiceTypeID = company["default_invoice_type_id"]
		serverDataReq.ListServerID = listServerID
		serverDataReq.Param.PerPage = 99999
		serverDataReq.UseBillingPrice = true
		serverDataReq.CurrencyCode = req.CurrencyCode
		serverDataReq.Category = "USAGE"

		listServerData, err := serverDataModel.GetServerDataFromRequest(conn, serverDataReq)
		if err != nil {
			logrus.Error("Error : ", err)

		}
		var serverDatas []dtServerData.ServerDataDataStruct
		// var total int = 0
		for _, each := range listServerData {
			if len(each["external_price"].(string)) == 0 {
				each["external_price"] = "0"
			}
			svrData := serverDataProcessor.CreateSingleServerDataStruct(each)
			serverDatas = append(serverDatas, svrData)

		}
		serverDatas = UniqueServerData(serverDatas, company["default_invoice_type_id"])
		if len(serverDatas) == 0 {
			logrus.Warn("Server data is empty for " + company["name"] + ". Generation skipped.")
			continue
		}
		// logrus.Info("Len Server Data : ", len(listServerData))

		invReq.ListExternalRootParentAccount = listExternalRootParentAccount

		var listInvoiceDetail []datastruct.InvoiceDetailStruct
		for _, each := range serverDatas {
			var invoiceDetail datastruct.InvoiceDetailStruct
			invoiceDetail.AccountID = each.AccountID
			invoiceDetail.ItemID = each.ItemID
			invoiceDetail.Qty = each.ExternalSMSCount
			invoiceDetail.Uom = each.Item.UOM
			invoiceDetail.ItemPrice = each.ExternalPrice
			invoiceDetail.Note = "-"
			invoiceDetail.BalanceType = each.ExternalBalanceType
			invoiceDetail.ServerID = each.ServerID
			invoiceDetail.LastUpdateUsername = req.LastUpdateUsername

			listInvoiceDetail = append(listInvoiceDetail, invoiceDetail)
		}
		invReq.ListInvoiceDetail = listInvoiceDetail
		// if company["default_invoice_type_id"] != "101" {
		// 	invReq.ServerID = "11"
		// }

		err = models.InsertInvoice(conn, invReq)

		if err != nil {
			logrus.Error("Error : ", err)
		}

		// }
		// else if company["default_invoice_type_id"] == "10" || company["default_invoice_type_id"] == "11" || company["default_invoice_type_id"] == "101" {
		// 	var serverDataReq dtServerData.ServerDataRequest
		// 	var listExternalRootParentAccount []string
		// 	for _, each := range listServerAccount {
		// 		if each["account_id"] == company["company_id"] && each["server_id"] == "20" {
		// 			listExternalRootParentAccount = append(listExternalRootParentAccount, each["external_account_id"])
		// 		}
		// 	}

		// 	if len(listExternalRootParentAccount) == 0 {
		// 		logrus.Warn("External Account ID is empty")
		// 		continue
		// 	}
		// 	serverDataReq.ListExternalRootParentAccount = listExternalRootParentAccount
		// 	serverDataReq.MonthUse = req.MonthUse
		// 	if company["default_invoice_type_id"] != "101" {
		// 		serverDataReq.ServerID = "20"
		// 	}
		// 	serverDataReq.Param.PerPage = 99999
		// 	serverDataReq.UseBillingPrice = true
		// 	serverDataReq.CurrencyCode = req.CurrencyCode

		// 	listServerData, err := serverDataModel.GetServerDataFromRequest(conn, serverDataReq)
		// 	if err != nil {
		// 		logrus.Error("Error : ", err)

		// 	}
		// 	var serverDatas []dtServerData.ServerDataDataStruct
		// 	// var total int = 0
		// 	for _, each := range listServerData {
		// 		if len(each["external_price"].(string)) == 0 {
		// 			each["external_price"] = "0"
		// 		}
		// 		svrData := serverDataProcessor.CreateSingleServerDataStruct(each)
		// 		serverDatas = append(serverDatas, svrData)

		// 	}
		// 	serverDatas = UniqueServerData(serverDatas, company["default_invoice_type_id"])

		// 	invReq.ListExternalRootParentAccount = listExternalRootParentAccount

		// 	var listInvoiceDetail []datastruct.InvoiceDetailStruct
		// 	for _, each := range serverDatas {
		// 		var invoiceDetail datastruct.InvoiceDetailStruct
		// 		invoiceDetail.AccountID = each.AccountID
		// 		invoiceDetail.ItemID = each.ItemID
		// 		invoiceDetail.Qty = each.ExternalSMSCount
		// 		invoiceDetail.Uom = each.Item.UOM
		// 		invoiceDetail.ItemPrice = each.ExternalPrice
		// 		invoiceDetail.Note = "-"
		// 		invoiceDetail.BalanceType = each.ExternalBalanceType
		// 		invoiceDetail.ServerID = each.ServerID
		// 		invoiceDetail.LastUpdateUsername = req.LastUpdateUsername

		// 		listInvoiceDetail = append(listInvoiceDetail, invoiceDetail)
		// 	}

		// 	invReq.ListInvoiceDetail = listInvoiceDetail
		// 	if company["default_invoice_type_id"] != "101" {
		// 		invReq.ServerID = "20"
		// 	}

		// 	err = models.InsertInvoice(conn, invReq)
		// 	if err != nil {
		// 		logrus.Error("Error : ", err)

		// 	}
		// } else if company["default_invoice_type_id"] == "5" {
		// 	var listInvoiceGroupByCompanyID []dtInvoiceGroup.InvoiceGroupDataStruct
		// 	for _, each := range listGroup {
		// 		var invGroup dtInvoiceGroup.InvoiceGroupDataStruct
		// 		invGroup = invoiceGroupProcessor.CreateSingleInvoiceGroupStruct(each)
		// 		if company["company_id"] == invGroup.CompanyID {
		// 			listInvoiceGroupByCompanyID = append(listInvoiceGroupByCompanyID, invGroup)
		// 		}
		// 	}
		// 	for _, each := range listInvoiceGroupByCompanyID {
		// 		var listAccountID []string
		// 		// identity = account_id
		// 		for _, detail := range each.InvoiceGroupDetail {
		// 			listAccountID = append(listAccountID, detail.Identity)
		// 		}

		// 		var svrAccountReq dtServerAccount.ServerAccountRequest
		// 		svrAccountReq.ListAccountID = listAccountID
		// 		svrAccountReq.Param.PerPage = 9999

		// 		mapServerAccounts, errGetServerAccounts := serverAccountModel.GetServerAccountFromRequest(conn, svrAccountReq)
		// 		if errGetServerAccounts != nil {
		// 			logrus.Error("Error in SQL : ", errGetServerAccounts)

		// 		}

		// 		var listExternalRootParentAccount []string
		// 		for _, each := range mapServerAccounts {
		// 			// svrAccount := serverAccountProcessor.CreateSingleServerAccountStruct(each)
		// 			listExternalRootParentAccount = append(listExternalRootParentAccount, each["external_account_id"])
		// 		}

		// 		var serverDataReq dtServerData.ServerDataRequest
		// 		serverDataReq.ListExternalRootParentAccount = listExternalRootParentAccount
		// 		serverDataReq.MonthUse = req.MonthUse
		// 		serverDataReq.Param.PerPage = 99999
		// 		serverDataReq.UseBillingPrice = true
		// 		serverDataReq.CurrencyCode = req.CurrencyCode

		// 		listServerData, err := serverDataModel.GetServerDataFromRequest(conn, serverDataReq)
		// 		if err != nil {
		// 			logrus.Error("Error : ", err)

		// 		}
		// 		var serverDatas []dtServerData.ServerDataDataStruct

		// 		// var total int = 0
		// 		for _, each := range listServerData {
		// 			if len(each["external_price"].(string)) == 0 {
		// 				each["external_price"] = "0"
		// 			}
		// 			svrData := serverDataProcessor.CreateSingleServerDataStruct(each)
		// 			serverDatas = append(serverDatas, svrData)

		// 		}
		// 		serverDatas = UniqueServerData(serverDatas, company["default_invoice_type_id"])

		// 		invReq.ListExternalRootParentAccount = listExternalRootParentAccount

		// 		var listInvoiceDetail []datastruct.InvoiceDetailStruct
		// 		for _, each := range serverDatas {
		// 			var invoiceDetail datastruct.InvoiceDetailStruct
		// 			invoiceDetail.AccountID = each.AccountID
		// 			invoiceDetail.ItemID = each.ItemID
		// 			invoiceDetail.Qty = each.ExternalSMSCount
		// 			invoiceDetail.Uom = each.Item.UOM
		// 			invoiceDetail.ItemPrice = each.ExternalPrice
		// 			invoiceDetail.Note = "-"
		// 			invoiceDetail.BalanceType = each.ExternalBalanceType
		// 			invoiceDetail.ServerID = each.ServerID
		// 			invoiceDetail.LastUpdateUsername = req.LastUpdateUsername

		// 			listInvoiceDetail = append(listInvoiceDetail, invoiceDetail)
		// 		}

		// 		invReq.ListInvoiceDetail = listInvoiceDetail

		// 		err = models.InsertInvoice(conn, invReq)
		// 		if err != nil {
		// 			logrus.Error("Error : ", err)

		// 		}

		// 	}
		// 	var getAccountFromServerDataReq dtServerData.ServerDataRequest
		// 	var accountReq dtAccount.AccountRequest
		// 	accountReq.CompanyID = company["company_id"]
		// 	accountReq.Param.PerPage = 9999
		// 	listAccount, _ := accountModel.GetAccountFromRequest(conn, accountReq)
		// 	var listAccountId []string
		// 	for _, each := range listAccount {
		// 		listAccountId = append(listAccountId, each["account_id"])
		// 	}

		// 	getAccountFromServerDataReq.MonthUse = req.MonthUse
		// 	getAccountFromServerDataReq.ListAccountID = listAccountId
		// 	getAccountFromServerDataReq.Param.PerPage = 9999

		// 	listAccountFromServerData, err := serverDataModel.GetAccountFromServerData(conn, getAccountFromServerDataReq)
		// 	if err != nil {
		// 		logrus.Error("Error : ", err)

		// 	}

		// 	if len(listAccountFromServerData) > 0 {
		// 		var accountIds []string
		// 		for _, each := range listAccountFromServerData {
		// 			accountIds = append(accountIds, each["account_id"])
		// 		}

		// 		var svrAccountReq dtServerAccount.ServerAccountRequest
		// 		svrAccountReq.ListAccountID = accountIds
		// 		svrAccountReq.Param.PerPage = 9999

		// 		mapServerAccounts, errGetServerAccounts := serverAccountModel.GetServerAccountFromRequest(conn, svrAccountReq)
		// 		if errGetServerAccounts != nil {
		// 			logrus.Error("Error : ", errGetServerAccounts)

		// 		}

		// 		var listExternalRootParentAccount []string
		// 		for _, each := range mapServerAccounts {
		// 			// svrAccount := serverAccountProcessor.CreateSingleServerAccountStruct(each)
		// 			listExternalRootParentAccount = append(listExternalRootParentAccount, each["external_account_id"])
		// 		}

		// 		if len(listExternalRootParentAccount) == 0 {
		// 			logrus.Warn("External Account ID is empty")
		// 			continue
		// 		}

		// 		var serverDataReq dtServerData.ServerDataRequest
		// 		serverDataReq.ListExternalRootParentAccount = listExternalRootParentAccount
		// 		serverDataReq.MonthUse = req.MonthUse
		// 		serverDataReq.Param.PerPage = 99999
		// 		serverDataReq.UseBillingPrice = true
		// 		serverDataReq.CurrencyCode = req.CurrencyCode

		// 		listServerData, err := serverDataModel.GetServerDataFromRequest(conn, serverDataReq)
		// 		if err != nil {
		// 			logrus.Error("Error : ", err)

		// 		}

		// 		var serverDatas []dtServerData.ServerDataDataStruct

		// 		// var total int = 0
		// 		for _, each := range listServerData {
		// 			if len(each["external_price"].(string)) == 0 {
		// 				each["external_price"] = "0"
		// 			}
		// 			svrData := serverDataProcessor.CreateSingleServerDataStruct(each)
		// 			serverDatas = append(serverDatas, svrData)

		// 		}
		// 		serverDatas = UniqueServerData(serverDatas, company["default_invoice_type_id"])

		// 		invReq.ListExternalRootParentAccount = listExternalRootParentAccount

		// 		var listInvoiceDetail []datastruct.InvoiceDetailStruct
		// 		for _, each := range serverDatas {
		// 			var invoiceDetail datastruct.InvoiceDetailStruct
		// 			invoiceDetail.AccountID = each.AccountID
		// 			invoiceDetail.ItemID = each.ItemID
		// 			invoiceDetail.Qty = each.ExternalSMSCount
		// 			invoiceDetail.Uom = each.Item.UOM
		// 			invoiceDetail.ItemPrice = each.ExternalPrice
		// 			invoiceDetail.Note = "-"
		// 			invoiceDetail.BalanceType = each.ExternalBalanceType
		// 			invoiceDetail.ServerID = each.ServerID
		// 			invoiceDetail.LastUpdateUsername = req.LastUpdateUsername

		// 			listInvoiceDetail = append(listInvoiceDetail, invoiceDetail)
		// 		}

		// 		invReq.ListInvoiceDetail = listInvoiceDetail

		// 		err = models.InsertInvoice(conn, invReq)
		// 		if err != nil {
		// 			logrus.Error("Error : ", err)

		// 		}

		// 	}

		// } else if company["default_invoice_type_id"] == "6" {
		// 	var listInvoiceGroupByCompanyID []dtInvoiceGroup.InvoiceGroupDataStruct
		// 	for _, each := range listGroup {
		// 		var invGroup dtInvoiceGroup.InvoiceGroupDataStruct
		// 		invGroup = invoiceGroupProcessor.CreateSingleInvoiceGroupStruct(each)
		// 		if company["company_id"] == invGroup.CompanyID {
		// 			listInvoiceGroupByCompanyID = append(listInvoiceGroupByCompanyID, invGroup)
		// 		}
		// 	}

		// 	for _, each := range listInvoiceGroupByCompanyID {

		// 		var listAccountID []string
		// 		// identity = account_id
		// 		// for _, detail := range each.InvoiceGroupDetail {
		// 		// 	listAccountID = append(listAccountID, detail.Identity)
		// 		// }

		// 		listAccountID = append(listAccountID, company["company_id"])

		// 		var svrAccountReq dtServerAccount.ServerAccountRequest
		// 		svrAccountReq.ListAccountID = listAccountID
		// 		svrAccountReq.Param.PerPage = 9999

		// 		mapServerAccounts, errGetServerAccounts := serverAccountModel.GetServerAccountFromRequest(conn, svrAccountReq)
		// 		if errGetServerAccounts != nil {
		// 			logrus.Error("Error in SQL : ", errGetServerAccounts)

		// 		}

		// 		var listExternalRootParentAccount []string
		// 		for _, each := range mapServerAccounts {
		// 			// svrAccount := serverAccountProcessor.CreateSingleServerAccountStruct(each)
		// 			listExternalRootParentAccount = append(listExternalRootParentAccount, each["external_account_id"])
		// 		}

		// 		var listUserID []string
		// 		// identity = user_id
		// 		for _, detail := range each.InvoiceGroupDetail {
		// 			if detail.Type == "usr" {
		// 				listUserID = append(listUserID, detail.Identity)
		// 			}
		// 		}

		// 		var serverDataReq dtServerData.ServerDataRequest
		// 		serverDataReq.ListUserID = listUserID
		// 		serverDataReq.ListExternalRootParentAccount = listExternalRootParentAccount
		// 		serverDataReq.MonthUse = req.MonthUse
		// 		serverDataReq.Param.PerPage = 99999
		// 		serverDataReq.UseBillingPrice = true
		// 		serverDataReq.CurrencyCode = req.CurrencyCode

		// 		listServerData, err := serverDataModel.GetServerDataFromRequest(conn, serverDataReq)
		// 		if err != nil {
		// 			logrus.Error("Error : ", err)

		// 		}

		// 		var serverDatas []dtServerData.ServerDataDataStruct

		// 		// var total int = 0
		// 		for _, each := range listServerData {
		// 			if len(each["external_price"].(string)) == 0 {
		// 				each["external_price"] = "0"
		// 			}
		// 			svrData := serverDataProcessor.CreateSingleServerDataStruct(each)
		// 			serverDatas = append(serverDatas, svrData)

		// 		}
		// 		serverDatas = UniqueServerData(serverDatas, company["default_invoice_type_id"])

		// 		invReq.ListExternalRootParentAccount = listExternalRootParentAccount

		// 		var listInvoiceDetail []datastruct.InvoiceDetailStruct
		// 		for _, each := range serverDatas {
		// 			var invoiceDetail datastruct.InvoiceDetailStruct
		// 			invoiceDetail.AccountID = each.AccountID
		// 			invoiceDetail.ItemID = each.ItemID
		// 			invoiceDetail.Qty = each.ExternalSMSCount
		// 			invoiceDetail.Uom = each.Item.UOM
		// 			invoiceDetail.ItemPrice = each.ExternalPrice
		// 			invoiceDetail.Note = "-"
		// 			invoiceDetail.BalanceType = each.ExternalBalanceType
		// 			invoiceDetail.ServerID = each.ServerID
		// 			invoiceDetail.ExternalUserID = each.ExternalUserID
		// 			invoiceDetail.ExternalUsername = each.ExternalUsername
		// 			invoiceDetail.LastUpdateUsername = req.LastUpdateUsername

		// 			listInvoiceDetail = append(listInvoiceDetail, invoiceDetail)
		// 		}
		// 		invReq.ListInvoiceDetail = listInvoiceDetail

		// 		err = models.InsertInvoice(conn, invReq)
		// 		if err != nil {
		// 			logrus.Error("Error : ", err)

		// 		}

		// 	}

		// }

		progress = progress + percentPerCompany
		if index == len(listCompany)-1 {
			progress = 100
		}

		models.UpdateProgressGenerateInvoiceHistory(conn, progress, batchID)
	}
}

func UniqueServerData(list []dtServerData.ServerDataDataStruct, invoiceTypeID string) []dtServerData.ServerDataDataStruct {
	var unique []dtServerData.ServerDataDataStruct
sampleLoop:
	for _, v := range list {
		var smsCount int
		var transCount int
		externalSmsCount, _ := strconv.Atoi(v.ExternalSMSCount)
		smsCount = externalSmsCount
		externalTranscount, _ := strconv.Atoi(v.ExternalTransCount)
		transCount = externalTranscount
		for i, u := range unique {
			if invoiceTypeID == "6" {
				if v.ExternalOperatorCode == u.ExternalOperatorCode && v.AccountID == u.AccountID && v.NewRoute == u.NewRoute && v.ExternalPrice == u.ExternalPrice && v.ServerID == u.ServerID && v.ExternalUserID == u.ExternalUserID {
					addSmsCount, _ := strconv.Atoi(u.ExternalSMSCount)
					addTransCount, _ := strconv.Atoi(u.ExternalTransCount)

					smsCount += addSmsCount
					transCount += addTransCount

					v.ExternalSMSCount = strconv.Itoa(smsCount)
					v.ExternalTransCount = strconv.Itoa(transCount)

					unique[i] = v
					continue sampleLoop
				}
			} else {
				if v.ExternalOperatorCode == u.ExternalOperatorCode && v.AccountID == u.AccountID && v.NewRoute == u.NewRoute && v.ExternalPrice == u.ExternalPrice && v.ServerID == u.ServerID {
					addSmsCount, _ := strconv.Atoi(u.ExternalSMSCount)
					addTransCount, _ := strconv.Atoi(u.ExternalTransCount)

					smsCount += addSmsCount
					transCount += addTransCount

					v.ExternalSMSCount = strconv.Itoa(smsCount)
					v.ExternalTransCount = strconv.Itoa(transCount)

					unique[i] = v
					continue sampleLoop
				}
			}
		}
		unique = append(unique, v)

	}
	sort.Slice(unique[:], func(i, j int) bool {
		return unique[i].Item.ItemName < unique[j].Item.ItemName
	})
	return unique
}

func DeleteInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	err := models.DeleteInvoice(conn, req)
	return err
}

func CancelInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	err := models.CancelInvoice(conn, req)
	return err
}

func PrintInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	err = models.PrintInvoice(conn, req)
	if err != nil {
		return err
	}

	return err
}

func UpdateReceivedDate(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	var err error

	if len(req.ReceiptLetterAttachment) > 0 {
		b64data := req.ReceiptLetterAttachment[strings.IndexByte(req.ReceiptLetterAttachment, ',')+1:]
		dec, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			// panic(err)
			return err
		}

		if _, err := os.Stat(config.Param.AttachmentFolder); os.IsNotExist(err) {
			errMkdir := os.MkdirAll(config.Param.AttachmentFolder, os.ModePerm)
			if errMkdir != nil {
				logrus.Error("ERROR-", err.Error())
				return errMkdir
			}
			// TODO: handle error
		}

		now := time.Now().Unix()
		fileName := "attachment_receipt_letter_" + req.InvoiceNo + "_" + strconv.FormatInt(now, 10) + ".jpg"
		path := config.Param.AttachmentFolder + "/" + fileName
		err = os.WriteFile(path, dec, 0600)
		// f, err := os.Create(config.Param.AttachmentFolder + "/attachment_invoice_" + req.InvoiceNo + ".jpg")
		if err != nil {
			return err
		}
		req.ReceiptLetterAttachment = fileName

	}

	err = models.UpdateReceivedDate(conn, req)
	if err != nil {
		return err
	}

	return err
}

func SummaryInvoice(conn *connections.Connections, req datastruct.InvoiceRequest) ([]datastruct.InvoiceDataStruct, error) {
	var output []datastruct.InvoiceDataStruct
	var err error

	// grab mapping data from model
	invoiceList, err := models.GetListInvoice(conn, req)
	if err != nil {
		return output, err
	}

	var listInvoiceID []string
	for _, invoice := range invoiceList {
		// single := CreateSingleInvoiceStruct(invoice)
		// output = append(output, single)
		listInvoiceID = append(listInvoiceID, invoice["invoice_id"].(string))
	}

	invoiceDetailList, err := models.GetListInvoiceDetail(conn, listInvoiceID)

	for _, each := range invoiceList {
		out := CreateInvoiceStruct(each, invoiceDetailList)
		output = append(output, out)
	}

	return output, err
}

func InquiryPayment(conn *connections.Connections, req datastruct.InquiryPaymentRequest) ([]datastruct.InquiryPaymentDataStruct, error) {
	var output []datastruct.InquiryPaymentDataStruct
	var tempListData []datastruct.InquiryPaymentDataStruct
	var err error

	// grab mapping data from model
	list, err := models.GetInquiryPayment(conn, req)
	if err != nil {
		return output, err
	}

	// var listInvoiceID []string
	for _, each := range list {
		// listInvoiceID = append(listInvoiceID, each["invoice_id"])

		single := CreateInquiryPaymentStruct(each)

		tempListData = append(tempListData, single)
	}

	var invoiceReq datastruct.InvoiceRequest
	// invoiceReq.ListInvoiceID = listInvoiceID
	invoiceReq.MonthUse = req.MonthUse
	invoiceReq.Param.PerPage = 99999
	listInvoice, err := models.GetListInvoice(conn, invoiceReq)
	if err != nil {
		return output, err
	}

	var listInvoiceID []string
	for _, each := range listInvoice {
		listInvoiceID = append(listInvoiceID, each["invoice_id"].(string))

	}
	listInvoiceDetail, err := models.GetListInvoiceDetail(conn, listInvoiceID)

	var listInvoiceStruct []datastruct.InvoiceDataStruct
	for _, each := range listInvoice {
		inv := CreateInvoiceStruct(each, listInvoiceDetail)
		listInvoiceStruct = append(listInvoiceStruct, inv)
	}

	if err != nil {
		return output, err
	}

	var proformaInvoiceReq datastruct.InvoiceRequest
	// invoiceReq.ListInvoiceID = listInvoiceID
	proformaInvoiceReq.MonthUse = req.MonthUse
	proformaInvoiceReq.Param.PerPage = 99999
	listProformaInvoice, err := proformaInvoiceModel.GetListInvoice(conn, proformaInvoiceReq)
	if err != nil {
		return output, err
	}

	var listProformaInvoiceID []string
	for _, each := range listProformaInvoice {
		listProformaInvoiceID = append(listProformaInvoiceID, each["invoice_id"].(string))
	}
	listProformaInvoiceDetail, err := proformaInvoiceModel.GetListInvoiceDetail(conn, listProformaInvoiceID)

	var listProformaInvoiceStruct []datastruct.InvoiceDataStruct
	for _, each := range listProformaInvoice {
		inv := CreateInvoiceStruct(each, listProformaInvoiceDetail)
		listProformaInvoiceStruct = append(listProformaInvoiceStruct, inv)
	}

	// listInvoiceDetail, err := models.GetListInvoiceDetail(conn, listInvoiceID)

	// if err != nil {
	// 	return output, err
	// }

	for _, each := range tempListData {
		var out datastruct.InquiryPaymentDataStruct
		out = each
		out.OutStanding = each.Amount

		for _, inv := range listInvoiceStruct {
			if inv.InvoiceStatus == "D" {
				continue
			}
			if inv.CompanyID == each.CompanyID {
				skip := true
				for _, detail := range inv.ListInvoiceDetail {
					if detail.AccountID == each.AccountID && inv.CompanyID == each.CompanyID {
						skip = false
						break
					}
				}

				if skip == false {
					grandTotalStr := inv.GrandTotal
					var grandTotal float64 = 0

					if len(grandTotalStr) > 0 {
						grandTotal, err = strconv.ParseFloat(grandTotalStr, 64)
						if err != nil {
							return output, err
						}
					}

					var ppn float64 = 0
					ppnStr := inv.PPNAmount
					if len(ppnStr) > 0 {
						ppn, err = strconv.ParseFloat(ppnStr, 64)
						if err != nil {
							return output, err
						}
					}

					subTotal := grandTotal - ppn

					subTotalStr := fmt.Sprintf("%2f", subTotal)
					// subTotalStr := strconv.Itoa(int(subTotal))

					out.Invoicing = subTotalStr

					var amount float64
					amountStr := out.Amount
					if len(amountStr) > 0 {
						amount, err = strconv.ParseFloat(amountStr, 64)
						if err != nil {
							return output, err
						}
					}

					outstanding := amount - subTotal

					var adjustmentAmount float64 = 0
					var adjustmenConfirmed bool = false
					for _, each := range inv.ListInvoiceDetail {
						if inv.InvoiceID == each.InvoiceID {
							if len(each.Adjustment) > 0 {
								// logrus.Info("Lihat Invoice Detail ID: ", each.InvoiceDetailID)
								var price float64 = 0
								price, err := strconv.ParseFloat(each.ItemPrice, 64)
								if err != nil {
									return output, err
								}

								var adjustmentSmscount float64 = 0
								adjustmentSmscount, err = strconv.ParseFloat(each.Adjustment, 64)
								// logrus.Info("Lihat adjustmentSmscount", adjustmentSmscount)
								if err != nil {
									return output, err
								}

								// logrus.Info("Lihat Perhitungan ", price*adjustmentSmscount)

								adjustmentAmount += price * adjustmentSmscount
								if len(each.AdjustmentConfirmationDate) > 0 {
									adjustmenConfirmed = true
								}

							}
						}
					}
					// logrus.Info("Lihat outstanding: ", outstanding)
					// logrus.Info("Lihat adjustmentAmount: ", adjustmentAmount)

					if adjustmenConfirmed {
						outstanding = outstanding - adjustmentAmount
					}
					out.OutStanding = fmt.Sprintf("%2f", outstanding)

					// output = append(output, out)
				}

				break
			}

		}

		for _, inv := range listProformaInvoiceStruct {
			if inv.InvoiceStatus == "D" {
				continue
			}
			if inv.CompanyID == each.CompanyID {
				skip := true
				for _, detail := range inv.ListInvoiceDetail {
					if detail.AccountID == each.AccountID && inv.CompanyID == each.CompanyID {
						skip = false
						break
					}
				}

				if skip == false {
					grandTotalStr := inv.GrandTotal
					var grandTotal float64 = 0

					if len(grandTotalStr) > 0 {
						grandTotal, err = strconv.ParseFloat(grandTotalStr, 64)
						if err != nil {
							return output, err
						}
					}

					var ppn float64 = 0
					ppnStr := inv.PPNAmount
					if len(ppnStr) > 0 {
						ppn, err = strconv.ParseFloat(ppnStr, 64)
						if err != nil {
							return output, err
						}
					}

					subTotal := grandTotal - ppn

					subTotalStr := fmt.Sprintf("%2f", subTotal)
					// subTotalStr := strconv.Itoa(int(subTotal))

					out.ProformaInvoiceAmount = subTotalStr

					// var amount float64
					// amountStr := out.Amount
					// if len(amountStr) > 0 {
					// 	amount, err = strconv.ParseFloat(amountStr, 64)
					// 	if err != nil {
					// 		return output, err
					// 	}
					// }

					// outstanding := amount - subTotal

					// var adjustmentAmount float64 = 0
					// var adjustmenConfirmed bool = false
					// for _, each := range inv.ListInvoiceDetail {
					// 	if inv.InvoiceID == each.InvoiceID {
					// 		if len(each.Adjustment) > 0 {
					// 			// logrus.Info("Lihat Invoice Detail ID: ", each.InvoiceDetailID)
					// 			var price float64 = 0
					// 			price, err := strconv.ParseFloat(each.ItemPrice, 64)
					// 			if err != nil {
					// 				return output, err
					// 			}

					// 			var adjustmentSmscount float64 = 0
					// 			adjustmentSmscount, err = strconv.ParseFloat(each.Adjustment, 64)
					// 			// logrus.Info("Lihat adjustmentSmscount", adjustmentSmscount)
					// 			if err != nil {
					// 				return output, err
					// 			}

					// 			// logrus.Info("Lihat Perhitungan ", price*adjustmentSmscount)

					// 			adjustmentAmount += price * adjustmentSmscount
					// 			if len(each.AdjustmentConfirmationDate) > 0 {
					// 				adjustmenConfirmed = true
					// 			}

					// 		}
					// 	}
					// }
					// // logrus.Info("Lihat outstanding: ", outstanding)
					// // logrus.Info("Lihat adjustmentAmount: ", adjustmentAmount)

					// if adjustmenConfirmed {
					// 	outstanding = outstanding - adjustmentAmount
					// }
					// out.OutStanding = fmt.Sprintf("%2f", outstanding)

					// output = append(output, out)
				}

				break
			}

		}

		output = append(output, out)
		// if len(each.InvoiceID) > 0 {
		// 	for _, dt := range listInvoice {
		// 		if each.InvoiceID == dt["invoice_id"].(string) {
		// 			out = each
		// 			// var subTotal int64
		// 			grandTotalStr := dt["grand_total"].(string)
		// 			var grandTotal float64 = 0
		// 			if len(grandTotalStr) > 0 {
		// 				grandTotal, err = strconv.ParseFloat(grandTotalStr, 64)
		// 				if err != nil {
		// 					return output, err
		// 				}
		// 			}

		// 			var ppn float64 = 0
		// 			ppnStr := dt["ppn_amount"].(string)
		// 			if len(ppnStr) > 0 {
		// 				ppn, err = strconv.ParseFloat(ppnStr, 64)
		// 				if err != nil {
		// 					return output, err
		// 				}
		// 			}

		// 			subTotal := grandTotal - ppn

		// 			subTotalStr := fmt.Sprintf("%2f", subTotal)
		// 			// subTotalStr := strconv.Itoa(int(subTotal))

		// 			out.Invoicing = subTotalStr

		// 			var amount float64
		// 			amountStr := out.Amount
		// 			if len(amountStr) > 0 {
		// 				amount, err = strconv.ParseFloat(amountStr, 64)
		// 				if err != nil {
		// 					return output, err
		// 				}
		// 			}

		// 			outstanding := amount - subTotal

		// 			out.OutStanding = fmt.Sprintf("%2f", outstanding)

		// 			output = append(output, out)

		// 			// for _, detail := range listInvoiceDetail {
		// 			// 	if detail["invoice_id"] == each["invoice_id"].(string) {
		// 			// 		qty ,_:= strconv.Atoi(detail["qty"])
		// 			// 		price ,_:= strconv.Atoi(detail["price"])

		// 			// 	}
		// 			// }

		// 			break
		// 		}
		// 	}
		// } else {
		// 	out = each
		// 	out.OutStanding = out.Amount
		// 	output = append(output, out)
		// }

	}

	if len(req.CompanyID) > 0 {
		detailsResult, err := models.GetInquiryPaymentDetail(conn, req)
		if err != nil {
			return output, err
		}

		for index, _ := range output {
			p := &output[index]
			var details []datastruct.InquiryPaymentDetailStruct
			for _, each := range detailsResult {
				if p.CompanyID == each["company_id"] {
					var dt datastruct.InquiryPaymentDetailStruct
					dt.CompanyID = each["company_id"]
					dt.CompanyName = each["company_name"]
					dt.AccountID = each["account_id"]
					dt.AccountName = each["account_name"]
					dt.Amount = each["amount"]
					dt.Sender = each["external_sender"]
					dt.ItemName = each["item_name"]
					dt.ItemID = each["item_id"]
					dt.InvoiceNo = each["invoice_no"]
					dt.InvoiceID = each["invoice_id"]
					dt.UserID = each["external_user_id"]
					dt.Username = each["external_username"]
					details = append(details, dt)
					// &single.InquiryPaymentDetail =
				}
			}
			p.InquiryPaymentDetail = details
		}
	}

	// invoiceDetailList, err := models.GetListInvoiceDetail(conn, listInvoiceID)

	// for _, each := range list {
	// 	out := CreateInvoiceStruct(each, invoiceDetailList)
	// 	output = append(output, out)
	// }

	return output, err
}

func CreateInvoiceStruct(invoice map[string]interface{}, detail []map[string]string) datastruct.InvoiceDataStruct {

	var data datastruct.InvoiceDataStruct

	// var single datastruct.InvoiceDataStruct
	data.InvoiceID, _ = invoice["invoice_id"].(string)
	data.InvoiceNo, _ = invoice["invoice_no"].(string)
	data.InvoiceDate, _ = invoice["invoice_date"].(string)
	data.InvoiceStatus, _ = invoice["invoicestatus"].(string)
	data.CompanyID, _ = invoice["company_id"].(string)
	data.MonthUse, _ = invoice["month_use"].(string)
	data.InvoiceTypeID, _ = invoice["inv_type_id"].(string)
	data.PrintCounter, _ = invoice["printcounter"].(string)
	data.PPN, _ = invoice["ppn"].(string)
	data.Note, _ = invoice["note"].(string)
	data.CancelDesc, _ = invoice["canceldesc"].(string)
	data.LastPrintUsername, _ = invoice["last_print_username"].(string)
	data.LastPrintDate, _ = invoice["last_print_date"].(string)
	data.CreatedAt, _ = invoice["created_at"].(string)
	data.CreatedBy, _ = invoice["created_by"].(string)
	data.LastUpdateUsername, _ = invoice["last_update_username"].(string)
	data.Attachment, _ = invoice["attachment"].(string)
	data.LastUpdateDate, _ = invoice["last_update_date"].(string)
	data.DiscountType, _ = invoice["discount_type"].(string)
	data.Discount, _ = invoice["discount"].(string)
	data.Paid, _ = invoice["paid"].(string)
	data.PaymentMethod, _ = invoice["payment_method"].(string)
	data.ExchangeRateDate, _ = invoice["exchange_rate_date"].(string)
	data.DueDate, _ = invoice["due_date"].(string)
	data.GrandTotal, _ = invoice["grand_total"].(string)
	data.PPNAmount, _ = invoice["ppn_amount"].(string)
	data.Sender, _ = invoice["sender"].(string)
	data.TaxInvoice, _ = invoice["tax_invoice"].(string)
	data.AdjustmentNote, _ = invoice["adjustment_note"].(string)
	data.ReceivedDate, _ = invoice["received_date"].(string)
	data.ReceiptLetterAttachment, _ = invoice["receipt_letter_attachment"].(string)

	var invoiceType datastruct.InvoiceTypeDataStruct
	invoiceType.InvoiceTypeID = invoice["invoice_type"].(map[string]interface{})["inv_type_id"].(string)
	invoiceType.InvoiceTypeName = invoice["invoice_type"].(map[string]interface{})["inv_type_name"].(string)
	invoiceType.ServerID = invoice["invoice_type"].(map[string]interface{})["server_id"].(string)
	invoiceType.Category = invoice["invoice_type"].(map[string]interface{})["category"].(string)
	invoiceType.LoadFromServer = invoice["invoice_type"].(map[string]interface{})["load_from_server"].(string)
	invoiceType.IsGroup = invoice["invoice_type"].(map[string]interface{})["is_group"].(string)
	invoiceType.GroupType = invoice["invoice_type"].(map[string]interface{})["group_type"].(string)
	invoiceType.CurrencyCode = invoice["invoice_type"].(map[string]interface{})["currency_code"].(string)

	data.InvoiceType = invoiceType

	var company dtCompany.CompanyDataStruct
	company.CompanyID = invoice["company"].(map[string]interface{})["company_id"].(string)
	company.Name = invoice["company"].(map[string]interface{})["name"].(string)
	company.Address1 = invoice["company"].(map[string]interface{})["address1"].(string)
	company.Address2 = invoice["company"].(map[string]interface{})["address2"].(string)
	company.City = invoice["company"].(map[string]interface{})["city"].(string)
	company.ContactPerson = invoice["company"].(map[string]interface{})["contact_person"].(string)
	company.ContactPersonPhone = invoice["company"].(map[string]interface{})["contact_person_phone"].(string)

	data.Company = company

	var invoiceDetails []datastruct.InvoiceDetailStruct

	for _, each := range detail {
		if data.InvoiceID == each["invoice_id"] {
			var dtDetail datastruct.InvoiceDetailStruct
			dtDetail.InvoiceDetailID = each["invoice_detail_id"]
			dtDetail.InvoiceID = each["invoice_id"]
			dtDetail.ItemID = each["itemid"]
			dtDetail.ItemPrice = each["item_price"]
			dtDetail.AccountID = each["account_id"]
			dtDetail.Qty = each["qty"]
			dtDetail.Uom = each["uom"]
			dtDetail.Note = each["note"]
			dtDetail.BalanceType = each["balance_type"]
			dtDetail.ServerID = each["server_id"]
			dtDetail.ExternalUserID = each["external_user_id"]
			dtDetail.ExternalUsername = each["external_username"]
			dtDetail.ExternalSender = each["external_sender"]
			dtDetail.Adjustment = each["adjustment"]
			dtDetail.AdjustmentConfirmationUsername = each["adjustment_confirmation_username"]
			dtDetail.AdjustmentConfirmationDate = each["adjustment_confirmation_date"]

			var item datastruct.ItemDataStruct
			item.ItemID = each["item_id"]
			item.ItemName = each["item_name"]
			item.Operator = each["operator"]
			item.Route = each["route"]
			item.Category = each["category"]
			item.UOM = each["uom"]

			var server dtServer.ServerDataStruct
			server.ServerID = each["server_id"]
			server.ServerName = each["server_name"]

			var account dtAccount.AccountDataStruct
			account.AccountID = each["account_id"]
			account.Name = each["name"]
			account.AccountType = each["account_type"]
			account.CompanyID = each["company_id"]
			account.NonTaxable = each["non_taxable"]

			dtDetail.Account = account
			dtDetail.Item = item
			dtDetail.Server = server

			invoiceDetails = append(invoiceDetails, dtDetail)
		}
	}

	data.ListInvoiceDetail = invoiceDetails

	return data
}

func CreateInquiryPaymentStruct(inq map[string]string) datastruct.InquiryPaymentDataStruct {
	var data datastruct.InquiryPaymentDataStruct
	data.AccountID = inq["account_id"]
	data.AccountName = inq["account_name"]
	data.CompanyID = inq["company_id"]
	data.CompanyName = inq["company_name"]
	data.InvoiceID = inq["invoice_id"]

	var amount float64 = 0
	if len(inq["amount"]) > 0 {
		amount, _ = strconv.ParseFloat(inq["amount"], 64)
	}
	data.Amount = fmt.Sprintf("%2f", amount)
	// data.Invoicing = inq["invoicing"]
	// data.OutStanding = inq["outstanding"]

	return data
}

func AdjustmentConfirmation(conn *connections.Connections, req datastruct.InvoiceRequest) error {
	// var output datastruct.InvoiceDataStruct
	var err error

	err = models.AdjustmentConfirmation(conn, req)
	if err != nil {
		return err
	}

	// jika tidak ada error, return single instance of single stub
	// single, err := models.GetSingleCompany(req.CompanyID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleCompanyStruct(single)
	return err
}

// for _, each := range tempListData {
// 	var out datastruct.InquiryPaymentDataStruct
// 	if len(each.InvoiceID) > 0 {
// 		for _, dt := range listInvoice {
// 			if each.InvoiceID == dt["invoice_id"].(string) {
// 				out = each
// 				// var subTotal int64
// 				grandTotalStr := dt["grand_total"].(string)
// 				var grandTotal float64 = 0
// 				if len(grandTotalStr) > 0 {
// 					grandTotal, err = strconv.ParseFloat(grandTotalStr, 64)
// 					if err != nil {
// 						return output, err
// 					}
// 				}

// 				var ppn float64 = 0
// 				ppnStr := dt["ppn_amount"].(string)
// 				if len(ppnStr) > 0 {
// 					ppn, err = strconv.ParseFloat(ppnStr, 64)
// 					if err != nil {
// 						return output, err
// 					}
// 				}

// 				subTotal := grandTotal - ppn

// 				subTotalStr := fmt.Sprintf("%2f", subTotal)
// 				// subTotalStr := strconv.Itoa(int(subTotal))

// 				out.Invoicing = subTotalStr

// 				var amount float64
// 				amountStr := out.Amount
// 				if len(amountStr) > 0 {
// 					amount, err = strconv.ParseFloat(amountStr, 64)
// 					if err != nil {
// 						return output, err
// 					}
// 				}

// 				outstanding := amount - subTotal

// 				out.OutStanding = fmt.Sprintf("%2f", outstanding)

// 				output = append(output, out)

// 				// for _, detail := range listInvoiceDetail {
// 				// 	if detail["invoice_id"] == each["invoice_id"].(string) {
// 				// 		qty ,_:= strconv.Atoi(detail["qty"])
// 				// 		price ,_:= strconv.Atoi(detail["price"])

// 				// 	}
// 				// }

// 				break
// 			}
// 		}
// 	} else {
// 		out = each
// 		out.OutStanding = out.Amount
// 		output = append(output, out)
// 	}

// }

package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/invoice-type/datastruct"
	"backendbillingdashboard/modules/invoice-type/models"
)

func GetListInvoiceType(conn *connections.Connections, req datastruct.InvoiceTypeRequest) ([]datastruct.InvoiceTypeDataStruct, error) {
	var output []datastruct.InvoiceTypeDataStruct
	var err error

	// grab mapping data from model
	invoiceTypeList, err := models.GetInvoiceTypeFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, invoiceType := range invoiceTypeList {
		single := CreateSingleInvoiceTypeStruct(invoiceType)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleInvoiceTypeStruct(invoiceType map[string]string) datastruct.InvoiceTypeDataStruct {
	var single datastruct.InvoiceTypeDataStruct
	single.InvoiceTypeID, _ = invoiceType["inv_type_id"]
	single.InvoiceTypeName, _ = invoiceType["inv_type_name"]
	single.ServerID, _ = invoiceType["server_id"]
	single.Category, _ = invoiceType["category"]
	single.LoadFromServer, _ = invoiceType["load_from_server"]
	single.IsGroup, _ = invoiceType["is_group"]
	single.GroupType, _ = invoiceType["group_type"]
	single.CurrencyCode, _ = invoiceType["currency_code"]
	single.LastUpdateUsername, _ = invoiceType["last_update_username"]
	single.LastUpdateDate, _ = invoiceType["last_update_date"]

	return single
}

func InsertInvoiceType(conn *connections.Connections, req datastruct.InvoiceTypeRequest) (datastruct.InvoiceTypeDataStruct, error) {
	var output datastruct.InvoiceTypeDataStruct
	var err error

	err = models.InsertInvoiceType(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single invoice-type
	single, err := models.GetSingleInvoiceType(req.InvoiceTypeID, conn)
	if err != nil {
		return output, err
	}

	output = CreateSingleInvoiceTypeStruct(single)
	return output, err
}

func UpdateInvoiceType(conn *connections.Connections, req datastruct.InvoiceTypeRequest) (datastruct.InvoiceTypeDataStruct, error) {
	var output datastruct.InvoiceTypeDataStruct
	var err error

	err = models.UpdateInvoiceType(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single invoice-type
	single, err := models.GetSingleInvoiceType(req.InvoiceTypeID, conn)
	if err != nil {
		return output, err
	}

	output = CreateSingleInvoiceTypeStruct(single)
	return output, err
}

func DeleteInvoiceType(conn *connections.Connections, req datastruct.InvoiceTypeRequest) error {
	err := models.DeleteInvoiceType(conn, req)
	return err
}

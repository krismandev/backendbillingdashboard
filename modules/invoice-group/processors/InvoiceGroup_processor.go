package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/invoice-group/datastruct"
	"backendbillingdashboard/modules/invoice-group/models"
)

func GetListInvoiceGroup(conn *connections.Connections, req datastruct.InvoiceGroupRequest) ([]datastruct.InvoiceGroupDataStruct, error) {
	var output []datastruct.InvoiceGroupDataStruct
	var err error

	// grab mapping data from model
	stubList, err := models.GetInvoiceGroupFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, stub := range stubList {
		single := CreateSingleInvoiceGroupStruct(stub)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleInvoiceGroupStruct(data map[string]interface{}) datastruct.InvoiceGroupDataStruct {
	var single datastruct.InvoiceGroupDataStruct
	single.GroupID, _ = data["group_id"].(string)
	single.GroupName, _ = data["group_name"].(string)
	single.CompanyID, _ = data["company_id"].(string)
	single.InvoiceTypeID, _ = data["invoice_type_id"].(string)

	var details []datastruct.InvoiceGroupDetailDataStruct
	for _, eachDetail := range data["company_invoice_group_detail"].([]map[string]interface{}) {
		var detail datastruct.InvoiceGroupDetailDataStruct
		detail.GroupID = eachDetail["group_id"].(string)
		detail.Identity = eachDetail["identity"].(string)
		detail.Type = eachDetail["type"].(string)
		details = append(details, detail)
	}
	single.InvoiceGroupDetail = details

	return single
}

func InsertInvoiceGroup(conn *connections.Connections, req datastruct.InvoiceGroupRequest) (datastruct.InvoiceGroupDataStruct, error) {
	var output datastruct.InvoiceGroupDataStruct
	var err error

	err = models.InsertInvoiceGroup(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	// single, err := models.GetSingleInvoiceGroup(req.GroupID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleInvoiceGroupStruct(single)
	return output, err
}

func UpdateInvoiceGroup(conn *connections.Connections, req datastruct.InvoiceGroupRequest) (datastruct.InvoiceGroupDataStruct, error) {
	var output datastruct.InvoiceGroupDataStruct
	var err error

	err = models.UpdateInvoiceGroup(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single stub
	single, err := models.GetSingleInvoiceGroup(req.GroupID, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleInvoiceGroupStruct(single)
	return output, err
}

func DeleteInvoiceGroup(conn *connections.Connections, req datastruct.InvoiceGroupRequest) error {
	err := models.DeleteInvoiceGroup(conn, req)
	return err
}

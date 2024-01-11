package processors

import (
	"backendbillingdashboard/connections"
	dtCompany "backendbillingdashboard/modules/company/datastruct"
	companyModel "backendbillingdashboard/modules/company/models"
	"backendbillingdashboard/modules/reconciliation/datastruct"
	"backendbillingdashboard/modules/reconciliation/models"
)

func GetListRecon(conn *connections.Connections, req datastruct.ReconRequest) ([]datastruct.ReconDataStruct, error) {
	var output []datastruct.ReconDataStruct
	var tempOutput []datastruct.ReconDataStruct
	var err error

	var listCompanyID []string

	var companyRequest dtCompany.CompanyRequest
	companyRequest.CompanyID = req.CompanyID
	companyRequest.Name = req.CompanyName

	companyResult, err := companyModel.GetCompanyFromRequest(conn, companyRequest)
	if err != nil {
		return output, err
	}

	for _, company := range companyResult {
		// single := CreateSingleReconStruct(stub)
		// output = append(output, single)

		var out datastruct.ReconDataStruct
		out.CompanyID = company["company_id"]
		out.CompanyName = company["name"]
		out.MonthUse = req.MonthUse

		tempOutput = append(tempOutput, out)
		listCompanyID = append(listCompanyID, company["company_id"])
	}

	// grab mapping data from model
	req.ListCompanyID = listCompanyID
	serverData, err := models.GetReconFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, each := range tempOutput {
		var out datastruct.ReconDataStruct
		out = each
		var listServerData []datastruct.ServerDataStruct
		for _, dt := range serverData {
			if dt["company_id"] == out.CompanyID {
				var data datastruct.ServerDataStruct
				data.AccountID = dt["account_id"]
				data.AccountName = dt["account_name"]
				data.ExternalOperatorCode = dt["external_operatorcode"]
				data.ExternalRoute = dt["external_route"]
				data.ServerID = dt["server_id"]
				data.ServerName = dt["server_name"]
				data.SmsCount = dt["smscount"]

				listServerData = append(listServerData, data)

			}
		}
		out.ListServerData = listServerData

		output = append(output, out)
	}

	return output, err
}

func CreateSingleReconStruct(data map[string]string) datastruct.ServerDataStruct {
	var single datastruct.ServerDataStruct
	// single.CompanyID, _ = stub["stubid"]
	// single.ReconName, _ = stub["stubname"]

	single.AccountID = data["account_id"]
	single.AccountName = data["account_name"]
	single.ExternalOperatorCode = data["external_operatorcode"]
	single.ExternalRoute = data["external_route"]
	single.SmsCount = data["smscount"]

	return single
}

func InsertRecon(conn *connections.Connections, req datastruct.ReconRequest) (datastruct.ReconDataStruct, error) {
	var output datastruct.ReconDataStruct
	var err error

	// err = models.InsertRecon(conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// // jika tidak ada error, return single instance of single stub
	// single, err := models.GetSingleRecon(req.ReconID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleReconStruct(single)
	return output, err
}

func UpdateRecon(conn *connections.Connections, req datastruct.ReconRequest) (datastruct.ReconDataStruct, error) {
	var output datastruct.ReconDataStruct
	var err error

	// err = models.UpdateRecon(conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// // jika tidak ada error, return single instance of single stub
	// single, err := models.GetSingleRecon(req.ReconID, conn, req)
	// if err != nil {
	// 	return output, err
	// }

	// output = CreateSingleReconStruct(single)
	return output, err
}

func DeleteRecon(conn *connections.Connections, req datastruct.ReconRequest) error {
	err := models.DeleteRecon(conn, req)
	return err
}

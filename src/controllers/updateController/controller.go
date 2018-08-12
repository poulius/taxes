package updateController

import (
	"modules/taxes"
	"net/http"
)

func Handle(request *http.Request) interface{} {
	var (
		taxes       taxes.Taxes
		response    response
		errResponse errResponse
	)
	err := taxes.Get()
	if err != nil {
		errResponse.Error = "Error: cannot load taxes data from file"
		return errResponse
	}

	taxType := request.URL.Query().Get("type")
	startDateStr := request.URL.Query().Get("startDate")
	endDateStr := request.URL.Query().Get("endDate")
	taxRateValueStr := request.URL.Query().Get("taxRateValue")
	municipalityID := request.URL.Query().Get("municipalityID")
	taxRateID := request.URL.Query().Get("taxRateID")

	err = taxes.Update(taxType, startDateStr, endDateStr, taxRateValueStr, municipalityID, taxRateID)
	if err != nil {
		errResponse.Error = "Cannot add tax rates. Error: " + err.Error()
		return errResponse
	}

	response.Message = "Taxes data was successfully updated"

	return response
}

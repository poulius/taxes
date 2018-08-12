package addController

import (
	"modules/api"
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

	municipality := request.URL.Query().Get("municipality")
	taxType := request.URL.Query().Get("type")
	startDateStr := request.URL.Query().Get("startDate")
	endDateStr := request.URL.Query().Get("endDate")
	taxRateValueStr := request.URL.Query().Get("taxRateValue")

	err = taxes.Add(municipality, taxType, startDateStr, endDateStr, taxRateValueStr)
	if err != nil {
		errResponse.Error = "Cannot add tax rates. Error: " + err.Error()
		return errResponse
	}

	// post to 3rd party service
	err = api.Post(municipality, taxType, startDateStr, endDateStr, taxRateValueStr)
	if err != nil {
		errResponse.Error = "Cannot post tax rates to 3rd party service. Error: " + err.Error()
		return errResponse
	}

	response.Message = "Taxes data was successfully added"

	return response
}

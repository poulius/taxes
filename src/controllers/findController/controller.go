package findController

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

	municipality := request.URL.Query().Get("municipality")
	dateStr := request.URL.Query().Get("date")

	taxRate, err := taxes.Find(municipality, dateStr)
	if err != nil {
		errResponse.Error = "Error: cannot find tax rates for this municipality and date"
		return errResponse
	}

	response.TaxRateValue = taxRate.TaxRateValue

	return response
}

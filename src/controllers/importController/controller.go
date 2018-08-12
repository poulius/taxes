package importController

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

	err = taxes.Import()
	if err != nil {
		errResponse.Error = "Error: cannot add tax rates"
		return errResponse
	}

	response.Message = "Taxes data was successfully imported"

	return response
}

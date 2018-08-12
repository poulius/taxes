package httpWorker

import (
	"controllers/addController"
	"controllers/findController"
	"controllers/importController"
	"controllers/updateController"
	"encoding/json"
	"net/http"
	"strings"
)

func handle(response http.ResponseWriter, request *http.Request) {
	var (
		controllerResponse interface{}
		jsonResponse       []byte
	)

	urlSegments := strings.Split(request.URL.Path, "/")

	if len(urlSegments) < 3 {
		response.Header().Set("Content-Type", "application/json")
		jsonResponse = []byte(`{"error":"Error: wrong params"}`)
		response.Write(jsonResponse)
		return
	} else {

		switch urlSegments[2] {
		case "add":
			controllerResponse = addController.Handle(request)
		case "import":
			controllerResponse = importController.Handle(request)
		case "find":
			controllerResponse = findController.Handle(request)
		case "update":
			controllerResponse = updateController.Handle(request)
		default:
			response.Header().Set("Content-Type", "application/json")
			jsonResponse = []byte(`{"error":"Error: wrong params"}`)
			response.Write(jsonResponse)
			return
		}

		jsonResponse, err := json.Marshal(controllerResponse)
		if err != nil {
			response.Header().Set("Content-Type", "application/json")
			jsonResponse = []byte(`{"error":"Error: wrong params"}`)
			response.Write(jsonResponse)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

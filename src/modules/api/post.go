package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	apiUrl string = "https://api.myjson.com/bins"
)

func Post(municipality, taxType, startDateStr, endDateStr, taxRateStr string) error {
	taxRateRequest := taxRate{taxType, startDateStr, endDateStr, taxRateStr}
	apiRequest := apiRequest{municipality, taxRateRequest}

	jsonApiRequest, err := json.Marshal(apiRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonApiRequest))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// printing this to get the link to submited data
	fmt.Println("Data pushed to:", string(body))

	return nil
}

package findController

type response struct {
	TaxRateValue float64 `json:"taxRateValue"`
}

type errResponse struct {
	Error string `json:"error"`
}

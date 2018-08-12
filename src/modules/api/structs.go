package api

type apiRequest struct {
	Municipality string  `json:"municipality"`
	TaxRate      taxRate `json:"taxRate"`
}

type taxRate struct {
	Type      string `json:"type"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	TaxRate   string `json:"taxRate"`
}

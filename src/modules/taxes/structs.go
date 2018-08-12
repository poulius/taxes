package taxes

type Taxes struct {
	Taxes []Tax `json:"taxes"`
}

type TaxRate struct {
	TaxRateID    int     `json:"taxRateID"`
	Type         string  `json:"type"`
	StartDate    string  `json:"startDate"`
	EndDate      string  `json:"endDate"`
	TaxRateValue float64 `json:"taxRate"`
}

type Tax struct {
	MunicipalityID string    `json:"municipalityID"`
	Municipality   string    `json:"municipality"`
	TaxRates       []TaxRate `json:"taxRates"`
}

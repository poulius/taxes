package taxes

import (
	"time"
)

func (taxes *Taxes) Find(municipality, dateStr string) (*TaxRate, error) {
	var (
		taxRates           []TaxRate
		taxRatesInTimeSpan []TaxRate
	)

	for i, tax := range taxes.Taxes {
		if tax.Municipality == municipality {
			taxRates = taxes.Taxes[i].TaxRates
			for ii, taxRate := range taxRates {
				startDate, err := time.Parse("2006-01-02", taxRate.StartDate)

				if err != nil {
					return nil, err
				}
				endDate, err := time.Parse("2006-01-02", taxRate.EndDate)
				if err != nil {
					return nil, err
				}
				date, err := time.Parse("2006-01-02", dateStr)
				if err != nil {
					return nil, err
				}

				startDate = startDate.Add(time.Second * time.Duration(-1))
				endDate = endDate.Add(time.Hour*time.Duration(23) + time.Minute*time.Duration(59) + time.Second*time.Duration(59))

				if inTimeSpan(startDate, endDate, date) {
					taxRatesInTimeSpan = append(taxRatesInTimeSpan, taxes.Taxes[i].TaxRates[ii])
					//return &taxes.Taxes[i].TaxRates[ii], nil
				}
			}
		}
	}

	if len(taxRatesInTimeSpan) > 0 {
		for _, taxRateInTimeSpan := range taxRatesInTimeSpan {
			if taxRateInTimeSpan.Type == "daily" {
				return &taxRateInTimeSpan, nil
			}
			if taxRateInTimeSpan.Type == "weekly" && !contains(taxRatesInTimeSpan, "daily") {
				return &taxRateInTimeSpan, nil
			}
			if taxRateInTimeSpan.Type == "monthly" && !contains(taxRatesInTimeSpan, "daily") && !contains(taxRatesInTimeSpan, "weekly") {
				return &taxRateInTimeSpan, nil
			}
			if taxRateInTimeSpan.Type == "yearly" && !contains(taxRatesInTimeSpan, "daily") && !contains(taxRatesInTimeSpan, "weekly") && !contains(taxRatesInTimeSpan, "monthly") {
				return &taxRateInTimeSpan, nil
			}
		}
	}

	return nil, nil
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func contains(taxRates []TaxRate, taxRateType string) bool {
	for _, taxRate := range taxRates {
		if taxRate.Type == taxRateType {
			return true
		}
	}
	return false
}

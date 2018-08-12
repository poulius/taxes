package taxes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"time"
)

func (taxes *Taxes) Update(taxType, startDateStr, endDateStr, taxRateValueStr, municipalityID, taxRateIDStr string) error {
	var (
		taxRateExists bool = false
	)

	if taxType == "" || startDateStr == "" || endDateStr == "" || taxRateValueStr == "" || municipalityID == "" || taxRateIDStr == "" {
		return errors.New("Required params are not provided")
	}

	taxRateValue, err := strconv.ParseFloat(taxRateValueStr, 64)
	if err != nil {
		return err
	}

	// just check if date format is ok
	dateFormat := "2006-01-02"
	_, err = time.Parse(dateFormat, startDateStr)
	if err != nil {
		return err
	}
	_, err = time.Parse(dateFormat, endDateStr)
	if err != nil {
		return err
	}

	taxRateID, err := strconv.Atoi(taxRateIDStr)
	if err != nil {
		return err
	}

	for i, tax := range taxes.Taxes {
		if tax.MunicipalityID == municipalityID {
			for j, taxRate := range taxes.Taxes[i].TaxRates {
				if taxRate.TaxRateID == taxRateID {
					taxes.Taxes[i].TaxRates[j] = TaxRate{TaxRateID: taxRateID, Type: taxType, StartDate: startDateStr, EndDate: endDateStr, TaxRateValue: taxRateValue}
					taxRateExists = true
				}
			}
		}
	}

	if !taxRateExists {
		return errors.New("Data could not be updated because provided municipality and tax rate does not exist.")
	}

	jsonFile, err := json.Marshal(taxes)
	if err != nil {
		return err
	}

	byteData := []byte(jsonFile)
	err = ioutil.WriteFile("data.json", byteData, 0644)
	if err != nil {
		return err
	}

	return nil
}

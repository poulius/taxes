package taxes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"time"
)

func (taxes *Taxes) Add(municipality, taxType, startDateStr, endDateStr, taxRateValueStr string) error {
	var (
		municipalityExists bool = false
	)

	if municipality == "" || taxType == "" || startDateStr == "" || endDateStr == "" || taxRateValueStr == "" {
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

	for i, tax := range taxes.Taxes {
		if tax.Municipality == municipality {
			// check if the same record exists in json file
			for _, taxRate := range taxes.Taxes[i].TaxRates {
				if taxRate.Type == taxType && taxRate.StartDate == startDateStr && taxRate.EndDate == endDateStr {
					return errors.New("For this municipality and date tada is already submited")
				}
			}

			taxes.Taxes[i].TaxRates = append(taxes.Taxes[i].TaxRates, TaxRate{TaxRateID: len(taxes.Taxes[i].TaxRates), Type: taxType, StartDate: startDateStr, EndDate: endDateStr, TaxRateValue: taxRateValue})
			municipalityExists = true
		}
	}

	if !municipalityExists {
		return errors.New("Provided municipality does not exist")
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

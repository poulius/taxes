package taxes

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func (taxes *Taxes) Get() error {

	jsonFile, err := os.Open("data.json")
	if err != nil {
		return err
	} else {
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			return err
		} else {
			err = json.Unmarshal(byteValue, taxes)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

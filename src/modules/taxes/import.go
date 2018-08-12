package taxes

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

func (taxes *Taxes) Import() error {
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		municipality := line[0]
		taxType := line[1]
		startDateStr := line[2]
		endDateStr := line[3]
		taxRateStr := line[4]

		taxes.Add(municipality, taxType, startDateStr, endDateStr, taxRateStr)

	}

	return nil
}

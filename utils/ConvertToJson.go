package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"
)

func validateAndBeautifyJSON(jsonStr string) (string, error) {
	// Validate the JSON by unmarshalling it into a generic map
	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return "", err // Return the error if JSON is invalid
	}

	// Beautify (pretty-print) the JSON using json.MarshalIndent
	beautified, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	// Return the beautified JSON as a string
	return string(beautified), nil
}

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func ConvertToJson(file multipart.File) (error, string) {
	// Read the CSV file
	csvReader := csv.NewReader(file)
	csvReader.LazyQuotes = true
	csvReader.TrimLeadingSpace = true
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return err, ""
	}

	// Convert records to JSON
	type RecordData map[string]string

	headers := records[0]

	var csvData []RecordData

	for _, record := range records[1:] {

		row := make(RecordData)

		for i, cell := range record {
			row[headers[i]] = cell
		}
		csvData = append(csvData, row)
	}

	// Convert JSON data to string
	jsonDataString, err := json.Marshal(csvData)

	r, err := validateAndBeautifyJSON(string(jsonDataString))

	fmt.Println(r)

	//return err,
	return err, r

}

package utils

import "time"

const (
	DateYMDFormatHyphen = "2006-01-02"
	DateDMYFormatSlash = "02/01/2006"
)

func FormatDate(inputDateString string, inputFormat string, outputFormat string) (string, error) {
	inputDate, err := time.Parse(inputFormat, inputDateString)
	if err != nil {
		return "", err
	}

	return inputDate.Format(outputFormat), nil
}
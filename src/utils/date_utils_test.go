package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	dateYMDFormatHyphen       = "2021-03-17"
	dateDMYFormatSlash        = "17/03/2021"
	dateWithErrorFormatHyphen = "2021&03&17"
)

func TestFormatDate(t *testing.T) {
	t.Parallel()

	t.Run("when FormatDate executes successfully", func(t *testing.T) {
		formattedDate, err := FormatDate(dateYMDFormatHyphen, DateYMDFormatHyphen, DateDMYFormatSlash)

		assert.NoError(t, err)
		assert.Equal(t, dateDMYFormatSlash, formattedDate)
	})

	t.Run("when FormatDate executes with an inputDate with error", func(t *testing.T) {
		formattedDate, err := FormatDate(dateWithErrorFormatHyphen, DateYMDFormatHyphen, DateDMYFormatSlash)

		assert.Error(t, err)
		assert.Equal(t, "", formattedDate)
		assert.Equal(t, "parsing time \"2021&03&17\" as \"2006-01-02\": cannot parse \"&03&17\" as \"-\"", err.Error())
	})
}

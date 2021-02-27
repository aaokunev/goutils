package str

import (
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
	"unicode"
)

func FirstToLower(str string) string {
	if len(str) == 0 {
		return str
	}
	result := []rune(str)
	result[0] = unicode.ToLower(result[0])
	return string(result)
}

func GetValue(value *string, defaultValue string) string {
	if value == nil {
		return defaultValue
	}

	return *value
}

func GetPointerOrNil(value string) *string {
	if len(value) == 0 {
		return nil
	}

	return &value
}

func ConvertStrToDate(str string, dateFormat string) *time.Time {
	if len(str) == 0 {
		return nil
	}

	datetime, err := time.Parse(dateFormat, str)

	if err != nil {
		panic(fmt.Sprintf("unable to parse date: %s", str))
	}

	return &datetime
}

func ConvertDateToStr(date *time.Time, dateFormat string) string {
	if date == nil {
		return ""
	}

	return date.Format(dateFormat)
}

func RemoveLineBreaks(str string) string {
	tmp := strings.Replace(str, "\n", "", -1)
	tmp = strings.Replace(tmp, "\r", "", -1)

	return tmp
}

func MustConvertToUUIDv4(str string) uuid.UUID {
	result, err := uuid.FromString(str)

	if err != nil {
		panic(errors.New(fmt.Sprintf("Unable to to convert string to UUIDv4: `%s`", str)))
	}

	return result
}

func MustConvertToUUIDv4Slice(strList []string) []uuid.UUID {
	if strList == nil || len(strList) == 0 {
		return make([]uuid.UUID, 0)
	}

	result := make([]uuid.UUID, 0, len(strList))

	for _, str := range strList {
		result = append(result, MustConvertToUUIDv4(str))
	}

	return result
}

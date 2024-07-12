package cases

import (
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

func ToSentence(s string) string {
	delimiter, _ := strconv.Atoi(" ")
	sentence := strcase.ToDelimited(s, uint8(delimiter))
	return strings.ToUpper(sentence[:1]) + sentence[1:]
}

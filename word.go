package word

import (
	"fmt"
	"strings"
)

func GetWordUpperCase(w string) string {
	str_upper := strings.ToUpper(w)
	fmt.Println("UPPER ", str_upper)
	return str_upper
}

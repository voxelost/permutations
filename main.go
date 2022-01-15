package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	prefix := "--"

	separator := GetOSArg(prefix, []string{"sep", "separator"}, " ")

	word_len, err := getOSArgInt(prefix, []string{"len", "length"}, 3)
	if err != nil {
		panic(err.Error())
	}

	charset := GetOSArg(prefix, []string{"chrs", "charset"}, "abcdefghijklmnopqrstuvwxyz")

	fmt.Println(generatePermutationsString(int(word_len), charset, separator))
}

func getOSArgInt(prefix string, keys []string, defaultValue int64) (int64, error) {
	noMatchString := ""
	valueStr := GetOSArg(prefix, keys, noMatchString)

	if valueStr != noMatchString {
		return strconv.ParseInt(valueStr, 10, 64)
	}

	return defaultValue, nil
}

func GetOSArg(prefix string, keys []string, defaultValue string) string {
	for _, arg := range os.Args[1:] {
		parsed := strings.Split(arg, "=")

		if strings.HasPrefix(parsed[0], prefix) {
			for _, key := range keys {
				if parsed[0][len(prefix):] == key {
					return parsed[1]
				}
			}
		}
	}

	return defaultValue
}

func generatePermutationsString(perm_len int, charset string, sep string) string {
	var out []string

	possibilities := math.Pow(float64(len(charset)), float64(perm_len))
	for i := 0; i < int(possibilities); i++ {
		var tempout string
		possibility := i

		for j := 0; j < perm_len; j++ {
			ch := possibility % len(charset)
			tempout = fmt.Sprintf("%c%s", charset[ch], tempout)
			possibility = possibility / len(charset)
		}

		out = append(out, tempout)
	}

	return strings.Join(out, sep)
}

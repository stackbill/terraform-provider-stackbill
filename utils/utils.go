package utils

import (
	"encoding/json"
	"regexp"
	"strings"
	"terraform-provider-stackbill/config"

	"github.com/TylerBrock/colorjson"
)

func GetProviderName() string {
	return config.Provider
}

// Camel case string to snake case
func CamelCaseStringToSnakeCase(camelCase string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(camelCase, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// Format json string
func FormatJsonString(jsonStr string) string {
	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4
	var obj map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &obj)
	s, _ := f.Marshal(obj)
	return string(s)
}

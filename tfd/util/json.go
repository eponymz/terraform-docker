package util

import "github.com/tidwall/gjson"

func ValidateJSON(json string) bool {
	return gjson.Valid(json)
}

func GetJSON(json string, path string) gjson.Result {
	result := gjson.Get(json, path)
	return result
}

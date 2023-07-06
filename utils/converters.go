package utils

import "strings"

func GetLabelAsStringWithSeparator(label map[string]string) string {
	returnableString := ""
	for k, v := range label {
		returnableString += k + ":" + v + ";"
	}
	return returnableString[:len(returnableString)-1]
}

func GetKeyIndexInfo(valueLookingFor string, key string) string {
	separatedString := strings.Split(key, "/")
	switch valueLookingFor {
	case "groupID":
		return separatedString[1]
	case "groupVersion":
		return separatedString[2]
	default:
		return ""
	}
}

package common

import (
	"errors"
	"strings"
)

func StrIsInSlice(element string, slice []string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

func StrRemoveProtocol(url string) (string, error) {
	index := strings.Index(url, "://")
	if index != -1 {
		return url[index+3:], nil
	}
	return "", errors.ErrUnsupported
}

func StrRemoveSpaces(input string) string {
	return strings.ReplaceAll(input, " ", "")
}

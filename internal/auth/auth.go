package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error){
	val := headers.Get("Authorization")
	if val==""{
		return "", errors.New("no authorization info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2{
		return "", errors.New("Malformed auth header")
	}
	if vals[0] != "ApiKey"{
		return "", errors.New("first part of auth header is wrong")
	}

	return vals[1], nil
}


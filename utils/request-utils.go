package utils

import (
	"errors"
	"log"
	"net/http"

	"github.com/XenZi/ARS-2022-23/idempotency"
)

func DoesKeyExistInTheCurrentSessionOfRequests(req *http.Request) (bool, error) {
	idempotencyMap := *idempotency.GetIdempotencyMap()
	key := req.Header.Get("x-idempotency-key")
	log.Println(key)
	val, err := idempotencyMap[key]
	log.Println(err)
	if err == true {
		return true, errors.New("Key is already existing")
	}
	idempotencyMap[key] = val
	return false, nil
}

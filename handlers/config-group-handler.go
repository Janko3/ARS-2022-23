package handlers

import (
	"fmt"
	"net/http"
)

func AddConfigGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}

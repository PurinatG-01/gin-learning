package handler

import "net/http"

type EventPort interface {
	Test(w http.ResponseWriter, r *http.Request)
}

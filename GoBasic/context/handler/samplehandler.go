package handler

import "net/http"

func hello(req *http.Request) {
	req.Context()
}

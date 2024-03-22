package main

import (
	"log"
	"net/http"
)

func TagsGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make([]map[string]string, 2)
		data[0] = map[string]string{
			"tag_id": "c96b47d7-c537-412b-a480-e8e793196650",
			"pass":   "H@EGOHDSKMV",
		}
		data[1] = map[string]string{
			"tag_id": "38091eee-25b4-43f0-b0b7-79355bfc2901",
			"pass":   "SDKFJASD!@#",
		}
		if err := encode(w, r, 200, data); err != nil {
			log.Println(err)
		}
	}
}

package handlers

import "github.com/gorilla/schema"

var queryDecoder = schema.NewDecoder()

func init() {
	queryDecoder.IgnoreUnknownKeys(true)
}

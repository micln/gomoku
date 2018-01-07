package main

import (
	"fmt"
)

type AjaxResponse struct {
	data interface{}
}

func newAjaxResponse(data interface{}) *AjaxResponse {
	return &AjaxResponse{data: data}
}

func (r *AjaxResponse) String() string {
	s := ""
	s += fmt.Sprintf("%v", r.data)
	return s
}

func AjaxSuccess(content interface{}, messages ...string) interface{} {
	return Ajax(true, content, messages...)
}
func AjaxError(content interface{}, messages ...string) interface{} {
	return Ajax(false, content, messages...)
}

func Ajax(success bool, content interface{}, messages ...string) interface{} {
	m := make(map[string]interface{})
	m[`success`] = success
	m[`content`] = content

	message := ``
	if len(messages) > 0 {
		message = messages[0]
	}

	m[`message`] = message

	return m
}

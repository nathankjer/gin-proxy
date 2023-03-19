package models

import "time"

type Request struct {
	Id             string    `json:"id"`
	Method         string    `json:"method"`
	Host           string    `json:"host"`
	Path           string    `json:"path"`
	ResponseStatus int       `json:"response_status"`
	ResponseBody   []byte    `json:"response_body"`
	CreatedAt      time.Time `json:"expires_at"`
}

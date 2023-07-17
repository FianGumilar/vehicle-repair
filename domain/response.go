package domain

import "time"

type ApiResponse struct {
	TraceID   string      `json:"trace_id"`
	CreatedAt time.Time   `json:"created_at"`
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

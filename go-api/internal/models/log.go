package models

type LogData struct {
	TimeStamp string `json:"timestamp" binding:"required"`
	Level     string `json:"level" binding:"required"`
	Service   string `json:"service" binding:"required"`
	Message   string `json:"message" binding:"required"`
}

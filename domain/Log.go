package domain

import "time"

type Level string

const (
	INFO Level = "Info"
	DEBUG Level = "Debug"
	ERROR Level = "Error"
)


type Log struct {
	Name string `json:"name"`
	Date time.Time `json:"date"`
	Message string `json:"message"`
	Level Level `json:"level"`
}

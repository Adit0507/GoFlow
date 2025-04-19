package main

import (
	"context"
	"log-service/data"
	"log-service/logs"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// writin log
	logEntry := data.LogEntry {
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		ans := &logs.LogResponse{Result: "failed"}
		return ans, err
	}

	ans := &logs.LogResponse{Result: "logged!"}

	return ans, nil
}

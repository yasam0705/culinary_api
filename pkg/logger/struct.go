package logger

import "go.uber.org/zap"

type Logger interface {
	Error(msg string, fields map[string]interface{})
	Debug(msg string, fields map[string]interface{})
	Info(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})
	Panic(msg string, fields map[string]interface{})
	Close()
}

type logger struct {
	log *zap.Logger
}

func (l *logger) Error(msg string, fields map[string]interface{}) {
	l.log.Error(msg, generateFields(fields)...)
}

func (l *logger) Debug(msg string, fields map[string]interface{}) {
	l.log.Debug(msg, generateFields(fields)...)
}

func (l *logger) Info(msg string, fields map[string]interface{}) {
	l.log.Info(msg, generateFields(fields)...)
}

func (l *logger) Fatal(msg string, fields map[string]interface{}) {
	l.log.Fatal(msg, generateFields(fields)...)
}

func (l *logger) Panic(msg string, fields map[string]interface{}) {
	l.log.Panic(msg, generateFields(fields)...)
}

func (l *logger) Close() {
	l.log.Sync()
}

func generateFields(fields map[string]interface{}) []zap.Field {
	var f = make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		f = append(f, zap.Any(k, v))
	}
	return f
}

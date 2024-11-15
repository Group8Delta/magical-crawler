package Logger

import (
	"magical-crwler/database"
	"magical-crwler/models"
)

type ILogger interface {
	Info(log models.Log)
	Debug(log models.Log)
	Warn(log models.Log)
	Error(log models.Log)
}
type Logger struct {
	repository database.IRepository
}

func NewLogger(repository database.IRepository) *Logger {
	return &Logger{repository: repository}
}
func (l *Logger) Info(msg string) {
	l.repository.AddLog(models.Log{
		Message:    msg,
		LogLevelID: 1,
	})
}
func (l *Logger) Debug(msg string) {
	l.repository.AddLog(models.Log{
		Message:    msg,
		LogLevelID: 2,
	})
}
func (l *Logger) Warn(msg string) {
	l.repository.AddLog(models.Log{
		Message:    msg,
		LogLevelID: 3,
	})
}
func (l *Logger) Error(msg string) {
	l.repository.AddLog(models.Log{
		Message:    msg,
		LogLevelID: 4,
	})
}

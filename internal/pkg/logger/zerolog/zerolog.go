package zerolog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	zlog zerolog.Logger
}

func New() (*Logger, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return &Logger{zlog: log.Logger}, nil
}

func (l *Logger) Info(msg string, fields ...any) {
	l.log(l.zlog.Info, msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...any) {
	l.log(l.zlog.Warn, msg, fields...)
}

func (l *Logger) Error(msg string, fields ...any) {
	l.log(l.zlog.Info, msg, fields...)
}

func (l *Logger) log(level func() *zerolog.Event, msg string, fields ...any) {
	event := level()
	for i := 0; i < len(fields)-1; i += 2 {
		if key, ok := fields[i].(string); ok {
			event = event.Interface(key, fields[i+1])
		}
	}
	event.Msg(msg)
}

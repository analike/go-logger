/**
 * @package go-zLog (2026)
 * @author Emmanuel Analike <emmanuel@analike.dev>
 * @created May 22, 2026; 8:30 AM
 */

package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
)

var zLog zerolog.Logger

func init() {
	out := zerolog.ConsoleWriter{TimeFormat: time.RFC3339, Out: os.Stdout}
	zLog = zerolog.New(out).With().Timestamp().Logger()
}

func addPrefix(msg string, p ...string) string {
	if len(p) > 0 {
		return fmt.Sprintf("[%s] %s", p[0], msg)
	}
	return msg
}

func caller() string {
	pr, file, no, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("%s:%d(%d)", file, no, pr)
	}
	return ""
}

func getMessage(val any) string {
	switch v := val.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	case error:
		return v.Error()
	default:
		return fmt.Sprintf("%+v", val)
	}
}

func Success(msg any, prefix ...string) {
	zLog.Info().Msg(addPrefix(getMessage(msg), prefix...))
}

func Info(msg any, prefix ...string) {
	zLog.Debug().Msg(addPrefix(getMessage(msg), prefix...))
}

func Error(msg any, prefix ...string) {
	zLog.Error().Str("trace", caller()).Msg(addPrefix(getMessage(msg), prefix...))
}

func Warn(msg any, prefix ...string) {
	zLog.Warn().Str("trace", caller()).Msg(addPrefix(getMessage(msg), prefix...))
	// zLog.Error().Str("trace", caller()).Msg(addPrefix(getMessage(msg), prefix...))
}

func Successf(format string, a ...any) {
	zLog.Info().Msg(fmt.Sprintf(format, a...))
}

func Infof(format string, a ...any) {
	zLog.Debug().Msg(fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...any) {
	zLog.Error().Str("trace", caller()).Msg(fmt.Sprintf(format, a...))
}

func Warnf(format string, a ...any) {
	zLog.Warn().Str("trace", caller()).Msg(fmt.Sprintf(format, a...))
}

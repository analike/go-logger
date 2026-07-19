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

// Error print ERR msg; adds a trace to the caller
func Error(msg any, prefix ...string) {
	zLog.Error().Str("trace", caller()).Msg(addPrefix(getMessage(msg), prefix...))
}

// Warn print WRN msg; adds a trace to the caller
func Warn(msg any, prefix ...string) {
	zLog.Warn().Str("trace", caller()).Msg(addPrefix(getMessage(msg), prefix...))
}

// Fatal print FTL msg; adds a trace to the caller
func Fatal(msg any, prefix ...string) {
	zLog.Fatal().Str("trace", caller()).Msg(addPrefix(getMessage(msg), prefix...))
}

// FatalMsg print FTL msg without trace to the caller
// exits program with code 1
func FatalMsg(msg any, prefix ...string) {
	zLog.Fatal().Msg(addPrefix(getMessage(msg), prefix...))
}

// ErrorMsg print ERR msg; without trace to the caller
func ErrorMsg(msg any, prefix ...string) {
	zLog.Error().Msg(addPrefix(getMessage(msg), prefix...))
}

// WarnMsg print WRN msg; without trace to the caller
func WarnMsg(msg any, prefix ...string) {
	zLog.Warn().Msg(addPrefix(getMessage(msg), prefix...))
}

// Successf prints a formatted INF message
func Successf(format string, a ...any) {
	zLog.Info().Msg(fmt.Sprintf(format, a...))
}

func Infof(format string, a ...any) {
	zLog.Debug().Msg(fmt.Sprintf(format, a...))
}

// Errorf prints formatted ERR message following defined pattern; adds trace to caller
func Errorf(format string, a ...any) {
	zLog.Error().Str("trace", caller()).Msg(fmt.Sprintf(format, a...))
}

// Fatalf prints formatted FTL message following defined pattern
// adding trace to caller and exiting with code 1
func Fatalf(format string, a ...any) {
	zLog.Fatal().Str("trace", caller()).Msg(fmt.Sprintf(format, a...))
}

// Warnf prints formatted ERR message following defined pattern; adds trace to caller
func Warnf(format string, a ...any) {
	zLog.Warn().Str("trace", caller()).Msg(fmt.Sprintf(format, a...))
}

// ErrorMsgf prints formatted ERR message following defined pattern without trace to caller
func ErrorMsgf(format string, a ...any) {
	zLog.Error().Msg(fmt.Sprintf(format, a...))
}

// WarnMsgf prints formatted WRN message following defined pattern without trace to caller
func WarnMsgf(format string, a ...any) {
	zLog.Warn().Msg(fmt.Sprintf(format, a...))
}

// FatalMsgf prints formatted FTL message following defined pattern without trace to caller
// exits the program with code 1
func FatalMsgf(format string, a ...any) {
	zLog.Fatal().Msg(fmt.Sprintf(format, a...))
}

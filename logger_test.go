/**
 * @package go-zLog (2026)
 * @author Emmanuel Analike <emmanuel@analike.dev>
 * @created May 22, 2026; 8:52 AM
 */

package logger

import (
	"errors"
	"testing"
)

var testMessage = "Lorem ipsum dolor sit amet"

func TestWarn(t *testing.T) {
	Warn(testMessage, "warn")
}

func TestInfo(t *testing.T) {
	Info(testMessage, "info")
}

func TestSuccess(t *testing.T) {
	Success(testMessage, "success")
}

func TestError(t *testing.T) {
	Error(testMessage, "error")
}

func TestWarnf(t *testing.T) {
	Warnf(testMessage+" [%s]", "warn-f")
}

func TestInfof(t *testing.T) {
	Infof(testMessage+" [%d] %+v", 64, "info-f")
}

func TestSuccessf(t *testing.T) {
	Successf("[%s] [%d] "+testMessage, "success-f", 2000)
}

func TestErrorf(t *testing.T) {
	Errorf("**%s** "+testMessage+" %+v", "error-f", errors.New("A typical error"))
}

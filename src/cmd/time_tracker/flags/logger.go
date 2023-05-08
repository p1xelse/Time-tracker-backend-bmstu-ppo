package flags

import (
	"fmt"
	"github.com/labstack/echo/v4"
	elog "github.com/labstack/gommon/log"
)

type LoggerFlags struct {
	LogLevel  uint8  `toml:"level"`
	LogHeader string `toml:"header"`
}

func (f LoggerFlags) Init(e *echo.Echo) echo.Logger {
	fmt.Println("header: ", f.LogHeader)
	fmt.Println("level: ", f.LogLevel)
	e.Logger.SetLevel(elog.Lvl(f.LogLevel))
	e.Logger.SetHeader(f.LogHeader)
	return e.Logger
}

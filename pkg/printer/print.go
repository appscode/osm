package printer

import (
	"fmt"

	"github.com/fatih/color"
)

func Info(i ...interface{}) {
	color.Set(color.FgCyan)
	fmt.Println(i...)
	color.Set(color.Reset)
}

func Version(i ...interface{}) {
	color.Set(color.FgBlue)
	fmt.Println(i...)
	color.Set(color.Reset)
}

func Error(i ...interface{}) {
	color.Set(color.FgRed, color.Bold)
	fmt.Println(i...)
	color.Set(color.Reset)
}

func Config(i ...interface{}) {
	color.Set(color.FgMagenta)
	fmt.Println(i...)
	color.Set(color.Reset)
}

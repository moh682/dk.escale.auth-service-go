package config

import "github.com/fatih/color"

var red = color.New(color.FgRed).Add(color.Underline)
var green = color.New(color.FgGreen).Add(color.Underline)

// Red only for panic
func Red(err string) {
	red.Println(err)
	panic(nil)
}

// Green only for logs
func Green(info string) {
	green.Println(info)
}

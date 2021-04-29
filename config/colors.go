package config

import "github.com/fatih/color"

var red = color.New(color.FgRed).Add(color.Underline)
var green = color.New(color.FgGreen).Add(color.Underline)
var cyan = color.New(color.FgCyan).Add(color.Underline)

// Red only for panic
func Red(err interface{}) string {
	return red.Sprint(err)
}

// Green only for logs
func Green(info interface{}) string {
	return green.Sprint(info)
}

// Cyan only for logs
func Cyan(info interface{}) string {
	return cyan.Sprint(info)
}

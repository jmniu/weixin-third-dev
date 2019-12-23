package logger

import "fmt"

func Infof(logStr string, args ...interface{}) {
	fmt.Printf("info: "+logStr+"\n", args...)
}

func Warnf(logStr string, args ...interface{}) {
	fmt.Printf("warn: "+logStr+"\n", args...)
}

func Fatalf(logStr string, args ...interface{}) {
	fmt.Printf("fatal: "+logStr+"\n", args...)
}

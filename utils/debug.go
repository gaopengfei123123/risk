package utils

import (
	"log"
	"os"
)

// 调试用
func DD(params ...interface{}) {
	log.Println(params...)
	os.Exit(0)
}

// 调试用， 带模板
func DDf(format string, params ...interface{}) {
	log.Printf(format+"\n", params...)
	os.Exit(0)
}

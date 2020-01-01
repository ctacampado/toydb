package main

import (
	"log"
	tl "toydb/pkg/toylog"
)

func main() {
	logger, err := tl.NewToyLog("toydb", tl.ERR, true)
	if nil != err {
		log.Fatal(err.Error())
	}
	defer logger.Close()

	logger.Info("hello world")
	logger.Debug("hello world with filename and line number")
	logger.Error("error message")
}

package toylog

import (
	"os"
	"testing"
)

func TestNewToyLoggerNIL(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file *os.File
	}{
		name: "logger",
		lvl:  DEBUG,
		file: os.Stdout,
	}

	_, err := NewToyLog(in.name, in.lvl, in.file, 12345)
	if nil != err {
		t.Log(err.Error())
	}
}

func TestNewToyLogNotNIL(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file *os.File
	}{
		name: "logger",
		lvl:  DEBUG,
		file: os.Stdout,
	}

	_, err := NewToyLog(in.name, in.lvl, in.file)
	if nil != err {
		t.Error(err.Error())
	}
}

func TestInfo(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file *os.File
		in   string
	}{
		name: "logger",
		lvl:  INFO,
		file: os.Stdout,
		in:   "hello world",
	}
	logger, err := NewToyLog(in.name, in.lvl, in.file)
	if nil != err {
		t.Error(err.Error())
	}
	logger.Info(in.in)
	logger.Debug(in.in)
}

func TestDebug(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file *os.File
		in   string
	}{
		name: "logger",
		lvl:  DEBUG,
		file: os.Stdout,
		in:   "hello world",
	}

	logger, err := NewToyLog(in.name, in.lvl, in.file)
	if nil != err {
		t.Error(err.Error())
	}
	logger.Debug(in.in)
}

func TestError(t *testing.T) {
	in := struct {
		name string
		lvl  LogLvl
		file *os.File
		in   string
	}{
		name: "logger",
		lvl:  ERR,
		file: os.Stdout,
		in:   "hello world",
	}

	logger, err := NewToyLog(in.name, in.lvl, in.file)
	if nil != err {
		t.Error(err.Error())
	}
	logger.Error(in.in)
}

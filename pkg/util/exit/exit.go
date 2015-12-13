package exit

import (
	"fmt"
	"os"

	"github.com/tmrts/tmplt/pkg/util/tlog"
)

const (
	CodeOK    = 0
	CodeFatal = 1
	CodeError = 2
)

func Fatal(err error) {
	tlog.Fatal(fmt.Sprint(err))

	os.Exit(CodeFatal)
}

func Error(err error) {
	tlog.Error(err.Error())

	os.Exit(CodeError)
}

func OK(fmtStr string, s ...interface{}) {
	tlog.Success(fmt.Sprintf(fmtStr, s...))

	os.Exit(CodeOK)
}
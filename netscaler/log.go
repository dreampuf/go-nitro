package netscaler

import (
	mlog "log"
	"os"
)

var log = mlog.New(os.Stdout, "[go-nitro]", mlog.Lshortfile)

func InitLog(logger *mlog.Logger) {
	log = logger
}

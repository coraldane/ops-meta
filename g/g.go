package g

import (
	"log"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

type JsonResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

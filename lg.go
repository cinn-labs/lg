package lg

import (
	"log"
	"runtime"
)

func GetTrace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func Error(err error) {
	fn := GetTrace()
	log.Println("********ERROR ON: ", fn, " - ", err)
}

func ValidationError(v interface{}, err error) {
	fn := GetTrace()
	log.Println("*VALIDATION ON: ", fn, " - ", err, " | ", v)
}

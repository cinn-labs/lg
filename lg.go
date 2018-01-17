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

func Error(err error, data ...interface{}) {
	fn := GetTrace()
	p := append([]interface{}{"********ERROR ON: ", fn, " - ", err, " | "}, data...)
	log.Println(p...)
}

func ValidationError(v interface{}, err error) {
	fn := GetTrace()
	log.Println("*VALIDATION ON: ", fn, " - ", err, " | ", v)
}

func Info(data ...interface{}) {
	fn := GetTrace()
	p := append([]interface{}{"**INFO: ", fn, " - "}, data...)
	log.Println(p...)
}

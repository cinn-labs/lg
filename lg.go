package lg

import (
	"log"
	"runtime"
)

func GetTrace(level int) string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(3+level, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func Error(err error, data ...interface{}) {
	fn := GetTrace(0)
	p := append([]interface{}{"********ERROR ON: ", fn, " - ", err, " | "}, data...)
	log.Println(p...)
}

func CError(err error, data ...interface{}) bool {
	if err != nil {
		fn := GetTrace(1)
		p := append([]interface{}{"********ERROR ON: ", fn, " - ", err, " | "}, data...)
		log.Println(p...)
	}
	return false
}

func ValidationError(v interface{}, err error) {
	fn := GetTrace(0)
	log.Println("*VALIDATION ON: ", fn, " - ", err, " | ", v)
}

func Info(data ...interface{}) {
	fn := GetTrace(0)
	p := append([]interface{}{"**INFO: ", fn, " - "}, data...)
	log.Println(p...)
}

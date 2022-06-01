package log

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Logger struct {
	log *log.Logger
}

var Log Logger

func init() {
	Log.log = log.New(os.Stdout, "", log.Ldate|log.Ltime)
}

func (l *Logger) LogRequest(r *http.Request) {
	requestInfo := fmt.Sprintf("%s %s%s ", r.Method, r.Host, r.RequestURI)
	l.LogInfo(requestInfo)

}

func (l *Logger) LogInfo(s string) {
	l.log.Println("INFO " + s)
}

func (l *Logger) LogDanger(s string) {
	l.log.Println("DANGER " + s)
}

func (l *Logger) LogWarning(s string) {
	l.log.Println("WARNING " + s)
}

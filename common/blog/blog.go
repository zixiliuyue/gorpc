package blog

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/golang/glog"
)

// This is temporary until we agree on log dirs and put those into each cmd.
func init() {
	flag.Set("logtostderr", "true")
}

// GlogWriter serves as a bridge between the standard log package and the glog package.
type GlogWriter struct{}

// Write implements the io.Writer interface.
func (writer GlogWriter) Write(data []byte) (n int, err error) {
	glog.InfoDepth(1, string(data))
	return len(data), nil
}

// Output for mgo logger
func (writer GlogWriter) Output(calldepth int, s string) error {
	glog.InfoDepth(calldepth, s)
	return nil
}

func (writer GlogWriter) Print(v ...interface{}) {
	glog.InfoDepth(1, v...)
}
func (writer GlogWriter) Printf(format string, v ...interface{}) {
	glog.InfoDepth(1, fmt.Sprintf("%v %v", format, v))
}

func (writer GlogWriter) Println(v ...interface{}) {
	glog.InfoDepth(1, v...)
}

var once sync.Once

// InitLogs initializes logs the way we want for blog. 是一个可以被多次调用但是只执行一次，若每次调用Do时传入参数f不同，但是只有第一个才会被执行。
//  log 等级分为４类：Info, Warning, Error, Fatal.
func InitLogs() {
	once.Do(func() {
		log.SetOutput(GlogWriter{})
		log.SetFlags(0)
		// The default glog flush interval is 30 seconds, which is frighteningly long.
		go func() {
			d := time.Duration(5 * time.Second)
			tick := time.Tick(d)
			for {
				select {
				case <-tick:
					glog.Flush()
				}
			}
		}()
	})
}

func CloseLogs() {
	glog.Flush()
}

var (
	Info        = glog.Infof
	Infof       = glog.Infof
	InfofDepthf = glog.Infof

	Warn  = glog.Warningf
	Warnf = glog.Warningf

	Error       = glog.Errorf
	Errorf      = glog.Errorf
	ErrorfDepth = glog.ErrorDepth

	Fatal  = glog.Fatal
	Fatalf = glog.Fatalf

	V = glog.V
)

func Debug(args ...interface{}) {
	if format, ok := (args[0]).(string); ok {
		glog.InfoDepth(1, fmt.Sprintf("%v %v", format, args[1:]))
	} else {
		glog.InfoDepth(1, args)
	}
}

func InfoJSON(format string, args ...interface{}) {
	params := []interface{}{}
	for _, arg := range args {
		if f, ok := arg.(errorFunc); ok {
			params = append(params, f.Error())
			continue
		}
		if f, ok := arg.(stringFunc); ok {
			params = append(params, f.String())
			continue
		}
		out, err := json.Marshal(arg)
		if err != nil {
			params = append(params, err.Error())
		}
		params = append(params, out)
	}
	glog.InfoDepth(1, fmt.Sprintf("%v %v", format, params))
}

func ErrorJSON(format string, args ...interface{}) {
	params := []interface{}{}
	for _, arg := range args {
		if f, ok := arg.(errorFunc); ok {
			params = append(params, f.Error())
			continue
		}
		if f, ok := arg.(stringFunc); ok {
			params = append(params, f.String())
			continue
		}
		out, err := json.Marshal(arg)
		if err != nil {
			params = append(params, err.Error())
		}
		params = append(params, out)
	}
	glog.ErrorDepth(1, fmt.Sprintf(format, params...))
}

type errorFunc interface {
	Error() string
}
type stringFunc interface {
	String() string
}

package logger

import (
	"runtime"
	"strings"
	"log"
	"fmt"
	"config/ini"
	"os"
	"io"
	"time"
)

const (
	PanicLev int = iota
	FatalLev
	ErrorLev
	WarnLev
	InfoLev
	DebugLev
)

var logLevel int
var defaultLogLevel = 4

func Init() {
	var err error
	var logPath string
	logLevel, err = ini.GetConfigToInt("Log", "logLevel")
	if err != nil {
		logLevel = defaultLogLevel
		Error(err.Error() + ",设为默认值:%d", defaultLogLevel)
	}
	logPath, err = ini.GetConfig("Log", "logPath")
	if err != nil {
		Error(err.Error())
	} else if strings.HasSuffix(logPath, "/") {
		logPath += time.Now().Format("20160102") + ".log"
	} else {
		logPath += "/" + time.Now().Format("20160102") + ".log"
	}
	_, err = os.Stat(logPath)
	if err != nil {
		var logFile *os.File
		logFile, err := os.Create(logPath)
		if err != nil {
			fmt.Println(err)
		}

		writers := []io.Writer{
			logFile,
			os.Stdout,
		}
		fileAndStdWriter := io.MultiWriter(writers...)
		log.SetOutput(fileAndStdWriter)
	} else {
		var logFile *os.File
		logFile, err = os.OpenFile(logPath, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0660)
		if err != nil {
			fmt.Println(err)
		}

		writers := []io.Writer{
			logFile,
			os.Stdout,
		}
		fileAndStdWriter := io.MultiWriter(writers...)
		log.SetOutput(fileAndStdWriter)
	}
}

func Info(format string, args ...interface{}) {
	if logLevel >= InfoLev {
		goid := goid()
		log.SetPrefix("[INFO] [Thread-" + goid + "] ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		format = fmt.Sprintf(format, args...)
		log.Output(2, format)
	}
}

func Debug(format string, args ...interface{}) {
	if logLevel >= DebugLev {
		goid := goid()
		log.SetPrefix("[DEBUG] [Thread-" + goid + "] ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		format = fmt.Sprintf(format, args...)
		log.Output(2, format)
	}
}

func Warn(format string, args ...interface{}) {
	if logLevel >= WarnLev {
		goid := goid()
		log.SetPrefix("[WARN] [Thread-" + goid + "] ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		format := fmt.Sprintf(format, args...)
		log.Output(2, format)
	}
}

func Error(format string, args ...interface{}) {
	if logLevel >= ErrorLev {
		goid := goid()
		log.SetPrefix("[ERROR] [Thread-" + goid + "] ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		format := fmt.Sprintf(format, args...)
		log.Output(2, format)
	}
}

func Fatal(format string, args ...interface{}) {
	if logLevel >= FatalLev {
		goid := goid()
		log.SetPrefix("[FATAL] [Thread-" + goid + "] ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		format := fmt.Sprintf(format, args...)
		log.Output(2, format)
	}
}

func goid() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idStr := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]

	return idStr
}

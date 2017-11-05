/*
logger 是系统日志的封装，主要在之上封装了Error，Info两个函数。并提供了跨日期
自动分割日志文件的功能。
可以在InitLogging 后直接使用logger.Error, logger.Info操作默认的日志对象。
也可以用logger.New 创建一个自己的日志对象。
*/
package log

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//logging 是一个默认的日志对象，提供全局的Error, Info函数供使用，必须调用InitLogging
//函数进行初始化
var logging *Logger

//初始化默认的日志对象，初始化后，就能使用Error，Info函数记录日志
func InitLogging(input_filename string) {
	var logFilePath = "./logs/" + input_filename
	var logFileFolder = filepath.Dir(logFilePath)
	os.Mkdir(logFileFolder, os.ModePerm)
	logFile, err := os.OpenFile("./logs/"+input_filename,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	logging = newLogger(logFile, log.Ldate|log.Ltime)
	logging.filename = input_filename
	logging.curFile = logFile
	logging.flag = log.Ldate | log.Ltime
	logging.runtimeCaller = 2
	logging.logFilePath = true
	logging.logFunc = false
	logging.startDate = time.Now().Format("2006-01-02")
}

//Error 默认日志对象方法，记录一条错误日志，需要先初始化
func Error(format string, v ...interface{}) {
	fmt.Printf("[Error] "+format+"\n", v...)
	logging.Error(format, v...)
}

//Errorln 默认日志对象方法，记录一条消息日志，需要先初始化
func Errorln(format string) {
	fmt.Println("[Error] " + format + "\n")
	logging.Errorln(format)
}

//Info 默认日志对象方法，记录一条消息日志，需要先初始化
func Info(format string, v ...interface{}) {
	fmt.Printf("[Info] "+format+"\n", v...)
	logging.Info(format, v...)
}

//Infoln 默认日志对象方法，记录一条消息日志，需要先初始化
func Infoln(format string) {
	fmt.Printf("[Info] " + format + "\n")
	logging.Infoln(format)
}

type Logger struct {
	level         int
	err           *log.Logger
	info          *log.Logger
	curFile       *os.File
	startDate     string
	filename      string
	runtimeCaller int
	logFilePath   bool
	logFunc       bool
	flag          int
	l             sync.Mutex // 锁住curFile文件的修改
}

func newLogger(w io.Writer, flag int) *Logger {
	logObj := new(Logger)
	logObj.err = log.New(w, "[ERROR] ", flag)
	logObj.info = log.New(w, "[INFO] ", flag)

	return logObj
}

//New 创建一个自己的日志对象。
func New(filename string, logFilePath bool, logFunc bool, flag int) *Logger {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	logFile, err := os.OpenFile(dir+"/logs/"+filename,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}

	result := newLogger(logFile, flag)
	result.filename = filename
	result.curFile = logFile
	result.runtimeCaller = 1
	result.logFilePath = logFilePath
	result.logFunc = logFunc
	result.startDate = time.Now().Format("2006-01-02")
	return result
}

//Error 记录一条错误日志
func (logobj *Logger) Error(format string, v ...interface{}) {
	logobj.willChange()

	logobj.l.Lock()
	defer logobj.l.Unlock()
	var buf bytes.Buffer
	funcName, file, line, ok := runtime.Caller(logobj.runtimeCaller)
	if ok {
		if logobj.logFilePath {
			buf.WriteString(filepath.Base(file))
			buf.WriteString(":")
			buf.WriteString(strconv.Itoa(line))
			buf.WriteString(" ")
		}
		if logobj.logFunc {
			buf.WriteString(runtime.FuncForPC(funcName).Name())
			buf.WriteString(" ")
		}
		buf.WriteString(format)
		format = buf.String()
	}
	logobj.err.Printf(format, v...)
}

//Errorln 打印一行错误日志
func (logobj *Logger) Errorln(format string) {
	logobj.willChange()

	logobj.l.Lock()
	defer logobj.l.Unlock()
	var buf bytes.Buffer
	funcName, file, line, ok := runtime.Caller(logobj.runtimeCaller)
	if ok {
		if logobj.logFilePath {
			buf.WriteString(filepath.Base(file))
			buf.WriteString(":")
			buf.WriteString(strconv.Itoa(line))
			buf.WriteString(" ")
		}
		if logobj.logFunc {
			buf.WriteString(runtime.FuncForPC(funcName).Name())
			buf.WriteString(" ")
		}
		buf.WriteString(format)
		format = buf.String()
	}
	logobj.err.Println(format)
}

//Info 记录一条消息日志
func (logobj *Logger) Info(format string, v ...interface{}) {
	logobj.willChange()

	logobj.l.Lock()
	defer logobj.l.Unlock()
	var buf bytes.Buffer
	funcName, file, line, ok := runtime.Caller(logobj.runtimeCaller)
	if ok {
		if logobj.logFilePath {
			buf.WriteString(filepath.Base(file))
			buf.WriteString(":")
			buf.WriteString(strconv.Itoa(line))
			buf.WriteString(" ")
		}
		if logobj.logFunc {
			buf.WriteString(runtime.FuncForPC(funcName).Name())
			buf.WriteString(" ")
		}
		buf.WriteString(format)
		format = buf.String()
	}
	logobj.info.Printf(format, v...)
}

//Infoln打印一行消息日志
func (logobj *Logger) Infoln(format string) {
	logobj.willChange()

	logobj.l.Lock()
	defer logobj.l.Unlock()
	var buf bytes.Buffer
	funcName, file, line, ok := runtime.Caller(logobj.runtimeCaller)
	if ok {
		if logobj.logFilePath {
			buf.WriteString(filepath.Base(file))
			buf.WriteString(":")
			buf.WriteString(strconv.Itoa(line))
			buf.WriteString(" ")
		}
		if logobj.logFunc {
			buf.WriteString(runtime.FuncForPC(funcName).Name())
			buf.WriteString(" ")
		}
		buf.WriteString(format)
		format = buf.String()
	}
	logobj.info.Println(format)
}

func (logobj *Logger) changeFile() {
	logobj.l.Lock()
	defer logobj.l.Unlock()
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := dir + "/logs/" + logobj.filename
	if logobj.curFile != nil {
		logobj.curFile.Close()
		nowTime := time.Now()
		time1dAgo := nowTime.Add(-1 * time.Hour * 24)
		os.Rename(filePath, filePath+"."+time1dAgo.Format("2006-01-02"))
	}
	logFile, _ := os.Create(dir + "/logs/" + logobj.filename)
	logobj.curFile = logFile
	logobj.err = log.New(logobj.curFile, "[ERROR]", logobj.flag)
	logobj.info = log.New(logobj.curFile, "[INFO]", logobj.flag)
}

func (logobj *Logger) willChange() {
	//跨日改时间
	nowDate := time.Now().Format("2006-01-02")
	if nowDate != logobj.startDate {
		logobj.startDate = nowDate
		logobj.changeFile()
	}
}

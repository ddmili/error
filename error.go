package error

import (
	"fmt"
	"runtime"
	"time"
)

type errorString struct {
	s string
}

type errorInfo struct {
	Time    string `json:"time"`
	Alarm   string `json:"alarm"`
	Message string `json:"message"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Func    string `json:"func"`
}

func (e *errorString) Error() string {
	return e.Error()
}

// NewInfo 生成错误实例
func NewInfo(text string) error {
	print("info", text)
	return fmt.Errorf(text)
}

// NewError 生成错误实例
func NewError(text string) error {
	print("error", text)
	return fmt.Errorf(text)
}

// NewFail 生成错误实例
func NewFail(text string) error {
	print("fail", text)
	return fmt.Errorf(text)
}

// NewPanic 生成错误实例
func NewPanic(text string) error {
	print("panic", text)
	return fmt.Errorf(text)
}

// print 打印错误信息
func print(level string, str string) {
	// 定义 文件名、行号、方法名
	_, fileName, line, ok := runtime.Caller(2)
	if ok {
		fmt.Printf("%v | %v | %v (%v:%v)\r\n", time.Now().Format("2006-1-2 15:04:05 "), level, str, fileName, line)
	}

}

// This sample program demonstrates how to use the base log package.
package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	/*
		Ldate = 1 << iota: 日期(YYYY/MM/DD)
		Ltime: 时间(HH:MM:SS)
		Lmicroseconds: 毫秒级时间(HH:MM:SS.NNNNN)
		Llongfile: 完整文件路径与行号(/a/b/c/d.go:10)
		Lshortfile: 最终文件名与行号(d.go:10)
		LstdFlags = Ldate | Ltime: 标准日志记录器初始值
	*/
}

func main() {
	// Println writes to the standard logger.
	// 仅写日志
	log.Println("message")

	// Fatalln is Println() followed by a call to os.Exit(1).
	// 写完日志后终止程序
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic().
	// 写完日志后打印调用栈并终止程序(除非程序执行recover函数)
	log.Panicln("panic message")

	// 对应格式化文本方法
	log.Printf("")
	log.Fatalln("")
	log.Panicln("")
}

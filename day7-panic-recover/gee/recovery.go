package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// 输出错误堆栈
func trace(message string) string {
	// 创建一个固定大小为32的数组来存储程序计数器(PC)值
	var pcs [32]uintptr
	// 获取调用栈信息，从第3层开始（跳过runtime.Callers和trace函数本身）
	n := runtime.Callers(3, pcs[:])
	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	// 遍历所有调用栈帧
	for _, pc := range pcs[:n] {
		// 获取每个PC对应的函数信息
		fn := runtime.FuncForPC(pc)
		// 获取文件名和行号
		file, line := fn.FileLine(pc)
		// 将信息写入字符串构建器
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

// 错误恢复中间件
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}

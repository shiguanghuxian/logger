/*
 * @Author: 时光弧线
 * @Date: 2017-09-08 16:14:36
 * @Last Modified by: 时光弧线
 * @Last Modified time: 2017-09-08 17:17:56
 */
package logger

import (
	"fmt"
	"os"
	"time"
)

// FormatFileName 生成文件名
type FormatFileName func(num ...int) (string, error)

// 默认日志格式化文件名
// return 文件路径 错误
func defaultFormatFileName(num ...int) (string, error) {
	_, err := os.Stat("logs")
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir("logs", 0777)
		if err != nil {
			return "", err
		}
	}
	var pathName string
	if len(num) > 0 {
		pathName = fmt.Sprintf("logs%serror_%s_%d.log", string(os.PathSeparator), time.Now().Format("2006-01-02"), num[0])
	} else {
		pathName = fmt.Sprintf("logs%serror_%s.log", string(os.PathSeparator), time.Now().Format("2006-01-02"))
	}
	// 判断文件是否存在
	_, err = os.Stat(pathName)
	if err != nil {
		if os.IsExist(err) {
			if len(num) > 0 {
				pathName, err = defaultFormatFileName(num[0] + 1)
			} else {
				pathName, err = defaultFormatFileName(1)
			}
		}
	}
	return pathName, nil
}

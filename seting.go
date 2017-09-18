/*
 * @Author: 时光弧线
 * @Date: 2017-09-08 17:09:41
 * @Last Modified by: 时光弧线
 * @Last Modified time: 2017-09-08 17:11:41
 */
package logger

/* 设置参数 */

// SetMaxFileSize 设置日志文件最大值
func SetMaxFileSize(size int64) {
	logger.maxFileSize = size
}

// SetFormatFileName 设置生成日志文件名规则
func SetFormatFileName(formatFileName FormatFileName) {
	logger.formatFileName = formatFileName
}

// SetTimedTask 设置定时检查文件时间
func SetTimedTask(timedTask int) {
	logger.timedTask = timedTask
}

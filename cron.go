/*
 * @Author: 时光弧线
 * @Date: 2017-09-08 16:59:11
 * @Last Modified by: 时光弧线
 * @Last Modified time: 2017-09-08 17:02:44
 */
package logger

import (
	"fmt"
	"log"

	"github.com/robfig/cron"
)

var (
	logCron *cron.Cron
)

// 定时任务，1每夜更新、2定时检查文件
func initCron() {
	logCron = cron.New()
	// 每夜执行
	logCron.AddFunc("@midnight", func() {
		err := ResetLogger()
		if err != nil {
			log.Println(err)
		}
	})
	// 定时检查文件大小
	logCron.AddFunc(fmt.Sprintf("@every %ds", logger.timedTask), func() {
		err := ResetLogger()
		if err != nil {
			log.Println(err)
		}
	})
	logCron.Start()
}

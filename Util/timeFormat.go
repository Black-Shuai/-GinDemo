package Util

import (
	"time"
)
//时间初始化
func InitTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

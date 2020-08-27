package common

import (
	"alarm/utils"
	"github.com/astaxie/beego/cache"
)

const JobPrefix = "Job_"

var CacheMemory cache.Cache

var AlarmCronTab *utils.Crontab

func DelJobById(jobId string) {
	if AlarmCronTab.IsExists(jobId) {
		AlarmCronTab.DelByID(jobId)
	}
}

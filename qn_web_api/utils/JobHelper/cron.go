package JobHelper

import (
	"lisijie/cron"
	"sync"

	"github.com/astaxie/beego"
)

var (
	mainCron *cron.Cron
	workPool chan bool
	lock     sync.Mutex
)

func init() {
	size := 10
	if size > 0 {
		workPool = make(chan bool, size)
	}
	mainCron = cron.New()
	mainCron.Start()
}

/*
*添加任务
@spec 任务时间
@job 任务
*/
func AddJob(spec string, job *Job) bool {
	lock.Lock()
	defer lock.Unlock()

	if GetEntryById(job.id) != nil {
		return false
	}
	err := mainCron.AddJob(spec, job)
	if err != nil {
		beego.Error("AddJob: ", err.Error())
		return false
	}
	return true
}

/*
*移除任务
 */
func RemoveJob(id int) {
	mainCron.RemoveJob(func(e *cron.Entry) bool {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				return true
			}
		}
		return false
	})
}

/*
*获取单个任务对象
 */
func GetEntryById(id int) *cron.Entry {
	entries := mainCron.Entries()
	for _, e := range entries {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				return e
			}
		}
	}
	return nil
}

/*
*获取多个任务对象
 */
func GetEntries(size int) []*cron.Entry {
	ret := mainCron.Entries()
	if len(ret) > size {
		return ret[:size]
	}
	return ret
}

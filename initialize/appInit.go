package initialize

import (
	"gin-blog/pkg/globals"
	"log"
)

func AppInit() {
	var err error
	globals.C, err = InitConfig()
	if err != nil {
		log.Println("配置文件加载出错")
		return
	}

	InitLogger()
	globals.E = InitEngine()
	globals.Db = InitDb(globals.C)
	globals.R = InitRedis(globals.C)
	globals.EventBus = InitEventBus()
	// 启动定时任务
	scheduler := InitTaskTimer()
	go scheduler.Run()
}

package initialize

import (
	"gin-blog/pkg/globals"
	"gin-blog/pkg/persistence"
	"time"
)

func InitTaskTimer() *persistence.TimeScheduler {

	articleLikeTask := persistence.NewArticleLikePersistTask(persistence.PersistParam{
		Db:  globals.Db,
		Rdb: globals.R,
	})

	articleReadingsTask := persistence.NewArticleReadPersistTask(persistence.PersistParam{
		Db:  globals.Db,
		Rdb: globals.R,
	})

	scheduler := persistence.NewScheduler([]persistence.Task{articleLikeTask, articleReadingsTask}, 5*time.Minute)
	return scheduler
}

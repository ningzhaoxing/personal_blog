package initialize

import (
	"fmt"
	"gin-blog/pkg/globals"
	"gin-blog/pkg/subscribes"
	"github.com/asaskevich/EventBus"
)

func InitEventBus() EventBus.Bus {
	bus := EventBus.New()

	err := bus.Subscribe(globals.ArticleReadingsTopic, subscribes.UpdateArticleReadingsSubscribe)

	if err != nil {
		panic(fmt.Errorf("initEventBus Error--> %v", err))
	}
	return bus
}

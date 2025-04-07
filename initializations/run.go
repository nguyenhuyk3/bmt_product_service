package initializations

import (
	"bmt_product_service/global"
	messagebroker "bmt_product_service/initializations/message_broker"
	"fmt"
)

func Run() {
	loadConfigs()
	initPostgreSql()

	messagebroker.InitFilmUploadConsummer()

	initRouter()

	r := initRouter()

	r.Run(fmt.Sprintf("0.0.0.0:%s", global.Config.Server.ServerPort))
}

package initializations

import (
	"bmt_product_service/global"
	"fmt"
)

func Run() {
	loadConfigs()
	initPostgreSql()
	initRouter()

	r := initRouter()

	r.Run(fmt.Sprintf("localhost:%s", global.Config.Server.ServerPort))
}

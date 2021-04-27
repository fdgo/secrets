package main

import (
	"fmt"
	_ "secret/initial"
	"secret/router"
	"secret/utils/config"
)

func main() {
	router.Router().Run(fmt.Sprintf(":%s", config.GloCfg.GetString("http.port")))
}

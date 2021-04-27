package initial

import "secret/utils/config"

func init() {
	config.GCfg(".")
	InitMysql()
	InitRedis()
}

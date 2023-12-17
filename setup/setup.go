package setup

import (
	"product-service/config"

	"github.com/sirupsen/logrus"
)

var Secrets = config.GetSecrets()

var Logger = logrus.New()

func init() {
	ConnectMongo()
	RedisConnection()
}

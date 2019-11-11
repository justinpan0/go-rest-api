package matching

import (
	logger "github.com/siddontang/go-log/log"
	"github.com/zimengpan/go-rest-api/service"
)

func StartEngine() {
	products, err := service.GetProducts()
	if err != nil {
		panic(err)
	}
	logger.Info("total products ", products)
	for _, product := range products {
		logger.Info("initialize kafka order reader for product ", product.Id)
		orderReader := NewKafkaOrderReader(product.Id, []string{"localhost:9092"})
		//snapshotStore := NewRedisSnapshotStore(product.Id)
		//logStore := NewKafkaLogStore(product.Id, gbeConfig.Kafka.Brokers)
		logger.Info("new engine ", product.Id)
		matchEngine := NewEngine(product, orderReader)
		logger.Info("match engine start ", product.Id)
		matchEngine.Start()
	}

	logger.Info("match engine ok")
}

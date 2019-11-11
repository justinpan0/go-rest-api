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
	logger.Info("total products", products)
	for _, product := range products {
		logger.Info("initialize kafka order reader for product", product.Id)
		orderReader := NewKafkaOrderReader(product.Id, []string{"localhost:9092"})
		//snapshotStore := NewRedisSnapshotStore(product.Id)
		//logStore := NewKafkaLogStore(product.Id, gbeConfig.Kafka.Brokers)
		matchEngine := NewEngine(product, orderReader)
		matchEngine.Start()
	}

	logger.Info("match engine ok")
}

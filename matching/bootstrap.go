package matching

import (
	logger "github.com/siddontang/go-log/log"
	"github.com/zimengpan/go-rest-api/main"
)

func StartEngine() {
	products, err := main.GetProducts()
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		orderReader := NewKafkaOrderReader(product.Id, gbeConfig.Kafka.Brokers)
		//snapshotStore := NewRedisSnapshotStore(product.Id)
		//logStore := NewKafkaLogStore(product.Id, gbeConfig.Kafka.Brokers)
		matchEngine := NewEngine(product, orderReader, logStore, snapshotStore)
		matchEngine.Start()
	}

	logger.Info("match engine ok")
}

package interactor

import "github.com/itouri/sgx-iaas/pkg/db/mongo"

var (
	mongoHandler *mongo.MongoHandler
)

func init() {
	/// そういやこれなにに使うんだ
	appc := mongo.AppConfig{
		Name:  "Endpoint",
		Port:  10000,
		Debug: true,
	}

	dbc := mongo.DbConfig{
		Host:     "localhost",
		Port:     27017,
		Database: "image-metadata",
	}

	mongoHandler = mongo.NewMongoHandler(appc, dbc)
}

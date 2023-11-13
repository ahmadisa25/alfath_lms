package mongo

import (
	"context"

	"flamingo.me/dingo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	//using code from novalagung.com
	injector.Bind(new(mongo.Database)).ToProvider(func() (*mongo.Database, error) {
		clientOptions := options.Client()
		clientOptions.ApplyURI("mongodb://localhost:27017")
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			return nil, err
		}

		err = client.Connect(context.Background())
		if err != nil {
			return nil, err
		}

		return client.Database("alfath_lms"), nil

	}).In(dingo.Singleton)
}

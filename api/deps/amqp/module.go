package amqp

import (
	"log"
	"os"

	"flamingo.me/dingo"
	"github.com/streadway/amqp"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(amqp.Connection)).ToProvider(func() *amqp.Connection {
		connString := fmt.gorm.DBSprintf("amqp://%s:%s@%s:%s/", os.Getenv("AMQP_USER"), os.Getenv("AMQP_GROUP"), os.Getenv("AMQP_HOST"), os.Getenv("AMQP_PORT"))
		conn, err := amqp.Dial(connString)

		if err != nil {
			log.Fatalf("%s: %s", "Failed to connect to Rabbit", err)
		}

		return conn

	}).In(dingo.Singleton)
}

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/myevents/configuration"
	"github.com/myevents/dblayer"
	_ "github.com/myevents/lib/msgqueue"
	_ "github.com/myevents/lib/msgqueue/amqp"
	"github.com/myevents/servicehandler"
	"github.com/streadway/amqp"
	 msgqueue_amqp "github.com/myevents/lib/msgqueue/amqp"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration file.")
	flag.Parse()

	config, _ := configuration.ExtractConfiguration(*confPath)
	conn, err := amqp.Dial(config.AMQPMessageBroker)

	if err!=nil{
		panic(err)
	}

	emitter,err:=msgqueue_amqp.NewAMQPEmitter(conn)
	if err!=nil{
		panic(err)
	}
	


	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.DatabaseType, config.DBConnection)

	// log.Fatal(servicehandler.ServeAPI(configuration.RestfulEPDefault, dbhandler))
	httpErrChan, httpsErrChan := servicehandler.ServeAPI(configuration.RestfulEPDefault, configuration.RestfulTLSEPDefault, dbhandler)
	done := make(chan bool)
	select {
	case err := <-httpErrChan:
		log.Fatal("Http error: ", err)
		done <- true
	case err := <-httpsErrChan:
		log.Fatal("Https error: ", err)
		done <- true
	}
	<-done
}

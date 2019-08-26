package main

import (
	dao "go-nsq/nsq/dao"
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	// eMumba => topicName, ch => channelName
	q, err := nsq.NewConsumer("eMumba", "ch", config)
	dao.FindError(err, "consume Failed")

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(message.Body)
		return nil
	}))

	er := q.ConnectToNSQD("127.0.0.1:4150")
	dao.FindError(er, "Connection to nsq failed")

	wg.Wait()

}

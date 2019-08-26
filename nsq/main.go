package main

import (
	dao "go-nsq/nsq/dao"
	"strconv"

	"github.com/nsqio/go-nsq"
)

func main() {
	//Connection
	config := nsq.NewConfig()
	w, err := nsq.NewProducer("localhost:4150", config)
	dao.FindError(err, "Producer failed")

	//publish to nsq(topicName => eMumba)
	for i := 0; i < 10; i++ {
		err := w.Publish("eMumba", []byte("test"+strconv.Itoa(i)))
		dao.FindError(err, "Publish failed")
	}
	w.Stop()
}

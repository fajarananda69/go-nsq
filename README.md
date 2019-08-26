# go-nsq

## Install and SETUP GO
Download and configure your workspace with latest version of Go and correct environment path.
- [Last Version](https://golang.org/dl/)
- [Windows](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/)
- [Linux](http://www.tecmint.com/install-go-in-linux/)

## install and SETUP NSQ
- install [nsq](https://nsq.io/deployment/installing.html) and setup nsq admin [klik here](https://nsq.io/overview/quick_start.html)

## Install Driver
Install dep [klik here](https://golang.github.io/dep/docs/installation.html) and Create new dep
```
dep init 
```
```
dep ensure -add github.com/nsqio/go-nsq
```

## Import
```
import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)
```

## Producer
```
    config := nsq.NewConfig()
	w, err := nsq.NewProducer("localhost:4150", config)
	dao.FindError(err, "Producer failed")

	for i := 0; i < 10; i++ {
		err := w.Publish("eMumba", []byte("test"+strconv.Itoa(i)))
		dao.FindError(err, "Publish failed")
	}
	w.Stop()
```

## Consumer
```
    wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, err := nsq.NewConsumer("eMumba", "ch", config)
	dao.FindError(err, "consume Failed")

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(message.Body)
		return nil
	}))

	er := q.ConnectToNSQD("127.0.0.1:4150")
	dao.FindError(er, "Connection to nsq failed")

	wg.Wait()
```

## Run
produce data to nsq
```
go run main.go
```

consume data from nsq
```
go run consumer.go
```
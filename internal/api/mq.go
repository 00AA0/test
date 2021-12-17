package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

type Msg struct {
	Id   int64  `json:"id"`
	Info string `json:"info"`
}

func StartConsumer() {
	MQConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test_consumer"),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithNameServer(primitive.NamesrvAddr{"127.0.0.1:9876"}),
		consumer.WithStrategy(consumer.AllocateByAveragely))

	if err != nil {
		fmt.Println(err)
	}

	f := func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, m := range msgs {
			if ctx.Err() != nil {
				return consumer.SuspendCurrentQueueAMoment, ctx.Err()
			}
			msg := &Msg{}
			json.Unmarshal(m.Body, msg)

			fmt.Println(msg)
			//if err := c.call(callback, m); err != nil {
			//	return consumer.SuspendCurrentQueueAMoment, nil
			//}
		}
		return 0, nil
	}

	err = MQConsumer.Subscribe("test_topic", consumer.MessageSelector{}, f)
	//err := api.MQConsumer.Subscribe("test_topic", selector, f)
	if err != nil {
		fmt.Println(err)
	}

	err = MQConsumer.Start()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Hour)
	err = MQConsumer.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}

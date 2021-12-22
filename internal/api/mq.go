package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Msg struct {
	Id   int64  `json:"id"`
	Info string `json:"info"`
}

var c *gin.Context

func StartProducer() (err error) {
	MQProducer, err := rocketmq.NewProducer(
		producer.WithGroupName("test_producer"),
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(0),
		producer.WithQueueSelector(producer.NewRoundRobinQueueSelector()),
		producer.WithSendMsgTimeout(time.Second*60),
		producer.WithInstanceName("test"))
	if err != nil {
		fmt.Println(err)
	}
	err = MQProducer.Start()
	if err != nil {
		fmt.Println(err)
	}

	msgList := make([]*primitive.Message, 0)
	for i := 1; i < 10; i++ {
		m := Msg{
			Id:   int64(i),
			Info: "test msg" + strconv.Itoa(i),
		}
		b, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err)
		}

		msg := primitive.NewMessage("test_topic", b)
		msg.WithTag("tag")
		//producer.NewHashQueueSelector()
		//msg.WithShardingKey("aa")
		msg.WithKeys([]string{"iii", "123"})
		//msg.WithProperty("name","micro services")
		msgList = append(msgList, msg)
	}

	m := Msg{
		Id:   int64(99),
		Info: "test msg" + strconv.Itoa(99),
	}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	msg := primitive.NewMessage("test_topic_x", b)
	msgList = append(msgList, msg)

	for _, message := range msgList {
		time.Sleep(time.Second)
		result, err := MQProducer.SendSync(c, message)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)

	}

	fmt.Println()

	time.Sleep(time.Hour)
	//err = MQProducer.Shutdown()
	//if err != nil {
	//	fmt.Printf("shutdown Consumer error: %s", err.Error())
	//}
	return
}

func encodeBatch(msgs ...*primitive.Message) *primitive.Message {
	if len(msgs) == 1 {
		return msgs[0]
	}

	// encode batch
	batch := new(primitive.Message)
	batch.Topic = msgs[0].Topic
	batch.Queue = msgs[0].Queue
	if len(msgs) > 1 {
		batch.Body = MarshalMessageBatch(msgs...)
		batch.Batch = true
	} else {
		batch.Body = msgs[0].Body
		batch.Flag = msgs[0].Flag
		batch.WithProperties(msgs[0].GetProperties())
		batch.TransactionId = msgs[0].TransactionId
	}
	return batch
}

func MarshalMessageBatch(msgs ...*primitive.Message) []byte {
	buffer := bytes.NewBufferString("")
	for _, msg := range msgs {
		data := msg.Marshal()
		buffer.Write(data)
	}
	return buffer.Bytes()
}

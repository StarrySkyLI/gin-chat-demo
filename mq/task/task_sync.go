package task

import (
	"chat/consts"
	"chat/mq"
	"chat/service"
	"context"
	"encoding/json"
	"fmt"
)

type SyncTask struct {
}

func (s *SyncTask) RunTaskCreate(ctx context.Context) error {
	rabbitMqQueue := consts.RabbitMqTaskQueue
	msgs, err := mq.ConsumeMessage(ctx, rabbitMqQueue)

	if err != nil {
		return err
	}
	var forever chan struct{}
	m := &consts.Msg{}

	go func() {
		for d := range msgs {

			// 落库

			_ = json.Unmarshal(d.Body, m)
			fmt.Println(m)
			err = service.InsertMsg(m.Datebase, m.Id, m.Content, m.Read, m.Expire)
			if err != nil {
				fmt.Println("InsertOneMsg Err", err)
			}

			d.Ack(false)

		}
	}()

	<-forever

	return nil
}

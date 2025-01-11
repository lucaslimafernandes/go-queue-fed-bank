package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lucaslimafernandes/go-queue-fed-bank/api/model"
)

func GetUserByID(userID int) (map[string]string, error) {

	ctx := context.Background()

	key := fmt.Sprintf("user_id:%d", userID)

	result, err := RDB.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	return result, err

}

func InsertUserQueue(userID int) error {

	var user model.QueueEntry

	ctx := context.Background()

	nrQueue, _ := IncrQueueNumber()

	user = model.QueueEntry{
		UserId:  userID,
		NrQueue: nrQueue,
		Date:    time.Now(),
		Status:  "waiting",
	}

	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	_, err = RDB.LPush(ctx, "userQueue", userData).Result()
	if err != nil {
		return err
	}

	return nil

}

func IncrQueueNumber() (int64, error) {

	ctx := context.Background()

	newQueueNr, err := RDB.Incr(ctx, "NR_FILA").Result()
	if err != nil {
		return 0, err
	}

	return newQueueNr, err

}

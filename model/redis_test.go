package model_test

import (
	"fmt"
	"testing"

	"github.com/pangxieke/sendmail/model"
	_ "github.com/pangxieke/sendmail/test"
	"github.com/stretchr/testify/assert"
)

func TestSyncRedis(t *testing.T) {
	assert := assert.New(t)
	key := "mail_notify"
	val := "{\"receivers\": \"pangxieke@126.com, pangxieke@126.com\", \"subject\": \"this is subject\", \"body\": \"this is body\", \"subtype\": \"html\"}"
	res, err := model.RedisClient.LPush(key, val).Result()

	assert.Nil(err)
	fmt.Println(res)
}

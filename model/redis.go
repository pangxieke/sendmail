package model

import (
	"encoding/json"
	"time"

	"github.com/pangxieke/sendmail/log"
)

func SyncRedis() (err error) {
	for {
		time.Sleep(2000 * time.Millisecond)

		// get data from redis list
		lLen, err := RedisClient.LLen("mail_notify").Result()

		if lLen == 0 {
			continue
		}
		str, err := RedisClient.LPop("mail_notify").Result()
		if err != nil {
			log.Info("redis LPop err ", err)
		}
		var mailData MailNotify
		if err = json.Unmarshal([]byte(str), &mailData); err != nil {
			log.Info("json.Unmarshal err", str)
		}

		//send mail
		go SendMail(mailData)
	}
}

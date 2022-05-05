package redis

import "time"

func Set(keyUserId string, value interface{}, duration time.Duration) (err error) {
	err = client.Set(keyUserId, value, duration).Err()
	return
}

func Get(keyUserId string) (value string, err error) {
	value, err = client.Get(keyUserId).Result()
	return
}

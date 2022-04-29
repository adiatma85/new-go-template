package redis

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/adiatma85/golang-alter-url-shortener/base62"
	"github.com/adiatma85/golang-alter-url-shortener/storage"
	"github.com/gomodule/redigo/redis"
)

type RedisClient struct {
	Pool *redis.Pool
}

// Return the instance of RedisClient
func New(host, port, password string) (storage.Service, error) {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		},
	}

	return &RedisClient{pool}, nil
}

// Function to save item in database
func (r *RedisClient) Save(url string, expires time.Time) (string, error) {
	conn := r.Pool.Get()
	defer conn.Close()

	var id uint64

	for used := true; used; used = r.isUsed(id) {
		id = rand.Uint64()
	}

	shortLink := storage.Item{
		Id:      id,
		URL:     url,
		Expires: expires.Format("2006-01-02 15:04:05.728046 +0300 EEST"),
		Visits:  0,
	}

	_, err := conn.Do("HMSET", redis.Args{"Shortener:" + strconv.FormatUint(id, 10)}.AddFlat(shortLink)...)

	if err != nil {
		return "", err
	}

	_, err = conn.Do("EXPIREAT", "Shortener:"+strconv.FormatUint(id, 10), expires.Unix())

	if err != nil {
		return "", err
	}

	return base62.Encode(id), nil
}

// Func to load
func (r *RedisClient) Load(code string) (string, error) {
	conn := r.Pool.Get()
	defer conn.Close()

	decodedId, err := base62.Decode(code)

	if err != nil {
		return "", err
	}

	urlString, err := redis.String(conn.Do("HGET", "Shortener:"+strconv.FormatUint(decodedId, 10), "url"))

	if err != nil {
		return "", err
	} else if len(urlString) == 0 {
		return "", storage.ErrNoLink
	}

	_, err = conn.Do("HINCRBY", "Shortener:"+strconv.FormatUint(decodedId, 10), "visits", 1)

	if err != nil {
		return "", err
	}

	return urlString, nil
}

// Func to load info of shortener link
func (r *RedisClient) LoadInfo(code string) (*storage.Item, error) {
	conn := r.Pool.Get()
	defer conn.Close()

	decodedId, err := base62.Decode(code)
	if err != nil {
		return nil, err
	}
	values, err := redis.Values(conn.Do("HGETALL", "Shortener:"+strconv.FormatUint(decodedId, 10)))

	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, storage.ErrNoLink
	}

	var shortLink storage.Item
	err = redis.ScanStruct(values, &shortLink)

	if err != nil {
		return nil, err
	}

	return &shortLink, nil
}

// Colse the pool of redis
func (r *RedisClient) Close() error {
	return r.Pool.Close()
}

// Helper function to check whether the id is used or not
func (r *RedisClient) isUsed(id uint64) bool {
	conn := r.Pool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", "Shortener:"+strconv.FormatUint(id, 10)))
	if err != nil {
		return false
	}
	return exists
}

// Helper function to check whether the id is exist or not
// func (r *RedisClient) isAvailable(id uint64) bool {
// 	conn := r.Pool.Get()
// 	defer conn.Close()

// 	exists, err := redis.Bool(conn.Do("EXISTS", "Shortener:"+strconv.FormatUint(id, 10)))
// 	if err != nil {
// 		return false
// 	}
// 	return !exists
// }

package redis_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)

	err := client.Close()
	assert.Nil(t, err)
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()

	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
}

func TestString(t *testing.T) {
	client.Set(ctx, "name", "boshir", time.Second*3)

	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "boshir", result)

	time.Sleep(time.Second * 5)
	result, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "boshir")
	client.RPush(ctx, "names", "yusuf")
	client.RPush(ctx, "names", "ahmad")

	assert.Equal(t, "boshir", client.LPop(ctx, "names").Val())
	assert.Equal(t, "yusuf", client.LPop(ctx, "names").Val())
	assert.Equal(t, "ahmad", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")
}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "boshir")
	client.SAdd(ctx, "students", "boshir")
	client.SAdd(ctx, "students", "yusuf")
	client.SAdd(ctx, "students", "yusuf")
	client.SAdd(ctx, "students", "ahmad")
	client.SAdd(ctx, "students", "ahmad")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"boshir", "yusuf", "ahmad"}, client.SMembers(ctx, "students").Val())
}

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "boshir"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 97, Member: "sukur"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 88, Member: "sadi"})

	assert.Equal(t, []string{"sadi", "sukur", "boshir"}, client.ZRange(ctx, "scores", 0, 2).Val())
	assert.Equal(t, "boshir", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "sukur", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "sadi", client.ZPopMax(ctx, "scores").Val()[0].Member)
}

func TestHash(t *testing.T) {
	client.HSet(ctx, "user:1", "id", "1")
	client.HSet(ctx, "user:1", "name", "boshir")
	client.HSet(ctx, "user:1", "email", "test@example.com")

	user := client.HGetAll(ctx, "user:1").Val()
	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "boshir", user["name"])
	assert.Equal(t, "test@example.com", user["email"])

	client.Del(ctx, "user:1")
}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "toko A",
		Longitude: 106.822702,
		Latitude:  -6.177590,
	})

	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "toko B",
		Longitude: 106.820889,
		Latitude:  -6.174964,
	})
	distance := client.GeoDist(ctx, "sellers", "toko A", "toko B", "km").Val()
	assert.Equal(t, 0.3543, distance)

	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude:  106.821825,
		Latitude:   -6.175105,
		Radius:     5,
		RadiusUnit: "km",
	}).Val()
	assert.Equal(t, []string{"toko A", "toko B"}, sellers)
}

func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "visitors", "boshir", "joko", "ahmad")
	client.PFAdd(ctx, "visitors", "boshir", "budi", "eko")
	client.PFAdd(ctx, "visitors", "budi", "joko", "suharto")
	assert.Equal(t, int64(6), client.PFCount(ctx, "visitors").Val())
}

func TestPipeliner(t *testing.T) {
	client.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "boshir", time.Second*5)
		pipeliner.SetEx(ctx, "address", "MY", time.Second*5)
		return nil
	})

	assert.Equal(t, "boshir", client.Get(ctx, "name").Val())
	assert.Equal(t, "MY", client.Get(ctx, "address").Val())
}

func TestTransaction(t *testing.T) {
	client.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "boshir", time.Second*5)
		pipeliner.SetEx(ctx, "address", "MY", time.Second*5)
		return nil

	})

	assert.Equal(t, "boshir", client.Get(ctx, "name").Val())
	assert.Equal(t, "MY", client.Get(ctx, "address").Val())
}

func TestPublishStream(t *testing.T) {
	for range 10 {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]any{
				"name":    "boshir",
				"address": "Malaysia",
			},
		}).Err()
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	client.XGroupCreate(ctx, "members", "group-1", "0")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
	client.XGroupCreateConsumer(ctx, "members", "group-2", "consumer-2")
}

func TestGetStream(t *testing.T) {
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    time.Second * 5,
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message.ID)
			fmt.Println(message.Values)
		}
	}

}

func TestSubscribePubSub(t *testing.T) {
	pubSub := client.Subscribe(ctx, "channel-1")
	for range 10 {
		message, err := pubSub.ReceiveMessage(ctx)
		assert.Nil(t, err)
		fmt.Println(message.Payload)
	}
	pubSub.Close()
}

func TestPublishPubSub(t *testing.T) {
	for i := range 10 {
		err := client.Publish(ctx, "channel-1", "Hello "+strconv.Itoa(i)).Err()
		assert.Nil(t, err)
	}
}

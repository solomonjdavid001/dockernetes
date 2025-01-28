package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/solomonjdavid001/Dockernetes/backend/config"
)

var (
	redisClient *redis.Client
	once        sync.Once
)

func getRedisClient() *redis.Client {
	once.Do(func() {
		redisClient = initRedisClient()
	})
	return redisClient
}

func initRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.GlobalConfig.Redis.Server,
	})
	return redisClient
}

func SetStringValue(key string, value string) error {
	client := getRedisClient()
	ctx := context.Background()

	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatal("Error setting Redis value: ", err)
		return err
	}
	return nil
}

func SetListValue(body SetRedisListRequest) error {
	client := getRedisClient()
	ctx := context.Background()
	var err error

	if body.PushToTail {
		err = client.LPush(ctx, body.Key, body.Element).Err()
	} else {
		err = client.RPush(ctx, body.Key, body.Element).Err()
	}
	if body.TTL > 0 {
		err = SetExpiry(client, ctx, body.Key, body.TTL)
	}
	if err != nil {
		fmt.Println("Error setting Redis value or ttl: ", err)
		return err
	}
	return nil
}

func SetHashValue(body SetRedisHashRequest) error {
	client := getRedisClient()
	ctx := context.Background()
	err := client.HSet(ctx, body.Key, body.Hash).Err()
	if body.TTL > 0 {
		err = SetExpiry(client, ctx, body.Key, body.TTL)
	}
	if err != nil {
		fmt.Println("Error setting Redis value or ttl: ", err)
		return err
	}
	return nil
}

func SetAddSetValue(body SetRedisSetRequest) error {
	client := getRedisClient()
	ctx := context.Background()
	err := client.SAdd(ctx, body.Key, body.Member).Err()
	if body.TTL > 0 {
		err = SetExpiry(client, ctx, body.Key, body.TTL)
	}
	if err != nil {
		fmt.Println("Error setting Redis value or ttl: ", err)
		return err
	}
	return nil
}

func SetAddSortedSetValue(body SetRedisSortedSetRequest) error {
	client := getRedisClient()
	ctx := context.Background()

	var redisZSet []redis.Z
	for member, score := range body.MemberScore {
		redisZSet = append(redisZSet, redis.Z{
			Score:  float64(score),
			Member: member,
		})
	}

	err := client.ZAdd(ctx, body.Key, redisZSet...).Err()
	if body.TTL > 0 {
		err = SetExpiry(client, ctx, body.Key, body.TTL)
	}
	if err != nil {
		fmt.Println("Error setting Redis value or ttl: ", err)
		return err
	}
	return nil
}

func UpdateRedisList(body UpdateRedisListRequest) error {
	client := getRedisClient()
	ctx := context.Background()

	err := client.LSet(ctx, body.Key, body.Index, body.Element).Err()
	if err != nil {
		fmt.Println("Error updating Redis list value: ", err)
		return err
	}
	return nil
}

func DeleteRedisSet(key string, member interface{}) error {
	client := getRedisClient()
	ctx := context.Background()

	err := client.SRem(ctx, key, member).Err()
	if err != nil {
		fmt.Println("Error deleting Redis set value: ", err)
		return err
	}
	return nil
}

func DeleteRedisSortedSet(key string, member interface{}) error {
	client := getRedisClient()
	ctx := context.Background()

	err := client.ZRem(ctx, key, member).Err()
	if err != nil {
		fmt.Println("Error deleting Redis sorted set value: ", err)
		return err
	}
	return nil
}

func DeleteRedisHash(key string, field string) error {
	client := getRedisClient()
	ctx := context.Background()

	err := client.HDel(ctx, key, field).Err()
	if err != nil {
		fmt.Println("Error deleting Redis hash field value: ", err)
		return err
	}
	return nil
}

func GetValue(key string) (string, error) {
	client := getRedisClient()
	ctx := context.Background()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("Error getting Redis value: ", err)
		return "", err
	}
	return val, nil
}

func GetAll() ([]byte, error) {
	client := getRedisClient()
	ctx := context.Background()

	iter := client.Scan(ctx, 0, "*", 0).Iterator()
	var keyInfos []KeyInfo
	for iter.Next(ctx) {
		var value string
		var listValue []string
		var hashValue map[string]string
		var setValue []string
		var sortedSetValue []redis.Z
		key := iter.Val()

		ttl, _ := client.TTL(ctx, iter.Val()).Result()

		size, _ := client.MemoryUsage(ctx, iter.Val()).Result()

		keyType, _ := client.Type(ctx, iter.Val()).Result()

		switch keyType {
		case "string":
			value, _ = client.Get(ctx, iter.Val()).Result()
		case "list":
			listValue, _ = client.LRange(ctx, key, 0, -1).Result()
		case "hash":
			hashValue, _ = client.HGetAll(ctx, key).Result()
		case "set":
			setValue, _ = client.SMembers(ctx, key).Result()
		case "zset":
			sortedSetValue, _ = client.ZRangeWithScores(ctx, key, 0, -1).Result()
		}

		keyInfos = append(keyInfos, KeyInfo{
			Key:            iter.Val(),
			Value:          value,
			TTL:            ttl.String(),
			Type:           keyType,
			Size:           size,
			ListValue:      listValue,
			SetValue:       setValue,
			SortedSetValue: sortedSetValue,
			HashValue:      hashValue,
		})
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(keyInfos)
	if err != nil {
		fmt.Printf("failed to marshal to JSON: %v", err)
		return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return jsonData, nil
}

func DeleteKey(key string) error {
	client := getRedisClient()
	ctx := context.Background()

	err := client.Del(ctx, key).Err()
	if err != nil {
		fmt.Printf("Error deleting Redis key: %v", err)
		return err
	}
	return nil
}

func SetKeyExpiry(body SetRedisExpiryRequest) error {
	client := getRedisClient()
	ctx := context.Background()

	err := SetExpiry(client, ctx, body.Key, body.TTL)
	if err != nil {
		fmt.Printf("Error setting Redis key expiry: %v", err)
		return err
	}
	return nil
}

func SetExpiry(client *redis.Client, ctx context.Context, key string, ttl uint64) error {
	err := client.Expire(ctx, key, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		fmt.Println("Error setting TTL:", err)
		return err
	}
	return nil
}

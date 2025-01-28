package redis

import "github.com/redis/go-redis/v9"

type SetRedisStringValueRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	TTL   uint64 `json:"ttl"`
}

type SetRedisListRequest struct {
	Key        string        `json:"key"`
	TTL        uint64        `json:"ttl"`
	PushToTail bool          `json:"pushToTail"`
	Element    []interface{} `json:"element"`
}

type SetRedisHashRequest struct {
	Key  string                 `json:"key"`
	TTL  uint64                 `json:"ttl"`
	Hash map[string]interface{} `json:"hash"`
}

type SetRedisSetRequest struct {
	Key    string        `json:"key"`
	TTL    uint64        `json:"ttl"`
	Member []interface{} `json:"member"`
}

type SetRedisSortedSetRequest struct {
	Key         string           `json:"key"`
	TTL         uint64           `json:"ttl"`
	MemberScore map[string]int64 `json:"memberScore"`
}

type UpdateRedisListRequest struct {
	Key     string      `json:"key"`
	Index   int64       `json:"index"`
	Element interface{} `json:"element"`
}

type SetRedisExpiryRequest struct {
	Key string `json:"key"`
	TTL uint64 `json:"ttl"`
}

type KeyInfo struct {
	Key            string            `json:"key"`
	Value          string            `json:"value,omitempty"`
	TTL            string            `json:"ttl"`
	Type           string            `json:"type"`
	Size           int64             `json:"size"`
	ListValue      []string          `json:"listValue,omitempty"`
	SetValue       []string          `json:"setValue,omitempty"`
	HashValue      map[string]string `json:"hashValue,omitempty"`
	SortedSetValue []redis.Z         `json:"sortedSetValue,omitempty"`
}

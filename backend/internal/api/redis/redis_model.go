package redis

type SetRedisValueRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

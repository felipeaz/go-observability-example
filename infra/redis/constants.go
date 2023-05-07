package redis

// errors
const (
	_failedToConnectToRedisServer = "failed to connect to redis server: %v\n"
	_failedToSetKey               = "failed to set key %s, value %s: %v\n"
	_failedToGetKey               = "failed to get key %s: %v\n"
	_failedToRemoveKey            = "failed to remove key %s: %v\n"
)

// actions
const (
	_getAction    = "GET"
	_setAction    = "SET"
	_deleteAction = "DEL"
)

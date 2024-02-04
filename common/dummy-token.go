package common

func DummyToken(key string) (value any, exists bool) {
	value, exists = key, true
	return
}

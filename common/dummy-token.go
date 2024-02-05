package common

func DummyToken(key string) (value string, exists bool) {
	value, exists = key, true
	return
}

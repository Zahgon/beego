package session

import (
	"crypto/cipher"
	"encoding/gob"
)

func init() {
	gob.Register([]interface{}{})
	gob.Register(map[int]interface{}{})
	gob.Register(map[string]interface{}{})
	gob.Register(map[interface{}]interface{}{})
	gob.Register(map[string]string{})
	gob.Register(map[int]string{})
	gob.Register(map[int]int{})
	gob.Register(map[int]int64{})
}

func EncodeGob(obj map[interface{}]interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DecodeGob(encoded []byte) (map[interface{}]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateRandomKey(strength int) []byte { _ = "STUB: not implemented"; return nil }

func encrypt(block cipher.Block, value []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decrypt(block cipher.Block, value []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeCookie(block cipher.Block, hashKey, name string, value map[interface{}]interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func decodeCookie(block cipher.Block, hashKey, name, value string, gcmaxlifetime int64) (map[interface{}]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encode(value []byte) []byte { _ = "STUB: not implemented"; return nil }

func decode(value []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

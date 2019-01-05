package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/gob"
)

func HashCode(data interface{}) (string, error) {
	hash := sha1.New()

	gobEncoder := gob.NewEncoder(hash)
	if err := gobEncoder.Encode(data); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))[:8], nil
}

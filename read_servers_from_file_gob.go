package main

import (
	"encoding/gob"
	"os"

	"github.com/google/brotli/go/cbrotli"
)

func ReadServersFromFileGob(file string) ([]Server, error) {
	gobBrReader, err := os.Open("./data/" + file + ".gob.br")
	if err != nil {
		return nil, err
	}
	defer gobBrReader.Close()
	gobReader := cbrotli.NewReader(gobBrReader)
	defer gobReader.Close()
	gobDecoder := gob.NewDecoder(gobReader)
	var servers []Server
	if err := gobDecoder.Decode(&servers); err != nil {
		return nil, err
	}
	return ParseServers(servers), nil
}

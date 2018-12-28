package main

import (
	"encoding/gob"
	"os"

	"github.com/google/brotli/go/cbrotli"
)

func WriteServersToFileGob(file string, servers []Server) (err error) {
	defer func() {
		if err == nil {
			_ = os.Rename("./data/"+file+".gob.br", "./data/"+file+".bak.gob.br")
			_ = os.Rename("./data/"+file+".gob.br.tmp", "./data/"+file+".gob.br")
		}
	}()
	gobBrFile, err := os.OpenFile("./data/"+file+".gob.br.tmp", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer gobBrFile.Close()
	gobWriter := cbrotli.NewWriter(gobBrFile, cbrotli.WriterOptions{Quality: 9})
	defer gobWriter.Close()
	gobEncoder := gob.NewEncoder(gobWriter)
	return gobEncoder.Encode(servers)
}

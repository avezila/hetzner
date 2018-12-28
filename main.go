package main

import (
	"log"
)

func main() {
	servers, err := ReadServersFromFileGob("servers")
	log.Println(len(servers), err)
	if err != nil || len(servers) < 100 {
		servers, err = ReadServersFromFileGob("servers.bak")
	}
	if err != nil || len(servers) < 100 {
		return
	}
	servers = ServersDedup(servers)
	log.Printf("%+v\n", servers[len(servers)-1])
	err = WriteServersToFileGob("servers", servers)
	if err != nil {
		log.Println(err)
		return
	}

	refreshNewServersDaemon()
}

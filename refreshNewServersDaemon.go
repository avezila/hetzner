package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

var uri = "https://www.hetzner.com/a_hz_serverboerse/live_data.json?m=1539681284257"

func refreshNewServersDaemon() {
	for {
		if err := refresh(); err != nil {
			log.Println(err)
			time.Sleep(time.Minute)
			continue
		}
		time.Sleep(time.Second * 5)
	}
}

var lastModified string

func refresh() error {
	log.Println("refresh", lastModified)
	client := http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept-Encoding", "gzip")
	if lastModified != "" {
		req.Header.Add("if-modified-since", lastModified)
	}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	log.Println("status", response.Status, response.Header.Get("Content-Encoding"))
	if response.StatusCode == http.StatusNotModified {
		return nil
	}
	lastModified = response.Header.Get("last-modified")

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return err
		}
		defer reader.Close()
	default:
		reader = response.Body
	}

	jsonDecoder := json.NewDecoder(reader)
	var hServers HServers
	if err := jsonDecoder.Decode(&hServers); err != nil {
		return err
	}

	if err := ReceiveNewServers(hServers.Servers); err != nil {
		return nil
	}

	return nil
}

func ReceiveNewServers(newServers []Server) error {
	newServers = ParseServers(newServers)
	log.Println("new Servers", len(newServers))
	servers, err := ReadServersFromFileGob("servers")
	if err != nil || len(servers) < 100 {
		servers, err = ReadServersFromFileGob("servers.bak")
	}
	if err != nil {
		return err
	}

	log.Println("old Servers", len(servers))
	servers = append(servers, newServers...)
	servers = ServersDedup(servers)
	log.Println("dedup servers", len(servers))

	if err := GenerateHtml(servers, newServers); err != nil {
		log.Println("Failed generate html", err)
	}
	return WriteServersToFileGob("servers", servers)
}

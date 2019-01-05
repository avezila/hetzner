package main

import (
	"html/template"
	"log"
	"os"
	"sort"
)

var ServersTopTemplate *template.Template

const WWWServersTopPath = WWWPath + "s/index.html"
const WWWServersNowPath = WWWPath + "s/now.html"

func init() {
	var err error
	ServersTopTemplate, err = template.ParseFiles("./servers.top.go.html")
	if err != nil {
		panic(err)
	}
}

func GenerateHtml(servers []Server, nowServers []Server) error {
	var eservers []EServer
	for _, server := range servers {
		eserver, err := server.EServer()
		if err != nil {
			log.Println("Failed server.EServer", err)
			continue
		}
		eservers = append(eservers, eserver)
	}
	aggregated := Aggregate(eservers)
	var aservers []AServer
	for _, eserver := range eservers {
		aservers = append(aservers, eserver.AServer(aggregated).Round())
	}

	for _, aserver := range aservers {
		if err := aserver.WriteServerHTML(); err != nil {
			log.Println("Failed write server html", err)
		}
	}
	if err := WriteSummary(aservers, WWWServersTopPath, aggregated); err != nil {
		return err
	}

	eservers = nil
	for _, server := range nowServers {
		eserver, err := server.EServer()
		if err != nil {
			log.Println("Failed server.EServer", err)
			continue
		}
		eservers = append(eservers, eserver)
	}
	aservers = nil
	for _, eserver := range eservers {
		aservers = append(aservers, eserver.AServer(aggregated).Round())
	}
	if err := SendTelegram(aservers); err != nil {
		log.Println("Failed SendTelegram", err)
	}
	return WriteSummary(aservers, WWWServersNowPath, aggregated)
}

func WriteSummary(aservers []AServer, path string, aggregated Aggregated) error {
	topCPU := append(aservers[:0:0], aservers...)
	sort.Slice(topCPU, func(i, j int) bool {
		return topCPU[i].CPUBenchPerEurAdj > topCPU[j].CPUBenchPerEurAdj
	})
	topRAM := append(aservers[:0:0], aservers...)
	sort.Slice(topRAM, func(i, j int) bool {
		return topRAM[i].RAMGbPerEurAdj > topRAM[j].RAMGbPerEurAdj
	})
	topHDD := append(aservers[:0:0], aservers...)
	sort.Slice(topHDD, func(i, j int) bool {
		return topHDD[i].HddGbPerEurAdj > topHDD[j].HddGbPerEurAdj
	})
	topSSD := append(aservers[:0:0], aservers...)
	sort.Slice(topSSD, func(i, j int) bool {
		return topSSD[i].SsdGbPerEurAdj > topSSD[j].SsdGbPerEurAdj
	})
	topAvg := append(aservers[:0:0], aservers...)
	sort.Slice(topAvg, func(i, j int) bool {
		return topAvg[i].Average > topAvg[j].Average
	})

	if err := os.MkdirAll(WWWServerPath, 0700); err != nil {
		return err
	}
	serversTopHTMLFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer serversTopHTMLFile.Close()
	return ServersTopTemplate.Execute(serversTopHTMLFile, struct {
		Aggregated Aggregated
		TopCPU     []AServer
		TopRAM     []AServer
		TopHDD     []AServer
		TopSSD     []AServer
		TopAvg     []AServer
	}{
		Aggregated: aggregated,
		TopCPU:     topCPU[0:20],
		TopRAM:     topRAM[0:20],
		TopHDD:     topHDD[0:20],
		TopSSD:     topSSD[0:20],
		TopAvg:     topAvg[0:20],
	})
}

package main

import (
	"log"
	"strconv"
)

type HServers struct {
	Hash         string `json:"hash"`
	MinMaxValues struct {
		MaxBenchmark int64   `json:"maxBenchmark"`
		MaxHDDCount  int64   `json:"maxHDDCount"`
		MaxHDDSize   int64   `json:"maxHDDSize"`
		MaxPrice     float64 `json:"maxPrice"`
		MaxRAM       int64   `json:"maxRam"`
		MinBenchmark int64   `json:"minBenchmark"`
		MinHDDCount  int64   `json:"minHDDCount"`
		MinHDDSize   int64   `json:"minHDDSize"`
		MinPrice     float64 `json:"minPrice"`
		MinRAM       int64   `json:"minRam"`
	} `json:"minMaxValues"`
	Servers []Server `json:"server"`
}

type Server struct {
	Bandwith      int64    `json:"bandwith"`
	CPU           string   `json:"cpu"`
	CPUBenchmark  int64    `json:"cpu_benchmark"`
	CPUCount      int64    `json:"cpu_count"`
	Datacenter    []string `json:"datacenter"`
	Description   []string `json:"description"`
	Dist          []string `json:"dist"`
	FixedPrice    bool     `json:"fixed_price"`
	Freetext      string   `json:"freetext"`
	HddCount      int64    `json:"hdd_count"`
	HddHr         string   `json:"hdd_hr"`
	HddSize       int64    `json:"hdd_size"`
	IsEcc         bool     `json:"is_ecc"`
	IsHighio      bool     `json:"is_highio"`
	Key           int64    `json:"key"`
	Name          string   `json:"name"`
	NextReduce    int64    `json:"next_reduce"`
	NextReduceHr  string   `json:"next_reduce_hr"`
	Price         string   `json:"price"`
	PriceEur      float64  `json:"price_eur"`
	PriceV        string   `json:"price_v"`
	PriceVEur     float64  `json:"price_v_eur"`
	RAM           int64    `json:"ram"`
	RAMHr         string   `json:"ram_hr"`
	SetupPrice    string   `json:"setup_price"`
	SetupPriceEur float64  `json:"setup_price_eur"`
	SpecialHdd    string   `json:"specialHdd"`
	Specials      []string `json:"specials"`
	Traffic       string   `json:"traffic"`
}

func (s Server) Parse() Server {
	s.PriceEur, _ = strconv.ParseFloat(s.Price, 64)
	s.PriceVEur, _ = strconv.ParseFloat(s.PriceV, 64)
	s.SetupPriceEur, _ = strconv.ParseFloat(s.SetupPrice, 64)
	return s
}

func ParseServers(servers []Server) []Server {
	for i, s := range servers {
		servers[i] = s.Parse()
	}
	return servers
}

func ServersDedup(servers []Server) []Server {
	uniqMap := map[string]Server{}
	var uniqSlice []Server
	for _, server := range servers {
		hash, err := server.HashCode()
		if err != nil {
			log.Println("failed hashcode server", err)
			uniqSlice = append(uniqSlice, server)
			continue
		}
		other, exists := uniqMap[hash]
		if !exists {
			uniqMap[hash] = server
			continue
		}
		swap := other.PriceEur > server.PriceEur
		if swap {
			uniqMap[hash] = server
			log.Printf("dedup replace %s %0.2f %0.2f \n%+v\n%+v\n\n", hash, other.PriceEur, server.PriceEur, other, server)
		}
	}
	for _, server := range uniqMap {
		uniqSlice = append(uniqSlice, server)
	}
	return uniqSlice
}

func (s Server) HashCode() (string, error) {
	s.Freetext = ""
	s.Key = 0
	s.Name = ""
	s.NextReduce = 0
	s.NextReduceHr = ""
	s.Price = ""
	s.PriceEur = 0
	s.PriceV = ""
	s.PriceVEur = 0
	s.SetupPrice = ""
	s.SetupPriceEur = 0
	s.Datacenter = nil
	s.RAMHr = ""
	s.Description = nil
	s.Freetext = ""
	return HashCode(s)
}

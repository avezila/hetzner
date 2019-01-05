package main

import (
	"math"
)

type AServer struct {
	EServer
	HddGbPerEurAdj    float64
	SsdGbPerEurAdj    float64
	RAMGbPerEurAdj    float64
	CPUBenchPerEurAdj float64
	Average           float64
}

func (e EServer) AServer(agg Aggregated) AServer {
	a := AServer{EServer: e}
	if agg.MaxHddGbPerEur != 0 {
		a.HddGbPerEurAdj = e.HddGbPerEur / agg.MaxHddGbPerEur
	}
	if agg.MaxSsdGbPerEur != 0 {
		a.SsdGbPerEurAdj = e.SsdGbPerEur / agg.MaxSsdGbPerEur
	}
	if agg.MaxRAMGbPerEur != 0 {
		a.RAMGbPerEurAdj = e.RAMGbPerEur / agg.MaxRAMGbPerEur
	}
	if agg.MaxCPUBenchPerEur != 0 {
		a.CPUBenchPerEurAdj = e.CPUBenchPerEur / agg.MaxCPUBenchPerEur
	}

	n := 0
	if a.HddGbPerEurAdj != 0 {
		a.Average += a.HddGbPerEurAdj * a.HddGbPerEurAdj
		n++
	}
	if a.SsdGbPerEurAdj != 0 {
		a.Average += a.SsdGbPerEurAdj * a.SsdGbPerEurAdj * 4
		n += 4
	}
	if a.RAMGbPerEurAdj != 0 {
		a.Average += a.RAMGbPerEurAdj * a.RAMGbPerEurAdj * 2
		n += 2
	}
	if a.CPUBenchPerEurAdj != 0 {
		a.Average += a.CPUBenchPerEurAdj * a.CPUBenchPerEurAdj * 8
		n += 8
	}
	if !a.OnlyHDD && !a.OnlySSD {
		a.Average++
		n++
	}
	if a.IsHighio {
		a.Average++
		n++
	}
	if a.HddCount > 2 {
		a.Average++
		n++
	}
	if a.IsEcc {
		a.Average++
		n++
	}
	if n > 0 {
		a.Average = math.Sqrt(a.Average) / float64(n)
	}

	return a
}

func (a AServer) Round() AServer {
	a.HddGbPerEurAdj = math.Ceil(a.HddGbPerEurAdj*100) / 100
	a.SsdGbPerEurAdj = math.Ceil(a.SsdGbPerEurAdj*100) / 100
	a.RAMGbPerEurAdj = math.Ceil(a.RAMGbPerEurAdj*100) / 100
	a.CPUBenchPerEurAdj = math.Ceil(a.CPUBenchPerEurAdj*100) / 100
	a.Average = math.Ceil(a.Average*100) / 100

	a.HddGbPerEur = math.Ceil(a.HddGbPerEur*100) / 100
	a.RAMGbPerEur = math.Ceil(a.RAMGbPerEur*100) / 100
	a.SsdGbPerEur = math.Ceil(a.SsdGbPerEur*100) / 100
	a.CPUBenchPerEur = math.Ceil(a.CPUBenchPerEur*100) / 100

	a.PriceEur = math.Ceil(a.PriceEur*100) / 100
	a.PriceVEur = math.Ceil(a.PriceVEur*100) / 100
	a.SetupPriceEur = math.Ceil(a.SetupPriceEur*100) / 100
	return a
}

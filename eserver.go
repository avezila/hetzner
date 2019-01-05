package main

import "strings"

type EServer struct {
	Server
	Hash           string
	OnlySSD        bool
	OnlyHDD        bool
	HddGbPerEur    float64
	SsdGbPerEur    float64
	RAMGbPerEur    float64
	CPUBenchPerEur float64
}

func (s Server) EServer() (EServer, error) {
	e := EServer{Server: s.Parse()}
	freeText := strings.ToLower(e.Freetext)

	e.OnlyHDD = strings.Contains(freeText, "hdd") && !strings.Contains(freeText, "ssd")
	e.OnlySSD = strings.Contains(freeText, "ssd") && !strings.Contains(freeText, "hdd")

	if e.PriceEur != 0 {
		if e.OnlySSD {
			e.SsdGbPerEur = float64(e.HddSize*e.HddCount) / e.PriceEur
		} else {
			e.HddGbPerEur = float64(e.HddSize*e.HddCount) / e.PriceEur
		}
		e.RAMGbPerEur = float64(e.RAM) / e.PriceEur
		e.CPUBenchPerEur = float64(e.CPUBenchmark) / e.PriceEur
	}

	hash, err := s.HashCode()
	if err != nil {
		return e, err
	}
	e.Hash = hash

	return e, nil
}

package main

type Aggregated struct {
	MinPrice        float64
	MinCPUBenchmark int64
	MinHddSize      int64
	MinSsdSize      int64
	MinRAM          int64

	MaxPrice        float64
	MaxCPUBenchmark int64
	MaxHddSize      int64
	MaxSsdSize      int64
	MaxRAM          int64

	MinHddGbPerEur    float64
	MinSsdGbPerEur    float64
	MinRAMGbPerEur    float64
	MinCPUBenchPerEur float64

	MaxHddGbPerEur    float64
	MaxSsdGbPerEur    float64
	MaxRAMGbPerEur    float64
	MaxCPUBenchPerEur float64
}

func Aggregate(servers []EServer) Aggregated {
	a := Aggregated{}
	for _, s := range servers {
		a.MinPrice = minF(a.MinPrice, s.PriceEur)
		a.MinCPUBenchmark = minI(a.MinCPUBenchmark, s.CPUBenchmark)
		if s.OnlySSD {
			a.MinSsdSize = minI(a.MinSsdSize, s.HddSize)
			a.MaxSsdSize = maxI(a.MaxSsdSize, s.HddSize)
		} else {
			a.MinHddSize = minI(a.MinHddSize, s.HddSize)
			a.MaxHddSize = maxI(a.MaxHddSize, s.HddSize)
		}
		a.MinRAM = minI(a.MinRAM, s.RAM)

		a.MaxPrice = maxF(a.MaxPrice, s.PriceEur)
		a.MaxCPUBenchmark = maxI(a.MaxCPUBenchmark, s.CPUBenchmark)
		a.MaxRAM = maxI(a.MaxRAM, s.RAM)

		a.MinHddGbPerEur = minF(a.MinHddGbPerEur, s.HddGbPerEur)
		a.MinSsdGbPerEur = minF(a.MinSsdGbPerEur, s.SsdGbPerEur)
		a.MinRAMGbPerEur = minF(a.MinRAMGbPerEur, s.RAMGbPerEur)
		a.MinCPUBenchPerEur = minF(a.MinCPUBenchPerEur, s.CPUBenchPerEur)

		a.MaxHddGbPerEur = maxF(a.MaxHddGbPerEur, s.HddGbPerEur)
		a.MaxSsdGbPerEur = maxF(a.MaxSsdGbPerEur, s.SsdGbPerEur)
		a.MaxRAMGbPerEur = maxF(a.MaxRAMGbPerEur, s.RAMGbPerEur)
		a.MaxCPUBenchPerEur = maxF(a.MaxCPUBenchPerEur, s.CPUBenchPerEur)
	}
	return a
}

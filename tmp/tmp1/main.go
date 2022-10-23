package main

import (
	"fmt"
	"sort"
)

// _ "net/http/pprof"

var g_StatTimeCost []int64

func main() {
	g_StatTimeCost = []int64{}
	for i := 1; i < 1001; i++ {
		g_StatTimeCost = append(g_StatTimeCost, int64(i))
	}
	printTimeCostProportion()
}

func printTimeCostProportion() {
	sort.Slice(g_StatTimeCost, func(i, j int) bool { return g_StatTimeCost[i] < g_StatTimeCost[j] })
	proportions := []float64{0.5, 0.6, 0.7, 0.8, 0.9}
	proportions = append(proportions, 0.91, 0.92, 0.93, 0.94, 0.95, 0.96, 0.97, 0.98, 0.99)
	proportions = append(proportions, 0.991, 0.992, 0.993, 0.994, 0.995, 0.996, 0.997, 0.998, 0.999)
	l := len(g_StatTimeCost)
	for _, v := range proportions {
		percent := v * 100
		idx := int(float64(l)*v) - 1
		if idx < 0 {
			idx = 0
		}
		respTime := g_StatTimeCost[idx]
		fmt.Printf("%.1f%% 响应时间小于 %dms\n", percent, respTime)
	}
	fmt.Printf("100%% 响应时间小于 %dms\n", g_StatTimeCost[len(g_StatTimeCost)-1])
}

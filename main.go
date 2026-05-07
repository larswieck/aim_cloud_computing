package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"deploy-tracker/tracker"
)

func main() {
	successful := envInt("SUCCESSFUL_DEPLOYS", 47)
	total := envInt("TOTAL_DEPLOYS", 50)
	deploys := envInt("DEPLOYS_THIS_WEEK", 12)
	days := envInt("DAYS", 7)
	outageMins := parseFloats(os.Getenv("OUTAGE_MINUTES"))

	rate := tracker.CalculateSuccessRate(successful, total)
	healthy := tracker.IsHealthy(rate, 95.0)
	mttr := tracker.MeanTimeToRecovery(outageMins)
	freq := tracker.DeploymentFrequency(deploys, days)

	fmt.Println("=== Deployment Health Report ===")
	fmt.Printf("Success Rate:          %.1f%%\n", rate)
	fmt.Printf("System Status:         %s\n", statusLabel(healthy))
	fmt.Printf("Mean Time to Recovery: %.1f min\n", mttr)
	fmt.Printf("Deployment Frequency:  %.2f deploys/day\n", freq)
}

func statusLabel(healthy bool) string {
	if healthy {
		return "HEALTHY"
	}
	return "DEGRADED"
}

func envInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}

func parseFloats(s string) []float64 {
	if s == "" {
		return nil
	}
	var out []float64
	for _, p := range strings.Split(s, ",") {
		f, err := strconv.ParseFloat(strings.TrimSpace(p), 64)
		if err == nil {
			out = append(out, f)
		}
	}
	return out
}

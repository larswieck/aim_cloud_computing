package tracker

// CalculateSuccessRate returns the deployment success rate as a percentage.
func CalculateSuccessRate(successful, total int) float64 {
	if total == 0 {
		return 0.0
	}
	return float64(successful) / float64(total) * 100
}

// IsHealthy reports whether a success rate meets the given threshold.
func IsHealthy(successRate, threshold float64) bool {
	return successRate >= threshold
}

// MeanTimeToRecovery returns the average outage duration in minutes.
func MeanTimeToRecovery(outageMins []float64) float64 {
	if len(outageMins) == 0 {
		return 0.0
	}
	var sum float64
	for _, m := range outageMins {
		sum += m
	}
	return sum / float64(len(outageMins))
}

// DeploymentFrequency returns the average number of deployments per day.
func DeploymentFrequency(deploys, days int) float64 {
	if days == 0 {
		return 0.0
	}
	return float64(deploys) / float64(days)
}

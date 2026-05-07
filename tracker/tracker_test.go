package tracker_test

import (
	"testing"

	"deploy-tracker/tracker"
)

func TestCalculateSuccessRate(t *testing.T) {
	got := tracker.CalculateSuccessRate(95, 100)
	if got != 95.0 {
		t.Errorf("CalculateSuccessRate(95, 100) = %f; want 95.0", got)
	}
}

func TestCalculateSuccessRateZeroTotal(t *testing.T) {
	got := tracker.CalculateSuccessRate(0, 0)
	if got != 0.0 {
		t.Errorf("CalculateSuccessRate(0, 0) = %f; want 0.0", got)
	}
}

func TestIsHealthy(t *testing.T) {
	if !tracker.IsHealthy(96.0, 95.0) {
		t.Error("IsHealthy(96.0, 95.0) = false; want true")
	}
	if tracker.IsHealthy(94.0, 95.0) {
		t.Error("IsHealthy(94.0, 95.0) = true; want false")
	}
}

func TestIsHealthyAtThreshold(t *testing.T) {
	if !tracker.IsHealthy(95.0, 95.0) {
		t.Error("IsHealthy(95.0, 95.0) = false; want true (threshold is inclusive)")
	}
}

func TestMeanTimeToRecovery(t *testing.T) {
	got := tracker.MeanTimeToRecovery([]float64{10, 20, 30})
	if got != 20.0 {
		t.Errorf("MeanTimeToRecovery([10,20,30]) = %f; want 20.0", got)
	}
}

func TestMeanTimeToRecoveryEmpty(t *testing.T) {
	got := tracker.MeanTimeToRecovery([]float64{})
	if got != 0.0 {
		t.Errorf("MeanTimeToRecovery([]) = %f; want 0.0", got)
	}
}

func TestDeploymentFrequencyNormal(t *testing.T) {
	got := tracker.DeploymentFrequency(14, 7)
	if got != 2.0 {
		t.Errorf("DeploymentFrequency(14, 7) = %f; want 2.0", got)
	}
}

func TestDeploymentFrequencyZeroDays(t *testing.T) {
	got := tracker.DeploymentFrequency(5, 0)
	if got != 0.0 {
		t.Errorf("DeploymentFrequency(5, 0) = %f; want 0.0", got)
	}
}

func TestDeploymentFrequencyZeroDeploys(t *testing.T) {
	got := tracker.DeploymentFrequency(0, 30)
	if got != 0.0 {
		t.Errorf("DeploymentFrequency(0, 30) = %f; want 0.0", got)
	}
}

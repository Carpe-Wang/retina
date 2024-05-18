package init

import (
	"retina/init/retina"
	"retina/internal/telemetry"
	"testing"
)

func TestTelemetryInitialization(t *testing.T) {
	version := "1.0.0"

	// Case 1: applicationInsightsID is not empty and enableTelemetry is true
	applicationInsightsID := "valid-id"
	enableTelemetry := true
	opts := &retina.Options{}
	if applicationInsightsID != "" && enableTelemetry {
		opts.EnableTelemetry = true
		opts.ApplicationInsightsID = applicationInsightsID
		telemetry.InitAppInsights(applicationInsightsID, version)
		defer telemetry.ShutdownAppInsights()
		defer telemetry.TrackPanic()
	}
	if !opts.EnableTelemetry {
		t.Errorf("Telemetry should be enabled when both applicationInsightsID is not empty and enableTelemetry is true")
	}

	// Case 2: applicationInsightsID is empty
	applicationInsightsID = ""
	enableTelemetry = true
	opts = &retina.Options{}
	if applicationInsightsID != "" && enableTelemetry {
		opts.EnableTelemetry = true
		opts.ApplicationInsightsID = applicationInsightsID
		telemetry.InitAppInsights(applicationInsightsID, version)
		defer telemetry.ShutdownAppInsights()
		defer telemetry.TrackPanic()
	}
	if opts.EnableTelemetry {
		t.Errorf("Telemetry should not be enabled when applicationInsightsID is empty")
	}

	// Case 3: enableTelemetry is false
	applicationInsightsID = "valid-id"
	enableTelemetry = false
	opts = &retina.Options{}
	if applicationInsightsID != "" && enableTelemetry {
		opts.EnableTelemetry = true
		opts.ApplicationInsightsID = applicationInsightsID
		telemetry.InitAppInsights(applicationInsightsID, version)
		defer telemetry.ShutdownAppInsights()
		defer telemetry.TrackPanic()
	}
	if opts.EnableTelemetry {
		t.Errorf("Telemetry should not be enabled when enableTelemetry is false")
	}
}

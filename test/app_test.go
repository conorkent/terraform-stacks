package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestAppWorkspaceUsesComputeInstanceIDOutput tests that the app workspace
// correctly uses the instance_id output from the compute workspace to form
// its output message.
func TestAppWorkspaceUsesComputeInstanceIDOutput(t *testing.T) {
	t.Parallel()

	// Mock instance ID for testing app workspace in isolation
	mockInstanceID := "i-1234567890abcdef0"

	appOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../app",
		Vars: map[string]interface{}{
			"region":      "us-east-1",
			"instance_id": mockInstanceID,
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, appOptions)
	terraform.InitAndApply(t, appOptions)

	// Verify the message output contains the instance_id
	message := terraform.Output(t, appOptions, "message")
	expectedMessage := fmt.Sprintf("App layer deployed on instance: %s", mockInstanceID)
	assert.Equal(t, expectedMessage, message, "Output message should contain the instance_id")
}

// TestAppWorkspaceWithActualComputeOutput tests the app workspace with
// real outputs from the compute workspace.
func TestAppWorkspaceWithActualComputeOutput(t *testing.T) {
	t.Parallel()

	// Create networking workspace
	networkingOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../networking",
		Vars: map[string]interface{}{
			"region": "us-west-2",
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, networkingOptions)
	terraform.InitAndApply(t, networkingOptions)

	vpcID := terraform.Output(t, networkingOptions, "vpc_id")

	// Create compute workspace
	computeOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../compute",
		Vars: map[string]interface{}{
			"region": "us-west-2",
			"vpc_id": vpcID,
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, computeOptions)
	terraform.InitAndApply(t, computeOptions)

	instanceID := terraform.Output(t, computeOptions, "instance_id")

	// Create app workspace with compute output
	appOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../app",
		Vars: map[string]interface{}{
			"region":      "us-west-2",
			"instance_id": instanceID,
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, appOptions)
	terraform.InitAndApply(t, appOptions)

	// Verify the message contains the actual instance ID
	message := terraform.Output(t, appOptions, "message")
	assert.Contains(t, message, instanceID, "Message should contain the actual instance ID from compute workspace")
	assert.Contains(t, message, "App layer deployed on instance:", "Message should have expected format")
}

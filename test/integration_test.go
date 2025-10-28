package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestFullStackIntegration tests the complete workflow:
// 1. Networking workspace creates VPC
// 2. Compute workspace creates EC2 instance after networking
// 3. App workspace uses compute output
// This simulates the dependency chain defined in stack.hcl
func TestFullStackIntegration(t *testing.T) {
	t.Parallel()

	region := "us-west-2"

	// Step 1: Deploy networking workspace
	networkingOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../networking",
		Vars: map[string]interface{}{
			"region": region,
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, networkingOptions)
	terraform.InitAndApply(t, networkingOptions)

	// Verify networking outputs
	vpcID := terraform.Output(t, networkingOptions, "vpc_id")
	assert.NotEmpty(t, vpcID, "Networking workspace should output VPC ID")
	assert.Regexp(t, "^vpc-[a-z0-9]+$", vpcID, "VPC ID should be valid")

	// Step 2: Deploy compute workspace (depends on networking)
	computeOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../compute",
		Vars: map[string]interface{}{
			"region": region,
			"vpc_id": vpcID,
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, computeOptions)
	terraform.InitAndApply(t, computeOptions)

	// Verify compute outputs
	instanceID := terraform.Output(t, computeOptions, "instance_id")
	assert.NotEmpty(t, instanceID, "Compute workspace should output instance ID")
	assert.Regexp(t, "^i-[a-z0-9]+$", instanceID, "Instance ID should be valid")

	// Step 3: Deploy app workspace (depends on compute)
	appOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../app",
		Vars: map[string]interface{}{
			"region":      region,
			"instance_id": instanceID,
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, appOptions)
	terraform.InitAndApply(t, appOptions)

	// Verify app outputs
	message := terraform.Output(t, appOptions, "message")
	assert.Contains(t, message, instanceID, "App message should contain instance ID from compute")
	assert.Equal(t, "App layer deployed on instance: "+instanceID, message, "Message should match expected format")
}

// TestWorkspaceDependencyOrder verifies that workspaces must be applied
// in the correct order as defined by depends_on in stack.hcl
func TestWorkspaceDependencyOrder(t *testing.T) {
	t.Parallel()

	region := "us-east-1"

	// Test that compute fails without networking outputs
	computeOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../compute",
		Vars: map[string]interface{}{
			"region": region,
			// Missing vpc_id - should fail or require manual input
		},
		NoColor: true,
	})

	// This should fail during plan/apply because vpc_id is required
	_, err := terraform.InitAndPlanE(t, computeOptions)
	assert.Error(t, err, "Compute workspace should require vpc_id from networking")

	// Test that app fails without compute outputs
	appOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../app",
		Vars: map[string]interface{}{
			"region": region,
			// Missing instance_id - should fail or require manual input
		},
		NoColor: true,
	})

	// This should fail during plan/apply because instance_id is required
	_, err = terraform.InitAndPlanE(t, appOptions)
	assert.Error(t, err, "App workspace should require instance_id from compute")
}

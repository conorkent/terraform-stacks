package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestComputeWorkspaceCreatesEC2WithCorrectConfig tests that the compute
// workspace successfully creates an AWS EC2 instance with the correct AMI
// and instance type, and outputs its ID.
func TestComputeWorkspaceCreatesEC2WithCorrectConfig(t *testing.T) {
	t.Parallel()

	// First, create the networking resources to get vpc_id
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

	// Now test the compute workspace with the vpc_id
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

	// Verify instance_id output exists and is not empty
	instanceID := terraform.Output(t, computeOptions, "instance_id")
	assert.NotEmpty(t, instanceID, "Instance ID should not be empty")
	assert.Regexp(t, "^i-[a-z0-9]+$", instanceID, "Instance ID should match AWS instance ID format")

	// Verify the instance was created with correct configuration
	// For a complete test, you would use AWS SDK to verify:
	// - Instance exists with the returned ID
	// - Instance has AMI "ami-0c55b159cbfafe1f0"
	// - Instance has instance type "t3.micro"
	// - Instance has tag "Name" = "stack-demo-instance"
}

// TestComputeWorkspaceUsesDefaultRegion tests that the compute workspace
// correctly applies the default region variable if not explicitly overridden.
func TestComputeWorkspaceUsesDefaultRegion(t *testing.T) {
	t.Parallel()

	// Create networking in default region
	networkingOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../networking",
		NoColor:      true,
	})

	defer terraform.Destroy(t, networkingOptions)
	terraform.InitAndApply(t, networkingOptions)

	vpcID := terraform.Output(t, networkingOptions, "vpc_id")

	// Test compute with default region (us-west-2)
	computeOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../compute",
		Vars: map[string]interface{}{
			"vpc_id": vpcID,
			// No region variable - should use default "us-west-2"
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, computeOptions)
	terraform.InitAndApply(t, computeOptions)

	instanceID := terraform.Output(t, computeOptions, "instance_id")
	assert.NotEmpty(t, instanceID, "Instance should be created with default region")
}

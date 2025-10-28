# Quick Start Guide

## Test Suite Overview

This test suite validates all 5 requirements:

1. ✅ **Networking workspace** creates VPC with correct tags and outputs
2. ✅ **Compute workspace** creates EC2 with correct AMI/type and outputs, after networking
3. ✅ **App workspace** uses `instance_id` output from compute correctly
4. ✅ **stack.hcl** defines correct workspace dependencies
5. ✅ **Default regions** are applied correctly when not overridden

## Quick Commands

### Configuration tests only (no AWS resources created)
```bash
make test-config
```

### Run all tests (creates AWS resources - costs may apply)
```bash
make test
```

### Run specific test
```bash
go test -v -run TestStackConfigDefinesCorrectDependencies
```

## Test Mapping

| Requirement | Test Function | File |
|------------|---------------|------|
| 1. Networking VPC | `TestNetworkingWorkspaceCreatesVPCWithCorrectTagsAndOutputs` | `networking_test.go` |
| 2. Compute EC2 | `TestComputeWorkspaceCreatesEC2WithCorrectConfig` | `compute_test.go` |
| 3. App uses instance_id | `TestAppWorkspaceUsesComputeInstanceIDOutput` | `app_test.go` |
| 4. Stack dependencies | `TestStackConfigDefinesCorrectDependencies` | `stack_config_test.go` |
| 5. Default regions | `TestNetworkingWorkspaceUsesDefaultRegion` + `TestComputeWorkspaceUsesDefaultRegion` | `networking_test.go` + `compute_test.go` |

## First Run

```bash
# Install dependencies
make install

# Run config tests (fast, no AWS resources)
make test-config

# If AWS credentials are configured, run full suite
make test
```

## AWS Configuration

For tests that create resources, ensure AWS credentials are configured:

```bash
export AWS_ACCESS_KEY_ID="your-key"
export AWS_SECRET_ACCESS_KEY="your-secret"
# OR
export AWS_PROFILE="your-profile"
```

**Note**: Tests automatically clean up resources, but charges may still apply.

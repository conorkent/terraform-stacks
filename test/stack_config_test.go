package test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestStackConfigDefinesCorrectDependencies tests that the stack.hcl
// configuration correctly defines the dependencies between the networking,
// compute, and app workspaces.
func TestStackConfigDefinesCorrectDependencies(t *testing.T) {
	t.Parallel()

	// Read the stack.hcl file
	stackContent, err := os.ReadFile("../stack.hcl")
	require.NoError(t, err, "Should be able to read stack.hcl")

	content := string(stackContent)

	// Verify stack definition exists
	assert.Contains(t, content, "stack \"three-tier\"", "Stack should be named 'three-tier'")

	// Verify networking workspace exists with no dependencies
	assert.Contains(t, content, "workspace \"networking\"", "Networking workspace should be defined")
	networkingWorkspace := extractWorkspaceBlock(content, "networking")
	assert.NotContains(t, networkingWorkspace, "depends_on", "Networking workspace should not have dependencies")

	// Verify compute workspace exists and depends on networking
	assert.Contains(t, content, "workspace \"compute\"", "Compute workspace should be defined")
	computeWorkspace := extractWorkspaceBlock(content, "compute")
	assert.Contains(t, computeWorkspace, "depends_on", "Compute workspace should have dependencies")
	assert.Contains(t, computeWorkspace, "[\"networking\"]", "Compute should depend on networking")

	// Verify app workspace exists and depends on compute
	assert.Contains(t, content, "workspace \"app\"", "App workspace should be defined")
	appWorkspace := extractWorkspaceBlock(content, "app")
	assert.Contains(t, appWorkspace, "depends_on", "App workspace should have dependencies")
	assert.Contains(t, appWorkspace, "[\"compute\"]", "App should depend on compute")
}

// TestStackConfigDefinesCorrectSources tests that each workspace
// points to the correct source directory.
func TestStackConfigDefinesCorrectSources(t *testing.T) {
	t.Parallel()

	stackContent, err := os.ReadFile("../stack.hcl")
	require.NoError(t, err, "Should be able to read stack.hcl")

	content := string(stackContent)

	// Verify workspace sources
	networkingWorkspace := extractWorkspaceBlock(content, "networking")
	assert.Contains(t, networkingWorkspace, "source = \"./networking\"", "Networking workspace should have correct source")

	computeWorkspace := extractWorkspaceBlock(content, "compute")
	assert.Contains(t, computeWorkspace, "source = \"./compute\"", "Compute workspace should have correct source")

	appWorkspace := extractWorkspaceBlock(content, "app")
	assert.Contains(t, appWorkspace, "source = \"./app\"", "App workspace should have correct source")
}

// TestStackConfigSetsDefaultRegion tests that the stack configuration
// sets a default region variable.
func TestStackConfigSetsDefaultRegion(t *testing.T) {
	t.Parallel()

	stackContent, err := os.ReadFile("../stack.hcl")
	require.NoError(t, err, "Should be able to read stack.hcl")

	content := string(stackContent)

	// Verify stack sets region variable
	stackBlock := extractStackBlock(content, "three-tier")
	assert.Contains(t, stackBlock, "variables", "Stack should define variables")
	assert.Contains(t, stackBlock, "region", "Stack should set region variable")
	assert.Contains(t, stackBlock, "us-east-1", "Stack should set region to us-east-1")
}

// Helper function to extract a workspace block from HCL content
func extractWorkspaceBlock(content, workspaceName string) string {
	startMarker := "workspace \"" + workspaceName + "\""
	startIdx := strings.Index(content, startMarker)
	if startIdx == -1 {
		return ""
	}

	// Find the opening brace
	openBraceIdx := strings.Index(content[startIdx:], "{")
	if openBraceIdx == -1 {
		return ""
	}

	// Find the matching closing brace
	braceCount := 1
	currentIdx := startIdx + openBraceIdx + 1
	for currentIdx < len(content) && braceCount > 0 {
		if content[currentIdx] == '{' {
			braceCount++
		} else if content[currentIdx] == '}' {
			braceCount--
		}
		currentIdx++
	}

	return content[startIdx:currentIdx]
}

// Helper function to extract a stack block from HCL content
func extractStackBlock(content, stackName string) string {
	startMarker := "stack \"" + stackName + "\""
	startIdx := strings.Index(content, startMarker)
	if startIdx == -1 {
		return ""
	}

	// Find the opening brace
	openBraceIdx := strings.Index(content[startIdx:], "{")
	if openBraceIdx == -1 {
		return ""
	}

	// Find the matching closing brace
	braceCount := 1
	currentIdx := startIdx + openBraceIdx + 1
	for currentIdx < len(content) && braceCount > 0 {
		if content[currentIdx] == '{' {
			braceCount++
		} else if content[currentIdx] == '}' {
			braceCount--
		}
		currentIdx++
	}

	return content[startIdx:currentIdx]
}

# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

This repository is for testing Terraform Cloud (TFC) Stacks functionality. The project is currently in early stages with minimal structure.

## Repository Purpose

- Test Terraform Cloud Stacks features and workflows
- Experiment with TFC stack configurations and deployments
- Development and validation of stack-based infrastructure patterns

## Expected Development Patterns

### Terraform Cloud Stacks
When working with TFC Stacks in this repository:
- Stacks are Terraform Cloud's approach to managing collections of infrastructure across multiple workspaces
- Stack configurations typically include deployment files (`deployments.tfdeploy.hcl`), component definitions, and orchestration logic
- TFC Stacks use HCL for both infrastructure definition and orchestration workflows

### Likely File Structure
As this repository grows, expect:
- `*.tfstack.hcl` files for stack configurations
- `deployments.tfdeploy.hcl` files for deployment definitions
- Component directories containing reusable Terraform modules
- `terraform.tf` or `*.tf` files for standard Terraform configurations

## Development Workflow

Since this is a testing repository, typical workflows will involve:
- Creating and modifying stack configurations
- Testing stack deployments against TFC
- Validating component interactions and dependencies
- Iterating on infrastructure patterns

## License

This project uses the MIT License.

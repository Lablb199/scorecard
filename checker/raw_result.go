// Copyright 2020 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package checker

import (
	"github.com/ossf/scorecard/v4/clients"
)

// RawResults contains results before a policy
// is applied.
//nolint
type RawResults struct {
	CIIBestPracticesResults     CIIBestPracticesData
	DangerousWorkflowResults    DangerousWorkflowData
	VulnerabilitiesResults      VulnerabilitiesData
	BinaryArtifactResults       BinaryArtifactData
	SecurityPolicyResults       SecurityPolicyData
	DependencyUpdateToolResults DependencyUpdateToolData
	BranchProtectionResults     BranchProtectionsData
	CodeReviewResults           CodeReviewData
	WebhookResults              WebhooksData
	ContributorsResults         ContributorsData
	MaintainedResults           MaintainedData
	SignedReleasesResults       SignedReleasesData
	FuzzingResults              FuzzingData
	LicenseResults              LicenseData
}

// FuzzingData represents different fuzzing done.
type FuzzingData struct {
	Fuzzers []Tool
}

// MaintainedData contains the raw results
// for the Maintained check.
type MaintainedData struct {
	Issues               []clients.Issue
	DefaultBranchCommits []clients.Commit
	ArchivedStatus       ArchivedStatus
}

// LicenseData contains the raw results
// for the License check.
type LicenseData struct {
	Files []File
}

// CodeReviewData contains the raw results
// for the Code-Review check.
type CodeReviewData struct {
	DefaultBranchCommits []clients.Commit
}

// ContributorsData represents contributor information.
type ContributorsData struct {
	Users []clients.Contributor
}

// VulnerabilitiesData contains the raw results
// for the Vulnerabilities check.
type VulnerabilitiesData struct {
	Vulnerabilities []clients.Vulnerability
}

// SecurityPolicyData contains the raw results
// for the Security-Policy check.
type SecurityPolicyData struct {
	// Files contains a list of files.
	Files []File
}

// BinaryArtifactData contains the raw results
// for the Binary-Artifact check.
type BinaryArtifactData struct {
	// Files contains a list of files.
	Files []File
}

// SignedReleasesData contains the raw results
// for the Signed-Releases check.
type SignedReleasesData struct {
	Releases []clients.Release
}

// DependencyUpdateToolData contains the raw results
// for the Dependency-Update-Tool check.
type DependencyUpdateToolData struct {
	// Tools contains a list of tools.
	// Note: we only populate one entry at most.
	Tools []Tool
}

// WebhooksData contains the raw results
// for the Webhook check.
type WebhooksData struct {
	Webhooks []clients.Webhook
}

// BranchProtectionsData contains the raw results
// for the Branch-Protection check.
type BranchProtectionsData struct {
	Branches []clients.BranchRef
}

// Tool represents a tool.
type Tool struct {
	URL  *string
	Desc *string
	File *File
	Name string
	// Runs of the tool.
	Runs []Run
	// Issues created by the tool.
	Issues []clients.Issue
	// Merge requests created by the tool.
	MergeRequests []clients.PullRequest

	// TODO: CodeCoverage, jsonWorkflowJob.
}

// Run represents a run.
type Run struct {
	URL string
	// TODO: add fields, e.g., Result=["success", "failure"]
}

// ArchivedStatus definess the archived status.
type ArchivedStatus struct {
	Status bool
	// TODO: add fields, e.g., date of archival.
}

// File represents a file.
type File struct {
	Path    string
	Snippet string   // Snippet of code
	Offset  uint     // Offset in the file of Path (line for source/text files).
	Type    FileType // Type of file.
	// TODO: add hash.
}

// CIIBestPracticesData contains data foor CIIBestPractices check.
type CIIBestPracticesData struct {
	Badge clients.BadgeLevel
}

// DangerousWorkflowType represents a type of dangerous workflow.
type DangerousWorkflowType string

const (
	// DangerousWorkflowScriptInjection represents a script injection.
	DangerousWorkflowScriptInjection DangerousWorkflowType = "scriptInjection"
	// DangerousWorkflowUntrustedCheckout represents an untrusted checkout.
	DangerousWorkflowUntrustedCheckout DangerousWorkflowType = "untrustedCheckout"
)

// DangerousWorkflowData contains raw results
// for dangerous workflow check.
type DangerousWorkflowData struct {
	Workflows []DangerousWorkflow
}

// DangerousWorkflow represents a dangerous workflow.
type DangerousWorkflow struct {
	Job  *WorkflowJob
	Type DangerousWorkflowType
	File File
}

// WorkflowJob reprresents a workflow job.
type WorkflowJob struct {
	Name *string
	ID   *string
}

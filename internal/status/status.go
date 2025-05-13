package status

import (
	"fmt"
	"os/exec"
	"strings"
)

// Type that represents the possible branch statuses
type BranchStatus rune

const (
	BranchStatusNone     BranchStatus = 0
	BranchStatusUnstaged BranchStatus = '*'
	BranchStatusStaged   BranchStatus = '+'
)

type GitInfo struct {
	BranchName string
	Status     BranchStatus
}

// Print a GitInfo object.
func (g *GitInfo) Print() {
	if g.Status == BranchStatusNone {
		fmt.Printf("[%s]", g.BranchName)
	} else {
		fmt.Printf("[%s %c]", g.BranchName, g.Status)
	}
}

// Check if GitInfo has been initialized.
func (g *GitInfo) IsEmpty() bool {
	return g.BranchName == ""
}

// Build the Git prompt
func GetGitInfo() (GitInfo, error) {
	name, err := GetGitBranch()
	if err != nil {
		return GitInfo{}, nil
	}

	status, err := GetBranchStatus()
	if err != nil {
		return GitInfo{}, nil
	}

	info := GitInfo{
		BranchName: name,
		Status:     status,
	}

	return info, nil
}

// Get the Git branch name
func GetGitBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	cleanOutput := strings.Split(string(output), "\n")

	return cleanOutput[0], nil
}

// Get the git branch status
func GetBranchStatus() (BranchStatus, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()

	if err != nil {
		return BranchStatusNone, err
	}

	cleanOutput := strings.Split(string(output), "\n")
	cleanOutput = cleanOutput[:len(cleanOutput)-1]

	if len(cleanOutput) <= 0 {
		return BranchStatusNone, nil
	}

	// Get first 2 characters from each line of 'git status --porcelain'
	var statuses []string
	for _, status := range cleanOutput {
		if len(status) >= 2 {
			statuses = append(statuses, status[:2])
		}
	}

	// Check for staged changes
	for _, status := range statuses {
		if len(status) != 2 {
			continue
		}

		if status[0] != ' ' && status[1] == ' ' {
			return BranchStatusStaged, nil
		}
	}

	return BranchStatusUnstaged, nil
}

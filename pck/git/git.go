package git

import (
	"io/ioutil"
	"os/exec"
	"strings"

	// go-git
	git "github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

// GetProjectTicket returns the project and ticket from the current branch
func GetProjectTicket() (string, error) {
	// Open repository
	r, err := git.PlainOpen(".")

	if err != nil {
		return "", err
	}

	// Get the branch name
	branch, err := r.Head()

	if err != nil {
		return "", err
	}

	// Branch name format: [PROJECT]-[TICKET]-[DESCRIPTION]
	// Example: AUR-2670-Add-Feature

	projectTicket := strings.Join(strings.Split(branch.Name().Short(), "-")[0:2], "-")

	// Print project and ticket
	return projectTicket, nil

}

func GetGitDiff() (string, error) {
	// Exec git diff
	diffCmd := exec.Command("git", "diff", "--cached")

	stdout, err := diffCmd.StdoutPipe()
	
	if err != nil {
		return "", err
	}
	
	diffCmd.Stderr = diffCmd.Stdout

	if err = diffCmd.Start(); err != nil {
		return "", err
	}
	
	gitDiff, err := ioutil.ReadAll(stdout)

	if err != nil {
		panic(err)
	}

	return string(gitDiff), nil
}

//get git commit log
func GetGitLog() (string, error) {

	// Exec git log
	logCmd := exec.Command("git", "log", "-5", "-p", "--pretty=%B")

	stdout, err := logCmd.StdoutPipe()
	
	if err != nil {
		return "", err
	}
	
	logCmd.Stderr = logCmd.Stdout

	if err = logCmd.Start(); err != nil {
		return "", err
	}
	
	gitLog, err := ioutil.ReadAll(stdout)

	if err != nil {
		panic(err)
	}

	return string(gitLog), nil
}

//get branch name
func GetBranchName() (string, error) {
	r, err := git.PlainOpen(".")

	if err != nil {
		return "", err
	}

	// Get the branch name
	branch, err := r.Head()

	if err != nil {
		return "", err
	}

	return branch.Name().Short(), nil

}
package model

import (
	"os/exec"
	"strings"
)

type GitAdapter struct{}

func (g *GitAdapter) Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	return cmd.Run()
}

func (g *GitAdapter) Diff() (string, error) {
	out, err := exec.Command("git", "diff").Output()
	return string(out), err
}

func (g *GitAdapter) CurrentBranch() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	return strings.TrimSpace(string(out)), err
}

type JJAdapter struct{}

func (j *JJAdapter) Commit(message string) error {
	// Jujutsu uses 'describe' to set messages on the current change
	cmd := exec.Command("jj", "describe", "-m", message)
	return cmd.Run()
}

func (j *JJAdapter) Diff() (string, error) {
	out, err := exec.Command("jj", "diff").Output()
	return string(out), err
}

func (j *JJAdapter) CurrentBranch() (string, error) {
	// JJ doesn't have "branches" in the same way, but it has bookmarks
	out, err := exec.Command("jj", "bookmark", "list").Output()
	return strings.TrimSpace(string(out)), err
}

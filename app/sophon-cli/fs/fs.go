package fs

import (
	"os"
	"os/exec"
	"path/filepath"
	"sophon-cli/term"
)

var Cwd string
var SophonDir string
var ProjectRoot string
var HomeSophonDir string
var CacheDir string

var HomeDir string
var HomeAuthPath string
var HomeAccountsPath string

func init() {
	var err error
	Cwd, err = os.Getwd()
	if err != nil {
		term.OutputErrorAndExit("Error getting current working directory: %v", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		term.OutputErrorAndExit("Couldn't find home dir: %v", err.Error())
	}
	HomeDir = home

	if os.Getenv("PLANDEX_ENV") == "development" {
		HomeSophonDir = filepath.Join(home, ".sophon-home-dev-v2")
	} else {
		HomeSophonDir = filepath.Join(home, ".sophon-home-v2")
	}

	// Create the home sophon directory if it doesn't exist
	err = os.MkdirAll(HomeSophonDir, os.ModePerm)
	if err != nil {
		term.OutputErrorAndExit(err.Error())
	}

	CacheDir = filepath.Join(HomeSophonDir, "cache")
	HomeAuthPath = filepath.Join(HomeSophonDir, "auth.json")
	HomeAccountsPath = filepath.Join(HomeSophonDir, "accounts.json")

	err = os.MkdirAll(filepath.Join(CacheDir, "tiktoken"), os.ModePerm)
	if err != nil {
		term.OutputErrorAndExit(err.Error())
	}
	err = os.Setenv("TIKTOKEN_CACHE_DIR", CacheDir)
	if err != nil {
		term.OutputErrorAndExit(err.Error())
	}

	FindSophonDir()
	if SophonDir != "" {
		ProjectRoot = Cwd
	}
}

func FindOrCreateSophon() (string, bool, error) {
	FindSophonDir()
	if SophonDir != "" {
		ProjectRoot = Cwd
		return SophonDir, false, nil
	}

	// Determine the directory path
	var dir string
	if os.Getenv("PLANDEX_ENV") == "development" {
		dir = filepath.Join(Cwd, ".sophon-dev-v2")
	} else {
		dir = filepath.Join(Cwd, ".sophon-v2")
	}

	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return "", false, err
	}
	SophonDir = dir
	ProjectRoot = Cwd

	return dir, true, nil
}

func ProjectRootIsGitRepo() bool {
	if ProjectRoot == "" {
		return false
	}

	return IsGitRepo(ProjectRoot)
}

func IsGitRepo(dir string) bool {
	isGitRepo := false

	if isCommandAvailable("git") {
		// check whether we're in a git repo
		cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")

		cmd.Dir = dir

		err := cmd.Run()

		if err == nil {
			isGitRepo = true
		}
	}

	return isGitRepo
}

func FindSophonDir() {
	SophonDir = findSophon(Cwd)
}

func findSophon(baseDir string) string {
	var dir string
	if os.Getenv("PLANDEX_ENV") == "development" {
		dir = filepath.Join(baseDir, ".sophon-dev-v2")
	} else {
		dir = filepath.Join(baseDir, ".sophon-v2")
	}
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return dir
	}

	return ""
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command(name, "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

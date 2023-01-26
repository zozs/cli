package testdata

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"testing"
)

func SetUpCiEnv(env map[string]string) error {
	for variable, value := range env {
		err := os.Setenv(variable, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetUpGitRepository(includeCommit bool) (string, error) {
	cwd, _ := os.Getwd()
	repoDir := cwd + "/testdata/"
	repo, err := git.PlainInit(repoDir, false)
	if err != nil {
		return cwd, err
	}
	w, err := repo.Worktree()
	if err != nil {
		return cwd, err
	}
	if includeCommit {
		_, err = w.Commit("Initial commit", &git.CommitOptions{Author: &object.Signature{Name: "author"}})
		if err != nil {
			return cwd, err
		}
	}

	err = os.Chdir(repoDir)
	return cwd, err
}

func TearDownGitRepository(dir string, t *testing.T) {
	cwd, _ := os.Getwd()
	repoDir := cwd + "/.git/"
	_, err := git.PlainOpen(repoDir)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}
	err = os.RemoveAll(repoDir)
	if err != nil {
		t.Fatal(err)
	}
}

func ResetEnv(ciEnv map[string]string, t *testing.T) {
	for _, variable := range ciEnv {
		err := os.Unsetenv(variable)
		if err != nil {
			t.Fatal(err)
		}
	}
}

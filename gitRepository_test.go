package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewGitRepository(t *testing.T) {
	workTree := "/home/marisa/proj/go-wyag"
	gitDir := filepath.Join(workTree, ".git")

	repo, _ := NewGitRepository(workTree, true)
	expectedRepo := &GitRepository{
		WorkTree: workTree,
		GitDir:   gitDir,
	}

	if repo.GitDir != expectedRepo.GitDir {
		assertStrings(t, repo.GitDir, expectedRepo.GitDir)
	}
}

func TestComputeRepoPath(t *testing.T) {
	workTree := "/home/marisa/proj/go-wyag"
	repo, _ := NewGitRepository(workTree, true)

	tests := []struct {
		name     string
		repoPath string
		want     string
	}{
		{
			"one path",
			repo.computeGitdirPath("alpha"),
			"/home/marisa/proj/go-wyag/.git/alpha",
		},
		{
			"multiple paths",
			repo.computeGitdirPath("alpha", "beta"),
			"/home/marisa/proj/go-wyag/.git/alpha/beta",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.repoPath != tt.want {
				assertStrings(t, tt.repoPath, tt.want)
			}
		})
	}
}

func TestMkdirGitdirPath(t *testing.T) {
	workTree := "/home/marisa/proj/go-wyag"
	repo, _ := NewGitRepository(workTree, true)

	// shouldMkdir = false tests
	tests1 := []struct {
		name  string
		paths []string
		want  string
	}{
		{
			"one path (shouldMkdir = false)",
			[]string{"alpha"},
			"",
		},
		{
			"multiple paths (shouldMkdir = false)",
			[]string{"alpha", "beta"},
			"",
		},
	}

	// shouldMkdir = true tests
	tests2 := []struct {
		name  string
		paths []string
		want  string
	}{
		{
			"one path (shouldMkdir = true)",
			[]string{"alpha"},
			"/home/marisa/proj/go-wyag/.git/alpha",
		},
		{
			"multiple paths (shouldMkdir = false)",
			[]string{"alpha", "beta"},
			"/home/marisa/proj/go-wyag/.git/alpha/beta",
		},
	}

	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			gitDirPath, err := repo.mkdirGitdirPath(false, tt.paths...)
			if err != nil {
				t.Error(err)
			}

			if gitDirPath != tt.want {
				assertStrings(t, gitDirPath, tt.want)
			}
		})
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			gitDirPath, err := repo.mkdirGitdirPath(true, tt.paths...)
			if err != nil {
				t.Error(err)
			}

			// Remove .git directory and all its children after
			// finishing one test so subsequent testings of
			// shouldMkdir = false tests work correctly
			defer os.RemoveAll(repo.GitDir)

			if gitDirPath != tt.want {
				assertStrings(t, gitDirPath, tt.want)
			}
		})
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but expected %q", got, want)
	}
}

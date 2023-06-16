package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type GitRepository struct {
	WorkTree string
	GitDir   string
	Conf     string
}

func NewGitRepository(workTreePath string, force bool) (*GitRepository, error) {
	repo := &GitRepository{
		WorkTree: workTreePath,
		GitDir:   filepath.Join(workTreePath, ".git"),
	}

	if _, err := os.Stat(repo.GitDir); os.IsNotExist(err) && !force {
		return nil, fmt.Errorf("Not a Git repository %s", workTreePath)
	}

	configDirPath, err := repo.makeGitdirFile(false, "config")
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(configDirPath)
	if err != nil {
		return nil, err
	}

	if configDirPath != "" && !os.IsNotExist(err) {
		cfg, err := ini.Load(repo.GitDir)
	}

	return repo, nil
}

// Compute path under repo's gitdir
func (repo *GitRepository) computeGitdirPath(paths ...string) string {
	gitDirPath := repo.GitDir
	for _, p := range paths {
		gitDirPath = filepath.Join(gitDirPath, p)
	}

	return gitDirPath
}

// Same as computeGitdirPath but mkdir paths if it doesn't exist
// and shouldMkdir is true
func (repo *GitRepository) mkdirGitdirPath(shouldMkdir bool, paths ...string) (string, error) {
	gitDirPath := repo.computeGitdirPath(paths...)

	fi, err := os.Stat(gitDirPath)

	if !os.IsNotExist(err) {
		if !fi.IsDir() {
			return "", fmt.Errorf("Not a directory %s", gitDirPath)
		}
		return gitDirPath, nil
	}

	if shouldMkdir {
		// Everyone can read, write and execute
		err := os.MkdirAll(gitDirPath, 0777)
		if err != nil {
			return "", err
		}
		return gitDirPath, nil
	}

	return "", nil
}

// Same as repo_path, but create dirname(*path) if absent.  For
// example, repo_file(r, \"refs\", \"remotes\", \"origin\", \"HEAD\") will create
// .git/refs/remotes/origin.
func (repo *GitRepository) makeGitdirFile(shouldMkdir bool, paths ...string) (string, error) {
	dirPath := paths[:len(paths)-1]
	gitDirPath, err := repo.mkdirGitdirPath(shouldMkdir, dirPath...)
	if err != nil {
		return "", err
	}

	if gitDirPath != "" {
		return repo.computeGitdirPath(paths...), nil
	}

	return gitDirPath, nil
}

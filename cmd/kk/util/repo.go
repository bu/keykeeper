package util

import (
	"log"
	"os"
	"path"
)

type Repo struct {
	path string
}

// GetPath returns the path of the repository.
func (r *Repo) GetPath() string {
	// if no custom path is set, get the default path
	if r.path == "" {
		// get the user home directory
		userHome, err := os.UserHomeDir()

		// if we cannot get the user home directory, panic, because we rely on it
		if err != nil {
			log.Panicf("Error getting user home directory: %v", err)
		}

		// set the password store path
		r.path = path.Join(userHome, ".password-store")
	}

	// return the path
	return r.path
}

func (r *Repo) IsReady() bool {
	_, err := os.Stat(r.GetPath())

	if os.IsNotExist(err) {
		return false
	}

	return true
}

// NewRepo creates a new repository object with the given path.
func NewRepo(path string) *Repo {
	return &Repo{
		path: path,
	}
}

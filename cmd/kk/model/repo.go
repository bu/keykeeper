package model

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"path/filepath"
)

type Repo struct {
	dirPath string
	config  *Config
}

func (r *Repo) IsReady() bool {
	stat, err := os.Stat(r.dirPath)

	if err != nil {
		return false
	}

	if !stat.IsDir() {
		return false
	}

	return true
}

func (r *Repo) CheckExists(reqPath string) bool {
	_, err := os.Stat(r.GetPath(reqPath))
	if err == nil {
		return true
	}

	_, err = os.Stat(r.GetFilepath(reqPath))
	if err == nil {
		return true
	}

	return false
}

// ExtractFileExtension is a function that extracts the file extension from the file path.
func (r *Repo) ExtractFileExtension(filePath string) string {
	return filepath.Ext(filePath)
}

func (r *Repo) GetPath(reqPath string) string {
	return path.Join(r.dirPath, reqPath)
}

func (r *Repo) GetFilepath(reqPath string) string {
	matches, err := filepath.Glob(r.GetPath(reqPath) + ".*")
	if err != nil {
		return ""
	}

	if len(matches) == 0 {
		return ""
	}

	return matches[0]
}

func NewRepo(dirPath string) (*Repo, error) {
	// if dirPath is empty, then we use .password-store under the user's home directory
	if dirPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		dirPath = path.Join(home, ".password-store")
	}

	// load the config if exists
	config, err := LoadConfig(path.Join(dirPath, ".keykeeper.yml"))
	if err != nil {
		return nil, err
	}

	return &Repo{
		dirPath: dirPath,
		config:  config,
	}, nil
}

func (r *Repo) CheckIsFile(filePath string) bool {
	stat, err := os.Stat(r.GetFilepath(filePath))

	if err != nil {
		return false
	}

	return !stat.IsDir()
}

// GetConfig is a function that returns the config.
func (r *Repo) GetConfig() *Config {
	return r.config

}

// RemoveFileExtension is a function that removes the file extension from the file path.
func (r *Repo) RemoveFileExtension(filePath string) string {
	return filePath[:len(filePath)-len(filepath.Ext(filePath))]
}

func LoadConfig(filePath string) (*Config, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		return &Config{}, nil
	}

	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

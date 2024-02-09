package file

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	flowDirName = "flow"
)

var (
	UserConfigPath       = ConfigDirPath() + "/config.yaml"
	DefaultWorkspacePath = CachedDataDirPath() + "/default"
	LatestCacheDataPath  = CachedDataDirPath() + "/latestcache"
)

func ConfigDirPath() string {
	dirname, err := os.UserConfigDir()
	if err != nil {
		panic(errors.Wrap(err, "unable to get config directory"))
	}
	return filepath.Join(dirname, flowDirName)
}

func CachedDataDirPath() string {
	dirname, err := os.UserCacheDir()
	if err != nil {
		panic(errors.Wrap(err, "unable to get cache directory"))
	}
	return filepath.Join(dirname, flowDirName)
}

func LatestCachedDataFilePath(cacheKey string) string {
	return filepath.Join(LatestCacheDataPath, cacheKey)
}

func EnsureConfigDir() error {
	if _, err := os.Stat(ConfigDirPath()); os.IsNotExist(err) {
		err = os.MkdirAll(ConfigDirPath(), 0750)
		if err != nil {
			return errors.Wrap(err, "unable to create config directory")
		}
	} else if err != nil {
		return errors.Wrap(err, "unable to check for config directory")
	}
	return nil
}

func EnsureCachedDataDir() error {
	if _, err := os.Stat(LatestCacheDataPath); os.IsNotExist(err) {
		err = os.MkdirAll(LatestCacheDataPath, 0750)
		if err != nil {
			return errors.Wrap(err, "unable to create cache directory")
		}
	} else if err != nil {
		return errors.Wrap(err, "unable to check for cache directory")
	}

	return nil
}

func EnsureDefaultWorkspace() error {
	if _, err := os.Stat(DefaultWorkspacePath); os.IsNotExist(err) {
		err = os.MkdirAll(DefaultWorkspacePath, 0750)
		if err != nil {
			return errors.Wrap(err, "unable to create default workspace directory")
		}
	} else if err != nil {
		return errors.Wrap(err, "unable to check for default workspace directory")
	}
	return nil
}

func EnsureWorkspaceDir(workspacePath string) error {
	if _, err := os.Stat(workspacePath); os.IsNotExist(err) {
		err = os.MkdirAll(workspacePath, 0750)
		if err != nil {
			return errors.Wrap(err, "unable to create workspace directory")
		}
	} else if err != nil {
		return errors.Wrap(err, "unable to check for workspace directory")
	}
	return nil
}

func EnsureWorkspaceConfig(workspaceName, workspacePath string) error {
	if _, err := os.Stat(filepath.Join(workspacePath, WorkspaceConfigFileName)); os.IsNotExist(err) {
		return InitWorkspaceConfig(workspaceName, workspacePath)
	} else if err != nil {
		return errors.Wrapf(err, "unable to check for workspace %s config file", workspaceName)
	}
	return nil
}

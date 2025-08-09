package paths

import (
	"os"
	"path/filepath"
	"sync"
)

var (
	dirCache   = make(map[string]string)
	cacheMutex sync.RWMutex
)

// GetDirectoryRoot returns the absolute path to a directory with the given name and caches it.
func GetDirectoryRoot(directoryName string) string {
	cacheMutex.RLock()
	if root, exists := dirCache[directoryName]; exists {
		cacheMutex.RUnlock()
		return root
	}
	cacheMutex.RUnlock()

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if root, exists := dirCache[directoryName]; exists {
		return root
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		targetPath := filepath.Join(dir, directoryName)
		if _, err := os.Stat(targetPath); err == nil {
			dirCache[directoryName] = targetPath
			return targetPath
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			panic("could not find directory named '" + directoryName + "'")
		}
		dir = parent
	}
}

// ResolvePath resolves a path relative to the specified directory.
func ResolvePath(directoryName string, relativePath string) string {
	return filepath.Join(GetDirectoryRoot(directoryName), relativePath)
}

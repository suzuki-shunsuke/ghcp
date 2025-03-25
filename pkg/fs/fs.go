package fs

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(FileSystem)),
	wire.Bind(new(Interface), new(*FileSystem)),
)

//go:generate mockgen -destination mock_fs/mock_fs.go github.com/suzuki-shunsuke/ghcp/pkg/fs Interface

type Interface interface {
	FindFiles(paths []string, filter FindFilesFilter) ([]File, error)
	ReadAsBase64EncodedContent(filename string) (string, error)
	SetDelete(delete bool)
}

// FindFilesFilter is an interface to filter directories and files.
type FindFilesFilter interface {
	SkipDir(path string) bool     // If true, it skips entering the directory
	ExcludeFile(path string) bool // If true, it excludes the file from the result
}

type nullFindFilesFilter struct{}

func (*nullFindFilesFilter) SkipDir(string) bool     { return false }
func (*nullFindFilesFilter) ExcludeFile(string) bool { return false }

type File struct {
	Path       string
	Executable bool
	Deleted    bool
}

// FileSystem provides manipulation of file system.
type FileSystem struct {
	Delete bool
}

func (fs *FileSystem) SetDelete(delete bool) {
	fs.Delete = delete
}

// FindFiles returns a list of files in the paths.
// If the filter is nil, it returns any files.
func (fs *FileSystem) FindFiles(paths []string, filter FindFilesFilter) ([]File, error) {
	if filter == nil {
		filter = &nullFindFilesFilter{}
	}
	var files []File
	for _, path := range paths {
		if fs.Delete {
			if _, err := os.Stat(path); err != nil {
				if !errors.Is(err, os.ErrNotExist) {
					return nil, fmt.Errorf("error while checking if the file exists: %w", err)
				}
				files = append(files, File{
					Path:    path,
					Deleted: true,
				})
				continue
			}
		}
		if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("error while walk: %w", err)
			}
			if info.Mode().IsDir() {
				if filter.SkipDir(path) {
					return filepath.SkipDir
				}
				return nil
			}
			if info.Mode().IsRegular() {
				if filter.ExcludeFile(path) {
					return nil
				}
				files = append(files, File{
					Path:       path,
					Executable: info.Mode()&0o100 != 0, // mask the executable bit of owner
				})
				return nil
			}
			return nil
		}); err != nil {
			return nil, fmt.Errorf("error while finding files in %s: %w", path, err)
		}
	}
	return files, nil
}

// ReadAsBase64EncodedContent returns content of the file as base64 encoded string.
func (fs *FileSystem) ReadAsBase64EncodedContent(filename string) (string, error) {
	r, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error while opening file %s: %w", filename, err)
	}
	defer r.Close() //nolint:errcheck
	var s strings.Builder
	e := base64.NewEncoder(base64.StdEncoding, &s)
	if _, err := io.Copy(e, r); err != nil {
		return "", fmt.Errorf("error while encoding file %s: %w", filename, err)
	}
	if err := e.Close(); err != nil {
		return "", fmt.Errorf("error while encoding file %s: %w", filename, err)
	}
	return s.String(), nil
}

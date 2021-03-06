package xos

import (
	"os"
)

const (
	// ModeDir represents a file mode in *nix is rwx-rx-rx
	ModeDir = os.FileMode(0755)
	// ModeRegularFile represents a file mode in *nix is rw-r-r
	ModeRegularFile = os.FileMode(0644)
)

func TempFile() string {
	return ""
}

// Touch creates an empty file only, returns error when file exist
func Touch(filename string) error {
	if _, err := os.Stat(filename); err == nil {
		return os.ErrExist
	}

	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	return fp.Close()
}

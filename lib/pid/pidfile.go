package pid

import (
	"fmt"
	"os"

	"github.com/eviltomorrow/philosopher-stone/lib/fs"
	"github.com/eviltomorrow/philosopher-stone/lib/runtimeutil"
)

func CreatePidFile(path string) (func() error, error) {
	file, err := fs.CreateFlockFile(path)
	if err != nil {
		return nil, err
	}

	file.WriteString(fmt.Sprintf("%d", runtimeutil.Pid))
	if err := file.Sync(); err != nil {
		file.Close()
		return nil, err
	}

	return func() error {
		if file != nil {
			if err := file.Close(); err != nil {
				return err
			}
			return os.Remove(path)
		}
		return fmt.Errorf("panic: pid file is nil")
	}, nil
}

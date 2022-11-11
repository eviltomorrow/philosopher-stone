package runtimeutil

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/eviltomorrow/philosopher-stone/lib/netutil"
	"github.com/eviltomorrow/philosopher-stone/lib/timeutil"
)

var (
	ExecutableName string
	ExecutableDir  string
	Pid            = os.Getpid()
	LaunchTime     = time.Now()
	HostName       string
	OS             = runtime.GOOS
	Arch           = runtime.GOARCH
	RunningTime    = func() string {
		return timeutil.FormatDuration(time.Since(LaunchTime))
	}
	IP string
)

func init() {
	path, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("panic: get Executable path failure, nest error: %v", err))
	}
	path, err = filepath.Abs(path)
	if err != nil {
		panic(fmt.Errorf("panic: abs RootDir failure, nest error: %v", err))
	}
	ExecutableDir = filepath.Dir(path)
	ExecutableName = filepath.Base(path)

	name, err := os.Hostname()
	if err == nil {
		HostName = name
	}

	localIP, err := netutil.GetLocalIP2()
	if err == nil {
		IP = localIP
	}
}

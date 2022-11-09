package buildinfo

import (
	"fmt"
	"runtime"
)

var (
	AppName     string = "Unknown"
	MainVersion string = "Unknown"
	GoVersion          = runtime.Version()
	GoOSArch           = runtime.GOOS + "/" + runtime.GOARCH
	GitSha      string = "Unknown"
	GitTag      string = "Unknown"
	GitBranch   string = "Unknown"
	BuildTime   string = "Unknown"
)

func GetVersion() string {
	return fmt.Sprintf(`Main Version(%s): %s
Git Sha: %s
Git Tag: %s
Git Branch: %s
Go Version: %s
GO OS/Arch: %s
Build Time: %s`, AppName, MainVersion, GitSha, GitTag, GitBranch, GoVersion, GoOSArch, BuildTime)
}

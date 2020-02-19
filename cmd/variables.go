package cmd

import "fmt"

const (
	appName = "mediawiki-cli"
	appDesc = "handy Mediawiki communities/wiki tool"
	appLongDesc = "More info soon"
)

type Version struct {
	Version string
	Commit string
	Date string
}

func (v Version) String() string {
	return fmt.Sprintf("version: %s (commit: %s), date: %s", v.Version, v.Commit, v.Date)
}

var version = Version{Commit: "dev", Version: "dev", Date: "n/a"}
package cmd

import "fmt"

const (
	appName     = "mediawiki-cli"
	appDesc     = "handy Mediawiki communities/wiki tool"
	appLongDesc = "More info soon"
)

var (
	version string
	commit  string
	date    string
)

func GetVersion() string {
	return fmt.Sprintf("version: %s (commit: %s), date: %s", version, commit, date)
}

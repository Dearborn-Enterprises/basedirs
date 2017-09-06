// Package basedirs provides cache, config and data directories with the simplest API possible.
//
// See https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html
package basedirs

import (
	"os"
	"os/user"
	"path/filepath"
)

var (
	Home string

	UserCache  string
	UserConfig string
	UserData   string
)

func init() {
	Home = os.Getenv("HOME")
	if Home == "" {
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
		Home = usr.HomeDir
	}

	UserCache = os.Getenv("XDG_CACHE_HOME")
	if UserCache == "" || !filepath.IsAbs(UserCache) {
		UserCache = filepath.Join(Home, ".cache")
	}

	UserConfig = os.Getenv("XDG_CONFIG_HOME")
	if UserConfig == "" || !filepath.IsAbs(UserConfig) {
		UserConfig = filepath.Join(Home, ".config")
	}

	UserData = os.Getenv("XDG_DATA_HOME")
	if UserData == "" || !filepath.IsAbs(UserData) {
		UserData = filepath.Join(Home, ".local", "share")
	}
}

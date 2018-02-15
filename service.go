package sourceroostersvc

import (
	"regexp"
	"time"
	"os"
)

type Rooster interface {
	FileExists(path string) bool
	IsProjectDir(dir string) bool
	IsSourceFile(filename string) bool
	IsResourceFile(filename string) bool
	GetLastUpdate(dir string) time.Duration
	HasCVS(dir string) bool
}

type rooster struct {
	sources struct {
		whitelist []string
		blacklist []string
	}
}

func (svc *rooster) FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func (svc *rooster) IsProjectDir(dir string) bool {
	// maven project
	if svc.FileExists(dir + "/pom.xml") && svc.FileExists(dir+"/src/") {
		return true
	}

	// any project
	if svc.FileExists(dir + "/Makefile") {
		return true
	}

	// c++ project
	if svc.FileExists(dir + "/CMakeLists.txt") {
		return true
	}

	// python project
	if svc.FileExists(dir + "/requirements.txt") {
		return true
	}

	// php + composer project
	if svc.FileExists(dir+"/composer.json") &&
		svc.FileExists(dir+"/composer.lock") &&
		svc.FileExists(dir+"/src/") {
		return true
	}

	// android project
	if svc.FileExists(dir+"/build.gradle") &&
		svc.FileExists(dir+"/local.properties") &&
		svc.FileExists(dir+"/settings.gradle") {
		return true
	}

	// golang project
	if svc.FileExists(dir+"/cmd") && svc.FileExists(dir+"/vendor") {
		return true
	}

	return false
}

func (svc *rooster) IsSourceFile(filename string) bool {
	for _, black := range svc.sources.blacklist {
		if ok, _ := regexp.MatchString(black, filename); ok {
			return false
		}
	}
	for _, white := range svc.sources.whitelist {
		if ok, _ := regexp.MatchString(white, filename); ok {
			return true
		}
	}
	return false
}

func (svc *rooster) IsResourceFile(filename string) bool {
	return false
}

func (svc *rooster) GetLastUpdate(dir string) time.Duration {
	return 0
}

func (svc *rooster) HasCVS(dir string) bool {
	return svc.FileExists(dir+"/.git/") || svc.FileExists(dir+"/.svn/")
}

func NewService(srcWhitelist, srcBlacklist []string) Rooster {
	return &rooster{}
}

package sourceroostersvc

import (
	"os"
	"regexp"
	"time"
)

type Rooster interface {
	FileExists(path string) bool
	IsProjectDir(dir string) bool
	DetectProject(dir string) Project
	IsSourceFile(filename string) bool
	IsResourceFile(filename string) bool
	GetLastUpdate(path string) (time.Time, error)
	HasCVS(dir string) bool
}

type rooster struct {
	sources []string
}

func (svc *rooster) FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func (svc *rooster) IsProjectDir(dir string) bool {
	// maven project
	if svc.FileExists(dir+"/pom.xml") && svc.FileExists(dir+"/src/") {
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

func (svc *rooster) DetectProject(dir string) Project {
	return Project{} // TODO: implement method
}

func (svc *rooster) IsSourceFile(filename string) bool {
	for _, white := range svc.sources {
		if ok, _ := regexp.MatchString(white, filename); ok {
			return true
		}
	}
	return false
}

func (svc *rooster) IsResourceFile(filename string) bool {
	return false
}

func (svc *rooster) GetLastUpdate(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

func (svc *rooster) HasCVS(dir string) bool {
	return svc.FileExists(dir+"/.git/") || svc.FileExists(dir+"/.svn/")
}

func NewService(srcWhitelist []string) Rooster {
	return &rooster{
		sources: srcWhitelist,
	}
}

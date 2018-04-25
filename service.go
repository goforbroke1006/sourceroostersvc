package sourceroostersvc

import (
	"os"
	"regexp"
	"strings"
	"time"
	fs "github.com/goforbroke1006/sourceroostersvc/filesystem"
)

type Rooster interface {
	IsProjectDir(dir string) bool
	DetectProject(dir string) Project
	IsSourceFile(filename string) bool
	IsResourceFile(filename string) bool
	GetLastUpdate(path string) (time.Time, error)
	GetCVSRemoteLink(dir string) string
}

type rooster struct {
	sources []string
}



func (svc *rooster) IsProjectDir(dir string) bool {
	// maven project
	if fs.FileExists(dir+"/pom.xml") && fs.FileExists(dir+"/src/") {
		return true
	}

	// any project
	if fs.FileExists(dir + "/Makefile") {
		return true
	}

	// c++ project
	if fs.FileExists(dir + "/CMakeLists.txt") {
		return true
	}

	// python project
	if fs.FileExists(dir + "/requirements.txt") {
		return true
	}

	// php + composer project
	if fs.FileExists(dir+"/composer.json") &&
		fs.FileExists(dir+"/composer.lock") &&
		fs.FileExists(dir+"/src/") {
		return true
	}

	// android project
	if fs.FileExists(dir+"/build.gradle") &&
		fs.FileExists(dir+"/local.properties") &&
		fs.FileExists(dir+"/settings.gradle") {
		return true
	}

	// golang project
	if fs.FileExists(dir+"/cmd") && fs.FileExists(dir+"/vendor") {
		return true
	}

	return false
}

func (svc *rooster) DetectProject(dir string) Project {
	return Project{
		Name: fs.GetDirSimpleName(dir),
		Path: dir,
	} // TODO: implement method
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

func (svc *rooster) GetCVSRemoteLinks(dir string) map[string]string {
	if fs.FileExists(dir+"/.git/") {
		return map[string]string {}
	}
	if fs.FileExists(dir+"/.svn/") {
		return map[string]string {}
	}
	return map[string]string {} // TODO: implement this method
}

func NewService(srcWhitelist []string) Rooster {
	return &rooster{
		sources: srcWhitelist,
	}
}

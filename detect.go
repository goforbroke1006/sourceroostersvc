package sourceroostersvc

import (
	"os"
	"path/filepath"
)

func IsProjectDir(dir string) bool {
	filepath.Dir(dir)

	// maven project
	if fileExists(dir + "/pom.xml") {
		return true
	}

	// any project
	if fileExists(dir + "/Makefile") {
		return true
	}

	// c++ project
	if fileExists(dir + "/CMakeLists.txt") {
		return true
	}

	// python project
	if fileExists(dir + "/requirements.txt") {
		return true
	}

	// php + composer project
	if fileExists(dir+"/composer.json") &&
		fileExists(dir+"/composer.lock") &&
		fileExists(dir+"/src/") {
		return true
	}

	// android project
	if fileExists(dir+"/build.gradle") &&
		fileExists(dir+"/local.properties") &&
		fileExists(dir+"/settings.gradle") {
		return true
	}

	// golang project
	if fileExists(dir+"/cmd") && fileExists(dir+"/vendor") {
		return true
	}

	return false
}

func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

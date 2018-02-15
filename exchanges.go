package sourceroostersvc

import "fmt"

type Project struct {
	Name    string
	Path    string
	CVSLink string
}

func (p Project) ToString() string {
	cvs := p.CVSLink
	if len(cvs) == 0 {
		cvs = "None"
	}
	return fmt.Sprint("[%s] %s", p.Name, cvs)
}

func NewProject(name, path, cvs string) Project {
	return Project{
		Name:    name,
		Path:    path,
		CVSLink: cvs,
	}
}

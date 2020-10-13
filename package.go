package plantuml

import (
	"io"
)

type Package struct {
	name     string
	children []Renderable
}

func NewPackage(name string) *Package {
	return &Package{name: name}
}

func (p *Package) Add(r ...Renderable) *Package {
	p.children = append(p.children, r...)
	return p
}

func (p *Package) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}

	//name package "ER model" to use the correct settings
	if p.name == "ER model" {
		w.Print("skinparam linetype ortho \n")
	}

	w.Print("package \"")
	w.Print(escapeP(p.name))
	w.Print("\" {\n")
	for _, child := range p.children {
		if err := child.Render(wr); err != nil {
			return err
		}
	}
	w.Print("}\n")

	return w.Err
}

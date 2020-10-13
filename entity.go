package plantuml

import (
	"io"
)


type RelationTo string
type RelationFr string


const (
	Bold		string 		= "**"
	Manda		string	 	= "* "

	FZeOn		RelationFr		= "|o--"
	FExOn		RelationFr		= "||--"
	FZeMa		RelationFr		= "}o--"
	FOnMa		RelationFr		= "}|--"

	TZeOn		RelationTo 		= "o|"
	TExOn		RelationTo		= "||"
	TZeMa		RelationTo		= "o{"
	TOnMa		RelationTo		= "|{"
)

type Entity struct {
	name	string
	alias	string
	attrs	[]Attribute
	conns	[]Connection
}

type Attribute struct {
	Name 	string
	DType 	string
	Prime	bool
	Manda 	bool
	Bold	bool
}

type Connection struct {
	Start 	string
	Dest 	string
	FType	RelationFr
	TType 	RelationTo
}

type CList struct {
	list	[]Connection
}


func NewEntity(name string, alias string) *Entity {
	return &Entity{
		name:   name,
		alias:	alias,
	}
}

func NewConnectionList() *CList {
	return &CList{}
}

func (cl *CList) AddConns(conns ...Connection) *CList {
	cl.list = append(cl.list, conns...)
	return cl
}

func (e *Entity) AddAttrs(attrs ...Attribute) *Entity {
	e.attrs = append(e.attrs, attrs...)
	return e
}

func (c *Connection) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Print(c.Start)
	w.Print(string(c.FType))
	w.Print(string(c.TType))
	w.Print(c.Dest + "\n")

	return w.Err
}

func (cl *CList) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}

	for _, conn := range cl.list {
		w.Print(conn.Start)
		w.Print(string(conn.FType))
		w.Print(string(conn.TType))
		w.Print(conn.Dest)
		w.Print("\n")
	}
	return w.Err
}

func (e *Entity) Render(wr io.Writer) error {
	w := strWriter{Writer: wr}
	w.Print("entity \"" + escapeP(e.name) + "\" as " + e.alias)
	w.Print(" {")
	w.Print("\n")

	for _, attr := range e.attrs {
		if attr.Manda {
			w.Print(Manda)
		}

		if attr.Bold {
			w.Print(Bold)
		}

		w.Print(attr.Name)
		w.Print(" : ")

		if len(attr.DType) > 40 {
			desc := SplitDesc(attr.DType)
			for _, v := range desc {
				w.Printf("%v \n", v )
			}
		} else {
			w.Print(attr.DType)
		}

		if attr.Bold {
			w.Print(Bold)
		}

		w.Print("\n")

		if attr.Prime {
			w.Print("--")
			w.Print("\n")
		}
	}

	w.Print("}")
	w.Print("\n")


	return w.Err
}

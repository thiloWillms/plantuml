package plantuml

import "io"




const (
	Bold		string 		= "**"
	Manda		string	 	= "* "

	FZeOn		string 		= "|o--"
	FExOn		string		= "||--"
	FZeMa		string		= "}o--"
	FOnMa		string		= "}|--"

	TZeOn		string 		= "o|"
	TExOn		string		= "||"
	TZeMa		string		= "o{"
	TOnMa		string		= "|{"
)

type Entity struct {
	name	string
	alias	string
	attrs	[]Attribute
	Conn	Connection
}

type Attribute struct {
	Name 	string
	DType 	string
	Prime	bool
	Manda 	bool
	Bold	bool
}

type Connection struct {
	fType	string
	tType 	string
	dest 	string
}

type []Connections struct {
listConn	[]Connection
}

func NewEntity(name string, alias string) *Entity {
	return &Entity{
		name:   name,
		alias:	alias,
	}
}

func (d *Entity) AddAttrs(attrs ...Attribute) *Entity {
	d.attrs = append(d.attrs, attrs...)
	return d
}

func (d *Entity) Render(wr io.Writer) error {
	var connections *[]Connections
	w := strWriter{Writer: wr}
	w.Print("entity \"" + escapeP(d.name) + "\" as " + d.alias)
	w.Print(" {")
	w.Print("\n")

	for _, attr := range d.attrs {
		if attr.Manda {
			w.Print(Manda)
		}

		if attr.Bold {
			w.Print(Bold)
		}

		w.Print(attr.Name)
		w.Print(" : ")
		w.Print(attr.DType)

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

	if d.Conn.dest != "" {
		connections = append(connections.listConn, d.Conn)
	}
	return w.Err
}

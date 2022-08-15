package models

type Parse struct {
	Url        string       // Url to parse
	HeaderSets []*HeaderSet // Header sets to parse
	Selection  Selection    // Selection to parse
}

type HeaderSet struct {
	Key   string // Key of header set
	Value string // Value of header set
}

type Selection struct {
	Find []*Find
}

type Find struct {
	Tag      string
	Class    string
	GetValue bool
	GetAttr  *GetAttr
	Find     []*Find
}

type GetAttr struct {
	Attr string
}

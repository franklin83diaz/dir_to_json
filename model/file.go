package model

type FileStructure struct {
	Name     string          `json:"name"`
	IsDir    bool            `json:"isDir"`
	Children []FileStructure `json:"children,omitempty"`
	Size     int64           `json:"size,omitempty"`
	Path     string          `json:"path"`
}

type FileExplore struct {
	Path   string
	Level  int
	Hidden bool
}

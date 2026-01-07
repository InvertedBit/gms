package models

type Media struct {
	Model
	FileName string
	FileType string
	Path     string
}

func (m *Media) GetURL() string {
	return "/media/" + m.ID.String() + "." + m.FileType
}

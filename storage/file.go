package storage

type File struct {
	Name     string
	Content  []byte
	Metadata map[string]string
}

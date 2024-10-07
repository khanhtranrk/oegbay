package operation

type Operation interface {
	MakeDir(load string, path string) error
	RemoveDir(load string, path string) error
	ReadFile(load string, path string) ([]byte, error)
	WriteFile(load string, path string, content []byte) error
	DeleteFile(load string, path string) error
}

package file

type Handler interface {
	GetFileExtension() string
	Decrypt([]byte) ([]byte, error)
	Encrypt([]byte) ([]byte, error)
}

package file

import "keykeeper/cmd/kk/model"

type NaClHandler struct {
	repo *model.Repo
}

func (h *NaClHandler) GetFileExtension() string {
	return ".nacl"
}

func (h *NaClHandler) Decrypt(data []byte) ([]byte, error) {
	return nil, nil
}

func (h *NaClHandler) Encrypt(data []byte) ([]byte, error) {
	return nil, nil
}

func NewNaClHandler(repo *model.Repo) *NaClHandler {
	return &NaClHandler{
		repo: repo,
	}
}

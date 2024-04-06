package file

import (
	"bytes"
	"fmt"
	"github.com/ProtonMail/go-crypto/openpgp"
	"golang.org/x/term"
	"io"
	"keykeeper/cmd/kk/model"
	"log"
	"os"
)

type OpenPgpHandler struct {
	repo *model.Repo
}

func NewPgpHandler(repo *model.Repo) *OpenPgpHandler {
	return &OpenPgpHandler{
		repo: repo,
	}
}

func (h *OpenPgpHandler) GetFileExtension() string {
	return ".gpg"
}

func (h *OpenPgpHandler) Decrypt(data []byte) ([]byte, error) {
	entityList, err := openpgp.ReadArmoredKeyRing(bytes.NewReader([]byte(h.repo.GetConfig().GpgPrivateKey)))
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	prompt := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		if symmetric {
			return nil, fmt.Errorf("prompt: message was marked as symmetrically encrypted")
		}
		if len(keys) == 0 {
			return nil, fmt.Errorf("prompt: no keys requested")
		}
		fmt.Print("\033[H\033[2J\033[3J")
		fmt.Print("Password for PGP Key (will not echo) : ")
		pass, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return nil, fmt.Errorf("prompt: error reading password: %s", err)
		}

		err = keys[0].PrivateKey.Decrypt(pass)
		if err != nil {
			return nil, fmt.Errorf("prompt: error decrypting key: %s", err)
		}

		return nil, nil
	}

	md, err := openpgp.ReadMessage(bytes.NewReader(data), entityList, prompt, nil)
	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, fmt.Errorf("error reading UnverifiedBody: %s", err)
	}

	return contents, nil
}

func (h *OpenPgpHandler) Encrypt(data []byte) ([]byte, error) {
	return nil, nil
}

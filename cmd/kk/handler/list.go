package handler

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/term"
	"keykeeper/cmd/kk/lang"
	"keykeeper/cmd/kk/model"
	"log"
	"os"
	"path"
	"strings"
)

func init() {
	prepareRepo()
	// register the handler
	registeredHandlers["list"] = &ListHandler{}
	// register the aliases
	registeredAliases["ls"] = &ListHandler{}
	registeredAliases["show"] = &ListHandler{}
	registeredAliases["cat"] = &ListHandler{}
	registeredAliases["get"] = &ListHandler{}
	registeredAliases["view"] = &ListHandler{}

	// register colors
	DirLevelColor = []colorStringFunc{
		color.BlueString,
		color.MagentaString,
		color.BlueString,
		color.YellowString,
		color.GreenString,
		color.RedString,
	}
}

var DirLevelColor []colorStringFunc

type colorStringFunc func(string, ...any) string

type ListHandler struct {
}

func (h *ListHandler) Handle(env *HandlerEnv) error {
	// list handler only process 1 argument
	if len(env.Args) != 1 {
		return fmt.Errorf("%s", lang.GetString("error.list.invalid_arguments", len(env.Args)))
	}

	// extract path
	requestPath := env.Args[0]

	// handle flags
	flags := model.ParseListFlag(env.Flags)

	// no color
	if flags.NoColor {
		color.NoColor = true
	}

	// if path is file, then we should show the file
	if env.Repo.CheckExists(requestPath) && env.Repo.CheckIsFile(requestPath) {
		return h.showFile(env, env.Repo.GetFilepath(requestPath), flags)
	}

	// start processing the list command
	headerFontStyle := color.New(color.FgCyan, color.Bold)

	if !flags.List && !flags.NoHeader && requestPath == "." {
		_, err := headerFontStyle.Println(lang.GetString("password_store"))
		if err != nil {
			return err
		}
	}

	if !flags.List && !flags.NoHeader && requestPath != "." {
		_, err := headerFontStyle.Println(requestPath)
		if err != nil {
			return err
		}
	}

	// list items
	return h.listItems(env, requestPath, flags.GetIndentCharacter(), flags, 0)
}

func (h *ListHandler) listItems(env *HandlerEnv, dirPath string, prefix string, flags *model.ListFlag, level uint64) error {
	items, err := os.ReadDir(env.Repo.GetPath(dirPath))
	if err != nil {
		return err
	}

	for _, item := range items {
		// skip hidden files
		if strings.HasPrefix(item.Name(), ".") {
			continue
		}

		// if current item is a directory, then we should list it recursively
		if item.IsDir() {
			if level >= uint64(len(DirLevelColor)) {
				level = level % uint64(len(DirLevelColor))
			}

			if flags.List {
				if !flags.ListFileOnly {
					fmt.Println(DirLevelColor[level](path.Join(dirPath, item.Name())))
				}
			} else {
				fmt.Println(DirLevelColor[level](prefix + item.Name()))
			}

			if flags.NoRecursive {
				continue
			}

			// recursively list the items
			err := h.listItems(
				env,
				path.Join(dirPath, item.Name()),
				prefix+flags.GetIndentCharacter(),
				flags,
				level+1,
			)

			if err != nil {
				return err
			}

			continue
		}

		// if this item is a file, then
		fileName := env.Repo.RemoveFileExtension(item.Name())

		if flags.List {
			fmt.Println(path.Join(dirPath, fileName))
		} else {
			fmt.Println(prefix + fileName)
		}
	}

	return nil
}

func (h *ListHandler) showFile(env *HandlerEnv, path string, _ *model.ListFlag) error {
	log.Println("showing file", path)

	fileExt := env.Repo.ExtractFileExtension(path)
	handler, ok := env.FileHandler[fileExt]

	if !ok {
		return fmt.Errorf("%s", lang.GetString("error.file_handler_not_found", fileExt))
	}

	// read the file
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// decrypt the file
	decryptedContent, err := handler.Decrypt(content)
	if err != nil {
		return err
	}

	// print the content
	fmt.Print("\033[H\033[2J\033[3J")
	fmt.Print(string(decryptedContent))
	fmt.Print("\n")

	return nil
}

func consolePrompot() {
	_, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	//defer term.Restore(int(os.Stdin.Fd()), oldState)
}

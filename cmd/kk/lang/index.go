package lang

import (
	"embed"
	"fmt"
	"strings"
)
import "gopkg.in/yaml.v3"

//go:embed *.yml
var packs embed.FS

// strings
var currentLang LangPack

func ConfigureLanguage(langs ...string) error {
	// language to load, default to english
	lang := "en_us"

	// if there are any languages provided, use the first not empty one
	for _, l := range langs {
		if l != "" {
			lang = l
			break
		}
	}

	// lower case
	lang = strings.ToLower(lang)

	// replace .utf-8 with .yml
	lang = strings.Replace(lang, ".utf-8", "", 1)
	lang = strings.Replace(lang, "-", "_", 2)

	// Check if the language file exists
	language, err := packs.ReadFile(lang + ".yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(language, &currentLang)
	if err != nil {
		return err
	}

	return nil
}

type LangInfo struct {
	Code   string `yaml:"Code"`
	Name   string `yaml:"Name"`
	Author string `yaml:"Author"`
}

type LangPack struct {
	Lang    LangInfo          `yaml:"Lang"`
	Strings map[string]string `yaml:"Strings"`
}

func GetLanguageCode() string {
	return currentLang.Lang.Code
}

func GetLanguageName() string {
	return currentLang.Lang.Name
}

func GetLanguageAuthor() string {
	return currentLang.Lang.Author
}

func GetString(key string, replaces ...any) string {
	str, ok := currentLang.Strings[key]

	if !ok {
		return key
	}

	return fmt.Sprintf(str, replaces...)
}

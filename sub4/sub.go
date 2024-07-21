package sub4

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed resource/locale.*.json
var localeFS embed.FS

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	if _, err := bundle.LoadMessageFileFS(localeFS, "resource/locale.en.json"); err != nil {
		panic(err)
	}
	if _, err := bundle.LoadMessageFileFS(localeFS, "resource/locale.ja.json"); err != nil {
		panic(err)
	}
}

func Print(lang string) {
	localizer := i18n.NewLocalizer(bundle, lang)
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "MSG1",
	}))
}

package sub4

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/bellwood4486/sample-go-embed/sub4/resource"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed resource/locale.*.json
var localeFS embed.FS

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	// ファイル名は以下の形式を満たす必要がある。
	// `{任意}.{言語}.{マーシャルする種類}`
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
		MessageID: resource.MSG1,
	}))
}

func Print2(lang string, name string) {
	localizer := i18n.NewLocalizer(bundle, lang)
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    resource.MSG2,
		TemplateData: map[string]any{"Name": name},
	}))
}

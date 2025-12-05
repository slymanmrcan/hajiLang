package runtime

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/slymanmrcan/hajilang/object"
)

var ScriptBaseDir string

func RegisterJSON(builtinEnv *object.Environment) {
	// json_read("data.json") -> Dosya içeriğini string olarak döner
	builtinEnv.Set("json_read", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "wrong number of args"}
			}
			path := args[0].(*object.String).Value

			// Eğer mutlak yol değilse, script dizinine göre yorumla
			if !filepath.IsAbs(path) && ScriptBaseDir != "" {
				path = filepath.Join(ScriptBaseDir, path)
			}

			data, err := os.ReadFile(path)
			if err != nil {
				return &object.Error{Message: err.Error()}
			}
			return &object.String{Value: string(data)}
		},
	})

	// json_write("data.json", "içerik") -> Dosyaya yazar
	builtinEnv.Set("json_write", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "wrong number of args"}
			}
			path := args[0].(*object.String).Value
			content := args[1].(*object.String).Value
			err := os.WriteFile(path, []byte(content), 0644)
			if err != nil {
				return &object.Error{Message: err.Error()}
			}
			return object.NULL
		},
	})

	// json_encode(hash) -> Hash'i JSON stringine çevirir
	builtinEnv.Set("json_encode", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Basit string birleştirme yerine Go'nun json kütüphanesini kullanmak
			// için Hash nesnesini geri Go Map'ine çeviren bir fonksiyon lazım.
			// Şimdilik string olarak basit dönelim veya Inspect kullanalım:
			return &object.String{Value: args[0].Inspect()}
		},
	})

	// json_decode(string) -> String'i Hash nesnesine çevirir
	builtinEnv.Set("json_decode", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			raw := args[0].(*object.String).Value
			var m interface{} // Map veya Array olabilir
			json.Unmarshal([]byte(raw), &m)

			// object/object.go içine eklediğimiz fonksiyonu kullanıyoruz
			return object.GoToHash(m)
		},
	})
}

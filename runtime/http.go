package runtime

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/slymanmrcan/hajilang/evaluator"
	"github.com/slymanmrcan/hajilang/lexer"
	"github.com/slymanmrcan/hajilang/object"
	"github.com/slymanmrcan/hajilang/parser"
)

// RunServer: Sunucuyu başlatan ana fonksiyon
func RunServer(scriptPath string) {

	// 1. Favicon hatasını sustur (Boş cevap dön)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	// 2. Tüm istekleri yakala
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		env := object.NewEnvironment()
		RegisterJSON(env)
		RegisterUtils(env)

		// İstek bilgilerini al
		method := r.Method
		bodyBytes, _ := io.ReadAll(r.Body)
		bodyStr := string(bodyBytes)

		// Basit ID yakalama (/api/posts/5 gibi)
		pathParts := strings.Split(r.URL.Path, "/")
		id := ""
		if len(pathParts) > 1 {
			// URL'nin son parçasını ID varsayalım
			id = pathParts[len(pathParts)-1]
		}

		// CTX hash'ini oluştur
		ctxMap := make(map[object.HashKey]object.HashPair)
		addToHash(ctxMap, "method", method)
		addToHash(ctxMap, "body", bodyStr)
		addToHash(ctxMap, "path", r.URL.Path)
		addToHash(ctxMap, "id", id)

		env.Set("CTX", &object.Hash{Pairs: ctxMap})

		// DİNAMİK DOSYA OKUMA
		scriptBytes, err := os.ReadFile(scriptPath)
		if err != nil {
			errMsg := fmt.Sprintf("Server dosyası bulunamadı: %s", scriptPath)
			fmt.Println(errMsg)
			http.Error(w, errMsg, 500)
			return
		}

		// Script'i parse et
		l := lexer.New(string(scriptBytes))
		p := parser.New(l)
		program := p.ParseProgram()

		// Parser hataları varsa
		if len(p.Errors()) > 0 {
			msg := strings.Join(p.Errors(), "\n")
			http.Error(w, "Haji Script Hatası:\n"+msg, 500)
			return
		}

		// Script'i çalıştır
		evaluated := evaluator.Eval(program, env)

		// Runtime hatası varsa
		if evaluated != nil && evaluated.Type() == object.ERROR_OBJ {
			errMsg := fmt.Sprintf("Script Çalışma Hatası: %s", evaluated.Inspect())
			fmt.Println("❌ " + errMsg)
			http.Error(w, errMsg, 500)
			return
		}

		// Cevabı gönder
		sendResponse(w, env)
	})

	fmt.Printf("🚀 Server çalışıyor. Dosya: %s\n", scriptPath)
	fmt.Println("👉 http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// sendResponse: Script'in oluşturduğu 'response' değişkenini okur ve HTTP cevabı verir
func sendResponse(w http.ResponseWriter, env *object.Environment) {
	obj, ok := env.Get("response")
	if !ok || obj == nil {
		fmt.Fprint(w, "Script çalıştı ama 'response' değişkeni tanımlanmadı.")
		return
	}

	// response bir Hash ise (status ve body içerir)
	if hash, ok := obj.(*object.Hash); ok {
		// Status Kodu
		statusKey := &object.String{Value: "status"}
		if pair, ok := hash.Pairs[statusKey.HashKey()]; ok {
			if intVal, ok := pair.Value.(*object.Integer); ok {
				w.WriteHeader(int(intVal.Value))
			}
		} else {
			w.WriteHeader(200)
		}

		// Body
		bodyKey := &object.String{Value: "body"}
		if pair, ok := hash.Pairs[bodyKey.HashKey()]; ok {
			// Eğer body string ise tırnakları temizle
			if strVal, ok := pair.Value.(*object.String); ok {
				fmt.Fprint(w, strVal.Value)
			} else {
				fmt.Fprint(w, pair.Value.Inspect())
			}
		}
	} else {
		// Sadece string döndüyse
		fmt.Fprint(w, obj.Inspect())
	}
}

// addToHash: Hash map'e string eklemek için yardımcı fonksiyon
func addToHash(m map[object.HashKey]object.HashPair, key string, val string) {
	k := &object.String{Value: key}
	v := &object.String{Value: val}
	m[k.HashKey()] = object.HashPair{Key: k, Value: v}
}

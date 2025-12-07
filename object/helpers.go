package object

// GoToHash - JSON'dan (Go interface{}) gelen veriyi HajiLang nesnesine çevirir
// API response'ları için kullanılır
func GoToHash(val interface{}) Object {
	switch v := val.(type) {
	case string:
		return &String{Value: v}
	case float64:
		// JSON'da sayılar float gelir, biz Integer'a çeviriyoruz
		return &Integer{Value: int64(v)}
	case int:
		return &Integer{Value: int64(v)}
	case bool:
		if v {
			return TRUE
		}
		return FALSE
	case map[string]interface{}:
		// JSON Objesini (Map) -> Haji Hash'ine çevir
		pairs := make(map[HashKey]HashPair)
		for k, val := range v {
			key := &String{Value: k}
			value := GoToHash(val) // İç içe (Recursive)
			pairs[key.HashKey()] = HashPair{Key: key, Value: value}
		}
		return &Hash{Pairs: pairs}
	case []interface{}:
		// JSON Listesini (Array) -> Haji Array'ine çevir
		elements := []Object{}
		for _, i := range v {
			elements = append(elements, GoToHash(i))
		}
		return &Array{Elements: elements}
	case nil:
		return NULL
	}
	return NULL
}

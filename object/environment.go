package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	im := make(map[string]bool) // immutable map
	return &Environment{store: s, immutable: im}
}

type Environment struct {
	store     map[string]Object
	immutable map[string]bool
	outer     *Environment
}

// Değişkeni getir (Okuma)
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Değişkeni kaydet (Yazma)
// Eğer değişken outer scope'ta varsa, orada güncelle
func (e *Environment) Set(name string, val Object) Object {
	// Önce bu scope'ta veya outer scope'ta var mı bak
	if _, ok := e.store[name]; ok {
		// Bu scope'ta var, burada güncelle
		e.store[name] = val
		return val
	}
	// Bu scope'ta yok, outer'da bakarak güncelle
	if e.outer != nil {
		if _, ok := e.outer.Get(name); ok {
			return e.outer.Set(name, val)
		}
	}
	// Hiçbir yerde yok, bu scope'ta yeni oluştur
	e.store[name] = val
	return val
}

// SetConst - Sabit değişken kaydet
func (e *Environment) SetConst(name string, val Object) Object {
	e.store[name] = val
	e.immutable[name] = true // ← Bu sabit!
	return val
}

// IsConst - Bu değişken sabit mi?
func (e *Environment) IsConst(name string) bool {
	return e.immutable[name]
}
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

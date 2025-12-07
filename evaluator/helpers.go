package evaluator

import (
	"fmt"

	"github.com/slymanmrcan/hajilang/object"
)

// nativeBoolToBooleanObject - Go bool'unu object.Boolean'a çevirir
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

// newError - Yeni error object oluşturur
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// isError - Object bir error mı kontrol eder
func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

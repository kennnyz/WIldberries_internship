package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

//var justString string // Избавляемся от глобальной переменной,
//глобальные переменные в пакетах, по сути, являются синглтонами,
//использующимися для неявного изменения состояний между слабо не очень связанными вещами,
//создавая прочную зависимость и делая код сложным для тестирования.

func createHugeString(i int) string {
	return strings.Repeat("o", i)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	// Переводим в срез рун, так как он безопасен в плане разделения строки на символы, а не на байты
	vr := []rune(v)
	return string(vr[:100])
}
func main() {
	justString := someFunc()
	fmt.Println(justString, len([]rune(justString)))
}

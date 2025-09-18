package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

// === PRINTING: fmt vs println ===
// Uso recomendado (2025):
// - Para salida en código real, diagnóstico y formateo, usar `fmt` (fmt.Println,
//   fmt.Printf) o un paquete de logging. `fmt` ofrece formato consistente y
//   opciones de control (printf verbs, precisión, etc.).
// - `println` es una función builtin útil para depuración rápida en ejemplos,
//   pero no se debe usar en producción porque su formato no está estandarizado.

// En go nil es nulo

func main() {
	// Integer signed: tipo `int` (implementación dependiente de la plataforma,
	// típicamente 32 o 64 bits). Pueden ser positivos o negativos.
	// Asignación explícita con tipo: `var name type = value`.
	var myIntVar int = -12
	fmt.Println("Mi variable int es:", myIntVar)

	// Unsigned integer: `uint` no permite valores negativos. Útil cuando se
	// necesita representar tamaños o contadores no negativos.
	var myUintVar uint = 12
	fmt.Println("Mi variable uint es:", myUintVar)

	// Floating point: `float64` para mayor precisión. También existe `float32`.
	var myFloatVar float64 = 3.14
	fmt.Println("Mi variable float64 es:", myFloatVar)

	// Complex: números complejos, p. ej. `complex128` (doble precisión).
	var myComplexVar complex128 = 1 + 2i
	fmt.Println("Mi variable complex128 es:", myComplexVar)

	// Tipos enteros con tamaño explícito: int8, uint8, int16, uint16, etc.
	// Se usan cuando se necesita controlar el tamaño (p. ej. binario, memoria).
	var myInt8Var int8 = 127
	fmt.Println("Mi variable int8 es:", myInt8Var)

	var myUint8Var uint8 = 255
	fmt.Println("Mi variable uint8 es:", myUint8Var)

	var myInt16Var int16 = 32767
	fmt.Println("Mi variable int16 es:", myInt16Var)

	var myUint16Var uint16 = 65535
	fmt.Println("Mi variable uint16 es:", myUint16Var)

	var myInt32Var int32 = 2147483647
	fmt.Println("Mi variable int32 es:", myInt32Var)

	var myUint32Var uint32 = 4294967295
	fmt.Println("Mi variable uint32 es:", myUint32Var)

	var myInt64Var int64 = 9223372036854775807
	fmt.Println("Mi variable int64 es:", myInt64Var)

	// Cuidado: literales muy grandes deben caber en el tipo asignado.
	var myUint64Var uint64 = 18446744073709551615
	fmt.Println("Mi variable uint64 es:", myUint64Var)

	var myFloat32Var float32 = 3.14
	fmt.Println("Mi variable float32 es:", myFloat32Var)

	var myFloat64Var float64 = 3.141592653589793
	fmt.Println("Mi variable float64 es:", myFloat64Var)

	// Strings: inmutables, secuencias de bytes UTF-8.
	var myStringVar string = "Hello, Go!"
	// Imprimir string con mensaje descriptivo
	fmt.Println("Mi variable string es:", myStringVar)

	// Booleanos: true/false
	var myBoolVar bool = true
	fmt.Println("Mi variable bool es:", myBoolVar)

	// Declaración corta (:=): inferencia de tipo, solo dentro de funciones.
	// Forma corta y habitual para variables locales.
	myShortFloatVar := 3.14
	fmt.Println("Mi variable float corta es:", myShortFloatVar)

	// Declaración múltiple: permite definir varias variables del mismo tipo.
	var a, b, c int = 1, 2, 3
	fmt.Println("Variables múltiples (a,b,c):", a, b, c)

	// Declaración múltiple con inferencia (forma corta).
	x, y := "foo", "bar"
	fmt.Println("Variables múltiples con inferencia (x,y):", x, y)

	// Constantes: inmutables y evaluadas en tiempo de compilación.
	const myConstVar string = "I am constant"
	fmt.Println("Mi constante string es:", myConstVar)

	// Constantes numéricas sin tipo explícito pueden ser usadas con flexibilidad
	// (numeric untyped constants). Se aconseja dar nombre y usar const
	const pi = 3.14159
	fmt.Println("Valor de pi:", pi)

	// Zero values: variables declaradas sin inicializador reciben el valor
	// por defecto (0 para números, "" para strings, false para bools).
	var zeroInt int
	var zeroString string
	var zeroBool bool
	fmt.Println("Zero values (int, string, bool):", zeroInt, zeroString, zeroBool)

	// Direcciones de memoria: usar &variable para obtener puntero.
	// Los punteros en Go se usan con seguridad; para impresión se muestra
	// la dirección de memoria. Para manipulación, declarar variables pointer.
	fmt.Println("Dirección de memoria de myIntVar:", &myIntVar)
	fmt.Println("Dirección de memoria de myFloatVar:", &myFloatVar)
	fmt.Println("Dirección de memoria de myStringVar:", &myStringVar)
	fmt.Println("Dirección de memoria de myBoolVar:", &myBoolVar)

	// Verbos de formato comunes (fmt.Printf):
	//  %d - enteros en base 10
	//  %b, %o, %x - binario/oct/hex
	//  %f - float con precisión (ej: %.2f dos decimales)
	//  %s - string
	//  %t - booleano
	//  %v - valor por defecto (útil para debugging)
	//  %T - muestra el tipo de la variable
	//  \n salto de línea, \t tabulación
	// Recomendación: usar fmt.Printf cuando necesites controlar formato y
	// precisión; usar %v/%T para inspección rápida.

	fmt.Printf("Formateado - Int (default): %d\n", myIntVar)
	fmt.Printf("Formateado - Uint (default): %d\n", myUintVar)
	fmt.Printf("Formateado - Float (default, 2 decimales): %.2f\n", myFloatVar)
	// Imprimimos complex separando parte real e imaginaria con formato
	fmt.Printf("Formateado - Complex: %.2f + %.2fi\n", real(myComplexVar), imag(myComplexVar))

	// Imprimir tipos con tamaño explícito usando fmt.Printf (con etiquetas)
	fmt.Printf("Formateado - Int8: %d\n", myInt8Var)
	fmt.Printf("Formateado - Uint8: %d\n", myUint8Var)
	fmt.Printf("Formateado - Int16: %d\n", myInt16Var)
	fmt.Printf("Formateado - Uint16: %d\n", myUint16Var)
	fmt.Printf("Formateado - Int32: %d\n", myInt32Var)
	fmt.Printf("Formateado - Uint32: %d\n", myUint32Var)
	fmt.Printf("Formateado - Int64: %d\n", myInt64Var)
	fmt.Printf("Formateado - Uint64: %d\n", myUint64Var)
	fmt.Printf("Formateado - Float32: %.2f\n", myFloat32Var)
	fmt.Printf("Formateado - Float64: %.6f\n", myFloat64Var)
	fmt.Printf("Formateado - String: %s\n", myStringVar)
	fmt.Printf("Formateado - Bool: %t\n", myBoolVar)
	fmt.Printf("Formateado - \n\tValor por defecto: %v\n\tTipo de dato: %T\n", myStringVar, myStringVar)

	// Uso de unsafe
	fmt.Printf("Uso de unsafe - \n\tTipo: %T, \n\tBytes: %d, \n\tBits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	// Uso de `` para strings literales sin escape
	rawString := `This is a raw string literal.
It can span multiple lines.
Special characters like \n and \t are not interpreted.`
	fmt.Println("El rawString es:\n", rawString)

	// Ejemplo de formateo en un scope diferente
	{
		fmt.Println()

		// Flotante a string
		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		// Entero a string
		intVar := 22
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		// De string a entero
		intVar1, err := strconv.ParseInt("1333", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVar1, intVar1)

		// De string a entero
		intVar2, err := strconv.ParseInt("aa1333", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVar2, intVar2)

		floatVar1, _ := strconv.ParseFloat("33.11", 64)
		fmt.Printf("type: %T, value: %f\n", floatVar1, floatVar1)

		// Alternativa simple cuando esperas decimal y un int:
		v3, err := strconv.Atoi("1333")
		fmt.Println("Atoi:", v3, err)

		// Tipo de dato byte
		var A byte = 'A'
		var a byte = 'a'
		fmt.Println("A:", A, "a:", a)

		var R byte = 82
		var s byte = 115
		fmt.Println("R:", R, "s:", s)

		fmt.Println(string(R), string(s))

		vector := []byte{65, 66, 67, 68, 69}
		fmt.Println("vector:", vector)
		fmt.Println("vector como string:", string(vector))
	}
}

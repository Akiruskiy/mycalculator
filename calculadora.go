package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type calc struct{} // creacion de un objeto

func (calc) operate(entrada string) int {
	operador := obtenerOperador(entrada)
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	fmt.Print("El resultado de la operacion es:")
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	default:
		fmt.Println(operador, "El operador no esta soportado")
		return 0
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada) //convertimos los string a entero con Atoi(el segundo valor guarda errores)
	return operador
}

func obtenerOperador(operation string) string {
	re := regexp.MustCompile("[^0-9]+")              //crea una variable tipo regexp para poder analizarla
	operator := re.FindAllString(operation, -1)      // parece que encuentra el operador
	var operador string = strings.Join(operator, "") // aca tampoco se que es lo que pasa
	return operador
}

/*func main() {
	scanner := bufio.NewScanner(os.Stdin) //para leer el input del usuario
	scanner.Scan()                        //asignamos para que escanee
	entrada := scanner.Text()             //operacion guarda el input
	//operador := leerEntrada()
	c := calc{}
	fmt.Println(c.operate(entrada))
	/*fmt.Println(valores, operador)
	fmt.Println(valores[0] + valores[1]) //concatenamos el string, no hace la operacion aritmetica */

	/*if error1 != nil || error2 != nil { // en caso de error imprime el error
		fmt.Print("Existen los siguientes errores: ")
		fmt.Println(error1, error2)

	} else {

	}*?*/
}
*/

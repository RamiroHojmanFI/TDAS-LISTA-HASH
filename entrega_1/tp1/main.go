package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineaInfija := scanner.Text()

		resultadoPostfijo := ConvertirAPostfija(lineaInfija)

		fmt.Println(resultadoPostfijo)
	}

}

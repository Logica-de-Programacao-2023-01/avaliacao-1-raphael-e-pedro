package q5

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.
import (
	"fmt"
	"strings"
)

func ProcessString(s string) string {
	package main



func main() {

	Entrada := "Entrada fornecida"

	Entrada = strings.Map(func(r rune) rune {
		switch r {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			return -1
		default:
			return r
		}
	}, Entrada)

	Entrada = strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
			return '.' + r
		}
		return r
	}, Entrada)

	Entrada = strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return r + ('a' - 'A')
		}
		return r
	}, Entrada)

	fmt.Println(Entrada)
}

	return ""
}

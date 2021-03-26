package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type DataNascimento struct {
	Dia int
	Mes int
	Ano int
}

type Pessoa struct {
	Nome   string
	Altura int
	DataNascimento
}

var dados [3]int

func CriarData() [3]int {

	var ddia, dmes, dano int

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	ddia = r1.Intn(30) + 1
	dmes = r1.Intn(12) + 1
	dano = r1.Intn(49) + 1950

	dados[0] = ddia
	dados[1] = dmes
	dados[2] = dano

	return dados
}

func ExibirMenu() {
	fmt.Println("Digite 1 para exibir todas pessoas cadastradas")
	fmt.Println("Digite 2 para cadastrar uma nova pessoa")
	fmt.Println("Digite 3 para organizar as pessoas por altura")
	fmt.Println("Digite 4 para sair do programa")
	fmt.Println("-----------------------------------------------------")

}

type OrganizarAltura []Pessoa

func (x OrganizarAltura) Len() int {
	return len(x)
}
func (x OrganizarAltura) Less(i, j int) bool {
	return x[i].Altura < x[j].Altura
}
func (x OrganizarAltura) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}


func main() {

	sliceStruct := []Pessoa{}

	var NomePessoa string
	var AlturaPessoa int
	var entrada int

	for {
		// MENU

		ExibirMenu()

		fmt.Scan(&entrada)

		switch entrada {

		// EXIBIR
		case 1:

			fmt.Println(sliceStruct)

			break
			// INSERIR DADOS
		case 2:
			fmt.Println("Digite o nome da pessoa:")

			fmt.Scan(&NomePessoa)

			fmt.Println("Digite a altura da pessoa em centímetros")
			fmt.Scan(&AlturaPessoa)

			CriarData()

			pessoinha := Pessoa{DataNascimento: DataNascimento{Dia: dados[0], Mes: dados[1], Ano: dados[2]}, Nome: NomePessoa, Altura: AlturaPessoa}

			sliceStruct = append(sliceStruct, pessoinha)

		case 3:

			sort.Sort(OrganizarAltura(sliceStruct))

			fmt.Println(sliceStruct)

		case 4:

			goto FIMPROGRAMA
		}

	}

FIMPROGRAMA:
	fmt.Println("Programa finalizado")

}

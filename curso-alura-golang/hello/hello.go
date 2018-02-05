package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 2
const delay = 4

func main() {

	intro()

	for {
		optionsMenu()
		command := getCommand()

		switch command {
		case 1:
			initTracking()
		case 2:
			fmt.Println("Exibindo Logs...")
			printLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}
	}
}

func intro() {
	name := "Dayana"
	version := 1.1
	fmt.Println("Olá sra", name)
	fmt.Println("A versão do programa é", version)
}

func optionsMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")
}

func getCommand() int {
	var commandUser int
	fmt.Scan(&commandUser)
	//fmt.Println("o tipo da váriavel commandUser é", reflect.TypeOf(commandUser))
	//fmt.Println("o endereço da minha váriavel na memória é", &commandUser)
	fmt.Println("o comando escolhido foi", commandUser)

	return commandUser
}

func initTracking() {
	fmt.Println("Monitorando...")

	sites := readFile()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			checkStatusCode(site)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}

func checkStatusCode(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problema. Status Code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readFile() []string {

	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)

		sites = append(sites, row)

		if err == io.EOF {
			break
		}

	}

	file.Close()
	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file))
}

// O que foi aprendido

// Cap2
/*
- Declaração de variáveis
	Valor padrão das variáveis sem valor inicial
- Inferência de tipos de variáveis
- Descobrir o tipo da variável
	Através da função TypeOf, do pacote reflect
- Declaração curta de variáveis :=
- Ler dados digitados do usuário
	Através das funções Scanf e Scan, do pacote fmt
- Mais convenções do Go
	Variáveis e imports não utilizados são deletados
*/

// Cap3
/*
- Controle de fluxo com if
	Sua condição não fica entre parênteses e deve sempre retornar um booleano
- Controle de fluxo com switch
	Se os casos não forem atendidos, será executado o código do caso default
- Introdução às funções
- Pacote os, para encerrar o programa
*/

// Cap4
/*
- Pacote net/http, com funcionalidades de acesso à internet, de requisições web
	Entre elas, a função http.Get, para fazer uma requisição GET para um site
- Funções com múltiplos retornos
	Identificador em branco, para ignorar um ou mais retornos de uma função
- Status de uma requisição
	Uma requisição de sucesso possui status code 200
- A instrução for, para deixar o nosso programa em loop eterno-
*/

// Cap5
/*
- Coleções: Arrays e Slices
- Trabalhar com Arrays e Slices
- Diferenças entre Arrays e Slices
- A estrutura de repetição for
- Operador de iteração range
	O range nos dá acesso a cada item da coleção, nos retornando a posição do item iterado e o próprio item daquela posição
- Constantes
- Trabalhando com o pacote time
*/

// Cap6
/*
- Detectar erros
- Trabalhar com arquivos
	Inclusive abrindo e fechando arquivos
- Ler um arquivo por completo com o pacote io/ioutil
- Ler linha a linha com o pacote bufio
	Criação de leitores
- Limpando strings com TrimSpace
*/

// Cap7
/*
- Abrir arquivos
- Se o arquivo não existir, nós o criamos
- As flags da função os.OpenFile
- Escrever em arquivos
- Converter tipos para string
- Trabalhar com tempo e formatá-lo
- Exibir o conteúdo de um arquivo
*/

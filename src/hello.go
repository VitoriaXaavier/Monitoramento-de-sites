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

const monitoramentos = 5
const delay = 1

func main() {

	introducao()

	for {

		exibeMenu()
		comando := leComando()

		// usando a instrução de controle de fluxo switch

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Opção invalida!")
			os.Exit(-1)
		}

	}

}

func introducao() {
	var nome string
	versao := 1.1

	fmt.Println("Digite o seu nome!")
	fmt.Scan(&nome)

	fmt.Println("Seja bem-vindo(a), ", nome, "!")
	fmt.Println("A versão atual do sistema é a ", versao)
}

func exibeMenu() {
	fmt.Println("Escolha uma opção!")
	fmt.Println("1- Iniciar o monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

}

func leComando() int {

	var comando int
	fmt.Scan(&comando)
	fmt.Println("Opção digita foi: ", comando)

	return comando

}

func iniciarMonitoramento() {

	fmt.Println("Iniciando o monitoramento a cada 1 minuto por 6 vezes")

	fmt.Println("")

	sites := leSitesDoArquivo()

	for i := 0; i <= monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando o site ", i, ":", site)
			testaSite(site)

		}
		time.Sleep(delay * time.Minute)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSite(site string) {

	resp, erro := http.Get(site)

	if erro != nil {
		fmt.Println("Ocorreu um erro", erro)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Foi carregado com suscesso!")
		registraLog(site, true)

	} else {
		fmt.Println("Não foi possivel carregar. Status Code ", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, erro := os.Open("sites.txt")

	if erro != nil {
		fmt.Println("Ocorreu um erro", erro)
	}

	leitor := bufio.NewReader(arquivo)

	for {

		linha, erro := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if erro == io.EOF {
			break
		}

	}
	arquivo.Close()
	return sites

}

func registraLog(site string, status bool) {

	arquivo, erro := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if erro != nil {
		fmt.Println(erro)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - "+ site + " -online? -" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func exibirLogs () {

	fmt.Println("Exibindo logs")

	arquivo, erro := ioutil.ReadFile("log.txt")

	if erro != nil {
		fmt.Println("Aconteceu um erro", erro)
	}

	fmt.Println(string(arquivo))
}
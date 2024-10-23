package main

import (
	"comida/menu"
	"comida/pedido"
	"fmt"
	"strings"
	"time"
)

func main() {
	var nome string
	var rua string
	fmt.Printf("%15s\n", "____|____ Bem-Vindo à Lanchonete Blue Burger ____|____")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("Digite seu Nome: ")
	fmt.Scan(&nome)
	fmt.Printf("Digite sua Rua: ")
	fmt.Scan(&rua)
	menu.PrintMenu()

	var controle string
	fmt.Println("Deseja fazer um pedido? [Y / N]")
	fmt.Scan(&controle)
	if controle == "y" || controle == "Y" {
		pedido.Pedidos()
	} else {
		fmt.Printf("Finalizando pedido...")
		time.Sleep(time.Second * 3)
		pedido.Extrato()
	}

}

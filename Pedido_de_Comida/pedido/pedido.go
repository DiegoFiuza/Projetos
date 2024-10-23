package pedido

import (
	"comida/menu"
	"fmt"
)

// variavel global
var pedidos []Pedido

type Pedido struct {
	ItemID         uint
	QuantidadeItem uint
}

func Pedidos() {
	menu.PrintMenu()

	var parar = false
	var quantidade uint
	var itemPedido uint

	for _, item := range menu.Itens {
		fmt.Printf("Número: %d Item: %s, Preço: %.2f\n", item.ID, item.Name, item.Price)
	}

	for !parar {
		fmt.Println("Digite o ID do item desejado: ")
		fmt.Scanln(&itemPedido)
		fmt.Println("Digite aa quantidade do item: ")
		fmt.Scanln(&quantidade)
		pedidos = append(pedidos, Pedido{ItemID: itemPedido, QuantidadeItem: quantidade})

		var continuar string
		fmt.Println("Deseja continuar: [Y / N]")
		fmt.Scan(&continuar)
		if continuar == "n" || continuar == "N" {
			parar = true
		}
	}
	Extrato()

}

func Extrato() {
	for _, item := range pedidos {
		fmt.Printf("Item ID: %d, Quantidade: %d\n", item.ItemID, item.QuantidadeItem)
	}
	total := 0.0
	var control string
	fmt.Println("Deseja pagar? ")
	fmt.Scan(&control)
	if control == "y" || control == "Y" {
		//
	}
}

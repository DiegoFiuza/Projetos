package menu

import (
	"fmt"
	"strings"
)

type Menu struct {
	ID    uint
	Name  string
	Price float64
}

var Itens = []Menu{
	{1, "Coca-cola", 12},
	{2, "Água", 4},
	{3, "Pepsi", 12},
	{4, "Hamburger", 17},
	{5, "Cheeseburger", 18},
	{6, "X-bacon", 22},
	{7, "Galisburger", 20},
	{8, "Smash Burger", 15},
	{9, "Chiken Wrap", 19},
	{10, "Hot-dog", 12},
	{11, "Cookie", 6},
	{12, "Sorvete", 4},
	{13, "Torta de chocolate", 9},
	{14, "Torta de morango", 9},
	{15, "Torta alemã", 13},
}

func PrintMenu() {
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "ID", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, item := range Itens {
		fmt.Printf(" %-7d %.2f    %-4s\n", item.ID, item.Price, item.Name)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}

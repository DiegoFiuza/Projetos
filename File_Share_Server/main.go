package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "admin" {
		return "$1$x9Qkf5mW$6U..H2KAPEAOi4SP0xsDp/"
	}
	return ""
}

func main() {

	//tratar erro caso o usuário nao passe a porta para subir o server
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretório> <porta>")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	porta := os.Args[2]
	//cria um File server com a biblioteca http e usa o Dir

	//autenticação do server
	authenticator := auth.NewBasicAuthenticator("Myserver.com", Secret)
	//funçao q associa url
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))
	fmt.Printf("Subindo server na porta %s", porta)
	log.Fatal(http.ListenAndServe(":"+porta, nil))
}

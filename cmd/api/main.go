package main

import "flag"

type application struct {
}

func main() {
	addr := flag.String("addr", ":8080", "server port")
	flag.Parse()

	app := application{}

	router := app.routes()

	router.Run(*addr)

}

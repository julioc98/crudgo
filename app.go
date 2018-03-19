package main

import (
	"github.com/julioc98/crudgo/server"
	"github.com/julioc98/crudgo/user"
)

func main() {
	user.Migrate()
	server.Listen()
}

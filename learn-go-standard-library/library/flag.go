package library

import (
	"flag"
	"fmt"
)

func Flag() {
  var username *string = flag.String("username", "root", "database username")
  var password *string = flag.String("password", "root", "database password")
  var host *string = flag.String("host", "localhost", "database host")
  var port *int = flag.Int("port", 5432, "database port")

  flag.Parse()
  fmt.Println("username", *username)
  fmt.Println("password", *password)
  fmt.Println("host", *host)
  fmt.Println("port", *port)

}

package database

var connection string 

func init() {
  connection = "MySql"
}

func GetDatabase() string {
  return connection
}

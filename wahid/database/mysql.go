package database

var connection string

func init() {
	connection = "MySQL Connection"
}

func GetDatabase() string {
	return connection
}
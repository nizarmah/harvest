package env

type Env struct {
	DB *DB
}

type DB struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

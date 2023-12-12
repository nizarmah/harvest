package env

type Env struct {
	DB *DB
}

type DB struct {
	Name     string
	Host     string
	Username string
	Password string
}

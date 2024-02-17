package env

type Env struct {
	BaseURL string

	DB   *DB
	SMTP *SMTP
}

type DB struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

type SMTP struct {
	Host     string
	Port     string
	Username string
	Password string
}

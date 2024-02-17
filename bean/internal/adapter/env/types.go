package env

type Env struct {
	BaseURL string

	DB    *DB
	Cache *Cache
	SMTP  *SMTP
}

type DB struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

type Cache struct {
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

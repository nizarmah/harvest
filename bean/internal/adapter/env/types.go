package env

type Env struct {
	BaseURL string

	FeatureFlags *FeatureFlags

	DB    *DB
	Cache *Cache
	SMTP  *SMTP

	BuyMeACoffee *BuyMeACoffee
}

type FeatureFlags struct {
	BypassHTTPS      bool
	BypassMembership bool
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

type BuyMeACoffee struct {
	WebhookSecret string
}

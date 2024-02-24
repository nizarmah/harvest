package env

func New() (*Env, error) {
	baseURL, err := lookup("BASE_URL")
	if err != nil {
		return nil, err
	}

	dbName, err := lookup("DB_NAME")
	if err != nil {
		return nil, err
	}

	dbHost, err := lookup("DB_HOST")
	if err != nil {
		return nil, err
	}

	dbPort, err := lookup("DB_PORT")
	if err != nil {
		return nil, err
	}

	dbUsername, err := lookup("DB_USERNAME")
	if err != nil {
		return nil, err
	}

	dbPassword, err := lookup("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	cacheHost, err := lookup("CACHE_HOST")
	if err != nil {
		return nil, err
	}

	cachePort, err := lookup("CACHE_PORT")
	if err != nil {
		return nil, err
	}

	cacheUsername, err := lookup("CACHE_USERNAME")
	if err != nil {
		return nil, err
	}

	cachePassword, err := lookup("CACHE_PASSWORD")
	if err != nil {
		return nil, err
	}

	smtpHost, err := lookup("SMTP_HOST")
	if err != nil {
		return nil, err
	}

	smtpPort, err := lookup("SMTP_PORT")
	if err != nil {
		return nil, err
	}

	smtpUsername, err := lookup("SMTP_USERNAME")
	if err != nil {
		return nil, err
	}

	smtpPassword, err := lookup("SMTP_PASSWORD")
	if err != nil {
		return nil, err
	}

	bmcWebhookSecret, err := lookup("BMC_WEBHOOK_SECRET")
	if err != nil {
		return nil, err
	}

	return &Env{
		BaseURL: baseURL,
		DB: &DB{
			Name:     dbName,
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUsername,
			Password: dbPassword,
		},
		Cache: &Cache{
			Host:     cacheHost,
			Port:     cachePort,
			Username: cacheUsername,
			Password: cachePassword,
		},
		SMTP: &SMTP{
			Host:     smtpHost,
			Port:     smtpPort,
			Username: smtpUsername,
			Password: smtpPassword,
		},
		BuyMeACoffee: &BuyMeACoffee{
			WebhookSecret: bmcWebhookSecret,
		},
	}, nil
}

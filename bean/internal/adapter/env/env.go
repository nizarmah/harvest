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

	return &Env{
		BaseURL: baseURL,
		DB: &DB{
			Name:     dbName,
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUsername,
			Password: dbPassword,
		},
		SMTP: &SMTP{
			Host:     smtpHost,
			Port:     smtpPort,
			Username: smtpUsername,
			Password: smtpPassword,
		},
	}, nil
}

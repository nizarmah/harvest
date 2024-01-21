package env

func New() (*Env, error) {
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

	return &Env{
		DB: &DB{
			Name:     dbName,
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUsername,
			Password: dbPassword,
		},
	}, nil
}

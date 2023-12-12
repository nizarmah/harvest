package env

func New() (*Env, error) {
	db, err := newDB()
	if err != nil {
		return nil, err
	}

	return &Env{
		DB: db,
	}, nil
}

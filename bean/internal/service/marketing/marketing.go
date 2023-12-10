package marketing

func Init() *Marketing {
	return &Marketing{
		Handlers: &Handlers{},
	}
}

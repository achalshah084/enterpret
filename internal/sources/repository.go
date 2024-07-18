package sources

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repository Repository) FindById(id string) (Sources, error) {
	return Sources{
		Id:   "1234",
		Type: Twitter,
	}, nil
}

package tenantSources

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repository Repository) FindById(id string) (TenantSource, error) {
	return TenantSource{
		Id:  "1",
		Key: "key",
	}, nil
}

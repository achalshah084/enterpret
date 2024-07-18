package tenants

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repository Repository) findById(id string) (Tenant, error) {
	return Tenant{
		Id:              "1",
		Name:            "Rapido",
		TenantSourceIds: []string{"1"},
	}, nil
}

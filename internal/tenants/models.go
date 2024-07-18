package tenants

type Tenant struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	TenantSourceIds []string `json:"tenantSourceIds"`
}

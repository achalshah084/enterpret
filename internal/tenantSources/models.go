package tenantSources

type TenantSource struct {
	Id       string `json:"id"`
	SourceId string `json:"sourceId"`
	Key      string `json:"config"`
	Handle   string `json:"handle"`
	TenantId string `json:"tenantId"`
}

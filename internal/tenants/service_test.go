package tenants

import (
	"enterpret/internal/feedBacks"
	"enterpret/internal/tenantSources"
	"testing"
)

func TestService_PullData(t *testing.T) {
	type fields struct {
		repository           Repository
		tenantSourcesService tenantSources.Service
		feedBacksRepository  feedBacks.Repository
	}
	type args struct {
		startDate string
		endDate   string
		id        string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Pull Data from twitter for the tenant (tenants Service)",
			fields:  fields{},
			args:    args{startDate: "2024-07-18", endDate: "2024-08-18", id: "1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository:           tt.fields.repository,
				tenantSourcesService: tt.fields.tenantSourcesService,
				feedBacksRepository:  tt.fields.feedBacksRepository,
			}
			if err := s.PullData(tt.args.startDate, tt.args.endDate, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PullData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package feedBacks

import (
	"enterpret/internal/sources"
	"enterpret/internal/tenantSources"
	"testing"
)

func TestService_PushData(t *testing.T) {
	type fields struct {
		repository              Repository
		twitterSource           sources.TwitterSource
		tenantSourcesRepository tenantSources.Repository
		sourcesRepository       sources.Repository
	}
	type args struct {
		feedBackId     string
		tenantSourceId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Push Data from twitter for the tenant (feedBacks Service)",
			fields:  fields{},
			args:    args{feedBackId: "123", tenantSourceId: "1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository:              tt.fields.repository,
				twitterSource:           tt.fields.twitterSource,
				tenantSourcesRepository: tt.fields.tenantSourcesRepository,
				sourcesRepository:       tt.fields.sourcesRepository,
			}
			if err := s.PushData(tt.args.feedBackId, tt.args.tenantSourceId); (err != nil) != tt.wantErr {
				t.Errorf("PushData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

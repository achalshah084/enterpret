package tenants

import (
	"enterpret/internal/feedBacks"
	"enterpret/internal/tenantSources"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		repository           Repository
		tenantSourcesService tenantSources.Service
		feedBacksRepository  feedBacks.Repository
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.repository, tt.args.tenantSourcesService, tt.args.feedBacksRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		// TODO: Add test cases.
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

package feedBacks

import (
	"enterpret/internal/sources"
	"enterpret/internal/tenantSources"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		repository              Repository
		twitterSource           sources.TwitterSource
		tenantSourcesRepository tenantSources.Repository
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
			if got := NewService(tt.args.repository, tt.args.twitterSource, tt.args.tenantSourcesRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_FindAll(t *testing.T) {
	type fields struct {
		repository              Repository
		twitterSource           sources.TwitterSource
		tenantSourcesRepository tenantSources.Repository
		sourcesRepository       sources.Repository
	}
	tests := []struct {
		name   string
		fields fields
		want   []shared.Feedback
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository:              tt.fields.repository,
				twitterSource:           tt.fields.twitterSource,
				tenantSourcesRepository: tt.fields.tenantSourcesRepository,
				sourcesRepository:       tt.fields.sourcesRepository,
			}
			if got := s.FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		// TODO: Add test cases.
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

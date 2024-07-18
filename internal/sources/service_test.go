package sources

import (
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		repository    Repository
		twitterSource TwitterSource
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
			if got := NewService(tt.args.repository, tt.args.twitterSource); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_PullData(t *testing.T) {
	type fields struct {
		repository    Repository
		twitterSource TwitterSource
	}
	type args struct {
		startDate      string
		endDate        string
		key            string
		sourceId       string
		tenantSourceId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []shared.Feedback
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository:    tt.fields.repository,
				twitterSource: tt.fields.twitterSource,
			}
			got, err := s.PullData(tt.args.startDate, tt.args.endDate, tt.args.key, tt.args.sourceId, tt.args.tenantSourceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PullData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PullData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

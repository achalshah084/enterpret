package tenantSources

import (
	"enterpret/internal/shared"
	"enterpret/internal/sources"
	"reflect"
	"testing"
)

func TestService_PullData(t *testing.T) {
	type fields struct {
		repository     Repository
		sourcesService sources.Service
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
		want    []shared.Feedback
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Pull Data from twitter for the tenant",
			fields: fields{},
			args:   args{startDate: "2024-07-18", endDate: "2024-08-18", id: "1"},
			want: []shared.Feedback{
				shared.Conversation{
					BaseFeedback: shared.BaseFeedback{
						Id:      "123",
						Content: "very well said",
						Type:    "conversation",
						Metadata: shared.TwitterMetadata{
							TenantSourceId: "1",
							Country:        "india",
						},
					},
					Participants: nil,
				},
				shared.Conversation{
					BaseFeedback: shared.BaseFeedback{
						Id:      "123",
						Content: "very well said",
						Type:    "conversation",
						Metadata: shared.TwitterMetadata{
							TenantSourceId: "1",
							Country:        "india",
						},
					},
					Participants: nil,
				},
				shared.Conversation{
					BaseFeedback: shared.BaseFeedback{
						Id:      "234",
						Content: "very well said",
						Type:    "conversation",
						Metadata: shared.TwitterMetadata{
							TenantSourceId: "1",
							Country:        "usa",
						},
					},
					Participants: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository:     tt.fields.repository,
				sourcesService: tt.fields.sourcesService,
			}
			got, err := s.PullData(tt.args.startDate, tt.args.endDate, tt.args.id)
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

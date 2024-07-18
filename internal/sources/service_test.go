package sources

import (
	"enterpret/internal/shared"
	"reflect"
	"testing"
)

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
		{
			name:   "Pull Data from twitter for the tenant (sources Service)",
			fields: fields{},
			args:   args{startDate: "2024-07-18", endDate: "2024-08-18"},
			want: []shared.Feedback{
				shared.Conversation{
					BaseFeedback: shared.BaseFeedback{
						Id:      "123",
						Content: "very well said",
						Type:    "conversation",
						Metadata: shared.TwitterMetadata{
							Country: "india",
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
							Country: "india",
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
							Country: "usa",
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

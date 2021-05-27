package eventgen

import (
	"log"
	"testing"
)

func Test_generateEvent(t *testing.T) {
	type args struct {
		eventType int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default test case",
			args: args{
				eventType: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateEvent(tt.args.eventType)
			log.Printf("generateEvent() = %v", got)
		})
	}
}

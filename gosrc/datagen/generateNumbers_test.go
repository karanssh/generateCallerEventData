package datagen

import (
	"log"
	"testing"
)

func TestRandomNumberInRange(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default test case",
			args: args{
				low:  9030000000,
				high: 9999999999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomNumberInRange(tt.args.low, tt.args.high)
			if !(got < tt.args.high || got > tt.args.low) {
				t.Errorf("RandomNumberInRange() = %v", got)
			}
			log.Printf("RandomNumberInRange() = %v", got)
		})
	}
}

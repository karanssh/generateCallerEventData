package datagen

import (
	"log"
	"testing"
)

func TestRandStringRunes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default test case",
			args: args{
				n: 22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomStringFromRunes(tt.args.n); len(got) > 0 {
				log.Printf("RandomStringFromRunes() = %v ", got)
				t.Logf("RandomStringFromRunes() = %v ", got)
			}
		})
	}
}

func TestRandomTime(t *testing.T) {

	//probably a better way to write test case exists but its pure func so no need to worry
	got := RandomTime()

	log.Printf("time generated is %s ", got)

}

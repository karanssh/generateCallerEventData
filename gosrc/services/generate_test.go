package services

import (
	"generateEventData/datadefinition"
	"reflect"
	"testing"
)

func TestGenerateJSONData(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "generic test 1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateJSONData(); (err != nil) != tt.wantErr {
				t.Errorf("GenerateJSONData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateCSVData(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "generic test 1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateCSVData(); (err != nil) != tt.wantErr {
				t.Errorf("GenerateCSVData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_generateItemData(t *testing.T) {
	tests := []struct {
		name string
		want []datadefinition.Event
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateItemData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateItemData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGenerateItemData(b *testing.B) {
	generateItemDataWithCount(10000, 10000, 10000, 10000)

}

func BenchmarkGenerateItemDataNewImpl(b *testing.B) {

	generateItemDataWithCountNewImpl(10000, 10000, 10000, 10000)
}

package main

import "testing"

func TestCar_ToString(t *testing.T) {
	type fields struct {
		Number string
		Model  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Car{
				Number: tt.fields.Number,
				Model:  tt.fields.Model,
			}
			if got := c.ToString(); got != tt.want {
				t.Errorf("Car.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

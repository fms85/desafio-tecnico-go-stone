package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{
			name: "should_sum_positive_numbers",
			a:    2.5,
			b:    3.7,
			want: 6.2,
		},
		{
			name: "should_sum_negative_numbers",
			a:    -10.0,
			b:    -5.0,
			want: -15.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.a, tt.b)
			assert.InDelta(t, tt.want, got, 0.0001, "Unexpected sum result")
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{
			name: "should_subtract_positive_numbers",
			a:    8.9,
			b:    4.2,
			want: 4.7,
		},
		{
			name: "should_subtract_negative_numbers",
			a:    -7.5,
			b:    -3.2,
			want: -4.3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sub(tt.a, tt.b)
			assert.InDelta(t, tt.want, got, 0.0001, "Unexpected subtraction result")
		})
	}
}

package internal

import (
	"reflect"
	"testing"
	"time"
)

func TestDaysBetween(t *testing.T) {
	from := time.Date(2023, time.December, 30, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)

	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		{
			name: "usual from -> to, with year and month cross",
			args: args{
				from: from,
				to:   to,
			},
			want: []time.Time{
				time.Date(2023, time.December, 30, 0, 0, 0, 0, time.UTC),
				time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
				time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "empty to -> from",
			args: args{
				from: to,
				to:   from,
			},
			want: []time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DaysBetween(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DaysBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

package min_time_to_visit_all_points

import "testing"

func Test_minTimeToVisitsAllPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				points: [][]int{{1, 1}, {3, 4}, {-1, 0}},
			},
			want: 7,
		},
		{
			name: "Example2",
			args: args{
				points: [][]int{{3, 2}, {-2, 2}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToVisitsAllPoints(tt.args.points); got != tt.want {
				t.Errorf("minTimeToVisitsAllPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

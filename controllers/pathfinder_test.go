package controllers

import (
	"lemin/models"
	"reflect"
	"testing"
)

func TestPathFinder(t *testing.T) {
	type args struct {
		colony models.Colony
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Valid path from start to end",
			args: args{
				colony: models.Colony{
					Rooms: map[string]*models.Room{
						"0": {Name: "0", Neighbours: []*models.Room{}},
						"1": {Name: "1", Neighbours: []*models.Room{}},
						"2": {Name: "2", Neighbours: []*models.Room{}},
						"3": {Name: "3", Neighbours: []*models.Room{}},
					},
					StartRoom:  models.Room{Name: "0"},
					EndRoom:    models.Room{Name: "3"},
					StartFound: true,
					EndFound:   true,
				},
			},
			want: [][]string{
				{"0", "1", "2", "3"},
			},
			wantErr: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathFinder(tt.args.colony)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathFinder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathFinder() = %v, want %v", got, tt.want)
			}
		})
	}
}

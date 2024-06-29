package pkg

import (
	"dir_to_json/model"
	"os"
	"reflect"
	"testing"
)

func Test_buildFileStructure(t *testing.T) {

	// create folder /tmp/test48541545695556
	os.Mkdir("/tmp/test48541545695556", 0777)
	// create file /tmp/test48541545695556/test
	file, _ := os.Create("/tmp/test48541545695556/test")
	file.Close()

	type args struct {
		fileExploremodel model.FileExplore
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Test 1 level 1", args: args{fileExploremodel: model.FileExplore{Path: "/tmp/test48541545695556", Level: 1, Hidden: false}}, want: "test", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildFileStructure(tt.args.fileExploremodel)

			if (err != nil) != tt.wantErr {
				t.Errorf("buildFileStructure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Children[0].Name, tt.want) {
				t.Errorf("buildFileStructure() = %v, want %v", got, tt.want)
			}
		})
	}

	// remove folder /tmp/test48541545695556
	os.RemoveAll("/tmp/test48541545695556")
}

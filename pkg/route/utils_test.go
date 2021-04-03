package route

import (
	"reflect"
	"testing"
)

func TestConvertStrToIntArray(t *testing.T) {
	given := "1,2,3,4,5"
	got, err := convertStrToIntArray(given)
	if err != nil {
		t.Fatal(err)
	}
	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v, given %v", got, want, given)
	}
}

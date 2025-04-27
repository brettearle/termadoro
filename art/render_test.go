package art

import (
	"reflect"
	"testing"
)

func TestRender(t *testing.T) {
	var testFrame1 Frame = []byte("-")
	var testFrame2 Frame = []byte("--")
	var testFrame3 Frame = []byte("---")
	var testFrame4 Frame = []byte("--->")
	t.Run("newReel", func(t *testing.T) {
		testFrames := []*Frame{&testFrame1, &testFrame2, &testFrame3, &testFrame4}
		got, _ := newReel(testFrames)
		want := Reel{
			count:  0,
			frames: testFrames,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %+v but got %+v", want, got)
		}
	})
}

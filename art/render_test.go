package art

import (
	"bytes"
	"reflect"
	"testing"
)

func setUp() ([]*Frame, Reel) {
	var testFrame1 Frame = []byte("-")
	var testFrame2 Frame = []byte("--")
	var testFrame3 Frame = []byte("---")
	var testFrame4 Frame = []byte("--->")

	testFrames := []*Frame{&testFrame1, &testFrame2, &testFrame3, &testFrame4}
	testReel := Reel{
		currentFrame: 0,
		frames:       testFrames,
	}
	return testFrames, testReel
}

func TestRender(t *testing.T) {
	t.Run("newReel", func(t *testing.T) {
		testFrames, want := setUp()
		got, _ := newReel(testFrames)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %+v but got %+v", want, got)
		}
	})

	t.Run("reel.Draw", func(t *testing.T) {
		testFrames, testReel := setUp()
		got := new(bytes.Buffer)
		testReel.Draw(got)
		want := string(*testFrames[0])
		if got.String() != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})

	t.Run("reel.Next", func(t *testing.T) {
		_, testReel := setUp()
		testReel.Next()
		got := testReel.currentFrame
		want := 1
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})

	t.Run("reel.Next should reset to 1st frame if at last frame of sequence", func(t *testing.T) {
		_, testReel := setUp()
		// set current frame to the last frame in sequence
		testReel.currentFrame = 3
		testReel.Next()
		got := testReel.currentFrame
		want := 0
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}

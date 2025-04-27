// art is the renderer for tdoro
package art

import "io"

// Is a single frame of a moving picture
type Frame []byte

// State of moving picture
type Reel struct {
	currentFrame int
	frames       []*Frame
}

func newReel(frs []*Frame) (Reel, error) {
	return Reel{
		currentFrame: 0,
		frames:       frs,
	}, nil

}

// Draws current frame in Reel to writer
func (r *Reel) Draw(w io.Writer) {
	w.Write(*r.frames[r.currentFrame])
}

// Moves currentFrame to next frame in sequence
// If last frame in sequence goes back to the first
func (r *Reel) Next() {
	if r.currentFrame == len(r.frames)-1 {
		r.currentFrame = 0
	} else {
		r.currentFrame += 1
	}

}

// TODO:Remove this, currently untested
func NewClockHeadWorkReel() Reel {
	var byteStr1 Frame = []byte(ClockHeadWork1)
	var byteStr2 Frame = []byte(ClockHeadWork2)
	var byteStr3 Frame = []byte(ClockHeadWork3)
	var byteStr4 Frame = []byte(ClockHeadWork4)

	return Reel{
		currentFrame: 0,
		frames:       []*Frame{&byteStr1, &byteStr2, &byteStr3, &byteStr4},
	}
}

// TODO:Remove this, currently untested
func NewClockHeadRestReel() Reel {
	var byteStr1 Frame = []byte(ClockHeadSleep1)
	var byteStr2 Frame = []byte(ClockHeadSleep2)
	var byteStr3 Frame = []byte(ClockHeadSleep3)
	var byteStr4 Frame = []byte(ClockHeadSleep4)
	var byteStr5 Frame = []byte(ClockHeadSleep5)

	return Reel{
		currentFrame: 0,
		frames:       []*Frame{&byteStr1, &byteStr2, &byteStr3, &byteStr4, &byteStr5},
	}
}

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

// TODO:
func NewClockHeadWorkReel() Reel {
	return Reel{}
}

// TODO:
func NewClockHeadRestReel() Reel {
	return Reel{}
}

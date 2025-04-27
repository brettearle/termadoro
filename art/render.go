package art

type Frame []byte

type Reel struct {
	count  int
	frames []*Frame
}

func newReel(frs []*Frame) (Reel, error) {
	return Reel{
		count:  0,
		frames: frs,
	}, nil

}

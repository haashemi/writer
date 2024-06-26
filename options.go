package writer

import "golang.org/x/text/unicode/bidi"

// DefaultOptions contains the recommended default options
var DefaultOptions = Options{
	Bidi:     true,
	Features: DefaultFeatures,
}

// Options holds the features used to modify, shape, and write the text.
type Options struct {
	Bidi     bool
	Features []Feature // Feature are the OpenType feature you want to enable.
	// TODO: Features maybe should be a struct instead of a slice?
}

// bidiText converts a bi-directional text logically to visually.
func bidiText(in string) (out string, err error) {
	p := bidi.Paragraph{}
	_, err = p.SetString(in)
	if err != nil {
		return
	}

	o, err := p.Order()
	if err != nil {
		return
	}

	mainDirection := p.Direction()

	for i := 0; i < o.NumRuns(); i++ {
		r := o.Run(i)

		switch r.Direction() {
		case bidi.LeftToRight:
			if mainDirection == bidi.LeftToRight {
				out += r.String()
			} else {
				out += bidi.ReverseString(r.String())
			}
		case bidi.RightToLeft:
			if mainDirection == bidi.RightToLeft {
				out += r.String()
			} else {
				out += bidi.ReverseString(r.String())
			}
		default:
			out += r.String()
		}
	}

	return
}

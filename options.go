package writer

import "golang.org/x/text/unicode/bidi"

type Options struct {
	Bidi bool
	// TODO: Align, Justify, IndentLeftMargin, IndentRightMargin, IndentFirstLine,
	// BeforeParagraphSpace, AfterParagraphSpace, TypeOptions, EastAsianFeatures?
}

func bidiText(in string) (out string, err error) {
	p := bidi.Paragraph{}
	p.SetString(in)

	o, err := p.Order()
	if err != nil {
		return
	}

	mainDirection := p.Direction()

	for i := range o.NumRuns() {
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

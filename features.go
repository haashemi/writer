package writer

import (
	"github.com/haashemi/go-harfbuzz/hb"
)

// Feature is just an alias to harfbuzz's Feature type.
type Feature = hb.Feature

// DefaultFeatures contains “enabled by default” features.
var DefaultFeatures = []Feature{
	StandardLigatures(),
	ContextualAlternatives(),
	Ornaments(),
	JustificationAlternatives(), // Should it be default or...?
}

// Learn more: https://learn.microsoft.com/en-us/typography/opentype/spec/features_ko#tag-liga
func StandardLigatures() Feature {
	return Feature{Tag: tag("liga"), Value: 1, Start: 0, End: 4294967295}
}

// Learn more: https://learn.microsoft.com/en-us/typography/opentype/spec/features_ae#tag-calt
func ContextualAlternatives() Feature {
	return Feature{Tag: tag("calt"), Value: 1, Start: 0, End: 4294967295}
}

// Learn more: https://learn.microsoft.com/en-us/typography/opentype/spec/features_ae#tag-dlig
func DiscretionaryLigatures() Feature {
	return Feature{Tag: tag("dlig"), Value: 1, Start: 0, End: 4294967295}
}

// Learn More: https://learn.microsoft.com/en-us/typography/opentype/spec/features_pt#tag-swsh
func Swash() Feature {
	return Feature{Tag: tag("swsh"), Value: 1, Start: 0, End: 4294967295}
}

// Learn more: https://learn.microsoft.com/en-us/typography/opentype/spec/features_ko#tag-onum
func OldStyle() Feature {
	return Feature{Tag: tag("onum"), Value: 1, Start: 0, End: 4294967295}
}

// Learn More: https://learn.microsoft.com/en-us/typography/opentype/spec/features_pt#tag-salt
func StylisticAlternatives() Feature {
	return Feature{Tag: tag("salt"), Value: 1, Start: 0, End: 4294967295}
}

// Learn More: https://learn.microsoft.com/en-us/typography/opentype/spec/features_pt#tag-titl
func TitlingAlternatives() Feature {
	return Feature{Tag: tag("titl"), Value: 1, Start: 0, End: 4294967295}
}

// Learn More: https://learn.microsoft.com/en-us/typography/opentype/spec/features_ko#tag-ornm
func Ornaments() Feature {
	return Feature{Tag: tag("ornm"), Value: 1, Start: 0, End: 4294967295}
}

// Learn More: https://learn.microsoft.com/en-us/typography/opentype/spec/features_ko#tag-ordn
func Ordinals() Feature {
	return Feature{Tag: tag("ordn"), Value: 1, Start: 0, End: 4294967295}
}

// Learn More: https://learn.microsoft.com/en-us/typography/opentype/spec/features_fj#tag-frac
func Fractions() Feature {
	return Feature{Tag: tag("frac"), Value: 1, Start: 0, End: 4294967295}
}

// Learn More: https://learn.microsoft.com/en-us/typography/opentype/spec/features_fj#tag-jalt
func JustificationAlternatives() Feature {
	return Feature{Tag: tag("jalt"), Value: 1, Start: 0, End: 4294967295}
}

// tag converts str to a Tag
func tag(str string) hb.Tag {
	if len(str) < 4 {
		return hb.Tag{}
	}
	return hb.Tag{str[3], str[2], str[1], str[0]}
}

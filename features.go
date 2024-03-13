package writer

import "github.com/haashemi/go-harfbuzz/hb"

type Feature = hb.Feature

var DefaultFeatures = []Feature{
	StandardLigatures(),
	ContextualAlternatives(),
	Ornaments(),
	JustificationAlternatives(), // Should it be default or...?
}

// Learn more: https://learn.microsoft.com/en-us/typography/opentype/spec/features_ko#tag-liga
func StandardLigatures() Feature {
	return Feature{Tag: hb.Tag{'l', 'i', 'g', 'a'}, Value: 1, Start: 0, End: 4294967295}
}

// Learn more: https://learn.microsoft.com/en-us/typography/opentype/spec/features_ae#tag-calt
func ContextualAlternatives() Feature {
	return Feature{Tag: hb.Tag{'c', 'a', 'l', 't'}, Value: 1, Start: 0, End: 4294967295}
}

// Learn more: https://learn.microsoft.com/en-us/typography/opentype/spec/features_ae#tag-dlig
func DiscretionaryLigatures() Feature {
	return Feature{Tag: hb.Tag{'d', 'l', 'i', 'g'}, Value: 1, Start: 0, End: 4294967295}
}

func Swash() Feature {
	return Feature{Tag: hb.Tag{'s', 'w', 's', 'h'}, Value: 1, Start: 0, End: 4294967295}
}

// Learn more: https://learn.microsoft.com/en-us/typography/opentype/spec/features_ko#tag-onum
func OldStyle() Feature {
	return Feature{Tag: hb.Tag{'o', 'n', 'u', 'm'}, Value: 1, Start: 0, End: 4294967295}
}

func StylisticAlternatives() Feature {
	return Feature{Tag: hb.Tag{'s', 'a', 'l', 't'}, Value: 1, Start: 0, End: 4294967295}
}

func TitlingAlternatives() Feature {
	return Feature{Tag: hb.Tag{'t', 'i', 't', 'l'}, Value: 1, Start: 0, End: 4294967295}
}

func Ornaments() Feature {
	return Feature{Tag: hb.Tag{'o', 'r', 'n', 'm'}, Value: 1, Start: 0, End: 4294967295}
}

func Ordinals() Feature {
	return Feature{Tag: hb.Tag{'o', 'r', 'd', 'n'}, Value: 1, Start: 0, End: 4294967295}
}

func Fractions() Feature {
	return Feature{Tag: hb.Tag{'f', 'r', 'a', 'c'}, Value: 1, Start: 0, End: 4294967295}
}

func JustificationAlternatives() Feature {
	return Feature{Tag: hb.Tag{'j', 'a', 'l', 't'}, Value: 1, Start: 0, End: 4294967295}
}

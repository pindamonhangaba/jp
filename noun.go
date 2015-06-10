package jp

type Noun struct {
	Stem string
}

func NewNoun(noun string) Noun {
	return Noun{Stem: noun}
}

func (noun *Noun) conjugate() []AdjConjugation {
	var adjs []AdjConjugation
	// using na-adjectives for nouns since they don't modify the last character like i-adjectives.
	for _, conj := range naAdjConjugations {
		adjconj := conj
		adjconj.Conjugated = noun.Stem + conj.Conjugation
		append(adjs, adjconj)
	}
	return adjs
}

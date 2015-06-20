package jp

type AdjConjugation struct {
	Polarity,
	Politeness,
	Tense,
	AdvEndChange,
	Conjugation,
	Example,
	Conjugated string
	TeForm bool
}

var iAdjConjugations []AdjConjugation = []AdjConjugation{
	// Informal (Common 普段)
	AdjConjugation{Polarity: "Affirmative", Politeness: "Informal (Common 普段)", Tense: "Present", AdvEndChange: "い", Conjugation: "", Example: "is/are", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Informal (Common 普段)", Tense: "Present", AdvEndChange: "く", Conjugation: "ない", Example: "is not/are not", TeForm: false},
	AdjConjugation{Polarity: "Affirmative", Politeness: "Informal (Common 普段)", Tense: "Past", AdvEndChange: "", Conjugation: "かった", Example: "was/were", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Informal (Common 普段)", Tense: "Past", AdvEndChange: "く", Conjugation: "なかった", Example: "was/were", TeForm: false},
	AdjConjugation{Polarity: "Tentative", Politeness: "Informal (Common 普段)", Tense: "", AdvEndChange: "", Conjugation: "かろう", Example: "is/are probably", TeForm: false},
	// Te-form
	AdjConjugation{Polarity: "", Politeness: "", Tense: "", AdvEndChange: "く", Conjugation: "て", Example: "", TeForm: true},
	// Formal (Polite 丁寧)
	AdjConjugation{Polarity: "Affirmative", Politeness: "Formal (Polite 丁寧)", Tense: "Present", AdvEndChange: "い", Conjugation: "です", Example: "is/are", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧)", Tense: "Present", AdvEndChange: "く", Conjugation: "ないです", Example: "is not/are not", TeForm: false},
	AdjConjugation{Polarity: "Affirmative", Politeness: "Formal (Polite 丁寧)", Tense: "Past", AdvEndChange: "", Conjugation: "かったです", Example: "was/were", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧)", Tense: "Past", AdvEndChange: "く", Conjugation: "なかったです", Example: "was/were", TeForm: false},
	AdjConjugation{Polarity: "Tentative", Politeness: "Formal (Polite 丁寧)", Tense: "", AdvEndChange: "い", Conjugation: "でしょう", Example: "is/are probably", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧))", Tense: "Present", AdvEndChange: "く", Conjugation: "ありません", Example: "is not/are not", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧)", Tense: "Past", AdvEndChange: "く", Conjugation: "ありませんでした", Example: "was/were", TeForm: false},
	// Conditional And Conjunctive
	AdjConjugation{Polarity: "", Politeness: "", Tense: "Conditional", AdvEndChange: "", Conjugation: "ければ", Example: "if [subject] is/was cold", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "", Tense: "Conditional", AdvEndChange: "く", Conjugation: "なければ", Example: "if [subject] is not/was not cold", TeForm: false},
	AdjConjugation{Polarity: "", Politeness: "", Tense: "Conjunctive", AdvEndChange: "く", Conjugation: "て", Example: "is/are cold ... and/so", TeForm: false},
	AdjConjugation{Polarity: "Negative ", Politeness: "", Tense: "Conjunctive", AdvEndChange: "く", Conjugation: "なくて", Example: "is not/are not cold ... and/so", TeForm: false},
}

var naAdjConjugations []AdjConjugation = []AdjConjugation{
	// Informal (Common 普段)
	AdjConjugation{Polarity: "Affirmative", Politeness: "Informal (Common 普段)", Tense: "Present", AdvEndChange: "", Conjugation: "だ", Example: "is/are", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Informal (Common 普段)", Tense: "Present", AdvEndChange: "", Conjugation: "ではない", Example: "is not/are not", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Informal (Common 普段)", Tense: "Present", AdvEndChange: "", Conjugation: "じゃない", Example: "is not/are not", TeForm: false},
	AdjConjugation{Polarity: "Affirmative", Politeness: "Informal (Common 普段)", Tense: "Past", AdvEndChange: "", Conjugation: "だった", Example: "was/were", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Informal (Common 普段)", Tense: "Past", AdvEndChange: "", Conjugation: "ではなかった", Example: "was/were", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Informal (Common 普段)", Tense: "Past", AdvEndChange: "", Conjugation: "じゃなかった", Example: "was/were", TeForm: false},
	AdjConjugation{Polarity: "Tentative", Politeness: "Informal (Common 普段)", Tense: "", AdvEndChange: "", Conjugation: "だろう", Example: "is/are probably", TeForm: false},
	// Te-form
	AdjConjugation{Polarity: "", Politeness: "", Tense: "", AdvEndChange: "", Conjugation: "で", Example: "", TeForm: true},
	// Formal (Polite 丁寧)
	AdjConjugation{Polarity: "Affirmative", Politeness: "Formal (Polite 丁寧)", Tense: "Present", AdvEndChange: "", Conjugation: "です", Example: "is/are", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧)", Tense: "Present", AdvEndChange: "", Conjugation: "ではありません", Example: "is not/are not", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧)", Tense: "Present", AdvEndChange: "", Conjugation: "じゃありません", Example: "is not/are not", TeForm: false},
	AdjConjugation{Polarity: "Affirmative", Politeness: "Formal (Polite 丁寧)", Tense: "Past", AdvEndChange: "", Conjugation: "でした", Example: "was/were", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧)", Tense: "Past", AdvEndChange: "", Conjugation: "ではありませんでした", Example: "was/were", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧)", Tense: "Past", AdvEndChange: "", Conjugation: "じゃありませんでした", Example: "was/were", TeForm: false},
	AdjConjugation{Polarity: "Tentative", Politeness: "Formal (Polite 丁寧)", Tense: "", AdvEndChange: "", Conjugation: "でしょう", Example: "is/are probably", TeForm: false},
	// extra
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧))", Tense: "Present", AdvEndChange: "", Conjugation: "ではないです", Example: "is not/are not", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧))", Tense: "Present", AdvEndChange: "", Conjugation: "じゃないです", Example: "is not/are not", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧))", Tense: "Past", AdvEndChange: "", Conjugation: "ではなかったです", Example: "is not/are not", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "Formal (Polite 丁寧))", Tense: "Past", AdvEndChange: "", Conjugation: "じゃなかったです", Example: "is not/are not", TeForm: false},
	// Conditional And Conjunctive
	AdjConjugation{Polarity: "", Politeness: "", Tense: "Conditional", AdvEndChange: "", Conjugation: "だったら", Example: "if [subject] is/was cold", TeForm: false},
	AdjConjugation{Polarity: "Negative", Politeness: "", Tense: "Conditional", AdvEndChange: "", Conjugation: "ではなかったら", Example: "if [subject] is not/was not cold", TeForm: false},
	AdjConjugation{Polarity: "", Politeness: "", Tense: "Conjunctive", AdvEndChange: "", Conjugation: "で", Example: "is/are cold ... and/so", TeForm: false},
	AdjConjugation{Polarity: "Negative ", Politeness: "", Tense: "Conjunctive", AdvEndChange: "", Conjugation: "ではなくて", Example: "is not/are not cold ... and/so", TeForm: false},
}

type NAAdjective struct {
	Stem string
}

type IAdjective struct {
	Raw,
	End,
	Stem string
}

func NewNAAdjective(adjective string) NAAdjective {
	return NAAdjective{Stem: adjective}
}

func NewIAdjective(adjective string) IAdjective {
	adj := IAdjective{Raw: adjective}
	adj.Stem = adjective[:len(adjective)-1]
	adj.Stem = adjective[len(adjective)-1:]
	return adj
}

func (adj *IAdjective) Conjugate() []AdjConjugation {
	var adjs []AdjConjugation
	for _, conj := range iAdjConjugations {
		adjconj := conj
		// treat special case いい
		if adj.Raw == "いい" && adjconj.Polarity != "Affirmative" && adjconj.Tense != "Present" {
			adjconj.Conjugated = "よ" + conj.AdvEndChange + conj.Conjugation
		} else {
			adjconj.Conjugated = adj.Stem + conj.AdvEndChange + conj.Conjugation
		}

		adjs = append(adjs, adjconj)
	}
	return adjs
}

func (adj *NAAdjective) Conjugate() []AdjConjugation {
	var adjs []AdjConjugation
	for _, conj := range naAdjConjugations {
		adjconj := conj
		adjconj.Conjugated = adj.Stem + conj.Conjugation
		adjs = append(adjs, adjconj)
	}
	return adjs
}

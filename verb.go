package jp

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
)

type VerbSimpleConjugation struct {
	KForm, Conjugated string
}

type VerbComplexConjugation struct {
	CForm, Politeness, Polarity, Conjugated string
}

type japkatResult struct {
	Basic   []VerbSimpleConjugation
	Complex []VerbComplexConjugation
}

type Verb struct {
	Stem, Raw string
}

func NewVerb(verb string) Verb {
	return Verb{Raw: verb}
}

func (verb *Verb) Conjugate() (error, []VerbSimpleConjugation, []VerbComplexConjugation) {
	cmd := exec.Command("JapKatsuyouCLI", verb.Raw)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	result := japkatResult{}
	err = json.Unmarshal(out.Bytes(), &result)
	if err != nil {
		log.Fatal(err)
	}

	verb.Stem = getStem(result.Complex)

	return err, result.Basic, result.Complex
}

func getStem(conjugations []VerbComplexConjugation) string {
	for _, value := range conjugations {
		if value.CForm == "Present" && value.Politeness == "Polite" && value.Polarity == "Affirmative" {
			// to runes
			runes := []rune(value.Conjugated)
			// remove masu (ます) and return stem
			return string(runes[:len(runes)-2])
		}
	}
	return ""
}

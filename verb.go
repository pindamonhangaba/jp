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
	basic   []VerbSimpleConjugation
	complex []VerbComplexConjugation
}

type Verb struct {
	Raw string
}

func NewVerb(verb string) Verb {
	return Verb{Raw: verb}
}

func (verb *Verb) conjugate() (error, []VerbSimpleConjugation, []VerbComplexConjugation) {
	cmd := exec.Command("JapKatsuyouCLI", Verb.Raw)
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

	return err, result.basic, result.complex
}

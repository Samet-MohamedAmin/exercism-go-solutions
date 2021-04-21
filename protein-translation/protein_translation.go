package protein

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

const (
	mappingFile = "map.json"
	proteinStop = "STOP"
)

type protein struct {
	Name   string
	Codons []string
}

type revMap map[string]string

var (
	// ErrStop is error stop
	ErrStop = errors.New("stop error")

	// ErrInvalidBase is invalid base error
	ErrInvalidBase = errors.New("invalid base error")

	mapping, _    = readMapping(mappingFile)
	revMapping, _ = genRevMapping(mapping)
)

func readMapping(filename string) (out []protein, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &out)
	return
}

func genRevMapping(ps []protein) (out revMap, err error) {
	out = revMap{}
	for _, p := range ps {
		for _, c := range p.Codons {
			out[c] = p.Name
		}
	}
	return
}

// FromCodon returns an error if protein type is invalid. If not returns protein type
func FromCodon(in string) (out string, err error) {
	switch {
	case len(in) == 0:
		return
	case len(in) < 3:
		err = ErrInvalidBase
		return
	}

	name, ok := revMapping[in[:3]]
	switch {
	case !ok:
		err = ErrInvalidBase
		return
	case name == proteinStop:
		err = ErrStop
		return
	}

	out = name
	return
}

// FromRNA returns an error if there is an invalid RNA sequence. If not returns protein sequence
func FromRNA(in string) (out []string, err error) {
	if len(in) == 0 {
		return
	}

	n, errCondon := FromCodon(in)
	switch errCondon {
	case ErrInvalidBase:
		err = errCondon
		return
	case ErrStop:
		return
	}

	ns, err := FromRNA(in[3:])
	out = append([]string{n}, ns...)
	return
}

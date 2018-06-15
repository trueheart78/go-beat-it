package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dataRecord struct {
	args  []string
	value string
	valid bool
}

func TestCmdName(t *testing.T) {
	var name string
	var err error

	cachedArgs := os.Args
	defer func() { os.Args = cachedArgs }()

	for _, d := range dataForTests() {
		os.Args = d.args
		name, err = Name()
		if d.valid {
			assert.Equal(t, d.value, name)
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func dataForTests() (ds []dataRecord) {
	ds = append(ds, dataRecord{args: []string{}, value: "", valid: false})
	ds = append(ds, dataRecord{args: []string{"mario"}, value: "mario", valid: true})
	ds = append(ds, dataRecord{args: []string{"dark", "souls", "3"}, value: "dark souls 3", valid: true})
	ds = append(ds, dataRecord{args: []string{"legend", "of", "zelda", "\t"}, value: "legend of zelda", valid: true})
	ds = append(ds, dataRecord{args: []string{"\tlegend", "of", "zelda", "\t"}, value: "legend of zelda", valid: true})
	return
}

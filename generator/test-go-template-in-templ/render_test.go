package testgotemplates

import (
	_ "embed"
	"testing"

	"github.com/a-h/templ/generator/htmldiff"
)

//go:embed expected.html
var expected string

func TestExample(t *testing.T) {
	component := Example()
	diff, err := htmldiff.Diff(component, expected)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}

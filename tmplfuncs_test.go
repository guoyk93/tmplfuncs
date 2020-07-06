package tmplfuncs

import (
	"bytes"
	"testing"
	"text/template"
)

func runTemplate(t *testing.T, tplStr string, data interface{}) string {
	var err error
	var tpl *template.Template
	if tpl, err = template.New("").Funcs(Funcs).Parse(tplStr); err != nil {
		t.Fatal(err)
		return ""
	}
	buf := &bytes.Buffer{}
	if err = tpl.Execute(buf, data); err != nil {
		t.Fatal(err)
		return ""
	}
	return buf.String()
}

func TestFuncs(t *testing.T) {
	data := map[string]interface{}{
		"Hello": "hello",
	}
	type testCase struct {
		tmpl string
		ret  string
	}
	testCases := []testCase{
		{
			tmpl: `{{ stringsContains .Hello "lo" }}`,
			ret:  "true",
		},
		{
			tmpl: `{{ if stringsContains .Hello "lo" }}hello{{end}}`,
			ret:  "hello",
		},
		{
			tmpl: `{{ range stringsFields "   he o  " }}x{{.}}{{end}}`,
			ret:  "xhexo",
		},
		{
			tmpl: `{{ stringsHasPrefix .Hello "he" }}`,
			ret:  "true",
		},
		{
			tmpl: `{{ stringsHasSuffix .Hello "lo" }}`,
			ret:  "true",
		},
		{
			tmpl: `{{ stringsRepeat .Hello 3 }}`,
			ret:  "hellohellohello",
		},
		{
			tmpl: `{{ stringsReplaceAll .Hello "lo" "ll" }}`,
			ret:  "helll",
		},
		{
			tmpl: `{{ stringsIndex .Hello "l" }}`,
			ret:  "2",
		},
		{
			tmpl: `{{ stringsLastIndex .Hello "l" }}`,
			ret:  "3",
		},
		{
			tmpl: `{{ range stringsSplit .Hello "ll" }}x{{.}}{{end}}`,
			ret:  "xhexo",
		},
		{
			tmpl: `{{ stringsToLower "HeLlo"}}`,
			ret:  "hello",
		},
		{
			tmpl: `{{ stringsToUpper "HeLlo"}}`,
			ret:  "HELLO",
		},
		{
			tmpl: `{{ stringsTrimPrefix .Hello "he"}}`,
			ret:  "llo",
		},
		{
			tmpl: `{{ stringsTrimSpace " hello "}}`,
			ret:  "hello",
		},
		{
			tmpl: `{{ stringsTrimSuffix .Hello "lo"}}`,
			ret:  "hel",
		},
		{
			tmpl: `{{ netResolveIPAddr "ip4" "localhost" }}`,
			ret:  "127.0.0.1",
		},
		{
			tmpl: `{{ k8sStatefulSetID }}`,
			ret:  "1",
		},
	}

	for _, testCase := range testCases {
		ret := runTemplate(t, testCase.tmpl, data)
		if ret != testCase.ret {
			t.Errorf("failed: %s, %s != %s", testCase.tmpl, ret, testCase.ret)
		}
	}
}

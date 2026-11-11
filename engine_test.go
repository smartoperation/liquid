package liquid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

var emptyBindings = map[string]interface{}{}

// There's a lot more tests in the filters and tags sub-packages.
// This collects a minimal set for testing end-to-end.
var liquidTests = []struct{ in, expected string }{
	{`{{ page.title }}`, "Introduction"},
	{`{% if x %}true{% endif %}`, "true"},
	{`{{ "upper" | upcase }}`, "UPPER"},
}

var testBindings = map[string]interface{}{
	"x":  123,
	"ar": []string{"first", "second", "third"},
	"page": map[string]interface{}{
		"title": "Introduction",
	},
}

func TestEngine_ParseAndRenderString(t *testing.T) {
	engine := NewEngine()
	for i, test := range liquidTests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			out, err := engine.ParseAndRenderString(test.in, testBindings)
			require.NoErrorf(t, err, test.in)
			require.Equalf(t, test.expected, out, test.in)
		})
	}
}

func TestEngine_ParseAndRenderString_ptr_to_hash(t *testing.T) {
	params := map[string]interface{}{
		"message": &map[string]interface{}{
			"Text": "hello",
			"jsonNumber" : json.Number("123"),
		},
	}
	engine := NewEngine()
	template := "{{ message.Text }} {{message.jsonNumber}}"
	str, err := engine.ParseAndRenderString(template, params)
	require.NoError(t, err)
	require.Equal(t, "hello 123", str)
}

type testStruct struct{ Text string }

func TestEngine_ParseAndRenderString_struct(t *testing.T) {
	params := map[string]interface{}{
		"message": testStruct{
			Text: "hello",
		},
	}
	engine := NewEngine()
	template := "{{ message.Text }}"
	str, err := engine.ParseAndRenderString(template, params)
	require.NoError(t, err)
	require.Equal(t, "hello", str)
}

func TestEngine_ParseAndRender_errors(t *testing.T) {
	_, err := NewEngine().ParseAndRenderString("{{ syntax error }}", emptyBindings)
	require.Error(t, err)
	_, err = NewEngine().ParseAndRenderString("{% if %}", emptyBindings)
	require.Error(t, err)
	_, err = NewEngine().ParseAndRenderString("{% undefined_tag %}", emptyBindings)
	require.Error(t, err)
	_, err = NewEngine().ParseAndRenderString("{% a | undefined_filter %}", emptyBindings)
	require.Error(t, err)
}

var disableTagsTests = []struct{ in, expected string }{
	{`{% include "LICENSE" %}`, `{% include "LICENSE" %}`},
	{`{% if true %}a{% else %}b{% endif %}`, `{% if true %}a{% else %}b{% endif %}`},
	{`{% for a in ar %}{{ a }}{% endfor %}`, `{% for a in ar %}{% endfor %}`},
	{`{%- assign y = 1 -%}`, `{%- assign y = 1 -%}`},
	{`{% undefined_tag %}`, `{% undefined_tag %}`},
	{`{% if %}`, `{% if %}`},
	{`{% {{ x }} %}`, `{% 123 %}`},
	{`{{ page.title }}`, "Introduction"},
	{`{{ "upper" | upcase }}`, "UPPER"},
}

func TestEngine_DisableTags(t *testing.T) {
	engine := NewEngine().DisableTags()
	for i, test := range disableTagsTests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			out, err := engine.ParseAndRenderString(test.in, testBindings)
			require.NoErrorf(t, err, test.in)
			require.Equalf(t, test.expected, out, test.in)
		})
	}
}

func TestEngine_DisableTags_delims(t *testing.T) {
	engine := NewEngine().Delims("OBJ(", ")OBJ", "TAG(", ")TAG").DisableTags()
	out, err := engine.ParseAndRenderString(`TAG( if true )TAGaTAG( endif )TAG OBJ( x )OBJ`, testBindings)
	require.NoError(t, err)
	require.Equal(t, `TAG( if true )TAGaTAG( endif )TAG 123`, out)
}

func TestEngine_withoutDisableTags(t *testing.T) {
	out, err := NewEngine().ParseAndRenderString(`{% include "LICENSE" %}`, testBindings)
	require.NoError(t, err)
	require.Contains(t, out, "MIT License")
}

func BenchmarkEngine_Parse(b *testing.B) {
	engine := NewEngine()
	buf := new(bytes.Buffer)
	for i := 0; i < 1000; i++ {
		_, err := io.WriteString(buf, `if{% if true %}true{% elsif %}elsif{% else %}else{% endif %}`)
		require.NoError(b, err)
		_, err = io.WriteString(buf, `loop{% for item in array %}loop{% break %}{% endfor %}`)
		require.NoError(b, err)
		_, err = io.WriteString(buf, `case{% case value %}{% when a %}{% when b %{% endcase %}`)
		require.NoError(b, err)
		_, err = io.WriteString(buf, `expr{{ a and b }}{{ a add: b }}`)
		require.NoError(b, err)
	}
	s := buf.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := engine.ParseTemplate(s)
		require.NoError(b, err)
	}
}

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extractSnippet(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		t.Run("without order", func(t *testing.T) {
			assert := assert.New(t)

			input := `#include <iostream>
#include <string>
using namespace std;

/* @snippet:hello */
string greet(const string &s) {
	return "hello " + s;
}
/* @endsnippet */

int main() {
	cout << greet("world") << "\n";
}
`

			expectedCode := `string greet(const string &s) {
	return "hello " + s;
}`

			snippet, ok := extractSnippet(input)
			assert.True(ok)
			assert.Equal(&Snippet{Title: "hello", Order: 0, Code: expectedCode}, snippet)
		})

		t.Run("with order", func(t *testing.T) {
			assert := assert.New(t)

			input := `#include <iostream>
#include <string>
using namespace std;

/* @snippet:hello */
/* @order:123 */
string greet(const string &s) {
	return "hello " + s;
}
/* @endsnippet */

int main() {
	cout << greet("world") << "\n";
}
`

			expectedCode := `string greet(const string &s) {
	return "hello " + s;
}`

			snippet, ok := extractSnippet(input)
			assert.True(ok)
			assert.Equal(&Snippet{Title: "hello", Order: 123, Code: expectedCode}, snippet)
		})
	})

	t.Run("no match", func(t *testing.T) {
		assert := assert.New(t)

		snippet, ok := extractSnippet("no match")
		assert.True(!ok)
		assert.Equal(&Snippet{Title: "", Order: 0, Code: ""}, snippet)
	})
}

func Test_writeXML(t *testing.T) {
	assert := assert.New(t)

	input := `int main() {
	cout << "$" << endl;
}`

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<CodeSnippets xmlns="http://schemas.microsoft.com/VisualStudio/2005/CodeSnippet">
	<CodeSnippet Format="1.0.0">
		<Header>
			<Title>hello</Title>
		</Header>
		<Snippet>
			<Code Language="CPP"><![CDATA[int main() {
	cout << "$$" << endl;
}]]></Code>
		</Snippet>
	</CodeSnippet>
</CodeSnippets>`

	actual := writeXML("hello", input)
	assert.Equal(expected, actual)
}

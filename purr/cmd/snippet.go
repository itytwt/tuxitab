package cmd

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// SnippetCmd is the object of `snippet` command settings.
type SnippetCmd struct {
	Config
	cmd *cobra.Command
}

// NewSnippetCmd creates a new `SnippetCmd` object.
func NewSnippetCmd(cfg Config) *SnippetCmd {
	ret := &SnippetCmd{
		Config: cfg,
	}

	ret.cmd = &cobra.Command{
		Use:   "snippet",
		Short: "Generate visual studio snippets",
		Args:  cobra.ExactArgs(1),
		Run:   ret.execute,
	}

	return ret
}

func (sc *SnippetCmd) execute(cmd *cobra.Command, args []string) {
	source := readSource(args[0])
	dict := map[string]Snippets{}

	for _, src := range source {
		code, err := ioutil.ReadFile(src)
		if err != nil {
			panic(err)
		}

		if cs, ok := extractSnippet(string(code)); ok {
			dict[cs.Title] = append(dict[cs.Title], cs)
		}
	}

	for title, list := range dict {
		sort.Sort(list)

		ol := []string{}
		for _, s := range list {
			ol = append(ol, s.Code)
		}

		xml := writeXML(title, strings.Join(ol, "\n"))
		outputToFile(sc.OutputFolder, title, xml)
	}

}

// Cmd of `SnippetCmd` returns its `cmd` field.
func (sc *SnippetCmd) Cmd() *cobra.Command {
	return sc.cmd
}

func readSource(root string) []string {
	ret := []string{}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if !info.IsDir() && isCppSource(path) {
			absPath, err := filepath.Abs(path)
			if err != nil {
				panic(err)
			}
			ret = append(ret, absPath)
		}

		return nil
	})

	return ret
}

func isCppSource(path string) bool {
	return strings.HasSuffix(path, ".cpp") || strings.HasSuffix(path, ".h")
}

// Snippet is a piece of extracted snippet with title and order information.
type Snippet struct {
	Title string
	Order int
	Code  string
}

// Snippets is a list of `Snippet`.
type Snippets []*Snippet

func (s Snippets) Len() int           { return len(s) }
func (s Snippets) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Snippets) Less(i, j int) bool { return s[i].Order < s[j].Order }

func extractSnippet(code string) (*Snippet, bool) {
	reg := regexp.MustCompile("\\/\\*\\s*@snippet:(\\w+)\\s*\\*\\/\\s*(?:\\/\\*\\s*@order:(\\w+)\\s*\\*\\/\\s*)?([\\s\\S]*?)\\s*\\/\\*\\s*@endsnippet\\s*\\*\\/")
	if res := reg.FindStringSubmatch(code); res != nil {
		if res[2] == "" {
			return &Snippet{Title: res[1], Code: res[3]}, true
		}
		order, err := strconv.Atoi(res[2])
		if err != nil {
			panic(err)
		}
		return &Snippet{Title: res[1], Order: order, Code: res[3]}, true
	}
	return &Snippet{Title: "", Order: 0, Code: ""}, false
}

// XMLCodeSnippets is the struct of a snippet object.
type XMLCodeSnippets struct {
	XMLName     xml.Name       `xml:"CodeSnippets"`
	XMLNS       string         `xml:"xmlns,attr"`
	CodeSnippet XMLCodeSnippet `xml:"CodeSnippet"`
}

// XMLCodeSnippet is the nested content of `XMLCodeSnippets`.
type XMLCodeSnippet struct {
	Format string  `xml:",attr"`
	Title  string  `xml:"Header>Title"`
	Code   XMLCode `xml:"Snippet>Code"`
}

// XMLCode is the actual content part.
type XMLCode struct {
	Language string `xml:",attr"`
	Data     string `xml:",cdata"`
}

func writeXML(title, snippet string) string {
	cs := XMLCodeSnippets{
		XMLNS: "http://schemas.microsoft.com/VisualStudio/2005/CodeSnippet",
		CodeSnippet: XMLCodeSnippet{
			Format: "1.0.0",
			Title:  title,
			Code: XMLCode{
				Language: "CPP",
				Data:     strings.Replace(snippet, "$", "$$", -1),
			},
		},
	}

	result, err := xml.MarshalIndent(cs, "", "\t")
	if err != nil {
		panic(err)
	}

	return xml.Header + string(result)
}

func outputToFile(outfolder, title, xml string) {
	path := outfolder
	err := os.MkdirAll(path, os.ModeDir)
	if err != nil {
		panic(err)
	}

	fout, err := os.Create(fmt.Sprintf((path + "/%v.snippet"), title))
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	if _, err := fout.WriteString(xml); err != nil {
		panic(err)
	}
}

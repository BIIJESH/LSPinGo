package analysis

import (
	"educationalsp/lsp"
	"fmt"
	"strings"
)

type State struct {
	//Map of filenames to contents
	//user Defined type
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}
func (s *State) OpenDocument(uri, text string) { //method no func name
	s.Documents[uri] = text
}
func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

//	func Use() {
//		mystate := NewState()
//		fmt.Println("this is state:", mystate)
//	}
func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// In real life, this would look up the type in our type analysis code...

	document := s.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Characters: %d", uri, len(document)),
		},
	}
}
func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	// In real life, this would look up the type in our type analysis code...

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				}, End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				}},
		},
	}
}
func (s *State) TextDocumentCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
	// In real life, this would look up the type in our type analysis code...
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "Neovim",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS C*de with a superior editor",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C*de",
				},
			}
			actions = append(actions, lsp.CodeAction{
				Title: "Censor to VS C*de",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}
	}

	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}

	return response
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim (btw)",
			Detail:        "Extremely good editor ",
			Documentation: "Fun to watch videos.",
		},
	}
	//ask your static analysis tool to provide completion
	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
	return response
}
func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}

}

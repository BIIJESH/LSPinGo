package analysis

import (
	"educationalsp/lsp"
	"fmt"
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

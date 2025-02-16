package lsp

type DidOpenTextDocumentNotification struct {
	Notification
	Params DidOpenTextDocument `json:"params"`
}
type DidOpenTextDocument struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

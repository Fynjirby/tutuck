module github.com/fynjirby/tutuck/client/cli

go 1.25.0

require (
	github.com/chzyer/readline v1.5.1
	github.com/fynjirby/tutuck/client/core v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
)

replace github.com/fynjirby/tutuck/client/core => ../core

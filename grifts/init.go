package grifts

import (
	"github.com/teten-nugraha/go_buffalo_restapi/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}

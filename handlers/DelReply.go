package handlers

import (
	"toropress/libs"
	"toropress/models"
)

type DeleteReplyHandler struct {
	libs.RootAuthHandler
}

func (self *DeleteReplyHandler) Get() {
	rid, _ := self.GetInt64(":rid")
	models.DelReply(rid)
	self.Ctx.Redirect(302, "/")
}

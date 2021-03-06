package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type INode interface {
	Eval() string
	Text() string
	SetParent(n INode)
	GetParent() INode
}
type Node struct {
	Text_    string
	Children []INode
	parent   INode
}

func (this Node) Eval() string {
	return ""
}
func (this Node) Text() string {
	return this.Text_
}
func (this *Node) SetParent(n INode) {
	if this.parent == nil {
		this.parent = n
	}
}
func (this *Node) GetParent() INode {
	return this.parent
}
func NewNode(ctx antlr.RuleContext) *Node {
	return &Node{Text_: ctx.GetText()}
}
func NewNodeNoCtx() *Node {
	return &Node{}
}

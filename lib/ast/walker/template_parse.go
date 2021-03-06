package walker

import (
	"github.com/champii/og/lib/ast"
)

type TemplateParse struct {
	AstWalker
	Root      ast.INode
	Templates *Templates
}

func (this *TemplateParse) StructType(n ast.INode) ast.INode {
	structType := n.(*ast.StructType)
	if structType.TemplateSpec != nil {
		types := []string{}
		for _, t := range structType.TemplateSpec.Result.Types {
			types = append(types, t.Eval())
		}
		this.Templates.Add(structType.Name, NewTemplate(structType.Name, types, structType))
	}
	return n
}
func (this *TemplateParse) Signature(n ast.INode) ast.INode {
	sig := n.(*ast.Signature)
	if sig.TemplateSpec != nil {
		if f, ok := sig.GetParent().(*ast.Function); ok {
			fDecl := f.GetParent().(*ast.FunctionDecl)
			types := []string{}
			for _, t := range sig.TemplateSpec.Result.Types {
				types = append(types, t.Eval())
			}
			this.Templates.Add(fDecl.Name, NewTemplate(fDecl.Name, types, fDecl))
		}
	}
	return n
}
func RunTemplateParse(tree ast.INode, templates *Templates) {
	templateParse := TemplateParse{
		Root:      tree,
		Templates: templates,
	}
	templateParse.type_ = &templateParse
	templateParse.Walk(tree)
}

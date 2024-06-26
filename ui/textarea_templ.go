// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"

const classTextarea = "flex min-h-[60px] w-full rounded-md border border-input bg-transparent px-3 py-2 text-sm shadow-sm placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50"

func Textarea(opts ...func(*templ.Attributes)) templ.Attributes {
	attrs := templ.Attributes{
		"class": classTextarea,
	}
	for _, o := range opts {
		o(&attrs)
	}
	return attrs
}

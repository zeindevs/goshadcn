// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"

const classButton = "inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 data-[disabled=true]:opacity-50"

const (
	classButtonVariantDefault     = "bg-primary text-primary-foreground shadow hover:bg-primary/90 rounded-md"
	classButtonVariantDestructive = "bg-destructive text-destructive-foreground shadow-sm hover:bg-destructive/90 rounded-md"
	classButtonVariantSecondary   = "bg-secondary text-secondary-foreground shadow-sm hover:bg-secondary/80 rounded-md"
	classButtonVariantOutline     = "border border-input bg-background shadow-sm hover:bg-accent hover:text-accent-foreground rounded-md"
	classButtonVariantGhost       = "hover:bg-accent hover:text-accent-foreground rounded-md"
	classButtonVariantLink        = "text-primary underline-offset-4 hover:underline rounded-md"
	classButtonVariantIcon        = "hover:bg-accent rounded-md"
)

const (
	classButtonSizeDefault = "h-9 px-4 py-2"
	classButtonSizeSmall   = "h-8 rounded-md px-3 text-xs"
	classButtonSizeLarge   = "h-10 rounded-md px-8"
	classButtonSizeIcon    = "h-9 w-9"
)

func Button(opts ...func(*templ.Attributes)) templ.Attributes {
	attrs := templ.Attributes{
		"class": classButton + " " + classButtonSizeDefault + " " + classButtonVariantDefault,
	}
	for _, o := range opts {
		o(&attrs)
	}
	return attrs
}

func ButtonVariant(v string, s ...string) func(*templ.Attributes) {
	return func(attrs *templ.Attributes) {
		btnsize := classButtonSizeDefault
		for _, bs := range s {
			switch bs {
			case "primary":
				btnsize = classButtonSizeDefault
			case "sm":
				btnsize = classButtonSizeSmall
			case "lg":
				btnsize = classButtonSizeLarge
			case "icon":
				btnsize = classButtonSizeIcon
			default:
				btnsize = classButtonSizeDefault
			}
		}
		attr := *attrs
		switch v {
		case "primary":
			attr["class"] = classButton + " " + btnsize + " " + classButtonVariantDefault
			break
		case "destructive":
			attr["class"] = classButton + " " + btnsize + " " + classButtonVariantDestructive
			break
		case "secondary":
			attr["class"] = classButton + " " + btnsize + " " + classButtonVariantSecondary
			break
		case "ghost":
			attr["class"] = classButton + " " + btnsize + " " + classButtonVariantGhost
			break
		case "outline":
			attr["class"] = classButton + " " + btnsize + " " + classButtonVariantOutline
			break
		case "link":
			attr["class"] = classButton + " " + btnsize + " " + classButtonVariantLink
			break
		case "icon":
			attr["class"] = attr["class"].(string) + " " + btnsize + " " + classButtonVariantIcon
			break
		}
	}
}

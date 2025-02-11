// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package shared

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Navbar() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"bg-gray-800 text-white py-4 px-4\"><nav class=\"container flex items-center justify-between\"><div class=\"text-xl font-semibold\">Shubham Sharma</div><ul class=\"flex space-x-4\"><li><a hx-get=\"/hero\" hx-target=\"#index\" hx-swap=\"innerHTML\" class=\"hover:text-yellow-500\">Home</a></li><li><a hx-get=\"/about\" hx-target=\"#index\" hx-swap=\"innerHTML\" class=\"hover:text-yellow-500\">About</a></li><li><a hx-get=\"/projects\" hx-target=\"#index\" hx-swap=\"innerHTML\" class=\"hover:text-yellow-500\">Projects</a></li></ul></nav></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

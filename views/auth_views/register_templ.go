// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package auth_views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/kyp0717/garage/views"

	"github.com/gofiber/fiber/v2"
)

func RegisterIndex(fromProtected bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<section class=\"card w-fit bg-base-200 shadow-xl mx-auto mb-8\"><div class=\"card-body pb-2\"><h1 class=\"card-title border-b border-b-slate-600 pb-[4px]\">Register User</h1><form hx-swap=\"transition:true\" class=\"rounded-xl drop-shadow-xl flex flex-col gap-4 w-96 p-8\" action=\"\" method=\"post\"><label class=\"flex flex-col justify-start gap-2\">Email: <input class=\"input input-bordered input-primary bg-slate-800\" type=\"email\" name=\"email\" required autofocus")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if fromProtected {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if fromProtected {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, " value=\"disabled \"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "></label> <label class=\"flex flex-col justify-start gap-2 relative\">Password: <input class=\"input input-bordered input-primary bg-slate-800\" type=\"password\" name=\"password\" required minlength=\"6\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if fromProtected {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if fromProtected {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, " value=\"disabled \"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "> <button title=\"View password\" type=\"button\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if fromProtected {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, " class=\"absolute top-12 right-3\" _=\"on click if [type of previous &lt;input/&gt;] == &#39;password&#39; then remove [@type=password] from previous &lt;input/&gt; then hide #eye then remove .hidden from #eye-slash else show #eye then add .hidden to #eye-slash then tell previous &lt;input/&gt; toggle [@type=password] end\"><svg id=\"eye\" xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" fill=\"currentColor\" viewBox=\"0 0 16 16\"><path d=\"M10.5 8a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0\"></path> <path d=\"M0 8s3-5.5 8-5.5S16 8 16 8s-3 5.5-8 5.5S0 8 0 8m8 3.5a3.5 3.5 0 1 0 0-7 3.5 3.5 0 0 0 0 7\"></path></svg> <svg id=\"eye-slash\" class=\"hidden\" xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" fill=\"currentColor\" viewBox=\"0 0 16 16\"><path d=\"m10.79 12.912-1.614-1.615a3.5 3.5 0 0 1-4.474-4.474l-2.06-2.06C.938 6.278 0 8 0 8s3 5.5 8 5.5a7 7 0 0 0 2.79-.588M5.21 3.088A7 7 0 0 1 8 2.5c5 0 8 5.5 8 5.5s-.939 1.721-2.641 3.238l-2.062-2.062a3.5 3.5 0 0 0-4.474-4.474z\"></path> <path d=\"M5.525 7.646a2.5 2.5 0 0 0 2.829 2.829zm4.95.708-2.829-2.83a2.5 2.5 0 0 1 2.829 2.829zm3.171 6-12-12 .708-.708 12 12z\"></path></svg></button></label> <label class=\"flex flex-col justify-start gap-2\">Username: <input class=\"input input-bordered input-primary bg-slate-800\" type=\"text\" name=\"username\" required minlength=\"4\" maxlength=\"64\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if fromProtected {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if fromProtected {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, " value=\"disabled \"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "></label><footer class=\"card-actions justify-end\"><button class=\"badge badge-primary px-6 py-4 hover:scale-[1.1]\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if fromProtected {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, ">Register User</button></footer></form></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func Register(
	page string,
	fromProtected, isError bool,
	msg fiber.Map,
	cmp templ.Component,
) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var3 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = cmp.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = views.Layout(page, fromProtected, isError, msg, "").Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate

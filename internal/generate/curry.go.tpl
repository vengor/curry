{{- /*gotype: github.com/vengor/curry/internal/generate.Data*/ -}}
package {{ .PackageName }}

{{ range .Funcs }}
// {{ .Name }} curries the first {{ len .CurriedArgIndexes }} arguments from a function with
{{- if .InputArguments }} {{ .InputArguments }}{{ end }}
{{- if .Variadic}}{{ if .InputArguments}} plus{{ end }} variadic{{ end }} arguments and {{ len .RetIndexes }} return values.
func {{ .Name }}[
{{- range $i, $idx := .CurriedArgIndexes }}{{- if $i }}, {{ end }}CurriedArg{{ $idx }}{{- end }}
{{- range $i, $idx := .ArgIndexes }}, Arg{{ $idx }}{{ end }}
{{- if .Variadic }}, Var{{ end -}}
{{- range $i, $idx := .RetIndexes }}, Ret{{ $idx }}{{ end }} any](
{{- range $i, $idx := .CurriedArgIndexes }}{{- if $i }}, {{ end }}c{{ $idx }} CurriedArg{{ $idx }}{{- end }}, f func(
{{- range $i, $idx := .CurriedArgIndexes }}{{- if $i }}, {{ end }}CurriedArg{{ $idx }}{{- end }}
{{- range $i, $idx := .ArgIndexes }}, Arg{{ $idx }}{{ end }}
{{- if .Variadic }}, ...Var{{ end -}})
{{- if .RetIndexes -}}
({{- range $i, $idx := .RetIndexes }}{{- if $i }}, {{ end }}Ret{{ $idx }}{{ end }})
{{- end }}) func(
{{- range $i, $idx := .ArgIndexes }}{{- if $i }}, {{ end }} Arg{{ $idx }}{{ end }}
{{- if .Variadic }}{{ if .ArgIndexes }}, {{ end }}...Var{{ end -}})
{{- if .RetIndexes -}}
({{- range $i, $idx := .RetIndexes }}{{- if $i }}, {{ end }}Ret{{ $idx }}{{ end }})
{{- end }} {
	return func(
        {{- range $i, $idx := .ArgIndexes }}{{- if $i }}, {{ end }}a{{ $idx}} Arg{{ $idx }}{{ end }}
        {{- if .Variadic }}{{ if .ArgIndexes }}, {{ end }}v ...Var{{ end -}}) (
        {{- range $i, $idx := .RetIndexes }}{{- if $i }}, {{ end }}Ret{{ $idx }}{{ end }}) {
        {{ if .RetIndexes }}return {{ end -}} f(
                {{- range $i, $idx := .CurriedArgIndexes }}{{- if $i }}, {{ end }}c{{ $idx }}{{- end -}}
                {{- range $i, $idx := .ArgIndexes }}, a{{ $idx}}{{ end -}}
                {{- if .Variadic }}, v...{{ end -}}
			)
    }
}
{{ end }}
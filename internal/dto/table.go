package dto

const Tmpl = `// Code generated by sql2struct. https://github.com/gangming/sql2struct
package {{.Package}}

{{if .ContainsTimeField}}import "time" {{end}}

//{{ .UpperCamelCaseName}} {{.Table.Comment}}
type {{ .UpperCamelCaseName}} struct {
{{range .Fields}} {{.UpperCamelCaseName}}  {{.Type}} {{.Tag}} {{if .Comment}} // {{.Comment}} {{end}}
{{end}}}

// TableName the name of table in database
func (t *{{.UpperCamelCaseName}}) TableName() string {
    return "{{.Table.Name}}"
}
`

type Table struct {
	Name               string  `sql:"name"`
	UpperCamelCaseName string  `sql:"upper_camel_case_name"`
	Comment            string  `sql:"comment"`
	Fields             []Field `sql:"fields"`
	ContainsTimeField  bool    `sql:"contains_time"`
}
type Field struct {
	IsPK               bool   `json:"is_pk"`
	Name               string `json:"name"`
	UpperCamelCaseName string `json:"upper_camel_case_name"`
	Type               string `json:"type"`
	Comment            string `json:"comment"`
	DefaultValue       string `json:"default_value"`
	Tag                string `json:"tag"`
}
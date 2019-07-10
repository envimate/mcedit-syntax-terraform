package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/helper/schema"
)

type terms struct {
	Keywords     map[string]bool
	KeywordColor string

	Resources     map[string]bool
	ResourceColor string

	DataSources     map[string]bool
	DataSourceColor string

	Providers     map[string]bool
	ProviderColor string

	Functions     map[string]bool
	FunctionColor string

	VarColor string
}

var t = &terms{
	Keywords:    make(map[string]bool),
	Resources:   make(map[string]bool),
	DataSources: make(map[string]bool),
	Providers:   make(map[string]bool),
	Functions:   make(map[string]bool),
}

func scan(r *schema.Resource, out map[string]bool) {
	for k, v := range r.Schema {
		if v.Elem != nil {
			c, ok := v.Elem.(*schema.Resource)
			if ok {
				scan(c, out)
			}
		}
		out[k] = true
	}
}

func main() {
	var (
		flags    *flag.FlagSet
		uc       string
		provider *schema.Provider
		tpl      *template.Template
		err      error
	)
	flags = flag.NewFlagSet("colors", flag.ExitOnError)

	flags.StringVar(&t.KeywordColor, "keyword", colorWhite, "color of keywords")
	flags.StringVar(&t.ResourceColor, "resource", colorBrightCyan, "color of resources")
	flags.StringVar(&t.DataSourceColor, "data-source", colorBrightCyan, "color of data sources")
	flags.StringVar(&t.ProviderColor, "provider", colorRed, "color of providers")
	flags.StringVar(&t.FunctionColor, "func", colorYellow, "color of functions")
	flags.StringVar(&t.VarColor, "var", colorYellow, "color of variables")

	flags.Parse(os.Args[1:])

	if uc = colors.exists(t.KeywordColor, t.ResourceColor, t.DataSourceColor, t.ProviderColor, t.FunctionColor, t.VarColor); uc != "" {
		fmt.Printf("unknown color %s\n", uc)
		flags.PrintDefaults()
		os.Exit(2)
	}

	provider = aws.Provider().(*schema.Provider)

	for k, v := range provider.ResourcesMap {
		scan(v, t.Keywords)
		t.Resources[k] = true
	}

	for k, v := range provider.DataSourcesMap {
		scan(v, t.Keywords)
		t.DataSources[k] = true
	}

	for k := range provider.Schema {
		t.Keywords[k] = true
	}

	for k := range config.Funcs() {
		t.Functions[k] = true
	}

	t.Providers["aws"] = true

	var addKeywords = []string{
		"private_key",
		"script_path",
		"bastion_user",
		"bastion_password",
		"bastion_private_key",
		"bastion_host",
		"bastion_port",
		"connection",
		"provisioner",
		"inline",
	}

	for _, k := range addKeywords {
		t.Keywords[k] = true
	}

	tpl = template.New("template")

	tpl, err = tpl.ParseFiles("template")

	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, t)

	if err != nil {
		panic(err)
	}
}

// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path"
	"strings"
	"text/template"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/provider/generators/resource/codegen"
	"github.com/mitchellh/cli"
)

var (
	cfTypeSchemaFile = flag.String("cfschema", "", "CloudFormation resource type schema file; required")
	packageName      = flag.String("package", "", "override package name for generated code")
	tfResourceType   = flag.String("resource", "", "Terraform resource type; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -resource <TF-resource-type> -cfschema <CF-type-schema-file> <generated-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 || *tfResourceType == "" || *cfTypeSchemaFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	destinationPackage := os.Getenv("GOPACKAGE")
	if *packageName != "" {
		destinationPackage = *packageName
	}

	filename := args[0]

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	generator := NewGenerator(ui, *tfResourceType, *cfTypeSchemaFile)

	if err := generator.Generate(destinationPackage, filename); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform %s resource: %s", *tfResourceType, err))
		os.Exit(1)
	}
}

type Generator struct {
	cfTypeSchemaFile string
	tfResourceType   string
	ui               cli.Ui
}

func NewGenerator(ui cli.Ui, tfResourceType, cfTypeSchemaFile string) *Generator {
	return &Generator{
		cfTypeSchemaFile: cfTypeSchemaFile,
		tfResourceType:   tfResourceType,
		ui: &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
	}
}

func (g *Generator) Infof(format string, a ...interface{}) {
	g.ui.Info(fmt.Sprintf(format, a...))
}

// Generate generates the resource's type factory into the specified file.
func (g *Generator) Generate(packageName, filename string) error {
	g.Infof("generating Terraform resource code for %q from %q into %q", g.tfResourceType, g.cfTypeSchemaFile, filename)

	// Create target directory.
	dirname := path.Dir(filename)
	err := os.MkdirAll(dirname, 0755)

	if err != nil {
		return fmt.Errorf("error creating target directory %s: %w", dirname, err)
	}

	resource, err := NewResource(g.tfResourceType, g.cfTypeSchemaFile)

	if err != nil {
		return fmt.Errorf("error reading CloudFormation resource schema for %s: %w", g.tfResourceType, err)
	}

	cfTypeName := *resource.CfResource.TypeName
	// https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html#schema-properties-typeName
	parts := strings.Split(cfTypeName, "::")

	if len(parts) != 3 || len(parts[2]) < 2 {
		return fmt.Errorf("incorrect format for CloudFormation Resource Provider Schema type name: %s", cfTypeName)
	}

	// e.g. "logGroup"
	factoryFunctionName := string(bytes.ToLower([]byte(parts[2][:1]))) + parts[2][1:]

	sb := strings.Builder{}
	codeEmitter := codegen.Emitter{
		CfResource: resource.CfResource,
		Writer:     &sb,
	}

	// Generate code for the CloudFormation root properties schema.
	codeFeatures, err := codeEmitter.EmitRootPropertiesSchema()

	if err != nil {
		return fmt.Errorf("error emitting root properties schema code: %w", err)
	}

	rootPropertiesSchema := sb.String()
	sb.Reset()

	templateData := TemplateData{
		CloudFormationTypeName: cfTypeName,
		FactoryFunctionName:    factoryFunctionName,
		HasUpdateMethod:        true,
		PackageName:            packageName,
		PrimaryIdentifierPath:  string(resource.CfResource.PrimaryIdentifier[0]),
		RootPropertiesSchema:   rootPropertiesSchema,
		SchemaVersion:          1,
		TerraformTypeName:      resource.TfType,
	}

	if codeFeatures&codegen.HasUpdatableProperty == 0 {
		templateData.HasUpdateMethod = false
	}

	if description := resource.CfResource.Description; description != nil {
		templateData.SchemaDescription = *description
	}

	for _, path := range resource.CfResource.WriteOnlyProperties {
		templateData.WriteOnlyPropertyPaths = append(templateData.WriteOnlyPropertyPaths, string(path))
	}

	tmpl, err := template.New("function").Parse(templateBody)

	if err != nil {
		return fmt.Errorf("error parsing function template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	generatedFileContents, err := format.Source(buffer.Bytes())

	if err != nil {
		g.Infof("%s", buffer.String())
		return fmt.Errorf("error formatting generated source code: %w", err)
	}

	f, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("error creating file (%s): %w", filename, err)
	}

	defer f.Close()

	_, err = f.Write(generatedFileContents)

	if err != nil {
		return fmt.Errorf("error writing to file (%s): %w", filename, err)
	}

	return nil
}

type TemplateData struct {
	CloudFormationTypeName string
	FactoryFunctionName    string
	HasUpdateMethod        bool
	PackageName            string
	PrimaryIdentifierPath  string
	RootPropertiesSchema   string
	SchemaDescription      string
	SchemaVersion          int64
	TerraformTypeName      string
	WriteOnlyPropertyPaths []string
}

var templateBody = `
// Code generated by generators/resource/main.go; DO NOT EDIT.

package {{ .PackageName }}

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("{{ .TerraformTypeName }}", {{ .FactoryFunctionName }})
}

// {{ .FactoryFunctionName }} returns the Terraform {{ .TerraformTypeName }} resource type.
// This Terraform resource type corresponds to the CloudFormation {{ .CloudFormationTypeName }} resource type.
func {{ .FactoryFunctionName }}(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := {{ .RootPropertiesSchema }}

{{ $tick := "` + "`" + `" }}

	schema := schema.Schema{
		Description: {{ $tick }}{{ .SchemaDescription }}{{ $tick }},
		Version:     {{ .SchemaVersion }},
		Attributes:  attributes,
	}

	var features ResourceTypeFeatures

{{ if .HasUpdateMethod }}
	features |= ResourceTypeHasUpdatableAttribute
{{- end }}

	resourceType, err := NewResourceType(
		"{{ .CloudFormationTypeName }}", // CloudFormation type name
		"{{ .TerraformTypeName }}", // Terraform type name
		schema, // Terraform schema
		"{{ .PrimaryIdentifierPath }}", // Primary identifier property path (JSON Pointer)
		[]string{
{{- range .WriteOnlyPropertyPaths }}
			"{{ . }}",
{{- end }}
		}, // Write-only property paths (JSON Pointer)
		features,
	)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema for %s:\n\n%v", "{{ .TerraformTypeName }}", schema)

	return resourceType, nil
}
`

type Resource struct {
	CfResource *cfschema.Resource
	TfType     string
}

func NewResource(resourceType, cfTypeSchemaFile string) (*Resource, error) {
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath(cfTypeSchemaFile)

	if err != nil {
		return nil, fmt.Errorf("error reading CloudFormation Resource Type Schema: %w", err)
	}

	resource, err := resourceSchema.Resource()

	if err != nil {
		return nil, fmt.Errorf("error parsing CloudFormation Resource Type Schema: %w", err)
	}

	// Ensure that there is one and only one primary identifier path.
	if got, expected := len(resource.PrimaryIdentifier), 1; got != expected {
		return nil, fmt.Errorf("expected %d primary identifier path(s), got: %d", expected, got)
	}

	if err := resource.Expand(); err != nil {
		return nil, fmt.Errorf("error expanding JSON Pointer references: %w", err)
	}

	return &Resource{
		CfResource: resource,
		TfType:     resourceType,
	}, nil
}
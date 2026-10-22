package main

// Run "go generate" to generate the docs for the registry: tfplugindocs builds the provider,
// dumps its schema through Terraform and renders docs/ from it, the templates in templates/ and
// the examples in examples/.
//go:generate go run github.com/fbreckle/terraform-plugin-docs/cmd/tfplugindocs

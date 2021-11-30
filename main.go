package main

import (
	"cdk.tf/go/stack/generated/kubernetes"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here
	kubernetes.NewKubernetesProvider(stack, jsii.String("provider"), &kubernetes.KubernetesProviderConfig{})

	label := map[string]*string{"app": jsii.String("hello")}

	ns := kubernetes.NewNamespace(stack, jsii.String("namespace"), &kubernetes.NamespaceConfig{
		Metadata: &kubernetes.NamespaceMetadata{
			Name:   jsii.String("test"),
			Labels: &label,
		},
	})

	ns.Metadata().Labels() // fails with "panic: unable to make an instance of unregistered interface interface {}"

	// should work
	// kubernetes.NewNamespace(stack, jsii.String("namespace"), &kubernetes.NamespaceConfig{
	// 	Metadata: &kubernetes.NamespaceMetadata{
	// 		Name:   jsii.String("another"),
	// 		Labels: cdktf.Token_AsAny(ns.Metadata().Labels()),
	// 	},
	// })

	// should work
	// kubernetes.NewPod(stack, jsii.String("pod"), &kubernetes.PodConfig{
	// 	Spec: &kubernetes.PodSpec{},
	// 	Metadata: &kubernetes.PodMetadata{
	// 		Name:   jsii.String("pod"),
	// 		Labels: ns.Metadata().Labels(),
	// 	},
	// })

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "test-cdktf-go-list-refs-k8s")

	app.Synth()
}

# test-cdktf-go-list-refs-k8s

Accessing (and using) list outputs should work in Go.
E.g. `namespace.Metadata().Labels()` throws instead of returning a resolvable or a token

```
panic: unable to make an instance of unregistered interface interface {}

    goroutine 1 [running]:
    github.com/aws/jsii-runtime-go/internal/kernel.(*Client).castAndSetToPtr(0xc00000c078, 0x1a5c040, 0xc0004fc820, 0x194, 0x1a730a0, 0xc0004f8d20, 0x15)
    	/Users/ansgar/projects/go/pkg/mod/github.com/aws/jsii-runtime-go@v1.37.0/internal/kernel/conversions.go:78 +0x1127
    github.com/aws/jsii-runtime-go/internal/kernel.(*Client).CastAndSetToPtr(0xc00000c078, 0x19f00e0, 0xc0004fc820, 0x1a730a0, 0xc0004f8d20)
    	/Users/ansgar/projects/go/pkg/mod/github.com/aws/jsii-runtime-go@v1.37.0/internal/kernel/conversions.go:22 +0x15a
    github.com/aws/jsii-runtime-go/runtime.Get(0x1bd7020, 0xc0004fc810, 0x1c1d276, 0x6, 0x19f00e0, 0xc0004fc820)
    	/Users/ansgar/projects/go/pkg/mod/github.com/aws/jsii-runtime-go@v1.37.0/runtime/runtime.go:298 +0x250
    cdk.tf/go/stack/generated/kubernetes.(*jsiiProxy_NamespaceMetadataOutputReference).Labels(0xc0004fc810, 0xacbae00, 0xc0004fc810)
    	/Users/ansgar/projects/playground/test-cdktf-go-list-refs-k8s/generated/kubernetes/kubernetes.go:230506 +0x7b
    main.NewMyStack(0xacbaa60, 0xc0004fc040, 0x1c29b57, 0x1b, 0xc0004fc040, 0x0)
    	/Users/ansgar/projects/playground/test-cdktf-go-list-refs-k8s/main.go:25 +0x325
    main.main()
    	/Users/ansgar/projects/playground/test-cdktf-go-list-refs-k8s/main.go:50 +0x7f
    exit status 2
```

## Reproduce by running
```
yarn fail
```
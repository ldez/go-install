# Bug with go install -mod=readonly <pkg@version>

- https://github.com/golang/go/issues/54908
- https://github.com/golangci/golangci-lint/discussions/3177

```console
$ docker run --rm -it golang:1.19 sh
# go install -mod=readonly github.com/ldez/go-install@5fca274bdce3e0aac69eefcc3dcaf867c7c8d9dc
go: downloading github.com/ldez/go-install v0.0.0-20220906210438-5fca274bdce3
go: downloading golang.org/x/exp v0.0.0-20220906200021-fcb1a314c389
# golang.org/x/exp/constraints
pkg/mod/golang.org/x/exp@v0.0.0-20220906200021-fcb1a314c389/constraints/constraints.go:13:2: embedding interface element ~int|~int8|~int16|~int32|~int64 requires go1.18 or later (-lang was set to go1.16; check go.mod)
pkg/mod/golang.org/x/exp@v0.0.0-20220906200021-fcb1a314c389/constraints/constraints.go:20:2: embedding interface element ~uint|~uint8|~uint16|~uint32|~uint64|~uintptr requires go1.18 or later (-lang was set to go1.16; check go.mod)
    pkg/mod/golang.org/x/exp@v0.0.0-20220906200021-fcb1a314c389/constraints/constraints.go:27:2: embedding interface element Signed|Unsigned requires go1.18 or later (-lang was set to go1.16; check go.mod)
pkg/mod/golang.org/x/exp@v0.0.0-20220906200021-fcb1a314c389/constraints/constraints.go:34:2: embedding interface element ~float32|~float64 requires go1.18 or later (-lang was set to go1.16; check go.mod)
pkg/mod/golang.org/x/exp@v0.0.0-20220906200021-fcb1a314c389/constraints/constraints.go:41:2: embedding interface element ~complex64|~complex128 requires go1.18 or later (-lang was set to go1.16; check go.mod)
pkg/mod/golang.org/x/exp@v0.0.0-20220906200021-fcb1a314c389/constraints/constraints.go:49:2: embedding interface element Integer|Float|~string requires go1.18 or later (-lang was set to go1.16; check go.mod)
# 
```

```console
$ docker run --rm -it golang:1.19 sh
# git clone https://github.com/ldez/go-install.git
Cloning into 'go-install'...
remote: Enumerating objects: 8, done.
remote: Counting objects: 100% (8/8), done.
remote: Compressing objects: 100% (6/6), done.
remote: Total 8 (delta 0), reused 8 (delta 0), pack-reused 0
Receiving objects: 100% (8/8), done.
# git checkout 5fca274bdce3e0aac69eefcc3dcaf867c7c8d9dc
# cd go-install
# go install -mod=readonly .
go: downloading golang.org/x/exp v0.0.0-20220906200021-fcb1a314c389
#
```

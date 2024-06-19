# pre-commit hook for go

## Usage

Add a .pre-commit-config.yaml at the root of your repository with this content:

```yaml
repos:
-   repo: https://github.com/sondalex/pre-commit-go
    rev: <rev> # should be a commit id or git tag
    hooks:
    -   id: gofmt 
    -   id: govet
```

## Features

* go format (`id: gofmt`)
* go vet: (`id: govet`)

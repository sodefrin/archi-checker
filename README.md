# archi-checker

![build-and-test](https://github.com/sodefrin/archi-checker/workflows/build-and-test/badge.svg)

This tool is a linter that checks dependencies based on import statements.

## Usage

Create default .archi-checker.uml.

```
$ archi-checkr -init  $(go list ./...)
$ cat .archi-checker.uml
default : github.com/google/go-cmp/cmp
default : github.com/google/go-cmp/cmp/cmpopts
default : github.com/sodefrin/archi-checker
default : golang.org/x/mod/modfile
```

Edit created .archi-checkr.uml to group package like below.

```
util : github.com/google/go-cmp/cmp
util : github.com/google/go-cmp/cmp/cmpopts
main : github.com/sodefrin/archi-checker
util : golang.org/x/mod/modfile
```

Then describe dependency between each package group.

```
util : github.com/google/go-cmp/cmp
util : github.com/google/go-cmp/cmp/cmpopts
main : github.com/sodefrin/archi-checker
util : golang.org/x/mod/modfile

main -> util
```

Then validate dependency.

```
$ archi-checker ($go list ./...)
```

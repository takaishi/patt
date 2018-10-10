# patt

## Description

* Create text file from template.
* Manage template files.

## Usage
### Add template

Template has front-matter thats has 'name' and 'destination'.

```
$ cat /path/to/template.md
+++
{
    "name": "test",
    "destination": "/path/to/destination/{{.Year}}{{.Month}}{{.Day}}.md"
}
+++
# {{.Year}}-{{.Month}}-{{.Day}}({{.Week}})

## Header2

- foo
- bar
```


```
$ patt add /path/to/template.md
```

### List templates

```
$ patt list
+------+---------------------------------------------+
| NAME |                   SOURCE                    |
+------+---------------------------------------------+
| foo  | /Users/r_takaishi/.patt.d/templates/foo.md  |
| test | /Users/r_takaishi/.patt.d/templates/test.md |
+------+---------------------------------------------+
```

### Create file from template

```
$ patt new test
```

```
$ cat  /path/to/destination/20170108.md
# 2017-01-05(Thu)

## Header2

- foo
- bar
```

### Delete template

```
$ patt delete foo
```

## Install

Use Homebrew:

```
$ brew install takaishi/fomulas/patt
```

Use `go get`:

```bash
$ go get -d github.com/takaishi/patt
```

## Contribution

1. Fork ([https://github.com/takaishi/patt/fork](https://github.com/takaishi/patt/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[takaishi](https://github.com/takaishi)

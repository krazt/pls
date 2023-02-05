# pls

`pls` is a CLI program that uses AI to tell you the command you are needing to accomplish a certain task from your
terminal.
Additionally, you can use it to generate and modify text and images, given a text description.

Its name is a short form of "please".

You can read more about it in the [design document](docs/design_document.md).

## Installation

```sh
go install github.com/krazt/pls@latest
```

## Usage

```sh
pls --help
```

## Examples

### Generate a command given a text description

You run:

```sh
pls "list only hidden files"
```

It outputs:

```text
Explanation:
    ls: list files in the current directory
    -a: list all files, including hidden files
    
    |: pipe the output of ls to the input of grep
    
    grep: search for a pattern in the input
    -v: invert the match, i.e. only show lines that do not match the pattern
    "^\.": the pattern is a regular expression that matches lines that start with a dot

ls -a | grep -v "^\."

Run the command? (y=yes, N=no, e=edit)
```

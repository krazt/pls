# pls

`pls` is a CLI program that uses AI to tell you the command you are needing to accomplish a certain task from your
terminal.
Additionally, you can use it to generate and modify text and images, given a text description.

Its name is a short form of "please". And careful — `pls` acts a little grumpy when asked out-of-scope questions.

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
ls -a | grep -v "^\."

Run the command? (y=yes, N=no, e=edit)
```

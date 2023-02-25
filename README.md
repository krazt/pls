# pls

`pls` is a CLI program that uses AI (OpenAI's GPT-3.5) to tell you the command you need to run in your terminal,
given your description in natural language.
Also, it can run the command for you.

Its name is a short form of "please". And careful — `pls` acts a little grumpy when asked out-of-scope questions.

You can read more about it in the [design document](docs/design_document.md).

## Installation

```sh
go install github.com/krazt/pls@latest
```

## Usage

You run:

```sh
pls list only hidden files
```

It outputs:

```text
ls -a | grep -v "^\."

Run the command? [y/N]
```

Write a shell command that would best accomplish the task described in the input the user has provided.

---

## Context

- OS: linux
- Arch: amd64
- Current Dir: /home/aubrey

## Input

"list only hidden files"

## Output

ls -a | grep "^."

---

## Context

- OS: linux
- Arch: amd64
- Current Dir: /home/raul/Downloads

## Input

"make a request to google and save to response.txt"

## Output

curl -s -o response.txt https://www.google.com

---

## Context

- OS: {{.OS}}
- Arch: {{.Arch}}
- Current Dir: {{.Dir}}

## Input

{{printf "%q" .Input}}

## Output

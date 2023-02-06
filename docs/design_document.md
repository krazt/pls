# Design Document

## Problem Statement

I'm a software developer using the terminal. Sometimes I don't remember or don't know what command(s) to use to achieve
certain tasks. I want a CLI program where I can just describe what I want to achieve in natural language, and then the
program outputs a command I can use for that.

Additionally, I want the program to be able to do the following:

- Generate text (to just print it to stdout). For example, to write a poem or some JSON or even to respond to "Hey, how
  are you doing!?"
- Generate an image given a text description
- Modify an image given a text description

## Solution

I'm going to build a CLI program that can do the following:

- Generate a command given a text description
- Generate text given a text description
- Generate an image given a text description
- Modify an image given a text description

## Implementation

Given the user input:

1. Predict the action
2. Generate the content

### Predicting the action

Available actions:

1. Unknown
2. Execute a command in the terminal
3. Perform a calculation
4. Get an answer that should be searched on the web
5. Chat
6. Generate text
7. Generate an image given a text description
8. Modify an image given a text description

### Generating the content

#### Action: Unknown

The user input is not recognized. The program should print an error message to stderr.

#### Action: Execute a command in the terminal

The user wants to execute a command in the terminal. The program should generate the command and execute it in the
terminal, asking the user for confirmation before executing it.

#### Action: Perform a calculation

The user wants to perform a calculation. The program should generate the calculation and print the result to stdout.

#### Action: Get an answer that should be searched on the web

The user wants to get an answer to a question. The program should generate a search query, search the web, and print
the answer to stdout.

#### Action: Chat

The user wants to have a conversation. The program should generate a response to the user's input as if they were
having a conversation, printing the response to stdout.

#### Action: Generate text

The user wants to generate some text. The program should generate the text and print it to stdout.

#### Action: Generate an image given a text description

The user wants to generate an image given a text description. The program should generate the image and save it to a
file.

#### Action: Modify an image given a text description

The user wants to modify an image given a text description. The program should modify the image and save it to a file.

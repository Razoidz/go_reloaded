# Go Reloaded

Go Reloaded is a command-line text processor written in Go. The current implementation converts hexadecimal values marked with `(hex)` into decimal numbers.

## Requirements

- Go 1.22.2 or newer

## Run

Pass an input file and an output file:

```sh
go run . sample.txt result.txt
```

The program reads the input, processes it, and writes the result to the output path. You can also build an executable:

```sh
go build -o go-reloaded .
./go-reloaded sample.txt result.txt
```

## Supported conversion

Place `(hex)` after a hexadecimal value. The marker is removed and the value is written in decimal.

```text
Input:  The number 1E (hex) is hexadecimal.
Output: The number 30 is hexadecimal.
```

Hexadecimal digits may use either uppercase or lowercase letters. Input is split on whitespace and joined with single spaces in the output, so the original spacing and line breaks are not preserved.

## Current limitations

Only the `(hex)` conversion is currently connected to the command-line program. Binary conversion, case modifiers, punctuation spacing, quote spacing, and article correction are not currently implemented in the executable. Invalid hexadecimal values are not reported as errors.

The command expects exactly two file paths. If the input file cannot be read, it prints an error message and exits. The output file is created or overwritten when processing succeeds.

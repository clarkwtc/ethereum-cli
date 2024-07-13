package utils

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type CommandLine struct {
    *bufio.Reader
}

func NewCommandLine() *CommandLine {
    return &CommandLine{bufio.NewReader(os.Stdin)}
}

func (commandLine *CommandLine) Input() (string, error) {
    input, err := commandLine.ReadString('\n')
    if err != nil {
        fmt.Println(err)
        return "", err
    }

    input = strings.TrimSpace(input)
    input = strings.ToLower(input)
    return input, nil
}

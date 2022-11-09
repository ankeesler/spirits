package menu

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Menu []Item

type IO struct {
	In  io.Reader
	Out io.Writer
}

func (m Menu) Run(ctx context.Context, io *IO) (context.Context, error) {
	for {
		fmt.Fprintln(io.Out, "--- Menu ---")
		for i, item := range m {
			fmt.Fprintf(io.Out, "%d\t%s\n", i, item.Title)
		}
		fmt.Fprintf(io.Out, "b\tGo back\n")
		fmt.Fprintln(io.Out, "---")

		selection, err := Input(io, "Selection: ")
		if err != nil {
			return ctx, fmt.Errorf("input: %w", err)
		}

		if selection == "b" {
			return ctx, nil
		}

		selectionInt, err := strconv.Atoi(selection)
		if err != nil || selectionInt < 0 || selectionInt >= len(m) {
			fmt.Fprintln(io.Out, "Invalid selection")
			continue
		}

		ctx, err = m[selectionInt].Run(ctx, io)
		if err != nil {
			return ctx, err
		}
	}
}

type Item struct {
	Title string
	Runner
}

type Runner interface {
	Run(context.Context, *IO) (context.Context, error)
}

type RunnerFunc func(context.Context, *IO) (context.Context, error)

func (f RunnerFunc) Run(ctx context.Context, io *IO) (context.Context, error) {
	return f(ctx, io)
}

func Input(io *IO, prompt string) (string, error) {
	r := bufio.NewReader(io.In)

	fmt.Print(prompt)

	selection, err := r.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}
	selection = strings.TrimSpace(selection)

	return selection, nil
}

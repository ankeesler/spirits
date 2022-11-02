package menu

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Menu[T any] []Item[T]

func (m Menu[T]) Run(ctx context.Context, out io.Writer, in io.Reader, t T) error {
	for {
		fmt.Fprintln(out, "--- Menu ---")
		for i, item := range m {
			fmt.Fprintf(out, "%d\t%s\n", i, item.Title)
		}
		fmt.Fprintf(out, "b\tGo back\n")
		fmt.Fprintln(out, "---")

		selection, err := Input(out, in, "Selection: ")
		if err != nil {
			return fmt.Errorf("input: %w", err)
		}

		if selection == "b" {
			return nil
		}

		selectionInt, err := strconv.Atoi(selection)
		if err != nil || selectionInt < 0 || selectionInt >= len(m) {
			fmt.Fprintln(out, "Invalid selection")
			continue
		}

		if err := m[selectionInt].Run(ctx, out, in, t); err != nil {
			return err
		}
	}
}

type Item[T any] struct {
	Title string
	Runner[T]
}

type Runner[T any] interface {
	Run(context.Context, io.Writer, io.Reader, T) error
}

type RunnerFunc[T any] func(context.Context, io.Writer, io.Reader, T) error

func (f RunnerFunc[T]) Run(ctx context.Context, out io.Writer, in io.Reader, t T) error {
	return f(ctx, out, in, t)
}

func Input(out io.Writer, in io.Reader, prompt string) (string, error) {
	r := bufio.NewReader(in)

	selection, err := r.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}
	selection = strings.TrimSpace(selection)

	return selection, nil
}

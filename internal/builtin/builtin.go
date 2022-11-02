package builtin

import (
	"context"
	"fmt"
	fspkg "io/fs"
	"time"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

type Creator[T any] interface {
	Create(context.Context, T, func(T) error) (T, error)
}

func Load[T proto.Message](
	fs fspkg.FS,
	creator Creator[T],
	newFunc func() T,
	validateFunc func(T) error,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := fspkg.WalkDir(fs, ".", func(path string, d fspkg.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		b, err := fspkg.ReadFile(fs, path)
		if err != nil {
			return fmt.Errorf("read file: %w", err)
		}

		t := newFunc()
		if err := prototext.Unmarshal(b, t); err != nil {
			return fmt.Errorf("unmarshal proto %s: %w", path, err)
		}

		if _, err := creator.Create(ctx, t, validateFunc); err != nil {
			return fmt.Errorf("create %s: %w", path, err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("walk dir: %w", err)
	}
	return nil
}

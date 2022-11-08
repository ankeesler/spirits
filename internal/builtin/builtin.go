package builtin

import (
	"context"
	"fmt"
	fspkg "io/fs"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

type Repo[T any] interface {
	Create(context.Context, T) (T, error)
}

func Load[T proto.Message, U any](
	ctx context.Context,
	fs fspkg.FS,
	newFunc func() T,
	convertFunc func(T) (U, error),
	repo Repo[U],
) error {
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

		u, err := convertFunc(t)
		if err != nil {
			return fmt.Errorf("convert: %w", err)
		}

		if _, err := repo.Create(ctx, u); err != nil {
			return fmt.Errorf("create %s: %w", path, err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("walk dir: %w", err)
	}
	return nil
}

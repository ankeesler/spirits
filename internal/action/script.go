package action

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ankeesler/spirits/internal/spirit"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

const ctxKey = "action.ctx"

type script struct {
	program *starlark.Program
}

func compile(source string) (*script, error) {
	s := &script{}

	predeclared, err := newPredeclared(
		&spirit.Spirit{},
		[]*spirit.Spirit{},
	)
	if err != nil {
		return nil, fmt.Errorf("get script predeclared symbols for compile: %w", err)
	}

	_, program, err := starlark.SourceProgram("<actionscript:source>", source, predeclared.Has)
	if err != nil {
		return nil, err
	}
	s.program = program

	return s, nil
}

func (s *script) Run(
	ctx context.Context, source *spirit.Spirit, targets []*spirit.Spirit) (context.Context, error) {
	out := bytes.NewBuffer([]byte{})
	thread := &starlark.Thread{
		Name: "<actionscript:load>",
		Print: func(thread *starlark.Thread, msg string) {
			msg = fmt.Sprintf("%s: %s", thread.Name, msg)
			fmt.Fprintf(out, msg)
			log.Printf(msg)
		},
	}
	predeclared, err := newPredeclared(source, targets)
	if err != nil {
		return ctx, fmt.Errorf("get script predeclared symbols for run: %w", err)
	}

	if err := s.run(ctx, thread, predeclared); err != nil {
		return ctx, fmt.Errorf("run script: %w (out: %q)", err, out.String())
	}

	return ctx, nil
}

func (s *script) run(
	ctx context.Context,
	thread *starlark.Thread,
	predeclared starlark.StringDict,
) error {
	type starlarkInitRet struct {
		globals starlark.StringDict
		err     error
	}
	done := make(chan *starlarkInitRet)
	defer close(done)

	thread.SetLocal(ctxKey, ctx)
	go func() {
		globals, err := s.program.Init(thread, predeclared)
		if err != nil {
			err = fmt.Errorf("script failed: %w (%s, %s)", err, thread.Local("resolve"), thread.Local("reject"))
		}
		done <- &starlarkInitRet{globals: globals, err: err}
	}()

	var initRet *starlarkInitRet
	select {
	case <-ctx.Done():
		thread.Cancel(ctx.Err().Error())
		initRet = <-done
	case initRet = <-done:
	}

	return initRet.err
}

func newPredeclared(
	source *spirit.Spirit,
	targets []*spirit.Spirit,
) (starlark.StringDict, error) {
	starlarkCtxStruct := starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
		"value": starlark.NewBuiltin(
			"ctx.value",
			func(
				thread *starlark.Thread,
				b *starlark.Builtin,
				args starlark.Tuple,
				kwargs []starlark.Tuple,
			) (starlark.Value, error) {
				var key starlark.Value
				if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 1, &key); err != nil {
					return nil, err
				}

				ctx, ok := thread.Local(ctxKey).(context.Context)
				if !ok {
					return nil, errors.New("missing thread local context")
				}

				val := ctx.Value(key)
				if val == nil {
					return starlark.None, nil
				}

				return starlark.String(val.(string)), nil
			},
		),
		"set_value": starlark.NewBuiltin(
			"ctx.value",
			func(
				thread *starlark.Thread,
				b *starlark.Builtin,
				args starlark.Tuple,
				kwargs []starlark.Tuple,
			) (starlark.Value, error) {
				var key, val starlark.Value
				if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 2, &key, &val); err != nil {
					return nil, err
				}
				return val, nil
			},
		),
	})

	var starlarkTargets []starlark.Value
	for _, target := range targets {
		starlarkTargets = append(starlarkTargets, newSpiritStarlarkStruct(target))
	}

	starlarkAbortFunc := starlark.NewBuiltin(
		"abort",
		func(
			thread *starlark.Thread,
			b *starlark.Builtin,
			args starlark.Tuple,
			kwargs []starlark.Tuple,
		) (starlark.Value, error) {
			thread.Cancel(fmt.Sprintf("abort(%v, %v)", args, kwargs))
			return starlark.None, nil
		},
	)

	starlarkActionStruct := starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
		"ctx":     starlarkCtxStruct,
		"source":  newSpiritStarlarkStruct(source),
		"targets": starlark.NewList(starlarkTargets),
		"abort":   starlarkAbortFunc,
	})

	return starlark.StringDict{
		"action": starlarkActionStruct,
	}, nil
}

func newSpiritStarlarkStruct(spirit *spirit.Spirit) *starlarkstruct.Struct {
	starlarkStatsDict := starlark.StringDict{}
	stats := spirit.Stats()
	addStatStarlarkBuitlins(
		starlarkStatsDict, "health",
		stats.Health, stats.SetHealth)
	addStatStarlarkBuitlins(
		starlarkStatsDict, "physical_power",
		stats.PhysicalPower, stats.SetPhysicalPower)
	addStatStarlarkBuitlins(
		starlarkStatsDict, "physical_constitution",
		stats.PhysicalConstitution, stats.SetPhysicalConstitution)
	addStatStarlarkBuitlins(
		starlarkStatsDict, "mental_power",
		stats.MentalPower, stats.SetMentalPower)
	addStatStarlarkBuitlins(
		starlarkStatsDict, "mental_constitution",
		stats.MentalConstitution, stats.SetMentalConstitution)
	addStatStarlarkBuitlins(
		starlarkStatsDict, "agility", stats.Health, stats.SetHealth)
	starlarkStatsStruct := starlarkstruct.FromStringDict(starlarkstruct.Default, starlarkStatsDict)

	return starlarkstruct.FromStringDict(starlarkstruct.Default, starlark.StringDict{
		"id":    starlark.String(spirit.ID()),
		"name":  starlark.String(spirit.Name()),
		"stats": starlarkStatsStruct,
	})
}

func addStatStarlarkBuitlins(
	starlarkDict starlark.StringDict,
	statName string,
	getterFunc func() int64,
	setterFunc func(int64),
) {
	getter := fmt.Sprintf("%s", statName)
	setter := fmt.Sprintf("set_%s", statName)

	starlarkDict[getter] = starlark.NewBuiltin(
		getter,
		func(
			thread *starlark.Thread,
			b *starlark.Builtin,
			args starlark.Tuple,
			kwargs []starlark.Tuple,
		) (starlark.Value, error) {
			return starlark.MakeInt64(getterFunc()), nil
		},
	)

	starlarkDict[setter] = starlark.NewBuiltin(
		setter,
		func(
			thread *starlark.Thread,
			b *starlark.Builtin,
			args starlark.Tuple,
			kwargs []starlark.Tuple,
		) (starlark.Value, error) {
			var newStat int64
			if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 1, &newStat); err != nil {
				return nil, err
			}
			setterFunc(newStat)
			return starlark.None, nil
		},
	)
}

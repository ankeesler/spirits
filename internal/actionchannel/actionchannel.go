package actionchannel

import (
	"context"
	"fmt"
	"sync"
)

type ActionChannel struct {
	m sync.Map
}

func (ac *ActionChannel) Post(
	namespace, battleName, battleGeneration, spiritName, spiritGeneration, actionName string,
) error {
	select {
	case ac.c(namespace, battleName, battleGeneration, spiritName, spiritGeneration) <- actionName:
	default:
		return fmt.Errorf(
			"no one listening for battleName %q battleGeneration %q spiritName %q spiritGeneration %q actions",
			battleName,
			battleGeneration,
			spiritName,
			spiritGeneration,
		)
	}
	return nil
}

func (ac *ActionChannel) Pend(
	ctx context.Context,
	namespace, battleName, battleGeneration, spiritName, spiritGeneration string,
) (string, error) {
	select {
	case actionName, ok := <-ac.c(namespace, battleName, battleGeneration, spiritName, spiritGeneration):
		if !ok {
			return "", fmt.Errorf(
				"channel closed for namespace %q battleName %q battleGeneration %q spiritName %q spiritGeneration %q actions",
				namespace,
				battleName,
				battleGeneration,
				spiritName,
				spiritGeneration,
			)
		}
		return actionName, nil
	case <-ctx.Done():
		return "", fmt.Errorf(
			"context canceled for namespace %q battleName %q battleGeneration %q spiritName %q spiritGeneration %q actions",
			namespace,
			battleName,
			battleGeneration,
			spiritName,
			spiritGeneration,
		)
	}
}

func (ac *ActionChannel) Close() {
	ac.m.Range(func(_, c any) bool {
		close(c.(chan string))
		return true
	})
}

func (ac *ActionChannel) c(
	namespace, battleName, battleGeneration, spiritName, spiritGeneration string,
) chan string {
	key := fmt.Sprintf(
		"%s-%s-%s-%s-%s",
		namespace,
		battleName,
		battleGeneration,
		spiritName,
		spiritGeneration,
	)
	c := make(chan string)
	cc, exists := ac.m.LoadOrStore(key, c)
	if exists {
		close(c)
		c = cc.(chan string)
	}
	return c
}

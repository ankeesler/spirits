package actionchannel

import (
	"context"
	"fmt"
	"sync"

	"k8s.io/klog/v2"
)

type ActionChannel struct {
	m sync.Map
}

func (ac *ActionChannel) Post(
	namespace, battleName, spiritName, actionName string,
) error {
	select {
	case ac.c(namespace, battleName, spiritName) <- actionName:
	default:
		return fmt.Errorf(
			"no one listening for battleName %q spiritName %q actions",
			battleName,
			spiritName,
		)
	}
	return nil
}

func (ac *ActionChannel) Pend(
	ctx context.Context,
	namespace, battleName, spiritName string,
) (string, error) {
	klog.V(1).InfoS("actionchannel pend", "namespace", namespace, "battleName", battleName, "spiritName", spiritName)
	select {
	case actionName, ok := <-ac.c(namespace, battleName, spiritName):
		if !ok {
			return "", fmt.Errorf(
				"channel closed for namespace %q battleName %q spiritName %q actions",
				namespace,
				battleName,
				spiritName,
			)
		}
		return actionName, nil
	case <-ctx.Done():
		return "", fmt.Errorf(
			"context canceled for namespace %q battleName %q spiritName %q actions: %q",
			namespace,
			battleName,
			spiritName,
			ctx.Err().Error(),
		)
	}
}

func (ac *ActionChannel) Close() {
	ac.m.Range(func(_, c any) bool {
		close(c.(chan string))
		return true
	})
}

func (ac *ActionChannel) c(namespace, battleName, spiritName string) chan string {
	key := fmt.Sprintf("%s-%s-%s", namespace, battleName, spiritName)
	c := make(chan string)
	cc, exists := ac.m.LoadOrStore(key, c)
	if exists {
		close(c)
		c = cc.(chan string)
	}
	return c
}

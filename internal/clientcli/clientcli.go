package clientcli

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	configpkg "sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

const (
	namespace = "default"

	createdBySpiritsClientLabelKey   = "spirits.ankeesler.github.com/created-by"
	createdBySpiritsClientLabelValue = "clientcli"
)

type config struct {
	ctx      context.Context
	command  string
	args     []string
	out, err io.Writer
	client   client.Client
}

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(spiritsv1alpha1.AddToScheme(scheme))
}

func Run(args []string, stdout, stderr io.Writer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	command := "battle"
	if len(args) > 1 {
		command = args[1]
	}

	client, err := createClient()
	if err != nil {
		return err
	}

	c := config{
		ctx:     ctx,
		command: command,
		args:    args[1:],
		out:     stdout,
		err:     stderr,
		client:  client,
	}
	runner, ok := map[string]func(*config) error{
		"spirit": runSpirit,
		"battle": runBattle,
		"action": runAction,
	}[command]
	if !ok {
		return fmt.Errorf("unknown command: %q", command)
	}
	if err := runner(&c); err != nil {
		return err
	}

	return nil
}

func createClient() (client.Client, error) {
	config, err := configpkg.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("load kubeconfig: %w", err)
	}

	client, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	return client, nil
}

func createOrPatch(c *config, obj client.Object, mutateFunc func() error) error {
	kind := reflect.TypeOf(obj).String()
	if len(obj.GetGenerateName()) > 0 {
		if err := c.client.Create(c.ctx, obj); err != nil {
			return fmt.Errorf("generate %s: %w", kind, err)
		}
	}
	if _, err := controllerutil.CreateOrPatch(c.ctx, c.client, obj, mutateFunc); err != nil {
		return fmt.Errorf("upsert %s %q: %w", kind, client.ObjectKeyFromObject(obj).String(), err)
	}
	return nil
}

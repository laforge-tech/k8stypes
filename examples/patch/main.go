package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	//demov1 "github.com/laforge-tech/k8stypes/pkg/api/demo/v1"
	"github.com/laforge-tech/k8stypes/pkg/client"

	demov1ac "github.com/laforge-tech/k8stypes/pkg/applyconfigurations/demo/v1"
	//demov1 "github.com/laforge-tech/k8stypes/pkg/api/demo/v1"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "20060102 15:04:05.000"})

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't get configuration")
	}

	cs, err := client.NewForConfig(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Can't get client")
	}

	ac := demov1ac.Demo("demo1", "default").WithSpec(
		demov1ac.DemoSpec().WithMessage("From sample 1"),
	)

	patched, err := cs.DemoV1().Demos("default").Apply(context.TODO(), ac, metav1.ApplyOptions{
		FieldManager: "sample",
		Force: true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Can't apply")
	}

	log.Info().Msgf("%s", patched.String())


	
	ac = demov1ac.Demo("demo1", "default").WithSpec(
		demov1ac.DemoSpec().WithMessage("From sample 2"),
	).WithStatus(
		demov1ac.DemoStatus().WithPhase("test"),
	)

	patched, err = cs.DemoV1().Demos("default").ApplyStatus(context.TODO(), ac, metav1.ApplyOptions{
		FieldManager: "sample",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Can't apply")
	}

	log.Info().Msgf("%s", patched.String())
}
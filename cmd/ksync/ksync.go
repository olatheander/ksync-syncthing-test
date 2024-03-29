package main

import (
	"fmt"
	"github.com/olatheander/ksync-syncthing-test/pkg/cli"
	"github.com/olatheander/ksync-syncthing-test/pkg/ksync"
	"github.com/olatheander/ksync-syncthing-test/pkg/ksync/cluster"
	"github.com/olatheander/ksync-syncthing-test/pkg/ksync/doctor"
	"github.com/spf13/pflag"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

)

var (
	globalUsage = `Inspect and sync files from remote containers.`

	rootCmd = &cobra.Command{
		Use:              "ksync",
		Short:            "Inspect and sync files from remote containers.",
		Long:             globalUsage,
		PersistentPreRun: initPersistent,
	}
)

func main() {
	rootCmd.AddCommand(
		(&initCmd{}).new(),
		(&watchCmd{}).new(),
	)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("%v", err)
	}
}

func init() {
	cobra.OnInitialize(func() {
		if err := cli.InitConfig("ksync"); err != nil {
			log.Fatal(err)
		}
	})

	if err := cli.DefaultFlags(rootCmd, "ksync"); err != nil {
		log.Fatal(err)
	}

	flags := rootCmd.PersistentFlags()
	localFlags(flags)
	remoteFlags(flags)
}

func localFlags(flags *pflag.FlagSet) {
	flags.StringP(
		"namespace",
		"n",
		"default",
		"namespace to use")
	if err := cli.BindFlag(
		viper.GetViper(), flags.Lookup("namespace"), "ksync"); err != nil {

		log.Fatal(err)
	}

	flags.String(
		"context",
		"",
		"name of the kubeconfig context to use")
	if err := cli.BindFlag(
		viper.GetViper(), flags.Lookup("context"), "ksync"); err != nil {

		log.Fatal(err)
	}

	flags.Int(
		"port",
		40322,
		"port on watch listens on locally")

	if err := cli.BindFlag(
		viper.GetViper(), flags.Lookup("port"), "ksync"); err != nil {

		log.Fatal(err)
	}

	flags.StringP(
		"output",
		"o",
		"pretty",
		"output format to use (e.g. \"json\")")

	if err := cli.BindFlag(viper.GetViper(), flags.Lookup("output"), "ksync"); err != nil {
		log.Fatal(err)
	}
}

func remoteFlags(flags *pflag.FlagSet) { //nolint: gocyclo
	flags.String(
		"image",
		fmt.Sprintf("vaporio/ksync:git-%s", ksync.GitCommit),
		"the image to use on the cluster")
	if err := flags.MarkHidden("image"); err != nil {
		log.Fatal(err)
	}

	if err := cli.BindFlag(
		viper.GetViper(), flags.Lookup("image"), "ksync"); err != nil {

		log.Fatal(err)
	}

	flags.String(
		"apikey",
		"ksync",
		"api key used for authentication with syncthing")
	if err := flags.MarkHidden("apikey"); err != nil {
		log.Fatal(err)
	}

	if err := cli.BindFlag(
		viper.GetViper(), flags.Lookup("apikey"), "ksync"); err != nil {

		log.Fatal(err)
	}

	flags.Int(
		"syncthing-port",
		8384,
		"port on which the syncthing server will listen")
	if err := flags.MarkHidden("syncthing-port"); err != nil {
		log.Fatal(err)
	}

	if err := cli.BindFlag(
		viper.GetViper(), flags.Lookup("syncthing-port"), "ksync"); err != nil {

		log.Fatal(err)
	}

	flags.String(
		"docker-root",
		"/var/lib/docker",
		"root directory of the docker storage (graph) driver")
	if err := flags.MarkHidden("docker-root"); err != nil {
		log.Fatal(err)
	}

	if err := cli.BindFlag(
		viper.GetViper(), flags.Lookup("docker-root"), "ksync"); err != nil {

		log.Fatal(err)
	}

	flags.String(
		"docker-socket",
		"/var/run/docker.sock",
		"path to the docker socket")
	if err := flags.MarkHidden("docker-socket"); err != nil {
		log.Fatal(err)
	}

	if err := cli.BindFlag(
		viper.GetViper(), flags.Lookup("docker-socket"), "ksync"); err != nil {

		log.Fatal(err)
	}

	flags.String(
		"daemonset-namespace",
		"kube-system",
		"Set the namespace for the remote DaemonSet to be deployed in.")
	if err := flags.MarkHidden("daemonset-namespace"); err != nil {
		log.Fatal(err)
	}

	if err := cli.BindFlag(
		viper.GetViper(), flags.Lookup("daemonset-namespace"), "ksync"); err != nil {

		log.Fatal(err)
	}
}

func initPersistent(cmd *cobra.Command, args []string) {
	cli.InitLogging()

	// This is a super special case where we don't want to initialize the k8s
	// client, instead waiting to test it as part of the doctor process.
	if !strings.HasPrefix(cmd.Use, "doctor") {
		initKubeClient()
	}

	cluster.SetImage(viper.GetString("image"))

	cluster.SetErrorHandlers()
}

func initKubeClient() {
	// The act of testing for a config, initializes the config.
	if err := doctor.IsClusterConfigValid(); err != nil {
		log.Fatal(err)
	}
}

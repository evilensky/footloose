package main

import (
	"github.com/spf13/cobra"

	"github.com/dlespiau/footloose/pkg/cluster"
)

var configCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a cluster configuration",
	RunE:  configCreate,
}

var configCreateOptions struct {
	file string
}

func init() {
	configCreateCmd.Flags().StringVarP(&configCreateOptions.file, "config", "c", Footloose, "Cluster configuration file")

	name := &defaultConfig.Cluster.Name
	configCreateCmd.PersistentFlags().StringVarP(name, "name", "n", *name, "Name of the cluster")

	private := &defaultConfig.Cluster.PrivateKey
	configCreateCmd.PersistentFlags().StringVarP(private, "key", "k", *private, "Name of the private and public key files")

	replicas := &defaultConfig.Machines[0].Count
	configCreateCmd.PersistentFlags().IntVarP(replicas, "replicas", "r", *replicas, "Number of machine replicas")

	privileged := &defaultConfig.Machines[0].Spec.Privileged
	configCreateCmd.PersistentFlags().BoolVar(privileged, "privileged", *privileged, "Create privileged containers")

	configCmd.AddCommand(configCreateCmd)
}

func configCreate(cmd *cobra.Command, args []string) error {
	cluster := cluster.New(defaultConfig)
	return cluster.Save(configCreateOptions.file)
}

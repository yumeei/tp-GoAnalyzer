package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gowatcher",
	Short: "GOwatcher est un outil pour vérifier l'accessibilité des URLS",
	Long: "Un outil CLI en Go pour vérifier l'état d'URLS",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func init() {

}
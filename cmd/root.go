package cmd

var rootCmd = &cobra.Command() {
	Use: "gowatcher",
	Short: "GOwatcher est un outil pour vérifier l'accessibilité des URLS",
	Long: "Un outil CLI en Go pour vérifier l'état d'URLS",

}

func Execute() {
	// if err := rootCmd.Execute(); err != 
}
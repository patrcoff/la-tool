package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func DiffCmd() *cobra.Command {
	var params1, params2 string

	cmd := &cobra.Command{
		Use:   "diff <source1> <source2>",
		Short: "Compare two Logic App configurations",
		Long: `Compare Logic App configurations from different sources such as directories, git repositories, or deployed instances.
Optional parameter files can be specified for each source using --params1 and --params2 flags.`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			source1 := args[0]
			source2 := args[1]
			fmt.Printf("Comparing Logic App configurations:\n  Source 1: %s\n  Source 2: %s\n", source1, source2)
			
			if params1 != "" {
				fmt.Printf("  Parameters for Source 1: %s\n", params1)
			}
			if params2 != "" {
				fmt.Printf("  Parameters for Source 2: %s\n", params2)
			}

			// Implement diff logic here
			// Use source1, source2, params1, and params2 in your diff implementation
            fmt.println("The implimentation of diff is yet to be...")
		},
	}

	cmd.Flags().StringVar(&params1, "params1", "", "Parameters file for source1 (optional)")
	cmd.Flags().StringVar(&params2, "params2", "", "Parameters file for source2 (optional)")

	return cmd
}

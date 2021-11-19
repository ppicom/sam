package reports

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/cmd"
)

func newGenerateCustomersReportCmd(generator ReportGenerator) *cobra.Command {
	return &cobra.Command{
		Use:         "generaInformeClients",
		Short:       "Genera l'informe dels clients",
		Long:        "Genera l'informe dels clients actius",
		Example:     `   generaInformeClients      Genera l'informe dels clients`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"ginfc",
			"generainfclients",
			"genera-inf-clients",
			"generarInfClients",
			"generarinfclients",
			"generar-inf-clients",
			"gcrep",
			"generateCustomersReport",
			"generatecustomersreport",
			"generate-customers-report",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := generator.Generate()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}

func init() {
	generator := NewCustomersReportGenerator()
	command := newGenerateCustomersReportCmd(generator)
	cmd.RootCmd.AddCommand(command)
}

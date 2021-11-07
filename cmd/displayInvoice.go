package cmd

import (
	"sam/adm"
	"strings"

	"github.com/spf13/cobra"
)

var displayInvoiceCmd = &cobra.Command{
	Use:         "mostraFactura codiFactura",
	Short:       "Mostra les dades d'una factura",
	Long:        "Mostra les dades d'uan factura indicant el seu codi",
	Example:     `   mostraFactura f-3945     Mostra les dades de la factura F-3945`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	Aliases: []string{
		"mfac",
		"mostrafactura", "mostra-factura",
		"mostrarFactura", "mostrarfactura", "mostrar-factura",
		"dinv",
		"displayInvoice", "displayinvoice", "display-invoice",
	},
	Args: ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		invoiceCode := strings.ToUpper(args[0])
		manager := adm.NewDisplayManager()
		_, err := manager.DisplayInvoice(invoiceCode)
		return err
	},
}

func init() {
	rootCmd.AddCommand(displayInvoiceCmd)
}
package cmd

import (
	"fmt"
	v "github.com/mehmetkule/vowels/vowels"
	"github.com/spf13/cobra"
)

const (
	ENGLISH = "AaEeIiOoUu"
	TURKISH = "ÄÅAaEeIıİiOoÖöUuÜüÃ¼"
	KURDISH = "AEÊIÎOUÛaeêiîouû"
)

type Config struct {
	Input      string `config:"input"`
	Output     string `config:"output"`
	VowelsType string `config:"vowelsType"`
}

var cfg Config

func init() {
	flags := RootCmd.PersistentFlags()
	flags.StringVar(&cfg.Input, "input", "", "reading file name")
	flags.StringVar(&cfg.Output, "output", "", "write file name")
	flags.StringVar(&cfg.VowelsType, "vowels", "t", "Vowels Type : ENGLISH = e ,TURKISH = t , KURDISH = k ")
	RootCmd.AddCommand(Vowels)
}

var Vowels = &cobra.Command{
	Use:   "vowels",
	Short: "Shor file reader data",
	Args:  cobra.MaximumNArgs(1),
	Run:   Data,
}

func Data(cmd *cobra.Command, args []string) {

	var vType string
	switch {
	case cfg.VowelsType == "e":
		vType = ENGLISH
	case cfg.VowelsType == "t":
		vType = TURKISH
	case cfg.VowelsType == "k":
		vType = KURDISH
	default:
		fmt.Printf("Invalid vowels type %s: please  ENGLISH = 'e' ,TURKISH = 't' , KURDISH = 'k' ", vType)
	}

	err := v.Delete(cfg.Input, cfg.Output, vType)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
}

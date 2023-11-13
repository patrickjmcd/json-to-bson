package main

import (
	"github.com/patrickjmcd/json-to-bson/pkg"
	"github.com/spf13/cobra"
	"log"
)

var (
	filename   string
	jsonString string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "Filename to convert")
	rootCmd.PersistentFlags().StringVarP(&jsonString, "json", "j", "", "JSON string to convert")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "json-to-bson",
	Short: "JSON to BSON Converter",
	Long:  `Convert JSON to BSON`,
	Run: func(cmd *cobra.Command, args []string) {
		if (filename == "" && jsonString == "") || (filename != "" && jsonString != "") {
			log.Fatal("Must provide either a filename or a JSON string")
		}
		if filename != "" {
			doc, err := pkg.ConvertFile(filename)
			if err != nil {
				log.Fatalf("error converting file %s: %v", filename, err)
			}
			log.Printf("%+v", doc)
		} else {
			doc, err := pkg.ConvertString([]byte(jsonString))
			if err != nil {
				log.Fatalf("error converting JSON string: %v", err)
			}
			log.Printf("%+v", doc)
		}
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("error executing root command: %v", err)
	}
}

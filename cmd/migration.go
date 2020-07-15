package cmd

import (
	"github.com/diptomondal007/go_clean_arch/conn"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	dialect = "mysql"
)

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "clean arch migration",
	Long:  `clean arch migration`,
	Args:  cobra.MinimumNArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if _, err := conn.GetDB(); err != nil {
			log.Println("MigrationDBConnErr:", err.Error())
			return err
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		db , _ := conn.GetDB()
		defer db.Close()

		wd, err := os.Getwd()
		if err != nil {
			log.Fatalln("MigrationGetWDError:", err.Error())
		}

		dir := wd + "/migrations"

		err = goose.SetDialect(dialect)
		if err != nil {
			log.Fatalln("MigrationSetDialectError:", err.Error())
		}

		if err := goose.Run(args[0], db.DB.DB(), dir, args...); err != nil {
			log.Fatalln("MigrationRunErr:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}

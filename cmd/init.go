package cmd

import (
	"fmt"

	"github.com/seqDemo/sqll"
	"github.com/seqDemo/util"
	"github.com/spf13/cobra"
)

func NewInitEnvCommand() *cobra.Command {
	//init database and table
	var dsn string
	cmd := &cobra.Command{
		Use:   "seq",
		Short: "create database and table ",
		RunE: func(cmd *cobra.Command, args []string) error {
			//var err error
			db, err := sqll.ConnectDB(dsn)
			if err != nil {
				return err
			}
			err = sqll.CreateDatabase(db, util.DatabaseSQL)
			if err != nil {
				return err
			}
			err = sqll.CreateTable(db, util.TableSQL)
			if err != nil {
				return err
			}
			err = sqll.QueryWithNoResult(db, util.RecordSQL)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&dsn, "dsn", "d", "", "database server dsn")
	return cmd
}

func NewInitCommand() *cobra.Command {
	//add sub command replay
	cmd := &cobra.Command{
		Use:   "init",
		Short: "init run environment",
	}
	fmt.Println("in init")
	cmd.AddCommand(NewInitEnvCommand())
	return cmd
}

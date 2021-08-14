package sub

import (
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile     string
	read        bool
	write       bool
	deleteFlag  bool
	sectionName string
	key         string
	value       string
)

func init() {

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", "", "ini file")
	rootCmd.PersistentFlags().BoolVarP(&read, "read", "r", false, "read the value of the key in the section of the file")
	rootCmd.PersistentFlags().BoolVarP(&write, "write", "w", false, "write the value of the key in the section of the file")
	rootCmd.PersistentFlags().BoolVarP(&deleteFlag, "deleteFlag", "D", false, "deleteFlag a whole section or a specified key of the file")
	rootCmd.PersistentFlags().StringVarP(&sectionName, "section", "s", "", "section title in the ini file")
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "key in the section")
	rootCmd.PersistentFlags().StringVarP(&value, "value", "v", "", "set the value of the key")

}

var rootCmd = &cobra.Command{
	Use:   "rwini",
	Short: "rwini is a read or write ini file tool",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := runClient()
		if err != nil {
			fmt.Println("Error info: ", err)
			_ = cmd.Usage()
			os.Exit(1)
		}
		return nil
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

func runClient() error {
	if len(cfgFile) == 0 {
		return errors.New(fmt.Sprintf("input ini file name"))
	}

	if (write && read) || (write && deleteFlag) || (read && deleteFlag) || !(read || write || deleteFlag) {
		return errors.New(fmt.Sprintf("read/write/deleteFlag input illegal, r: %v, w: %v, D:%v ", read, write, deleteFlag))
	}

	if write {
		return writeIni()
	}

	if read {
		return readIni()
	}

	if deleteFlag {
		return deleteIni()
	}

	return nil
}

func readIni() error {
	if len(sectionName) == 0 || len(key) == 0 {
		return errors.New(fmt.Sprintf("input illegal, sectionname: %s, key: %s", sectionName, key))
	}
	cfg, err := ini.Load(cfgFile)
	if err != nil {
		return err
	}
	fmt.Println(cfg.Section(sectionName).Key(key).String())
	return nil
}

func writeIni() error {
	if len(sectionName) == 0 || len(key) == 0 || len(value) == 0 {
		return errors.New(fmt.Sprintf("input illegal, sectionname: %s, key: %s, value: %s", sectionName, key, value))
	}
	cfg, err := ini.Load(cfgFile)
	if err != nil {
		return err
	}
	cfg.Section(sectionName).Key(key).SetValue(value)
	err = cfg.SaveTo(cfgFile)
	if err != nil {
		return err
	}
	return nil
}

func deleteIni() error {
	if len(sectionName) == 0 {
		return errors.New(fmt.Sprintf("input illegal, sectionname: %s", sectionName))
	}
	cfg, err := ini.Load(cfgFile)
	if err != nil {
		return err
	}
	if len(key) == 0 {
		cfg.DeleteSection(sectionName)
	} else {
		cfg.Section(sectionName).DeleteKey(key)
	}

	err = cfg.SaveTo(cfgFile)
	if err != nil {
		return err
	}
	return nil
}

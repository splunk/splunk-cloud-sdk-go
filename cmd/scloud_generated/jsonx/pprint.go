package jsonx

import (
	"encoding/json"
	"strings"

	"github.com/spf13/cobra"
)

func pprint(cmd *cobra.Command, value interface{}, asErr bool) {
	if value == nil {
		return
	}
	switch vt := value.(type) {
	case string:
		if asErr {
			cmd.PrintErr(vt)
		} else {
			cmd.Print(vt)
		}
		if !strings.HasSuffix(vt, "\n") {
			if asErr {
				cmd.PrintErrln()
			} else {
				cmd.Println()
			}
		}
	default:
		var encoder *json.Encoder
		if asErr {
			encoder = json.NewEncoder(cmd.OutOrStderr())
		} else {
			encoder = json.NewEncoder(cmd.OutOrStdout())
		}
		encoder.SetIndent("", "    ")
		err := encoder.Encode(value)
		if err != nil {
			panic("jsonx.pprint: encoder.Encode(value) error: %s" + err.Error())
		}
	}
}

func Pprint(cmd *cobra.Command, value interface{}) {
	pprint(cmd, value, false)
}

func PprintErr(cmd *cobra.Command, value interface{}) {
	pprint(cmd, value, true)
}

package jsonx

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Pprint(cmd *cobra.Command, value interface{}) {
	if value == nil {
		return
	}
	switch vt := value.(type) {
	case string:
		cmd.Print(vt)
		if !strings.HasSuffix(vt, "\n") {
			cmd.Println()
		}
	default:
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "    ")
		err := encoder.Encode(value)
		if err != nil {
			panic("jsonx.Pprint: encoder.Encode(value) error: %s" + err.Error())
		}
	}
}

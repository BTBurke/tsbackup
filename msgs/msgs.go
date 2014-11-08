package msgs

import "fmt"

const NoArgsHelp string = `If this is your first time running tsbackup, try tsbackup setup.  For help, run tsbackup --help.`

const Help string = `tsbackup help message container`

func UnknownCmd(cmd string) string {
	return fmt.Sprintf("Unknown command: %s. For help, try tsbackup --help.", cmd)
}

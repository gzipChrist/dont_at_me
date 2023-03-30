package components

import (
	"fmt"
	"github.com/gzipchrist/dont_at_me/pkg/style"
)

var Header = `
      |               /        ____                   
    __|   __   _  _    _|_    / __,\    _  _  _    _  
   /  |  /  \_/ |/ |    |    | /  | |  / |/ |/ |  |/  
   \_/|_/\__/   |  |_/  |_/  | \_/|/     |  |  |_/|__/
                              \____/
`

var Prompt = fmt.Sprintf("  Enter a username to check availability %s", style.Dim.Colorize("[q to quit]"))

var TextInput = "  @ "

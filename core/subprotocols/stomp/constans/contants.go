package constans

import "strings"

const EndLineStringRune = '\n'

const TerminalByte byte = 0

var SupportVersion = []string{"v10.stomp", "v11.stomp"}
var SupportVersionInString = strings.Join(SupportVersion, ",")

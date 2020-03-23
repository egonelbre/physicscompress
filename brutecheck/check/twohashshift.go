package check

func TwoHash_XorXor_Shift015(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 1)) ^ (name[1] << 5) {
	case 0:
		return name == "switch"
	case 16:
		return name == "if"
	case 27:
		return name == "interface"
	case 35:
		return name == "const"
	case 42:
		return name == "goto"
	case 44:
		return name == "go"
	case 46:
		return name == "continue"
	case 47:
		return name == "for"
	case 64:
		return name == "select"
	case 66:
		return name == "return"
	case 78:
		return name == "else"
	case 96:
		return name == "struct"
	case 104:
		return name == "func"
	case 109:
		return name == "defer"
	case 111:
		return name == "default"
	case 116:
		return name == "import"
	case 129:
		return name == "break"
	case 193:
		return name == "range"
	case 194:
		return name == "chan"
	case 199:
		return name == "package"
	case 204:
		return name == "type"
	case 207:
		return name == "var"
	case 226:
		return name == "case"
	case 231:
		return name == "fallthrough"
	case 249:
		return name == "map"
	}
	return false
}

func TwoHash_XorXor_Shift016(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 1)) ^ (name[1] << 6) {
	case 3:
		return name == "const"
	case 10:
		return name == "goto"
	case 12:
		return name == "go"
	case 14:
		return name == "continue"
	case 15:
		return name == "for"
	case 32:
		return name == "switch"
	case 65:
		return name == "break"
	case 80:
		return name == "if"
	case 91:
		return name == "interface"
	case 130:
		return name == "case"
	case 135:
		return name == "fallthrough"
	case 136:
		return name == "func"
	case 141:
		return name == "defer"
	case 143:
		return name == "default"
	case 148:
		return name == "import"
	case 153:
		return name == "map"
	case 160:
		return name == "select"
	case 161:
		return name == "range"
	case 162:
		return name == "return"
	case 167:
		return name == "package"
	case 172:
		return name == "type"
	case 175:
		return name == "var"
	case 194:
		return name == "chan"
	case 206:
		return name == "else"
	case 224:
		return name == "struct"
	}
	return false
}

func TwoHash_XorXor_Shift021(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 2)) ^ (name[1] << 1) {
	case 0:
		return name == "select"
	case 4:
		return name == "return"
	case 5:
		return name == "package"
	case 15:
		return name == "range"
	case 25:
		return name == "var"
	case 34:
		return name == "struct"
	case 36:
		return name == "switch"
	case 38:
		return name == "type"
	case 64:
		return name == "go"
	case 69:
		return name == "for"
	case 70:
		return name == "goto"
	case 72:
		return name == "else"
	case 74:
		return name == "case"
	case 81:
		return name == "fallthrough"
	case 87:
		return name == "const"
	case 88:
		return name == "chan"
	case 90:
		return name == "continue"
	case 93:
		return name == "default"
	case 95:
		return name == "defer"
	case 105:
		return name == "break"
	case 106:
		return name == "if"
	case 113:
		return name == "interface"
	case 117:
		return name == "map"
	case 118:
		return name == "func"
	case 120:
		return name == "import"
	}
	return false
}

func TwoHash_XorXor_Shift024(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 2)) ^ (name[1] << 4) {
	case 8:
		return name == "chan"
	case 68:
		return name == "type"
	case 77:
		return name == "interface"
	case 80:
		return name == "else"
	case 104:
		return name == "goto"
	case 107:
		return name == "for"
	case 110:
		return name == "go"
	case 114:
		return name == "import"
	case 116:
		return name == "continue"
	case 121:
		return name == "const"
	case 131:
		return name == "fallthrough"
	case 138:
		return name == "struct"
	case 152:
		return name == "case"
	case 154:
		return name == "select"
	case 158:
		return name == "return"
	case 167:
		return name == "map"
	case 173:
		return name == "break"
	case 186:
		return name == "switch"
	case 197:
		return name == "defer"
	case 198:
		return name == "if"
	case 199:
		return name == "default"
	case 203:
		return name == "var"
	case 204:
		return name == "func"
	case 215:
		return name == "package"
	case 221:
		return name == "range"
	}
	return false
}

func TwoHash_XorXor_Shift025(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 2)) ^ (name[1] << 5) {
	case 2:
		return name == "import"
	case 16:
		return name == "else"
	case 42:
		return name == "switch"
	case 53:
		return name == "defer"
	case 55:
		return name == "default"
	case 60:
		return name == "func"
	case 74:
		return name == "struct"
	case 100:
		return name == "continue"
	case 102:
		return name == "if"
	case 105:
		return name == "const"
	case 106:
		return name == "select"
	case 109:
		return name == "interface"
	case 110:
		return name == "return"
	case 120:
		return name == "goto"
	case 123:
		return name == "for"
	case 126:
		return name == "go"
	case 136:
		return name == "chan"
	case 151:
		return name == "map"
	case 168:
		return name == "case"
	case 179:
		return name == "fallthrough"
	case 205:
		return name == "break"
	case 231:
		return name == "package"
	case 237:
		return name == "range"
	case 244:
		return name == "type"
	case 251:
		return name == "var"
	}
	return false
}

func TwoHash_XorXor_Shift026(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 2)) ^ (name[1] << 6) {
	case 10:
		return name == "switch"
	case 13:
		return name == "break"
	case 38:
		return name == "if"
	case 45:
		return name == "interface"
	case 68:
		return name == "continue"
	case 73:
		return name == "const"
	case 88:
		return name == "goto"
	case 91:
		return name == "for"
	case 94:
		return name == "go"
	case 135:
		return name == "package"
	case 136:
		return name == "chan"
	case 138:
		return name == "select"
	case 141:
		return name == "range"
	case 142:
		return name == "return"
	case 144:
		return name == "else"
	case 148:
		return name == "type"
	case 155:
		return name == "var"
	case 200:
		return name == "case"
	case 202:
		return name == "struct"
	case 211:
		return name == "fallthrough"
	case 213:
		return name == "defer"
	case 215:
		return name == "default"
	case 220:
		return name == "func"
	case 226:
		return name == "import"
	case 247:
		return name == "map"
	}
	return false
}

func TwoHash_XorXor_Shift032(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 3)) ^ (name[1] << 2) {
	case 2:
		return name == "return"
	case 3:
		return name == "package"
	case 10:
		return name == "select"
	case 17:
		return name == "range"
	case 55:
		return name == "var"
	case 64:
		return name == "type"
	case 66:
		return name == "switch"
	case 78:
		return name == "struct"
	case 128:
		return name == "goto"
	case 134:
		return name == "go"
	case 143:
		return name == "for"
	case 152:
		return name == "case"
	case 156:
		return name == "else"
	case 161:
		return name == "const"
	case 172:
		return name == "continue"
	case 177:
		return name == "defer"
	case 179:
		return name == "default"
	case 188:
		return name == "chan"
	case 191:
		return name == "fallthrough"
	case 210:
		return name == "if"
	case 221:
		return name == "break"
	case 224:
		return name == "func"
	case 239:
		return name == "map"
	case 249:
		return name == "interface"
	case 250:
		return name == "import"
	}
	return false
}

func TwoHash_XorXor_Shift034(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 3)) ^ (name[1] << 4) {
	case 12:
		return name == "case"
	case 42:
		return name == "if"
	case 43:
		return name == "fallthrough"
	case 52:
		return name == "type"
	case 53:
		return name == "break"
	case 100:
		return name == "func"
	case 117:
		return name == "defer"
	case 119:
		return name == "default"
	case 123:
		return name == "map"
	case 133:
		return name == "range"
	case 151:
		return name == "package"
	case 156:
		return name == "chan"
	case 158:
		return name == "import"
	case 161:
		return name == "interface"
	case 163:
		return name == "var"
	case 195:
		return name == "for"
	case 198:
		return name == "return"
	case 202:
		return name == "go"
	case 204:
		return name == "goto"
	case 206:
		return name == "select"
	case 222:
		return name == "struct"
	case 224:
		return name == "continue"
	case 236:
		return name == "else"
	case 237:
		return name == "const"
	case 238:
		return name == "switch"
	}
	return false
}

func TwoHash_XorXor_Shift035(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 3)) ^ (name[1] << 5) {
	case 27:
		return name == "fallthrough"
	case 28:
		return name == "chan"
	case 30:
		return name == "struct"
	case 54:
		return name == "return"
	case 60:
		return name == "case"
	case 62:
		return name == "select"
	case 75:
		return name == "map"
	case 85:
		return name == "break"
	case 126:
		return name == "switch"
	case 129:
		return name == "interface"
	case 132:
		return name == "type"
	case 133:
		return name == "defer"
	case 135:
		return name == "default"
	case 138:
		return name == "if"
	case 147:
		return name == "var"
	case 148:
		return name == "func"
	case 167:
		return name == "package"
	case 172:
		return name == "else"
	case 181:
		return name == "range"
	case 211:
		return name == "for"
	case 218:
		return name == "go"
	case 220:
		return name == "goto"
	case 238:
		return name == "import"
	case 240:
		return name == "continue"
	case 253:
		return name == "const"
	}
	return false
}

func TwoHash_XorXor_Shift041(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 4)) ^ (name[1] << 1) {
	case 17:
		return name == "map"
	case 69:
		return name == "interface"
	case 76:
		return name == "import"
	case 94:
		return name == "if"
	case 140:
		return name == "else"
	case 141:
		return name == "default"
	case 142:
		return name == "func"
	case 143:
		return name == "defer"
	case 161:
		return name == "var"
	case 169:
		return name == "fallthrough"
	case 170:
		return name == "goto"
	case 172:
		return name == "go"
	case 182:
		return name == "type"
	case 189:
		return name == "for"
	case 193:
		return name == "break"
	case 197:
		return name == "package"
	case 216:
		return name == "switch"
	case 222:
		return name == "struct"
	case 228:
		return name == "chan"
	case 230:
		return name == "continue"
	case 231:
		return name == "range"
	case 235:
		return name == "const"
	case 236:
		return name == "return"
	case 246:
		return name == "case"
	case 252:
		return name == "select"
	}
	return false
}

func TwoHash_XorXor_Shift043(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 4)) ^ (name[1] << 3) {
	case 10:
		return name == "go"
	case 12:
		return name == "goto"
	case 14:
		return name == "return"
	case 15:
		return name == "package"
	case 27:
		return name == "for"
	case 30:
		return name == "select"
	case 45:
		return name == "range"
	case 52:
		return name == "else"
	case 60:
		return name == "case"
	case 64:
		return name == "continue"
	case 77:
		return name == "const"
	case 99:
		return name == "fallthrough"
	case 107:
		return name == "var"
	case 109:
		return name == "defer"
	case 111:
		return name == "default"
	case 116:
		return name == "chan"
	case 140:
		return name == "type"
	case 142:
		return name == "switch"
	case 150:
		return name == "struct"
	case 162:
		return name == "if"
	case 181:
		return name == "break"
	case 204:
		return name == "func"
	case 219:
		return name == "map"
	case 233:
		return name == "interface"
	case 254:
		return name == "import"
	}
	return false
}

func TwoHash_XorXor_Shift045(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 4)) ^ (name[1] << 5) {
	case 5:
		return name == "range"
	case 20:
		return name == "case"
	case 39:
		return name == "package"
	case 52:
		return name == "chan"
	case 54:
		return name == "import"
	case 67:
		return name == "var"
	case 75:
		return name == "fallthrough"
	case 82:
		return name == "if"
	case 89:
		return name == "interface"
	case 100:
		return name == "type"
	case 101:
		return name == "break"
	case 131:
		return name == "for"
	case 134:
		return name == "return"
	case 146:
		return name == "go"
	case 148:
		return name == "goto"
	case 150:
		return name == "select"
	case 182:
		return name == "struct"
	case 196:
		return name == "func"
	case 212:
		return name == "else"
	case 213:
		return name == "const"
	case 214:
		return name == "switch"
	case 216:
		return name == "continue"
	case 229:
		return name == "defer"
	case 231:
		return name == "default"
	case 243:
		return name == "map"
	}
	return false
}

func TwoHash_XorXor_Shift046(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 4)) ^ (name[1] << 6) {
	case 4:
		return name == "type"
	case 5:
		return name == "defer"
	case 7:
		return name == "default"
	case 18:
		return name == "if"
	case 25:
		return name == "interface"
	case 35:
		return name == "var"
	case 36:
		return name == "func"
	case 43:
		return name == "fallthrough"
	case 52:
		return name == "chan"
	case 54:
		return name == "struct"
	case 71:
		return name == "package"
	case 84:
		return name == "else"
	case 101:
		return name == "range"
	case 102:
		return name == "return"
	case 116:
		return name == "case"
	case 118:
		return name == "select"
	case 147:
		return name == "map"
	case 163:
		return name == "for"
	case 165:
		return name == "break"
	case 178:
		return name == "go"
	case 180:
		return name == "goto"
	case 214:
		return name == "import"
	case 245:
		return name == "const"
	case 246:
		return name == "switch"
	case 248:
		return name == "continue"
	}
	return false
}

func TwoHash_XorXor_Shift050(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 5)) ^ (name[1] << 0) {
	case 3:
		return name == "select"
	case 5:
		return name == "case"
	case 7:
		return name == "continue"
	case 10:
		return name == "const"
	case 12:
		return name == "chan"
	case 17:
		return name == "switch"
	case 18:
		return name == "struct"
	case 35:
		return name == "return"
	case 36:
		return name == "range"
	case 55:
		return name == "break"
	case 68:
		return name == "if"
	case 71:
		return name == "interface"
	case 75:
		return name == "import"
	case 102:
		return name == "package"
	case 139:
		return name == "goto"
	case 141:
		return name == "go"
	case 162:
		return name == "var"
	case 170:
		return name == "fallthrough"
	case 172:
		return name == "for"
	case 177:
		return name == "func"
	case 194:
		return name == "map"
	case 200:
		return name == "else"
	case 224:
		return name == "defer"
	case 226:
		return name == "default"
	case 253:
		return name == "type"
	}
	return false
}

func TwoHash_XorXor_Shift051(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 5)) ^ (name[1] << 1) {
	case 1:
		return name == "var"
	case 9:
		return name == "fallthrough"
	case 29:
		return name == "for"
	case 46:
		return name == "func"
	case 58:
		return name == "goto"
	case 60:
		return name == "go"
	case 77:
		return name == "default"
	case 79:
		return name == "defer"
	case 97:
		return name == "map"
	case 118:
		return name == "type"
	case 124:
		return name == "else"
	case 135:
		return name == "range"
	case 136:
		return name == "switch"
	case 140:
		return name == "return"
	case 142:
		return name == "struct"
	case 161:
		return name == "break"
	case 166:
		return name == "case"
	case 172:
		return name == "select"
	case 180:
		return name == "chan"
	case 182:
		return name == "continue"
	case 187:
		return name == "const"
	case 197:
		return name == "package"
	case 238:
		return name == "if"
	case 245:
		return name == "interface"
	case 252:
		return name == "import"
	}
	return false
}

func TwoHash_XorXor_Shift060(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 6)) ^ (name[1] << 0) {
	case 34:
		return name == "map"
	case 36:
		return name == "if"
	case 39:
		return name == "interface"
	case 40:
		return name == "else"
	case 43:
		return name == "import"
	case 96:
		return name == "defer"
	case 98:
		return name == "default"
	case 102:
		return name == "package"
	case 125:
		return name == "type"
	case 163:
		return name == "select"
	case 165:
		return name == "case"
	case 167:
		return name == "continue"
	case 170:
		return name == "const"
	case 171:
		return name == "goto"
	case 172:
		return name == "chan"
	case 173:
		return name == "go"
	case 177:
		return name == "switch"
	case 178:
		return name == "struct"
	case 226:
		return name == "var"
	case 227:
		return name == "return"
	case 228:
		return name == "range"
	case 234:
		return name == "fallthrough"
	case 236:
		return name == "for"
	case 241:
		return name == "func"
	case 247:
		return name == "break"
	}
	return false
}

func TwoHash_XorXor_Shift062(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 6)) ^ (name[1] << 2) {
	case 1:
		return name == "range"
	case 7:
		return name == "var"
	case 15:
		return name == "fallthrough"
	case 18:
		return name == "return"
	case 22:
		return name == "struct"
	case 26:
		return name == "switch"
	case 63:
		return name == "for"
	case 64:
		return name == "case"
	case 77:
		return name == "break"
	case 80:
		return name == "func"
	case 82:
		return name == "select"
	case 100:
		return name == "chan"
	case 116:
		return name == "continue"
	case 120:
		return name == "goto"
	case 121:
		return name == "const"
	case 126:
		return name == "go"
	case 131:
		return name == "package"
	case 145:
		return name == "defer"
	case 147:
		return name == "default"
	case 199:
		return name == "map"
	case 218:
		return name == "if"
	case 224:
		return name == "type"
	case 241:
		return name == "interface"
	case 242:
		return name == "import"
	case 244:
		return name == "else"
	}
	return false
}

func TwoHash_XorXor_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 7)) ^ (name[1] << 3) {
	case 3:
		return name == "fallthrough"
	case 11:
		return name == "var"
	case 13:
		return name == "range"
	case 15:
		return name == "package"
	case 38:
		return name == "struct"
	case 45:
		return name == "defer"
	case 46:
		return name == "return"
	case 47:
		return name == "default"
	case 62:
		return name == "switch"
	case 123:
		return name == "for"
	case 139:
		return name == "map"
	case 140:
		return name == "case"
	case 149:
		return name == "break"
	case 172:
		return name == "func"
	case 174:
		return name == "select"
	case 178:
		return name == "if"
	case 196:
		return name == "chan"
	case 204:
		return name == "type"
	case 228:
		return name == "else"
	case 238:
		return name == "import"
	case 240:
		return name == "continue"
	case 249:
		return name == "interface"
	case 250:
		return name == "go"
	case 252:
		return name == "goto"
	case 253:
		return name == "const"
	}
	return false
}

func TwoHash_XorXor_Shift130(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) ^ (name[0] << 3)) ^ (name[1] << 0) {
	case 15:
		return name == "map"
	case 41:
		return name == "import"
	case 42:
		return name == "if"
	case 52:
		return name == "interface"
	case 71:
		return name == "fallthrough"
	case 75:
		return name == "default"
	case 76:
		return name == "else"
	case 77:
		return name == "func"
	case 79:
		return name == "defer"
	case 83:
		return name == "go"
	case 89:
		return name == "for"
	case 95:
		return name == "goto"
	case 103:
		return name == "continue"
	case 104:
		return name == "break"
	case 113:
		return name == "case"
	case 120:
		return name == "chan"
	case 125:
		return name == "const"
	case 209:
		return name == "type"
	case 215:
		return name == "var"
	case 224:
		return name == "struct"
	case 227:
		return name == "switch"
	case 239:
		return name == "package"
	case 241:
		return name == "select"
	case 249:
		return name == "return"
	case 251:
		return name == "range"
	}
	return false
}

func TwoHash_XorXor_Shift132(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) ^ (name[0] << 3)) ^ (name[1] << 2) {
	case 0:
		return name == "select"
	case 8:
		return name == "return"
	case 10:
		return name == "package"
	case 30:
		return name == "range"
	case 50:
		return name == "var"
	case 68:
		return name == "struct"
	case 72:
		return name == "switch"
	case 76:
		return name == "type"
	case 128:
		return name == "go"
	case 138:
		return name == "for"
	case 140:
		return name == "goto"
	case 144:
		return name == "else"
	case 148:
		return name == "case"
	case 162:
		return name == "fallthrough"
	case 174:
		return name == "const"
	case 176:
		return name == "chan"
	case 180:
		return name == "continue"
	case 186:
		return name == "default"
	case 190:
		return name == "defer"
	case 210:
		return name == "break"
	case 212:
		return name == "if"
	case 226:
		return name == "interface"
	case 234:
		return name == "map"
	case 236:
		return name == "func"
	case 240:
		return name == "import"
	}
	return false
}

func TwoHash_XorXor_Shift135(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) ^ (name[0] << 3)) ^ (name[1] << 5) {
	case 6:
		return name == "fallthrough"
	case 16:
		return name == "chan"
	case 20:
		return name == "struct"
	case 48:
		return name == "case"
	case 52:
		return name == "select"
	case 60:
		return name == "return"
	case 78:
		return name == "map"
	case 90:
		return name == "break"
	case 116:
		return name == "switch"
	case 136:
		return name == "type"
	case 138:
		return name == "defer"
	case 140:
		return name == "if"
	case 142:
		return name == "default"
	case 150:
		return name == "var"
	case 152:
		return name == "func"
	case 154:
		return name == "interface"
	case 160:
		return name == "else"
	case 174:
		return name == "package"
	case 186:
		return name == "range"
	case 208:
		return name == "goto"
	case 214:
		return name == "for"
	case 220:
		return name == "go"
	case 228:
		return name == "import"
	case 232:
		return name == "continue"
	case 242:
		return name == "const"
	}
	return false
}

func TwoHash_XorXor_Shift143(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) ^ (name[0] << 4)) ^ (name[1] << 3) {
	case 0:
		return name == "goto"
	case 4:
		return name == "return"
	case 6:
		return name == "package"
	case 12:
		return name == "go"
	case 20:
		return name == "select"
	case 30:
		return name == "for"
	case 34:
		return name == "range"
	case 48:
		return name == "case"
	case 56:
		return name == "else"
	case 66:
		return name == "const"
	case 88:
		return name == "continue"
	case 98:
		return name == "defer"
	case 102:
		return name == "default"
	case 110:
		return name == "var"
	case 120:
		return name == "chan"
	case 126:
		return name == "fallthrough"
	case 128:
		return name == "type"
	case 132:
		return name == "switch"
	case 156:
		return name == "struct"
	case 164:
		return name == "if"
	case 186:
		return name == "break"
	case 192:
		return name == "func"
	case 222:
		return name == "map"
	case 242:
		return name == "interface"
	case 244:
		return name == "import"
	}
	return false
}

func TwoHash_XorXor_Shift145(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) ^ (name[0] << 4)) ^ (name[1] << 5) {
	case 10:
		return name == "range"
	case 24:
		return name == "case"
	case 46:
		return name == "package"
	case 56:
		return name == "chan"
	case 60:
		return name == "import"
	case 66:
		return name == "interface"
	case 70:
		return name == "var"
	case 84:
		return name == "if"
	case 86:
		return name == "fallthrough"
	case 104:
		return name == "type"
	case 106:
		return name == "break"
	case 134:
		return name == "for"
	case 140:
		return name == "return"
	case 148:
		return name == "go"
	case 152:
		return name == "goto"
	case 156:
		return name == "select"
	case 188:
		return name == "struct"
	case 192:
		return name == "continue"
	case 200:
		return name == "func"
	case 216:
		return name == "else"
	case 218:
		return name == "const"
	case 220:
		return name == "switch"
	case 234:
		return name == "defer"
	case 238:
		return name == "default"
	case 246:
		return name == "map"
	}
	return false
}

func TwoHash_XorXor_Shift146(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) ^ (name[0] << 4)) ^ (name[1] << 6) {
	case 2:
		return name == "interface"
	case 8:
		return name == "type"
	case 10:
		return name == "defer"
	case 14:
		return name == "default"
	case 20:
		return name == "if"
	case 38:
		return name == "var"
	case 40:
		return name == "func"
	case 54:
		return name == "fallthrough"
	case 56:
		return name == "chan"
	case 60:
		return name == "struct"
	case 78:
		return name == "package"
	case 88:
		return name == "else"
	case 106:
		return name == "range"
	case 108:
		return name == "return"
	case 120:
		return name == "case"
	case 124:
		return name == "select"
	case 150:
		return name == "map"
	case 166:
		return name == "for"
	case 170:
		return name == "break"
	case 180:
		return name == "go"
	case 184:
		return name == "goto"
	case 220:
		return name == "import"
	case 224:
		return name == "continue"
	case 250:
		return name == "const"
	case 252:
		return name == "switch"
	}
	return false
}

func TwoHash_XorXor_Shift161(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) ^ (name[0] << 6)) ^ (name[1] << 1) {
	case 6:
		return name == "select"
	case 10:
		return name == "case"
	case 14:
		return name == "continue"
	case 20:
		return name == "const"
	case 22:
		return name == "goto"
	case 24:
		return name == "chan"
	case 26:
		return name == "go"
	case 34:
		return name == "switch"
	case 36:
		return name == "struct"
	case 68:
		return name == "var"
	case 70:
		return name == "return"
	case 72:
		return name == "range"
	case 84:
		return name == "fallthrough"
	case 88:
		return name == "for"
	case 98:
		return name == "func"
	case 110:
		return name == "break"
	case 132:
		return name == "map"
	case 136:
		return name == "if"
	case 142:
		return name == "interface"
	case 144:
		return name == "else"
	case 150:
		return name == "import"
	case 192:
		return name == "defer"
	case 196:
		return name == "default"
	case 204:
		return name == "package"
	case 250:
		return name == "type"
	}
	return false
}

func TwoHash_XorXor_Shift240(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) ^ (name[0] << 4)) ^ (name[1] << 0) {
	case 3:
		return name == "for"
	case 5:
		return name == "func"
	case 13:
		return name == "var"
	case 15:
		return name == "goto"
	case 23:
		return name == "go"
	case 41:
		return name == "type"
	case 44:
		return name == "else"
	case 45:
		return name == "fallthrough"
	case 49:
		return name == "defer"
	case 57:
		return name == "default"
	case 65:
		return name == "case"
	case 70:
		return name == "break"
	case 72:
		return name == "chan"
	case 75:
		return name == "const"
	case 77:
		return name == "select"
	case 85:
		return name == "range"
	case 92:
		return name == "struct"
	case 93:
		return name == "return"
	case 95:
		return name == "switch"
	case 125:
		return name == "package"
	case 127:
		return name == "continue"
	case 189:
		return name == "map"
	case 218:
		return name == "interface"
	case 229:
		return name == "import"
	case 254:
		return name == "if"
	}
	return false
}

func TwoHash_XorXor_Shift246(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) ^ (name[0] << 4)) ^ (name[1] << 6) {
	case 12:
		return name == "fallthrough"
	case 16:
		return name == "type"
	case 20:
		return name == "defer"
	case 24:
		return name == "if"
	case 28:
		return name == "default"
	case 32:
		return name == "chan"
	case 40:
		return name == "struct"
	case 44:
		return name == "var"
	case 48:
		return name == "func"
	case 52:
		return name == "interface"
	case 64:
		return name == "else"
	case 92:
		return name == "package"
	case 96:
		return name == "case"
	case 104:
		return name == "select"
	case 116:
		return name == "range"
	case 120:
		return name == "return"
	case 156:
		return name == "map"
	case 160:
		return name == "goto"
	case 172:
		return name == "for"
	case 180:
		return name == "break"
	case 184:
		return name == "go"
	case 200:
		return name == "import"
	case 208:
		return name == "continue"
	case 228:
		return name == "const"
	case 232:
		return name == "switch"
	}
	return false
}

func TwoHash_XorXor_Shift250(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) ^ (name[0] << 5)) ^ (name[1] << 0) {
	case 12:
		return name == "struct"
	case 15:
		return name == "switch"
	case 17:
		return name == "case"
	case 24:
		return name == "chan"
	case 27:
		return name == "const"
	case 29:
		return name == "select"
	case 38:
		return name == "break"
	case 47:
		return name == "continue"
	case 53:
		return name == "range"
	case 61:
		return name == "return"
	case 78:
		return name == "if"
	case 85:
		return name == "import"
	case 106:
		return name == "interface"
	case 125:
		return name == "package"
	case 135:
		return name == "go"
	case 141:
		return name == "fallthrough"
	case 159:
		return name == "goto"
	case 163:
		return name == "for"
	case 165:
		return name == "func"
	case 173:
		return name == "var"
	case 205:
		return name == "map"
	case 220:
		return name == "else"
	case 233:
		return name == "type"
	case 241:
		return name == "defer"
	case 249:
		return name == "default"
	}
	return false
}

func TwoHash_XorXor_Shift260(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) ^ (name[0] << 6)) ^ (name[1] << 0) {
	case 10:
		return name == "interface"
	case 45:
		return name == "map"
	case 46:
		return name == "if"
	case 53:
		return name == "import"
	case 60:
		return name == "else"
	case 105:
		return name == "type"
	case 113:
		return name == "defer"
	case 121:
		return name == "default"
	case 125:
		return name == "package"
	case 143:
		return name == "continue"
	case 167:
		return name == "go"
	case 172:
		return name == "struct"
	case 175:
		return name == "switch"
	case 177:
		return name == "case"
	case 184:
		return name == "chan"
	case 187:
		return name == "const"
	case 189:
		return name == "select"
	case 191:
		return name == "goto"
	case 205:
		return name == "fallthrough"
	case 227:
		return name == "for"
	case 229:
		return name == "func"
	case 230:
		return name == "break"
	case 237:
		return name == "var"
	case 245:
		return name == "range"
	case 253:
		return name == "return"
	}
	return false
}

func TwoHash_XorXor_Shift303(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) ^ (name[0] << 0)) ^ (name[1] << 3) {
	case 3:
		return name == "chan"
	case 6:
		return name == "for"
	case 15:
		return name == "go"
	case 37:
		return name == "else"
	case 49:
		return name == "import"
	case 51:
		return name == "const"
	case 54:
		return name == "fallthrough"
	case 63:
		return name == "goto"
	case 64:
		return name == "package"
	case 73:
		return name == "if"
	case 75:
		return name == "case"
	case 81:
		return name == "interface"
	case 82:
		return name == "range"
	case 91:
		return name == "continue"
	case 100:
		return name == "defer"
	case 102:
		return name == "var"
	case 106:
		return name == "return"
	case 107:
		return name == "select"
	case 116:
		return name == "default"
	case 125:
		return name == "map"
	case 156:
		return name == "type"
	case 218:
		return name == "break"
	case 227:
		return name == "struct"
	case 238:
		return name == "func"
	case 251:
		return name == "switch"
	}
	return false
}

func TwoHash_XorXor_Shift304(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) ^ (name[0] << 0)) ^ (name[1] << 4) {
	case 3:
		return name == "struct"
	case 12:
		return name == "default"
	case 18:
		return name == "return"
	case 19:
		return name == "select"
	case 22:
		return name == "func"
	case 25:
		return name == "if"
	case 28:
		return name == "defer"
	case 46:
		return name == "fallthrough"
	case 51:
		return name == "switch"
	case 74:
		return name == "range"
	case 83:
		return name == "case"
	case 88:
		return name == "package"
	case 101:
		return name == "map"
	case 106:
		return name == "break"
	case 126:
		return name == "var"
	case 133:
		return name == "else"
	case 135:
		return name == "go"
	case 137:
		return name == "import"
	case 142:
		return name == "for"
	case 183:
		return name == "goto"
	case 187:
		return name == "const"
	case 193:
		return name == "interface"
	case 195:
		return name == "chan"
	case 196:
		return name == "type"
	case 211:
		return name == "continue"
	}
	return false
}

func TwoHash_XorXor_Shift351(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) ^ (name[0] << 5)) ^ (name[1] << 1) {
	case 6:
		return name == "for"
	case 10:
		return name == "func"
	case 26:
		return name == "var"
	case 30:
		return name == "goto"
	case 46:
		return name == "go"
	case 82:
		return name == "type"
	case 88:
		return name == "else"
	case 90:
		return name == "fallthrough"
	case 98:
		return name == "defer"
	case 114:
		return name == "default"
	case 122:
		return name == "map"
	case 130:
		return name == "case"
	case 140:
		return name == "break"
	case 144:
		return name == "chan"
	case 150:
		return name == "const"
	case 154:
		return name == "select"
	case 170:
		return name == "range"
	case 180:
		return name == "interface"
	case 184:
		return name == "struct"
	case 186:
		return name == "return"
	case 190:
		return name == "switch"
	case 202:
		return name == "import"
	case 250:
		return name == "package"
	case 252:
		return name == "if"
	case 254:
		return name == "continue"
	}
	return false
}

func TwoHash_XorXor_Shift361(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) ^ (name[0] << 6)) ^ (name[1] << 1) {
	case 14:
		return name == "go"
	case 24:
		return name == "struct"
	case 26:
		return name == "fallthrough"
	case 30:
		return name == "switch"
	case 34:
		return name == "case"
	case 48:
		return name == "chan"
	case 54:
		return name == "const"
	case 58:
		return name == "select"
	case 62:
		return name == "goto"
	case 70:
		return name == "for"
	case 74:
		return name == "func"
	case 76:
		return name == "break"
	case 90:
		return name == "var"
	case 94:
		return name == "continue"
	case 106:
		return name == "range"
	case 122:
		return name == "return"
	case 154:
		return name == "map"
	case 156:
		return name == "if"
	case 170:
		return name == "import"
	case 184:
		return name == "else"
	case 210:
		return name == "type"
	case 212:
		return name == "interface"
	case 226:
		return name == "defer"
	case 242:
		return name == "default"
	case 250:
		return name == "package"
	}
	return false
}

func TwoHash_XorXor_Shift402(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) ^ (name[0] << 0)) ^ (name[1] << 2) {
	case 65:
		return name == "interface"
	case 82:
		return name == "fallthrough"
	case 95:
		return name == "continue"
	case 128:
		return name == "default"
	case 131:
		return name == "chan"
	case 132:
		return name == "package"
	case 134:
		return name == "return"
	case 135:
		return name == "select"
	case 143:
		return name == "const"
	case 149:
		return name == "else"
	case 155:
		return name == "goto"
	case 160:
		return name == "defer"
	case 166:
		return name == "range"
	case 167:
		return name == "case"
	case 189:
		return name == "import"
	case 194:
		return name == "var"
	case 195:
		return name == "struct"
	case 207:
		return name == "switch"
	case 208:
		return name == "type"
	case 209:
		return name == "if"
	case 217:
		return name == "map"
	case 234:
		return name == "for"
	case 242:
		return name == "func"
	case 250:
		return name == "break"
	case 251:
		return name == "go"
	}
	return false
}

func TwoHash_XorXor_Shift403(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) ^ (name[0] << 0)) ^ (name[1] << 3) {
	case 8:
		return name == "package"
	case 28:
		return name == "defer"
	case 42:
		return name == "range"
	case 43:
		return name == "case"
	case 46:
		return name == "for"
	case 58:
		return name == "return"
	case 59:
		return name == "select"
	case 60:
		return name == "default"
	case 63:
		return name == "go"
	case 69:
		return name == "else"
	case 75:
		return name == "const"
	case 78:
		return name == "var"
	case 85:
		return name == "map"
	case 95:
		return name == "goto"
	case 97:
		return name == "import"
	case 99:
		return name == "chan"
	case 121:
		return name == "if"
	case 137:
		return name == "interface"
	case 142:
		return name == "func"
	case 155:
		return name == "continue"
	case 162:
		return name == "break"
	case 171:
		return name == "switch"
	case 179:
		return name == "struct"
	case 222:
		return name == "fallthrough"
	case 252:
		return name == "type"
	}
	return false
}

func TwoHash_XorXor_Shift404(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) ^ (name[0] << 0)) ^ (name[1] << 4) {
	case 16:
		return name == "package"
	case 18:
		return name == "break"
	case 19:
		return name == "continue"
	case 25:
		return name == "interface"
	case 41:
		return name == "if"
	case 50:
		return name == "range"
	case 51:
		return name == "case"
	case 66:
		return name == "return"
	case 67:
		return name == "select"
	case 68:
		return name == "default"
	case 77:
		return name == "map"
	case 83:
		return name == "struct"
	case 86:
		return name == "var"
	case 99:
		return name == "switch"
	case 100:
		return name == "defer"
	case 118:
		return name == "func"
	case 163:
		return name == "chan"
	case 164:
		return name == "type"
	case 166:
		return name == "for"
	case 183:
		return name == "go"
	case 195:
		return name == "const"
	case 198:
		return name == "fallthrough"
	case 215:
		return name == "goto"
	case 217:
		return name == "import"
	case 229:
		return name == "else"
	}
	return false
}

func TwoHash_XorXor_Shift414(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) ^ (name[0] << 1)) ^ (name[1] << 4) {
	case 6:
		return name == "chan"
	case 12:
		return name == "for"
	case 30:
		return name == "go"
	case 56:
		return name == "type"
	case 74:
		return name == "else"
	case 98:
		return name == "import"
	case 102:
		return name == "const"
	case 108:
		return name == "fallthrough"
	case 126:
		return name == "goto"
	case 128:
		return name == "package"
	case 146:
		return name == "if"
	case 150:
		return name == "case"
	case 162:
		return name == "interface"
	case 164:
		return name == "range"
	case 180:
		return name == "break"
	case 182:
		return name == "continue"
	case 198:
		return name == "struct"
	case 200:
		return name == "defer"
	case 204:
		return name == "var"
	case 212:
		return name == "return"
	case 214:
		return name == "select"
	case 220:
		return name == "func"
	case 232:
		return name == "default"
	case 246:
		return name == "switch"
	case 250:
		return name == "map"
	}
	return false
}

func TwoHash_XorXor_Shift501(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) ^ (name[0] << 0)) ^ (name[1] << 1) {
	case 6:
		return name == "type"
	case 12:
		return name == "func"
	case 14:
		return name == "defer"
	case 16:
		return name == "range"
	case 29:
		return name == "const"
	case 33:
		return name == "case"
	case 38:
		return name == "break"
	case 51:
		return name == "chan"
	case 57:
		return name == "goto"
	case 61:
		return name == "else"
	case 78:
		return name == "default"
	case 82:
		return name == "package"
	case 91:
		return name == "struct"
	case 93:
		return name == "switch"
	case 115:
		return name == "import"
	case 120:
		return name == "return"
	case 121:
		return name == "select"
	case 149:
		return name == "interface"
	case 189:
		return name == "continue"
	case 196:
		return name == "fallthrough"
	case 207:
		return name == "map"
	case 212:
		return name == "var"
	case 216:
		return name == "for"
	case 229:
		return name == "if"
	case 249:
		return name == "go"
	}
	return false
}

func TwoHash_XorXor_Shift503(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) ^ (name[0] << 0)) ^ (name[1] << 3) {
	case 5:
		return name == "map"
	case 11:
		return name == "switch"
	case 14:
		return name == "fallthrough"
	case 19:
		return name == "struct"
	case 25:
		return name == "if"
	case 27:
		return name == "continue"
	case 30:
		return name == "var"
	case 57:
		return name == "interface"
	case 60:
		return name == "type"
	case 78:
		return name == "func"
	case 82:
		return name == "break"
	case 95:
		return name == "go"
	case 126:
		return name == "for"
	case 133:
		return name == "else"
	case 152:
		return name == "package"
	case 154:
		return name == "return"
	case 155:
		return name == "select"
	case 159:
		return name == "goto"
	case 163:
		return name == "chan"
	case 172:
		return name == "default"
	case 187:
		return name == "const"
	case 193:
		return name == "import"
	case 218:
		return name == "range"
	case 235:
		return name == "case"
	case 236:
		return name == "defer"
	}
	return false
}

func TwoHash_XorXor_Shift505(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) ^ (name[0] << 0)) ^ (name[1] << 5) {
	case 7:
		return name == "goto"
	case 9:
		return name == "import"
	case 18:
		return name == "return"
	case 19:
		return name == "select"
	case 35:
		return name == "const"
	case 36:
		return name == "default"
	case 38:
		return name == "fallthrough"
	case 45:
		return name == "map"
	case 51:
		return name == "struct"
	case 54:
		return name == "var"
	case 70:
		return name == "func"
	case 83:
		return name == "switch"
	case 100:
		return name == "defer"
	case 101:
		return name == "else"
	case 130:
		return name == "break"
	case 131:
		return name == "continue"
	case 137:
		return name == "interface"
	case 176:
		return name == "package"
	case 195:
		return name == "case"
	case 199:
		return name == "go"
	case 212:
		return name == "type"
	case 227:
		return name == "chan"
	case 230:
		return name == "for"
	case 233:
		return name == "if"
	case 242:
		return name == "range"
	}
	return false
}

func TwoHash_XorXor_Shift510(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) ^ (name[0] << 1)) ^ (name[1] << 0) {
	case 9:
		return name == "const"
	case 13:
		return name == "defer"
	case 17:
		return name == "type"
	case 22:
		return name == "break"
	case 33:
		return name == "goto"
	case 37:
		return name == "range"
	case 38:
		return name == "else"
	case 39:
		return name == "case"
	case 46:
		return name == "chan"
	case 57:
		return name == "func"
	case 65:
		return name == "return"
	case 67:
		return name == "select"
	case 77:
		return name == "default"
	case 81:
		return name == "switch"
	case 82:
		return name == "struct"
	case 97:
		return name == "package"
	case 127:
		return name == "import"
	case 156:
		return name == "interface"
	case 169:
		return name == "continue"
	case 195:
		return name == "for"
	case 205:
		return name == "fallthrough"
	case 219:
		return name == "map"
	case 225:
		return name == "go"
	case 237:
		return name == "var"
	case 244:
		return name == "if"
	}
	return false
}

func TwoHash_XorXor_Shift513(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) ^ (name[0] << 1)) ^ (name[1] << 3) {
	case 0:
		return name == "default"
	case 6:
		return name == "chan"
	case 8:
		return name == "package"
	case 12:
		return name == "return"
	case 14:
		return name == "select"
	case 30:
		return name == "const"
	case 42:
		return name == "else"
	case 54:
		return name == "goto"
	case 64:
		return name == "defer"
	case 76:
		return name == "range"
	case 78:
		return name == "case"
	case 122:
		return name == "import"
	case 130:
		return name == "interface"
	case 132:
		return name == "var"
	case 134:
		return name == "struct"
	case 158:
		return name == "switch"
	case 160:
		return name == "type"
	case 162:
		return name == "if"
	case 164:
		return name == "fallthrough"
	case 178:
		return name == "map"
	case 190:
		return name == "continue"
	case 212:
		return name == "for"
	case 228:
		return name == "func"
	case 244:
		return name == "break"
	case 246:
		return name == "go"
	}
	return false
}

func TwoHash_XorXor_Shift525(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) ^ (name[0] << 2)) ^ (name[1] << 5) {
	case 0:
		return name == "package"
	case 12:
		return name == "chan"
	case 24:
		return name == "for"
	case 36:
		return name == "if"
	case 44:
		return name == "case"
	case 60:
		return name == "go"
	case 68:
		return name == "interface"
	case 72:
		return name == "range"
	case 104:
		return name == "break"
	case 108:
		return name == "continue"
	case 112:
		return name == "type"
	case 140:
		return name == "struct"
	case 144:
		return name == "defer"
	case 148:
		return name == "else"
	case 152:
		return name == "var"
	case 168:
		return name == "return"
	case 172:
		return name == "select"
	case 184:
		return name == "func"
	case 196:
		return name == "import"
	case 204:
		return name == "const"
	case 208:
		return name == "default"
	case 216:
		return name == "fallthrough"
	case 236:
		return name == "switch"
	case 244:
		return name == "map"
	case 252:
		return name == "goto"
	}
	return false
}

func TwoHash_XorXor_Shift602(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 0)) ^ (name[1] << 2) {
	case 26:
		return name == "for"
	case 34:
		return name == "fallthrough"
	case 35:
		return name == "struct"
	case 41:
		return name == "map"
	case 47:
		return name == "switch"
	case 48:
		return name == "default"
	case 50:
		return name == "var"
	case 52:
		return name == "package"
	case 91:
		return name == "go"
	case 93:
		return name == "import"
	case 102:
		return name == "return"
	case 103:
		return name == "select"
	case 113:
		return name == "if"
	case 144:
		return name == "type"
	case 145:
		return name == "interface"
	case 159:
		return name == "const"
	case 176:
		return name == "defer"
	case 178:
		return name == "func"
	case 182:
		return name == "range"
	case 195:
		return name == "chan"
	case 213:
		return name == "else"
	case 219:
		return name == "goto"
	case 223:
		return name == "continue"
	case 231:
		return name == "case"
	case 234:
		return name == "break"
	}
	return false
}

func TwoHash_XorXor_Shift603(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 0)) ^ (name[1] << 3) {
	case 5:
		return name == "else"
	case 12:
		return name == "defer"
	case 27:
		return name == "continue"
	case 31:
		return name == "goto"
	case 35:
		return name == "chan"
	case 58:
		return name == "range"
	case 75:
		return name == "switch"
	case 83:
		return name == "struct"
	case 89:
		return name == "interface"
	case 91:
		return name == "const"
	case 107:
		return name == "case"
	case 129:
		return name == "import"
	case 140:
		return name == "default"
	case 159:
		return name == "go"
	case 165:
		return name == "map"
	case 174:
		return name == "fallthrough"
	case 178:
		return name == "break"
	case 184:
		return name == "package"
	case 188:
		return name == "type"
	case 190:
		return name == "var"
	case 206:
		return name == "func"
	case 217:
		return name == "if"
	case 218:
		return name == "return"
	case 219:
		return name == "select"
	case 222:
		return name == "for"
	}
	return false
}

func TwoHash_XorXor_Shift604(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 0)) ^ (name[1] << 4) {
	case 2:
		return name == "break"
	case 23:
		return name == "go"
	case 34:
		return name == "range"
	case 54:
		return name == "func"
	case 57:
		return name == "import"
	case 86:
		return name == "for"
	case 115:
		return name == "case"
	case 116:
		return name == "defer"
	case 131:
		return name == "switch"
	case 137:
		return name == "if"
	case 147:
		return name == "continue"
	case 151:
		return name == "goto"
	case 160:
		return name == "package"
	case 162:
		return name == "return"
	case 163:
		return name == "select"
	case 165:
		return name == "else"
	case 166:
		return name == "var"
	case 179:
		return name == "struct"
	case 182:
		return name == "fallthrough"
	case 189:
		return name == "map"
	case 201:
		return name == "interface"
	case 211:
		return name == "const"
	case 227:
		return name == "chan"
	case 228:
		return name == "type"
	case 244:
		return name == "default"
	}
	return false
}

func TwoHash_XorXor_Shift605(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 0)) ^ (name[1] << 5) {
	case 4:
		return name == "default"
	case 7:
		return name == "go"
	case 18:
		return name == "range"
	case 19:
		return name == "switch"
	case 41:
		return name == "if"
	case 67:
		return name == "case"
	case 70:
		return name == "for"
	case 73:
		return name == "import"
	case 82:
		return name == "return"
	case 83:
		return name == "select"
	case 84:
		return name == "type"
	case 98:
		return name == "break"
	case 99:
		return name == "chan"
	case 115:
		return name == "struct"
	case 131:
		return name == "continue"
	case 132:
		return name == "defer"
	case 134:
		return name == "fallthrough"
	case 135:
		return name == "goto"
	case 141:
		return name == "map"
	case 144:
		return name == "package"
	case 150:
		return name == "var"
	case 195:
		return name == "const"
	case 198:
		return name == "func"
	case 229:
		return name == "else"
	case 233:
		return name == "interface"
	}
	return false
}

func TwoHash_XorXor_Shift630(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 3)) ^ (name[1] << 0) {
	case 5:
		return name == "defer"
	case 17:
		return name == "var"
	case 33:
		return name == "package"
	case 34:
		return name == "break"
	case 55:
		return name == "const"
	case 68:
		return name == "else"
	case 69:
		return name == "func"
	case 87:
		return name == "goto"
	case 102:
		return name == "interface"
	case 108:
		return name == "struct"
	case 111:
		return name == "switch"
	case 112:
		return name == "chan"
	case 117:
		return name == "return"
	case 119:
		return name == "continue"
	case 121:
		return name == "case"
	case 125:
		return name == "select"
	case 133:
		return name == "default"
	case 145:
		return name == "fallthrough"
	case 159:
		return name == "for"
	case 165:
		return name == "import"
	case 174:
		return name == "if"
	case 177:
		return name == "range"
	case 201:
		return name == "map"
	case 215:
		return name == "go"
	case 217:
		return name == "type"
	}
	return false
}

func TwoHash_XorAdd_Shift016(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 1)) + (name[1] << 6) {
	case 2:
		return name == "case"
	case 7:
		return name == "fallthrough"
	case 8:
		return name == "func"
	case 13:
		return name == "defer"
	case 15:
		return name == "default"
	case 20:
		return name == "import"
	case 25:
		return name == "map"
	case 32:
		return name == "select"
	case 33:
		return name == "range"
	case 34:
		return name == "return"
	case 39:
		return name == "package"
	case 44:
		return name == "type"
	case 47:
		return name == "var"
	case 65:
		return name == "break"
	case 80:
		return name == "if"
	case 91:
		return name == "interface"
	case 131:
		return name == "const"
	case 138:
		return name == "goto"
	case 140:
		return name == "go"
	case 142:
		return name == "continue"
	case 143:
		return name == "for"
	case 160:
		return name == "switch"
	case 194:
		return name == "chan"
	case 206:
		return name == "else"
	case 224:
		return name == "struct"
	}
	return false
}

func TwoHash_XorAdd_Shift024(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 2)) + (name[1] << 4) {
	case 6:
		return name == "if"
	case 8:
		return name == "chan"
	case 10:
		return name == "struct"
	case 26:
		return name == "select"
	case 30:
		return name == "return"
	case 58:
		return name == "switch"
	case 80:
		return name == "else"
	case 100:
		return name == "type"
	case 114:
		return name == "import"
	case 116:
		return name == "continue"
	case 121:
		return name == "const"
	case 136:
		return name == "goto"
	case 139:
		return name == "for"
	case 141:
		return name == "interface"
	case 142:
		return name == "go"
	case 152:
		return name == "case"
	case 163:
		return name == "fallthrough"
	case 173:
		return name == "break"
	case 199:
		return name == "map"
	case 215:
		return name == "package"
	case 221:
		return name == "range"
	case 229:
		return name == "defer"
	case 231:
		return name == "default"
	case 235:
		return name == "var"
	case 236:
		return name == "func"
	}
	return false
}

func TwoHash_XorAdd_Shift025(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 2)) + (name[1] << 5) {
	case 16:
		return name == "else"
	case 53:
		return name == "defer"
	case 55:
		return name == "default"
	case 60:
		return name == "func"
	case 66:
		return name == "import"
	case 74:
		return name == "struct"
	case 100:
		return name == "continue"
	case 102:
		return name == "if"
	case 105:
		return name == "const"
	case 106:
		return name == "select"
	case 109:
		return name == "interface"
	case 110:
		return name == "return"
	case 120:
		return name == "goto"
	case 123:
		return name == "for"
	case 126:
		return name == "go"
	case 136:
		return name == "chan"
	case 168:
		return name == "case"
	case 170:
		return name == "switch"
	case 179:
		return name == "fallthrough"
	case 205:
		return name == "break"
	case 215:
		return name == "map"
	case 231:
		return name == "package"
	case 237:
		return name == "range"
	case 244:
		return name == "type"
	case 251:
		return name == "var"
	}
	return false
}

func TwoHash_XorAdd_Shift031(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 3)) + (name[1] << 1) {
	case 4:
		return name == "else"
	case 17:
		return name == "for"
	case 22:
		return name == "if"
	case 24:
		return name == "go"
	case 26:
		return name == "goto"
	case 29:
		return name == "interface"
	case 30:
		return name == "func"
	case 40:
		return name == "import"
	case 45:
		return name == "map"
	case 73:
		return name == "package"
	case 87:
		return name == "range"
	case 96:
		return name == "return"
	case 104:
		return name == "select"
	case 117:
		return name == "var"
	case 134:
		return name == "struct"
	case 140:
		return name == "switch"
	case 150:
		return name == "type"
	case 222:
		return name == "case"
	case 236:
		return name == "chan"
	case 238:
		return name == "continue"
	case 239:
		return name == "defer"
	case 241:
		return name == "default"
	case 249:
		return name == "break"
	case 251:
		return name == "const"
	case 253:
		return name == "fallthrough"
	}
	return false
}

func TwoHash_XorAdd_Shift033(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 3)) + (name[1] << 3) {
	case 36:
		return name == "case"
	case 62:
		return name == "struct"
	case 67:
		return name == "fallthrough"
	case 77:
		return name == "defer"
	case 79:
		return name == "default"
	case 86:
		return name == "switch"
	case 92:
		return name == "chan"
	case 108:
		return name == "type"
	case 115:
		return name == "map"
	case 122:
		return name == "if"
	case 136:
		return name == "continue"
	case 140:
		return name == "else"
	case 143:
		return name == "package"
	case 149:
		return name == "const"
	case 157:
		return name == "range"
	case 165:
		return name == "break"
	case 171:
		return name == "for"
	case 177:
		return name == "interface"
	case 178:
		return name == "go"
	case 180:
		return name == "goto"
	case 182:
		return name == "import"
	case 187:
		return name == "var"
	case 190:
		return name == "return"
	case 198:
		return name == "select"
	case 220:
		return name == "func"
	}
	return false
}

func TwoHash_XorAdd_Shift041(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 4)) + (name[1] << 1) {
	case 0:
		return name == "select"
	case 4:
		return name == "chan"
	case 9:
		return name == "break"
	case 15:
		return name == "defer"
	case 17:
		return name == "default"
	case 19:
		return name == "const"
	case 22:
		return name == "continue"
	case 30:
		return name == "struct"
	case 36:
		return name == "switch"
	case 37:
		return name == "var"
	case 44:
		return name == "else"
	case 45:
		return name == "fallthrough"
	case 54:
		return name == "type"
	case 65:
		return name == "for"
	case 78:
		return name == "func"
	case 80:
		return name == "go"
	case 82:
		return name == "goto"
	case 94:
		return name == "if"
	case 112:
		return name == "import"
	case 117:
		return name == "interface"
	case 149:
		return name == "map"
	case 201:
		return name == "package"
	case 231:
		return name == "range"
	case 240:
		return name == "return"
	case 246:
		return name == "case"
	}
	return false
}

func TwoHash_XorAdd_Shift042(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 4)) + (name[1] << 2) {
	case 4:
		return name == "else"
	case 6:
		return name == "struct"
	case 18:
		return name == "switch"
	case 31:
		return name == "for"
	case 40:
		return name == "type"
	case 42:
		return name == "if"
	case 46:
		return name == "go"
	case 48:
		return name == "goto"
	case 56:
		return name == "func"
	case 74:
		return name == "import"
	case 81:
		return name == "interface"
	case 87:
		return name == "map"
	case 139:
		return name == "package"
	case 169:
		return name == "range"
	case 184:
		return name == "case"
	case 186:
		return name == "return"
	case 202:
		return name == "select"
	case 212:
		return name == "chan"
	case 217:
		return name == "defer"
	case 219:
		return name == "default"
	case 231:
		return name == "var"
	case 237:
		return name == "break"
	case 239:
		return name == "fallthrough"
	case 241:
		return name == "const"
	case 244:
		return name == "continue"
	}
	return false
}

func TwoHash_XorAdd_Shift050(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 5)) + (name[1] << 0) {
	case 4:
		return name == "map"
	case 16:
		return name == "else"
	case 36:
		return name == "var"
	case 44:
		return name == "fallthrough"
	case 50:
		return name == "for"
	case 57:
		return name == "func"
	case 81:
		return name == "go"
	case 83:
		return name == "goto"
	case 104:
		return name == "package"
	case 136:
		return name == "if"
	case 147:
		return name == "import"
	case 151:
		return name == "interface"
	case 166:
		return name == "range"
	case 171:
		return name == "return"
	case 183:
		return name == "break"
	case 197:
		return name == "case"
	case 203:
		return name == "select"
	case 204:
		return name == "chan"
	case 212:
		return name == "const"
	case 215:
		return name == "continue"
	case 218:
		return name == "struct"
	case 221:
		return name == "switch"
	case 234:
		return name == "defer"
	case 236:
		return name == "default"
	case 253:
		return name == "type"
	}
	return false
}

func TwoHash_XorAdd_Shift051(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 5)) + (name[1] << 1) {
	case 0:
		return name == "import"
	case 5:
		return name == "interface"
	case 7:
		return name == "range"
	case 16:
		return name == "return"
	case 38:
		return name == "case"
	case 41:
		return name == "break"
	case 48:
		return name == "select"
	case 52:
		return name == "chan"
	case 67:
		return name == "const"
	case 70:
		return name == "continue"
	case 78:
		return name == "struct"
	case 79:
		return name == "defer"
	case 81:
		return name == "default"
	case 84:
		return name == "switch"
	case 101:
		return name == "map"
	case 118:
		return name == "type"
	case 124:
		return name == "else"
	case 133:
		return name == "var"
	case 141:
		return name == "fallthrough"
	case 161:
		return name == "for"
	case 174:
		return name == "func"
	case 192:
		return name == "go"
	case 194:
		return name == "goto"
	case 201:
		return name == "package"
	case 238:
		return name == "if"
	}
	return false
}

func TwoHash_XorAdd_Shift060(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 6)) + (name[1] << 0) {
	case 37:
		return name == "case"
	case 43:
		return name == "select"
	case 44:
		return name == "chan"
	case 49:
		return name == "go"
	case 51:
		return name == "goto"
	case 52:
		return name == "const"
	case 55:
		return name == "continue"
	case 58:
		return name == "struct"
	case 61:
		return name == "switch"
	case 104:
		return name == "package"
	case 106:
		return name == "defer"
	case 108:
		return name == "default"
	case 125:
		return name == "type"
	case 164:
		return name == "map"
	case 168:
		return name == "if"
	case 176:
		return name == "else"
	case 179:
		return name == "import"
	case 183:
		return name == "interface"
	case 228:
		return name == "var"
	case 230:
		return name == "range"
	case 235:
		return name == "return"
	case 236:
		return name == "fallthrough"
	case 242:
		return name == "for"
	case 247:
		return name == "break"
	case 249:
		return name == "func"
	}
	return false
}

func TwoHash_XorAdd_Shift061(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 6)) + (name[1] << 1) {
	case 5:
		return name == "map"
	case 14:
		return name == "if"
	case 28:
		return name == "else"
	case 32:
		return name == "import"
	case 37:
		return name == "interface"
	case 69:
		return name == "var"
	case 71:
		return name == "range"
	case 77:
		return name == "fallthrough"
	case 80:
		return name == "return"
	case 97:
		return name == "for"
	case 105:
		return name == "break"
	case 110:
		return name == "func"
	case 134:
		return name == "case"
	case 144:
		return name == "select"
	case 148:
		return name == "chan"
	case 160:
		return name == "go"
	case 162:
		return name == "goto"
	case 163:
		return name == "const"
	case 166:
		return name == "continue"
	case 174:
		return name == "struct"
	case 180:
		return name == "switch"
	case 201:
		return name == "package"
	case 207:
		return name == "defer"
	case 209:
		return name == "default"
	case 246:
		return name == "type"
	}
	return false
}

func TwoHash_XorAdd_Shift062(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 6)) + (name[1] << 2) {
	case 1:
		return name == "interface"
	case 7:
		return name == "var"
	case 9:
		return name == "range"
	case 15:
		return name == "fallthrough"
	case 26:
		return name == "return"
	case 63:
		return name == "for"
	case 72:
		return name == "case"
	case 77:
		return name == "break"
	case 88:
		return name == "func"
	case 90:
		return name == "select"
	case 100:
		return name == "chan"
	case 126:
		return name == "go"
	case 128:
		return name == "goto"
	case 129:
		return name == "const"
	case 132:
		return name == "continue"
	case 139:
		return name == "package"
	case 150:
		return name == "struct"
	case 153:
		return name == "defer"
	case 155:
		return name == "default"
	case 162:
		return name == "switch"
	case 199:
		return name == "map"
	case 218:
		return name == "if"
	case 232:
		return name == "type"
	case 244:
		return name == "else"
	case 250:
		return name == "import"
	}
	return false
}

func TwoHash_XorAdd_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) ^ (name[0] << 7)) + (name[1] << 3) {
	case 0:
		return name == "continue"
	case 11:
		return name == "var"
	case 13:
		return name == "range"
	case 15:
		return name == "package"
	case 19:
		return name == "fallthrough"
	case 38:
		return name == "struct"
	case 45:
		return name == "defer"
	case 46:
		return name == "return"
	case 47:
		return name == "default"
	case 62:
		return name == "switch"
	case 123:
		return name == "for"
	case 139:
		return name == "map"
	case 140:
		return name == "case"
	case 149:
		return name == "break"
	case 172:
		return name == "func"
	case 174:
		return name == "select"
	case 178:
		return name == "if"
	case 196:
		return name == "chan"
	case 204:
		return name == "type"
	case 228:
		return name == "else"
	case 238:
		return name == "import"
	case 249:
		return name == "interface"
	case 250:
		return name == "go"
	case 252:
		return name == "goto"
	case 253:
		return name == "const"
	}
	return false
}

func TwoHash_XorAdd_Shift150(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) ^ (name[0] << 5)) + (name[1] << 0) {
	case 1:
		return name == "type"
	case 7:
		return name == "map"
	case 20:
		return name == "else"
	case 39:
		return name == "var"
	case 53:
		return name == "for"
	case 55:
		return name == "fallthrough"
	case 61:
		return name == "func"
	case 83:
		return name == "go"
	case 87:
		return name == "goto"
	case 111:
		return name == "package"
	case 138:
		return name == "if"
	case 153:
		return name == "import"
	case 160:
		return name == "interface"
	case 171:
		return name == "range"
	case 177:
		return name == "return"
	case 188:
		return name == "break"
	case 201:
		return name == "case"
	case 208:
		return name == "chan"
	case 209:
		return name == "select"
	case 217:
		return name == "const"
	case 223:
		return name == "continue"
	case 224:
		return name == "struct"
	case 227:
		return name == "switch"
	case 239:
		return name == "defer"
	case 243:
		return name == "default"
	}
	return false
}

func TwoHash_XorAdd_Shift161(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) ^ (name[0] << 6)) + (name[1] << 1) {
	case 8:
		return name == "map"
	case 16:
		return name == "if"
	case 32:
		return name == "else"
	case 38:
		return name == "import"
	case 46:
		return name == "interface"
	case 72:
		return name == "var"
	case 76:
		return name == "range"
	case 86:
		return name == "return"
	case 88:
		return name == "fallthrough"
	case 100:
		return name == "for"
	case 110:
		return name == "break"
	case 114:
		return name == "func"
	case 138:
		return name == "case"
	case 150:
		return name == "select"
	case 152:
		return name == "chan"
	case 162:
		return name == "go"
	case 166:
		return name == "goto"
	case 168:
		return name == "const"
	case 174:
		return name == "continue"
	case 180:
		return name == "struct"
	case 186:
		return name == "switch"
	case 208:
		return name == "package"
	case 212:
		return name == "defer"
	case 216:
		return name == "default"
	case 250:
		return name == "type"
	}
	return false
}

func TwoHash_XorAdd_Shift240(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) ^ (name[0] << 4)) + (name[1] << 0) {
	case 34:
		return name == "interface"
	case 61:
		return name == "map"
	case 125:
		return name == "package"
	case 127:
		return name == "continue"
	case 129:
		return name == "case"
	case 136:
		return name == "chan"
	case 141:
		return name == "select"
	case 147:
		return name == "const"
	case 149:
		return name == "range"
	case 156:
		return name == "struct"
	case 157:
		return name == "return"
	case 159:
		return name == "switch"
	case 166:
		return name == "break"
	case 172:
		return name == "else"
	case 173:
		return name == "fallthrough"
	case 185:
		return name == "defer"
	case 193:
		return name == "default"
	case 201:
		return name == "type"
	case 205:
		return name == "var"
	case 207:
		return name == "goto"
	case 219:
		return name == "for"
	case 229:
		return name == "func"
	case 231:
		return name == "go"
	case 245:
		return name == "import"
	case 254:
		return name == "if"
	}
	return false
}

func TwoHash_XorAdd_Shift250(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) ^ (name[0] << 5)) + (name[1] << 0) {
	case 1:
		return name == "default"
	case 9:
		return name == "type"
	case 13:
		return name == "map"
	case 28:
		return name == "else"
	case 45:
		return name == "var"
	case 59:
		return name == "for"
	case 69:
		return name == "func"
	case 77:
		return name == "fallthrough"
	case 87:
		return name == "go"
	case 95:
		return name == "goto"
	case 114:
		return name == "interface"
	case 125:
		return name == "package"
	case 142:
		return name == "if"
	case 165:
		return name == "import"
	case 175:
		return name == "continue"
	case 181:
		return name == "range"
	case 189:
		return name == "return"
	case 198:
		return name == "break"
	case 209:
		return name == "case"
	case 216:
		return name == "chan"
	case 221:
		return name == "select"
	case 227:
		return name == "const"
	case 236:
		return name == "struct"
	case 239:
		return name == "switch"
	case 249:
		return name == "defer"
	}
	return false
}

func TwoHash_XorAdd_Shift302(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) ^ (name[0] << 0)) + (name[1] << 2) {
	case 3:
		return name == "goto"
	case 7:
		return name == "const"
	case 13:
		return name == "import"
	case 17:
		return name == "if"
	case 18:
		return name == "break"
	case 19:
		return name == "struct"
	case 26:
		return name == "func"
	case 31:
		return name == "switch"
	case 51:
		return name == "go"
	case 56:
		return name == "type"
	case 58:
		return name == "for"
	case 194:
		return name == "fallthrough"
	case 199:
		return name == "case"
	case 204:
		return name == "package"
	case 214:
		return name == "return"
	case 215:
		return name == "select"
	case 217:
		return name == "interface"
	case 222:
		return name == "range"
	case 223:
		return name == "continue"
	case 224:
		return name == "defer"
	case 227:
		return name == "chan"
	case 240:
		return name == "default"
	case 242:
		return name == "var"
	case 245:
		return name == "else"
	case 249:
		return name == "map"
	}
	return false
}

func TwoHash_XorAdd_Shift303(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) ^ (name[0] << 0)) + (name[1] << 3) {
	case 28:
		return name == "type"
	case 70:
		return name == "fallthrough"
	case 75:
		return name == "case"
	case 80:
		return name == "package"
	case 98:
		return name == "range"
	case 106:
		return name == "return"
	case 107:
		return name == "select"
	case 116:
		return name == "defer"
	case 118:
		return name == "var"
	case 125:
		return name == "map"
	case 131:
		return name == "chan"
	case 132:
		return name == "default"
	case 145:
		return name == "interface"
	case 155:
		return name == "continue"
	case 165:
		return name == "else"
	case 169:
		return name == "if"
	case 191:
		return name == "goto"
	case 193:
		return name == "import"
	case 195:
		return name == "const"
	case 218:
		return name == "break"
	case 227:
		return name == "struct"
	case 238:
		return name == "func"
	case 239:
		return name == "go"
	case 246:
		return name == "for"
	case 251:
		return name == "switch"
	}
	return false
}

func TwoHash_XorAdd_Shift351(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) ^ (name[0] << 5)) + (name[1] << 1) {
	case 2:
		return name == "case"
	case 16:
		return name == "chan"
	case 26:
		return name == "select"
	case 38:
		return name == "const"
	case 42:
		return name == "range"
	case 56:
		return name == "struct"
	case 58:
		return name == "return"
	case 62:
		return name == "switch"
	case 68:
		return name == "interface"
	case 76:
		return name == "break"
	case 88:
		return name == "else"
	case 90:
		return name == "fallthrough"
	case 114:
		return name == "defer"
	case 122:
		return name == "map"
	case 130:
		return name == "default"
	case 146:
		return name == "type"
	case 154:
		return name == "var"
	case 158:
		return name == "goto"
	case 182:
		return name == "for"
	case 202:
		return name == "func"
	case 206:
		return name == "go"
	case 234:
		return name == "import"
	case 250:
		return name == "package"
	case 252:
		return name == "if"
	case 254:
		return name == "continue"
	}
	return false
}

func TwoHash_XorAdd_Shift361(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) ^ (name[0] << 6)) + (name[1] << 1) {
	case 2:
		return name == "default"
	case 18:
		return name == "type"
	case 26:
		return name == "map"
	case 28:
		return name == "if"
	case 56:
		return name == "else"
	case 74:
		return name == "import"
	case 90:
		return name == "var"
	case 94:
		return name == "continue"
	case 106:
		return name == "range"
	case 118:
		return name == "for"
	case 122:
		return name == "return"
	case 138:
		return name == "func"
	case 140:
		return name == "break"
	case 154:
		return name == "fallthrough"
	case 162:
		return name == "case"
	case 174:
		return name == "go"
	case 176:
		return name == "chan"
	case 186:
		return name == "select"
	case 190:
		return name == "goto"
	case 198:
		return name == "const"
	case 216:
		return name == "struct"
	case 222:
		return name == "switch"
	case 228:
		return name == "interface"
	case 242:
		return name == "defer"
	case 250:
		return name == "package"
	}
	return false
}

func TwoHash_XorAdd_Shift400(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) ^ (name[0] << 0)) + (name[1] << 0) {
	case 55:
		return name == "fallthrough"
	case 82:
		return name == "continue"
	case 97:
		return name == "package"
	case 103:
		return name == "interface"
	case 118:
		return name == "import"
	case 119:
		return name == "return"
	case 120:
		return name == "select"
	case 121:
		return name == "default"
	case 131:
		return name == "range"
	case 132:
		return name == "case"
	case 135:
		return name == "struct"
	case 138:
		return name == "switch"
	case 139:
		return name == "chan"
	case 145:
		return name == "else"
	case 150:
		return name == "goto"
	case 153:
		return name == "defer"
	case 155:
		return name == "func"
	case 162:
		return name == "const"
	case 164:
		return name == "break"
	case 167:
		return name == "var"
	case 173:
		return name == "type"
	case 175:
		return name == "if"
	case 182:
		return name == "go"
	case 190:
		return name == "map"
	case 197:
		return name == "for"
	}
	return false
}

func TwoHash_XorAdd_Shift401(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) ^ (name[0] << 0)) + (name[1] << 1) {
	case 1:
		return name == "switch"
	case 5:
		return name == "goto"
	case 8:
		return name == "var"
	case 16:
		return name == "func"
	case 17:
		return name == "const"
	case 21:
		return name == "if"
	case 22:
		return name == "break"
	case 31:
		return name == "map"
	case 37:
		return name == "go"
	case 38:
		return name == "type"
	case 52:
		return name == "for"
	case 152:
		return name == "fallthrough"
	case 193:
		return name == "continue"
	case 194:
		return name == "package"
	case 213:
		return name == "interface"
	case 220:
		return name == "return"
	case 221:
		return name == "select"
	case 222:
		return name == "default"
	case 227:
		return name == "import"
	case 228:
		return name == "range"
	case 229:
		return name == "case"
	case 243:
		return name == "chan"
	case 251:
		return name == "struct"
	case 253:
		return name == "else"
	case 254:
		return name == "defer"
	}
	return false
}

func TwoHash_XorAdd_Shift405(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) ^ (name[0] << 0)) + (name[1] << 5) {
	case 7:
		return name == "goto"
	case 9:
		return name == "if"
	case 19:
		return name == "const"
	case 32:
		return name == "package"
	case 35:
		return name == "chan"
	case 39:
		return name == "go"
	case 54:
		return name == "for"
	case 66:
		return name == "range"
	case 67:
		return name == "case"
	case 84:
		return name == "type"
	case 102:
		return name == "var"
	case 114:
		return name == "break"
	case 125:
		return name == "map"
	case 147:
		return name == "struct"
	case 165:
		return name == "else"
	case 169:
		return name == "import"
	case 178:
		return name == "return"
	case 179:
		return name == "select"
	case 180:
		return name == "default"
	case 185:
		return name == "interface"
	case 195:
		return name == "continue"
	case 198:
		return name == "func"
	case 212:
		return name == "defer"
	case 243:
		return name == "switch"
	case 246:
		return name == "fallthrough"
	}
	return false
}

func TwoHash_XorAdd_Shift413(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) ^ (name[0] << 1)) + (name[1] << 3) {
	case 6:
		return name == "goto"
	case 14:
		return name == "const"
	case 26:
		return name == "import"
	case 34:
		return name == "if"
	case 36:
		return name == "break"
	case 38:
		return name == "struct"
	case 52:
		return name == "func"
	case 62:
		return name == "switch"
	case 102:
		return name == "go"
	case 112:
		return name == "type"
	case 116:
		return name == "for"
	case 132:
		return name == "fallthrough"
	case 142:
		return name == "case"
	case 152:
		return name == "package"
	case 172:
		return name == "return"
	case 174:
		return name == "select"
	case 178:
		return name == "interface"
	case 188:
		return name == "range"
	case 190:
		return name == "continue"
	case 192:
		return name == "defer"
	case 198:
		return name == "chan"
	case 224:
		return name == "default"
	case 228:
		return name == "var"
	case 234:
		return name == "else"
	case 242:
		return name == "map"
	}
	return false
}

func TwoHash_XorAdd_Shift500(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) ^ (name[0] << 0)) + (name[1] << 0) {
	case 22:
		return name == "import"
	case 23:
		return name == "return"
	case 24:
		return name == "select"
	case 39:
		return name == "struct"
	case 41:
		return name == "defer"
	case 42:
		return name == "switch"
	case 50:
		return name == "const"
	case 51:
		return name == "range"
	case 52:
		return name == "break"
	case 68:
		return name == "case"
	case 75:
		return name == "chan"
	case 81:
		return name == "else"
	case 86:
		return name == "goto"
	case 91:
		return name == "func"
	case 103:
		return name == "fallthrough"
	case 109:
		return name == "type"
	case 110:
		return name == "map"
	case 117:
		return name == "for"
	case 119:
		return name == "var"
	case 143:
		return name == "if"
	case 150:
		return name == "go"
	case 183:
		return name == "interface"
	case 210:
		return name == "continue"
	case 233:
		return name == "default"
	case 241:
		return name == "package"
	}
	return false
}

func TwoHash_XorAdd_Shift505(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) ^ (name[0] << 0)) + (name[1] << 5) {
	case 2:
		return name == "break"
	case 3:
		return name == "case"
	case 7:
		return name == "go"
	case 9:
		return name == "interface"
	case 20:
		return name == "type"
	case 36:
		return name == "default"
	case 38:
		return name == "fallthrough"
	case 45:
		return name == "map"
	case 51:
		return name == "struct"
	case 54:
		return name == "var"
	case 67:
		return name == "continue"
	case 73:
		return name == "import"
	case 82:
		return name == "return"
	case 83:
		return name == "select"
	case 100:
		return name == "defer"
	case 101:
		return name == "else"
	case 134:
		return name == "func"
	case 147:
		return name == "switch"
	case 163:
		return name == "const"
	case 176:
		return name == "package"
	case 199:
		return name == "goto"
	case 227:
		return name == "chan"
	case 230:
		return name == "for"
	case 233:
		return name == "if"
	case 242:
		return name == "range"
	}
	return false
}

func TwoHash_XorAdd_Shift511(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) ^ (name[0] << 1)) + (name[1] << 1) {
	case 6:
		return name == "range"
	case 8:
		return name == "case"
	case 14:
		return name == "struct"
	case 20:
		return name == "switch"
	case 22:
		return name == "chan"
	case 34:
		return name == "else"
	case 44:
		return name == "goto"
	case 50:
		return name == "defer"
	case 54:
		return name == "func"
	case 68:
		return name == "const"
	case 72:
		return name == "break"
	case 78:
		return name == "var"
	case 90:
		return name == "type"
	case 94:
		return name == "if"
	case 108:
		return name == "go"
	case 110:
		return name == "fallthrough"
	case 124:
		return name == "map"
	case 138:
		return name == "for"
	case 164:
		return name == "continue"
	case 194:
		return name == "package"
	case 206:
		return name == "interface"
	case 236:
		return name == "import"
	case 238:
		return name == "return"
	case 240:
		return name == "select"
	case 242:
		return name == "default"
	}
	return false
}

func TwoHash_XorAdd_Shift512(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) ^ (name[0] << 1)) + (name[1] << 2) {
	case 2:
		return name == "switch"
	case 10:
		return name == "goto"
	case 16:
		return name == "var"
	case 32:
		return name == "func"
	case 34:
		return name == "const"
	case 42:
		return name == "if"
	case 44:
		return name == "break"
	case 48:
		return name == "fallthrough"
	case 62:
		return name == "map"
	case 74:
		return name == "go"
	case 76:
		return name == "type"
	case 104:
		return name == "for"
	case 130:
		return name == "continue"
	case 132:
		return name == "package"
	case 170:
		return name == "interface"
	case 184:
		return name == "return"
	case 186:
		return name == "select"
	case 188:
		return name == "default"
	case 198:
		return name == "import"
	case 200:
		return name == "range"
	case 202:
		return name == "case"
	case 230:
		return name == "chan"
	case 246:
		return name == "struct"
	case 250:
		return name == "else"
	case 252:
		return name == "defer"
	}
	return false
}

func TwoHash_XorAdd_Shift601(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 0)) + (name[1] << 1) {
	case 1:
		return name == "const"
	case 5:
		return name == "interface"
	case 6:
		return name == "break"
	case 37:
		return name == "case"
	case 51:
		return name == "chan"
	case 61:
		return name == "else"
	case 65:
		return name == "continue"
	case 69:
		return name == "goto"
	case 80:
		return name == "func"
	case 102:
		return name == "type"
	case 104:
		return name == "fallthrough"
	case 110:
		return name == "default"
	case 111:
		return name == "map"
	case 114:
		return name == "package"
	case 120:
		return name == "var"
	case 132:
		return name == "for"
	case 181:
		return name == "if"
	case 188:
		return name == "return"
	case 189:
		return name == "select"
	case 195:
		return name == "import"
	case 197:
		return name == "go"
	case 219:
		return name == "struct"
	case 225:
		return name == "switch"
	case 238:
		return name == "defer"
	case 244:
		return name == "range"
	}
	return false
}

func TwoHash_XorAdd_Shift603(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 0)) + (name[1] << 3) {
	case 14:
		return name == "func"
	case 25:
		return name == "if"
	case 26:
		return name == "return"
	case 27:
		return name == "select"
	case 30:
		return name == "for"
	case 58:
		return name == "range"
	case 60:
		return name == "type"
	case 76:
		return name == "defer"
	case 81:
		return name == "import"
	case 95:
		return name == "go"
	case 107:
		return name == "case"
	case 147:
		return name == "struct"
	case 153:
		return name == "interface"
	case 155:
		return name == "const"
	case 163:
		return name == "chan"
	case 171:
		return name == "switch"
	case 174:
		return name == "fallthrough"
	case 178:
		return name == "break"
	case 181:
		return name == "map"
	case 184:
		return name == "package"
	case 190:
		return name == "var"
	case 197:
		return name == "else"
	case 204:
		return name == "default"
	case 219:
		return name == "continue"
	case 223:
		return name == "goto"
	}
	return false
}

func TwoHash_XorAdd_Shift605(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 0)) + (name[1] << 5) {
	case 3:
		return name == "const"
	case 6:
		return name == "func"
	case 67:
		return name == "continue"
	case 68:
		return name == "default"
	case 71:
		return name == "goto"
	case 82:
		return name == "range"
	case 98:
		return name == "break"
	case 99:
		return name == "chan"
	case 115:
		return name == "struct"
	case 131:
		return name == "case"
	case 134:
		return name == "for"
	case 137:
		return name == "import"
	case 146:
		return name == "return"
	case 147:
		return name == "select"
	case 148:
		return name == "type"
	case 169:
		return name == "if"
	case 196:
		return name == "defer"
	case 198:
		return name == "fallthrough"
	case 199:
		return name == "go"
	case 205:
		return name == "map"
	case 208:
		return name == "package"
	case 211:
		return name == "switch"
	case 214:
		return name == "var"
	case 229:
		return name == "else"
	case 233:
		return name == "interface"
	}
	return false
}

func TwoHash_XorAdd_Shift606(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 0)) + (name[1] << 6) {
	case 35:
		return name == "continue"
	case 39:
		return name == "goto"
	case 41:
		return name == "import"
	case 50:
		return name == "return"
	case 51:
		return name == "select"
	case 99:
		return name == "chan"
	case 100:
		return name == "defer"
	case 101:
		return name == "else"
	case 102:
		return name == "for"
	case 105:
		return name == "if"
	case 114:
		return name == "range"
	case 162:
		return name == "break"
	case 163:
		return name == "case"
	case 166:
		return name == "func"
	case 167:
		return name == "go"
	case 169:
		return name == "interface"
	case 179:
		return name == "switch"
	case 180:
		return name == "type"
	case 227:
		return name == "const"
	case 228:
		return name == "default"
	case 230:
		return name == "fallthrough"
	case 237:
		return name == "map"
	case 240:
		return name == "package"
	case 243:
		return name == "struct"
	case 246:
		return name == "var"
	}
	return false
}

func TwoHash_XorAdd_Shift616(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 1)) + (name[1] << 6) {
	case 4:
		return name == "break"
	case 6:
		return name == "case"
	case 12:
		return name == "func"
	case 14:
		return name == "go"
	case 18:
		return name == "interface"
	case 38:
		return name == "switch"
	case 40:
		return name == "type"
	case 70:
		return name == "const"
	case 72:
		return name == "default"
	case 76:
		return name == "fallthrough"
	case 90:
		return name == "map"
	case 96:
		return name == "package"
	case 102:
		return name == "struct"
	case 108:
		return name == "var"
	case 134:
		return name == "continue"
	case 142:
		return name == "goto"
	case 146:
		return name == "import"
	case 164:
		return name == "return"
	case 166:
		return name == "select"
	case 198:
		return name == "chan"
	case 200:
		return name == "defer"
	case 202:
		return name == "else"
	case 204:
		return name == "for"
	case 210:
		return name == "if"
	case 228:
		return name == "range"
	}
	return false
}

func TwoHash_XorAdd_Shift620(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 2)) + (name[1] << 0) {
	case 0:
		return name == "else"
	case 11:
		return name == "goto"
	case 13:
		return name == "func"
	case 53:
		return name == "defer"
	case 58:
		return name == "break"
	case 59:
		return name == "const"
	case 73:
		return name == "type"
	case 82:
		return name == "interface"
	case 97:
		return name == "package"
	case 121:
		return name == "var"
	case 138:
		return name == "if"
	case 139:
		return name == "go"
	case 145:
		return name == "import"
	case 173:
		return name == "return"
	case 177:
		return name == "select"
	case 181:
		return name == "default"
	case 185:
		return name == "fallthrough"
	case 192:
		return name == "struct"
	case 195:
		return name == "switch"
	case 199:
		return name == "for"
	case 213:
		return name == "map"
	case 233:
		return name == "range"
	case 237:
		return name == "case"
	case 244:
		return name == "chan"
	case 251:
		return name == "continue"
	}
	return false
}

func TwoHash_XorAdd_Shift630(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) ^ (name[0] << 3)) + (name[1] << 0) {
	case 9:
		return name == "map"
	case 25:
		return name == "type"
	case 39:
		return name == "go"
	case 46:
		return name == "if"
	case 49:
		return name == "range"
	case 53:
		return name == "import"
	case 69:
		return name == "default"
	case 81:
		return name == "fallthrough"
	case 95:
		return name == "for"
	case 117:
		return name == "return"
	case 118:
		return name == "interface"
	case 121:
		return name == "case"
	case 125:
		return name == "select"
	case 128:
		return name == "chan"
	case 135:
		return name == "continue"
	case 140:
		return name == "struct"
	case 143:
		return name == "switch"
	case 148:
		return name == "else"
	case 161:
		return name == "package"
	case 165:
		return name == "func"
	case 167:
		return name == "goto"
	case 194:
		return name == "break"
	case 197:
		return name == "defer"
	case 199:
		return name == "const"
	case 209:
		return name == "var"
	}
	return false
}

func TwoHash_AddXor_Shift015(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 1)) ^ (name[1] << 5) {
	case 12:
		return name == "switch"
	case 20:
		return name == "if"
	case 27:
		return name == "interface"
	case 43:
		return name == "const"
	case 46:
		return name == "continue"
	case 47:
		return name == "for"
	case 48:
		return name == "go"
	case 50:
		return name == "goto"
	case 74:
		return name == "return"
	case 76:
		return name == "select"
	case 78:
		return name == "else"
	case 108:
		return name == "struct"
	case 109:
		return name == "defer"
	case 111:
		return name == "default"
	case 112:
		return name == "func"
	case 120:
		return name == "import"
	case 137:
		return name == "break"
	case 199:
		return name == "package"
	case 201:
		return name == "range"
	case 202:
		return name == "chan"
	case 204:
		return name == "type"
	case 207:
		return name == "var"
	case 234:
		return name == "case"
	case 247:
		return name == "fallthrough"
	case 253:
		return name == "map"
	}
	return false
}

func TwoHash_AddXor_Shift024(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 2)) ^ (name[1] << 4) {
	case 16:
		return name == "chan"
	case 68:
		return name == "type"
	case 77:
		return name == "interface"
	case 80:
		return name == "goto"
	case 88:
		return name == "else"
	case 97:
		return name == "const"
	case 100:
		return name == "continue"
	case 107:
		return name == "for"
	case 110:
		return name == "go"
	case 122:
		return name == "import"
	case 128:
		return name == "case"
	case 130:
		return name == "select"
	case 146:
		return name == "struct"
	case 158:
		return name == "return"
	case 162:
		return name == "switch"
	case 167:
		return name == "map"
	case 173:
		return name == "break"
	case 179:
		return name == "fallthrough"
	case 197:
		return name == "defer"
	case 198:
		return name == "if"
	case 199:
		return name == "default"
	case 203:
		return name == "var"
	case 204:
		return name == "func"
	case 215:
		return name == "package"
	case 221:
		return name == "range"
	}
	return false
}

func TwoHash_AddXor_Shift025(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 2)) ^ (name[1] << 5) {
	case 10:
		return name == "import"
	case 24:
		return name == "else"
	case 50:
		return name == "switch"
	case 53:
		return name == "defer"
	case 55:
		return name == "default"
	case 60:
		return name == "func"
	case 64:
		return name == "goto"
	case 82:
		return name == "struct"
	case 102:
		return name == "if"
	case 109:
		return name == "interface"
	case 110:
		return name == "return"
	case 113:
		return name == "const"
	case 114:
		return name == "select"
	case 116:
		return name == "continue"
	case 123:
		return name == "for"
	case 126:
		return name == "go"
	case 131:
		return name == "fallthrough"
	case 144:
		return name == "chan"
	case 151:
		return name == "map"
	case 176:
		return name == "case"
	case 205:
		return name == "break"
	case 231:
		return name == "package"
	case 237:
		return name == "range"
	case 244:
		return name == "type"
	case 251:
		return name == "var"
	}
	return false
}

func TwoHash_AddXor_Shift026(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 2)) ^ (name[1] << 6) {
	case 13:
		return name == "break"
	case 18:
		return name == "switch"
	case 38:
		return name == "if"
	case 45:
		return name == "interface"
	case 81:
		return name == "const"
	case 84:
		return name == "continue"
	case 91:
		return name == "for"
	case 94:
		return name == "go"
	case 96:
		return name == "goto"
	case 135:
		return name == "package"
	case 141:
		return name == "range"
	case 142:
		return name == "return"
	case 144:
		return name == "chan"
	case 146:
		return name == "select"
	case 148:
		return name == "type"
	case 152:
		return name == "else"
	case 155:
		return name == "var"
	case 208:
		return name == "case"
	case 210:
		return name == "struct"
	case 213:
		return name == "defer"
	case 215:
		return name == "default"
	case 220:
		return name == "func"
	case 227:
		return name == "fallthrough"
	case 234:
		return name == "import"
	case 247:
		return name == "map"
	}
	return false
}

func TwoHash_AddXor_Shift034(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 3)) ^ (name[1] << 4) {
	case 12:
		return name == "case"
	case 42:
		return name == "if"
	case 43:
		return name == "fallthrough"
	case 52:
		return name == "type"
	case 53:
		return name == "break"
	case 100:
		return name == "func"
	case 117:
		return name == "defer"
	case 119:
		return name == "default"
	case 123:
		return name == "map"
	case 133:
		return name == "range"
	case 151:
		return name == "package"
	case 156:
		return name == "chan"
	case 158:
		return name == "import"
	case 163:
		return name == "var"
	case 177:
		return name == "interface"
	case 195:
		return name == "for"
	case 198:
		return name == "return"
	case 202:
		return name == "go"
	case 204:
		return name == "goto"
	case 206:
		return name == "select"
	case 208:
		return name == "continue"
	case 222:
		return name == "struct"
	case 236:
		return name == "else"
	case 237:
		return name == "const"
	case 238:
		return name == "switch"
	}
	return false
}

func TwoHash_AddXor_Shift035(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 3)) ^ (name[1] << 5) {
	case 27:
		return name == "fallthrough"
	case 28:
		return name == "chan"
	case 30:
		return name == "struct"
	case 54:
		return name == "return"
	case 60:
		return name == "case"
	case 62:
		return name == "select"
	case 75:
		return name == "map"
	case 85:
		return name == "break"
	case 126:
		return name == "switch"
	case 132:
		return name == "type"
	case 133:
		return name == "defer"
	case 135:
		return name == "default"
	case 138:
		return name == "if"
	case 145:
		return name == "interface"
	case 147:
		return name == "var"
	case 148:
		return name == "func"
	case 167:
		return name == "package"
	case 172:
		return name == "else"
	case 181:
		return name == "range"
	case 192:
		return name == "continue"
	case 211:
		return name == "for"
	case 218:
		return name == "go"
	case 220:
		return name == "goto"
	case 238:
		return name == "import"
	case 253:
		return name == "const"
	}
	return false
}

func TwoHash_AddXor_Shift041(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 4)) ^ (name[1] << 1) {
	case 17:
		return name == "map"
	case 69:
		return name == "interface"
	case 76:
		return name == "import"
	case 94:
		return name == "if"
	case 140:
		return name == "else"
	case 141:
		return name == "default"
	case 142:
		return name == "func"
	case 143:
		return name == "defer"
	case 161:
		return name == "var"
	case 169:
		return name == "fallthrough"
	case 170:
		return name == "goto"
	case 172:
		return name == "go"
	case 182:
		return name == "type"
	case 189:
		return name == "for"
	case 193:
		return name == "break"
	case 197:
		return name == "package"
	case 216:
		return name == "switch"
	case 222:
		return name == "struct"
	case 228:
		return name == "chan"
	case 230:
		return name == "continue"
	case 231:
		return name == "range"
	case 235:
		return name == "const"
	case 236:
		return name == "return"
	case 246:
		return name == "case"
	case 252:
		return name == "select"
	}
	return false
}

func TwoHash_AddXor_Shift043(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 4)) ^ (name[1] << 3) {
	case 10:
		return name == "go"
	case 12:
		return name == "goto"
	case 14:
		return name == "return"
	case 15:
		return name == "package"
	case 27:
		return name == "for"
	case 30:
		return name == "select"
	case 45:
		return name == "range"
	case 52:
		return name == "else"
	case 60:
		return name == "case"
	case 64:
		return name == "continue"
	case 77:
		return name == "const"
	case 99:
		return name == "fallthrough"
	case 107:
		return name == "var"
	case 109:
		return name == "defer"
	case 111:
		return name == "default"
	case 116:
		return name == "chan"
	case 140:
		return name == "type"
	case 142:
		return name == "switch"
	case 150:
		return name == "struct"
	case 162:
		return name == "if"
	case 181:
		return name == "break"
	case 204:
		return name == "func"
	case 219:
		return name == "map"
	case 233:
		return name == "interface"
	case 254:
		return name == "import"
	}
	return false
}

func TwoHash_AddXor_Shift045(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 4)) ^ (name[1] << 5) {
	case 5:
		return name == "range"
	case 20:
		return name == "case"
	case 39:
		return name == "package"
	case 52:
		return name == "chan"
	case 54:
		return name == "import"
	case 67:
		return name == "var"
	case 75:
		return name == "fallthrough"
	case 82:
		return name == "if"
	case 89:
		return name == "interface"
	case 100:
		return name == "type"
	case 101:
		return name == "break"
	case 131:
		return name == "for"
	case 134:
		return name == "return"
	case 146:
		return name == "go"
	case 148:
		return name == "goto"
	case 150:
		return name == "select"
	case 182:
		return name == "struct"
	case 196:
		return name == "func"
	case 212:
		return name == "else"
	case 213:
		return name == "const"
	case 214:
		return name == "switch"
	case 216:
		return name == "continue"
	case 229:
		return name == "defer"
	case 231:
		return name == "default"
	case 243:
		return name == "map"
	}
	return false
}

func TwoHash_AddXor_Shift046(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 4)) ^ (name[1] << 6) {
	case 4:
		return name == "type"
	case 5:
		return name == "defer"
	case 7:
		return name == "default"
	case 18:
		return name == "if"
	case 25:
		return name == "interface"
	case 35:
		return name == "var"
	case 36:
		return name == "func"
	case 43:
		return name == "fallthrough"
	case 52:
		return name == "chan"
	case 54:
		return name == "struct"
	case 71:
		return name == "package"
	case 84:
		return name == "else"
	case 101:
		return name == "range"
	case 102:
		return name == "return"
	case 116:
		return name == "case"
	case 118:
		return name == "select"
	case 147:
		return name == "map"
	case 163:
		return name == "for"
	case 165:
		return name == "break"
	case 178:
		return name == "go"
	case 180:
		return name == "goto"
	case 214:
		return name == "import"
	case 245:
		return name == "const"
	case 246:
		return name == "switch"
	case 248:
		return name == "continue"
	}
	return false
}

func TwoHash_AddXor_Shift050(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 5)) ^ (name[1] << 0) {
	case 3:
		return name == "select"
	case 5:
		return name == "case"
	case 7:
		return name == "continue"
	case 10:
		return name == "const"
	case 12:
		return name == "chan"
	case 17:
		return name == "switch"
	case 18:
		return name == "struct"
	case 35:
		return name == "return"
	case 36:
		return name == "range"
	case 55:
		return name == "break"
	case 68:
		return name == "if"
	case 71:
		return name == "interface"
	case 75:
		return name == "import"
	case 102:
		return name == "package"
	case 139:
		return name == "goto"
	case 141:
		return name == "go"
	case 162:
		return name == "var"
	case 170:
		return name == "fallthrough"
	case 172:
		return name == "for"
	case 177:
		return name == "func"
	case 194:
		return name == "map"
	case 200:
		return name == "else"
	case 224:
		return name == "defer"
	case 226:
		return name == "default"
	case 253:
		return name == "type"
	}
	return false
}

func TwoHash_AddXor_Shift051(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 5)) ^ (name[1] << 1) {
	case 1:
		return name == "var"
	case 9:
		return name == "fallthrough"
	case 29:
		return name == "for"
	case 46:
		return name == "func"
	case 58:
		return name == "goto"
	case 60:
		return name == "go"
	case 77:
		return name == "default"
	case 79:
		return name == "defer"
	case 97:
		return name == "map"
	case 118:
		return name == "type"
	case 124:
		return name == "else"
	case 135:
		return name == "range"
	case 136:
		return name == "switch"
	case 140:
		return name == "return"
	case 142:
		return name == "struct"
	case 161:
		return name == "break"
	case 166:
		return name == "case"
	case 172:
		return name == "select"
	case 180:
		return name == "chan"
	case 182:
		return name == "continue"
	case 187:
		return name == "const"
	case 197:
		return name == "package"
	case 238:
		return name == "if"
	case 245:
		return name == "interface"
	case 252:
		return name == "import"
	}
	return false
}

func TwoHash_AddXor_Shift060(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 6)) ^ (name[1] << 0) {
	case 34:
		return name == "map"
	case 36:
		return name == "if"
	case 39:
		return name == "interface"
	case 40:
		return name == "else"
	case 43:
		return name == "import"
	case 96:
		return name == "defer"
	case 98:
		return name == "default"
	case 102:
		return name == "package"
	case 125:
		return name == "type"
	case 163:
		return name == "select"
	case 165:
		return name == "case"
	case 167:
		return name == "continue"
	case 170:
		return name == "const"
	case 171:
		return name == "goto"
	case 172:
		return name == "chan"
	case 173:
		return name == "go"
	case 177:
		return name == "switch"
	case 178:
		return name == "struct"
	case 226:
		return name == "var"
	case 227:
		return name == "return"
	case 228:
		return name == "range"
	case 234:
		return name == "fallthrough"
	case 236:
		return name == "for"
	case 241:
		return name == "func"
	case 247:
		return name == "break"
	}
	return false
}

func TwoHash_AddXor_Shift062(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 6)) ^ (name[1] << 2) {
	case 1:
		return name == "range"
	case 7:
		return name == "var"
	case 15:
		return name == "fallthrough"
	case 18:
		return name == "return"
	case 22:
		return name == "struct"
	case 26:
		return name == "switch"
	case 63:
		return name == "for"
	case 64:
		return name == "case"
	case 77:
		return name == "break"
	case 80:
		return name == "func"
	case 82:
		return name == "select"
	case 100:
		return name == "chan"
	case 116:
		return name == "continue"
	case 120:
		return name == "goto"
	case 121:
		return name == "const"
	case 126:
		return name == "go"
	case 131:
		return name == "package"
	case 145:
		return name == "defer"
	case 147:
		return name == "default"
	case 199:
		return name == "map"
	case 218:
		return name == "if"
	case 224:
		return name == "type"
	case 241:
		return name == "interface"
	case 242:
		return name == "import"
	case 244:
		return name == "else"
	}
	return false
}

func TwoHash_AddXor_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 7)) ^ (name[1] << 3) {
	case 3:
		return name == "fallthrough"
	case 11:
		return name == "var"
	case 13:
		return name == "range"
	case 15:
		return name == "package"
	case 38:
		return name == "struct"
	case 45:
		return name == "defer"
	case 46:
		return name == "return"
	case 47:
		return name == "default"
	case 62:
		return name == "switch"
	case 123:
		return name == "for"
	case 139:
		return name == "map"
	case 140:
		return name == "case"
	case 149:
		return name == "break"
	case 172:
		return name == "func"
	case 174:
		return name == "select"
	case 178:
		return name == "if"
	case 196:
		return name == "chan"
	case 204:
		return name == "type"
	case 228:
		return name == "else"
	case 238:
		return name == "import"
	case 240:
		return name == "continue"
	case 249:
		return name == "interface"
	case 250:
		return name == "go"
	case 252:
		return name == "goto"
	case 253:
		return name == "const"
	}
	return false
}

func TwoHash_AddXor_Shift135(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) + (name[0] << 3)) ^ (name[1] << 5) {
	case 0:
		return name == "case"
	case 4:
		return name == "select"
	case 32:
		return name == "chan"
	case 36:
		return name == "struct"
	case 60:
		return name == "return"
	case 68:
		return name == "switch"
	case 78:
		return name == "map"
	case 90:
		return name == "break"
	case 102:
		return name == "fallthrough"
	case 136:
		return name == "type"
	case 138:
		return name == "defer"
	case 140:
		return name == "if"
	case 142:
		return name == "default"
	case 150:
		return name == "var"
	case 152:
		return name == "func"
	case 154:
		return name == "interface"
	case 160:
		return name == "goto"
	case 174:
		return name == "package"
	case 176:
		return name == "else"
	case 186:
		return name == "range"
	case 194:
		return name == "const"
	case 200:
		return name == "continue"
	case 214:
		return name == "for"
	case 220:
		return name == "go"
	case 244:
		return name == "import"
	}
	return false
}

func TwoHash_AddXor_Shift145(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) + (name[0] << 4)) ^ (name[1] << 5) {
	case 10:
		return name == "range"
	case 24:
		return name == "case"
	case 46:
		return name == "package"
	case 56:
		return name == "chan"
	case 60:
		return name == "import"
	case 70:
		return name == "var"
	case 84:
		return name == "if"
	case 86:
		return name == "fallthrough"
	case 98:
		return name == "interface"
	case 104:
		return name == "type"
	case 106:
		return name == "break"
	case 134:
		return name == "for"
	case 140:
		return name == "return"
	case 148:
		return name == "go"
	case 152:
		return name == "goto"
	case 156:
		return name == "select"
	case 160:
		return name == "continue"
	case 188:
		return name == "struct"
	case 200:
		return name == "func"
	case 216:
		return name == "else"
	case 218:
		return name == "const"
	case 220:
		return name == "switch"
	case 234:
		return name == "defer"
	case 238:
		return name == "default"
	case 246:
		return name == "map"
	}
	return false
}

func TwoHash_AddXor_Shift146(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) + (name[0] << 4)) ^ (name[1] << 6) {
	case 8:
		return name == "type"
	case 10:
		return name == "defer"
	case 14:
		return name == "default"
	case 20:
		return name == "if"
	case 34:
		return name == "interface"
	case 38:
		return name == "var"
	case 40:
		return name == "func"
	case 54:
		return name == "fallthrough"
	case 56:
		return name == "chan"
	case 60:
		return name == "struct"
	case 78:
		return name == "package"
	case 88:
		return name == "else"
	case 106:
		return name == "range"
	case 108:
		return name == "return"
	case 120:
		return name == "case"
	case 124:
		return name == "select"
	case 128:
		return name == "continue"
	case 150:
		return name == "map"
	case 166:
		return name == "for"
	case 170:
		return name == "break"
	case 180:
		return name == "go"
	case 184:
		return name == "goto"
	case 220:
		return name == "import"
	case 250:
		return name == "const"
	case 252:
		return name == "switch"
	}
	return false
}

func TwoHash_AddXor_Shift161(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) + (name[0] << 6)) ^ (name[1] << 1) {
	case 6:
		return name == "select"
	case 10:
		return name == "case"
	case 14:
		return name == "continue"
	case 20:
		return name == "const"
	case 22:
		return name == "goto"
	case 24:
		return name == "chan"
	case 26:
		return name == "go"
	case 34:
		return name == "switch"
	case 36:
		return name == "struct"
	case 68:
		return name == "var"
	case 70:
		return name == "return"
	case 72:
		return name == "range"
	case 84:
		return name == "fallthrough"
	case 88:
		return name == "for"
	case 98:
		return name == "func"
	case 110:
		return name == "break"
	case 132:
		return name == "map"
	case 136:
		return name == "if"
	case 142:
		return name == "interface"
	case 144:
		return name == "else"
	case 150:
		return name == "import"
	case 192:
		return name == "defer"
	case 196:
		return name == "default"
	case 204:
		return name == "package"
	case 250:
		return name == "type"
	}
	return false
}

func TwoHash_AddXor_Shift230(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) + (name[0] << 3)) ^ (name[1] << 0) {
	case 2:
		return name == "interface"
	case 13:
		return name == "import"
	case 21:
		return name == "map"
	case 39:
		return name == "goto"
	case 47:
		return name == "go"
	case 53:
		return name == "func"
	case 54:
		return name == "if"
	case 61:
		return name == "fallthrough"
	case 64:
		return name == "chan"
	case 67:
		return name == "const"
	case 73:
		return name == "case"
	case 81:
		return name == "defer"
	case 83:
		return name == "for"
	case 84:
		return name == "else"
	case 86:
		return name == "break"
	case 87:
		return name == "continue"
	case 89:
		return name == "default"
	case 196:
		return name == "struct"
	case 197:
		return name == "range"
	case 199:
		return name == "switch"
	case 201:
		return name == "type"
	case 205:
		return name == "return"
	case 213:
		return name == "select"
	case 221:
		return name == "var"
	case 253:
		return name == "package"
	}
	return false
}

func TwoHash_AddXor_Shift250(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) + (name[0] << 5)) ^ (name[1] << 0) {
	case 12:
		return name == "struct"
	case 15:
		return name == "switch"
	case 17:
		return name == "case"
	case 24:
		return name == "chan"
	case 27:
		return name == "const"
	case 29:
		return name == "select"
	case 38:
		return name == "break"
	case 42:
		return name == "interface"
	case 53:
		return name == "range"
	case 61:
		return name == "return"
	case 78:
		return name == "if"
	case 85:
		return name == "import"
	case 125:
		return name == "package"
	case 135:
		return name == "go"
	case 141:
		return name == "fallthrough"
	case 159:
		return name == "goto"
	case 163:
		return name == "for"
	case 165:
		return name == "func"
	case 173:
		return name == "var"
	case 205:
		return name == "map"
	case 220:
		return name == "else"
	case 233:
		return name == "type"
	case 239:
		return name == "continue"
	case 241:
		return name == "defer"
	case 249:
		return name == "default"
	}
	return false
}

func TwoHash_AddXor_Shift260(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) + (name[0] << 6)) ^ (name[1] << 0) {
	case 10:
		return name == "interface"
	case 45:
		return name == "map"
	case 46:
		return name == "if"
	case 53:
		return name == "import"
	case 60:
		return name == "else"
	case 105:
		return name == "type"
	case 113:
		return name == "defer"
	case 121:
		return name == "default"
	case 125:
		return name == "package"
	case 143:
		return name == "continue"
	case 167:
		return name == "go"
	case 172:
		return name == "struct"
	case 175:
		return name == "switch"
	case 177:
		return name == "case"
	case 184:
		return name == "chan"
	case 187:
		return name == "const"
	case 189:
		return name == "select"
	case 191:
		return name == "goto"
	case 205:
		return name == "fallthrough"
	case 227:
		return name == "for"
	case 229:
		return name == "func"
	case 230:
		return name == "break"
	case 237:
		return name == "var"
	case 245:
		return name == "range"
	case 253:
		return name == "return"
	}
	return false
}

func TwoHash_AddXor_Shift301(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) + (name[0] << 0)) ^ (name[1] << 1) {
	case 65:
		return name == "case"
	case 67:
		return name == "import"
	case 70:
		return name == "defer"
	case 71:
		return name == "map"
	case 75:
		return name == "struct"
	case 76:
		return name == "var"
	case 77:
		return name == "switch"
	case 83:
		return name == "chan"
	case 85:
		return name == "const"
	case 86:
		return name == "default"
	case 88:
		return name == "range"
	case 89:
		return name == "goto"
	case 93:
		return name == "else"
	case 102:
		return name == "type"
	case 104:
		return name == "return"
	case 105:
		return name == "select"
	case 106:
		return name == "package"
	case 108:
		return name == "func"
	case 109:
		return name == "interface"
	case 110:
		return name == "break"
	case 124:
		return name == "fallthrough"
	case 125:
		return name == "continue"
	case 160:
		return name == "for"
	case 169:
		return name == "go"
	case 181:
		return name == "if"
	}
	return false
}

func TwoHash_AddXor_Shift304(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) + (name[0] << 0)) ^ (name[1] << 4) {
	case 3:
		return name == "chan"
	case 4:
		return name == "type"
	case 25:
		return name == "if"
	case 69:
		return name == "else"
	case 73:
		return name == "import"
	case 81:
		return name == "interface"
	case 83:
		return name == "continue"
	case 119:
		return name == "goto"
	case 123:
		return name == "const"
	case 135:
		return name == "go"
	case 138:
		return name == "range"
	case 142:
		return name == "for"
	case 147:
		return name == "case"
	case 149:
		return name == "map"
	case 158:
		return name == "var"
	case 170:
		return name == "break"
	case 174:
		return name == "fallthrough"
	case 184:
		return name == "package"
	case 204:
		return name == "default"
	case 211:
		return name == "switch"
	case 214:
		return name == "func"
	case 220:
		return name == "defer"
	case 227:
		return name == "struct"
	case 242:
		return name == "return"
	case 243:
		return name == "select"
	}
	return false
}

func TwoHash_AddXor_Shift361(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) + (name[0] << 6)) ^ (name[1] << 1) {
	case 14:
		return name == "go"
	case 24:
		return name == "struct"
	case 26:
		return name == "fallthrough"
	case 30:
		return name == "switch"
	case 34:
		return name == "case"
	case 48:
		return name == "chan"
	case 54:
		return name == "const"
	case 58:
		return name == "select"
	case 62:
		return name == "goto"
	case 70:
		return name == "for"
	case 74:
		return name == "func"
	case 76:
		return name == "break"
	case 84:
		return name == "interface"
	case 90:
		return name == "var"
	case 106:
		return name == "range"
	case 122:
		return name == "return"
	case 154:
		return name == "map"
	case 156:
		return name == "if"
	case 170:
		return name == "import"
	case 184:
		return name == "else"
	case 210:
		return name == "type"
	case 222:
		return name == "continue"
	case 226:
		return name == "defer"
	case 242:
		return name == "default"
	case 250:
		return name == "package"
	}
	return false
}

func TwoHash_AddXor_Shift403(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) + (name[0] << 0)) ^ (name[1] << 3) {
	case 14:
		return name == "func"
	case 30:
		return name == "fallthrough"
	case 34:
		return name == "break"
	case 107:
		return name == "switch"
	case 115:
		return name == "struct"
	case 124:
		return name == "type"
	case 137:
		return name == "interface"
	case 149:
		return name == "map"
	case 155:
		return name == "continue"
	case 156:
		return name == "defer"
	case 161:
		return name == "import"
	case 171:
		return name == "case"
	case 174:
		return name == "var"
	case 185:
		return name == "if"
	case 197:
		return name == "else"
	case 202:
		return name == "range"
	case 203:
		return name == "const"
	case 223:
		return name == "goto"
	case 227:
		return name == "chan"
	case 232:
		return name == "package"
	case 238:
		return name == "for"
	case 250:
		return name == "return"
	case 251:
		return name == "select"
	case 252:
		return name == "default"
	case 255:
		return name == "go"
	}
	return false
}

func TwoHash_AddXor_Shift412(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) + (name[0] << 1)) ^ (name[1] << 2) {
	case 64:
		return name == "for"
	case 82:
		return name == "go"
	case 106:
		return name == "if"
	case 130:
		return name == "case"
	case 134:
		return name == "import"
	case 140:
		return name == "defer"
	case 142:
		return name == "map"
	case 150:
		return name == "struct"
	case 152:
		return name == "var"
	case 154:
		return name == "switch"
	case 166:
		return name == "chan"
	case 170:
		return name == "const"
	case 172:
		return name == "default"
	case 176:
		return name == "range"
	case 178:
		return name == "goto"
	case 186:
		return name == "else"
	case 204:
		return name == "type"
	case 208:
		return name == "return"
	case 210:
		return name == "select"
	case 212:
		return name == "package"
	case 216:
		return name == "func"
	case 218:
		return name == "interface"
	case 220:
		return name == "break"
	case 248:
		return name == "fallthrough"
	case 250:
		return name == "continue"
	}
	return false
}

func TwoHash_AddXor_Shift510(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) + (name[0] << 1)) ^ (name[1] << 0) {
	case 9:
		return name == "const"
	case 13:
		return name == "defer"
	case 17:
		return name == "type"
	case 22:
		return name == "break"
	case 33:
		return name == "goto"
	case 38:
		return name == "else"
	case 39:
		return name == "case"
	case 45:
		return name == "var"
	case 46:
		return name == "chan"
	case 57:
		return name == "func"
	case 67:
		return name == "for"
	case 77:
		return name == "fallthrough"
	case 91:
		return name == "map"
	case 97:
		return name == "go"
	case 116:
		return name == "if"
	case 156:
		return name == "interface"
	case 161:
		return name == "package"
	case 169:
		return name == "continue"
	case 193:
		return name == "return"
	case 195:
		return name == "select"
	case 205:
		return name == "default"
	case 209:
		return name == "switch"
	case 210:
		return name == "struct"
	case 229:
		return name == "range"
	case 255:
		return name == "import"
	}
	return false
}

func TwoHash_AddXor_Shift630(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) + (name[0] << 3)) ^ (name[1] << 0) {
	case 5:
		return name == "defer"
	case 17:
		return name == "var"
	case 33:
		return name == "package"
	case 34:
		return name == "break"
	case 55:
		return name == "const"
	case 68:
		return name == "else"
	case 69:
		return name == "func"
	case 73:
		return name == "map"
	case 87:
		return name == "goto"
	case 108:
		return name == "struct"
	case 111:
		return name == "switch"
	case 112:
		return name == "chan"
	case 117:
		return name == "return"
	case 119:
		return name == "continue"
	case 121:
		return name == "case"
	case 125:
		return name == "select"
	case 133:
		return name == "default"
	case 145:
		return name == "fallthrough"
	case 159:
		return name == "for"
	case 165:
		return name == "import"
	case 174:
		return name == "if"
	case 177:
		return name == "range"
	case 215:
		return name == "go"
	case 217:
		return name == "type"
	case 230:
		return name == "interface"
	}
	return false
}

func TwoHash_AddAdd_Shift022(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 2)) + (name[1] << 2) {
	case 20:
		return name == "case"
	case 39:
		return name == "fallthrough"
	case 41:
		return name == "defer"
	case 43:
		return name == "default"
	case 48:
		return name == "chan"
	case 59:
		return name == "map"
	case 62:
		return name == "if"
	case 72:
		return name == "else"
	case 75:
		return name == "package"
	case 77:
		return name == "const"
	case 80:
		return name == "continue"
	case 81:
		return name == "range"
	case 85:
		return name == "break"
	case 87:
		return name == "for"
	case 90:
		return name == "go"
	case 92:
		return name == "goto"
	case 94:
		return name == "import"
	case 95:
		return name == "var"
	case 98:
		return name == "return"
	case 101:
		return name == "interface"
	case 102:
		return name == "select"
	case 112:
		return name == "func"
	case 162:
		return name == "struct"
	case 174:
		return name == "switch"
	case 184:
		return name == "type"
	}
	return false
}

func TwoHash_AddAdd_Shift024(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 2)) + (name[1] << 4) {
	case 6:
		return name == "if"
	case 16:
		return name == "chan"
	case 18:
		return name == "struct"
	case 30:
		return name == "return"
	case 34:
		return name == "select"
	case 66:
		return name == "switch"
	case 88:
		return name == "else"
	case 100:
		return name == "type"
	case 122:
		return name == "import"
	case 129:
		return name == "const"
	case 132:
		return name == "continue"
	case 139:
		return name == "for"
	case 141:
		return name == "interface"
	case 142:
		return name == "go"
	case 144:
		return name == "goto"
	case 160:
		return name == "case"
	case 173:
		return name == "break"
	case 179:
		return name == "fallthrough"
	case 199:
		return name == "map"
	case 215:
		return name == "package"
	case 221:
		return name == "range"
	case 229:
		return name == "defer"
	case 231:
		return name == "default"
	case 235:
		return name == "var"
	case 236:
		return name == "func"
	}
	return false
}

func TwoHash_AddAdd_Shift025(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 2)) + (name[1] << 5) {
	case 24:
		return name == "else"
	case 53:
		return name == "defer"
	case 55:
		return name == "default"
	case 60:
		return name == "func"
	case 74:
		return name == "import"
	case 82:
		return name == "struct"
	case 102:
		return name == "if"
	case 109:
		return name == "interface"
	case 110:
		return name == "return"
	case 113:
		return name == "const"
	case 114:
		return name == "select"
	case 116:
		return name == "continue"
	case 123:
		return name == "for"
	case 126:
		return name == "go"
	case 128:
		return name == "goto"
	case 144:
		return name == "chan"
	case 176:
		return name == "case"
	case 178:
		return name == "switch"
	case 195:
		return name == "fallthrough"
	case 205:
		return name == "break"
	case 215:
		return name == "map"
	case 231:
		return name == "package"
	case 237:
		return name == "range"
	case 244:
		return name == "type"
	case 251:
		return name == "var"
	}
	return false
}

func TwoHash_AddAdd_Shift033(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 3)) + (name[1] << 3) {
	case 36:
		return name == "case"
	case 62:
		return name == "struct"
	case 67:
		return name == "fallthrough"
	case 77:
		return name == "defer"
	case 79:
		return name == "default"
	case 86:
		return name == "switch"
	case 92:
		return name == "chan"
	case 108:
		return name == "type"
	case 115:
		return name == "map"
	case 122:
		return name == "if"
	case 140:
		return name == "else"
	case 143:
		return name == "package"
	case 149:
		return name == "const"
	case 152:
		return name == "continue"
	case 157:
		return name == "range"
	case 165:
		return name == "break"
	case 171:
		return name == "for"
	case 178:
		return name == "go"
	case 180:
		return name == "goto"
	case 182:
		return name == "import"
	case 187:
		return name == "var"
	case 190:
		return name == "return"
	case 193:
		return name == "interface"
	case 198:
		return name == "select"
	case 220:
		return name == "func"
	}
	return false
}

func TwoHash_AddAdd_Shift041(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 4)) + (name[1] << 1) {
	case 0:
		return name == "select"
	case 4:
		return name == "chan"
	case 9:
		return name == "break"
	case 15:
		return name == "defer"
	case 17:
		return name == "default"
	case 19:
		return name == "const"
	case 22:
		return name == "continue"
	case 30:
		return name == "struct"
	case 36:
		return name == "switch"
	case 37:
		return name == "var"
	case 44:
		return name == "else"
	case 45:
		return name == "fallthrough"
	case 54:
		return name == "type"
	case 65:
		return name == "for"
	case 78:
		return name == "func"
	case 80:
		return name == "go"
	case 82:
		return name == "goto"
	case 94:
		return name == "if"
	case 112:
		return name == "import"
	case 117:
		return name == "interface"
	case 149:
		return name == "map"
	case 201:
		return name == "package"
	case 231:
		return name == "range"
	case 240:
		return name == "return"
	case 246:
		return name == "case"
	}
	return false
}

func TwoHash_AddAdd_Shift042(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 4)) + (name[1] << 2) {
	case 4:
		return name == "else"
	case 6:
		return name == "struct"
	case 18:
		return name == "switch"
	case 31:
		return name == "for"
	case 40:
		return name == "type"
	case 42:
		return name == "if"
	case 46:
		return name == "go"
	case 48:
		return name == "goto"
	case 56:
		return name == "func"
	case 74:
		return name == "import"
	case 81:
		return name == "interface"
	case 87:
		return name == "map"
	case 139:
		return name == "package"
	case 169:
		return name == "range"
	case 184:
		return name == "case"
	case 186:
		return name == "return"
	case 202:
		return name == "select"
	case 212:
		return name == "chan"
	case 217:
		return name == "defer"
	case 219:
		return name == "default"
	case 231:
		return name == "var"
	case 237:
		return name == "break"
	case 239:
		return name == "fallthrough"
	case 241:
		return name == "const"
	case 244:
		return name == "continue"
	}
	return false
}

func TwoHash_AddAdd_Shift050(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 5)) + (name[1] << 0) {
	case 4:
		return name == "map"
	case 16:
		return name == "else"
	case 36:
		return name == "var"
	case 44:
		return name == "fallthrough"
	case 50:
		return name == "for"
	case 57:
		return name == "func"
	case 81:
		return name == "go"
	case 83:
		return name == "goto"
	case 104:
		return name == "package"
	case 136:
		return name == "if"
	case 147:
		return name == "import"
	case 151:
		return name == "interface"
	case 166:
		return name == "range"
	case 171:
		return name == "return"
	case 183:
		return name == "break"
	case 197:
		return name == "case"
	case 203:
		return name == "select"
	case 204:
		return name == "chan"
	case 212:
		return name == "const"
	case 215:
		return name == "continue"
	case 218:
		return name == "struct"
	case 221:
		return name == "switch"
	case 234:
		return name == "defer"
	case 236:
		return name == "default"
	case 253:
		return name == "type"
	}
	return false
}

func TwoHash_AddAdd_Shift051(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 5)) + (name[1] << 1) {
	case 0:
		return name == "import"
	case 5:
		return name == "interface"
	case 7:
		return name == "range"
	case 16:
		return name == "return"
	case 38:
		return name == "case"
	case 41:
		return name == "break"
	case 48:
		return name == "select"
	case 52:
		return name == "chan"
	case 67:
		return name == "const"
	case 70:
		return name == "continue"
	case 78:
		return name == "struct"
	case 79:
		return name == "defer"
	case 81:
		return name == "default"
	case 84:
		return name == "switch"
	case 101:
		return name == "map"
	case 118:
		return name == "type"
	case 124:
		return name == "else"
	case 133:
		return name == "var"
	case 141:
		return name == "fallthrough"
	case 161:
		return name == "for"
	case 174:
		return name == "func"
	case 192:
		return name == "go"
	case 194:
		return name == "goto"
	case 201:
		return name == "package"
	case 238:
		return name == "if"
	}
	return false
}

func TwoHash_AddAdd_Shift060(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 6)) + (name[1] << 0) {
	case 37:
		return name == "case"
	case 43:
		return name == "select"
	case 44:
		return name == "chan"
	case 49:
		return name == "go"
	case 51:
		return name == "goto"
	case 52:
		return name == "const"
	case 55:
		return name == "continue"
	case 58:
		return name == "struct"
	case 61:
		return name == "switch"
	case 104:
		return name == "package"
	case 106:
		return name == "defer"
	case 108:
		return name == "default"
	case 125:
		return name == "type"
	case 164:
		return name == "map"
	case 168:
		return name == "if"
	case 176:
		return name == "else"
	case 179:
		return name == "import"
	case 183:
		return name == "interface"
	case 228:
		return name == "var"
	case 230:
		return name == "range"
	case 235:
		return name == "return"
	case 236:
		return name == "fallthrough"
	case 242:
		return name == "for"
	case 247:
		return name == "break"
	case 249:
		return name == "func"
	}
	return false
}

func TwoHash_AddAdd_Shift061(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 6)) + (name[1] << 1) {
	case 5:
		return name == "map"
	case 14:
		return name == "if"
	case 28:
		return name == "else"
	case 32:
		return name == "import"
	case 37:
		return name == "interface"
	case 69:
		return name == "var"
	case 71:
		return name == "range"
	case 77:
		return name == "fallthrough"
	case 80:
		return name == "return"
	case 97:
		return name == "for"
	case 105:
		return name == "break"
	case 110:
		return name == "func"
	case 134:
		return name == "case"
	case 144:
		return name == "select"
	case 148:
		return name == "chan"
	case 160:
		return name == "go"
	case 162:
		return name == "goto"
	case 163:
		return name == "const"
	case 166:
		return name == "continue"
	case 174:
		return name == "struct"
	case 180:
		return name == "switch"
	case 201:
		return name == "package"
	case 207:
		return name == "defer"
	case 209:
		return name == "default"
	case 246:
		return name == "type"
	}
	return false
}

func TwoHash_AddAdd_Shift062(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 6)) + (name[1] << 2) {
	case 1:
		return name == "interface"
	case 7:
		return name == "var"
	case 9:
		return name == "range"
	case 15:
		return name == "fallthrough"
	case 26:
		return name == "return"
	case 63:
		return name == "for"
	case 72:
		return name == "case"
	case 77:
		return name == "break"
	case 88:
		return name == "func"
	case 90:
		return name == "select"
	case 100:
		return name == "chan"
	case 126:
		return name == "go"
	case 128:
		return name == "goto"
	case 129:
		return name == "const"
	case 132:
		return name == "continue"
	case 139:
		return name == "package"
	case 150:
		return name == "struct"
	case 153:
		return name == "defer"
	case 155:
		return name == "default"
	case 162:
		return name == "switch"
	case 199:
		return name == "map"
	case 218:
		return name == "if"
	case 232:
		return name == "type"
	case 244:
		return name == "else"
	case 250:
		return name == "import"
	}
	return false
}

func TwoHash_AddAdd_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) + (name[0] << 7)) + (name[1] << 3) {
	case 0:
		return name == "continue"
	case 11:
		return name == "var"
	case 13:
		return name == "range"
	case 15:
		return name == "package"
	case 19:
		return name == "fallthrough"
	case 38:
		return name == "struct"
	case 45:
		return name == "defer"
	case 46:
		return name == "return"
	case 47:
		return name == "default"
	case 62:
		return name == "switch"
	case 123:
		return name == "for"
	case 139:
		return name == "map"
	case 140:
		return name == "case"
	case 149:
		return name == "break"
	case 172:
		return name == "func"
	case 174:
		return name == "select"
	case 178:
		return name == "if"
	case 196:
		return name == "chan"
	case 204:
		return name == "type"
	case 228:
		return name == "else"
	case 238:
		return name == "import"
	case 249:
		return name == "interface"
	case 250:
		return name == "go"
	case 252:
		return name == "goto"
	case 253:
		return name == "const"
	}
	return false
}

func TwoHash_AddAdd_Shift130(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) + (name[0] << 3)) + (name[1] << 0) {
	case 1:
		return name == "return"
	case 9:
		return name == "select"
	case 23:
		return name == "var"
	case 24:
		return name == "struct"
	case 27:
		return name == "switch"
	case 33:
		return name == "type"
	case 129:
		return name == "case"
	case 136:
		return name == "chan"
	case 140:
		return name == "break"
	case 143:
		return name == "defer"
	case 145:
		return name == "const"
	case 147:
		return name == "default"
	case 151:
		return name == "continue"
	case 156:
		return name == "else"
	case 165:
		return name == "for"
	case 167:
		return name == "fallthrough"
	case 171:
		return name == "go"
	case 173:
		return name == "func"
	case 175:
		return name == "goto"
	case 178:
		return name == "if"
	case 193:
		return name == "import"
	case 200:
		return name == "interface"
	case 207:
		return name == "map"
	case 239:
		return name == "package"
	case 251:
		return name == "range"
	}
	return false
}

func TwoHash_AddAdd_Shift133(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) + (name[0] << 3)) + (name[1] << 3) {
	case 40:
		return name == "case"
	case 68:
		return name == "struct"
	case 78:
		return name == "fallthrough"
	case 82:
		return name == "defer"
	case 86:
		return name == "default"
	case 92:
		return name == "switch"
	case 96:
		return name == "chan"
	case 112:
		return name == "type"
	case 118:
		return name == "map"
	case 124:
		return name == "if"
	case 144:
		return name == "else"
	case 150:
		return name == "package"
	case 154:
		return name == "const"
	case 160:
		return name == "continue"
	case 162:
		return name == "range"
	case 170:
		return name == "break"
	case 174:
		return name == "for"
	case 180:
		return name == "go"
	case 184:
		return name == "goto"
	case 188:
		return name == "import"
	case 190:
		return name == "var"
	case 196:
		return name == "return"
	case 202:
		return name == "interface"
	case 204:
		return name == "select"
	case 224:
		return name == "func"
	}
	return false
}

func TwoHash_AddAdd_Shift150(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) + (name[0] << 5)) + (name[1] << 0) {
	case 1:
		return name == "type"
	case 7:
		return name == "map"
	case 20:
		return name == "else"
	case 39:
		return name == "var"
	case 53:
		return name == "for"
	case 55:
		return name == "fallthrough"
	case 61:
		return name == "func"
	case 83:
		return name == "go"
	case 87:
		return name == "goto"
	case 111:
		return name == "package"
	case 138:
		return name == "if"
	case 153:
		return name == "import"
	case 160:
		return name == "interface"
	case 171:
		return name == "range"
	case 177:
		return name == "return"
	case 188:
		return name == "break"
	case 201:
		return name == "case"
	case 208:
		return name == "chan"
	case 209:
		return name == "select"
	case 217:
		return name == "const"
	case 223:
		return name == "continue"
	case 224:
		return name == "struct"
	case 227:
		return name == "switch"
	case 239:
		return name == "defer"
	case 243:
		return name == "default"
	}
	return false
}

func TwoHash_AddAdd_Shift161(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) + (name[0] << 6)) + (name[1] << 1) {
	case 8:
		return name == "map"
	case 16:
		return name == "if"
	case 32:
		return name == "else"
	case 38:
		return name == "import"
	case 46:
		return name == "interface"
	case 72:
		return name == "var"
	case 76:
		return name == "range"
	case 86:
		return name == "return"
	case 88:
		return name == "fallthrough"
	case 100:
		return name == "for"
	case 110:
		return name == "break"
	case 114:
		return name == "func"
	case 138:
		return name == "case"
	case 150:
		return name == "select"
	case 152:
		return name == "chan"
	case 162:
		return name == "go"
	case 166:
		return name == "goto"
	case 168:
		return name == "const"
	case 174:
		return name == "continue"
	case 180:
		return name == "struct"
	case 186:
		return name == "switch"
	case 208:
		return name == "package"
	case 212:
		return name == "defer"
	case 216:
		return name == "default"
	case 250:
		return name == "type"
	}
	return false
}

func TwoHash_AddAdd_Shift230(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) + (name[0] << 3)) + (name[1] << 0) {
	case 5:
		return name == "range"
	case 13:
		return name == "return"
	case 21:
		return name == "select"
	case 29:
		return name == "var"
	case 36:
		return name == "struct"
	case 39:
		return name == "switch"
	case 41:
		return name == "type"
	case 137:
		return name == "case"
	case 144:
		return name == "chan"
	case 150:
		return name == "break"
	case 153:
		return name == "defer"
	case 155:
		return name == "const"
	case 161:
		return name == "default"
	case 164:
		return name == "else"
	case 167:
		return name == "continue"
	case 171:
		return name == "for"
	case 175:
		return name == "go"
	case 181:
		return name == "func"
	case 182:
		return name == "if"
	case 183:
		return name == "goto"
	case 189:
		return name == "fallthrough"
	case 205:
		return name == "import"
	case 213:
		return name == "map"
	case 218:
		return name == "interface"
	case 253:
		return name == "package"
	}
	return false
}

func TwoHash_AddAdd_Shift302(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) + (name[0] << 0)) + (name[1] << 2) {
	case 7:
		return name == "case"
	case 9:
		return name == "map"
	case 17:
		return name == "if"
	case 18:
		return name == "var"
	case 30:
		return name == "range"
	case 32:
		return name == "defer"
	case 35:
		return name == "chan"
	case 44:
		return name == "package"
	case 48:
		return name == "default"
	case 51:
		return name == "go"
	case 53:
		return name == "else"
	case 54:
		return name == "return"
	case 55:
		return name == "select"
	case 58:
		return name == "for"
	case 66:
		return name == "fallthrough"
	case 67:
		return name == "goto"
	case 71:
		return name == "const"
	case 77:
		return name == "import"
	case 82:
		return name == "break"
	case 90:
		return name == "func"
	case 95:
		return name == "continue"
	case 105:
		return name == "interface"
	case 115:
		return name == "struct"
	case 120:
		return name == "type"
	case 127:
		return name == "switch"
	}
	return false
}

func TwoHash_AddAdd_Shift303(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) + (name[0] << 0)) + (name[1] << 3) {
	case 1:
		return name == "import"
	case 3:
		return name == "const"
	case 26:
		return name == "break"
	case 27:
		return name == "continue"
	case 33:
		return name == "interface"
	case 46:
		return name == "func"
	case 67:
		return name == "struct"
	case 91:
		return name == "switch"
	case 92:
		return name == "type"
	case 139:
		return name == "case"
	case 141:
		return name == "map"
	case 150:
		return name == "var"
	case 162:
		return name == "range"
	case 169:
		return name == "if"
	case 176:
		return name == "package"
	case 180:
		return name == "defer"
	case 195:
		return name == "chan"
	case 196:
		return name == "default"
	case 198:
		return name == "fallthrough"
	case 202:
		return name == "return"
	case 203:
		return name == "select"
	case 229:
		return name == "else"
	case 239:
		return name == "go"
	case 246:
		return name == "for"
	case 255:
		return name == "goto"
	}
	return false
}

func TwoHash_AddAdd_Shift400(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) + (name[0] << 0)) + (name[1] << 0) {
	case 4:
		return name == "case"
	case 5:
		return name == "for"
	case 7:
		return name == "var"
	case 11:
		return name == "chan"
	case 17:
		return name == "else"
	case 22:
		return name == "goto"
	case 25:
		return name == "defer"
	case 27:
		return name == "func"
	case 34:
		return name == "const"
	case 35:
		return name == "range"
	case 36:
		return name == "break"
	case 45:
		return name == "type"
	case 54:
		return name == "import"
	case 55:
		return name == "return"
	case 56:
		return name == "select"
	case 57:
		return name == "default"
	case 65:
		return name == "package"
	case 71:
		return name == "struct"
	case 74:
		return name == "switch"
	case 82:
		return name == "continue"
	case 103:
		return name == "interface"
	case 119:
		return name == "fallthrough"
	case 239:
		return name == "if"
	case 246:
		return name == "go"
	case 254:
		return name == "map"
	}
	return false
}

func TwoHash_AddAdd_Shift403(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) + (name[0] << 0)) + (name[1] << 3) {
	case 5:
		return name == "else"
	case 14:
		return name == "for"
	case 30:
		return name == "fallthrough"
	case 31:
		return name == "goto"
	case 43:
		return name == "const"
	case 49:
		return name == "import"
	case 66:
		return name == "break"
	case 78:
		return name == "func"
	case 91:
		return name == "continue"
	case 105:
		return name == "interface"
	case 115:
		return name == "struct"
	case 124:
		return name == "type"
	case 139:
		return name == "switch"
	case 165:
		return name == "map"
	case 171:
		return name == "case"
	case 174:
		return name == "var"
	case 185:
		return name == "if"
	case 202:
		return name == "range"
	case 220:
		return name == "defer"
	case 227:
		return name == "chan"
	case 232:
		return name == "package"
	case 250:
		return name == "return"
	case 251:
		return name == "select"
	case 252:
		return name == "default"
	case 255:
		return name == "go"
	}
	return false
}

func TwoHash_AddAdd_Shift413(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 4) + (name[0] << 1)) + (name[1] << 3) {
	case 14:
		return name == "case"
	case 18:
		return name == "map"
	case 34:
		return name == "if"
	case 36:
		return name == "var"
	case 60:
		return name == "range"
	case 64:
		return name == "defer"
	case 70:
		return name == "chan"
	case 88:
		return name == "package"
	case 96:
		return name == "default"
	case 102:
		return name == "go"
	case 106:
		return name == "else"
	case 108:
		return name == "return"
	case 110:
		return name == "select"
	case 116:
		return name == "for"
	case 132:
		return name == "fallthrough"
	case 134:
		return name == "goto"
	case 142:
		return name == "const"
	case 154:
		return name == "import"
	case 164:
		return name == "break"
	case 180:
		return name == "func"
	case 190:
		return name == "continue"
	case 210:
		return name == "interface"
	case 230:
		return name == "struct"
	case 240:
		return name == "type"
	case 254:
		return name == "switch"
	}
	return false
}

func TwoHash_AddAdd_Shift500(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) + (name[0] << 0)) + (name[1] << 0) {
	case 15:
		return name == "if"
	case 22:
		return name == "go"
	case 39:
		return name == "fallthrough"
	case 46:
		return name == "map"
	case 53:
		return name == "for"
	case 55:
		return name == "var"
	case 68:
		return name == "case"
	case 75:
		return name == "chan"
	case 81:
		return name == "else"
	case 86:
		return name == "goto"
	case 91:
		return name == "func"
	case 105:
		return name == "defer"
	case 109:
		return name == "type"
	case 114:
		return name == "const"
	case 115:
		return name == "range"
	case 116:
		return name == "break"
	case 150:
		return name == "import"
	case 151:
		return name == "return"
	case 152:
		return name == "select"
	case 167:
		return name == "struct"
	case 169:
		return name == "default"
	case 170:
		return name == "switch"
	case 177:
		return name == "package"
	case 210:
		return name == "continue"
	case 247:
		return name == "interface"
	}
	return false
}

func TwoHash_AddAdd_Shift511(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 5) + (name[0] << 1)) + (name[1] << 1) {
	case 8:
		return name == "case"
	case 10:
		return name == "for"
	case 14:
		return name == "var"
	case 22:
		return name == "chan"
	case 34:
		return name == "else"
	case 44:
		return name == "goto"
	case 50:
		return name == "defer"
	case 54:
		return name == "func"
	case 68:
		return name == "const"
	case 70:
		return name == "range"
	case 72:
		return name == "break"
	case 90:
		return name == "type"
	case 108:
		return name == "import"
	case 110:
		return name == "return"
	case 112:
		return name == "select"
	case 114:
		return name == "default"
	case 130:
		return name == "package"
	case 142:
		return name == "struct"
	case 148:
		return name == "switch"
	case 164:
		return name == "continue"
	case 206:
		return name == "interface"
	case 222:
		return name == "if"
	case 236:
		return name == "go"
	case 238:
		return name == "fallthrough"
	case 252:
		return name == "map"
	}
	return false
}

func TwoHash_AddAdd_Shift601(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) + (name[0] << 0)) + (name[1] << 1) {
	case 4:
		return name == "for"
	case 37:
		return name == "case"
	case 51:
		return name == "chan"
	case 61:
		return name == "else"
	case 65:
		return name == "continue"
	case 69:
		return name == "goto"
	case 80:
		return name == "func"
	case 102:
		return name == "type"
	case 110:
		return name == "defer"
	case 116:
		return name == "range"
	case 129:
		return name == "const"
	case 133:
		return name == "interface"
	case 134:
		return name == "break"
	case 181:
		return name == "if"
	case 188:
		return name == "return"
	case 189:
		return name == "select"
	case 195:
		return name == "import"
	case 197:
		return name == "go"
	case 219:
		return name == "struct"
	case 225:
		return name == "switch"
	case 232:
		return name == "fallthrough"
	case 238:
		return name == "default"
	case 239:
		return name == "map"
	case 242:
		return name == "package"
	case 248:
		return name == "var"
	}
	return false
}

func TwoHash_AddAdd_Shift602(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) + (name[0] << 0)) + (name[1] << 2) {
	case 3:
		return name == "chan"
	case 21:
		return name == "else"
	case 31:
		return name == "continue"
	case 35:
		return name == "goto"
	case 54:
		return name == "range"
	case 56:
		return name == "defer"
	case 58:
		return name == "func"
	case 88:
		return name == "type"
	case 95:
		return name == "const"
	case 97:
		return name == "interface"
	case 106:
		return name == "break"
	case 129:
		return name == "if"
	case 134:
		return name == "return"
	case 135:
		return name == "select"
	case 157:
		return name == "import"
	case 163:
		return name == "go"
	case 170:
		return name == "fallthrough"
	case 177:
		return name == "map"
	case 180:
		return name == "package"
	case 184:
		return name == "default"
	case 186:
		return name == "var"
	case 195:
		return name == "struct"
	case 207:
		return name == "switch"
	case 226:
		return name == "for"
	case 231:
		return name == "case"
	}
	return false
}

func TwoHash_AddAdd_Shift620(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) + (name[0] << 2)) + (name[1] << 0) {
	case 0:
		return name == "else"
	case 11:
		return name == "goto"
	case 13:
		return name == "func"
	case 53:
		return name == "defer"
	case 58:
		return name == "break"
	case 59:
		return name == "const"
	case 73:
		return name == "type"
	case 82:
		return name == "interface"
	case 105:
		return name == "range"
	case 138:
		return name == "if"
	case 139:
		return name == "go"
	case 145:
		return name == "import"
	case 173:
		return name == "return"
	case 177:
		return name == "select"
	case 181:
		return name == "default"
	case 185:
		return name == "fallthrough"
	case 192:
		return name == "struct"
	case 195:
		return name == "switch"
	case 199:
		return name == "for"
	case 213:
		return name == "map"
	case 225:
		return name == "package"
	case 237:
		return name == "case"
	case 244:
		return name == "chan"
	case 249:
		return name == "var"
	case 251:
		return name == "continue"
	}
	return false
}

func TwoHash_AddAdd_Shift630(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 6) + (name[0] << 3)) + (name[1] << 0) {
	case 25:
		return name == "type"
	case 39:
		return name == "go"
	case 46:
		return name == "if"
	case 49:
		return name == "range"
	case 53:
		return name == "import"
	case 69:
		return name == "default"
	case 81:
		return name == "fallthrough"
	case 95:
		return name == "for"
	case 117:
		return name == "return"
	case 121:
		return name == "case"
	case 125:
		return name == "select"
	case 128:
		return name == "chan"
	case 135:
		return name == "continue"
	case 137:
		return name == "map"
	case 140:
		return name == "struct"
	case 143:
		return name == "switch"
	case 148:
		return name == "else"
	case 161:
		return name == "package"
	case 165:
		return name == "func"
	case 167:
		return name == "goto"
	case 194:
		return name == "break"
	case 197:
		return name == "defer"
	case 199:
		return name == "const"
	case 209:
		return name == "var"
	case 246:
		return name == "interface"
	}
	return false
}

func TwoHash_OrXor_Shift032(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 3)) ^ (name[1] << 2) {
	case 2:
		return name == "return"
	case 3:
		return name == "package"
	case 10:
		return name == "select"
	case 17:
		return name == "range"
	case 55:
		return name == "var"
	case 64:
		return name == "type"
	case 66:
		return name == "switch"
	case 78:
		return name == "struct"
	case 128:
		return name == "goto"
	case 134:
		return name == "go"
	case 143:
		return name == "for"
	case 152:
		return name == "case"
	case 156:
		return name == "else"
	case 161:
		return name == "const"
	case 164:
		return name == "continue"
	case 177:
		return name == "defer"
	case 179:
		return name == "default"
	case 188:
		return name == "chan"
	case 191:
		return name == "fallthrough"
	case 210:
		return name == "if"
	case 221:
		return name == "break"
	case 224:
		return name == "func"
	case 239:
		return name == "map"
	case 241:
		return name == "interface"
	case 250:
		return name == "import"
	}
	return false
}

func TwoHash_OrXor_Shift034(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 3)) ^ (name[1] << 4) {
	case 12:
		return name == "case"
	case 42:
		return name == "if"
	case 43:
		return name == "fallthrough"
	case 52:
		return name == "type"
	case 53:
		return name == "break"
	case 100:
		return name == "func"
	case 117:
		return name == "defer"
	case 119:
		return name == "default"
	case 123:
		return name == "map"
	case 133:
		return name == "range"
	case 151:
		return name == "package"
	case 156:
		return name == "chan"
	case 158:
		return name == "import"
	case 163:
		return name == "var"
	case 169:
		return name == "interface"
	case 195:
		return name == "for"
	case 198:
		return name == "return"
	case 202:
		return name == "go"
	case 204:
		return name == "goto"
	case 206:
		return name == "select"
	case 222:
		return name == "struct"
	case 232:
		return name == "continue"
	case 236:
		return name == "else"
	case 237:
		return name == "const"
	case 238:
		return name == "switch"
	}
	return false
}

func TwoHash_OrXor_Shift035(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 3)) ^ (name[1] << 5) {
	case 27:
		return name == "fallthrough"
	case 28:
		return name == "chan"
	case 30:
		return name == "struct"
	case 54:
		return name == "return"
	case 60:
		return name == "case"
	case 62:
		return name == "select"
	case 75:
		return name == "map"
	case 85:
		return name == "break"
	case 126:
		return name == "switch"
	case 132:
		return name == "type"
	case 133:
		return name == "defer"
	case 135:
		return name == "default"
	case 137:
		return name == "interface"
	case 138:
		return name == "if"
	case 147:
		return name == "var"
	case 148:
		return name == "func"
	case 167:
		return name == "package"
	case 172:
		return name == "else"
	case 181:
		return name == "range"
	case 211:
		return name == "for"
	case 218:
		return name == "go"
	case 220:
		return name == "goto"
	case 238:
		return name == "import"
	case 248:
		return name == "continue"
	case 253:
		return name == "const"
	}
	return false
}

func TwoHash_OrXor_Shift041(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 4)) ^ (name[1] << 1) {
	case 17:
		return name == "map"
	case 69:
		return name == "interface"
	case 76:
		return name == "import"
	case 94:
		return name == "if"
	case 140:
		return name == "else"
	case 141:
		return name == "default"
	case 142:
		return name == "func"
	case 143:
		return name == "defer"
	case 161:
		return name == "var"
	case 169:
		return name == "fallthrough"
	case 170:
		return name == "goto"
	case 172:
		return name == "go"
	case 182:
		return name == "type"
	case 189:
		return name == "for"
	case 193:
		return name == "break"
	case 197:
		return name == "package"
	case 216:
		return name == "switch"
	case 222:
		return name == "struct"
	case 228:
		return name == "chan"
	case 230:
		return name == "continue"
	case 231:
		return name == "range"
	case 235:
		return name == "const"
	case 236:
		return name == "return"
	case 246:
		return name == "case"
	case 252:
		return name == "select"
	}
	return false
}

func TwoHash_OrXor_Shift043(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 4)) ^ (name[1] << 3) {
	case 10:
		return name == "go"
	case 12:
		return name == "goto"
	case 14:
		return name == "return"
	case 15:
		return name == "package"
	case 27:
		return name == "for"
	case 30:
		return name == "select"
	case 45:
		return name == "range"
	case 52:
		return name == "else"
	case 60:
		return name == "case"
	case 64:
		return name == "continue"
	case 77:
		return name == "const"
	case 99:
		return name == "fallthrough"
	case 107:
		return name == "var"
	case 109:
		return name == "defer"
	case 111:
		return name == "default"
	case 116:
		return name == "chan"
	case 140:
		return name == "type"
	case 142:
		return name == "switch"
	case 150:
		return name == "struct"
	case 162:
		return name == "if"
	case 181:
		return name == "break"
	case 204:
		return name == "func"
	case 219:
		return name == "map"
	case 233:
		return name == "interface"
	case 254:
		return name == "import"
	}
	return false
}

func TwoHash_OrXor_Shift045(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 4)) ^ (name[1] << 5) {
	case 5:
		return name == "range"
	case 20:
		return name == "case"
	case 39:
		return name == "package"
	case 52:
		return name == "chan"
	case 54:
		return name == "import"
	case 67:
		return name == "var"
	case 75:
		return name == "fallthrough"
	case 82:
		return name == "if"
	case 89:
		return name == "interface"
	case 100:
		return name == "type"
	case 101:
		return name == "break"
	case 131:
		return name == "for"
	case 134:
		return name == "return"
	case 146:
		return name == "go"
	case 148:
		return name == "goto"
	case 150:
		return name == "select"
	case 182:
		return name == "struct"
	case 196:
		return name == "func"
	case 212:
		return name == "else"
	case 213:
		return name == "const"
	case 214:
		return name == "switch"
	case 216:
		return name == "continue"
	case 229:
		return name == "defer"
	case 231:
		return name == "default"
	case 243:
		return name == "map"
	}
	return false
}

func TwoHash_OrXor_Shift046(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 4)) ^ (name[1] << 6) {
	case 4:
		return name == "type"
	case 5:
		return name == "defer"
	case 7:
		return name == "default"
	case 18:
		return name == "if"
	case 25:
		return name == "interface"
	case 35:
		return name == "var"
	case 36:
		return name == "func"
	case 43:
		return name == "fallthrough"
	case 52:
		return name == "chan"
	case 54:
		return name == "struct"
	case 71:
		return name == "package"
	case 84:
		return name == "else"
	case 101:
		return name == "range"
	case 102:
		return name == "return"
	case 116:
		return name == "case"
	case 118:
		return name == "select"
	case 147:
		return name == "map"
	case 163:
		return name == "for"
	case 165:
		return name == "break"
	case 178:
		return name == "go"
	case 180:
		return name == "goto"
	case 214:
		return name == "import"
	case 245:
		return name == "const"
	case 246:
		return name == "switch"
	case 248:
		return name == "continue"
	}
	return false
}

func TwoHash_OrXor_Shift050(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 5)) ^ (name[1] << 0) {
	case 3:
		return name == "select"
	case 5:
		return name == "case"
	case 7:
		return name == "continue"
	case 10:
		return name == "const"
	case 12:
		return name == "chan"
	case 17:
		return name == "switch"
	case 18:
		return name == "struct"
	case 35:
		return name == "return"
	case 36:
		return name == "range"
	case 55:
		return name == "break"
	case 68:
		return name == "if"
	case 71:
		return name == "interface"
	case 75:
		return name == "import"
	case 102:
		return name == "package"
	case 139:
		return name == "goto"
	case 141:
		return name == "go"
	case 162:
		return name == "var"
	case 170:
		return name == "fallthrough"
	case 172:
		return name == "for"
	case 177:
		return name == "func"
	case 194:
		return name == "map"
	case 200:
		return name == "else"
	case 224:
		return name == "defer"
	case 226:
		return name == "default"
	case 253:
		return name == "type"
	}
	return false
}

func TwoHash_OrXor_Shift051(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 5)) ^ (name[1] << 1) {
	case 1:
		return name == "var"
	case 9:
		return name == "fallthrough"
	case 29:
		return name == "for"
	case 46:
		return name == "func"
	case 58:
		return name == "goto"
	case 60:
		return name == "go"
	case 77:
		return name == "default"
	case 79:
		return name == "defer"
	case 97:
		return name == "map"
	case 118:
		return name == "type"
	case 124:
		return name == "else"
	case 135:
		return name == "range"
	case 136:
		return name == "switch"
	case 140:
		return name == "return"
	case 142:
		return name == "struct"
	case 161:
		return name == "break"
	case 166:
		return name == "case"
	case 172:
		return name == "select"
	case 180:
		return name == "chan"
	case 182:
		return name == "continue"
	case 187:
		return name == "const"
	case 197:
		return name == "package"
	case 238:
		return name == "if"
	case 245:
		return name == "interface"
	case 252:
		return name == "import"
	}
	return false
}

func TwoHash_OrXor_Shift060(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 6)) ^ (name[1] << 0) {
	case 34:
		return name == "map"
	case 36:
		return name == "if"
	case 39:
		return name == "interface"
	case 40:
		return name == "else"
	case 43:
		return name == "import"
	case 96:
		return name == "defer"
	case 98:
		return name == "default"
	case 102:
		return name == "package"
	case 125:
		return name == "type"
	case 163:
		return name == "select"
	case 165:
		return name == "case"
	case 167:
		return name == "continue"
	case 170:
		return name == "const"
	case 171:
		return name == "goto"
	case 172:
		return name == "chan"
	case 173:
		return name == "go"
	case 177:
		return name == "switch"
	case 178:
		return name == "struct"
	case 226:
		return name == "var"
	case 227:
		return name == "return"
	case 228:
		return name == "range"
	case 234:
		return name == "fallthrough"
	case 236:
		return name == "for"
	case 241:
		return name == "func"
	case 247:
		return name == "break"
	}
	return false
}

func TwoHash_OrXor_Shift062(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 6)) ^ (name[1] << 2) {
	case 1:
		return name == "range"
	case 7:
		return name == "var"
	case 15:
		return name == "fallthrough"
	case 18:
		return name == "return"
	case 22:
		return name == "struct"
	case 26:
		return name == "switch"
	case 63:
		return name == "for"
	case 64:
		return name == "case"
	case 77:
		return name == "break"
	case 80:
		return name == "func"
	case 82:
		return name == "select"
	case 100:
		return name == "chan"
	case 116:
		return name == "continue"
	case 120:
		return name == "goto"
	case 121:
		return name == "const"
	case 126:
		return name == "go"
	case 131:
		return name == "package"
	case 145:
		return name == "defer"
	case 147:
		return name == "default"
	case 199:
		return name == "map"
	case 218:
		return name == "if"
	case 224:
		return name == "type"
	case 241:
		return name == "interface"
	case 242:
		return name == "import"
	case 244:
		return name == "else"
	}
	return false
}

func TwoHash_OrXor_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 7)) ^ (name[1] << 3) {
	case 3:
		return name == "fallthrough"
	case 11:
		return name == "var"
	case 13:
		return name == "range"
	case 15:
		return name == "package"
	case 38:
		return name == "struct"
	case 45:
		return name == "defer"
	case 46:
		return name == "return"
	case 47:
		return name == "default"
	case 62:
		return name == "switch"
	case 123:
		return name == "for"
	case 139:
		return name == "map"
	case 140:
		return name == "case"
	case 149:
		return name == "break"
	case 172:
		return name == "func"
	case 174:
		return name == "select"
	case 178:
		return name == "if"
	case 196:
		return name == "chan"
	case 204:
		return name == "type"
	case 228:
		return name == "else"
	case 238:
		return name == "import"
	case 240:
		return name == "continue"
	case 249:
		return name == "interface"
	case 250:
		return name == "go"
	case 252:
		return name == "goto"
	case 253:
		return name == "const"
	}
	return false
}

func TwoHash_OrXor_Shift143(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) | (name[0] << 4)) ^ (name[1] << 3) {
	case 0:
		return name == "goto"
	case 4:
		return name == "return"
	case 6:
		return name == "package"
	case 12:
		return name == "go"
	case 20:
		return name == "select"
	case 30:
		return name == "for"
	case 34:
		return name == "range"
	case 48:
		return name == "case"
	case 56:
		return name == "else"
	case 66:
		return name == "const"
	case 72:
		return name == "continue"
	case 98:
		return name == "defer"
	case 102:
		return name == "default"
	case 110:
		return name == "var"
	case 120:
		return name == "chan"
	case 126:
		return name == "fallthrough"
	case 128:
		return name == "type"
	case 132:
		return name == "switch"
	case 156:
		return name == "struct"
	case 164:
		return name == "if"
	case 186:
		return name == "break"
	case 192:
		return name == "func"
	case 222:
		return name == "map"
	case 226:
		return name == "interface"
	case 244:
		return name == "import"
	}
	return false
}

func TwoHash_OrXor_Shift145(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) | (name[0] << 4)) ^ (name[1] << 5) {
	case 10:
		return name == "range"
	case 24:
		return name == "case"
	case 46:
		return name == "package"
	case 56:
		return name == "chan"
	case 60:
		return name == "import"
	case 70:
		return name == "var"
	case 82:
		return name == "interface"
	case 84:
		return name == "if"
	case 86:
		return name == "fallthrough"
	case 104:
		return name == "type"
	case 106:
		return name == "break"
	case 134:
		return name == "for"
	case 140:
		return name == "return"
	case 148:
		return name == "go"
	case 152:
		return name == "goto"
	case 156:
		return name == "select"
	case 188:
		return name == "struct"
	case 200:
		return name == "func"
	case 208:
		return name == "continue"
	case 216:
		return name == "else"
	case 218:
		return name == "const"
	case 220:
		return name == "switch"
	case 234:
		return name == "defer"
	case 238:
		return name == "default"
	case 246:
		return name == "map"
	}
	return false
}

func TwoHash_OrXor_Shift146(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) | (name[0] << 4)) ^ (name[1] << 6) {
	case 8:
		return name == "type"
	case 10:
		return name == "defer"
	case 14:
		return name == "default"
	case 18:
		return name == "interface"
	case 20:
		return name == "if"
	case 38:
		return name == "var"
	case 40:
		return name == "func"
	case 54:
		return name == "fallthrough"
	case 56:
		return name == "chan"
	case 60:
		return name == "struct"
	case 78:
		return name == "package"
	case 88:
		return name == "else"
	case 106:
		return name == "range"
	case 108:
		return name == "return"
	case 120:
		return name == "case"
	case 124:
		return name == "select"
	case 150:
		return name == "map"
	case 166:
		return name == "for"
	case 170:
		return name == "break"
	case 180:
		return name == "go"
	case 184:
		return name == "goto"
	case 220:
		return name == "import"
	case 240:
		return name == "continue"
	case 250:
		return name == "const"
	case 252:
		return name == "switch"
	}
	return false
}

func TwoHash_OrXor_Shift161(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) | (name[0] << 6)) ^ (name[1] << 1) {
	case 6:
		return name == "select"
	case 10:
		return name == "case"
	case 14:
		return name == "continue"
	case 20:
		return name == "const"
	case 22:
		return name == "goto"
	case 24:
		return name == "chan"
	case 26:
		return name == "go"
	case 34:
		return name == "switch"
	case 36:
		return name == "struct"
	case 68:
		return name == "var"
	case 70:
		return name == "return"
	case 72:
		return name == "range"
	case 84:
		return name == "fallthrough"
	case 88:
		return name == "for"
	case 98:
		return name == "func"
	case 110:
		return name == "break"
	case 132:
		return name == "map"
	case 136:
		return name == "if"
	case 142:
		return name == "interface"
	case 144:
		return name == "else"
	case 150:
		return name == "import"
	case 192:
		return name == "defer"
	case 196:
		return name == "default"
	case 204:
		return name == "package"
	case 250:
		return name == "type"
	}
	return false
}

func TwoHash_OrXor_Shift260(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) | (name[0] << 6)) ^ (name[1] << 0) {
	case 10:
		return name == "interface"
	case 45:
		return name == "map"
	case 46:
		return name == "if"
	case 53:
		return name == "import"
	case 60:
		return name == "else"
	case 105:
		return name == "type"
	case 113:
		return name == "defer"
	case 121:
		return name == "default"
	case 125:
		return name == "package"
	case 143:
		return name == "continue"
	case 167:
		return name == "go"
	case 172:
		return name == "struct"
	case 175:
		return name == "switch"
	case 177:
		return name == "case"
	case 184:
		return name == "chan"
	case 187:
		return name == "const"
	case 189:
		return name == "select"
	case 191:
		return name == "goto"
	case 205:
		return name == "fallthrough"
	case 227:
		return name == "for"
	case 229:
		return name == "func"
	case 230:
		return name == "break"
	case 237:
		return name == "var"
	case 245:
		return name == "range"
	case 253:
		return name == "return"
	}
	return false
}

func TwoHash_OrAdd_Shift031(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 3)) + (name[1] << 1) {
	case 4:
		return name == "else"
	case 17:
		return name == "for"
	case 22:
		return name == "if"
	case 24:
		return name == "go"
	case 26:
		return name == "goto"
	case 30:
		return name == "func"
	case 37:
		return name == "interface"
	case 40:
		return name == "import"
	case 45:
		return name == "map"
	case 73:
		return name == "package"
	case 87:
		return name == "range"
	case 96:
		return name == "return"
	case 104:
		return name == "select"
	case 117:
		return name == "var"
	case 134:
		return name == "struct"
	case 140:
		return name == "switch"
	case 150:
		return name == "type"
	case 222:
		return name == "case"
	case 236:
		return name == "chan"
	case 239:
		return name == "defer"
	case 241:
		return name == "default"
	case 246:
		return name == "continue"
	case 249:
		return name == "break"
	case 251:
		return name == "const"
	case 253:
		return name == "fallthrough"
	}
	return false
}

func TwoHash_OrAdd_Shift033(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 3)) + (name[1] << 3) {
	case 36:
		return name == "case"
	case 62:
		return name == "struct"
	case 67:
		return name == "fallthrough"
	case 77:
		return name == "defer"
	case 79:
		return name == "default"
	case 86:
		return name == "switch"
	case 92:
		return name == "chan"
	case 108:
		return name == "type"
	case 115:
		return name == "map"
	case 122:
		return name == "if"
	case 140:
		return name == "else"
	case 143:
		return name == "package"
	case 144:
		return name == "continue"
	case 149:
		return name == "const"
	case 157:
		return name == "range"
	case 165:
		return name == "break"
	case 171:
		return name == "for"
	case 178:
		return name == "go"
	case 180:
		return name == "goto"
	case 182:
		return name == "import"
	case 185:
		return name == "interface"
	case 187:
		return name == "var"
	case 190:
		return name == "return"
	case 198:
		return name == "select"
	case 220:
		return name == "func"
	}
	return false
}

func TwoHash_OrAdd_Shift041(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 4)) + (name[1] << 1) {
	case 0:
		return name == "select"
	case 4:
		return name == "chan"
	case 9:
		return name == "break"
	case 15:
		return name == "defer"
	case 17:
		return name == "default"
	case 19:
		return name == "const"
	case 22:
		return name == "continue"
	case 30:
		return name == "struct"
	case 36:
		return name == "switch"
	case 37:
		return name == "var"
	case 44:
		return name == "else"
	case 45:
		return name == "fallthrough"
	case 54:
		return name == "type"
	case 65:
		return name == "for"
	case 78:
		return name == "func"
	case 80:
		return name == "go"
	case 82:
		return name == "goto"
	case 94:
		return name == "if"
	case 112:
		return name == "import"
	case 117:
		return name == "interface"
	case 149:
		return name == "map"
	case 201:
		return name == "package"
	case 231:
		return name == "range"
	case 240:
		return name == "return"
	case 246:
		return name == "case"
	}
	return false
}

func TwoHash_OrAdd_Shift042(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 4)) + (name[1] << 2) {
	case 4:
		return name == "else"
	case 6:
		return name == "struct"
	case 18:
		return name == "switch"
	case 31:
		return name == "for"
	case 40:
		return name == "type"
	case 42:
		return name == "if"
	case 46:
		return name == "go"
	case 48:
		return name == "goto"
	case 56:
		return name == "func"
	case 74:
		return name == "import"
	case 81:
		return name == "interface"
	case 87:
		return name == "map"
	case 139:
		return name == "package"
	case 169:
		return name == "range"
	case 184:
		return name == "case"
	case 186:
		return name == "return"
	case 202:
		return name == "select"
	case 212:
		return name == "chan"
	case 217:
		return name == "defer"
	case 219:
		return name == "default"
	case 231:
		return name == "var"
	case 237:
		return name == "break"
	case 239:
		return name == "fallthrough"
	case 241:
		return name == "const"
	case 244:
		return name == "continue"
	}
	return false
}

func TwoHash_OrAdd_Shift050(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 5)) + (name[1] << 0) {
	case 4:
		return name == "map"
	case 16:
		return name == "else"
	case 36:
		return name == "var"
	case 44:
		return name == "fallthrough"
	case 50:
		return name == "for"
	case 57:
		return name == "func"
	case 81:
		return name == "go"
	case 83:
		return name == "goto"
	case 104:
		return name == "package"
	case 136:
		return name == "if"
	case 147:
		return name == "import"
	case 151:
		return name == "interface"
	case 166:
		return name == "range"
	case 171:
		return name == "return"
	case 183:
		return name == "break"
	case 197:
		return name == "case"
	case 203:
		return name == "select"
	case 204:
		return name == "chan"
	case 212:
		return name == "const"
	case 215:
		return name == "continue"
	case 218:
		return name == "struct"
	case 221:
		return name == "switch"
	case 234:
		return name == "defer"
	case 236:
		return name == "default"
	case 253:
		return name == "type"
	}
	return false
}

func TwoHash_OrAdd_Shift051(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 5)) + (name[1] << 1) {
	case 0:
		return name == "import"
	case 5:
		return name == "interface"
	case 7:
		return name == "range"
	case 16:
		return name == "return"
	case 38:
		return name == "case"
	case 41:
		return name == "break"
	case 48:
		return name == "select"
	case 52:
		return name == "chan"
	case 67:
		return name == "const"
	case 70:
		return name == "continue"
	case 78:
		return name == "struct"
	case 79:
		return name == "defer"
	case 81:
		return name == "default"
	case 84:
		return name == "switch"
	case 101:
		return name == "map"
	case 118:
		return name == "type"
	case 124:
		return name == "else"
	case 133:
		return name == "var"
	case 141:
		return name == "fallthrough"
	case 161:
		return name == "for"
	case 174:
		return name == "func"
	case 192:
		return name == "go"
	case 194:
		return name == "goto"
	case 201:
		return name == "package"
	case 238:
		return name == "if"
	}
	return false
}

func TwoHash_OrAdd_Shift060(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 6)) + (name[1] << 0) {
	case 37:
		return name == "case"
	case 43:
		return name == "select"
	case 44:
		return name == "chan"
	case 49:
		return name == "go"
	case 51:
		return name == "goto"
	case 52:
		return name == "const"
	case 55:
		return name == "continue"
	case 58:
		return name == "struct"
	case 61:
		return name == "switch"
	case 104:
		return name == "package"
	case 106:
		return name == "defer"
	case 108:
		return name == "default"
	case 125:
		return name == "type"
	case 164:
		return name == "map"
	case 168:
		return name == "if"
	case 176:
		return name == "else"
	case 179:
		return name == "import"
	case 183:
		return name == "interface"
	case 228:
		return name == "var"
	case 230:
		return name == "range"
	case 235:
		return name == "return"
	case 236:
		return name == "fallthrough"
	case 242:
		return name == "for"
	case 247:
		return name == "break"
	case 249:
		return name == "func"
	}
	return false
}

func TwoHash_OrAdd_Shift061(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 6)) + (name[1] << 1) {
	case 5:
		return name == "map"
	case 14:
		return name == "if"
	case 28:
		return name == "else"
	case 32:
		return name == "import"
	case 37:
		return name == "interface"
	case 69:
		return name == "var"
	case 71:
		return name == "range"
	case 77:
		return name == "fallthrough"
	case 80:
		return name == "return"
	case 97:
		return name == "for"
	case 105:
		return name == "break"
	case 110:
		return name == "func"
	case 134:
		return name == "case"
	case 144:
		return name == "select"
	case 148:
		return name == "chan"
	case 160:
		return name == "go"
	case 162:
		return name == "goto"
	case 163:
		return name == "const"
	case 166:
		return name == "continue"
	case 174:
		return name == "struct"
	case 180:
		return name == "switch"
	case 201:
		return name == "package"
	case 207:
		return name == "defer"
	case 209:
		return name == "default"
	case 246:
		return name == "type"
	}
	return false
}

func TwoHash_OrAdd_Shift062(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 6)) + (name[1] << 2) {
	case 1:
		return name == "interface"
	case 7:
		return name == "var"
	case 9:
		return name == "range"
	case 15:
		return name == "fallthrough"
	case 26:
		return name == "return"
	case 63:
		return name == "for"
	case 72:
		return name == "case"
	case 77:
		return name == "break"
	case 88:
		return name == "func"
	case 90:
		return name == "select"
	case 100:
		return name == "chan"
	case 126:
		return name == "go"
	case 128:
		return name == "goto"
	case 129:
		return name == "const"
	case 132:
		return name == "continue"
	case 139:
		return name == "package"
	case 150:
		return name == "struct"
	case 153:
		return name == "defer"
	case 155:
		return name == "default"
	case 162:
		return name == "switch"
	case 199:
		return name == "map"
	case 218:
		return name == "if"
	case 232:
		return name == "type"
	case 244:
		return name == "else"
	case 250:
		return name == "import"
	}
	return false
}

func TwoHash_OrAdd_Shift073(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 0) | (name[0] << 7)) + (name[1] << 3) {
	case 0:
		return name == "continue"
	case 11:
		return name == "var"
	case 13:
		return name == "range"
	case 15:
		return name == "package"
	case 19:
		return name == "fallthrough"
	case 38:
		return name == "struct"
	case 45:
		return name == "defer"
	case 46:
		return name == "return"
	case 47:
		return name == "default"
	case 62:
		return name == "switch"
	case 123:
		return name == "for"
	case 139:
		return name == "map"
	case 140:
		return name == "case"
	case 149:
		return name == "break"
	case 172:
		return name == "func"
	case 174:
		return name == "select"
	case 178:
		return name == "if"
	case 196:
		return name == "chan"
	case 204:
		return name == "type"
	case 228:
		return name == "else"
	case 238:
		return name == "import"
	case 249:
		return name == "interface"
	case 250:
		return name == "go"
	case 252:
		return name == "goto"
	case 253:
		return name == "const"
	}
	return false
}

func TwoHash_OrAdd_Shift150(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) | (name[0] << 5)) + (name[1] << 0) {
	case 1:
		return name == "type"
	case 7:
		return name == "map"
	case 20:
		return name == "else"
	case 39:
		return name == "var"
	case 53:
		return name == "for"
	case 55:
		return name == "fallthrough"
	case 61:
		return name == "func"
	case 83:
		return name == "go"
	case 87:
		return name == "goto"
	case 111:
		return name == "package"
	case 138:
		return name == "if"
	case 153:
		return name == "import"
	case 160:
		return name == "interface"
	case 171:
		return name == "range"
	case 177:
		return name == "return"
	case 188:
		return name == "break"
	case 201:
		return name == "case"
	case 208:
		return name == "chan"
	case 209:
		return name == "select"
	case 217:
		return name == "const"
	case 223:
		return name == "continue"
	case 224:
		return name == "struct"
	case 227:
		return name == "switch"
	case 239:
		return name == "defer"
	case 243:
		return name == "default"
	}
	return false
}

func TwoHash_OrAdd_Shift161(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 1) | (name[0] << 6)) + (name[1] << 1) {
	case 8:
		return name == "map"
	case 16:
		return name == "if"
	case 32:
		return name == "else"
	case 38:
		return name == "import"
	case 46:
		return name == "interface"
	case 72:
		return name == "var"
	case 76:
		return name == "range"
	case 86:
		return name == "return"
	case 88:
		return name == "fallthrough"
	case 100:
		return name == "for"
	case 110:
		return name == "break"
	case 114:
		return name == "func"
	case 138:
		return name == "case"
	case 150:
		return name == "select"
	case 152:
		return name == "chan"
	case 162:
		return name == "go"
	case 166:
		return name == "goto"
	case 168:
		return name == "const"
	case 174:
		return name == "continue"
	case 180:
		return name == "struct"
	case 186:
		return name == "switch"
	case 208:
		return name == "package"
	case 212:
		return name == "defer"
	case 216:
		return name == "default"
	case 250:
		return name == "type"
	}
	return false
}

func TwoHash_OrAdd_Shift250(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 2) | (name[0] << 5)) + (name[1] << 0) {
	case 1:
		return name == "default"
	case 9:
		return name == "type"
	case 13:
		return name == "map"
	case 28:
		return name == "else"
	case 45:
		return name == "var"
	case 59:
		return name == "for"
	case 69:
		return name == "func"
	case 77:
		return name == "fallthrough"
	case 87:
		return name == "go"
	case 95:
		return name == "goto"
	case 125:
		return name == "package"
	case 142:
		return name == "if"
	case 146:
		return name == "interface"
	case 165:
		return name == "import"
	case 181:
		return name == "range"
	case 189:
		return name == "return"
	case 198:
		return name == "break"
	case 207:
		return name == "continue"
	case 209:
		return name == "case"
	case 216:
		return name == "chan"
	case 221:
		return name == "select"
	case 227:
		return name == "const"
	case 236:
		return name == "struct"
	case 239:
		return name == "switch"
	case 249:
		return name == "defer"
	}
	return false
}

func TwoHash_OrAdd_Shift361(name string) bool {
	if len(name) < 2 {
		return false
	}
	switch ((byte(len(name)) << 3) | (name[0] << 6)) + (name[1] << 1) {
	case 2:
		return name == "default"
	case 18:
		return name == "type"
	case 26:
		return name == "map"
	case 28:
		return name == "if"
	case 36:
		return name == "interface"
	case 56:
		return name == "else"
	case 74:
		return name == "import"
	case 90:
		return name == "var"
	case 106:
		return name == "range"
	case 118:
		return name == "for"
	case 122:
		return name == "return"
	case 138:
		return name == "func"
	case 140:
		return name == "break"
	case 154:
		return name == "fallthrough"
	case 158:
		return name == "continue"
	case 162:
		return name == "case"
	case 174:
		return name == "go"
	case 176:
		return name == "chan"
	case 186:
		return name == "select"
	case 190:
		return name == "goto"
	case 198:
		return name == "const"
	case 216:
		return name == "struct"
	case 222:
		return name == "switch"
	case 242:
		return name == "defer"
	case 250:
		return name == "package"
	}
	return false
}

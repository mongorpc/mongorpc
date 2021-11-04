// Code generated from MongoRPCRules.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // MongoRPCRules

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 55, 290,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 64, 10, 3, 12, 3, 14, 3, 67,
	11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 73, 10, 4, 12, 4, 14, 4, 76, 11, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 85, 10, 5, 12, 5, 14, 5,
	88, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 104, 10, 9, 12, 9, 14, 9, 107, 11, 9, 3,
	9, 3, 9, 3, 9, 5, 9, 112, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 5, 11, 122, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 5,
	13, 129, 10, 13, 3, 13, 3, 13, 7, 13, 133, 10, 13, 12, 13, 14, 13, 136,
	11, 13, 3, 14, 3, 14, 3, 15, 5, 15, 141, 10, 15, 3, 15, 3, 15, 7, 15, 145,
	10, 15, 12, 15, 14, 15, 148, 11, 15, 3, 16, 3, 16, 3, 17, 5, 17, 153, 10,
	17, 3, 17, 3, 17, 7, 17, 157, 10, 17, 12, 17, 14, 17, 160, 11, 17, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 179, 10, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 5,
	22, 193, 10, 22, 3, 22, 3, 22, 7, 22, 197, 10, 22, 12, 22, 14, 22, 200,
	11, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 215, 10, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 247, 10, 22, 3, 22, 3, 22,
	7, 22, 251, 10, 22, 12, 22, 14, 22, 254, 11, 22, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 5, 25, 265, 10, 25, 6, 25, 267,
	10, 25, 13, 25, 14, 25, 268, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 5, 28, 284, 10, 28, 6,
	28, 286, 10, 28, 13, 28, 14, 28, 287, 3, 28, 2, 3, 42, 29, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
	46, 48, 50, 52, 54, 2, 11, 4, 2, 36, 36, 40, 45, 5, 2, 33, 35, 37, 49,
	52, 52, 4, 2, 24, 24, 28, 28, 3, 2, 38, 39, 3, 2, 16, 21, 3, 2, 22, 23,
	3, 2, 25, 26, 4, 2, 27, 31, 53, 53, 3, 2, 36, 37, 2, 304, 2, 56, 3, 2,
	2, 2, 4, 60, 3, 2, 2, 2, 6, 68, 3, 2, 2, 2, 8, 79, 3, 2, 2, 2, 10, 91,
	3, 2, 2, 2, 12, 93, 3, 2, 2, 2, 14, 95, 3, 2, 2, 2, 16, 99, 3, 2, 2, 2,
	18, 115, 3, 2, 2, 2, 20, 117, 3, 2, 2, 2, 22, 125, 3, 2, 2, 2, 24, 128,
	3, 2, 2, 2, 26, 137, 3, 2, 2, 2, 28, 140, 3, 2, 2, 2, 30, 149, 3, 2, 2,
	2, 32, 152, 3, 2, 2, 2, 34, 161, 3, 2, 2, 2, 36, 178, 3, 2, 2, 2, 38, 180,
	3, 2, 2, 2, 40, 186, 3, 2, 2, 2, 42, 214, 3, 2, 2, 2, 44, 255, 3, 2, 2,
	2, 46, 257, 3, 2, 2, 2, 48, 266, 3, 2, 2, 2, 50, 270, 3, 2, 2, 2, 52, 275,
	3, 2, 2, 2, 54, 285, 3, 2, 2, 2, 56, 57, 7, 49, 2, 2, 57, 58, 5, 4, 3,
	2, 58, 59, 5, 6, 4, 2, 59, 3, 3, 2, 2, 2, 60, 65, 5, 40, 21, 2, 61, 62,
	7, 12, 2, 2, 62, 64, 5, 40, 21, 2, 63, 61, 3, 2, 2, 2, 64, 67, 3, 2, 2,
	2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 5, 3, 2, 2, 2, 67, 65, 3,
	2, 2, 2, 68, 74, 7, 5, 2, 2, 69, 73, 5, 14, 8, 2, 70, 73, 5, 12, 7, 2,
	71, 73, 5, 34, 18, 2, 72, 69, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 71, 3,
	2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75,
	77, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 78, 7, 6, 2, 2, 78, 7, 3, 2, 2,
	2, 79, 86, 7, 5, 2, 2, 80, 85, 5, 16, 9, 2, 81, 85, 5, 14, 8, 2, 82, 85,
	7, 55, 2, 2, 83, 85, 5, 34, 18, 2, 84, 80, 3, 2, 2, 2, 84, 81, 3, 2, 2,
	2, 84, 82, 3, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84,
	3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2,
	89, 90, 7, 6, 2, 2, 90, 9, 3, 2, 2, 2, 91, 92, 9, 2, 2, 2, 92, 11, 3, 2,
	2, 2, 93, 94, 7, 55, 2, 2, 94, 13, 3, 2, 2, 2, 95, 96, 7, 34, 2, 2, 96,
	97, 5, 54, 28, 2, 97, 98, 5, 8, 5, 2, 98, 15, 3, 2, 2, 2, 99, 100, 7, 33,
	2, 2, 100, 105, 5, 10, 6, 2, 101, 102, 7, 14, 2, 2, 102, 104, 5, 10, 6,
	2, 103, 101, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105,
	106, 3, 2, 2, 2, 106, 111, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 109,
	7, 13, 2, 2, 109, 110, 7, 35, 2, 2, 110, 112, 5, 42, 22, 2, 111, 108, 3,
	2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 7, 15, 2,
	2, 114, 17, 3, 2, 2, 2, 115, 116, 7, 52, 2, 2, 116, 19, 3, 2, 2, 2, 117,
	118, 7, 5, 2, 2, 118, 121, 7, 52, 2, 2, 119, 120, 7, 3, 2, 2, 120, 122,
	7, 4, 2, 2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 123, 3, 2,
	2, 2, 123, 124, 7, 6, 2, 2, 124, 21, 3, 2, 2, 2, 125, 126, 5, 42, 22, 2,
	126, 23, 3, 2, 2, 2, 127, 129, 5, 22, 12, 2, 128, 127, 3, 2, 2, 2, 128,
	129, 3, 2, 2, 2, 129, 134, 3, 2, 2, 2, 130, 131, 7, 14, 2, 2, 131, 133,
	5, 22, 12, 2, 132, 130, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132, 3,
	2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 25, 3, 2, 2, 2, 136, 134, 3, 2, 2,
	2, 137, 138, 5, 42, 22, 2, 138, 27, 3, 2, 2, 2, 139, 141, 5, 26, 14, 2,
	140, 139, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 146, 3, 2, 2, 2, 142,
	143, 7, 14, 2, 2, 143, 145, 5, 26, 14, 2, 144, 142, 3, 2, 2, 2, 145, 148,
	3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 29, 3, 2,
	2, 2, 148, 146, 3, 2, 2, 2, 149, 150, 7, 52, 2, 2, 150, 31, 3, 2, 2, 2,
	151, 153, 5, 30, 16, 2, 152, 151, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153,
	158, 3, 2, 2, 2, 154, 155, 7, 14, 2, 2, 155, 157, 5, 30, 16, 2, 156, 154,
	3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 159, 3, 2,
	2, 2, 159, 33, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161, 162, 7, 46, 2, 2,
	162, 163, 7, 52, 2, 2, 163, 164, 7, 7, 2, 2, 164, 165, 5, 32, 17, 2, 165,
	166, 7, 8, 2, 2, 166, 167, 7, 5, 2, 2, 167, 168, 7, 47, 2, 2, 168, 169,
	5, 42, 22, 2, 169, 170, 7, 15, 2, 2, 170, 171, 7, 6, 2, 2, 171, 35, 3,
	2, 2, 2, 172, 173, 7, 12, 2, 2, 173, 179, 5, 40, 21, 2, 174, 175, 7, 10,
	2, 2, 175, 176, 5, 42, 22, 2, 176, 177, 7, 11, 2, 2, 177, 179, 3, 2, 2,
	2, 178, 172, 3, 2, 2, 2, 178, 174, 3, 2, 2, 2, 179, 37, 3, 2, 2, 2, 180,
	181, 7, 12, 2, 2, 181, 182, 5, 40, 21, 2, 182, 183, 7, 7, 2, 2, 183, 184,
	5, 24, 13, 2, 184, 185, 7, 8, 2, 2, 185, 39, 3, 2, 2, 2, 186, 187, 9, 3,
	2, 2, 187, 41, 3, 2, 2, 2, 188, 189, 8, 22, 1, 2, 189, 215, 7, 48, 2, 2,
	190, 192, 7, 10, 2, 2, 191, 193, 5, 42, 22, 2, 192, 191, 3, 2, 2, 2, 192,
	193, 3, 2, 2, 2, 193, 198, 3, 2, 2, 2, 194, 195, 7, 14, 2, 2, 195, 197,
	5, 42, 22, 2, 196, 194, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2, 198, 196, 3,
	2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 201, 3, 2, 2, 2, 200, 198, 3, 2, 2,
	2, 201, 215, 7, 11, 2, 2, 202, 203, 9, 4, 2, 2, 203, 215, 5, 42, 22, 10,
	204, 215, 7, 51, 2, 2, 205, 215, 7, 50, 2, 2, 206, 215, 9, 5, 2, 2, 207,
	215, 7, 52, 2, 2, 208, 215, 5, 50, 26, 2, 209, 215, 5, 52, 27, 2, 210,
	211, 7, 7, 2, 2, 211, 212, 5, 42, 22, 2, 212, 213, 7, 8, 2, 2, 213, 215,
	3, 2, 2, 2, 214, 188, 3, 2, 2, 2, 214, 190, 3, 2, 2, 2, 214, 202, 3, 2,
	2, 2, 214, 204, 3, 2, 2, 2, 214, 205, 3, 2, 2, 2, 214, 206, 3, 2, 2, 2,
	214, 207, 3, 2, 2, 2, 214, 208, 3, 2, 2, 2, 214, 209, 3, 2, 2, 2, 214,
	210, 3, 2, 2, 2, 215, 252, 3, 2, 2, 2, 216, 217, 12, 19, 2, 2, 217, 218,
	9, 6, 2, 2, 218, 251, 5, 42, 22, 20, 219, 220, 12, 18, 2, 2, 220, 221,
	9, 7, 2, 2, 221, 251, 5, 42, 22, 19, 222, 223, 12, 17, 2, 2, 223, 224,
	9, 8, 2, 2, 224, 251, 5, 42, 22, 18, 225, 226, 12, 16, 2, 2, 226, 227,
	9, 9, 2, 2, 227, 251, 5, 42, 22, 17, 228, 229, 12, 15, 2, 2, 229, 230,
	7, 32, 2, 2, 230, 251, 5, 42, 22, 16, 231, 232, 12, 14, 2, 2, 232, 233,
	7, 12, 2, 2, 233, 251, 5, 40, 21, 2, 234, 235, 12, 13, 2, 2, 235, 236,
	7, 12, 2, 2, 236, 237, 5, 40, 21, 2, 237, 238, 7, 7, 2, 2, 238, 239, 5,
	28, 15, 2, 239, 240, 7, 8, 2, 2, 240, 251, 3, 2, 2, 2, 241, 242, 12, 12,
	2, 2, 242, 243, 7, 10, 2, 2, 243, 246, 5, 42, 22, 2, 244, 245, 7, 13, 2,
	2, 245, 247, 5, 42, 22, 2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2,
	247, 248, 3, 2, 2, 2, 248, 249, 7, 11, 2, 2, 249, 251, 3, 2, 2, 2, 250,
	216, 3, 2, 2, 2, 250, 219, 3, 2, 2, 2, 250, 222, 3, 2, 2, 2, 250, 225,
	3, 2, 2, 2, 250, 228, 3, 2, 2, 2, 250, 231, 3, 2, 2, 2, 250, 234, 3, 2,
	2, 2, 250, 241, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2,
	252, 253, 3, 2, 2, 2, 253, 43, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 255, 256,
	7, 52, 2, 2, 256, 45, 3, 2, 2, 2, 257, 258, 7, 9, 2, 2, 258, 259, 5, 42,
	22, 2, 259, 260, 7, 8, 2, 2, 260, 47, 3, 2, 2, 2, 261, 264, 7, 53, 2, 2,
	262, 265, 5, 18, 10, 2, 263, 265, 5, 46, 24, 2, 264, 262, 3, 2, 2, 2, 264,
	263, 3, 2, 2, 2, 265, 267, 3, 2, 2, 2, 266, 261, 3, 2, 2, 2, 267, 268,
	3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 49, 3, 2,
	2, 2, 270, 271, 9, 10, 2, 2, 271, 272, 7, 7, 2, 2, 272, 273, 5, 48, 25,
	2, 273, 274, 7, 8, 2, 2, 274, 51, 3, 2, 2, 2, 275, 276, 7, 52, 2, 2, 276,
	277, 7, 7, 2, 2, 277, 278, 5, 24, 13, 2, 278, 279, 7, 8, 2, 2, 279, 53,
	3, 2, 2, 2, 280, 283, 7, 53, 2, 2, 281, 284, 7, 52, 2, 2, 282, 284, 5,
	20, 11, 2, 283, 281, 3, 2, 2, 2, 283, 282, 3, 2, 2, 2, 284, 286, 3, 2,
	2, 2, 285, 280, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2,
	287, 288, 3, 2, 2, 2, 288, 55, 3, 2, 2, 2, 27, 65, 72, 74, 84, 86, 105,
	111, 121, 128, 134, 140, 146, 152, 158, 178, 192, 198, 214, 246, 250, 252,
	264, 268, 283, 287,
}
var literalNames = []string{
	"", "'='", "'**'", "'{'", "'}'", "'('", "')'", "'$('", "'['", "']'", "'.'",
	"':'", "','", "';'", "'<'", "'<='", "'>='", "'>'", "'=='", "'!='", "'&&'",
	"'||'", "'!'", "'&'", "'|'", "'+'", "'-'", "'*'", "'^'", "'%'", "'in'",
	"'allow'", "'match'", "'if'", "'get'", "'exists'", "'true'", "'false'",
	"'list'", "'create'", "'update'", "'read'", "'write'", "'delete'", "'function'",
	"'return'", "'null'", "'service'", "", "", "", "'/'",
}
var symbolicNames = []string{
	"", "", "", "CurlyOpen", "CurlyClose", "BracketOpen", "BracketClose", "PathVariableBracket",
	"SquareBracketOpen", "SquareBracketClose", "Dot", "Colon", "Comma", "Semicolon",
	"LessThan", "LessOrEqual", "GreaterOrEqual", "GreaterThan", "Equals", "Unequal",
	"LogicalAnd", "LogicalOr", "LogicalNot", "BinaryAnd", "BinaryOr", "ArithmeticAdd",
	"ArithmeticSub", "ArithmeticMul", "ArithmeticExp", "ArithmeticModus", "InOperator",
	"Allow", "Match", "If", "Get", "Exists", "True", "False", "List", "Create",
	"Update", "Read", "Write", "Delete", "Function", "Return", "Null", "Service",
	"Number", "StringExpression", "Identifier", "Slash", "WS", "Comment",
}

var ruleNames = []string{
	"service", "namespaceIdentifier", "namespaceBlock", "matchBlock", "allowKey",
	"comment", "matcher", "allow", "getPathVariable", "pathVariable", "arg",
	"arguments", "memberArg", "memberArguments", "argDeclaration", "argDeclarations",
	"functionDeclaration", "fieldReference", "memberFunction", "id", "expression",
	"objectReference", "getPathExpressionVariable", "getPath", "ruleFunctionCall",
	"functionCall", "matchPath",
}

type MongoRPCRulesParser struct {
	*antlr.BaseParser
}

// NewMongoRPCRulesParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *MongoRPCRulesParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMongoRPCRulesParser(input antlr.TokenStream) *MongoRPCRulesParser {
	this := new(MongoRPCRulesParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "MongoRPCRules.g4"

	return this
}

// MongoRPCRulesParser tokens.
const (
	MongoRPCRulesParserEOF                 = antlr.TokenEOF
	MongoRPCRulesParserT__0                = 1
	MongoRPCRulesParserT__1                = 2
	MongoRPCRulesParserCurlyOpen           = 3
	MongoRPCRulesParserCurlyClose          = 4
	MongoRPCRulesParserBracketOpen         = 5
	MongoRPCRulesParserBracketClose        = 6
	MongoRPCRulesParserPathVariableBracket = 7
	MongoRPCRulesParserSquareBracketOpen   = 8
	MongoRPCRulesParserSquareBracketClose  = 9
	MongoRPCRulesParserDot                 = 10
	MongoRPCRulesParserColon               = 11
	MongoRPCRulesParserComma               = 12
	MongoRPCRulesParserSemicolon           = 13
	MongoRPCRulesParserLessThan            = 14
	MongoRPCRulesParserLessOrEqual         = 15
	MongoRPCRulesParserGreaterOrEqual      = 16
	MongoRPCRulesParserGreaterThan         = 17
	MongoRPCRulesParserEquals              = 18
	MongoRPCRulesParserUnequal             = 19
	MongoRPCRulesParserLogicalAnd          = 20
	MongoRPCRulesParserLogicalOr           = 21
	MongoRPCRulesParserLogicalNot          = 22
	MongoRPCRulesParserBinaryAnd           = 23
	MongoRPCRulesParserBinaryOr            = 24
	MongoRPCRulesParserArithmeticAdd       = 25
	MongoRPCRulesParserArithmeticSub       = 26
	MongoRPCRulesParserArithmeticMul       = 27
	MongoRPCRulesParserArithmeticExp       = 28
	MongoRPCRulesParserArithmeticModus     = 29
	MongoRPCRulesParserInOperator          = 30
	MongoRPCRulesParserAllow               = 31
	MongoRPCRulesParserMatch               = 32
	MongoRPCRulesParserIf                  = 33
	MongoRPCRulesParserGet                 = 34
	MongoRPCRulesParserExists              = 35
	MongoRPCRulesParserTrue                = 36
	MongoRPCRulesParserFalse               = 37
	MongoRPCRulesParserList                = 38
	MongoRPCRulesParserCreate              = 39
	MongoRPCRulesParserUpdate              = 40
	MongoRPCRulesParserRead                = 41
	MongoRPCRulesParserWrite               = 42
	MongoRPCRulesParserDelete              = 43
	MongoRPCRulesParserFunction            = 44
	MongoRPCRulesParserReturn              = 45
	MongoRPCRulesParserNull                = 46
	MongoRPCRulesParserService             = 47
	MongoRPCRulesParserNumber              = 48
	MongoRPCRulesParserStringExpression    = 49
	MongoRPCRulesParserIdentifier          = 50
	MongoRPCRulesParserSlash               = 51
	MongoRPCRulesParserWS                  = 52
	MongoRPCRulesParserComment             = 53
)

// MongoRPCRulesParser rules.
const (
	MongoRPCRulesParserRULE_service                   = 0
	MongoRPCRulesParserRULE_namespaceIdentifier       = 1
	MongoRPCRulesParserRULE_namespaceBlock            = 2
	MongoRPCRulesParserRULE_matchBlock                = 3
	MongoRPCRulesParserRULE_allowKey                  = 4
	MongoRPCRulesParserRULE_comment                   = 5
	MongoRPCRulesParserRULE_matcher                   = 6
	MongoRPCRulesParserRULE_allow                     = 7
	MongoRPCRulesParserRULE_getPathVariable           = 8
	MongoRPCRulesParserRULE_pathVariable              = 9
	MongoRPCRulesParserRULE_arg                       = 10
	MongoRPCRulesParserRULE_arguments                 = 11
	MongoRPCRulesParserRULE_memberArg                 = 12
	MongoRPCRulesParserRULE_memberArguments           = 13
	MongoRPCRulesParserRULE_argDeclaration            = 14
	MongoRPCRulesParserRULE_argDeclarations           = 15
	MongoRPCRulesParserRULE_functionDeclaration       = 16
	MongoRPCRulesParserRULE_fieldReference            = 17
	MongoRPCRulesParserRULE_memberFunction            = 18
	MongoRPCRulesParserRULE_id                        = 19
	MongoRPCRulesParserRULE_expression                = 20
	MongoRPCRulesParserRULE_objectReference           = 21
	MongoRPCRulesParserRULE_getPathExpressionVariable = 22
	MongoRPCRulesParserRULE_getPath                   = 23
	MongoRPCRulesParserRULE_ruleFunctionCall          = 24
	MongoRPCRulesParserRULE_functionCall              = 25
	MongoRPCRulesParserRULE_matchPath                 = 26
)

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_service
	return p
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) Service() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserService, 0)
}

func (s *ServiceContext) NamespaceIdentifier() INamespaceIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceIdentifierContext)
}

func (s *ServiceContext) NamespaceBlock() INamespaceBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceBlockContext)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitService(s)
	}
}

func (p *MongoRPCRulesParser) Service() (localctx IServiceContext) {
	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MongoRPCRulesParserRULE_service)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(MongoRPCRulesParserService)
	}
	{
		p.SetState(55)
		p.NamespaceIdentifier()
	}
	{
		p.SetState(56)
		p.NamespaceBlock()
	}

	return localctx
}

// INamespaceIdentifierContext is an interface to support dynamic dispatch.
type INamespaceIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceIdentifierContext differentiates from other interfaces.
	IsNamespaceIdentifierContext()
}

type NamespaceIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceIdentifierContext() *NamespaceIdentifierContext {
	var p = new(NamespaceIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_namespaceIdentifier
	return p
}

func (*NamespaceIdentifierContext) IsNamespaceIdentifierContext() {}

func NewNamespaceIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceIdentifierContext {
	var p = new(NamespaceIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_namespaceIdentifier

	return p
}

func (s *NamespaceIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceIdentifierContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *NamespaceIdentifierContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *NamespaceIdentifierContext) AllDot() []antlr.TerminalNode {
	return s.GetTokens(MongoRPCRulesParserDot)
}

func (s *NamespaceIdentifierContext) Dot(i int) antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserDot, i)
}

func (s *NamespaceIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterNamespaceIdentifier(s)
	}
}

func (s *NamespaceIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitNamespaceIdentifier(s)
	}
}

func (p *MongoRPCRulesParser) NamespaceIdentifier() (localctx INamespaceIdentifierContext) {
	localctx = NewNamespaceIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MongoRPCRulesParserRULE_namespaceIdentifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Id()
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MongoRPCRulesParserDot {
		{
			p.SetState(59)
			p.Match(MongoRPCRulesParserDot)
		}
		{
			p.SetState(60)
			p.Id()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INamespaceBlockContext is an interface to support dynamic dispatch.
type INamespaceBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceBlockContext differentiates from other interfaces.
	IsNamespaceBlockContext()
}

type NamespaceBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceBlockContext() *NamespaceBlockContext {
	var p = new(NamespaceBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_namespaceBlock
	return p
}

func (*NamespaceBlockContext) IsNamespaceBlockContext() {}

func NewNamespaceBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceBlockContext {
	var p = new(NamespaceBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_namespaceBlock

	return p
}

func (s *NamespaceBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceBlockContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserCurlyOpen, 0)
}

func (s *NamespaceBlockContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserCurlyClose, 0)
}

func (s *NamespaceBlockContext) AllMatcher() []IMatcherContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMatcherContext)(nil)).Elem())
	var tst = make([]IMatcherContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMatcherContext)
		}
	}

	return tst
}

func (s *NamespaceBlockContext) Matcher(i int) IMatcherContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatcherContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMatcherContext)
}

func (s *NamespaceBlockContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *NamespaceBlockContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *NamespaceBlockContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem())
	var tst = make([]IFunctionDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDeclarationContext)
		}
	}

	return tst
}

func (s *NamespaceBlockContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *NamespaceBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterNamespaceBlock(s)
	}
}

func (s *NamespaceBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitNamespaceBlock(s)
	}
}

func (p *MongoRPCRulesParser) NamespaceBlock() (localctx INamespaceBlockContext) {
	localctx = NewNamespaceBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MongoRPCRulesParserRULE_namespaceBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(MongoRPCRulesParserCurlyOpen)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(MongoRPCRulesParserMatch-32))|(1<<(MongoRPCRulesParserFunction-32))|(1<<(MongoRPCRulesParserComment-32)))) != 0 {
		p.SetState(70)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MongoRPCRulesParserMatch:
			{
				p.SetState(67)
				p.Matcher()
			}

		case MongoRPCRulesParserComment:
			{
				p.SetState(68)
				p.Comment()
			}

		case MongoRPCRulesParserFunction:
			{
				p.SetState(69)
				p.FunctionDeclaration()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Match(MongoRPCRulesParserCurlyClose)
	}

	return localctx
}

// IMatchBlockContext is an interface to support dynamic dispatch.
type IMatchBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatchBlockContext differentiates from other interfaces.
	IsMatchBlockContext()
}

type MatchBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchBlockContext() *MatchBlockContext {
	var p = new(MatchBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_matchBlock
	return p
}

func (*MatchBlockContext) IsMatchBlockContext() {}

func NewMatchBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchBlockContext {
	var p = new(MatchBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_matchBlock

	return p
}

func (s *MatchBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchBlockContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserCurlyOpen, 0)
}

func (s *MatchBlockContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserCurlyClose, 0)
}

func (s *MatchBlockContext) AllAllow() []IAllowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAllowContext)(nil)).Elem())
	var tst = make([]IAllowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAllowContext)
		}
	}

	return tst
}

func (s *MatchBlockContext) Allow(i int) IAllowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAllowContext)
}

func (s *MatchBlockContext) AllMatcher() []IMatcherContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMatcherContext)(nil)).Elem())
	var tst = make([]IMatcherContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMatcherContext)
		}
	}

	return tst
}

func (s *MatchBlockContext) Matcher(i int) IMatcherContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatcherContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMatcherContext)
}

func (s *MatchBlockContext) AllComment() []antlr.TerminalNode {
	return s.GetTokens(MongoRPCRulesParserComment)
}

func (s *MatchBlockContext) Comment(i int) antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserComment, i)
}

func (s *MatchBlockContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem())
	var tst = make([]IFunctionDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDeclarationContext)
		}
	}

	return tst
}

func (s *MatchBlockContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *MatchBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterMatchBlock(s)
	}
}

func (s *MatchBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitMatchBlock(s)
	}
}

func (p *MongoRPCRulesParser) MatchBlock() (localctx IMatchBlockContext) {
	localctx = NewMatchBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MongoRPCRulesParserRULE_matchBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(MongoRPCRulesParserCurlyOpen)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(MongoRPCRulesParserAllow-31))|(1<<(MongoRPCRulesParserMatch-31))|(1<<(MongoRPCRulesParserFunction-31))|(1<<(MongoRPCRulesParserComment-31)))) != 0 {
		p.SetState(82)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MongoRPCRulesParserAllow:
			{
				p.SetState(78)
				p.Allow()
			}

		case MongoRPCRulesParserMatch:
			{
				p.SetState(79)
				p.Matcher()
			}

		case MongoRPCRulesParserComment:
			{
				p.SetState(80)
				p.Match(MongoRPCRulesParserComment)
			}

		case MongoRPCRulesParserFunction:
			{
				p.SetState(81)
				p.FunctionDeclaration()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(MongoRPCRulesParserCurlyClose)
	}

	return localctx
}

// IAllowKeyContext is an interface to support dynamic dispatch.
type IAllowKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllowKeyContext differentiates from other interfaces.
	IsAllowKeyContext()
}

type AllowKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllowKeyContext() *AllowKeyContext {
	var p = new(AllowKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_allowKey
	return p
}

func (*AllowKeyContext) IsAllowKeyContext() {}

func NewAllowKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllowKeyContext {
	var p = new(AllowKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_allowKey

	return p
}

func (s *AllowKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *AllowKeyContext) Read() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserRead, 0)
}

func (s *AllowKeyContext) Write() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserWrite, 0)
}

func (s *AllowKeyContext) Update() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserUpdate, 0)
}

func (s *AllowKeyContext) Delete() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserDelete, 0)
}

func (s *AllowKeyContext) Create() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserCreate, 0)
}

func (s *AllowKeyContext) List() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserList, 0)
}

func (s *AllowKeyContext) Get() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserGet, 0)
}

func (s *AllowKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllowKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllowKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterAllowKey(s)
	}
}

func (s *AllowKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitAllowKey(s)
	}
}

func (p *MongoRPCRulesParser) AllowKey() (localctx IAllowKeyContext) {
	localctx = NewAllowKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MongoRPCRulesParserRULE_allowKey)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(MongoRPCRulesParserGet-34))|(1<<(MongoRPCRulesParserList-34))|(1<<(MongoRPCRulesParserCreate-34))|(1<<(MongoRPCRulesParserUpdate-34))|(1<<(MongoRPCRulesParserRead-34))|(1<<(MongoRPCRulesParserWrite-34))|(1<<(MongoRPCRulesParserDelete-34)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) Comment() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserComment, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *MongoRPCRulesParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MongoRPCRulesParserRULE_comment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(MongoRPCRulesParserComment)
	}

	return localctx
}

// IMatcherContext is an interface to support dynamic dispatch.
type IMatcherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatcherContext differentiates from other interfaces.
	IsMatcherContext()
}

type MatcherContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatcherContext() *MatcherContext {
	var p = new(MatcherContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_matcher
	return p
}

func (*MatcherContext) IsMatcherContext() {}

func NewMatcherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatcherContext {
	var p = new(MatcherContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_matcher

	return p
}

func (s *MatcherContext) GetParser() antlr.Parser { return s.parser }

func (s *MatcherContext) Match() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserMatch, 0)
}

func (s *MatcherContext) MatchPath() IMatchPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchPathContext)
}

func (s *MatcherContext) MatchBlock() IMatchBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchBlockContext)
}

func (s *MatcherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatcherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatcherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterMatcher(s)
	}
}

func (s *MatcherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitMatcher(s)
	}
}

func (p *MongoRPCRulesParser) Matcher() (localctx IMatcherContext) {
	localctx = NewMatcherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MongoRPCRulesParserRULE_matcher)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(MongoRPCRulesParserMatch)
	}
	{
		p.SetState(94)
		p.MatchPath()
	}
	{
		p.SetState(95)
		p.MatchBlock()
	}

	return localctx
}

// IAllowContext is an interface to support dynamic dispatch.
type IAllowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllowContext differentiates from other interfaces.
	IsAllowContext()
}

type AllowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllowContext() *AllowContext {
	var p = new(AllowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_allow
	return p
}

func (*AllowContext) IsAllowContext() {}

func NewAllowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllowContext {
	var p = new(AllowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_allow

	return p
}

func (s *AllowContext) GetParser() antlr.Parser { return s.parser }

func (s *AllowContext) Allow() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserAllow, 0)
}

func (s *AllowContext) AllAllowKey() []IAllowKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAllowKeyContext)(nil)).Elem())
	var tst = make([]IAllowKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAllowKeyContext)
		}
	}

	return tst
}

func (s *AllowContext) AllowKey(i int) IAllowKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllowKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAllowKeyContext)
}

func (s *AllowContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSemicolon, 0)
}

func (s *AllowContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(MongoRPCRulesParserComma)
}

func (s *AllowContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserComma, i)
}

func (s *AllowContext) Colon() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserColon, 0)
}

func (s *AllowContext) If() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIf, 0)
}

func (s *AllowContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AllowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterAllow(s)
	}
}

func (s *AllowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitAllow(s)
	}
}

func (p *MongoRPCRulesParser) Allow() (localctx IAllowContext) {
	localctx = NewAllowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MongoRPCRulesParserRULE_allow)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(MongoRPCRulesParserAllow)
	}
	{
		p.SetState(98)
		p.AllowKey()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MongoRPCRulesParserComma {
		{
			p.SetState(99)
			p.Match(MongoRPCRulesParserComma)
		}
		{
			p.SetState(100)
			p.AllowKey()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MongoRPCRulesParserColon {
		{
			p.SetState(106)
			p.Match(MongoRPCRulesParserColon)
		}
		{
			p.SetState(107)
			p.Match(MongoRPCRulesParserIf)
		}
		{
			p.SetState(108)
			p.expression(0)
		}

	}
	{
		p.SetState(111)
		p.Match(MongoRPCRulesParserSemicolon)
	}

	return localctx
}

// IGetPathVariableContext is an interface to support dynamic dispatch.
type IGetPathVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetPathVariableContext differentiates from other interfaces.
	IsGetPathVariableContext()
}

type GetPathVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetPathVariableContext() *GetPathVariableContext {
	var p = new(GetPathVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_getPathVariable
	return p
}

func (*GetPathVariableContext) IsGetPathVariableContext() {}

func NewGetPathVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetPathVariableContext {
	var p = new(GetPathVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_getPathVariable

	return p
}

func (s *GetPathVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *GetPathVariableContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIdentifier, 0)
}

func (s *GetPathVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetPathVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetPathVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterGetPathVariable(s)
	}
}

func (s *GetPathVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitGetPathVariable(s)
	}
}

func (p *MongoRPCRulesParser) GetPathVariable() (localctx IGetPathVariableContext) {
	localctx = NewGetPathVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MongoRPCRulesParserRULE_getPathVariable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(MongoRPCRulesParserIdentifier)
	}

	return localctx
}

// IPathVariableContext is an interface to support dynamic dispatch.
type IPathVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathVariableContext differentiates from other interfaces.
	IsPathVariableContext()
}

type PathVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathVariableContext() *PathVariableContext {
	var p = new(PathVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_pathVariable
	return p
}

func (*PathVariableContext) IsPathVariableContext() {}

func NewPathVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathVariableContext {
	var p = new(PathVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_pathVariable

	return p
}

func (s *PathVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *PathVariableContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserCurlyOpen, 0)
}

func (s *PathVariableContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIdentifier, 0)
}

func (s *PathVariableContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserCurlyClose, 0)
}

func (s *PathVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterPathVariable(s)
	}
}

func (s *PathVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitPathVariable(s)
	}
}

func (p *MongoRPCRulesParser) PathVariable() (localctx IPathVariableContext) {
	localctx = NewPathVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MongoRPCRulesParserRULE_pathVariable)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(MongoRPCRulesParserCurlyOpen)
	}
	{
		p.SetState(116)
		p.Match(MongoRPCRulesParserIdentifier)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MongoRPCRulesParserT__0 {
		{
			p.SetState(117)
			p.Match(MongoRPCRulesParserT__0)
		}
		{
			p.SetState(118)
			p.Match(MongoRPCRulesParserT__1)
		}

	}
	{
		p.SetState(121)
		p.Match(MongoRPCRulesParserCurlyClose)
	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *MongoRPCRulesParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MongoRPCRulesParserRULE_arg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.expression(0)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArg() []IArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgContext)(nil)).Elem())
	var tst = make([]IArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Arg(i int) IArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *ArgumentsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(MongoRPCRulesParserComma)
}

func (s *ArgumentsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserComma, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *MongoRPCRulesParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MongoRPCRulesParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MongoRPCRulesParserBracketOpen)|(1<<MongoRPCRulesParserSquareBracketOpen)|(1<<MongoRPCRulesParserLogicalNot)|(1<<MongoRPCRulesParserArithmeticSub))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(MongoRPCRulesParserGet-34))|(1<<(MongoRPCRulesParserExists-34))|(1<<(MongoRPCRulesParserTrue-34))|(1<<(MongoRPCRulesParserFalse-34))|(1<<(MongoRPCRulesParserNull-34))|(1<<(MongoRPCRulesParserNumber-34))|(1<<(MongoRPCRulesParserStringExpression-34))|(1<<(MongoRPCRulesParserIdentifier-34)))) != 0) {
		{
			p.SetState(125)
			p.Arg()
		}

	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MongoRPCRulesParserComma {
		{
			p.SetState(128)
			p.Match(MongoRPCRulesParserComma)
		}
		{
			p.SetState(129)
			p.Arg()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMemberArgContext is an interface to support dynamic dispatch.
type IMemberArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberArgContext differentiates from other interfaces.
	IsMemberArgContext()
}

type MemberArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberArgContext() *MemberArgContext {
	var p = new(MemberArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_memberArg
	return p
}

func (*MemberArgContext) IsMemberArgContext() {}

func NewMemberArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberArgContext {
	var p = new(MemberArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_memberArg

	return p
}

func (s *MemberArgContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberArgContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterMemberArg(s)
	}
}

func (s *MemberArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitMemberArg(s)
	}
}

func (p *MongoRPCRulesParser) MemberArg() (localctx IMemberArgContext) {
	localctx = NewMemberArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MongoRPCRulesParserRULE_memberArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.expression(0)
	}

	return localctx
}

// IMemberArgumentsContext is an interface to support dynamic dispatch.
type IMemberArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberArgumentsContext differentiates from other interfaces.
	IsMemberArgumentsContext()
}

type MemberArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberArgumentsContext() *MemberArgumentsContext {
	var p = new(MemberArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_memberArguments
	return p
}

func (*MemberArgumentsContext) IsMemberArgumentsContext() {}

func NewMemberArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberArgumentsContext {
	var p = new(MemberArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_memberArguments

	return p
}

func (s *MemberArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberArgumentsContext) AllMemberArg() []IMemberArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemberArgContext)(nil)).Elem())
	var tst = make([]IMemberArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemberArgContext)
		}
	}

	return tst
}

func (s *MemberArgumentsContext) MemberArg(i int) IMemberArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemberArgContext)
}

func (s *MemberArgumentsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(MongoRPCRulesParserComma)
}

func (s *MemberArgumentsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserComma, i)
}

func (s *MemberArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterMemberArguments(s)
	}
}

func (s *MemberArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitMemberArguments(s)
	}
}

func (p *MongoRPCRulesParser) MemberArguments() (localctx IMemberArgumentsContext) {
	localctx = NewMemberArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MongoRPCRulesParserRULE_memberArguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MongoRPCRulesParserBracketOpen)|(1<<MongoRPCRulesParserSquareBracketOpen)|(1<<MongoRPCRulesParserLogicalNot)|(1<<MongoRPCRulesParserArithmeticSub))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(MongoRPCRulesParserGet-34))|(1<<(MongoRPCRulesParserExists-34))|(1<<(MongoRPCRulesParserTrue-34))|(1<<(MongoRPCRulesParserFalse-34))|(1<<(MongoRPCRulesParserNull-34))|(1<<(MongoRPCRulesParserNumber-34))|(1<<(MongoRPCRulesParserStringExpression-34))|(1<<(MongoRPCRulesParserIdentifier-34)))) != 0) {
		{
			p.SetState(137)
			p.MemberArg()
		}

	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MongoRPCRulesParserComma {
		{
			p.SetState(140)
			p.Match(MongoRPCRulesParserComma)
		}
		{
			p.SetState(141)
			p.MemberArg()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArgDeclarationContext is an interface to support dynamic dispatch.
type IArgDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgDeclarationContext differentiates from other interfaces.
	IsArgDeclarationContext()
}

type ArgDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgDeclarationContext() *ArgDeclarationContext {
	var p = new(ArgDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_argDeclaration
	return p
}

func (*ArgDeclarationContext) IsArgDeclarationContext() {}

func NewArgDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgDeclarationContext {
	var p = new(ArgDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_argDeclaration

	return p
}

func (s *ArgDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIdentifier, 0)
}

func (s *ArgDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterArgDeclaration(s)
	}
}

func (s *ArgDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitArgDeclaration(s)
	}
}

func (p *MongoRPCRulesParser) ArgDeclaration() (localctx IArgDeclarationContext) {
	localctx = NewArgDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MongoRPCRulesParserRULE_argDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(MongoRPCRulesParserIdentifier)
	}

	return localctx
}

// IArgDeclarationsContext is an interface to support dynamic dispatch.
type IArgDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgDeclarationsContext differentiates from other interfaces.
	IsArgDeclarationsContext()
}

type ArgDeclarationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgDeclarationsContext() *ArgDeclarationsContext {
	var p = new(ArgDeclarationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_argDeclarations
	return p
}

func (*ArgDeclarationsContext) IsArgDeclarationsContext() {}

func NewArgDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgDeclarationsContext {
	var p = new(ArgDeclarationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_argDeclarations

	return p
}

func (s *ArgDeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgDeclarationsContext) AllArgDeclaration() []IArgDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgDeclarationContext)(nil)).Elem())
	var tst = make([]IArgDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgDeclarationContext)
		}
	}

	return tst
}

func (s *ArgDeclarationsContext) ArgDeclaration(i int) IArgDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgDeclarationContext)
}

func (s *ArgDeclarationsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(MongoRPCRulesParserComma)
}

func (s *ArgDeclarationsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserComma, i)
}

func (s *ArgDeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgDeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgDeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterArgDeclarations(s)
	}
}

func (s *ArgDeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitArgDeclarations(s)
	}
}

func (p *MongoRPCRulesParser) ArgDeclarations() (localctx IArgDeclarationsContext) {
	localctx = NewArgDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MongoRPCRulesParserRULE_argDeclarations)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MongoRPCRulesParserIdentifier {
		{
			p.SetState(149)
			p.ArgDeclaration()
		}

	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MongoRPCRulesParserComma {
		{
			p.SetState(152)
			p.Match(MongoRPCRulesParserComma)
		}
		{
			p.SetState(153)
			p.ArgDeclaration()
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) Function() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserFunction, 0)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketOpen, 0)
}

func (s *FunctionDeclarationContext) ArgDeclarations() IArgDeclarationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgDeclarationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgDeclarationsContext)
}

func (s *FunctionDeclarationContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketClose, 0)
}

func (s *FunctionDeclarationContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserCurlyOpen, 0)
}

func (s *FunctionDeclarationContext) Return() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserReturn, 0)
}

func (s *FunctionDeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionDeclarationContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSemicolon, 0)
}

func (s *FunctionDeclarationContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserCurlyClose, 0)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *MongoRPCRulesParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MongoRPCRulesParserRULE_functionDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(MongoRPCRulesParserFunction)
	}
	{
		p.SetState(160)
		p.Match(MongoRPCRulesParserIdentifier)
	}
	{
		p.SetState(161)
		p.Match(MongoRPCRulesParserBracketOpen)
	}
	{
		p.SetState(162)
		p.ArgDeclarations()
	}
	{
		p.SetState(163)
		p.Match(MongoRPCRulesParserBracketClose)
	}
	{
		p.SetState(164)
		p.Match(MongoRPCRulesParserCurlyOpen)
	}
	{
		p.SetState(165)
		p.Match(MongoRPCRulesParserReturn)
	}
	{
		p.SetState(166)
		p.expression(0)
	}
	{
		p.SetState(167)
		p.Match(MongoRPCRulesParserSemicolon)
	}
	{
		p.SetState(168)
		p.Match(MongoRPCRulesParserCurlyClose)
	}

	return localctx
}

// IFieldReferenceContext is an interface to support dynamic dispatch.
type IFieldReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldReferenceContext differentiates from other interfaces.
	IsFieldReferenceContext()
}

type FieldReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldReferenceContext() *FieldReferenceContext {
	var p = new(FieldReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_fieldReference
	return p
}

func (*FieldReferenceContext) IsFieldReferenceContext() {}

func NewFieldReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldReferenceContext {
	var p = new(FieldReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_fieldReference

	return p
}

func (s *FieldReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldReferenceContext) CopyFrom(ctx *FieldReferenceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FieldReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldReferenceWithIdentifierContext struct {
	*FieldReferenceContext
}

func NewFieldReferenceWithIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldReferenceWithIdentifierContext {
	var p = new(FieldReferenceWithIdentifierContext)

	p.FieldReferenceContext = NewEmptyFieldReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldReferenceContext))

	return p
}

func (s *FieldReferenceWithIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldReferenceWithIdentifierContext) Dot() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserDot, 0)
}

func (s *FieldReferenceWithIdentifierContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *FieldReferenceWithIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterFieldReferenceWithIdentifier(s)
	}
}

func (s *FieldReferenceWithIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitFieldReferenceWithIdentifier(s)
	}
}

type FieldReferenceWithMemberRefContext struct {
	*FieldReferenceContext
}

func NewFieldReferenceWithMemberRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldReferenceWithMemberRefContext {
	var p = new(FieldReferenceWithMemberRefContext)

	p.FieldReferenceContext = NewEmptyFieldReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldReferenceContext))

	return p
}

func (s *FieldReferenceWithMemberRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldReferenceWithMemberRefContext) SquareBracketOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSquareBracketOpen, 0)
}

func (s *FieldReferenceWithMemberRefContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldReferenceWithMemberRefContext) SquareBracketClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSquareBracketClose, 0)
}

func (s *FieldReferenceWithMemberRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterFieldReferenceWithMemberRef(s)
	}
}

func (s *FieldReferenceWithMemberRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitFieldReferenceWithMemberRef(s)
	}
}

func (p *MongoRPCRulesParser) FieldReference() (localctx IFieldReferenceContext) {
	localctx = NewFieldReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MongoRPCRulesParserRULE_fieldReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MongoRPCRulesParserDot:
		localctx = NewFieldReferenceWithIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Match(MongoRPCRulesParserDot)
		}
		{
			p.SetState(171)
			p.Id()
		}

	case MongoRPCRulesParserSquareBracketOpen:
		localctx = NewFieldReferenceWithMemberRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Match(MongoRPCRulesParserSquareBracketOpen)
		}
		{
			p.SetState(173)
			p.expression(0)
		}
		{
			p.SetState(174)
			p.Match(MongoRPCRulesParserSquareBracketClose)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMemberFunctionContext is an interface to support dynamic dispatch.
type IMemberFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberFunctionContext differentiates from other interfaces.
	IsMemberFunctionContext()
}

type MemberFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberFunctionContext() *MemberFunctionContext {
	var p = new(MemberFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_memberFunction
	return p
}

func (*MemberFunctionContext) IsMemberFunctionContext() {}

func NewMemberFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberFunctionContext {
	var p = new(MemberFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_memberFunction

	return p
}

func (s *MemberFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberFunctionContext) Dot() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserDot, 0)
}

func (s *MemberFunctionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *MemberFunctionContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketOpen, 0)
}

func (s *MemberFunctionContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *MemberFunctionContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketClose, 0)
}

func (s *MemberFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterMemberFunction(s)
	}
}

func (s *MemberFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitMemberFunction(s)
	}
}

func (p *MongoRPCRulesParser) MemberFunction() (localctx IMemberFunctionContext) {
	localctx = NewMemberFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MongoRPCRulesParserRULE_memberFunction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(MongoRPCRulesParserDot)
	}
	{
		p.SetState(179)
		p.Id()
	}
	{
		p.SetState(180)
		p.Match(MongoRPCRulesParserBracketOpen)
	}
	{
		p.SetState(181)
		p.Arguments()
	}
	{
		p.SetState(182)
		p.Match(MongoRPCRulesParserBracketClose)
	}

	return localctx
}

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_id
	return p
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIdentifier, 0)
}

func (s *IdContext) Allow() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserAllow, 0)
}

func (s *IdContext) Match() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserMatch, 0)
}

func (s *IdContext) If() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIf, 0)
}

func (s *IdContext) Exists() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserExists, 0)
}

func (s *IdContext) True() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserTrue, 0)
}

func (s *IdContext) False() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserFalse, 0)
}

func (s *IdContext) List() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserList, 0)
}

func (s *IdContext) Create() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserCreate, 0)
}

func (s *IdContext) Update() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserUpdate, 0)
}

func (s *IdContext) Read() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserRead, 0)
}

func (s *IdContext) Write() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserWrite, 0)
}

func (s *IdContext) Delete() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserDelete, 0)
}

func (s *IdContext) Function() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserFunction, 0)
}

func (s *IdContext) Return() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserReturn, 0)
}

func (s *IdContext) Null() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserNull, 0)
}

func (s *IdContext) Service() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserService, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitId(s)
	}
}

func (p *MongoRPCRulesParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MongoRPCRulesParserRULE_id)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(MongoRPCRulesParserAllow-31))|(1<<(MongoRPCRulesParserMatch-31))|(1<<(MongoRPCRulesParserIf-31))|(1<<(MongoRPCRulesParserExists-31))|(1<<(MongoRPCRulesParserTrue-31))|(1<<(MongoRPCRulesParserFalse-31))|(1<<(MongoRPCRulesParserList-31))|(1<<(MongoRPCRulesParserCreate-31))|(1<<(MongoRPCRulesParserUpdate-31))|(1<<(MongoRPCRulesParserRead-31))|(1<<(MongoRPCRulesParserWrite-31))|(1<<(MongoRPCRulesParserDelete-31))|(1<<(MongoRPCRulesParserFunction-31))|(1<<(MongoRPCRulesParserReturn-31))|(1<<(MongoRPCRulesParserNull-31))|(1<<(MongoRPCRulesParserService-31))|(1<<(MongoRPCRulesParserIdentifier-31)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayExpressionContext struct {
	*ExpressionContext
}

func NewArrayExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExpressionContext {
	var p = new(ArrayExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExpressionContext) SquareBracketOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSquareBracketOpen, 0)
}

func (s *ArrayExpressionContext) SquareBracketClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSquareBracketClose, 0)
}

func (s *ArrayExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArrayExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayExpressionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(MongoRPCRulesParserComma)
}

func (s *ArrayExpressionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserComma, i)
}

func (s *ArrayExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterArrayExpression(s)
	}
}

func (s *ArrayExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitArrayExpression(s)
	}
}

type NumberExpressionContext struct {
	*ExpressionContext
}

func NewNumberExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExpressionContext {
	var p = new(NumberExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserNumber, 0)
}

func (s *NumberExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterNumberExpression(s)
	}
}

func (s *NumberExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitNumberExpression(s)
	}
}

type ObjectReferenceExpressionContext struct {
	*ExpressionContext
}

func NewObjectReferenceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectReferenceExpressionContext {
	var p = new(ObjectReferenceExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ObjectReferenceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectReferenceExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIdentifier, 0)
}

func (s *ObjectReferenceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterObjectReferenceExpression(s)
	}
}

func (s *ObjectReferenceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitObjectReferenceExpression(s)
	}
}

type ParenthesisExpressionContext struct {
	*ExpressionContext
}

func NewParenthesisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisExpressionContext {
	var p = new(ParenthesisExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisExpressionContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketOpen, 0)
}

func (s *ParenthesisExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesisExpressionContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketClose, 0)
}

func (s *ParenthesisExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterParenthesisExpression(s)
	}
}

func (s *ParenthesisExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitParenthesisExpression(s)
	}
}

type ArithmeticExpressionContext struct {
	*ExpressionContext
}

func NewArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticExpressionContext {
	var p = new(ArithmeticExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArithmeticExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArithmeticExpressionContext) ArithmeticAdd() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserArithmeticAdd, 0)
}

func (s *ArithmeticExpressionContext) ArithmeticSub() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserArithmeticSub, 0)
}

func (s *ArithmeticExpressionContext) ArithmeticMul() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserArithmeticMul, 0)
}

func (s *ArithmeticExpressionContext) Slash() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSlash, 0)
}

func (s *ArithmeticExpressionContext) ArithmeticExp() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserArithmeticExp, 0)
}

func (s *ArithmeticExpressionContext) ArithmeticModus() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserArithmeticModus, 0)
}

func (s *ArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterArithmeticExpression(s)
	}
}

func (s *ArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitArithmeticExpression(s)
	}
}

type MemberReferenceExpressionContext struct {
	*ExpressionContext
}

func NewMemberReferenceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberReferenceExpressionContext {
	var p = new(MemberReferenceExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MemberReferenceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberReferenceExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberReferenceExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserDot, 0)
}

func (s *MemberReferenceExpressionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *MemberReferenceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterMemberReferenceExpression(s)
	}
}

func (s *MemberReferenceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitMemberReferenceExpression(s)
	}
}

type BooleanExpressionContext struct {
	*ExpressionContext
}

func NewBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExpressionContext) True() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserTrue, 0)
}

func (s *BooleanExpressionContext) False() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserFalse, 0)
}

func (s *BooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterBooleanExpression(s)
	}
}

func (s *BooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitBooleanExpression(s)
	}
}

type FunctionExpressionContext struct {
	*ExpressionContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

type CompareExpressionContext struct {
	*ExpressionContext
}

func NewCompareExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpressionContext {
	var p = new(CompareExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompareExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CompareExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompareExpressionContext) LessThan() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserLessThan, 0)
}

func (s *CompareExpressionContext) LessOrEqual() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserLessOrEqual, 0)
}

func (s *CompareExpressionContext) Equals() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserEquals, 0)
}

func (s *CompareExpressionContext) Unequal() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserUnequal, 0)
}

func (s *CompareExpressionContext) GreaterThan() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserGreaterThan, 0)
}

func (s *CompareExpressionContext) GreaterOrEqual() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserGreaterOrEqual, 0)
}

func (s *CompareExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterCompareExpression(s)
	}
}

func (s *CompareExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitCompareExpression(s)
	}
}

type BinaryExpressionContext struct {
	*ExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) BinaryAnd() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBinaryAnd, 0)
}

func (s *BinaryExpressionContext) BinaryOr() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBinaryOr, 0)
}

func (s *BinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitBinaryExpression(s)
	}
}

type LogicalExpressionContext struct {
	*ExpressionContext
}

func NewLogicalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogicalExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalExpressionContext) LogicalAnd() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserLogicalAnd, 0)
}

func (s *LogicalExpressionContext) LogicalOr() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserLogicalOr, 0)
}

func (s *LogicalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitLogicalExpression(s)
	}
}

type GetExpressionContext struct {
	*ExpressionContext
}

func NewGetExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetExpressionContext {
	var p = new(GetExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *GetExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetExpressionContext) RuleFunctionCall() IRuleFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleFunctionCallContext)
}

func (s *GetExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterGetExpression(s)
	}
}

func (s *GetExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitGetExpression(s)
	}
}

type InExpressionContext struct {
	*ExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *InExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InExpressionContext) InOperator() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserInOperator, 0)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitInExpression(s)
	}
}

type StringExpressionContext struct {
	*ExpressionContext
}

func NewStringExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExpressionContext {
	var p = new(StringExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *StringExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExpressionContext) StringExpression() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserStringExpression, 0)
}

func (s *StringExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterStringExpression(s)
	}
}

func (s *StringExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitStringExpression(s)
	}
}

type NullExpressionContext struct {
	*ExpressionContext
}

func NewNullExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullExpressionContext {
	var p = new(NullExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NullExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullExpressionContext) Null() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserNull, 0)
}

func (s *NullExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterNullExpression(s)
	}
}

func (s *NullExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitNullExpression(s)
	}
}

type RangeExpressionContext struct {
	*ExpressionContext
}

func NewRangeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeExpressionContext {
	var p = new(RangeExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *RangeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *RangeExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RangeExpressionContext) SquareBracketOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSquareBracketOpen, 0)
}

func (s *RangeExpressionContext) SquareBracketClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSquareBracketClose, 0)
}

func (s *RangeExpressionContext) Colon() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserColon, 0)
}

func (s *RangeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterRangeExpression(s)
	}
}

func (s *RangeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitRangeExpression(s)
	}
}

type UnaryExpressionContext struct {
	*ExpressionContext
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) LogicalNot() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserLogicalNot, 0)
}

func (s *UnaryExpressionContext) ArithmeticSub() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserArithmeticSub, 0)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

type MemberFunctionExpressionContext struct {
	*ExpressionContext
}

func NewMemberFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberFunctionExpressionContext {
	var p = new(MemberFunctionExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MemberFunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberFunctionExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberFunctionExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserDot, 0)
}

func (s *MemberFunctionExpressionContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *MemberFunctionExpressionContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketOpen, 0)
}

func (s *MemberFunctionExpressionContext) MemberArguments() IMemberArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberArgumentsContext)
}

func (s *MemberFunctionExpressionContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketClose, 0)
}

func (s *MemberFunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterMemberFunctionExpression(s)
	}
}

func (s *MemberFunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitMemberFunctionExpression(s)
	}
}

func (p *MongoRPCRulesParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *MongoRPCRulesParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, MongoRPCRulesParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNullExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(187)
			p.Match(MongoRPCRulesParserNull)
		}

	case 2:
		localctx = NewArrayExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(188)
			p.Match(MongoRPCRulesParserSquareBracketOpen)
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MongoRPCRulesParserBracketOpen)|(1<<MongoRPCRulesParserSquareBracketOpen)|(1<<MongoRPCRulesParserLogicalNot)|(1<<MongoRPCRulesParserArithmeticSub))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(MongoRPCRulesParserGet-34))|(1<<(MongoRPCRulesParserExists-34))|(1<<(MongoRPCRulesParserTrue-34))|(1<<(MongoRPCRulesParserFalse-34))|(1<<(MongoRPCRulesParserNull-34))|(1<<(MongoRPCRulesParserNumber-34))|(1<<(MongoRPCRulesParserStringExpression-34))|(1<<(MongoRPCRulesParserIdentifier-34)))) != 0) {
			{
				p.SetState(189)
				p.expression(0)
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MongoRPCRulesParserComma {
			{
				p.SetState(192)
				p.Match(MongoRPCRulesParserComma)
			}
			{
				p.SetState(193)
				p.expression(0)
			}

			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(199)
			p.Match(MongoRPCRulesParserSquareBracketClose)
		}

	case 3:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(200)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MongoRPCRulesParserLogicalNot || _la == MongoRPCRulesParserArithmeticSub) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(201)
			p.expression(8)
		}

	case 4:
		localctx = NewStringExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(202)
			p.Match(MongoRPCRulesParserStringExpression)
		}

	case 5:
		localctx = NewNumberExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(203)
			p.Match(MongoRPCRulesParserNumber)
		}

	case 6:
		localctx = NewBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(204)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MongoRPCRulesParserTrue || _la == MongoRPCRulesParserFalse) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 7:
		localctx = NewObjectReferenceExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(205)
			p.Match(MongoRPCRulesParserIdentifier)
		}

	case 8:
		localctx = NewGetExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(206)
			p.RuleFunctionCall()
		}

	case 9:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(207)
			p.FunctionCall()
		}

	case 10:
		localctx = NewParenthesisExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(208)
			p.Match(MongoRPCRulesParserBracketOpen)
		}
		{
			p.SetState(209)
			p.expression(0)
		}
		{
			p.SetState(210)
			p.Match(MongoRPCRulesParserBracketClose)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(248)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCompareExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MongoRPCRulesParserRULE_expression)
				p.SetState(214)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(215)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MongoRPCRulesParserLessThan)|(1<<MongoRPCRulesParserLessOrEqual)|(1<<MongoRPCRulesParserGreaterOrEqual)|(1<<MongoRPCRulesParserGreaterThan)|(1<<MongoRPCRulesParserEquals)|(1<<MongoRPCRulesParserUnequal))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(216)
					p.expression(18)
				}

			case 2:
				localctx = NewLogicalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MongoRPCRulesParserRULE_expression)
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(218)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MongoRPCRulesParserLogicalAnd || _la == MongoRPCRulesParserLogicalOr) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(219)
					p.expression(17)
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MongoRPCRulesParserRULE_expression)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(221)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MongoRPCRulesParserBinaryAnd || _la == MongoRPCRulesParserBinaryOr) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(222)
					p.expression(16)
				}

			case 4:
				localctx = NewArithmeticExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MongoRPCRulesParserRULE_expression)
				p.SetState(223)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(224)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-25)&-(0x1f+1)) == 0 && ((1<<uint((_la-25)))&((1<<(MongoRPCRulesParserArithmeticAdd-25))|(1<<(MongoRPCRulesParserArithmeticSub-25))|(1<<(MongoRPCRulesParserArithmeticMul-25))|(1<<(MongoRPCRulesParserArithmeticExp-25))|(1<<(MongoRPCRulesParserArithmeticModus-25))|(1<<(MongoRPCRulesParserSlash-25)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(225)
					p.expression(15)
				}

			case 5:
				localctx = NewInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MongoRPCRulesParserRULE_expression)
				p.SetState(226)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(227)
					p.Match(MongoRPCRulesParserInOperator)
				}
				{
					p.SetState(228)
					p.expression(14)
				}

			case 6:
				localctx = NewMemberReferenceExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MongoRPCRulesParserRULE_expression)
				p.SetState(229)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(230)
					p.Match(MongoRPCRulesParserDot)
				}
				{
					p.SetState(231)
					p.Id()
				}

			case 7:
				localctx = NewMemberFunctionExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MongoRPCRulesParserRULE_expression)
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(233)
					p.Match(MongoRPCRulesParserDot)
				}
				{
					p.SetState(234)
					p.Id()
				}
				{
					p.SetState(235)
					p.Match(MongoRPCRulesParserBracketOpen)
				}
				{
					p.SetState(236)
					p.MemberArguments()
				}
				{
					p.SetState(237)
					p.Match(MongoRPCRulesParserBracketClose)
				}

			case 8:
				localctx = NewRangeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MongoRPCRulesParserRULE_expression)
				p.SetState(239)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(240)
					p.Match(MongoRPCRulesParserSquareBracketOpen)
				}
				{
					p.SetState(241)
					p.expression(0)
				}
				p.SetState(244)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == MongoRPCRulesParserColon {
					{
						p.SetState(242)
						p.Match(MongoRPCRulesParserColon)
					}
					{
						p.SetState(243)
						p.expression(0)
					}

				}
				{
					p.SetState(246)
					p.Match(MongoRPCRulesParserSquareBracketClose)
				}

			}

		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IObjectReferenceContext is an interface to support dynamic dispatch.
type IObjectReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectReferenceContext differentiates from other interfaces.
	IsObjectReferenceContext()
}

type ObjectReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectReferenceContext() *ObjectReferenceContext {
	var p = new(ObjectReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_objectReference
	return p
}

func (*ObjectReferenceContext) IsObjectReferenceContext() {}

func NewObjectReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectReferenceContext {
	var p = new(ObjectReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_objectReference

	return p
}

func (s *ObjectReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIdentifier, 0)
}

func (s *ObjectReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterObjectReference(s)
	}
}

func (s *ObjectReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitObjectReference(s)
	}
}

func (p *MongoRPCRulesParser) ObjectReference() (localctx IObjectReferenceContext) {
	localctx = NewObjectReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MongoRPCRulesParserRULE_objectReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(MongoRPCRulesParserIdentifier)
	}

	return localctx
}

// IGetPathExpressionVariableContext is an interface to support dynamic dispatch.
type IGetPathExpressionVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetPathExpressionVariableContext differentiates from other interfaces.
	IsGetPathExpressionVariableContext()
}

type GetPathExpressionVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetPathExpressionVariableContext() *GetPathExpressionVariableContext {
	var p = new(GetPathExpressionVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_getPathExpressionVariable
	return p
}

func (*GetPathExpressionVariableContext) IsGetPathExpressionVariableContext() {}

func NewGetPathExpressionVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetPathExpressionVariableContext {
	var p = new(GetPathExpressionVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_getPathExpressionVariable

	return p
}

func (s *GetPathExpressionVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *GetPathExpressionVariableContext) PathVariableBracket() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserPathVariableBracket, 0)
}

func (s *GetPathExpressionVariableContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GetPathExpressionVariableContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketClose, 0)
}

func (s *GetPathExpressionVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetPathExpressionVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetPathExpressionVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterGetPathExpressionVariable(s)
	}
}

func (s *GetPathExpressionVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitGetPathExpressionVariable(s)
	}
}

func (p *MongoRPCRulesParser) GetPathExpressionVariable() (localctx IGetPathExpressionVariableContext) {
	localctx = NewGetPathExpressionVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MongoRPCRulesParserRULE_getPathExpressionVariable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(MongoRPCRulesParserPathVariableBracket)
	}
	{
		p.SetState(256)
		p.expression(0)
	}
	{
		p.SetState(257)
		p.Match(MongoRPCRulesParserBracketClose)
	}

	return localctx
}

// IGetPathContext is an interface to support dynamic dispatch.
type IGetPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetPathContext differentiates from other interfaces.
	IsGetPathContext()
}

type GetPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetPathContext() *GetPathContext {
	var p = new(GetPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_getPath
	return p
}

func (*GetPathContext) IsGetPathContext() {}

func NewGetPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetPathContext {
	var p = new(GetPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_getPath

	return p
}

func (s *GetPathContext) GetParser() antlr.Parser { return s.parser }

func (s *GetPathContext) AllSlash() []antlr.TerminalNode {
	return s.GetTokens(MongoRPCRulesParserSlash)
}

func (s *GetPathContext) Slash(i int) antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSlash, i)
}

func (s *GetPathContext) AllGetPathVariable() []IGetPathVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGetPathVariableContext)(nil)).Elem())
	var tst = make([]IGetPathVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGetPathVariableContext)
		}
	}

	return tst
}

func (s *GetPathContext) GetPathVariable(i int) IGetPathVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetPathVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGetPathVariableContext)
}

func (s *GetPathContext) AllGetPathExpressionVariable() []IGetPathExpressionVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGetPathExpressionVariableContext)(nil)).Elem())
	var tst = make([]IGetPathExpressionVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGetPathExpressionVariableContext)
		}
	}

	return tst
}

func (s *GetPathContext) GetPathExpressionVariable(i int) IGetPathExpressionVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetPathExpressionVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGetPathExpressionVariableContext)
}

func (s *GetPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterGetPath(s)
	}
}

func (s *GetPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitGetPath(s)
	}
}

func (p *MongoRPCRulesParser) GetPath() (localctx IGetPathContext) {
	localctx = NewGetPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MongoRPCRulesParserRULE_getPath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MongoRPCRulesParserSlash {
		{
			p.SetState(259)
			p.Match(MongoRPCRulesParserSlash)
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MongoRPCRulesParserIdentifier:
			{
				p.SetState(260)
				p.GetPathVariable()
			}

		case MongoRPCRulesParserPathVariableBracket:
			{
				p.SetState(261)
				p.GetPathExpressionVariable()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRuleFunctionCallContext is an interface to support dynamic dispatch.
type IRuleFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleFunctionCallContext differentiates from other interfaces.
	IsRuleFunctionCallContext()
}

type RuleFunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleFunctionCallContext() *RuleFunctionCallContext {
	var p = new(RuleFunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_ruleFunctionCall
	return p
}

func (*RuleFunctionCallContext) IsRuleFunctionCallContext() {}

func NewRuleFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleFunctionCallContext {
	var p = new(RuleFunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_ruleFunctionCall

	return p
}

func (s *RuleFunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleFunctionCallContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketOpen, 0)
}

func (s *RuleFunctionCallContext) GetPath() IGetPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetPathContext)
}

func (s *RuleFunctionCallContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketClose, 0)
}

func (s *RuleFunctionCallContext) Get() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserGet, 0)
}

func (s *RuleFunctionCallContext) Exists() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserExists, 0)
}

func (s *RuleFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleFunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterRuleFunctionCall(s)
	}
}

func (s *RuleFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitRuleFunctionCall(s)
	}
}

func (p *MongoRPCRulesParser) RuleFunctionCall() (localctx IRuleFunctionCallContext) {
	localctx = NewRuleFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MongoRPCRulesParserRULE_ruleFunctionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MongoRPCRulesParserGet || _la == MongoRPCRulesParserExists) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(269)
		p.Match(MongoRPCRulesParserBracketOpen)
	}
	{
		p.SetState(270)
		p.GetPath()
	}
	{
		p.SetState(271)
		p.Match(MongoRPCRulesParserBracketClose)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIdentifier, 0)
}

func (s *FunctionCallContext) BracketOpen() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketOpen, 0)
}

func (s *FunctionCallContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FunctionCallContext) BracketClose() antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserBracketClose, 0)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *MongoRPCRulesParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MongoRPCRulesParserRULE_functionCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(MongoRPCRulesParserIdentifier)
	}
	{
		p.SetState(274)
		p.Match(MongoRPCRulesParserBracketOpen)
	}
	{
		p.SetState(275)
		p.Arguments()
	}
	{
		p.SetState(276)
		p.Match(MongoRPCRulesParserBracketClose)
	}

	return localctx
}

// IMatchPathContext is an interface to support dynamic dispatch.
type IMatchPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatchPathContext differentiates from other interfaces.
	IsMatchPathContext()
}

type MatchPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchPathContext() *MatchPathContext {
	var p = new(MatchPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MongoRPCRulesParserRULE_matchPath
	return p
}

func (*MatchPathContext) IsMatchPathContext() {}

func NewMatchPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchPathContext {
	var p = new(MatchPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MongoRPCRulesParserRULE_matchPath

	return p
}

func (s *MatchPathContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchPathContext) AllSlash() []antlr.TerminalNode {
	return s.GetTokens(MongoRPCRulesParserSlash)
}

func (s *MatchPathContext) Slash(i int) antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserSlash, i)
}

func (s *MatchPathContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(MongoRPCRulesParserIdentifier)
}

func (s *MatchPathContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(MongoRPCRulesParserIdentifier, i)
}

func (s *MatchPathContext) AllPathVariable() []IPathVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPathVariableContext)(nil)).Elem())
	var tst = make([]IPathVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPathVariableContext)
		}
	}

	return tst
}

func (s *MatchPathContext) PathVariable(i int) IPathVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPathVariableContext)
}

func (s *MatchPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.EnterMatchPath(s)
	}
}

func (s *MatchPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MongoRPCRulesListener); ok {
		listenerT.ExitMatchPath(s)
	}
}

func (p *MongoRPCRulesParser) MatchPath() (localctx IMatchPathContext) {
	localctx = NewMatchPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MongoRPCRulesParserRULE_matchPath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MongoRPCRulesParserSlash {
		{
			p.SetState(278)
			p.Match(MongoRPCRulesParserSlash)
		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MongoRPCRulesParserIdentifier:
			{
				p.SetState(279)
				p.Match(MongoRPCRulesParserIdentifier)
			}

		case MongoRPCRulesParserCurlyOpen:
			{
				p.SetState(280)
				p.PathVariable()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *MongoRPCRulesParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 20:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MongoRPCRulesParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

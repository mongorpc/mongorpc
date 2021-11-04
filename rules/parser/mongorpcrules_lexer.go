// Code generated from MongoRPCRules.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 55, 334,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3,
	40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 49, 6, 49, 283, 10, 49, 13, 49, 14, 49, 284, 3, 49, 3,
	49, 6, 49, 289, 10, 49, 13, 49, 14, 49, 290, 5, 49, 293, 10, 49, 3, 50,
	3, 50, 7, 50, 297, 10, 50, 12, 50, 14, 50, 300, 11, 50, 3, 50, 3, 50, 3,
	50, 7, 50, 305, 10, 50, 12, 50, 14, 50, 308, 11, 50, 3, 50, 5, 50, 311,
	10, 50, 3, 51, 3, 51, 7, 51, 315, 10, 51, 12, 51, 14, 51, 318, 11, 51,
	3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 7,
	54, 330, 10, 54, 12, 54, 14, 54, 333, 11, 54, 2, 2, 55, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63,
	33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81,
	42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99,
	51, 101, 52, 103, 53, 105, 54, 107, 55, 3, 2, 8, 3, 2, 41, 41, 3, 2, 36,
	36, 5, 2, 67, 92, 97, 97, 99, 124, 7, 2, 47, 47, 50, 59, 67, 92, 97, 97,
	99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 3, 2, 12, 12, 2, 341, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2,
	2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2,
	2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3,
	2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73,
	3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2,
	81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2,
	2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2,
	2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3,
	2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 3, 109, 3, 2, 2, 2, 5,
	111, 3, 2, 2, 2, 7, 114, 3, 2, 2, 2, 9, 116, 3, 2, 2, 2, 11, 118, 3, 2,
	2, 2, 13, 120, 3, 2, 2, 2, 15, 122, 3, 2, 2, 2, 17, 125, 3, 2, 2, 2, 19,
	127, 3, 2, 2, 2, 21, 129, 3, 2, 2, 2, 23, 131, 3, 2, 2, 2, 25, 133, 3,
	2, 2, 2, 27, 135, 3, 2, 2, 2, 29, 137, 3, 2, 2, 2, 31, 139, 3, 2, 2, 2,
	33, 142, 3, 2, 2, 2, 35, 145, 3, 2, 2, 2, 37, 147, 3, 2, 2, 2, 39, 150,
	3, 2, 2, 2, 41, 153, 3, 2, 2, 2, 43, 156, 3, 2, 2, 2, 45, 159, 3, 2, 2,
	2, 47, 161, 3, 2, 2, 2, 49, 163, 3, 2, 2, 2, 51, 165, 3, 2, 2, 2, 53, 167,
	3, 2, 2, 2, 55, 169, 3, 2, 2, 2, 57, 171, 3, 2, 2, 2, 59, 173, 3, 2, 2,
	2, 61, 175, 3, 2, 2, 2, 63, 178, 3, 2, 2, 2, 65, 184, 3, 2, 2, 2, 67, 190,
	3, 2, 2, 2, 69, 193, 3, 2, 2, 2, 71, 197, 3, 2, 2, 2, 73, 204, 3, 2, 2,
	2, 75, 209, 3, 2, 2, 2, 77, 215, 3, 2, 2, 2, 79, 220, 3, 2, 2, 2, 81, 227,
	3, 2, 2, 2, 83, 234, 3, 2, 2, 2, 85, 239, 3, 2, 2, 2, 87, 245, 3, 2, 2,
	2, 89, 252, 3, 2, 2, 2, 91, 261, 3, 2, 2, 2, 93, 268, 3, 2, 2, 2, 95, 273,
	3, 2, 2, 2, 97, 282, 3, 2, 2, 2, 99, 310, 3, 2, 2, 2, 101, 312, 3, 2, 2,
	2, 103, 319, 3, 2, 2, 2, 105, 321, 3, 2, 2, 2, 107, 325, 3, 2, 2, 2, 109,
	110, 7, 63, 2, 2, 110, 4, 3, 2, 2, 2, 111, 112, 7, 44, 2, 2, 112, 113,
	7, 44, 2, 2, 113, 6, 3, 2, 2, 2, 114, 115, 7, 125, 2, 2, 115, 8, 3, 2,
	2, 2, 116, 117, 7, 127, 2, 2, 117, 10, 3, 2, 2, 2, 118, 119, 7, 42, 2,
	2, 119, 12, 3, 2, 2, 2, 120, 121, 7, 43, 2, 2, 121, 14, 3, 2, 2, 2, 122,
	123, 7, 38, 2, 2, 123, 124, 7, 42, 2, 2, 124, 16, 3, 2, 2, 2, 125, 126,
	7, 93, 2, 2, 126, 18, 3, 2, 2, 2, 127, 128, 7, 95, 2, 2, 128, 20, 3, 2,
	2, 2, 129, 130, 7, 48, 2, 2, 130, 22, 3, 2, 2, 2, 131, 132, 7, 60, 2, 2,
	132, 24, 3, 2, 2, 2, 133, 134, 7, 46, 2, 2, 134, 26, 3, 2, 2, 2, 135, 136,
	7, 61, 2, 2, 136, 28, 3, 2, 2, 2, 137, 138, 7, 62, 2, 2, 138, 30, 3, 2,
	2, 2, 139, 140, 7, 62, 2, 2, 140, 141, 7, 63, 2, 2, 141, 32, 3, 2, 2, 2,
	142, 143, 7, 64, 2, 2, 143, 144, 7, 63, 2, 2, 144, 34, 3, 2, 2, 2, 145,
	146, 7, 64, 2, 2, 146, 36, 3, 2, 2, 2, 147, 148, 7, 63, 2, 2, 148, 149,
	7, 63, 2, 2, 149, 38, 3, 2, 2, 2, 150, 151, 7, 35, 2, 2, 151, 152, 7, 63,
	2, 2, 152, 40, 3, 2, 2, 2, 153, 154, 7, 40, 2, 2, 154, 155, 7, 40, 2, 2,
	155, 42, 3, 2, 2, 2, 156, 157, 7, 126, 2, 2, 157, 158, 7, 126, 2, 2, 158,
	44, 3, 2, 2, 2, 159, 160, 7, 35, 2, 2, 160, 46, 3, 2, 2, 2, 161, 162, 7,
	40, 2, 2, 162, 48, 3, 2, 2, 2, 163, 164, 7, 126, 2, 2, 164, 50, 3, 2, 2,
	2, 165, 166, 7, 45, 2, 2, 166, 52, 3, 2, 2, 2, 167, 168, 7, 47, 2, 2, 168,
	54, 3, 2, 2, 2, 169, 170, 7, 44, 2, 2, 170, 56, 3, 2, 2, 2, 171, 172, 7,
	96, 2, 2, 172, 58, 3, 2, 2, 2, 173, 174, 7, 39, 2, 2, 174, 60, 3, 2, 2,
	2, 175, 176, 7, 107, 2, 2, 176, 177, 7, 112, 2, 2, 177, 62, 3, 2, 2, 2,
	178, 179, 7, 99, 2, 2, 179, 180, 7, 110, 2, 2, 180, 181, 7, 110, 2, 2,
	181, 182, 7, 113, 2, 2, 182, 183, 7, 121, 2, 2, 183, 64, 3, 2, 2, 2, 184,
	185, 7, 111, 2, 2, 185, 186, 7, 99, 2, 2, 186, 187, 7, 118, 2, 2, 187,
	188, 7, 101, 2, 2, 188, 189, 7, 106, 2, 2, 189, 66, 3, 2, 2, 2, 190, 191,
	7, 107, 2, 2, 191, 192, 7, 104, 2, 2, 192, 68, 3, 2, 2, 2, 193, 194, 7,
	105, 2, 2, 194, 195, 7, 103, 2, 2, 195, 196, 7, 118, 2, 2, 196, 70, 3,
	2, 2, 2, 197, 198, 7, 103, 2, 2, 198, 199, 7, 122, 2, 2, 199, 200, 7, 107,
	2, 2, 200, 201, 7, 117, 2, 2, 201, 202, 7, 118, 2, 2, 202, 203, 7, 117,
	2, 2, 203, 72, 3, 2, 2, 2, 204, 205, 7, 118, 2, 2, 205, 206, 7, 116, 2,
	2, 206, 207, 7, 119, 2, 2, 207, 208, 7, 103, 2, 2, 208, 74, 3, 2, 2, 2,
	209, 210, 7, 104, 2, 2, 210, 211, 7, 99, 2, 2, 211, 212, 7, 110, 2, 2,
	212, 213, 7, 117, 2, 2, 213, 214, 7, 103, 2, 2, 214, 76, 3, 2, 2, 2, 215,
	216, 7, 110, 2, 2, 216, 217, 7, 107, 2, 2, 217, 218, 7, 117, 2, 2, 218,
	219, 7, 118, 2, 2, 219, 78, 3, 2, 2, 2, 220, 221, 7, 101, 2, 2, 221, 222,
	7, 116, 2, 2, 222, 223, 7, 103, 2, 2, 223, 224, 7, 99, 2, 2, 224, 225,
	7, 118, 2, 2, 225, 226, 7, 103, 2, 2, 226, 80, 3, 2, 2, 2, 227, 228, 7,
	119, 2, 2, 228, 229, 7, 114, 2, 2, 229, 230, 7, 102, 2, 2, 230, 231, 7,
	99, 2, 2, 231, 232, 7, 118, 2, 2, 232, 233, 7, 103, 2, 2, 233, 82, 3, 2,
	2, 2, 234, 235, 7, 116, 2, 2, 235, 236, 7, 103, 2, 2, 236, 237, 7, 99,
	2, 2, 237, 238, 7, 102, 2, 2, 238, 84, 3, 2, 2, 2, 239, 240, 7, 121, 2,
	2, 240, 241, 7, 116, 2, 2, 241, 242, 7, 107, 2, 2, 242, 243, 7, 118, 2,
	2, 243, 244, 7, 103, 2, 2, 244, 86, 3, 2, 2, 2, 245, 246, 7, 102, 2, 2,
	246, 247, 7, 103, 2, 2, 247, 248, 7, 110, 2, 2, 248, 249, 7, 103, 2, 2,
	249, 250, 7, 118, 2, 2, 250, 251, 7, 103, 2, 2, 251, 88, 3, 2, 2, 2, 252,
	253, 7, 104, 2, 2, 253, 254, 7, 119, 2, 2, 254, 255, 7, 112, 2, 2, 255,
	256, 7, 101, 2, 2, 256, 257, 7, 118, 2, 2, 257, 258, 7, 107, 2, 2, 258,
	259, 7, 113, 2, 2, 259, 260, 7, 112, 2, 2, 260, 90, 3, 2, 2, 2, 261, 262,
	7, 116, 2, 2, 262, 263, 7, 103, 2, 2, 263, 264, 7, 118, 2, 2, 264, 265,
	7, 119, 2, 2, 265, 266, 7, 116, 2, 2, 266, 267, 7, 112, 2, 2, 267, 92,
	3, 2, 2, 2, 268, 269, 7, 112, 2, 2, 269, 270, 7, 119, 2, 2, 270, 271, 7,
	110, 2, 2, 271, 272, 7, 110, 2, 2, 272, 94, 3, 2, 2, 2, 273, 274, 7, 117,
	2, 2, 274, 275, 7, 103, 2, 2, 275, 276, 7, 116, 2, 2, 276, 277, 7, 120,
	2, 2, 277, 278, 7, 107, 2, 2, 278, 279, 7, 101, 2, 2, 279, 280, 7, 103,
	2, 2, 280, 96, 3, 2, 2, 2, 281, 283, 4, 50, 59, 2, 282, 281, 3, 2, 2, 2,
	283, 284, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285,
	292, 3, 2, 2, 2, 286, 288, 7, 48, 2, 2, 287, 289, 4, 50, 59, 2, 288, 287,
	3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290, 291, 3, 2,
	2, 2, 291, 293, 3, 2, 2, 2, 292, 286, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2,
	293, 98, 3, 2, 2, 2, 294, 298, 7, 41, 2, 2, 295, 297, 10, 2, 2, 2, 296,
	295, 3, 2, 2, 2, 297, 300, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2, 298, 299,
	3, 2, 2, 2, 299, 301, 3, 2, 2, 2, 300, 298, 3, 2, 2, 2, 301, 311, 7, 41,
	2, 2, 302, 306, 7, 36, 2, 2, 303, 305, 10, 3, 2, 2, 304, 303, 3, 2, 2,
	2, 305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307,
	309, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 311, 7, 36, 2, 2, 310, 294,
	3, 2, 2, 2, 310, 302, 3, 2, 2, 2, 311, 100, 3, 2, 2, 2, 312, 316, 9, 4,
	2, 2, 313, 315, 9, 5, 2, 2, 314, 313, 3, 2, 2, 2, 315, 318, 3, 2, 2, 2,
	316, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 102, 3, 2, 2, 2, 318,
	316, 3, 2, 2, 2, 319, 320, 7, 49, 2, 2, 320, 104, 3, 2, 2, 2, 321, 322,
	9, 6, 2, 2, 322, 323, 3, 2, 2, 2, 323, 324, 8, 53, 2, 2, 324, 106, 3, 2,
	2, 2, 325, 326, 7, 49, 2, 2, 326, 327, 7, 49, 2, 2, 327, 331, 3, 2, 2,
	2, 328, 330, 10, 7, 2, 2, 329, 328, 3, 2, 2, 2, 330, 333, 3, 2, 2, 2, 331,
	329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 108, 3, 2, 2, 2, 333, 331,
	3, 2, 2, 2, 11, 2, 284, 290, 292, 298, 306, 310, 316, 331, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'**'", "'{'", "'}'", "'('", "')'", "'$('", "'['", "']'", "'.'",
	"':'", "','", "';'", "'<'", "'<='", "'>='", "'>'", "'=='", "'!='", "'&&'",
	"'||'", "'!'", "'&'", "'|'", "'+'", "'-'", "'*'", "'^'", "'%'", "'in'",
	"'allow'", "'match'", "'if'", "'get'", "'exists'", "'true'", "'false'",
	"'list'", "'create'", "'update'", "'read'", "'write'", "'delete'", "'function'",
	"'return'", "'null'", "'service'", "", "", "", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "CurlyOpen", "CurlyClose", "BracketOpen", "BracketClose", "PathVariableBracket",
	"SquareBracketOpen", "SquareBracketClose", "Dot", "Colon", "Comma", "Semicolon",
	"LessThan", "LessOrEqual", "GreaterOrEqual", "GreaterThan", "Equals", "Unequal",
	"LogicalAnd", "LogicalOr", "LogicalNot", "BinaryAnd", "BinaryOr", "ArithmeticAdd",
	"ArithmeticSub", "ArithmeticMul", "ArithmeticExp", "ArithmeticModus", "InOperator",
	"Allow", "Match", "If", "Get", "Exists", "True", "False", "List", "Create",
	"Update", "Read", "Write", "Delete", "Function", "Return", "Null", "Service",
	"Number", "StringExpression", "Identifier", "Slash", "WS", "Comment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "CurlyOpen", "CurlyClose", "BracketOpen", "BracketClose",
	"PathVariableBracket", "SquareBracketOpen", "SquareBracketClose", "Dot",
	"Colon", "Comma", "Semicolon", "LessThan", "LessOrEqual", "GreaterOrEqual",
	"GreaterThan", "Equals", "Unequal", "LogicalAnd", "LogicalOr", "LogicalNot",
	"BinaryAnd", "BinaryOr", "ArithmeticAdd", "ArithmeticSub", "ArithmeticMul",
	"ArithmeticExp", "ArithmeticModus", "InOperator", "Allow", "Match", "If",
	"Get", "Exists", "True", "False", "List", "Create", "Update", "Read", "Write",
	"Delete", "Function", "Return", "Null", "Service", "Number", "StringExpression",
	"Identifier", "Slash", "WS", "Comment",
}

type MongoRPCRulesLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewMongoRPCRulesLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *MongoRPCRulesLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMongoRPCRulesLexer(input antlr.CharStream) *MongoRPCRulesLexer {
	l := new(MongoRPCRulesLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "MongoRPCRules.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MongoRPCRulesLexer tokens.
const (
	MongoRPCRulesLexerT__0                = 1
	MongoRPCRulesLexerT__1                = 2
	MongoRPCRulesLexerCurlyOpen           = 3
	MongoRPCRulesLexerCurlyClose          = 4
	MongoRPCRulesLexerBracketOpen         = 5
	MongoRPCRulesLexerBracketClose        = 6
	MongoRPCRulesLexerPathVariableBracket = 7
	MongoRPCRulesLexerSquareBracketOpen   = 8
	MongoRPCRulesLexerSquareBracketClose  = 9
	MongoRPCRulesLexerDot                 = 10
	MongoRPCRulesLexerColon               = 11
	MongoRPCRulesLexerComma               = 12
	MongoRPCRulesLexerSemicolon           = 13
	MongoRPCRulesLexerLessThan            = 14
	MongoRPCRulesLexerLessOrEqual         = 15
	MongoRPCRulesLexerGreaterOrEqual      = 16
	MongoRPCRulesLexerGreaterThan         = 17
	MongoRPCRulesLexerEquals              = 18
	MongoRPCRulesLexerUnequal             = 19
	MongoRPCRulesLexerLogicalAnd          = 20
	MongoRPCRulesLexerLogicalOr           = 21
	MongoRPCRulesLexerLogicalNot          = 22
	MongoRPCRulesLexerBinaryAnd           = 23
	MongoRPCRulesLexerBinaryOr            = 24
	MongoRPCRulesLexerArithmeticAdd       = 25
	MongoRPCRulesLexerArithmeticSub       = 26
	MongoRPCRulesLexerArithmeticMul       = 27
	MongoRPCRulesLexerArithmeticExp       = 28
	MongoRPCRulesLexerArithmeticModus     = 29
	MongoRPCRulesLexerInOperator          = 30
	MongoRPCRulesLexerAllow               = 31
	MongoRPCRulesLexerMatch               = 32
	MongoRPCRulesLexerIf                  = 33
	MongoRPCRulesLexerGet                 = 34
	MongoRPCRulesLexerExists              = 35
	MongoRPCRulesLexerTrue                = 36
	MongoRPCRulesLexerFalse               = 37
	MongoRPCRulesLexerList                = 38
	MongoRPCRulesLexerCreate              = 39
	MongoRPCRulesLexerUpdate              = 40
	MongoRPCRulesLexerRead                = 41
	MongoRPCRulesLexerWrite               = 42
	MongoRPCRulesLexerDelete              = 43
	MongoRPCRulesLexerFunction            = 44
	MongoRPCRulesLexerReturn              = 45
	MongoRPCRulesLexerNull                = 46
	MongoRPCRulesLexerService             = 47
	MongoRPCRulesLexerNumber              = 48
	MongoRPCRulesLexerStringExpression    = 49
	MongoRPCRulesLexerIdentifier          = 50
	MongoRPCRulesLexerSlash               = 51
	MongoRPCRulesLexerWS                  = 52
	MongoRPCRulesLexerComment             = 53
)

// Code generated from C:/Users/User/Laboratory/golang/src/github.com/newm4n/grule-rule-engine/antlr\grulev2.g4 by ANTLR 4.8. DO NOT EDIT.

package grulev2

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 43, 419,
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
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 30, 3, 30, 5, 30, 200, 10, 30, 3, 30, 6, 30, 203, 10,
	30, 13, 30, 14, 30, 204, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3,
	34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 41, 3, 41, 7, 41, 259, 10, 41, 12, 41, 14, 41, 262, 11, 41, 3,
	42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47,
	3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3,
	51, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 55, 3, 55,
	3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 3, 60, 3,
	61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64,
	3, 64, 7, 64, 318, 10, 64, 12, 64, 14, 64, 321, 11, 64, 3, 64, 3, 64, 3,
	65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 7, 65, 331, 10, 65, 12, 65, 14,
	65, 334, 11, 65, 3, 65, 3, 65, 3, 66, 6, 66, 339, 10, 66, 13, 66, 14, 66,
	340, 3, 67, 6, 67, 344, 10, 67, 13, 67, 14, 67, 345, 5, 67, 348, 10, 67,
	3, 67, 3, 67, 6, 67, 352, 10, 67, 13, 67, 14, 67, 353, 3, 67, 6, 67, 357,
	10, 67, 13, 67, 14, 67, 358, 3, 67, 3, 67, 3, 67, 3, 67, 6, 67, 365, 10,
	67, 13, 67, 14, 67, 366, 5, 67, 369, 10, 67, 3, 67, 3, 67, 6, 67, 373,
	10, 67, 13, 67, 14, 67, 374, 3, 67, 3, 67, 3, 67, 6, 67, 380, 10, 67, 13,
	67, 14, 67, 381, 3, 67, 3, 67, 5, 67, 386, 10, 67, 3, 68, 6, 68, 389, 10,
	68, 13, 68, 14, 68, 390, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 7, 69,
	399, 10, 69, 12, 69, 14, 69, 402, 11, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3,
	69, 3, 70, 3, 70, 3, 70, 3, 70, 7, 70, 413, 10, 70, 12, 70, 14, 70, 416,
	11, 70, 3, 70, 3, 70, 3, 400, 2, 71, 3, 3, 5, 2, 7, 2, 9, 2, 11, 2, 13,
	2, 15, 2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 27, 2, 29, 2, 31, 2, 33, 2,
	35, 2, 37, 2, 39, 2, 41, 2, 43, 2, 45, 2, 47, 2, 49, 2, 51, 2, 53, 2, 55,
	2, 57, 2, 59, 2, 61, 4, 63, 5, 65, 6, 67, 7, 69, 8, 71, 9, 73, 10, 75,
	11, 77, 12, 79, 13, 81, 14, 83, 15, 85, 16, 87, 17, 89, 18, 91, 19, 93,
	20, 95, 21, 97, 22, 99, 23, 101, 24, 103, 25, 105, 26, 107, 27, 109, 28,
	111, 29, 113, 30, 115, 31, 117, 32, 119, 33, 121, 34, 123, 35, 125, 36,
	127, 37, 129, 38, 131, 39, 133, 40, 135, 41, 137, 42, 139, 43, 3, 2, 35,
	3, 2, 50, 59, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69,
	101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72,
	104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75,
	107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78,
	110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81,
	113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84,
	116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87,
	119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90,
	122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 4, 2, 67, 92,
	99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 4, 2, 36, 36, 94, 94, 4, 2, 41,
	41, 94, 94, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 12, 12, 15, 15, 2, 414,
	2, 3, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2,
	2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2,
	2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3,
	2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89,
	3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2,
	97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2,
	2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111,
	3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2,
	2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3,
	2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2,
	133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2,
	2, 2, 3, 141, 3, 2, 2, 2, 5, 143, 3, 2, 2, 2, 7, 145, 3, 2, 2, 2, 9, 147,
	3, 2, 2, 2, 11, 149, 3, 2, 2, 2, 13, 151, 3, 2, 2, 2, 15, 153, 3, 2, 2,
	2, 17, 155, 3, 2, 2, 2, 19, 157, 3, 2, 2, 2, 21, 159, 3, 2, 2, 2, 23, 161,
	3, 2, 2, 2, 25, 163, 3, 2, 2, 2, 27, 165, 3, 2, 2, 2, 29, 167, 3, 2, 2,
	2, 31, 169, 3, 2, 2, 2, 33, 171, 3, 2, 2, 2, 35, 173, 3, 2, 2, 2, 37, 175,
	3, 2, 2, 2, 39, 177, 3, 2, 2, 2, 41, 179, 3, 2, 2, 2, 43, 181, 3, 2, 2,
	2, 45, 183, 3, 2, 2, 2, 47, 185, 3, 2, 2, 2, 49, 187, 3, 2, 2, 2, 51, 189,
	3, 2, 2, 2, 53, 191, 3, 2, 2, 2, 55, 193, 3, 2, 2, 2, 57, 195, 3, 2, 2,
	2, 59, 197, 3, 2, 2, 2, 61, 206, 3, 2, 2, 2, 63, 211, 3, 2, 2, 2, 65, 216,
	3, 2, 2, 2, 67, 221, 3, 2, 2, 2, 69, 224, 3, 2, 2, 2, 71, 227, 3, 2, 2,
	2, 73, 232, 3, 2, 2, 2, 75, 238, 3, 2, 2, 2, 77, 243, 3, 2, 2, 2, 79, 247,
	3, 2, 2, 2, 81, 256, 3, 2, 2, 2, 83, 263, 3, 2, 2, 2, 85, 265, 3, 2, 2,
	2, 87, 267, 3, 2, 2, 2, 89, 269, 3, 2, 2, 2, 91, 271, 3, 2, 2, 2, 93, 273,
	3, 2, 2, 2, 95, 276, 3, 2, 2, 2, 97, 278, 3, 2, 2, 2, 99, 280, 3, 2, 2,
	2, 101, 282, 3, 2, 2, 2, 103, 285, 3, 2, 2, 2, 105, 288, 3, 2, 2, 2, 107,
	291, 3, 2, 2, 2, 109, 293, 3, 2, 2, 2, 111, 295, 3, 2, 2, 2, 113, 297,
	3, 2, 2, 2, 115, 299, 3, 2, 2, 2, 117, 301, 3, 2, 2, 2, 119, 303, 3, 2,
	2, 2, 121, 305, 3, 2, 2, 2, 123, 307, 3, 2, 2, 2, 125, 309, 3, 2, 2, 2,
	127, 311, 3, 2, 2, 2, 129, 324, 3, 2, 2, 2, 131, 338, 3, 2, 2, 2, 133,
	385, 3, 2, 2, 2, 135, 388, 3, 2, 2, 2, 137, 394, 3, 2, 2, 2, 139, 408,
	3, 2, 2, 2, 141, 142, 7, 46, 2, 2, 142, 4, 3, 2, 2, 2, 143, 144, 9, 2,
	2, 2, 144, 6, 3, 2, 2, 2, 145, 146, 9, 3, 2, 2, 146, 8, 3, 2, 2, 2, 147,
	148, 9, 4, 2, 2, 148, 10, 3, 2, 2, 2, 149, 150, 9, 5, 2, 2, 150, 12, 3,
	2, 2, 2, 151, 152, 9, 6, 2, 2, 152, 14, 3, 2, 2, 2, 153, 154, 9, 7, 2,
	2, 154, 16, 3, 2, 2, 2, 155, 156, 9, 8, 2, 2, 156, 18, 3, 2, 2, 2, 157,
	158, 9, 9, 2, 2, 158, 20, 3, 2, 2, 2, 159, 160, 9, 10, 2, 2, 160, 22, 3,
	2, 2, 2, 161, 162, 9, 11, 2, 2, 162, 24, 3, 2, 2, 2, 163, 164, 9, 12, 2,
	2, 164, 26, 3, 2, 2, 2, 165, 166, 9, 13, 2, 2, 166, 28, 3, 2, 2, 2, 167,
	168, 9, 14, 2, 2, 168, 30, 3, 2, 2, 2, 169, 170, 9, 15, 2, 2, 170, 32,
	3, 2, 2, 2, 171, 172, 9, 16, 2, 2, 172, 34, 3, 2, 2, 2, 173, 174, 9, 17,
	2, 2, 174, 36, 3, 2, 2, 2, 175, 176, 9, 18, 2, 2, 176, 38, 3, 2, 2, 2,
	177, 178, 9, 19, 2, 2, 178, 40, 3, 2, 2, 2, 179, 180, 9, 20, 2, 2, 180,
	42, 3, 2, 2, 2, 181, 182, 9, 21, 2, 2, 182, 44, 3, 2, 2, 2, 183, 184, 9,
	22, 2, 2, 184, 46, 3, 2, 2, 2, 185, 186, 9, 23, 2, 2, 186, 48, 3, 2, 2,
	2, 187, 188, 9, 24, 2, 2, 188, 50, 3, 2, 2, 2, 189, 190, 9, 25, 2, 2, 190,
	52, 3, 2, 2, 2, 191, 192, 9, 26, 2, 2, 192, 54, 3, 2, 2, 2, 193, 194, 9,
	27, 2, 2, 194, 56, 3, 2, 2, 2, 195, 196, 9, 28, 2, 2, 196, 58, 3, 2, 2,
	2, 197, 199, 7, 71, 2, 2, 198, 200, 7, 47, 2, 2, 199, 198, 3, 2, 2, 2,
	199, 200, 3, 2, 2, 2, 200, 202, 3, 2, 2, 2, 201, 203, 5, 5, 3, 2, 202,
	201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204, 205,
	3, 2, 2, 2, 205, 60, 3, 2, 2, 2, 206, 207, 5, 41, 21, 2, 207, 208, 5, 47,
	24, 2, 208, 209, 5, 29, 15, 2, 209, 210, 5, 15, 8, 2, 210, 62, 3, 2, 2,
	2, 211, 212, 5, 51, 26, 2, 212, 213, 5, 21, 11, 2, 213, 214, 5, 15, 8,
	2, 214, 215, 5, 33, 17, 2, 215, 64, 3, 2, 2, 2, 216, 217, 5, 45, 23, 2,
	217, 218, 5, 21, 11, 2, 218, 219, 5, 15, 8, 2, 219, 220, 5, 33, 17, 2,
	220, 66, 3, 2, 2, 2, 221, 222, 7, 40, 2, 2, 222, 223, 7, 40, 2, 2, 223,
	68, 3, 2, 2, 2, 224, 225, 7, 126, 2, 2, 225, 226, 7, 126, 2, 2, 226, 70,
	3, 2, 2, 2, 227, 228, 5, 45, 23, 2, 228, 229, 5, 41, 21, 2, 229, 230, 5,
	47, 24, 2, 230, 231, 5, 15, 8, 2, 231, 72, 3, 2, 2, 2, 232, 233, 5, 17,
	9, 2, 233, 234, 5, 7, 4, 2, 234, 235, 5, 29, 15, 2, 235, 236, 5, 43, 22,
	2, 236, 237, 5, 15, 8, 2, 237, 74, 3, 2, 2, 2, 238, 239, 5, 33, 17, 2,
	239, 240, 5, 47, 24, 2, 240, 241, 5, 29, 15, 2, 241, 242, 5, 29, 15, 2,
	242, 76, 3, 2, 2, 2, 243, 244, 5, 33, 17, 2, 244, 245, 5, 35, 18, 2, 245,
	246, 5, 45, 23, 2, 246, 78, 3, 2, 2, 2, 247, 248, 5, 43, 22, 2, 248, 249,
	5, 7, 4, 2, 249, 250, 5, 29, 15, 2, 250, 251, 5, 23, 12, 2, 251, 252, 5,
	15, 8, 2, 252, 253, 5, 33, 17, 2, 253, 254, 5, 11, 6, 2, 254, 255, 5, 15,
	8, 2, 255, 80, 3, 2, 2, 2, 256, 260, 9, 29, 2, 2, 257, 259, 9, 30, 2, 2,
	258, 257, 3, 2, 2, 2, 259, 262, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 260,
	261, 3, 2, 2, 2, 261, 82, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 263, 264, 7,
	45, 2, 2, 264, 84, 3, 2, 2, 2, 265, 266, 7, 47, 2, 2, 266, 86, 3, 2, 2,
	2, 267, 268, 7, 49, 2, 2, 268, 88, 3, 2, 2, 2, 269, 270, 7, 44, 2, 2, 270,
	90, 3, 2, 2, 2, 271, 272, 7, 39, 2, 2, 272, 92, 3, 2, 2, 2, 273, 274, 7,
	63, 2, 2, 274, 275, 7, 63, 2, 2, 275, 94, 3, 2, 2, 2, 276, 277, 7, 63,
	2, 2, 277, 96, 3, 2, 2, 2, 278, 279, 7, 64, 2, 2, 279, 98, 3, 2, 2, 2,
	280, 281, 7, 62, 2, 2, 281, 100, 3, 2, 2, 2, 282, 283, 7, 64, 2, 2, 283,
	284, 7, 63, 2, 2, 284, 102, 3, 2, 2, 2, 285, 286, 7, 62, 2, 2, 286, 287,
	7, 63, 2, 2, 287, 104, 3, 2, 2, 2, 288, 289, 7, 35, 2, 2, 289, 290, 7,
	63, 2, 2, 290, 106, 3, 2, 2, 2, 291, 292, 7, 40, 2, 2, 292, 108, 3, 2,
	2, 2, 293, 294, 7, 126, 2, 2, 294, 110, 3, 2, 2, 2, 295, 296, 7, 61, 2,
	2, 296, 112, 3, 2, 2, 2, 297, 298, 7, 125, 2, 2, 298, 114, 3, 2, 2, 2,
	299, 300, 7, 127, 2, 2, 300, 116, 3, 2, 2, 2, 301, 302, 7, 42, 2, 2, 302,
	118, 3, 2, 2, 2, 303, 304, 7, 43, 2, 2, 304, 120, 3, 2, 2, 2, 305, 306,
	7, 93, 2, 2, 306, 122, 3, 2, 2, 2, 307, 308, 7, 95, 2, 2, 308, 124, 3,
	2, 2, 2, 309, 310, 7, 48, 2, 2, 310, 126, 3, 2, 2, 2, 311, 319, 7, 36,
	2, 2, 312, 313, 7, 94, 2, 2, 313, 318, 11, 2, 2, 2, 314, 315, 7, 36, 2,
	2, 315, 318, 7, 36, 2, 2, 316, 318, 10, 31, 2, 2, 317, 312, 3, 2, 2, 2,
	317, 314, 3, 2, 2, 2, 317, 316, 3, 2, 2, 2, 318, 321, 3, 2, 2, 2, 319,
	317, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 322, 3, 2, 2, 2, 321, 319,
	3, 2, 2, 2, 322, 323, 7, 36, 2, 2, 323, 128, 3, 2, 2, 2, 324, 332, 7, 41,
	2, 2, 325, 326, 7, 94, 2, 2, 326, 331, 11, 2, 2, 2, 327, 328, 7, 41, 2,
	2, 328, 331, 7, 41, 2, 2, 329, 331, 10, 32, 2, 2, 330, 325, 3, 2, 2, 2,
	330, 327, 3, 2, 2, 2, 330, 329, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2, 332,
	330, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 335, 3, 2, 2, 2, 334, 332,
	3, 2, 2, 2, 335, 336, 7, 41, 2, 2, 336, 130, 3, 2, 2, 2, 337, 339, 5, 5,
	3, 2, 338, 337, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 338, 3, 2, 2, 2,
	340, 341, 3, 2, 2, 2, 341, 132, 3, 2, 2, 2, 342, 344, 5, 5, 3, 2, 343,
	342, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 345, 346,
	3, 2, 2, 2, 346, 348, 3, 2, 2, 2, 347, 343, 3, 2, 2, 2, 347, 348, 3, 2,
	2, 2, 348, 349, 3, 2, 2, 2, 349, 351, 7, 48, 2, 2, 350, 352, 5, 5, 3, 2,
	351, 350, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 351, 3, 2, 2, 2, 353,
	354, 3, 2, 2, 2, 354, 386, 3, 2, 2, 2, 355, 357, 5, 5, 3, 2, 356, 355,
	3, 2, 2, 2, 357, 358, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 358, 359, 3, 2,
	2, 2, 359, 360, 3, 2, 2, 2, 360, 361, 7, 48, 2, 2, 361, 362, 5, 59, 30,
	2, 362, 386, 3, 2, 2, 2, 363, 365, 5, 5, 3, 2, 364, 363, 3, 2, 2, 2, 365,
	366, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 369,
	3, 2, 2, 2, 368, 364, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 370, 3, 2,
	2, 2, 370, 372, 7, 48, 2, 2, 371, 373, 5, 5, 3, 2, 372, 371, 3, 2, 2, 2,
	373, 374, 3, 2, 2, 2, 374, 372, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375,
	376, 3, 2, 2, 2, 376, 377, 5, 59, 30, 2, 377, 386, 3, 2, 2, 2, 378, 380,
	5, 5, 3, 2, 379, 378, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 379, 3, 2,
	2, 2, 381, 382, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 384, 5, 59, 30,
	2, 384, 386, 3, 2, 2, 2, 385, 347, 3, 2, 2, 2, 385, 356, 3, 2, 2, 2, 385,
	368, 3, 2, 2, 2, 385, 379, 3, 2, 2, 2, 386, 134, 3, 2, 2, 2, 387, 389,
	9, 33, 2, 2, 388, 387, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 388, 3, 2,
	2, 2, 390, 391, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 393, 8, 68, 2, 2,
	393, 136, 3, 2, 2, 2, 394, 395, 7, 49, 2, 2, 395, 396, 7, 44, 2, 2, 396,
	400, 3, 2, 2, 2, 397, 399, 11, 2, 2, 2, 398, 397, 3, 2, 2, 2, 399, 402,
	3, 2, 2, 2, 400, 401, 3, 2, 2, 2, 400, 398, 3, 2, 2, 2, 401, 403, 3, 2,
	2, 2, 402, 400, 3, 2, 2, 2, 403, 404, 7, 44, 2, 2, 404, 405, 7, 49, 2,
	2, 405, 406, 3, 2, 2, 2, 406, 407, 8, 69, 3, 2, 407, 138, 3, 2, 2, 2, 408,
	409, 7, 49, 2, 2, 409, 410, 7, 49, 2, 2, 410, 414, 3, 2, 2, 2, 411, 413,
	10, 34, 2, 2, 412, 411, 3, 2, 2, 2, 413, 416, 3, 2, 2, 2, 414, 412, 3,
	2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 417, 3, 2, 2, 2, 416, 414, 3, 2, 2,
	2, 417, 418, 8, 70, 4, 2, 418, 140, 3, 2, 2, 2, 23, 2, 199, 204, 260, 317,
	319, 330, 332, 340, 345, 347, 353, 358, 366, 368, 374, 381, 385, 390, 400,
	414, 5, 3, 68, 2, 3, 69, 3, 3, 70, 4,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "", "", "", "'&&'", "'||'", "", "", "", "", "", "", "'+'", "'-'",
	"'/'", "'*'", "'%'", "'=='", "'='", "'>'", "'<'", "'>='", "'<='", "'!='",
	"'&'", "'|'", "';'", "'{'", "'}'", "'('", "')'", "'['", "']'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "RULE", "WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NULL_LITERAL",
	"NOT", "SALIENCE", "SIMPLENAME", "PLUS", "MINUS", "DIV", "MUL", "MOD",
	"EQUALS", "ASSIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "BITAND", "BITOR",
	"SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET", "LS_BRACKET",
	"RS_BRACKET", "DOT", "DQUOTA_STRING", "SQUOTA_STRING", "DECIMAL_LITERAL",
	"REAL_LITERAL", "SPACE", "COMMENT", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "DEC_DIGIT", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y",
	"Z", "EXPONENT_NUM_PART", "RULE", "WHEN", "THEN", "AND", "OR", "TRUE",
	"FALSE", "NULL_LITERAL", "NOT", "SALIENCE", "SIMPLENAME", "PLUS", "MINUS",
	"DIV", "MUL", "MOD", "EQUALS", "ASSIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS",
	"BITAND", "BITOR", "SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET",
	"LS_BRACKET", "RS_BRACKET", "DOT", "DQUOTA_STRING", "SQUOTA_STRING", "DECIMAL_LITERAL",
	"REAL_LITERAL", "SPACE", "COMMENT", "LINE_COMMENT",
}

type grulev2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func Newgrulev2Lexer(input antlr.CharStream) *grulev2Lexer {

	l := new(grulev2Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "grulev2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// grulev2Lexer tokens.
const (
	grulev2LexerT__0            = 1
	grulev2LexerRULE            = 2
	grulev2LexerWHEN            = 3
	grulev2LexerTHEN            = 4
	grulev2LexerAND             = 5
	grulev2LexerOR              = 6
	grulev2LexerTRUE            = 7
	grulev2LexerFALSE           = 8
	grulev2LexerNULL_LITERAL    = 9
	grulev2LexerNOT             = 10
	grulev2LexerSALIENCE        = 11
	grulev2LexerSIMPLENAME      = 12
	grulev2LexerPLUS            = 13
	grulev2LexerMINUS           = 14
	grulev2LexerDIV             = 15
	grulev2LexerMUL             = 16
	grulev2LexerMOD             = 17
	grulev2LexerEQUALS          = 18
	grulev2LexerASSIGN          = 19
	grulev2LexerGT              = 20
	grulev2LexerLT              = 21
	grulev2LexerGTE             = 22
	grulev2LexerLTE             = 23
	grulev2LexerNOTEQUALS       = 24
	grulev2LexerBITAND          = 25
	grulev2LexerBITOR           = 26
	grulev2LexerSEMICOLON       = 27
	grulev2LexerLR_BRACE        = 28
	grulev2LexerRR_BRACE        = 29
	grulev2LexerLR_BRACKET      = 30
	grulev2LexerRR_BRACKET      = 31
	grulev2LexerLS_BRACKET      = 32
	grulev2LexerRS_BRACKET      = 33
	grulev2LexerDOT             = 34
	grulev2LexerDQUOTA_STRING   = 35
	grulev2LexerSQUOTA_STRING   = 36
	grulev2LexerDECIMAL_LITERAL = 37
	grulev2LexerREAL_LITERAL    = 38
	grulev2LexerSPACE           = 39
	grulev2LexerCOMMENT         = 40
	grulev2LexerLINE_COMMENT    = 41
)

func (l *grulev2Lexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 66:
		l.SPACE_Action(localctx, actionIndex)

	case 67:
		l.COMMENT_Action(localctx, actionIndex)

	case 68:
		l.LINE_COMMENT_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *grulev2Lexer) SPACE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		l.Skip()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *grulev2Lexer) COMMENT_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 1:
		l.Skip()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *grulev2Lexer) LINE_COMMENT_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 2:
		l.Skip()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
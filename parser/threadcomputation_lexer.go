// Code generated from ThreadComputation.g4 by ANTLR 4.9.3. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 26, 301,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 5, 13, 99, 10, 13,
	3, 13, 6, 13, 102, 10, 13, 13, 13, 14, 13, 103, 3, 14, 3, 14, 3, 15, 3,
	15, 7, 15, 110, 10, 15, 12, 15, 14, 15, 113, 11, 15, 3, 16, 5, 16, 116,
	10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 7, 17, 122, 10, 17, 12, 17, 14, 17,
	125, 11, 17, 3, 17, 6, 17, 128, 10, 17, 13, 17, 14, 17, 129, 5, 17, 132,
	10, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 5, 21, 155, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 5, 22, 175, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 181,
	10, 23, 12, 23, 14, 23, 184, 11, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 6, 24, 192, 10, 24, 13, 24, 14, 24, 193, 3, 24, 3, 24, 3, 25, 5,
	25, 199, 10, 25, 3, 25, 3, 25, 5, 25, 203, 10, 25, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 5, 29, 214, 10, 29, 3, 29, 6,
	29, 217, 10, 29, 13, 29, 14, 29, 218, 3, 29, 5, 29, 222, 10, 29, 3, 29,
	3, 29, 5, 29, 226, 10, 29, 3, 29, 6, 29, 229, 10, 29, 13, 29, 14, 29, 230,
	3, 29, 5, 29, 234, 10, 29, 3, 29, 5, 29, 237, 10, 29, 3, 30, 7, 30, 240,
	10, 30, 12, 30, 14, 30, 243, 11, 30, 3, 30, 3, 30, 6, 30, 247, 10, 30,
	13, 30, 14, 30, 248, 3, 30, 6, 30, 252, 10, 30, 13, 30, 14, 30, 253, 3,
	30, 5, 30, 257, 10, 30, 3, 31, 5, 31, 260, 10, 31, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 32, 3, 32, 5, 32, 268, 10, 32, 3, 32, 7, 32, 271, 10, 32, 12,
	32, 14, 32, 274, 11, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 281,
	10, 32, 3, 32, 7, 32, 284, 10, 32, 12, 32, 14, 32, 287, 11, 32, 3, 32,
	5, 32, 290, 10, 32, 3, 33, 5, 33, 293, 10, 33, 3, 33, 5, 33, 296, 10, 33,
	3, 34, 3, 34, 5, 34, 300, 10, 34, 3, 182, 2, 35, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29,
	16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47,
	25, 49, 26, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67,
	2, 3, 2, 14, 4, 2, 65, 65, 98, 98, 17, 2, 35, 35, 40, 40, 44, 47, 49, 49,
	61, 64, 96, 97, 126, 126, 128, 128, 174, 174, 217, 217, 249, 249, 8595,
	8595, 8597, 8597, 8712, 8713, 8745, 8746, 4, 2, 11, 11, 34, 34, 3, 2, 51,
	59, 3, 2, 50, 59, 4, 2, 45, 45, 47, 47, 4, 2, 71, 71, 103, 103, 6, 2, 12,
	12, 15, 15, 41, 41, 94, 94, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 4, 2,
	67, 92, 99, 124, 4, 2, 65, 65, 98, 124, 4, 2, 48, 48, 50, 59, 2, 327, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2,
	2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3,
	2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49,
	3, 2, 2, 2, 3, 69, 3, 2, 2, 2, 5, 71, 3, 2, 2, 2, 7, 73, 3, 2, 2, 2, 9,
	75, 3, 2, 2, 2, 11, 77, 3, 2, 2, 2, 13, 79, 3, 2, 2, 2, 15, 81, 3, 2, 2,
	2, 17, 83, 3, 2, 2, 2, 19, 91, 3, 2, 2, 2, 21, 93, 3, 2, 2, 2, 23, 95,
	3, 2, 2, 2, 25, 98, 3, 2, 2, 2, 27, 105, 3, 2, 2, 2, 29, 107, 3, 2, 2,
	2, 31, 115, 3, 2, 2, 2, 33, 131, 3, 2, 2, 2, 35, 133, 3, 2, 2, 2, 37, 135,
	3, 2, 2, 2, 39, 137, 3, 2, 2, 2, 41, 154, 3, 2, 2, 2, 43, 174, 3, 2, 2,
	2, 45, 176, 3, 2, 2, 2, 47, 191, 3, 2, 2, 2, 49, 202, 3, 2, 2, 2, 51, 206,
	3, 2, 2, 2, 53, 208, 3, 2, 2, 2, 55, 210, 3, 2, 2, 2, 57, 236, 3, 2, 2,
	2, 59, 256, 3, 2, 2, 2, 61, 259, 3, 2, 2, 2, 63, 289, 3, 2, 2, 2, 65, 295,
	3, 2, 2, 2, 67, 299, 3, 2, 2, 2, 69, 70, 7, 93, 2, 2, 70, 4, 3, 2, 2, 2,
	71, 72, 7, 95, 2, 2, 72, 6, 3, 2, 2, 2, 73, 74, 7, 66, 2, 2, 74, 8, 3,
	2, 2, 2, 75, 76, 7, 38, 2, 2, 76, 10, 3, 2, 2, 2, 77, 78, 7, 60, 2, 2,
	78, 12, 3, 2, 2, 2, 79, 80, 7, 42, 2, 2, 80, 14, 3, 2, 2, 2, 81, 82, 7,
	43, 2, 2, 82, 16, 3, 2, 2, 2, 83, 84, 7, 110, 2, 2, 84, 85, 7, 99, 2, 2,
	85, 86, 7, 111, 2, 2, 86, 87, 7, 100, 2, 2, 87, 88, 7, 102, 2, 2, 88, 89,
	7, 99, 2, 2, 89, 90, 7, 94, 2, 2, 90, 18, 3, 2, 2, 2, 91, 92, 7, 94, 2,
	2, 92, 20, 3, 2, 2, 2, 93, 94, 7, 125, 2, 2, 94, 22, 3, 2, 2, 2, 95, 96,
	7, 127, 2, 2, 96, 24, 3, 2, 2, 2, 97, 99, 5, 27, 14, 2, 98, 97, 3, 2, 2,
	2, 98, 99, 3, 2, 2, 2, 99, 101, 3, 2, 2, 2, 100, 102, 5, 39, 20, 2, 101,
	100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104,
	3, 2, 2, 2, 104, 26, 3, 2, 2, 2, 105, 106, 9, 2, 2, 2, 106, 28, 3, 2, 2,
	2, 107, 111, 5, 65, 33, 2, 108, 110, 5, 67, 34, 2, 109, 108, 3, 2, 2, 2,
	110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112,
	30, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114, 116, 5, 55, 28, 2, 115, 114,
	3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 118, 5, 33,
	17, 2, 118, 32, 3, 2, 2, 2, 119, 123, 5, 51, 26, 2, 120, 122, 5, 53, 27,
	2, 121, 120, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123,
	124, 3, 2, 2, 2, 124, 132, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 126, 128,
	7, 50, 2, 2, 127, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 127, 3, 2,
	2, 2, 129, 130, 3, 2, 2, 2, 130, 132, 3, 2, 2, 2, 131, 119, 3, 2, 2, 2,
	131, 127, 3, 2, 2, 2, 132, 34, 3, 2, 2, 2, 133, 134, 5, 57, 29, 2, 134,
	36, 3, 2, 2, 2, 135, 136, 5, 63, 32, 2, 136, 38, 3, 2, 2, 2, 137, 138,
	9, 3, 2, 2, 138, 40, 3, 2, 2, 2, 139, 140, 7, 37, 2, 2, 140, 141, 7, 118,
	2, 2, 141, 142, 7, 116, 2, 2, 142, 143, 7, 119, 2, 2, 143, 155, 7, 103,
	2, 2, 144, 145, 7, 37, 2, 2, 145, 146, 7, 86, 2, 2, 146, 147, 7, 116, 2,
	2, 147, 148, 7, 119, 2, 2, 148, 155, 7, 103, 2, 2, 149, 150, 7, 37, 2,
	2, 150, 151, 7, 86, 2, 2, 151, 152, 7, 84, 2, 2, 152, 153, 7, 87, 2, 2,
	153, 155, 7, 71, 2, 2, 154, 139, 3, 2, 2, 2, 154, 144, 3, 2, 2, 2, 154,
	149, 3, 2, 2, 2, 155, 42, 3, 2, 2, 2, 156, 157, 7, 37, 2, 2, 157, 158,
	7, 104, 2, 2, 158, 159, 7, 99, 2, 2, 159, 160, 7, 110, 2, 2, 160, 161,
	7, 117, 2, 2, 161, 175, 7, 103, 2, 2, 162, 163, 7, 37, 2, 2, 163, 164,
	7, 72, 2, 2, 164, 165, 7, 99, 2, 2, 165, 166, 7, 110, 2, 2, 166, 167, 7,
	117, 2, 2, 167, 175, 7, 103, 2, 2, 168, 169, 7, 37, 2, 2, 169, 170, 7,
	72, 2, 2, 170, 171, 7, 67, 2, 2, 171, 172, 7, 78, 2, 2, 172, 173, 7, 85,
	2, 2, 173, 175, 7, 71, 2, 2, 174, 156, 3, 2, 2, 2, 174, 162, 3, 2, 2, 2,
	174, 168, 3, 2, 2, 2, 175, 44, 3, 2, 2, 2, 176, 177, 7, 49, 2, 2, 177,
	178, 7, 44, 2, 2, 178, 182, 3, 2, 2, 2, 179, 181, 11, 2, 2, 2, 180, 179,
	3, 2, 2, 2, 181, 184, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 182, 180, 3, 2,
	2, 2, 183, 185, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 186, 7, 44, 2, 2,
	186, 187, 7, 49, 2, 2, 187, 188, 3, 2, 2, 2, 188, 189, 8, 23, 2, 2, 189,
	46, 3, 2, 2, 2, 190, 192, 9, 4, 2, 2, 191, 190, 3, 2, 2, 2, 192, 193, 3,
	2, 2, 2, 193, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 195, 3, 2, 2,
	2, 195, 196, 8, 24, 2, 2, 196, 48, 3, 2, 2, 2, 197, 199, 7, 15, 2, 2, 198,
	197, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 203,
	7, 12, 2, 2, 201, 203, 7, 15, 2, 2, 202, 198, 3, 2, 2, 2, 202, 201, 3,
	2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 205, 8, 25, 2, 2, 205, 50, 3, 2, 2,
	2, 206, 207, 9, 5, 2, 2, 207, 52, 3, 2, 2, 2, 208, 209, 9, 6, 2, 2, 209,
	54, 3, 2, 2, 2, 210, 211, 9, 7, 2, 2, 211, 56, 3, 2, 2, 2, 212, 214, 5,
	55, 28, 2, 213, 212, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 221, 3, 2,
	2, 2, 215, 217, 9, 6, 2, 2, 216, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2,
	218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 222, 3, 2, 2, 2, 220,
	222, 5, 59, 30, 2, 221, 216, 3, 2, 2, 2, 221, 220, 3, 2, 2, 2, 222, 223,
	3, 2, 2, 2, 223, 225, 9, 8, 2, 2, 224, 226, 9, 7, 2, 2, 225, 224, 3, 2,
	2, 2, 225, 226, 3, 2, 2, 2, 226, 228, 3, 2, 2, 2, 227, 229, 9, 6, 2, 2,
	228, 227, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 230,
	231, 3, 2, 2, 2, 231, 237, 3, 2, 2, 2, 232, 234, 5, 55, 28, 2, 233, 232,
	3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 237, 5, 59,
	30, 2, 236, 213, 3, 2, 2, 2, 236, 233, 3, 2, 2, 2, 237, 58, 3, 2, 2, 2,
	238, 240, 9, 6, 2, 2, 239, 238, 3, 2, 2, 2, 240, 243, 3, 2, 2, 2, 241,
	239, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 244, 3, 2, 2, 2, 243, 241,
	3, 2, 2, 2, 244, 246, 7, 48, 2, 2, 245, 247, 9, 6, 2, 2, 246, 245, 3, 2,
	2, 2, 247, 248, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2,
	249, 257, 3, 2, 2, 2, 250, 252, 9, 6, 2, 2, 251, 250, 3, 2, 2, 2, 252,
	253, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 255,
	3, 2, 2, 2, 255, 257, 7, 48, 2, 2, 256, 241, 3, 2, 2, 2, 256, 251, 3, 2,
	2, 2, 257, 60, 3, 2, 2, 2, 258, 260, 7, 15, 2, 2, 259, 258, 3, 2, 2, 2,
	259, 260, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 262, 7, 12, 2, 2, 262,
	62, 3, 2, 2, 2, 263, 272, 7, 41, 2, 2, 264, 267, 7, 94, 2, 2, 265, 268,
	5, 61, 31, 2, 266, 268, 11, 2, 2, 2, 267, 265, 3, 2, 2, 2, 267, 266, 3,
	2, 2, 2, 268, 271, 3, 2, 2, 2, 269, 271, 10, 9, 2, 2, 270, 264, 3, 2, 2,
	2, 270, 269, 3, 2, 2, 2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272,
	273, 3, 2, 2, 2, 273, 275, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275, 290,
	7, 41, 2, 2, 276, 285, 7, 36, 2, 2, 277, 280, 7, 94, 2, 2, 278, 281, 5,
	61, 31, 2, 279, 281, 11, 2, 2, 2, 280, 278, 3, 2, 2, 2, 280, 279, 3, 2,
	2, 2, 281, 284, 3, 2, 2, 2, 282, 284, 10, 10, 2, 2, 283, 277, 3, 2, 2,
	2, 283, 282, 3, 2, 2, 2, 284, 287, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 285,
	286, 3, 2, 2, 2, 286, 288, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 288, 290,
	7, 36, 2, 2, 289, 263, 3, 2, 2, 2, 289, 276, 3, 2, 2, 2, 290, 64, 3, 2,
	2, 2, 291, 293, 9, 11, 2, 2, 292, 291, 3, 2, 2, 2, 293, 296, 3, 2, 2, 2,
	294, 296, 9, 12, 2, 2, 295, 292, 3, 2, 2, 2, 295, 294, 3, 2, 2, 2, 296,
	66, 3, 2, 2, 2, 297, 300, 5, 65, 33, 2, 298, 300, 9, 13, 2, 2, 299, 297,
	3, 2, 2, 2, 299, 298, 3, 2, 2, 2, 300, 68, 3, 2, 2, 2, 38, 2, 98, 103,
	111, 115, 123, 129, 131, 154, 174, 182, 193, 198, 202, 213, 218, 221, 225,
	230, 233, 236, 241, 248, 253, 256, 259, 267, 270, 272, 280, 283, 285, 289,
	292, 295, 299, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "']'", "'@'", "'$'", "':'", "'('", "')'", "'lambda\\'", "'\\'",
	"'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "OPS", "OPS_START", "NAME",
	"INTEGER", "DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING", "OP", "TRUE", "FALSE",
	"BLOCK_COMMENT", "WhiteSpace", "NewLine",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "OPS", "OPS_START", "NAME", "INTEGER", "DECIMAL_INTEGER",
	"FLOAT_NUMBER", "STRING", "OP", "TRUE", "FALSE", "BLOCK_COMMENT", "WhiteSpace",
	"NewLine", "NON_ZERO_DIGIT", "DIGIT", "SIGN", "EXPONENT_OR_POINT_FLOAT",
	"POINT_FLOAT", "RN", "SHORT_STRING", "ID_START", "ID_CONTINUE",
}

type ThreadComputationLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewThreadComputationLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ThreadComputationLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewThreadComputationLexer(input antlr.CharStream) *ThreadComputationLexer {
	l := new(ThreadComputationLexer)
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
	l.GrammarFileName = "ThreadComputation.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ThreadComputationLexer tokens.
const (
	ThreadComputationLexerT__0            = 1
	ThreadComputationLexerT__1            = 2
	ThreadComputationLexerT__2            = 3
	ThreadComputationLexerT__3            = 4
	ThreadComputationLexerT__4            = 5
	ThreadComputationLexerT__5            = 6
	ThreadComputationLexerT__6            = 7
	ThreadComputationLexerT__7            = 8
	ThreadComputationLexerT__8            = 9
	ThreadComputationLexerT__9            = 10
	ThreadComputationLexerT__10           = 11
	ThreadComputationLexerOPS             = 12
	ThreadComputationLexerOPS_START       = 13
	ThreadComputationLexerNAME            = 14
	ThreadComputationLexerINTEGER         = 15
	ThreadComputationLexerDECIMAL_INTEGER = 16
	ThreadComputationLexerFLOAT_NUMBER    = 17
	ThreadComputationLexerSTRING          = 18
	ThreadComputationLexerOP              = 19
	ThreadComputationLexerTRUE            = 20
	ThreadComputationLexerFALSE           = 21
	ThreadComputationLexerBLOCK_COMMENT   = 22
	ThreadComputationLexerWhiteSpace      = 23
	ThreadComputationLexerNewLine         = 24
)

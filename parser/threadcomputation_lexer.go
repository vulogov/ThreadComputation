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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 328,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 16, 5, 16, 126, 10, 16, 3, 16, 6, 16, 129, 10, 16, 13,
	16, 14, 16, 130, 3, 17, 3, 17, 3, 18, 3, 18, 7, 18, 137, 10, 18, 12, 18,
	14, 18, 140, 11, 18, 3, 19, 5, 19, 143, 10, 19, 3, 19, 3, 19, 3, 20, 3,
	20, 7, 20, 149, 10, 20, 12, 20, 14, 20, 152, 11, 20, 3, 20, 6, 20, 155,
	10, 20, 13, 20, 14, 20, 156, 5, 20, 159, 10, 20, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 182, 10, 24, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 202, 10, 25, 3,
	26, 3, 26, 3, 26, 3, 26, 7, 26, 208, 10, 26, 12, 26, 14, 26, 211, 11, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 6, 27, 219, 10, 27, 13, 27, 14,
	27, 220, 3, 27, 3, 27, 3, 28, 5, 28, 226, 10, 28, 3, 28, 3, 28, 5, 28,
	230, 10, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3,
	32, 5, 32, 241, 10, 32, 3, 32, 6, 32, 244, 10, 32, 13, 32, 14, 32, 245,
	3, 32, 5, 32, 249, 10, 32, 3, 32, 3, 32, 5, 32, 253, 10, 32, 3, 32, 6,
	32, 256, 10, 32, 13, 32, 14, 32, 257, 3, 32, 5, 32, 261, 10, 32, 3, 32,
	5, 32, 264, 10, 32, 3, 33, 7, 33, 267, 10, 33, 12, 33, 14, 33, 270, 11,
	33, 3, 33, 3, 33, 6, 33, 274, 10, 33, 13, 33, 14, 33, 275, 3, 33, 6, 33,
	279, 10, 33, 13, 33, 14, 33, 280, 3, 33, 5, 33, 284, 10, 33, 3, 34, 5,
	34, 287, 10, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 295,
	10, 35, 3, 35, 7, 35, 298, 10, 35, 12, 35, 14, 35, 301, 11, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 308, 10, 35, 3, 35, 7, 35, 311, 10,
	35, 12, 35, 14, 35, 314, 11, 35, 3, 35, 5, 35, 317, 10, 35, 3, 36, 5, 36,
	320, 10, 36, 3, 36, 5, 36, 323, 10, 36, 3, 37, 3, 37, 5, 37, 327, 10, 37,
	3, 209, 2, 38, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	29, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67, 2, 69, 2, 71, 2, 73, 2, 3, 2,
	14, 4, 2, 65, 65, 98, 98, 17, 2, 35, 35, 40, 40, 44, 47, 49, 49, 61, 64,
	96, 97, 126, 126, 128, 128, 174, 174, 217, 217, 249, 249, 8595, 8595, 8597,
	8597, 8712, 8713, 8745, 8746, 4, 2, 11, 11, 34, 34, 3, 2, 51, 59, 3, 2,
	50, 59, 4, 2, 45, 45, 47, 47, 4, 2, 71, 71, 103, 103, 6, 2, 12, 12, 15,
	15, 41, 41, 94, 94, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 4, 2, 67, 92,
	99, 124, 4, 2, 65, 65, 98, 124, 4, 2, 48, 48, 50, 59, 2, 354, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2,
	2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 3, 75, 3,
	2, 2, 2, 5, 77, 3, 2, 2, 2, 7, 79, 3, 2, 2, 2, 9, 81, 3, 2, 2, 2, 11, 83,
	3, 2, 2, 2, 13, 85, 3, 2, 2, 2, 15, 87, 3, 2, 2, 2, 17, 89, 3, 2, 2, 2,
	19, 97, 3, 2, 2, 2, 21, 99, 3, 2, 2, 2, 23, 105, 3, 2, 2, 2, 25, 112, 3,
	2, 2, 2, 27, 120, 3, 2, 2, 2, 29, 122, 3, 2, 2, 2, 31, 125, 3, 2, 2, 2,
	33, 132, 3, 2, 2, 2, 35, 134, 3, 2, 2, 2, 37, 142, 3, 2, 2, 2, 39, 158,
	3, 2, 2, 2, 41, 160, 3, 2, 2, 2, 43, 162, 3, 2, 2, 2, 45, 164, 3, 2, 2,
	2, 47, 181, 3, 2, 2, 2, 49, 201, 3, 2, 2, 2, 51, 203, 3, 2, 2, 2, 53, 218,
	3, 2, 2, 2, 55, 229, 3, 2, 2, 2, 57, 233, 3, 2, 2, 2, 59, 235, 3, 2, 2,
	2, 61, 237, 3, 2, 2, 2, 63, 263, 3, 2, 2, 2, 65, 283, 3, 2, 2, 2, 67, 286,
	3, 2, 2, 2, 69, 316, 3, 2, 2, 2, 71, 322, 3, 2, 2, 2, 73, 326, 3, 2, 2,
	2, 75, 76, 7, 93, 2, 2, 76, 4, 3, 2, 2, 2, 77, 78, 7, 95, 2, 2, 78, 6,
	3, 2, 2, 2, 79, 80, 7, 66, 2, 2, 80, 8, 3, 2, 2, 2, 81, 82, 7, 38, 2, 2,
	82, 10, 3, 2, 2, 2, 83, 84, 7, 60, 2, 2, 84, 12, 3, 2, 2, 2, 85, 86, 7,
	42, 2, 2, 86, 14, 3, 2, 2, 2, 87, 88, 7, 43, 2, 2, 88, 16, 3, 2, 2, 2,
	89, 90, 7, 110, 2, 2, 90, 91, 7, 99, 2, 2, 91, 92, 7, 111, 2, 2, 92, 93,
	7, 100, 2, 2, 93, 94, 7, 102, 2, 2, 94, 95, 7, 99, 2, 2, 95, 96, 7, 94,
	2, 2, 96, 18, 3, 2, 2, 2, 97, 98, 7, 94, 2, 2, 98, 20, 3, 2, 2, 2, 99,
	100, 7, 118, 2, 2, 100, 101, 7, 116, 2, 2, 101, 102, 7, 119, 2, 2, 102,
	103, 7, 103, 2, 2, 103, 104, 7, 94, 2, 2, 104, 22, 3, 2, 2, 2, 105, 106,
	7, 104, 2, 2, 106, 107, 7, 99, 2, 2, 107, 108, 7, 110, 2, 2, 108, 109,
	7, 117, 2, 2, 109, 110, 7, 103, 2, 2, 110, 111, 7, 94, 2, 2, 111, 24, 3,
	2, 2, 2, 112, 113, 7, 104, 2, 2, 113, 114, 7, 107, 2, 2, 114, 115, 7, 110,
	2, 2, 115, 116, 7, 118, 2, 2, 116, 117, 7, 103, 2, 2, 117, 118, 7, 116,
	2, 2, 118, 119, 7, 94, 2, 2, 119, 26, 3, 2, 2, 2, 120, 121, 7, 125, 2,
	2, 121, 28, 3, 2, 2, 2, 122, 123, 7, 127, 2, 2, 123, 30, 3, 2, 2, 2, 124,
	126, 5, 33, 17, 2, 125, 124, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 128,
	3, 2, 2, 2, 127, 129, 5, 45, 23, 2, 128, 127, 3, 2, 2, 2, 129, 130, 3,
	2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 32, 3, 2, 2,
	2, 132, 133, 9, 2, 2, 2, 133, 34, 3, 2, 2, 2, 134, 138, 5, 71, 36, 2, 135,
	137, 5, 73, 37, 2, 136, 135, 3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138, 136,
	3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 36, 3, 2, 2, 2, 140, 138, 3, 2,
	2, 2, 141, 143, 5, 61, 31, 2, 142, 141, 3, 2, 2, 2, 142, 143, 3, 2, 2,
	2, 143, 144, 3, 2, 2, 2, 144, 145, 5, 39, 20, 2, 145, 38, 3, 2, 2, 2, 146,
	150, 5, 57, 29, 2, 147, 149, 5, 59, 30, 2, 148, 147, 3, 2, 2, 2, 149, 152,
	3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 159, 3, 2,
	2, 2, 152, 150, 3, 2, 2, 2, 153, 155, 7, 50, 2, 2, 154, 153, 3, 2, 2, 2,
	155, 156, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157,
	159, 3, 2, 2, 2, 158, 146, 3, 2, 2, 2, 158, 154, 3, 2, 2, 2, 159, 40, 3,
	2, 2, 2, 160, 161, 5, 63, 32, 2, 161, 42, 3, 2, 2, 2, 162, 163, 5, 69,
	35, 2, 163, 44, 3, 2, 2, 2, 164, 165, 9, 3, 2, 2, 165, 46, 3, 2, 2, 2,
	166, 167, 7, 37, 2, 2, 167, 168, 7, 118, 2, 2, 168, 169, 7, 116, 2, 2,
	169, 170, 7, 119, 2, 2, 170, 182, 7, 103, 2, 2, 171, 172, 7, 37, 2, 2,
	172, 173, 7, 86, 2, 2, 173, 174, 7, 116, 2, 2, 174, 175, 7, 119, 2, 2,
	175, 182, 7, 103, 2, 2, 176, 177, 7, 37, 2, 2, 177, 178, 7, 86, 2, 2, 178,
	179, 7, 84, 2, 2, 179, 180, 7, 87, 2, 2, 180, 182, 7, 71, 2, 2, 181, 166,
	3, 2, 2, 2, 181, 171, 3, 2, 2, 2, 181, 176, 3, 2, 2, 2, 182, 48, 3, 2,
	2, 2, 183, 184, 7, 37, 2, 2, 184, 185, 7, 104, 2, 2, 185, 186, 7, 99, 2,
	2, 186, 187, 7, 110, 2, 2, 187, 188, 7, 117, 2, 2, 188, 202, 7, 103, 2,
	2, 189, 190, 7, 37, 2, 2, 190, 191, 7, 72, 2, 2, 191, 192, 7, 99, 2, 2,
	192, 193, 7, 110, 2, 2, 193, 194, 7, 117, 2, 2, 194, 202, 7, 103, 2, 2,
	195, 196, 7, 37, 2, 2, 196, 197, 7, 72, 2, 2, 197, 198, 7, 67, 2, 2, 198,
	199, 7, 78, 2, 2, 199, 200, 7, 85, 2, 2, 200, 202, 7, 71, 2, 2, 201, 183,
	3, 2, 2, 2, 201, 189, 3, 2, 2, 2, 201, 195, 3, 2, 2, 2, 202, 50, 3, 2,
	2, 2, 203, 204, 7, 49, 2, 2, 204, 205, 7, 44, 2, 2, 205, 209, 3, 2, 2,
	2, 206, 208, 11, 2, 2, 2, 207, 206, 3, 2, 2, 2, 208, 211, 3, 2, 2, 2, 209,
	210, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 210, 212, 3, 2, 2, 2, 211, 209,
	3, 2, 2, 2, 212, 213, 7, 44, 2, 2, 213, 214, 7, 49, 2, 2, 214, 215, 3,
	2, 2, 2, 215, 216, 8, 26, 2, 2, 216, 52, 3, 2, 2, 2, 217, 219, 9, 4, 2,
	2, 218, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220,
	221, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 223, 8, 27, 2, 2, 223, 54,
	3, 2, 2, 2, 224, 226, 7, 15, 2, 2, 225, 224, 3, 2, 2, 2, 225, 226, 3, 2,
	2, 2, 226, 227, 3, 2, 2, 2, 227, 230, 7, 12, 2, 2, 228, 230, 7, 15, 2,
	2, 229, 225, 3, 2, 2, 2, 229, 228, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231,
	232, 8, 28, 2, 2, 232, 56, 3, 2, 2, 2, 233, 234, 9, 5, 2, 2, 234, 58, 3,
	2, 2, 2, 235, 236, 9, 6, 2, 2, 236, 60, 3, 2, 2, 2, 237, 238, 9, 7, 2,
	2, 238, 62, 3, 2, 2, 2, 239, 241, 5, 61, 31, 2, 240, 239, 3, 2, 2, 2, 240,
	241, 3, 2, 2, 2, 241, 248, 3, 2, 2, 2, 242, 244, 9, 6, 2, 2, 243, 242,
	3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2,
	2, 2, 246, 249, 3, 2, 2, 2, 247, 249, 5, 65, 33, 2, 248, 243, 3, 2, 2,
	2, 248, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 252, 9, 8, 2, 2, 251,
	253, 9, 7, 2, 2, 252, 251, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 255,
	3, 2, 2, 2, 254, 256, 9, 6, 2, 2, 255, 254, 3, 2, 2, 2, 256, 257, 3, 2,
	2, 2, 257, 255, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 264, 3, 2, 2, 2,
	259, 261, 5, 61, 31, 2, 260, 259, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261,
	262, 3, 2, 2, 2, 262, 264, 5, 65, 33, 2, 263, 240, 3, 2, 2, 2, 263, 260,
	3, 2, 2, 2, 264, 64, 3, 2, 2, 2, 265, 267, 9, 6, 2, 2, 266, 265, 3, 2,
	2, 2, 267, 270, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2,
	269, 271, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 271, 273, 7, 48, 2, 2, 272,
	274, 9, 6, 2, 2, 273, 272, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 273,
	3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 284, 3, 2, 2, 2, 277, 279, 9, 6,
	2, 2, 278, 277, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2,
	280, 281, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 284, 7, 48, 2, 2, 283,
	268, 3, 2, 2, 2, 283, 278, 3, 2, 2, 2, 284, 66, 3, 2, 2, 2, 285, 287, 7,
	15, 2, 2, 286, 285, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 288, 3, 2, 2,
	2, 288, 289, 7, 12, 2, 2, 289, 68, 3, 2, 2, 2, 290, 299, 7, 41, 2, 2, 291,
	294, 7, 94, 2, 2, 292, 295, 5, 67, 34, 2, 293, 295, 11, 2, 2, 2, 294, 292,
	3, 2, 2, 2, 294, 293, 3, 2, 2, 2, 295, 298, 3, 2, 2, 2, 296, 298, 10, 9,
	2, 2, 297, 291, 3, 2, 2, 2, 297, 296, 3, 2, 2, 2, 298, 301, 3, 2, 2, 2,
	299, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 302, 3, 2, 2, 2, 301,
	299, 3, 2, 2, 2, 302, 317, 7, 41, 2, 2, 303, 312, 7, 36, 2, 2, 304, 307,
	7, 94, 2, 2, 305, 308, 5, 67, 34, 2, 306, 308, 11, 2, 2, 2, 307, 305, 3,
	2, 2, 2, 307, 306, 3, 2, 2, 2, 308, 311, 3, 2, 2, 2, 309, 311, 10, 10,
	2, 2, 310, 304, 3, 2, 2, 2, 310, 309, 3, 2, 2, 2, 311, 314, 3, 2, 2, 2,
	312, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 315, 3, 2, 2, 2, 314,
	312, 3, 2, 2, 2, 315, 317, 7, 36, 2, 2, 316, 290, 3, 2, 2, 2, 316, 303,
	3, 2, 2, 2, 317, 70, 3, 2, 2, 2, 318, 320, 9, 11, 2, 2, 319, 318, 3, 2,
	2, 2, 320, 323, 3, 2, 2, 2, 321, 323, 9, 12, 2, 2, 322, 319, 3, 2, 2, 2,
	322, 321, 3, 2, 2, 2, 323, 72, 3, 2, 2, 2, 324, 327, 5, 71, 36, 2, 325,
	327, 9, 13, 2, 2, 326, 324, 3, 2, 2, 2, 326, 325, 3, 2, 2, 2, 327, 74,
	3, 2, 2, 2, 38, 2, 125, 130, 138, 142, 150, 156, 158, 181, 201, 209, 220,
	225, 229, 240, 245, 248, 252, 257, 260, 263, 268, 275, 280, 283, 286, 294,
	297, 299, 307, 310, 312, 316, 319, 322, 326, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "']'", "'@'", "'$'", "':'", "'('", "')'", "'lambda\\'", "'\\'",
	"'true\\'", "'false\\'", "'filter\\'", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "OPS", "OPS_START",
	"NAME", "INTEGER", "DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING", "OP", "TRUE",
	"FALSE", "BLOCK_COMMENT", "WhiteSpace", "NewLine",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "OPS", "OPS_START", "NAME",
	"INTEGER", "DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING", "OP", "TRUE", "FALSE",
	"BLOCK_COMMENT", "WhiteSpace", "NewLine", "NON_ZERO_DIGIT", "DIGIT", "SIGN",
	"EXPONENT_OR_POINT_FLOAT", "POINT_FLOAT", "RN", "SHORT_STRING", "ID_START",
	"ID_CONTINUE",
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
	ThreadComputationLexerT__11           = 12
	ThreadComputationLexerT__12           = 13
	ThreadComputationLexerT__13           = 14
	ThreadComputationLexerOPS             = 15
	ThreadComputationLexerOPS_START       = 16
	ThreadComputationLexerNAME            = 17
	ThreadComputationLexerINTEGER         = 18
	ThreadComputationLexerDECIMAL_INTEGER = 19
	ThreadComputationLexerFLOAT_NUMBER    = 20
	ThreadComputationLexerSTRING          = 21
	ThreadComputationLexerOP              = 22
	ThreadComputationLexerTRUE            = 23
	ThreadComputationLexerFALSE           = 24
	ThreadComputationLexerBLOCK_COMMENT   = 25
	ThreadComputationLexerWhiteSpace      = 26
	ThreadComputationLexerNewLine         = 27
)

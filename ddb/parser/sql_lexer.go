// Code generated from Sql.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 55, 470, 
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
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4, 
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 
	4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 
	10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 
	3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 
	18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 
	3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 
	22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 
	3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 
	26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 
	3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 
	30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 
	3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 
	38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 
	3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 
	43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 
	3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 
	53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 
	3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 
	64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 
	3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74, 3, 
	74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77, 3, 77, 3, 78, 6, 78, 424, 10, 78, 
	13, 78, 14, 78, 425, 3, 79, 3, 79, 7, 79, 430, 10, 79, 12, 79, 14, 79, 
	433, 11, 79, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 7, 
	80, 443, 10, 80, 12, 80, 14, 80, 446, 11, 80, 3, 80, 3, 80, 3, 81, 3, 81, 
	6, 81, 452, 10, 81, 13, 81, 14, 81, 453, 3, 81, 3, 81, 3, 82, 3, 82, 6, 
	82, 460, 10, 82, 13, 82, 14, 82, 461, 3, 83, 6, 83, 465, 10, 83, 13, 83, 
	14, 83, 466, 3, 83, 3, 83, 2, 2, 84, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 
	87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 2, 99, 2, 101, 2, 103, 2, 105, 
	2, 107, 2, 109, 2, 111, 2, 113, 2, 115, 2, 117, 2, 119, 2, 121, 2, 123, 
	2, 125, 2, 127, 2, 129, 2, 131, 2, 133, 2, 135, 2, 137, 2, 139, 2, 141, 
	2, 143, 2, 145, 2, 147, 2, 149, 2, 151, 2, 153, 2, 155, 50, 157, 51, 159, 
	52, 161, 53, 163, 54, 165, 55, 3, 2, 37, 4, 2, 67, 67, 99, 99, 4, 2, 68, 
	68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 
	71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 
	74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 
	77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 
	80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 
	83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 
	86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 
	89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 
	92, 124, 124, 3, 2, 50, 59, 4, 2, 50, 59, 67, 72, 4, 2, 67, 92, 99, 124, 
	5, 2, 67, 92, 99, 124, 19970, 40871, 8, 2, 48, 48, 50, 59, 67, 92, 99, 
	124, 19970, 40871, 65290, 65291, 3, 2, 41, 41, 6, 2, 11, 12, 15, 15, 34, 
	34, 36, 36, 6, 2, 11, 12, 15, 15, 34, 34, 60, 60, 5, 2, 11, 12, 15, 15, 
	34, 34, 2, 449, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 
	3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 
	55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 
	2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 
	2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 
	2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 
	2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 
	3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 
	2, 159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2, 165, 3, 
	2, 2, 2, 3, 167, 3, 2, 2, 2, 5, 169, 3, 2, 2, 2, 7, 171, 3, 2, 2, 2, 9, 
	173, 3, 2, 2, 2, 11, 175, 3, 2, 2, 2, 13, 177, 3, 2, 2, 2, 15, 180, 3, 
	2, 2, 2, 17, 182, 3, 2, 2, 2, 19, 184, 3, 2, 2, 2, 21, 186, 3, 2, 2, 2, 
	23, 188, 3, 2, 2, 2, 25, 190, 3, 2, 2, 2, 27, 193, 3, 2, 2, 2, 29, 196, 
	3, 2, 2, 2, 31, 199, 3, 2, 2, 2, 33, 201, 3, 2, 2, 2, 35, 204, 3, 2, 2, 
	2, 37, 207, 3, 2, 2, 2, 39, 214, 3, 2, 2, 2, 41, 219, 3, 2, 2, 2, 43, 223, 
	3, 2, 2, 2, 45, 227, 3, 2, 2, 2, 47, 231, 3, 2, 2, 2, 49, 235, 3, 2, 2, 
	2, 51, 241, 3, 2, 2, 2, 53, 250, 3, 2, 2, 2, 55, 256, 3, 2, 2, 2, 57, 262, 
	3, 2, 2, 2, 59, 265, 3, 2, 2, 2, 61, 271, 3, 2, 2, 2, 63, 277, 3, 2, 2, 
	2, 65, 283, 3, 2, 2, 2, 67, 288, 3, 2, 2, 2, 69, 295, 3, 2, 2, 2, 71, 299, 
	3, 2, 2, 2, 73, 302, 3, 2, 2, 2, 75, 310, 3, 2, 2, 2, 77, 314, 3, 2, 2, 
	2, 79, 317, 3, 2, 2, 2, 81, 322, 3, 2, 2, 2, 83, 325, 3, 2, 2, 2, 85, 329, 
	3, 2, 2, 2, 87, 334, 3, 2, 2, 2, 89, 340, 3, 2, 2, 2, 91, 347, 3, 2, 2, 
	2, 93, 352, 3, 2, 2, 2, 95, 359, 3, 2, 2, 2, 97, 364, 3, 2, 2, 2, 99, 366, 
	3, 2, 2, 2, 101, 368, 3, 2, 2, 2, 103, 370, 3, 2, 2, 2, 105, 372, 3, 2, 
	2, 2, 107, 374, 3, 2, 2, 2, 109, 376, 3, 2, 2, 2, 111, 378, 3, 2, 2, 2, 
	113, 380, 3, 2, 2, 2, 115, 382, 3, 2, 2, 2, 117, 384, 3, 2, 2, 2, 119, 
	386, 3, 2, 2, 2, 121, 388, 3, 2, 2, 2, 123, 390, 3, 2, 2, 2, 125, 392, 
	3, 2, 2, 2, 127, 394, 3, 2, 2, 2, 129, 396, 3, 2, 2, 2, 131, 398, 3, 2, 
	2, 2, 133, 400, 3, 2, 2, 2, 135, 402, 3, 2, 2, 2, 137, 404, 3, 2, 2, 2, 
	139, 406, 3, 2, 2, 2, 141, 408, 3, 2, 2, 2, 143, 410, 3, 2, 2, 2, 145, 
	412, 3, 2, 2, 2, 147, 414, 3, 2, 2, 2, 149, 416, 3, 2, 2, 2, 151, 418, 
	3, 2, 2, 2, 153, 420, 3, 2, 2, 2, 155, 423, 3, 2, 2, 2, 157, 427, 3, 2, 
	2, 2, 159, 434, 3, 2, 2, 2, 161, 449, 3, 2, 2, 2, 163, 457, 3, 2, 2, 2, 
	165, 464, 3, 2, 2, 2, 167, 168, 7, 44, 2, 2, 168, 4, 3, 2, 2, 2, 169, 170, 
	7, 46, 2, 2, 170, 6, 3, 2, 2, 2, 171, 172, 7, 49, 2, 2, 172, 8, 3, 2, 2, 
	2, 173, 174, 7, 45, 2, 2, 174, 10, 3, 2, 2, 2, 175, 176, 7, 47, 2, 2, 176, 
	12, 3, 2, 2, 2, 177, 178, 7, 126, 2, 2, 178, 179, 7, 126, 2, 2, 179, 14, 
	3, 2, 2, 2, 180, 181, 7, 42, 2, 2, 181, 16, 3, 2, 2, 2, 182, 183, 7, 43, 
	2, 2, 183, 18, 3, 2, 2, 2, 184, 185, 7, 63, 2, 2, 185, 20, 3, 2, 2, 2, 
	186, 187, 7, 64, 2, 2, 187, 22, 3, 2, 2, 2, 188, 189, 7, 62, 2, 2, 189, 
	24, 3, 2, 2, 2, 190, 191, 7, 64, 2, 2, 191, 192, 7, 63, 2, 2, 192, 26, 
	3, 2, 2, 2, 193, 194, 7, 62, 2, 2, 194, 195, 7, 63, 2, 2, 195, 28, 3, 2, 
	2, 2, 196, 197, 7, 62, 2, 2, 197, 198, 7, 64, 2, 2, 198, 30, 3, 2, 2, 2, 
	199, 200, 7, 128, 2, 2, 200, 32, 3, 2, 2, 2, 201, 202, 7, 35, 2, 2, 202, 
	203, 7, 128, 2, 2, 203, 34, 3, 2, 2, 2, 204, 205, 5, 97, 49, 2, 205, 206, 
	5, 133, 67, 2, 206, 36, 3, 2, 2, 2, 207, 208, 5, 133, 67, 2, 208, 209, 
	5, 105, 53, 2, 209, 210, 5, 119, 60, 2, 210, 211, 5, 105, 53, 2, 211, 212, 
	5, 101, 51, 2, 212, 213, 5, 135, 68, 2, 213, 38, 3, 2, 2, 2, 214, 215, 
	5, 107, 54, 2, 215, 216, 5, 131, 66, 2, 216, 217, 5, 125, 63, 2, 217, 218, 
	5, 121, 61, 2, 218, 40, 3, 2, 2, 2, 219, 220, 5, 121, 61, 2, 220, 221, 
	5, 97, 49, 2, 221, 222, 5, 143, 72, 2, 222, 42, 3, 2, 2, 2, 223, 224, 5, 
	133, 67, 2, 224, 225, 5, 137, 69, 2, 225, 226, 5, 121, 61, 2, 226, 44, 
	3, 2, 2, 2, 227, 228, 5, 97, 49, 2, 228, 229, 5, 139, 70, 2, 229, 230, 
	5, 109, 55, 2, 230, 46, 3, 2, 2, 2, 231, 232, 5, 121, 61, 2, 232, 233, 
	5, 113, 57, 2, 233, 234, 5, 123, 62, 2, 234, 48, 3, 2, 2, 2, 235, 236, 
	5, 101, 51, 2, 236, 237, 5, 125, 63, 2, 237, 238, 5, 137, 69, 2, 238, 239, 
	5, 123, 62, 2, 239, 240, 5, 135, 68, 2, 240, 50, 3, 2, 2, 2, 241, 242, 
	5, 103, 52, 2, 242, 243, 5, 113, 57, 2, 243, 244, 5, 133, 67, 2, 244, 245, 
	5, 135, 68, 2, 245, 246, 5, 113, 57, 2, 246, 247, 5, 123, 62, 2, 247, 248, 
	5, 101, 51, 2, 248, 249, 5, 135, 68, 2, 249, 52, 3, 2, 2, 2, 250, 251, 
	5, 141, 71, 2, 251, 252, 5, 111, 56, 2, 252, 253, 5, 105, 53, 2, 253, 254, 
	5, 131, 66, 2, 254, 255, 5, 105, 53, 2, 255, 54, 3, 2, 2, 2, 256, 257, 
	5, 109, 55, 2, 257, 258, 5, 131, 66, 2, 258, 259, 5, 125, 63, 2, 259, 260, 
	5, 137, 69, 2, 260, 261, 5, 127, 64, 2, 261, 56, 3, 2, 2, 2, 262, 263, 
	5, 99, 50, 2, 263, 264, 5, 145, 73, 2, 264, 58, 3, 2, 2, 2, 265, 266, 5, 
	125, 63, 2, 266, 267, 5, 131, 66, 2, 267, 268, 5, 103, 52, 2, 268, 269, 
	5, 105, 53, 2, 269, 270, 5, 131, 66, 2, 270, 60, 3, 2, 2, 2, 271, 272, 
	5, 123, 62, 2, 272, 273, 5, 137, 69, 2, 273, 274, 5, 119, 60, 2, 274, 275, 
	5, 119, 60, 2, 275, 276, 5, 133, 67, 2, 276, 62, 3, 2, 2, 2, 277, 278, 
	5, 107, 54, 2, 278, 279, 5, 113, 57, 2, 279, 280, 5, 131, 66, 2, 280, 281, 
	5, 133, 67, 2, 281, 282, 5, 135, 68, 2, 282, 64, 3, 2, 2, 2, 283, 284, 
	5, 119, 60, 2, 284, 285, 5, 97, 49, 2, 285, 286, 5, 133, 67, 2, 286, 287, 
	5, 135, 68, 2, 287, 66, 3, 2, 2, 2, 288, 289, 5, 111, 56, 2, 289, 290, 
	5, 97, 49, 2, 290, 291, 5, 139, 70, 2, 291, 292, 5, 113, 57, 2, 292, 293, 
	5, 123, 62, 2, 293, 294, 5, 109, 55, 2, 294, 68, 3, 2, 2, 2, 295, 296, 
	5, 123, 62, 2, 296, 297, 5, 125, 63, 2, 297, 298, 5, 135, 68, 2, 298, 70, 
	3, 2, 2, 2, 299, 300, 5, 113, 57, 2, 300, 301, 5, 133, 67, 2, 301, 72, 
	3, 2, 2, 2, 302, 303, 5, 99, 50, 2, 303, 304, 5, 105, 53, 2, 304, 305, 
	5, 135, 68, 2, 305, 306, 5, 141, 71, 2, 306, 307, 5, 105, 53, 2, 307, 308, 
	5, 105, 53, 2, 308, 309, 5, 123, 62, 2, 309, 74, 3, 2, 2, 2, 310, 311, 
	5, 97, 49, 2, 311, 312, 5, 123, 62, 2, 312, 313, 5, 103, 52, 2, 313, 76, 
	3, 2, 2, 2, 314, 315, 5, 113, 57, 2, 315, 316, 5, 123, 62, 2, 316, 78, 
	3, 2, 2, 2, 317, 318, 5, 123, 62, 2, 318, 319, 5, 137, 69, 2, 319, 320, 
	5, 119, 60, 2, 320, 321, 5, 119, 60, 2, 321, 80, 3, 2, 2, 2, 322, 323, 
	5, 125, 63, 2, 323, 324, 5, 131, 66, 2, 324, 82, 3, 2, 2, 2, 325, 326, 
	5, 97, 49, 2, 326, 327, 5, 133, 67, 2, 327, 328, 5, 101, 51, 2, 328, 84, 
	3, 2, 2, 2, 329, 330, 5, 103, 52, 2, 330, 331, 5, 105, 53, 2, 331, 332, 
	5, 133, 67, 2, 332, 333, 5, 101, 51, 2, 333, 86, 3, 2, 2, 2, 334, 335, 
	5, 119, 60, 2, 335, 336, 5, 113, 57, 2, 336, 337, 5, 121, 61, 2, 337, 338, 
	5, 113, 57, 2, 338, 339, 5, 135, 68, 2, 339, 88, 3, 2, 2, 2, 340, 341, 
	5, 125, 63, 2, 341, 342, 5, 107, 54, 2, 342, 343, 5, 107, 54, 2, 343, 344, 
	5, 133, 67, 2, 344, 345, 5, 105, 53, 2, 345, 346, 5, 135, 68, 2, 346, 90, 
	3, 2, 2, 2, 347, 348, 5, 119, 60, 2, 348, 349, 5, 113, 57, 2, 349, 350, 
	5, 117, 59, 2, 350, 351, 5, 105, 53, 2, 351, 92, 3, 2, 2, 2, 352, 353, 
	5, 105, 53, 2, 353, 354, 5, 143, 72, 2, 354, 355, 5, 113, 57, 2, 355, 356, 
	5, 133, 67, 2, 356, 357, 5, 135, 68, 2, 357, 358, 5, 133, 67, 2, 358, 94, 
	3, 2, 2, 2, 359, 360, 5, 101, 51, 2, 360, 361, 5, 97, 49, 2, 361, 362, 
	5, 133, 67, 2, 362, 363, 5, 135, 68, 2, 363, 96, 3, 2, 2, 2, 364, 365, 
	9, 2, 2, 2, 365, 98, 3, 2, 2, 2, 366, 367, 9, 3, 2, 2, 367, 100, 3, 2, 
	2, 2, 368, 369, 9, 4, 2, 2, 369, 102, 3, 2, 2, 2, 370, 371, 9, 5, 2, 2, 
	371, 104, 3, 2, 2, 2, 372, 373, 9, 6, 2, 2, 373, 106, 3, 2, 2, 2, 374, 
	375, 9, 7, 2, 2, 375, 108, 3, 2, 2, 2, 376, 377, 9, 8, 2, 2, 377, 110, 
	3, 2, 2, 2, 378, 379, 9, 9, 2, 2, 379, 112, 3, 2, 2, 2, 380, 381, 9, 10, 
	2, 2, 381, 114, 3, 2, 2, 2, 382, 383, 9, 11, 2, 2, 383, 116, 3, 2, 2, 2, 
	384, 385, 9, 12, 2, 2, 385, 118, 3, 2, 2, 2, 386, 387, 9, 13, 2, 2, 387, 
	120, 3, 2, 2, 2, 388, 389, 9, 14, 2, 2, 389, 122, 3, 2, 2, 2, 390, 391, 
	9, 15, 2, 2, 391, 124, 3, 2, 2, 2, 392, 393, 9, 16, 2, 2, 393, 126, 3, 
	2, 2, 2, 394, 395, 9, 17, 2, 2, 395, 128, 3, 2, 2, 2, 396, 397, 9, 18, 
	2, 2, 397, 130, 3, 2, 2, 2, 398, 399, 9, 19, 2, 2, 399, 132, 3, 2, 2, 2, 
	400, 401, 9, 20, 2, 2, 401, 134, 3, 2, 2, 2, 402, 403, 9, 21, 2, 2, 403, 
	136, 3, 2, 2, 2, 404, 405, 9, 22, 2, 2, 405, 138, 3, 2, 2, 2, 406, 407, 
	9, 23, 2, 2, 407, 140, 3, 2, 2, 2, 408, 409, 9, 24, 2, 2, 409, 142, 3, 
	2, 2, 2, 410, 411, 9, 25, 2, 2, 411, 144, 3, 2, 2, 2, 412, 413, 9, 26, 
	2, 2, 413, 146, 3, 2, 2, 2, 414, 415, 9, 27, 2, 2, 415, 148, 3, 2, 2, 2, 
	416, 417, 9, 28, 2, 2, 417, 150, 3, 2, 2, 2, 418, 419, 9, 29, 2, 2, 419, 
	152, 3, 2, 2, 2, 420, 421, 9, 30, 2, 2, 421, 154, 3, 2, 2, 2, 422, 424, 
	5, 149, 75, 2, 423, 422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 423, 3, 
	2, 2, 2, 425, 426, 3, 2, 2, 2, 426, 156, 3, 2, 2, 2, 427, 431, 9, 31, 2, 
	2, 428, 430, 9, 32, 2, 2, 429, 428, 3, 2, 2, 2, 430, 433, 3, 2, 2, 2, 431, 
	429, 3, 2, 2, 2, 431, 432, 3, 2, 2, 2, 432, 158, 3, 2, 2, 2, 433, 431, 
	3, 2, 2, 2, 434, 444, 7, 41, 2, 2, 435, 436, 7, 94, 2, 2, 436, 443, 7, 
	94, 2, 2, 437, 438, 7, 41, 2, 2, 438, 443, 7, 41, 2, 2, 439, 440, 7, 94, 
	2, 2, 440, 443, 7, 41, 2, 2, 441, 443, 10, 33, 2, 2, 442, 435, 3, 2, 2, 
	2, 442, 437, 3, 2, 2, 2, 442, 439, 3, 2, 2, 2, 442, 441, 3, 2, 2, 2, 443, 
	446, 3, 2, 2, 2, 444, 442, 3, 2, 2, 2, 444, 445, 3, 2, 2, 2, 445, 447, 
	3, 2, 2, 2, 446, 444, 3, 2, 2, 2, 447, 448, 7, 41, 2, 2, 448, 160, 3, 2, 
	2, 2, 449, 451, 7, 36, 2, 2, 450, 452, 10, 34, 2, 2, 451, 450, 3, 2, 2, 
	2, 452, 453, 3, 2, 2, 2, 453, 451, 3, 2, 2, 2, 453, 454, 3, 2, 2, 2, 454, 
	455, 3, 2, 2, 2, 455, 456, 7, 36, 2, 2, 456, 162, 3, 2, 2, 2, 457, 459, 
	7, 60, 2, 2, 458, 460, 10, 35, 2, 2, 459, 458, 3, 2, 2, 2, 460, 461, 3, 
	2, 2, 2, 461, 459, 3, 2, 2, 2, 461, 462, 3, 2, 2, 2, 462, 164, 3, 2, 2, 
	2, 463, 465, 9, 36, 2, 2, 464, 463, 3, 2, 2, 2, 465, 466, 3, 2, 2, 2, 466, 
	464, 3, 2, 2, 2, 466, 467, 3, 2, 2, 2, 467, 468, 3, 2, 2, 2, 468, 469, 
	8, 83, 2, 2, 469, 166, 3, 2, 2, 2, 10, 2, 425, 431, 442, 444, 453, 461, 
	466, 3, 8, 2, 2,
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
	"", "'*'", "','", "'/'", "'+'", "'-'", "'||'", "'('", "')'", "'='", "'>'", 
	"'<'", "'>='", "'<='", "'<>'", "'~'", "'!~'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "AS", 
	"SELECT", "FROM", "MAX", "SUM", "AVG", "MIN", "COUNT", "DISTINCT", "WHERE", 
	"GROUP", "BY", "ORDER", "NULLS", "FIRST", "LAST", "HAVING", "NOT", "IS", 
	"BETWEEN", "AND", "IN", "NULL", "OR", "ASC", "DESC", "LIMIT", "OFFSET", 
	"LIKE", "EXISTS", "CAST", "DECIMAL_LITERAL", "ID", "TEXT_STRING", "TEXT_ALIAS", 
	"BIND_VARIABLE", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "AS", "SELECT", 
	"FROM", "MAX", "SUM", "AVG", "MIN", "COUNT", "DISTINCT", "WHERE", "GROUP", 
	"BY", "ORDER", "NULLS", "FIRST", "LAST", "HAVING", "NOT", "IS", "BETWEEN", 
	"AND", "IN", "NULL", "OR", "ASC", "DESC", "LIMIT", "OFFSET", "LIKE", "EXISTS", 
	"CAST", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", 
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "DEC_DIGIT", 
	"HEX_DIGIT", "LETTER", "DECIMAL_LITERAL", "ID", "TEXT_STRING", "TEXT_ALIAS", 
	"BIND_VARIABLE", "WS",
}

type SqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSqlLexer(input antlr.CharStream) *SqlLexer {

	l := new(SqlLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Sql.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SqlLexer tokens.
const (
	SqlLexerT__0 = 1
	SqlLexerT__1 = 2
	SqlLexerT__2 = 3
	SqlLexerT__3 = 4
	SqlLexerT__4 = 5
	SqlLexerT__5 = 6
	SqlLexerT__6 = 7
	SqlLexerT__7 = 8
	SqlLexerT__8 = 9
	SqlLexerT__9 = 10
	SqlLexerT__10 = 11
	SqlLexerT__11 = 12
	SqlLexerT__12 = 13
	SqlLexerT__13 = 14
	SqlLexerT__14 = 15
	SqlLexerT__15 = 16
	SqlLexerAS = 17
	SqlLexerSELECT = 18
	SqlLexerFROM = 19
	SqlLexerMAX = 20
	SqlLexerSUM = 21
	SqlLexerAVG = 22
	SqlLexerMIN = 23
	SqlLexerCOUNT = 24
	SqlLexerDISTINCT = 25
	SqlLexerWHERE = 26
	SqlLexerGROUP = 27
	SqlLexerBY = 28
	SqlLexerORDER = 29
	SqlLexerNULLS = 30
	SqlLexerFIRST = 31
	SqlLexerLAST = 32
	SqlLexerHAVING = 33
	SqlLexerNOT = 34
	SqlLexerIS = 35
	SqlLexerBETWEEN = 36
	SqlLexerAND = 37
	SqlLexerIN = 38
	SqlLexerNULL = 39
	SqlLexerOR = 40
	SqlLexerASC = 41
	SqlLexerDESC = 42
	SqlLexerLIMIT = 43
	SqlLexerOFFSET = 44
	SqlLexerLIKE = 45
	SqlLexerEXISTS = 46
	SqlLexerCAST = 47
	SqlLexerDECIMAL_LITERAL = 48
	SqlLexerID = 49
	SqlLexerTEXT_STRING = 50
	SqlLexerTEXT_ALIAS = 51
	SqlLexerBIND_VARIABLE = 52
	SqlLexerWS = 53
)

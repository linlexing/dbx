// Code generated from Sql.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Sql

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 365, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 
	4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 78, 10, 
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 96, 10, 12, 3, 12, 5, 
	12, 99, 10, 12, 3, 12, 5, 12, 102, 10, 12, 3, 12, 5, 12, 105, 10, 12, 3, 
	12, 5, 12, 108, 10, 12, 3, 12, 5, 12, 111, 10, 12, 3, 12, 3, 12, 3, 12, 
	3, 12, 7, 12, 117, 10, 12, 12, 12, 14, 12, 120, 11, 12, 3, 13, 3, 13, 5, 
	13, 124, 10, 13, 3, 13, 3, 13, 7, 13, 128, 10, 13, 12, 13, 14, 13, 131, 
	11, 13, 3, 14, 3, 14, 5, 14, 135, 10, 14, 3, 14, 5, 14, 138, 10, 14, 3, 
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 148, 10, 15, 
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 159, 
	10, 15, 12, 15, 14, 15, 162, 11, 15, 3, 16, 3, 16, 3, 16, 5, 16, 167, 10, 
	16, 3, 17, 3, 17, 5, 17, 171, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 
	3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 182, 10, 18, 3, 18, 5, 18, 185, 10, 
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 
	3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 201, 10, 19, 3, 20, 3, 20, 3, 20, 7, 
	20, 206, 10, 20, 12, 20, 14, 20, 209, 11, 20, 3, 21, 3, 21, 5, 21, 213, 
	10, 21, 3, 21, 3, 21, 3, 21, 5, 21, 218, 10, 21, 7, 21, 220, 10, 21, 12, 
	21, 14, 21, 223, 11, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 230, 
	10, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 
	3, 24, 5, 24, 242, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 
	24, 5, 24, 251, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 258, 
	10, 24, 12, 24, 14, 24, 261, 11, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 
	267, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 276, 
	10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 284, 10, 24, 3, 
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 294, 10, 24, 
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 302, 10, 24, 3, 24, 3, 
	24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 310, 10, 24, 12, 24, 14, 24, 313, 
	11, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 7, 26, 322, 10, 
	26, 12, 26, 14, 26, 325, 11, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 
	29, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 337, 10, 29, 12, 29, 14, 29, 340, 
	11, 29, 3, 30, 3, 30, 5, 30, 344, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 
	30, 350, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 356, 10, 31, 3, 31, 
	3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 363, 10, 31, 3, 31, 2, 5, 22, 28, 46, 
	32, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 2, 9, 4, 2, 58, 58, 60, 
	60, 3, 2, 50, 52, 4, 2, 3, 3, 5, 5, 3, 2, 6, 7, 3, 2, 22, 25, 3, 2, 11, 
	18, 3, 2, 43, 44, 2, 388, 2, 62, 3, 2, 2, 2, 4, 64, 3, 2, 2, 2, 6, 66, 
	3, 2, 2, 2, 8, 68, 3, 2, 2, 2, 10, 70, 3, 2, 2, 2, 12, 72, 3, 2, 2, 2, 
	14, 75, 3, 2, 2, 2, 16, 79, 3, 2, 2, 2, 18, 81, 3, 2, 2, 2, 20, 83, 3, 
	2, 2, 2, 22, 85, 3, 2, 2, 2, 24, 123, 3, 2, 2, 2, 26, 132, 3, 2, 2, 2, 
	28, 147, 3, 2, 2, 2, 30, 166, 3, 2, 2, 2, 32, 170, 3, 2, 2, 2, 34, 184, 
	3, 2, 2, 2, 36, 200, 3, 2, 2, 2, 38, 202, 3, 2, 2, 2, 40, 210, 3, 2, 2, 
	2, 42, 229, 3, 2, 2, 2, 44, 231, 3, 2, 2, 2, 46, 301, 3, 2, 2, 2, 48, 314, 
	3, 2, 2, 2, 50, 316, 3, 2, 2, 2, 52, 326, 3, 2, 2, 2, 54, 328, 3, 2, 2, 
	2, 56, 331, 3, 2, 2, 2, 58, 341, 3, 2, 2, 2, 60, 351, 3, 2, 2, 2, 62, 63, 
	7, 58, 2, 2, 63, 3, 3, 2, 2, 2, 64, 65, 7, 58, 2, 2, 65, 5, 3, 2, 2, 2, 
	66, 67, 7, 58, 2, 2, 67, 7, 3, 2, 2, 2, 68, 69, 7, 58, 2, 2, 69, 9, 3, 
	2, 2, 2, 70, 71, 9, 2, 2, 2, 71, 11, 3, 2, 2, 2, 72, 73, 9, 3, 2, 2, 73, 
	74, 7, 53, 2, 2, 74, 13, 3, 2, 2, 2, 75, 77, 7, 55, 2, 2, 76, 78, 7, 56, 
	2, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 15, 3, 2, 2, 2, 79, 80, 
	7, 57, 2, 2, 80, 17, 3, 2, 2, 2, 81, 82, 7, 59, 2, 2, 82, 19, 3, 2, 2, 
	2, 83, 84, 7, 61, 2, 2, 84, 21, 3, 2, 2, 2, 85, 86, 8, 12, 1, 2, 86, 87, 
	7, 20, 2, 2, 87, 88, 5, 24, 13, 2, 88, 89, 7, 21, 2, 2, 89, 95, 5, 40, 
	21, 2, 90, 91, 5, 12, 7, 2, 91, 92, 5, 40, 21, 2, 92, 93, 7, 54, 2, 2, 
	93, 94, 5, 46, 24, 2, 94, 96, 3, 2, 2, 2, 95, 90, 3, 2, 2, 2, 95, 96, 3, 
	2, 2, 2, 96, 98, 3, 2, 2, 2, 97, 99, 5, 44, 23, 2, 98, 97, 3, 2, 2, 2, 
	98, 99, 3, 2, 2, 2, 99, 101, 3, 2, 2, 2, 100, 102, 5, 50, 26, 2, 101, 100, 
	3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 104, 3, 2, 2, 2, 103, 105, 5, 54, 
	28, 2, 104, 103, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 
	106, 108, 5, 56, 29, 2, 107, 106, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 
	110, 3, 2, 2, 2, 109, 111, 5, 60, 31, 2, 110, 109, 3, 2, 2, 2, 110, 111, 
	3, 2, 2, 2, 111, 118, 3, 2, 2, 2, 112, 113, 12, 3, 2, 2, 113, 114, 5, 14, 
	8, 2, 114, 115, 5, 22, 12, 4, 115, 117, 3, 2, 2, 2, 116, 112, 3, 2, 2, 
	2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 
	23, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 124, 7, 3, 2, 2, 122, 124, 5, 
	26, 14, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3, 2, 2, 2, 124, 129, 3, 2, 
	2, 2, 125, 126, 7, 4, 2, 2, 126, 128, 5, 26, 14, 2, 127, 125, 3, 2, 2, 
	2, 128, 131, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 
	25, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 132, 137, 5, 28, 15, 2, 133, 135, 
	7, 19, 2, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 3, 2, 
	2, 2, 136, 138, 5, 10, 6, 2, 137, 134, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 
	138, 27, 3, 2, 2, 2, 139, 140, 8, 15, 1, 2, 140, 148, 5, 2, 2, 2, 141, 
	148, 5, 32, 17, 2, 142, 148, 5, 30, 16, 2, 143, 144, 7, 9, 2, 2, 144, 145, 
	5, 28, 15, 2, 145, 146, 7, 10, 2, 2, 146, 148, 3, 2, 2, 2, 147, 139, 3, 
	2, 2, 2, 147, 141, 3, 2, 2, 2, 147, 142, 3, 2, 2, 2, 147, 143, 3, 2, 2, 
	2, 148, 160, 3, 2, 2, 2, 149, 150, 12, 6, 2, 2, 150, 151, 9, 4, 2, 2, 151, 
	159, 5, 28, 15, 7, 152, 153, 12, 5, 2, 2, 153, 154, 9, 5, 2, 2, 154, 159, 
	5, 28, 15, 6, 155, 156, 12, 4, 2, 2, 156, 157, 7, 8, 2, 2, 157, 159, 5, 
	28, 15, 5, 158, 149, 3, 2, 2, 2, 158, 152, 3, 2, 2, 2, 158, 155, 3, 2, 
	2, 2, 159, 162, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 
	161, 29, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 163, 167, 5, 16, 9, 2, 164, 
	167, 5, 18, 10, 2, 165, 167, 5, 20, 11, 2, 166, 163, 3, 2, 2, 2, 166, 164, 
	3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 167, 31, 3, 2, 2, 2, 168, 171, 5, 34, 
	18, 2, 169, 171, 5, 36, 19, 2, 170, 168, 3, 2, 2, 2, 170, 169, 3, 2, 2, 
	2, 171, 33, 3, 2, 2, 2, 172, 173, 9, 6, 2, 2, 173, 174, 7, 9, 2, 2, 174, 
	175, 5, 38, 20, 2, 175, 176, 7, 10, 2, 2, 176, 185, 3, 2, 2, 2, 177, 178, 
	7, 26, 2, 2, 178, 181, 7, 9, 2, 2, 179, 182, 7, 3, 2, 2, 180, 182, 5, 38, 
	20, 2, 181, 179, 3, 2, 2, 2, 181, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 
	183, 185, 7, 10, 2, 2, 184, 172, 3, 2, 2, 2, 184, 177, 3, 2, 2, 2, 185, 
	35, 3, 2, 2, 2, 186, 187, 5, 8, 5, 2, 187, 188, 7, 9, 2, 2, 188, 189, 5, 
	38, 20, 2, 189, 190, 7, 10, 2, 2, 190, 201, 3, 2, 2, 2, 191, 192, 7, 27, 
	2, 2, 192, 201, 5, 38, 20, 2, 193, 194, 7, 49, 2, 2, 194, 195, 7, 9, 2, 
	2, 195, 196, 5, 38, 20, 2, 196, 197, 7, 19, 2, 2, 197, 198, 5, 6, 4, 2, 
	198, 199, 7, 10, 2, 2, 199, 201, 3, 2, 2, 2, 200, 186, 3, 2, 2, 2, 200, 
	191, 3, 2, 2, 2, 200, 193, 3, 2, 2, 2, 201, 37, 3, 2, 2, 2, 202, 207, 5, 
	28, 15, 2, 203, 204, 7, 4, 2, 2, 204, 206, 5, 28, 15, 2, 205, 203, 3, 2, 
	2, 2, 206, 209, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 
	208, 39, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 210, 212, 5, 42, 22, 2, 211, 
	213, 5, 10, 6, 2, 212, 211, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 221, 
	3, 2, 2, 2, 214, 215, 7, 4, 2, 2, 215, 217, 5, 42, 22, 2, 216, 218, 5, 
	10, 6, 2, 217, 216, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 220, 3, 2, 2, 
	2, 219, 214, 3, 2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221, 
	222, 3, 2, 2, 2, 222, 41, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 224, 230, 5, 
	4, 3, 2, 225, 226, 7, 9, 2, 2, 226, 227, 5, 22, 12, 2, 227, 228, 7, 10, 
	2, 2, 228, 230, 3, 2, 2, 2, 229, 224, 3, 2, 2, 2, 229, 225, 3, 2, 2, 2, 
	230, 43, 3, 2, 2, 2, 231, 232, 7, 28, 2, 2, 232, 233, 5, 46, 24, 2, 233, 
	45, 3, 2, 2, 2, 234, 235, 8, 24, 1, 2, 235, 236, 5, 28, 15, 2, 236, 237, 
	5, 48, 25, 2, 237, 238, 5, 28, 15, 2, 238, 302, 3, 2, 2, 2, 239, 241, 5, 
	28, 15, 2, 240, 242, 7, 36, 2, 2, 241, 240, 3, 2, 2, 2, 241, 242, 3, 2, 
	2, 2, 242, 243, 3, 2, 2, 2, 243, 244, 7, 38, 2, 2, 244, 245, 5, 28, 15, 
	2, 245, 246, 7, 39, 2, 2, 246, 247, 5, 28, 15, 2, 247, 302, 3, 2, 2, 2, 
	248, 250, 5, 28, 15, 2, 249, 251, 7, 36, 2, 2, 250, 249, 3, 2, 2, 2, 250, 
	251, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 7, 40, 2, 2, 253, 254, 
	7, 9, 2, 2, 254, 259, 5, 28, 15, 2, 255, 256, 7, 4, 2, 2, 256, 258, 5, 
	28, 15, 2, 257, 255, 3, 2, 2, 2, 258, 261, 3, 2, 2, 2, 259, 257, 3, 2, 
	2, 2, 259, 260, 3, 2, 2, 2, 260, 262, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 
	262, 263, 7, 10, 2, 2, 263, 302, 3, 2, 2, 2, 264, 266, 5, 28, 15, 2, 265, 
	267, 7, 36, 2, 2, 266, 265, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 268, 
	3, 2, 2, 2, 268, 269, 7, 40, 2, 2, 269, 270, 7, 9, 2, 2, 270, 271, 5, 22, 
	12, 2, 271, 272, 7, 10, 2, 2, 272, 302, 3, 2, 2, 2, 273, 275, 5, 28, 15, 
	2, 274, 276, 7, 36, 2, 2, 275, 274, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 
	277, 3, 2, 2, 2, 277, 278, 7, 47, 2, 2, 278, 279, 5, 28, 15, 2, 279, 302, 
	3, 2, 2, 2, 280, 281, 5, 28, 15, 2, 281, 283, 7, 37, 2, 2, 282, 284, 7, 
	36, 2, 2, 283, 282, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 285, 3, 2, 2, 
	2, 285, 286, 7, 41, 2, 2, 286, 302, 3, 2, 2, 2, 287, 288, 7, 48, 2, 2, 
	288, 289, 7, 9, 2, 2, 289, 290, 5, 22, 12, 2, 290, 291, 7, 10, 2, 2, 291, 
	302, 3, 2, 2, 2, 292, 294, 7, 62, 2, 2, 293, 292, 3, 2, 2, 2, 293, 294, 
	3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 296, 7, 9, 2, 2, 296, 297, 5, 46, 
	24, 2, 297, 298, 7, 10, 2, 2, 298, 302, 3, 2, 2, 2, 299, 300, 7, 36, 2, 
	2, 300, 302, 5, 46, 24, 5, 301, 234, 3, 2, 2, 2, 301, 239, 3, 2, 2, 2, 
	301, 248, 3, 2, 2, 2, 301, 264, 3, 2, 2, 2, 301, 273, 3, 2, 2, 2, 301, 
	280, 3, 2, 2, 2, 301, 287, 3, 2, 2, 2, 301, 293, 3, 2, 2, 2, 301, 299, 
	3, 2, 2, 2, 302, 311, 3, 2, 2, 2, 303, 304, 12, 4, 2, 2, 304, 305, 7, 39, 
	2, 2, 305, 310, 5, 46, 24, 5, 306, 307, 12, 3, 2, 2, 307, 308, 7, 42, 2, 
	2, 308, 310, 5, 46, 24, 4, 309, 303, 3, 2, 2, 2, 309, 306, 3, 2, 2, 2, 
	310, 313, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2, 312, 
	47, 3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 314, 315, 9, 7, 2, 2, 315, 49, 3, 
	2, 2, 2, 316, 317, 7, 29, 2, 2, 317, 318, 7, 30, 2, 2, 318, 323, 5, 52, 
	27, 2, 319, 320, 7, 4, 2, 2, 320, 322, 5, 52, 27, 2, 321, 319, 3, 2, 2, 
	2, 322, 325, 3, 2, 2, 2, 323, 321, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 
	51, 3, 2, 2, 2, 325, 323, 3, 2, 2, 2, 326, 327, 5, 28, 15, 2, 327, 53, 
	3, 2, 2, 2, 328, 329, 7, 35, 2, 2, 329, 330, 5, 46, 24, 2, 330, 55, 3, 
	2, 2, 2, 331, 332, 7, 31, 2, 2, 332, 333, 7, 30, 2, 2, 333, 338, 5, 58, 
	30, 2, 334, 335, 7, 4, 2, 2, 335, 337, 5, 58, 30, 2, 336, 334, 3, 2, 2, 
	2, 337, 340, 3, 2, 2, 2, 338, 336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 
	57, 3, 2, 2, 2, 340, 338, 3, 2, 2, 2, 341, 343, 5, 28, 15, 2, 342, 344, 
	9, 8, 2, 2, 343, 342, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 349, 3, 2, 
	2, 2, 345, 346, 7, 32, 2, 2, 346, 350, 7, 33, 2, 2, 347, 348, 7, 32, 2, 
	2, 348, 350, 7, 34, 2, 2, 349, 345, 3, 2, 2, 2, 349, 347, 3, 2, 2, 2, 349, 
	350, 3, 2, 2, 2, 350, 59, 3, 2, 2, 2, 351, 362, 7, 45, 2, 2, 352, 353, 
	5, 16, 9, 2, 353, 354, 7, 4, 2, 2, 354, 356, 3, 2, 2, 2, 355, 352, 3, 2, 
	2, 2, 355, 356, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 363, 5, 16, 9, 2, 
	358, 359, 5, 16, 9, 2, 359, 360, 7, 46, 2, 2, 360, 361, 5, 16, 9, 2, 361, 
	363, 3, 2, 2, 2, 362, 355, 3, 2, 2, 2, 362, 358, 3, 2, 2, 2, 363, 61, 3, 
	2, 2, 2, 43, 77, 95, 98, 101, 104, 107, 110, 118, 123, 129, 134, 137, 147, 
	158, 160, 166, 170, 181, 184, 200, 207, 212, 217, 221, 229, 241, 250, 259, 
	266, 275, 283, 293, 301, 309, 311, 323, 338, 343, 349, 355, 362,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'*'", "','", "'/'", "'+'", "'-'", "'||'", "'('", "')'", "'='", "'>'", 
	"'<'", "'>='", "'<='", "'<>'", "'~'", "'!~'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "AS", 
	"SELECT", "FROM", "MAX", "SUM", "AVG", "MIN", "COUNT", "DISTINCT", "WHERE", 
	"GROUP", "BY", "ORDER", "NULLS", "FIRST", "LAST", "HAVING", "NOT", "IS", 
	"BETWEEN", "AND", "IN", "NULL", "OR", "ASC", "DESC", "LIMIT", "OFFSET", 
	"LIKE", "EXISTS", "CAST", "INNER", "LEFT", "RIGHT", "JOIN", "ON", "UNION", 
	"ALL", "DECIMAL_LITERAL", "ID", "TEXT_STRING", "TEXT_ALIAS", "BIND_VARIABLE", 
	"COMMENT", "WS",
}

var ruleNames = []string{
	"columnName", "tableName", "typeName", "functionName", "alias", "join", 
	"union", "decimalLiteral", "textLiteral", "bind_variables", "selectStatement", 
	"selectElements", "selectElement", "expr", "value", "functionCall", "aggregateFunction", 
	"commonFunction", "functionArg", "tableSources", "tableSource", "whereClause", 
	"logicExpression", "comparisonOperator", "groupByClause", "groupByItem", 
	"havingClause", "orderByClause", "orderByExpression", "limitClause",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SqlParser struct {
	*antlr.BaseParser
}

func NewSqlParser(input antlr.TokenStream) *SqlParser {
	this := new(SqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sql.g4"

	return this
}

// SqlParser tokens.
const (
	SqlParserEOF = antlr.TokenEOF
	SqlParserT__0 = 1
	SqlParserT__1 = 2
	SqlParserT__2 = 3
	SqlParserT__3 = 4
	SqlParserT__4 = 5
	SqlParserT__5 = 6
	SqlParserT__6 = 7
	SqlParserT__7 = 8
	SqlParserT__8 = 9
	SqlParserT__9 = 10
	SqlParserT__10 = 11
	SqlParserT__11 = 12
	SqlParserT__12 = 13
	SqlParserT__13 = 14
	SqlParserT__14 = 15
	SqlParserT__15 = 16
	SqlParserAS = 17
	SqlParserSELECT = 18
	SqlParserFROM = 19
	SqlParserMAX = 20
	SqlParserSUM = 21
	SqlParserAVG = 22
	SqlParserMIN = 23
	SqlParserCOUNT = 24
	SqlParserDISTINCT = 25
	SqlParserWHERE = 26
	SqlParserGROUP = 27
	SqlParserBY = 28
	SqlParserORDER = 29
	SqlParserNULLS = 30
	SqlParserFIRST = 31
	SqlParserLAST = 32
	SqlParserHAVING = 33
	SqlParserNOT = 34
	SqlParserIS = 35
	SqlParserBETWEEN = 36
	SqlParserAND = 37
	SqlParserIN = 38
	SqlParserNULL = 39
	SqlParserOR = 40
	SqlParserASC = 41
	SqlParserDESC = 42
	SqlParserLIMIT = 43
	SqlParserOFFSET = 44
	SqlParserLIKE = 45
	SqlParserEXISTS = 46
	SqlParserCAST = 47
	SqlParserINNER = 48
	SqlParserLEFT = 49
	SqlParserRIGHT = 50
	SqlParserJOIN = 51
	SqlParserON = 52
	SqlParserUNION = 53
	SqlParserALL = 54
	SqlParserDECIMAL_LITERAL = 55
	SqlParserID = 56
	SqlParserTEXT_STRING = 57
	SqlParserTEXT_ALIAS = 58
	SqlParserBIND_VARIABLE = 59
	SqlParserCOMMENT = 60
	SqlParserWS = 61
)

// SqlParser rules.
const (
	SqlParserRULE_columnName = 0
	SqlParserRULE_tableName = 1
	SqlParserRULE_typeName = 2
	SqlParserRULE_functionName = 3
	SqlParserRULE_alias = 4
	SqlParserRULE_join = 5
	SqlParserRULE_union = 6
	SqlParserRULE_decimalLiteral = 7
	SqlParserRULE_textLiteral = 8
	SqlParserRULE_bind_variables = 9
	SqlParserRULE_selectStatement = 10
	SqlParserRULE_selectElements = 11
	SqlParserRULE_selectElement = 12
	SqlParserRULE_expr = 13
	SqlParserRULE_value = 14
	SqlParserRULE_functionCall = 15
	SqlParserRULE_aggregateFunction = 16
	SqlParserRULE_commonFunction = 17
	SqlParserRULE_functionArg = 18
	SqlParserRULE_tableSources = 19
	SqlParserRULE_tableSource = 20
	SqlParserRULE_whereClause = 21
	SqlParserRULE_logicExpression = 22
	SqlParserRULE_comparisonOperator = 23
	SqlParserRULE_groupByClause = 24
	SqlParserRULE_groupByItem = 25
	SqlParserRULE_havingClause = 26
	SqlParserRULE_orderByClause = 27
	SqlParserRULE_orderByExpression = 28
	SqlParserRULE_limitClause = 29
)

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_columnName
	return p
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (s *ColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) ColumnName() (localctx IColumnNameContext) {
	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SqlParserRULE_columnName)

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
		p.SetState(60)
		p.Match(SqlParserID)
	}



	return localctx
}


// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (s *TableNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitTableName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SqlParserRULE_tableName)

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
		p.SetState(62)
		p.Match(SqlParserID)
	}



	return localctx
}


// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SqlParserRULE_typeName)

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
		p.SetState(64)
		p.Match(SqlParserID)
	}



	return localctx
}


// IFunctionNameContext is an interface to support dynamic dispatch.
type IFunctionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionNameContext differentiates from other interfaces.
	IsFunctionNameContext()
}

type FunctionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionNameContext() *FunctionNameContext {
	var p = new(FunctionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_functionName
	return p
}

func (*FunctionNameContext) IsFunctionNameContext() {}

func NewFunctionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNameContext {
	var p = new(FunctionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_functionName

	return p
}

func (s *FunctionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNameContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *FunctionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterFunctionName(s)
	}
}

func (s *FunctionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitFunctionName(s)
	}
}

func (s *FunctionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitFunctionName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) FunctionName() (localctx IFunctionNameContext) {
	localctx = NewFunctionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SqlParserRULE_functionName)

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
		p.Match(SqlParserID)
	}



	return localctx
}


// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *AliasContext) TEXT_ALIAS() antlr.TerminalNode {
	return s.GetToken(SqlParserTEXT_ALIAS, 0)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SqlParserRULE_alias)
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
		p.SetState(68)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SqlParserID || _la == SqlParserTEXT_ALIAS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IJoinContext is an interface to support dynamic dispatch.
type IJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJoinContext differentiates from other interfaces.
	IsJoinContext()
}

type JoinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinContext() *JoinContext {
	var p = new(JoinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_join
	return p
}

func (*JoinContext) IsJoinContext() {}

func NewJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinContext {
	var p = new(JoinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_join

	return p
}

func (s *JoinContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinContext) JOIN() antlr.TerminalNode {
	return s.GetToken(SqlParserJOIN, 0)
}

func (s *JoinContext) INNER() antlr.TerminalNode {
	return s.GetToken(SqlParserINNER, 0)
}

func (s *JoinContext) LEFT() antlr.TerminalNode {
	return s.GetToken(SqlParserLEFT, 0)
}

func (s *JoinContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(SqlParserRIGHT, 0)
}

func (s *JoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *JoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterJoin(s)
	}
}

func (s *JoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitJoin(s)
	}
}

func (s *JoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitJoin(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) Join() (localctx IJoinContext) {
	localctx = NewJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SqlParserRULE_join)
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
		p.SetState(70)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 48)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 48))) & ((1 << (SqlParserINNER - 48)) | (1 << (SqlParserLEFT - 48)) | (1 << (SqlParserRIGHT - 48)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(71)
		p.Match(SqlParserJOIN)
	}



	return localctx
}


// IUnionContext is an interface to support dynamic dispatch.
type IUnionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionContext differentiates from other interfaces.
	IsUnionContext()
}

type UnionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionContext() *UnionContext {
	var p = new(UnionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_union
	return p
}

func (*UnionContext) IsUnionContext() {}

func NewUnionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionContext {
	var p = new(UnionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_union

	return p
}

func (s *UnionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionContext) UNION() antlr.TerminalNode {
	return s.GetToken(SqlParserUNION, 0)
}

func (s *UnionContext) ALL() antlr.TerminalNode {
	return s.GetToken(SqlParserALL, 0)
}

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitUnion(s)
	}
}

func (s *UnionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitUnion(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) Union() (localctx IUnionContext) {
	localctx = NewUnionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SqlParserRULE_union)
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
		p.SetState(73)
		p.Match(SqlParserUNION)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == SqlParserALL {
		{
			p.SetState(74)
			p.Match(SqlParserALL)
		}

	}



	return localctx
}


// IDecimalLiteralContext is an interface to support dynamic dispatch.
type IDecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalLiteralContext differentiates from other interfaces.
	IsDecimalLiteralContext()
}

type DecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalLiteralContext() *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SqlParserRULE_decimalLiteral)

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
		p.Match(SqlParserDECIMAL_LITERAL)
	}



	return localctx
}


// ITextLiteralContext is an interface to support dynamic dispatch.
type ITextLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextLiteralContext differentiates from other interfaces.
	IsTextLiteralContext()
}

type TextLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextLiteralContext() *TextLiteralContext {
	var p = new(TextLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_textLiteral
	return p
}

func (*TextLiteralContext) IsTextLiteralContext() {}

func NewTextLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextLiteralContext {
	var p = new(TextLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_textLiteral

	return p
}

func (s *TextLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *TextLiteralContext) TEXT_STRING() antlr.TerminalNode {
	return s.GetToken(SqlParserTEXT_STRING, 0)
}

func (s *TextLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TextLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTextLiteral(s)
	}
}

func (s *TextLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTextLiteral(s)
	}
}

func (s *TextLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitTextLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) TextLiteral() (localctx ITextLiteralContext) {
	localctx = NewTextLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SqlParserRULE_textLiteral)

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
		p.SetState(79)
		p.Match(SqlParserTEXT_STRING)
	}



	return localctx
}


// IBind_variablesContext is an interface to support dynamic dispatch.
type IBind_variablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBind_variablesContext differentiates from other interfaces.
	IsBind_variablesContext()
}

type Bind_variablesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBind_variablesContext() *Bind_variablesContext {
	var p = new(Bind_variablesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_bind_variables
	return p
}

func (*Bind_variablesContext) IsBind_variablesContext() {}

func NewBind_variablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bind_variablesContext {
	var p = new(Bind_variablesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_bind_variables

	return p
}

func (s *Bind_variablesContext) GetParser() antlr.Parser { return s.parser }

func (s *Bind_variablesContext) BIND_VARIABLE() antlr.TerminalNode {
	return s.GetToken(SqlParserBIND_VARIABLE, 0)
}

func (s *Bind_variablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bind_variablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Bind_variablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterBind_variables(s)
	}
}

func (s *Bind_variablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitBind_variables(s)
	}
}

func (s *Bind_variablesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitBind_variables(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) Bind_variables() (localctx IBind_variablesContext) {
	localctx = NewBind_variablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SqlParserRULE_bind_variables)

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
		p.SetState(81)
		p.Match(SqlParserBIND_VARIABLE)
	}



	return localctx
}


// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectStatement
	return p
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectStatement

	return p
}

func (s *SelectStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatementContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SqlParserSELECT, 0)
}

func (s *SelectStatementContext) SelectElements() ISelectElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectElementsContext)
}

func (s *SelectStatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(SqlParserFROM, 0)
}

func (s *SelectStatementContext) AllTableSources() []ITableSourcesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableSourcesContext)(nil)).Elem())
	var tst = make([]ITableSourcesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableSourcesContext)
		}
	}

	return tst
}

func (s *SelectStatementContext) TableSources(i int) ITableSourcesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableSourcesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableSourcesContext)
}

func (s *SelectStatementContext) Join() IJoinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJoinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJoinContext)
}

func (s *SelectStatementContext) ON() antlr.TerminalNode {
	return s.GetToken(SqlParserON, 0)
}

func (s *SelectStatementContext) LogicExpression() ILogicExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicExpressionContext)
}

func (s *SelectStatementContext) WhereClause() IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SelectStatementContext) GroupByClause() IGroupByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupByClauseContext)
}

func (s *SelectStatementContext) HavingClause() IHavingClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHavingClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHavingClauseContext)
}

func (s *SelectStatementContext) OrderByClause() IOrderByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *SelectStatementContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *SelectStatementContext) AllSelectStatement() []ISelectStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem())
	var tst = make([]ISelectStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectStatementContext)
		}
	}

	return tst
}

func (s *SelectStatementContext) SelectStatement(i int) ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *SelectStatementContext) Union() IUnionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSelectStatement(s)
	}
}

func (s *SelectStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSelectStatement(s)
	}
}

func (s *SelectStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitSelectStatement(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *SqlParser) SelectStatement() (localctx ISelectStatementContext) {
	return p.selectStatement(0)
}

func (p *SqlParser) selectStatement(_p int) (localctx ISelectStatementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISelectStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, SqlParserRULE_selectStatement, _p)

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
	{
		p.SetState(84)
		p.Match(SqlParserSELECT)
	}
	{
		p.SetState(85)
		p.SelectElements()
	}
	{
		p.SetState(86)
		p.Match(SqlParserFROM)
	}
	{
		p.SetState(87)
		p.TableSources()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(88)
			p.Join()
		}
		{
			p.SetState(89)
			p.TableSources()
		}
		{
			p.SetState(90)
			p.Match(SqlParserON)
		}
		{
			p.SetState(91)
			p.logicExpression(0)
		}


	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(95)
			p.WhereClause()
		}


	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(98)
			p.GroupByClause()
		}


	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(101)
			p.HavingClause()
		}


	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(104)
			p.OrderByClause()
		}


	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(107)
			p.LimitClause()
		}


	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSelectStatementContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_selectStatement)
			p.SetState(110)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(111)
				p.Union()
			}
			{
				p.SetState(112)
				p.selectStatement(2)
			}


		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}



	return localctx
}


// ISelectElementsContext is an interface to support dynamic dispatch.
type ISelectElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStar returns the star token.
	GetStar() antlr.Token 


	// SetStar sets the star token.
	SetStar(antlr.Token) 


	// IsSelectElementsContext differentiates from other interfaces.
	IsSelectElementsContext()
}

type SelectElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	star antlr.Token
}

func NewEmptySelectElementsContext() *SelectElementsContext {
	var p = new(SelectElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectElements
	return p
}

func (*SelectElementsContext) IsSelectElementsContext() {}

func NewSelectElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementsContext {
	var p = new(SelectElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectElements

	return p
}

func (s *SelectElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementsContext) GetStar() antlr.Token { return s.star }


func (s *SelectElementsContext) SetStar(v antlr.Token) { s.star = v }


func (s *SelectElementsContext) AllSelectElement() []ISelectElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectElementContext)(nil)).Elem())
	var tst = make([]ISelectElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectElementContext)
		}
	}

	return tst
}

func (s *SelectElementsContext) SelectElement(i int) ISelectElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectElementContext)
}

func (s *SelectElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSelectElements(s)
	}
}

func (s *SelectElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSelectElements(s)
	}
}

func (s *SelectElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitSelectElements(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) SelectElements() (localctx ISelectElementsContext) {
	localctx = NewSelectElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SqlParserRULE_selectElements)
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
	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserT__0:
		{
			p.SetState(119)

			var _m = p.Match(SqlParserT__0)

			localctx.(*SelectElementsContext).star = _m
		}


	case SqlParserT__6, SqlParserMAX, SqlParserSUM, SqlParserAVG, SqlParserMIN, SqlParserCOUNT, SqlParserDISTINCT, SqlParserCAST, SqlParserDECIMAL_LITERAL, SqlParserID, SqlParserTEXT_STRING, SqlParserBIND_VARIABLE:
		{
			p.SetState(120)
			p.SelectElement()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == SqlParserT__1 {
		{
			p.SetState(123)
			p.Match(SqlParserT__1)
		}
		{
			p.SetState(124)
			p.SelectElement()
		}


		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ISelectElementContext is an interface to support dynamic dispatch.
type ISelectElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectElementContext differentiates from other interfaces.
	IsSelectElementContext()
}

type SelectElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectElementContext() *SelectElementContext {
	var p = new(SelectElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectElement
	return p
}

func (*SelectElementContext) IsSelectElementContext() {}

func NewSelectElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementContext {
	var p = new(SelectElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectElement

	return p
}

func (s *SelectElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SelectElementContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *SelectElementContext) AS() antlr.TerminalNode {
	return s.GetToken(SqlParserAS, 0)
}

func (s *SelectElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSelectElement(s)
	}
}

func (s *SelectElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSelectElement(s)
	}
}

func (s *SelectElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitSelectElement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) SelectElement() (localctx ISelectElementContext) {
	localctx = NewSelectElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SqlParserRULE_selectElement)
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
		p.SetState(130)
		p.expr(0)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == SqlParserAS || _la == SqlParserID || _la == SqlParserTEXT_ALIAS {
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == SqlParserAS {
			{
				p.SetState(131)
				p.Match(SqlParserAS)
			}

		}
		{
			p.SetState(134)
			p.Alias()
		}

	}



	return localctx
}


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ExprContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExprContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *SqlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SqlParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, SqlParserRULE_expr, _p)
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
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(138)
			p.ColumnName()
		}


	case 2:
		{
			p.SetState(139)
			p.FunctionCall()
		}


	case 3:
		{
			p.SetState(140)
			p.Value()
		}


	case 4:
		{
			p.SetState(141)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(142)
			p.expr(0)
		}
		{
			p.SetState(143)
			p.Match(SqlParserT__7)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expr)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(148)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SqlParserT__0 || _la == SqlParserT__2) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(149)
					p.expr(5)
				}


			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expr)
				p.SetState(150)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(151)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SqlParserT__3 || _la == SqlParserT__4) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(152)
					p.expr(4)
				}


			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expr)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}

				{
					p.SetState(154)
					p.Match(SqlParserT__5)
				}

				{
					p.SetState(155)
					p.expr(3)
				}

			}

		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}



	return localctx
}


// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *ValueContext) TextLiteral() ITextLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextLiteralContext)
}

func (s *ValueContext) Bind_variables() IBind_variablesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBind_variablesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBind_variablesContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SqlParserRULE_value)

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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.DecimalLiteral()
		}


	case SqlParserTEXT_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.TextLiteral()
		}


	case SqlParserBIND_VARIABLE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.Bind_variables()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = SqlParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) AggregateFunction() IAggregateFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAggregateFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAggregateFunctionContext)
}

func (s *FunctionCallContext) CommonFunction() ICommonFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommonFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommonFunctionContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SqlParserRULE_functionCall)

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

	p.SetState(168)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserMAX, SqlParserSUM, SqlParserAVG, SqlParserMIN, SqlParserCOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.AggregateFunction()
		}


	case SqlParserDISTINCT, SqlParserCAST, SqlParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.CommonFunction()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IAggregateFunctionContext is an interface to support dynamic dispatch.
type IAggregateFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStarArg returns the starArg token.
	GetStarArg() antlr.Token 


	// SetStarArg sets the starArg token.
	SetStarArg(antlr.Token) 


	// IsAggregateFunctionContext differentiates from other interfaces.
	IsAggregateFunctionContext()
}

type AggregateFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	starArg antlr.Token
}

func NewEmptyAggregateFunctionContext() *AggregateFunctionContext {
	var p = new(AggregateFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_aggregateFunction
	return p
}

func (*AggregateFunctionContext) IsAggregateFunctionContext() {}

func NewAggregateFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateFunctionContext {
	var p = new(AggregateFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_aggregateFunction

	return p
}

func (s *AggregateFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateFunctionContext) GetStarArg() antlr.Token { return s.starArg }


func (s *AggregateFunctionContext) SetStarArg(v antlr.Token) { s.starArg = v }


func (s *AggregateFunctionContext) FunctionArg() IFunctionArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgContext)
}

func (s *AggregateFunctionContext) AVG() antlr.TerminalNode {
	return s.GetToken(SqlParserAVG, 0)
}

func (s *AggregateFunctionContext) MAX() antlr.TerminalNode {
	return s.GetToken(SqlParserMAX, 0)
}

func (s *AggregateFunctionContext) MIN() antlr.TerminalNode {
	return s.GetToken(SqlParserMIN, 0)
}

func (s *AggregateFunctionContext) SUM() antlr.TerminalNode {
	return s.GetToken(SqlParserSUM, 0)
}

func (s *AggregateFunctionContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SqlParserCOUNT, 0)
}

func (s *AggregateFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AggregateFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterAggregateFunction(s)
	}
}

func (s *AggregateFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitAggregateFunction(s)
	}
}

func (s *AggregateFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitAggregateFunction(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) AggregateFunction() (localctx IAggregateFunctionContext) {
	localctx = NewAggregateFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SqlParserRULE_aggregateFunction)
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

	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserMAX, SqlParserSUM, SqlParserAVG, SqlParserMIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			_la = p.GetTokenStream().LA(1)

			if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << SqlParserMAX) | (1 << SqlParserSUM) | (1 << SqlParserAVG) | (1 << SqlParserMIN))) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(171)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(172)
			p.FunctionArg()
		}
		{
			p.SetState(173)
			p.Match(SqlParserT__7)
		}


	case SqlParserCOUNT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(SqlParserCOUNT)
		}
		{
			p.SetState(176)
			p.Match(SqlParserT__6)
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SqlParserT__0:
			{
				p.SetState(177)

				var _m = p.Match(SqlParserT__0)

				localctx.(*AggregateFunctionContext).starArg = _m
			}


		case SqlParserT__6, SqlParserMAX, SqlParserSUM, SqlParserAVG, SqlParserMIN, SqlParserCOUNT, SqlParserDISTINCT, SqlParserCAST, SqlParserDECIMAL_LITERAL, SqlParserID, SqlParserTEXT_STRING, SqlParserBIND_VARIABLE:
			{
				p.SetState(178)
				p.FunctionArg()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(181)
			p.Match(SqlParserT__7)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ICommonFunctionContext is an interface to support dynamic dispatch.
type ICommonFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommonFunctionContext differentiates from other interfaces.
	IsCommonFunctionContext()
}

type CommonFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommonFunctionContext() *CommonFunctionContext {
	var p = new(CommonFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_commonFunction
	return p
}

func (*CommonFunctionContext) IsCommonFunctionContext() {}

func NewCommonFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommonFunctionContext {
	var p = new(CommonFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_commonFunction

	return p
}

func (s *CommonFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *CommonFunctionContext) FunctionName() IFunctionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *CommonFunctionContext) FunctionArg() IFunctionArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgContext)
}

func (s *CommonFunctionContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(SqlParserDISTINCT, 0)
}

func (s *CommonFunctionContext) CAST() antlr.TerminalNode {
	return s.GetToken(SqlParserCAST, 0)
}

func (s *CommonFunctionContext) AS() antlr.TerminalNode {
	return s.GetToken(SqlParserAS, 0)
}

func (s *CommonFunctionContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *CommonFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommonFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CommonFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterCommonFunction(s)
	}
}

func (s *CommonFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitCommonFunction(s)
	}
}

func (s *CommonFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitCommonFunction(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) CommonFunction() (localctx ICommonFunctionContext) {
	localctx = NewCommonFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SqlParserRULE_commonFunction)

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

	p.SetState(198)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.FunctionName()
		}
		{
			p.SetState(185)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(186)
			p.FunctionArg()
		}
		{
			p.SetState(187)
			p.Match(SqlParserT__7)
		}


	case SqlParserDISTINCT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Match(SqlParserDISTINCT)
		}
		{
			p.SetState(190)
			p.FunctionArg()
		}


	case SqlParserCAST:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.Match(SqlParserCAST)
		}
		{
			p.SetState(192)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(193)
			p.FunctionArg()
		}
		{
			p.SetState(194)
			p.Match(SqlParserAS)
		}
		{
			p.SetState(195)
			p.TypeName()
		}
		{
			p.SetState(196)
			p.Match(SqlParserT__7)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IFunctionArgContext is an interface to support dynamic dispatch.
type IFunctionArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgContext differentiates from other interfaces.
	IsFunctionArgContext()
}

type FunctionArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgContext() *FunctionArgContext {
	var p = new(FunctionArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_functionArg
	return p
}

func (*FunctionArgContext) IsFunctionArgContext() {}

func NewFunctionArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgContext {
	var p = new(FunctionArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_functionArg

	return p
}

func (s *FunctionArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *FunctionArgContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterFunctionArg(s)
	}
}

func (s *FunctionArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitFunctionArg(s)
	}
}

func (s *FunctionArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitFunctionArg(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) FunctionArg() (localctx IFunctionArgContext) {
	localctx = NewFunctionArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SqlParserRULE_functionArg)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.expr(0)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(201)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(202)
				p.expr(0)
			}


		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}



	return localctx
}


// ITableSourcesContext is an interface to support dynamic dispatch.
type ITableSourcesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableSourcesContext differentiates from other interfaces.
	IsTableSourcesContext()
}

type TableSourcesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableSourcesContext() *TableSourcesContext {
	var p = new(TableSourcesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableSources
	return p
}

func (*TableSourcesContext) IsTableSourcesContext() {}

func NewTableSourcesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableSourcesContext {
	var p = new(TableSourcesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableSources

	return p
}

func (s *TableSourcesContext) GetParser() antlr.Parser { return s.parser }

func (s *TableSourcesContext) AllTableSource() []ITableSourceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableSourceContext)(nil)).Elem())
	var tst = make([]ITableSourceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableSourceContext)
		}
	}

	return tst
}

func (s *TableSourcesContext) TableSource(i int) ITableSourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableSourceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableSourceContext)
}

func (s *TableSourcesContext) AllAlias() []IAliasContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAliasContext)(nil)).Elem())
	var tst = make([]IAliasContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAliasContext)
		}
	}

	return tst
}

func (s *TableSourcesContext) Alias(i int) IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *TableSourcesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSourcesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableSourcesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTableSources(s)
	}
}

func (s *TableSourcesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTableSources(s)
	}
}

func (s *TableSourcesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitTableSources(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) TableSources() (localctx ITableSourcesContext) {
	localctx = NewTableSourcesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SqlParserRULE_tableSources)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.TableSource()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(209)
			p.Alias()
		}


	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(212)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(213)
				p.TableSource()
			}
			p.SetState(215)
			p.GetErrorHandler().Sync(p)


			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(214)
					p.Alias()
				}


			}


		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}



	return localctx
}


// ITableSourceContext is an interface to support dynamic dispatch.
type ITableSourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableSourceContext differentiates from other interfaces.
	IsTableSourceContext()
}

type TableSourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableSourceContext() *TableSourceContext {
	var p = new(TableSourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableSource
	return p
}

func (*TableSourceContext) IsTableSourceContext() {}

func NewTableSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableSourceContext {
	var p = new(TableSourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableSource

	return p
}

func (s *TableSourceContext) GetParser() antlr.Parser { return s.parser }

func (s *TableSourceContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *TableSourceContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *TableSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableSourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTableSource(s)
	}
}

func (s *TableSourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTableSource(s)
	}
}

func (s *TableSourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitTableSource(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) TableSource() (localctx ITableSourceContext) {
	localctx = NewTableSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SqlParserRULE_tableSource)

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

	p.SetState(227)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.TableName()
		}


	case SqlParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(224)
			p.selectStatement(0)
		}
		{
			p.SetState(225)
			p.Match(SqlParserT__7)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SqlParserWHERE, 0)
}

func (s *WhereClauseContext) LogicExpression() ILogicExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicExpressionContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (s *WhereClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitWhereClause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SqlParserRULE_whereClause)

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
		p.SetState(229)
		p.Match(SqlParserWHERE)
	}
	{
		p.SetState(230)
		p.logicExpression(0)
	}



	return localctx
}


// ILogicExpressionContext is an interface to support dynamic dispatch.
type ILogicExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeftBracket returns the leftBracket token.
	GetLeftBracket() antlr.Token 

	// GetRightBracket returns the rightBracket token.
	GetRightBracket() antlr.Token 

	// GetLogicalOperator returns the logicalOperator token.
	GetLogicalOperator() antlr.Token 


	// SetLeftBracket sets the leftBracket token.
	SetLeftBracket(antlr.Token) 

	// SetRightBracket sets the rightBracket token.
	SetRightBracket(antlr.Token) 

	// SetLogicalOperator sets the logicalOperator token.
	SetLogicalOperator(antlr.Token) 


	// IsLogicExpressionContext differentiates from other interfaces.
	IsLogicExpressionContext()
}

type LogicExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	leftBracket antlr.Token
	rightBracket antlr.Token
	logicalOperator antlr.Token
}

func NewEmptyLogicExpressionContext() *LogicExpressionContext {
	var p = new(LogicExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_logicExpression
	return p
}

func (*LogicExpressionContext) IsLogicExpressionContext() {}

func NewLogicExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicExpressionContext {
	var p = new(LogicExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_logicExpression

	return p
}

func (s *LogicExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicExpressionContext) GetLeftBracket() antlr.Token { return s.leftBracket }

func (s *LogicExpressionContext) GetRightBracket() antlr.Token { return s.rightBracket }

func (s *LogicExpressionContext) GetLogicalOperator() antlr.Token { return s.logicalOperator }


func (s *LogicExpressionContext) SetLeftBracket(v antlr.Token) { s.leftBracket = v }

func (s *LogicExpressionContext) SetRightBracket(v antlr.Token) { s.rightBracket = v }

func (s *LogicExpressionContext) SetLogicalOperator(v antlr.Token) { s.logicalOperator = v }


func (s *LogicExpressionContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LogicExpressionContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogicExpressionContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *LogicExpressionContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(SqlParserBETWEEN, 0)
}

func (s *LogicExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(SqlParserAND, 0)
}

func (s *LogicExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *LogicExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(SqlParserIN, 0)
}

func (s *LogicExpressionContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *LogicExpressionContext) LIKE() antlr.TerminalNode {
	return s.GetToken(SqlParserLIKE, 0)
}

func (s *LogicExpressionContext) IS() antlr.TerminalNode {
	return s.GetToken(SqlParserIS, 0)
}

func (s *LogicExpressionContext) NULL() antlr.TerminalNode {
	return s.GetToken(SqlParserNULL, 0)
}

func (s *LogicExpressionContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(SqlParserEXISTS, 0)
}

func (s *LogicExpressionContext) AllLogicExpression() []ILogicExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILogicExpressionContext)(nil)).Elem())
	var tst = make([]ILogicExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILogicExpressionContext)
		}
	}

	return tst
}

func (s *LogicExpressionContext) LogicExpression(i int) ILogicExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILogicExpressionContext)
}

func (s *LogicExpressionContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMENT, 0)
}

func (s *LogicExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(SqlParserOR, 0)
}

func (s *LogicExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LogicExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterLogicExpression(s)
	}
}

func (s *LogicExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitLogicExpression(s)
	}
}

func (s *LogicExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitLogicExpression(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *SqlParser) LogicExpression() (localctx ILogicExpressionContext) {
	return p.logicExpression(0)
}

func (p *SqlParser) logicExpression(_p int) (localctx ILogicExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLogicExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogicExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, SqlParserRULE_logicExpression, _p)
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
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(233)
			p.expr(0)
		}
		{
			p.SetState(234)
			p.ComparisonOperator()
		}
		{
			p.SetState(235)
			p.expr(0)
		}


	case 2:
		{
			p.SetState(237)
			p.expr(0)
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == SqlParserNOT {
			{
				p.SetState(238)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(241)
			p.Match(SqlParserBETWEEN)
		}
		{
			p.SetState(242)
			p.expr(0)
		}
		{
			p.SetState(243)
			p.Match(SqlParserAND)
		}
		{
			p.SetState(244)
			p.expr(0)
		}


	case 3:
		{
			p.SetState(246)
			p.expr(0)
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == SqlParserNOT {
			{
				p.SetState(247)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(250)
			p.Match(SqlParserIN)
		}
		{
			p.SetState(251)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(252)
			p.expr(0)
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == SqlParserT__1 {
			{
				p.SetState(253)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(254)
				p.expr(0)
			}


			p.SetState(259)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(260)
			p.Match(SqlParserT__7)
		}


	case 4:
		{
			p.SetState(262)
			p.expr(0)
		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == SqlParserNOT {
			{
				p.SetState(263)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(266)
			p.Match(SqlParserIN)
		}
		{
			p.SetState(267)
			p.Match(SqlParserT__6)
		}
		{
			p.SetState(268)
			p.selectStatement(0)
		}
		{
			p.SetState(269)
			p.Match(SqlParserT__7)
		}


	case 5:
		{
			p.SetState(271)
			p.expr(0)
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == SqlParserNOT {
			{
				p.SetState(272)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(275)
			p.Match(SqlParserLIKE)
		}
		{
			p.SetState(276)
			p.expr(0)
		}


	case 6:
		{
			p.SetState(278)
			p.expr(0)
		}
		{
			p.SetState(279)
			p.Match(SqlParserIS)
		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == SqlParserNOT {
			{
				p.SetState(280)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(283)
			p.Match(SqlParserNULL)
		}


	case 7:
		{
			p.SetState(285)
			p.Match(SqlParserEXISTS)
		}
		{
			p.SetState(286)

			var _m = p.Match(SqlParserT__6)

			localctx.(*LogicExpressionContext).leftBracket = _m
		}
		{
			p.SetState(287)
			p.selectStatement(0)
		}
		{
			p.SetState(288)

			var _m = p.Match(SqlParserT__7)

			localctx.(*LogicExpressionContext).rightBracket = _m
		}


	case 8:
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == SqlParserCOMMENT {
			{
				p.SetState(290)
				p.Match(SqlParserCOMMENT)
			}

		}
		{
			p.SetState(293)

			var _m = p.Match(SqlParserT__6)

			localctx.(*LogicExpressionContext).leftBracket = _m
		}
		{
			p.SetState(294)
			p.logicExpression(0)
		}
		{
			p.SetState(295)

			var _m = p.Match(SqlParserT__7)

			localctx.(*LogicExpressionContext).rightBracket = _m
		}


	case 9:
		{
			p.SetState(297)
			p.Match(SqlParserNOT)
		}
		{
			p.SetState(298)
			p.logicExpression(3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(307)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogicExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_logicExpression)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(302)

					var _m = p.Match(SqlParserAND)

					localctx.(*LogicExpressionContext).logicalOperator = _m
				}
				{
					p.SetState(303)
					p.logicExpression(3)
				}


			case 2:
				localctx = NewLogicExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_logicExpression)
				p.SetState(304)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(305)

					var _m = p.Match(SqlParserOR)

					localctx.(*LogicExpressionContext).logicalOperator = _m
				}
				{
					p.SetState(306)
					p.logicExpression(2)
				}

			}

		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}



	return localctx
}


// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SqlParserRULE_comparisonOperator)
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
		p.SetState(312)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << SqlParserT__8) | (1 << SqlParserT__9) | (1 << SqlParserT__10) | (1 << SqlParserT__11) | (1 << SqlParserT__12) | (1 << SqlParserT__13) | (1 << SqlParserT__14) | (1 << SqlParserT__15))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IGroupByClauseContext is an interface to support dynamic dispatch.
type IGroupByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupByClauseContext differentiates from other interfaces.
	IsGroupByClauseContext()
}

type GroupByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByClauseContext() *GroupByClauseContext {
	var p = new(GroupByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_groupByClause
	return p
}

func (*GroupByClauseContext) IsGroupByClauseContext() {}

func NewGroupByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByClauseContext {
	var p = new(GroupByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_groupByClause

	return p
}

func (s *GroupByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByClauseContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SqlParserGROUP, 0)
}

func (s *GroupByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *GroupByClauseContext) AllGroupByItem() []IGroupByItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroupByItemContext)(nil)).Elem())
	var tst = make([]IGroupByItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroupByItemContext)
		}
	}

	return tst
}

func (s *GroupByClauseContext) GroupByItem(i int) IGroupByItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupByItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroupByItemContext)
}

func (s *GroupByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GroupByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterGroupByClause(s)
	}
}

func (s *GroupByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitGroupByClause(s)
	}
}

func (s *GroupByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitGroupByClause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) GroupByClause() (localctx IGroupByClauseContext) {
	localctx = NewGroupByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SqlParserRULE_groupByClause)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(SqlParserGROUP)
	}
	{
		p.SetState(315)
		p.Match(SqlParserBY)
	}
	{
		p.SetState(316)
		p.GroupByItem()
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(317)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(318)
				p.GroupByItem()
			}


		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}



	return localctx
}


// IGroupByItemContext is an interface to support dynamic dispatch.
type IGroupByItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupByItemContext differentiates from other interfaces.
	IsGroupByItemContext()
}

type GroupByItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByItemContext() *GroupByItemContext {
	var p = new(GroupByItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_groupByItem
	return p
}

func (*GroupByItemContext) IsGroupByItemContext() {}

func NewGroupByItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByItemContext {
	var p = new(GroupByItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_groupByItem

	return p
}

func (s *GroupByItemContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByItemContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GroupByItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GroupByItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterGroupByItem(s)
	}
}

func (s *GroupByItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitGroupByItem(s)
	}
}

func (s *GroupByItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitGroupByItem(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) GroupByItem() (localctx IGroupByItemContext) {
	localctx = NewGroupByItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SqlParserRULE_groupByItem)

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
		p.SetState(324)
		p.expr(0)
	}



	return localctx
}


// IHavingClauseContext is an interface to support dynamic dispatch.
type IHavingClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHavingClauseContext differentiates from other interfaces.
	IsHavingClauseContext()
}

type HavingClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHavingClauseContext() *HavingClauseContext {
	var p = new(HavingClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_havingClause
	return p
}

func (*HavingClauseContext) IsHavingClauseContext() {}

func NewHavingClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingClauseContext {
	var p = new(HavingClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_havingClause

	return p
}

func (s *HavingClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *HavingClauseContext) HAVING() antlr.TerminalNode {
	return s.GetToken(SqlParserHAVING, 0)
}

func (s *HavingClauseContext) LogicExpression() ILogicExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicExpressionContext)
}

func (s *HavingClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *HavingClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterHavingClause(s)
	}
}

func (s *HavingClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitHavingClause(s)
	}
}

func (s *HavingClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitHavingClause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) HavingClause() (localctx IHavingClauseContext) {
	localctx = NewHavingClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SqlParserRULE_havingClause)

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
		p.SetState(326)
		p.Match(SqlParserHAVING)
	}
	{
		p.SetState(327)
		p.logicExpression(0)
	}



	return localctx
}


// IOrderByClauseContext is an interface to support dynamic dispatch.
type IOrderByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByClauseContext differentiates from other interfaces.
	IsOrderByClauseContext()
}

type OrderByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByClauseContext() *OrderByClauseContext {
	var p = new(OrderByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_orderByClause
	return p
}

func (*OrderByClauseContext) IsOrderByClauseContext() {}

func NewOrderByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByClauseContext {
	var p = new(OrderByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_orderByClause

	return p
}

func (s *OrderByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByClauseContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SqlParserORDER, 0)
}

func (s *OrderByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *OrderByClauseContext) AllOrderByExpression() []IOrderByExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrderByExpressionContext)(nil)).Elem())
	var tst = make([]IOrderByExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrderByExpressionContext)
		}
	}

	return tst
}

func (s *OrderByClauseContext) OrderByExpression(i int) IOrderByExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrderByExpressionContext)
}

func (s *OrderByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OrderByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterOrderByClause(s)
	}
}

func (s *OrderByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitOrderByClause(s)
	}
}

func (s *OrderByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitOrderByClause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) OrderByClause() (localctx IOrderByClauseContext) {
	localctx = NewOrderByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SqlParserRULE_orderByClause)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Match(SqlParserORDER)
	}
	{
		p.SetState(330)
		p.Match(SqlParserBY)
	}
	{
		p.SetState(331)
		p.OrderByExpression()
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(332)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(333)
				p.OrderByExpression()
			}


		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}



	return localctx
}


// IOrderByExpressionContext is an interface to support dynamic dispatch.
type IOrderByExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOrder returns the order token.
	GetOrder() antlr.Token 


	// SetOrder sets the order token.
	SetOrder(antlr.Token) 


	// IsOrderByExpressionContext differentiates from other interfaces.
	IsOrderByExpressionContext()
}

type OrderByExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	order antlr.Token
}

func NewEmptyOrderByExpressionContext() *OrderByExpressionContext {
	var p = new(OrderByExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_orderByExpression
	return p
}

func (*OrderByExpressionContext) IsOrderByExpressionContext() {}

func NewOrderByExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByExpressionContext {
	var p = new(OrderByExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_orderByExpression

	return p
}

func (s *OrderByExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByExpressionContext) GetOrder() antlr.Token { return s.order }


func (s *OrderByExpressionContext) SetOrder(v antlr.Token) { s.order = v }


func (s *OrderByExpressionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrderByExpressionContext) ASC() antlr.TerminalNode {
	return s.GetToken(SqlParserASC, 0)
}

func (s *OrderByExpressionContext) DESC() antlr.TerminalNode {
	return s.GetToken(SqlParserDESC, 0)
}

func (s *OrderByExpressionContext) NULLS() antlr.TerminalNode {
	return s.GetToken(SqlParserNULLS, 0)
}

func (s *OrderByExpressionContext) FIRST() antlr.TerminalNode {
	return s.GetToken(SqlParserFIRST, 0)
}

func (s *OrderByExpressionContext) LAST() antlr.TerminalNode {
	return s.GetToken(SqlParserLAST, 0)
}

func (s *OrderByExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OrderByExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterOrderByExpression(s)
	}
}

func (s *OrderByExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitOrderByExpression(s)
	}
}

func (s *OrderByExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitOrderByExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) OrderByExpression() (localctx IOrderByExpressionContext) {
	localctx = NewOrderByExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SqlParserRULE_orderByExpression)
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
		p.SetState(339)
		p.expr(0)
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(340)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderByExpressionContext).order = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserASC || _la == SqlParserDESC) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OrderByExpressionContext).order = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(343)
			p.Match(SqlParserNULLS)
		}
		{
			p.SetState(344)
			p.Match(SqlParserFIRST)
		}


	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(345)
			p.Match(SqlParserNULLS)
		}
		{
			p.SetState(346)
			p.Match(SqlParserLAST)
		}



	}



	return localctx
}


// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOffset returns the offset rule contexts.
	GetOffset() IDecimalLiteralContext

	// GetLimit returns the limit rule contexts.
	GetLimit() IDecimalLiteralContext


	// SetOffset sets the offset rule contexts.
	SetOffset(IDecimalLiteralContext)

	// SetLimit sets the limit rule contexts.
	SetLimit(IDecimalLiteralContext)


	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	offset IDecimalLiteralContext 
	limit IDecimalLiteralContext 
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) GetOffset() IDecimalLiteralContext { return s.offset }

func (s *LimitClauseContext) GetLimit() IDecimalLiteralContext { return s.limit }


func (s *LimitClauseContext) SetOffset(v IDecimalLiteralContext) { s.offset = v }

func (s *LimitClauseContext) SetLimit(v IDecimalLiteralContext) { s.limit = v }


func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SqlParserLIMIT, 0)
}

func (s *LimitClauseContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(SqlParserOFFSET, 0)
}

func (s *LimitClauseContext) AllDecimalLiteral() []IDecimalLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem())
	var tst = make([]IDecimalLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecimalLiteralContext)
		}
	}

	return tst
}

func (s *LimitClauseContext) DecimalLiteral(i int) IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SqlParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SqlParserRULE_limitClause)

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
		p.SetState(349)
		p.Match(SqlParserLIMIT)
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.SetState(353)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(350)

				var _x = p.DecimalLiteral()


				localctx.(*LimitClauseContext).offset = _x
			}
			{
				p.SetState(351)
				p.Match(SqlParserT__1)
			}


		}
		{
			p.SetState(355)

			var _x = p.DecimalLiteral()


			localctx.(*LimitClauseContext).limit = _x
		}


	case 2:
		{
			p.SetState(356)

			var _x = p.DecimalLiteral()


			localctx.(*LimitClauseContext).limit = _x
		}
		{
			p.SetState(357)
			p.Match(SqlParserOFFSET)
		}
		{
			p.SetState(358)

			var _x = p.DecimalLiteral()


			localctx.(*LimitClauseContext).offset = _x
		}

	}



	return localctx
}


func (p *SqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
			var t *SelectStatementContext = nil
			if localctx != nil { t = localctx.(*SelectStatementContext) }
			return p.SelectStatement_Sempred(t, predIndex)

	case 13:
			var t *ExprContext = nil
			if localctx != nil { t = localctx.(*ExprContext) }
			return p.Expr_Sempred(t, predIndex)

	case 22:
			var t *LogicExpressionContext = nil
			if localctx != nil { t = localctx.(*LogicExpressionContext) }
			return p.LogicExpression_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SqlParser) SelectStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) LogicExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}


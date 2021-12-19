// Here is an array containing both numbers as well as other arrays, which in turn contain numbers and arrays:
// array = [  1,
//            2,
//            3,
//            [4, 5, 6], //s1
//            7,
//            [8, //s2
//              [9, 10, 11,
//                [12, 13, 14]
//              ]
//            ],
//            [15, 16, 17, 18, 19, //s3
//              [20, 21, 22, //a
//                [23, 24, 25, //b
//                  [26, 27, 29] //c
//                ], 30, 31
//              ], 32
//            ], 33
//         ]
// Write a recursive function that prints all the numbers (and just numbers).

package main

import (
	"fmt"
)

func makeMixedSlice() []interface{} {
	var mixedSlice []interface{}
	mixedSlice = append(mixedSlice, 1, 2.222, "three")
	var sli1 []interface{}
	sli1 = append(sli1, 4, 5, 6)
	mixedSlice = append(mixedSlice, sli1, 7)

	var sli2b []interface{}
	sli2b = append(sli2b, 12, 13, 14)
	var sli2a []interface{}
	sli2a = append(sli2a, 9, 10, 11, sli2b)
	var sli2 []interface{}
	sli2 = append(sli2, 8, sli2a)
	mixedSlice = append(mixedSlice, sli2)

	var sli3c []interface{}
	sli3c = append(sli3c, 26, 27, 29)
	var sli3b []interface{}
	sli3b = append(sli3b, 23, 24, 25, sli3c)
	var sli3a []interface{}
	sli3a = append(sli3a, 20, 21, 22, sli3b, 30, 31)
	var sli3 []interface{}
	sli3 = append(sli3, 15, 16, 17, 18, 19, sli3a, 32)
	mixedSlice = append(mixedSlice, sli3, 33)

	return mixedSlice
}

func printNestedNums(mixedSlice []interface{}) {
	for _, val := range mixedSlice {
		if innerMixedSlice, ok := val.([]interface{}); ok {
			printNestedNums(innerMixedSlice)
		} else {
			fmt.Println(val)
		}

		// switch val.(type) {
		// case []interface{}:
		//   innerMixedSlice := val.([]interface{}) // confident i've already checked type
		//   printNestedNums(innerMixedSlice)
		// default:
		//   fmt.Println(val)
		// }
	}
}

func main() {
	mixedSlice := makeMixedSlice()
	fmt.Println(mixedSlice)
	printNestedNums(mixedSlice)
}

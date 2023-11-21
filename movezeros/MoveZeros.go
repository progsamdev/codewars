/*
codeWars: https://www.codewars.com/users/sammsilva
problem: https://www.codewars.com/kata/52597aa56021e91c93000cb0/solutions/go
*/

package movezeros

func MoveZeros(arr []int) []int {
	var auxArr []int
	var count int
	for i, v := range arr {
		if arr[i] != 0 {
			auxArr = append(auxArr, v)
			continue
		}
		count++
	}
	auxArr = append(auxArr, make([]int, count)...)
	return auxArr
}

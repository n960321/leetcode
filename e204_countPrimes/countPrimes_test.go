package e204_countPrimes

import "testing"

/*

https://leetcode.com/problems/count-primes/

Count the number of prime numbers less than a non-negative number, n.



Example 1:

Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
Example 2:

Input: n = 0
Output: 0
Example 3:

Input: n = 1
Output: 0
*/

func countPrimes(n int) int {
	//先建造一個map 指定通通都是 質數
	flagArray := make([]bool, n)
	for q := 0; q < len(flagArray); q++ {
		flagArray[q] = true
	}
	result := 0
	for i := 2; i < n; i++ {
		if flagArray[i] == true {
			// is Primes
			result++
			// rm it's multiples
			j := 2
			// 將此質數的倍數 通通拉出來改狀態
			for i*j < n {
				flagArray[i*j] = false
				j++
			}
		}
	}
	return result
}

func countPrimes_v1(n int) int {
	if n <= 2 {
		return 0
	}
	primes := []int{}
OUTER:
	for v := 2; v <= n; v++ {

		if len(primes) == 0 {
			primes = append(primes, v)
			continue
		}
		for _, vv := range primes {
			if v%vv == 0 {
				continue OUTER
			}
		}
		primes = append(primes, v)
	}

	return len(primes)
}

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 10,
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

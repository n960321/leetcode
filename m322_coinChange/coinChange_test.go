package m322_coinChange

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.com/problems/coin-change/

You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Example 3:

Input: coins = [1], amount = 0
Output: 0
Example 4:

Input: coins = [1], amount = 1
Output: 1
Example 5:

Input: coins = [1], amount = 2
Output: 2
---
1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

*/
func coinChange_v1(coins []int, amount int) int {
	r := -1
	l := []int{}
	if len(coins) < 1 {
		return -1
	}

	sort.Ints(coins)
OUTER:
	for amount > 0 {
		for i := len(coins) - 1; i >= 0; i-- {
			v := coins[i]
			if amount-v >= 0 {
				amount -= v
				l = append(l, v)
				fmt.Printf("===== %+v , %v \n", l, amount)
				continue OUTER
			}
			if amount-v < 0 {
				coins = append(coins[:i], coins[i+1:]...)
				fmt.Printf("----- %+v , %v\n", coins, amount)
				continue OUTER
			}
		}
		if len(coins) == 0 {
			return -1
		}
	}
	r = len(l)
	return r
}

const maxInt = int(^uint(0)>>1) - 1

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = maxInt
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
				fmt.Printf("===== [%v] %+v \n", i, dp[i])
			}
		}
	}

	if dp[amount] == maxInt {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases
		// {
		// 	name: "1",
		// 	args: args{coins: []int{1, 2, 5}, amount: 11},
		// 	want: 3,
		// },
		{
			name: "2",
			args: args{coins: []int{186, 419, 83, 408}, amount: 6249},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

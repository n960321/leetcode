package e121_Best_Time_to_Buy_and_Sell_Stock

import "testing"

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/
// func TestBestTimeToBuyAndSell(t *testing.T) {
// 	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
// 	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
// 	fmt.Println(maxProfit([]int{2, 4, 1}))
// }

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				prices: []int{2, 1, 2, 1, 0, 1, 2},
			},
			want: 2,
		},
		{
			name: "5",
			args: args{
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func maxProfit(prices []int) int {
	ans := 0
	buy := 0
	sell := 0

	for i := range prices {
		if i == 0 || (buy > prices[i]) {
			buy = prices[i]
			sell = 0
			continue
		}

		if sell == 0 || sell < prices[i] {
			sell = prices[i]
			if _ans := sell - buy; _ans > ans {
				ans = _ans
				continue
			}
			sell = 0
		}
	}
	if ans < 0 {
		ans = 0
	}
	return ans
}

// minP := 999999
// maxP := 0
// i := 0
// for i < len(prices) {
// 	if prices[i] < minP {
// 		minP = prices[i]
// 	} else if prices[i] - minP > maxP {
// 		maxP = prices[i] - minP
// 	}
// 	i += 1
// }
// return maxP

// https://leetcode.com/problems/online-stock-span/

package leetcode_go

type StockSpanner struct {
	stk []StockAndCnt
}

type StockAndCnt struct {
	price int
	cnt   int
}

func ConstructorP901() StockSpanner {
	return StockSpanner{stk: []StockAndCnt{}}
}

func (this *StockSpanner) Next(price int) int {
	cur := 1
	for len(this.stk) > 0 && price >= this.stk[len(this.stk)-1].price {
		cur += this.stk[len(this.stk)-1].cnt
		this.stk = this.stk[:len(this.stk)-1]
	}
	this.stk = append(this.stk, StockAndCnt{price, cur})
	return cur
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

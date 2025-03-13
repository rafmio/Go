package arrays

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{From: "Chrys",
			To:  "Riya",
			Sum: 100,
		},
		{
			From: "Adil",
			To:   "Chrys",
			Sum:  200,
		}}

	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}

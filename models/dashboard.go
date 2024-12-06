package models

type Summary struct {
	TotalUser   int
	TotalSales  int
	TotalOrder  int
	TotalProduc int
}

type Revenue struct {
	Month        string
	TotalEarning int
}

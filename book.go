package malgova

// AllocateCash book
func (b *Book) AllocateCash(Money float64) {
	b.CashAllocated = Money
	b.Cash = b.CashAllocated
}

// PlaceMarketOrder book
func (b *Book) plac
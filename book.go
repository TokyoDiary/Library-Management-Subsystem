package malgova

// AllocateCash book
func (b *Book) AllocateCash(Money float64) {
	b.CashAllocated = Money
	b.Cash = b.CashAllocated
}

// PlaceMarketOrder book
func (b *Book) placeMarketOrder(Qty int) {
	b.PendingOrderQuantity = Qty
	b.IsMarketOrder = true
}

// PlaceMarketOrder book
func (b *Book) placeLimitOrder(Qty int, Price float64) {
	b.PendingOrderQuantity = Qty
	b.IsMar
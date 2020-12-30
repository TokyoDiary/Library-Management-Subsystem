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
	b.IsMarketOrder = false
	b.PendingOrderPrice = Price
}

// QuantityAffordable book
func (b *Book) QuantityAffordable(Price float64) int {
	if Price <= b.Cash {
		return int(b.Cash / Price)
	}
	return 0
}

// Buy Order
func (b *Book) Buy(Qty int) {
	
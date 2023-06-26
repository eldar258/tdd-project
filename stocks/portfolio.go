package stocks

type Portfolio []Money

func (p Portfolio) Add(m Money) Portfolio {
	return append(p, m)
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, m := range p {
		total += m.amount
	}
	return Money{amount: total, currency: currency}
}

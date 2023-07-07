package stocks

import "errors"

type Portfolio []Money

func (p Portfolio) Add(m Money) Portfolio {
	return append(p, m)
}

func (p Portfolio) Evaluate(currency string) (Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)

	for _, m := range p {
		if value, ok := convert(m, currency); ok {
			total += value
		} else {
			failedConversions = append(failedConversions, m.currency, "->", currency, ",")
		}
	}

	if len(failedConversions) == 0 {
		return Money{amount: total, currency: currency}, nil
	}

	failures := "["
	for _, el := range failedConversions {
		failures += el
	}
	failures += "]"
	return NewMoney(0, ""), errors.New("Missing exchange rate(s):" + failures)
}

func convert(money Money, currency string) (float64, bool) {
	exchangeRates := map[string]float64{
		"EUR->USD": 1.2,
		"USD->KRW": 1100,
	}

	if currency == money.currency {
		return money.amount, true
	}
	rate, ok := exchangeRates[money.currency+"->"+currency]
	return money.amount * rate, ok
}

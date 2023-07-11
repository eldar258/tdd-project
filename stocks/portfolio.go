package stocks

import "errors"

type Portfolio []Money

func (p Portfolio) Add(m Money) Portfolio {
	return append(p, m)
}

func (p Portfolio) Evaluate(bank Bank, currency string) (*Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)

	for _, m := range p {
		if value, err := bank.Convert(m, currency); err == nil {
			total += value.amount
		} else {
			failedConversions = append(failedConversions, err.Error(), ",")
		}
	}

	if len(failedConversions) == 0 {
		return &Money{amount: total, currency: currency}, nil
	}

	failures := "["
	for _, el := range failedConversions {
		failures += el
	}
	failures += "]"
	return nil, errors.New("Missing exchange rate(s):" + failures)
}

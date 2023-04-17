package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king", "queen", "jack", "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sumCards := ParseCard(card1) + ParseCard(card2)
	parsedDealerCard := ParseCard(dealerCard)

	switch {
	case ParseCard(card1) == 11 && ParseCard(card2) == 11:
		return "P"
	case sumCards == 21 && (parsedDealerCard != 11 && parsedDealerCard != 10):
		return "W"
	case sumCards == 21 && (parsedDealerCard == 11 || parsedDealerCard == 10):
		return "S"
	case sumCards >= 17 && sumCards <= 20:
		return "S"
	case sumCards >= 12 && sumCards <= 16 && parsedDealerCard >= 7:
		return "H"
	case sumCards >= 12 && sumCards <= 16:
		return "S"
	default:
		return "H"
	}
}

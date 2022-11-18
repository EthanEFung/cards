package cards

import (
    "sort"
    "math/rand"
)

// Suit represents the french suit of a standard 52 card deck.
type Suit int

const (
	Suitless Suit = iota
	Diamonds
	Clubs
	Hearts
	Spades
)

// Rank represents one of 13 standard card ranks.
type Rank int

const (
	Rankless Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

//go:generate stringer -type=Suit,Rank

// Card represents one of 52 standard french suited cards or a joker.
type Card struct {
	Suit Suit
	Rank Rank
}

// String returns a string representation of the card.
func (c Card) String() string {
	if c.Rank == Rankless && c.Suit == Suitless {
		return "Joker"
	}
	return c.Rank.String() + " of " + c.Suit.String()
}

var suits = []Suit{Spades, Diamonds, Clubs, Hearts}

var ranks = []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// Deck represents an array of cards
type Deck []Card

// New returns 52 standard french suited cards in standard order:
// Spades before Diamonds, Diamonds before Clubs, and Clubs before Hearts.
// Cards are also sorted by ranks: Ace through King in ascending order.
func New() Deck {
	cards := []Card{}
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, Card{suit, rank})
		}
	}
	return cards
}

// AddJokers will append the number of jokers specified in the first argument.
func (d *Deck) AddJokers(n int) {
    for i := 0; i < n; i++ {
        *d = append(*d, Card{})
    }
}

// SortBy will order the deck according to the `less` function passed in the
// first argument
func (d Deck) SortBy(less func(i, j int) bool) Deck {
    sort.Slice(d, less)
    return d
}

// Shuffle will randomize the ordering of the deck when given a uniq int64 to seed the
// RNG algorithm.
func (d Deck) Shuffle(seed int64) {
    rand.Seed(seed)
    rand.Shuffle(len(d), func(i, j int) {
        d[i], d[j] = d[j], d[i]
    })
}

// Omit will remove the card if the `remove` function passed as the first argument
// returns true. The `remove` function passes the index of the card in
// question as the first argument in its signature.
func (d *Deck) Omit(remove func(i int) bool) {
    remaining := Deck{}
    for i, card := range *d {
        if remove(i) {
            continue
        }
        remaining = append(remaining, card)
    }
    *d = remaining
}

// Multiply will copy the deck of cards by the multiple specified in the first
// argument of the function.
func (d *Deck) Multiply(x int) {
    cards := make(Deck, 0, len(*d)*x)
    for i := 0; i < x; i++ {
        cards = append(cards, *d...)
    }
    *d = cards
}


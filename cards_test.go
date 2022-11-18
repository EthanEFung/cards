package cards

import (
    "testing"
)

func TestCard(t *testing.T) {
    ace := Card{Spades, Ace}
    want := "Ace of Spades"
    have := ace.String()
    if want != have {
        t.Fatalf("Expected Card.String() to return '"+want+"', but got: '"+have+"'")
    }

    joker := Card{}
    want = "Joker"
    have = joker.String()
    if want != have {
        t.Fatalf("Expected Suitless and Rankless card to return 'Joker', but returned: '%s'", joker.String())
    }
}

func TestNewLength(t *testing.T) {
    deck := New()
    if len(deck) != 52 {
        t.Fatalf("Expected New() to return 52 cards, but got %d", len(deck))
    }
}

func TestNewSuits(t *testing.T) {
    deck := New()
    counts := make(map[Suit]int, len(deck))
    for _, card := range deck {
        counts[card.Suit]++
    }
    if len(counts) != 4 {
        t.Fatalf("Expected New() to return 4 different kinds of suits but returned %d", len(counts))
    }
    for suit, count := range counts {
        if count != 13 {
            t.Fatalf("Expected New() to return 13 %v but instead returned %d", suit.String(), count)
        }
    }
}

func TestNewRanks(t *testing.T) {
    deck := New()
    counts := make(map[Rank]int, len(deck))
    for _, card := range deck {
        counts[card.Rank]++
    }
    if len(counts) != 13 {
        t.Fatalf("Expected New() to return 4 different kinds of ranks but returned %d", len(counts))
    }
    for rank, count := range counts {
        if count != 4 {
            t.Fatalf("Expected New() to return 4 %vs but instead returned %d", rank.String(), count)
        }
    }
}

func TestDeckAddJokers(t *testing.T) {
    deck := New()
    deck.AddJokers(0)
    if len(deck) != 52 {
        t.Fatalf("Called deck.AddJokers(0) not expecting jokers to be added, but deck length is %d", len(deck))
    }
    deck.AddJokers(1)
    if len(deck) != 53 {
        t.Fatalf("Called deck.AddJokers(1) expecting 1 joker to be added, but deck length is %d", len(deck))
    }
    deck = New()
    deck.AddJokers(2)
    if len(deck) != 54 {
        t.Fatalf("Called deck.AddJokers(2) expecting 2 jokers to be added, but deck length is %d", len(deck))
    }
}

func TestDeckSortBy(t *testing.T) {
    deck := New()
    
    // TODO: create a struct of some sensible sorting functions
    aceLow := func(i, j int) bool {
        cardA, cardB := deck[i], deck[j]
        if cardA.Suit == cardB.Suit {
            return cardA.Rank < cardB.Rank
        }
        return cardA.Suit < cardB.Suit
    }
    inspect := func(d Deck) map[string]int {
        indices := make(map[string]int, len(d))
        for i, card := range d {
            indices[card.String()] = i
        }
        return indices
    }

    deck.SortBy(aceLow)

    order := inspect(deck)

    if order["Ace of Hearts"] > order["King of Hearts"] {
        t.Fatalf("Called deck.SortBy(aceLow) expecting the Ace of Hearts to come before the King")
    }
    if order["Ace of Hearts"] > order["Two of Hearts"] {
        t.Fatalf("Called deck.SortBy(aceLow) expecting the Ace of Hearts to come before the Two")
    }
    if order["Ace of Hearts"] > order["Ace of Spades"] {
        t.Fatalf("Called deck.SortBy(aceLow) expecting the Ace of Hearts to come before the Ace of Spades")
    }
}

func TestDeckShuffle (t *testing.T) {
    standard := New()
    shuffled := New()

    shuffled.Shuffle(0)

    var diff bool
    for i := range standard {
        if standard[i] != shuffled[i] {
            diff = true
        }
    }
    if !diff {
        t.Fatalf("Called deck.Shuffle expecting the order to change, but the order did not change.")
    }
}

func TestDeckOmit(t *testing.T) {
    deck := New()
    isAce := func(i int) bool {
        return deck[i].Rank == Ace 
    }
    deck.Omit(isAce)
    for _, card := range deck {
        if card.Rank == Ace {
            t.Fatalf("Called deck.Omit(isAce), but one or more Aces remained in the deck")
        }
    }
}

func TestDeckMultiply(t *testing.T) {
    deck := New()
    deck.Multiply(0)
    if len(deck) != 0 {
        t.Fatalf("Multiplied the deck by 0, and received %d cards", len(deck))
    }

    deck = New()
    deck.Multiply(1)
    if len(deck) != 52 {
        t.Fatalf("Multiplied the deck by 1, and received %d cards", len(deck))
    }

    deck = New()
    deck.Multiply(2)
    if len(deck) != 104 {
        t.Fatalf("Multiplied the deck by 2 but received %d cards", len(deck))
    }
    deck.Multiply(2)
    if len(deck) != 52*2*2 {
        t.Fatalf("Multiplied the deck again by 2 but received %d cards. Expected 52 * 2 * 2", len(deck))
    }
}

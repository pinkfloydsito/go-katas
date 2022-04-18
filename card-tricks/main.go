package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var userCards []int
var deck []int = GetDeck()
var playing = true
var currentTurn  = 0

func main() {
	Start()
}

func Start(){
	var choice string
	for playing {

		currentTurn += 1		
		ClearScreen()

		fmt.Printf("___________________________________________\n\n")
		fmt.Println("Hand: ", userCards)
		fmt.Println("Deck: ", deck)
		fmt.Println("Turn ", currentTurn)
		fmt.Println("Choose your action")

		MenuActions()

		fmt.Scanln(&choice)
		
		switch choice {
		case "1":
			FavoriteCards()
		case "2":
			GetCardFromDeck()
		case "3":
			ExChangeCardFromDeck()
		case "4":
			AddCardsOnTopOfHand()
		case "5":
			RemoveCardFromHand()
		case "6":
			playing = false
		default:
			fmt.Println("Unknown choice")
		}
	}
}

func MenuActions() {
	fmt.Println("1. Favorite Cards")
	fmt.Println("2. Retrieve Card from Stack")
	fmt.Println("3. Exchange Card in the stack")
	fmt.Println("4. Add Card to the top of the stack")
	fmt.Println("5. Remove Card")
	fmt.Println("6. Exit")
}


func FavoriteCards() {
	favoriteCards := []int { 2, 6, 9 }

	for _, value := range favoriteCards {
		// favoriteCards[idx]
		idx, err := FindCardIndex(deck, value)
		if err == nil { // found then remove
			deck = RemoveCardIndex(deck, idx)
		}
		
		if !HasCard(userCards, value) { // not found then append
			userCards = append(userCards, value)
		}
	}

	fmt.Println("Favorite Cards added to your current hand")
}

func GetDeck() []int {
	deck := make([]int, 10)

	for idx := range deck {
		deck[idx] =  idx
	}

	return deck
}

func FindCardIndex(cards []int, card int) (int, error) {
	for idx, n := range  cards {
		if n == card {
			return idx, nil
		}
	}

	return -1, errors.New("index not found")
}

func RemoveCardIndex(cards []int, idx int) []int {
	return append(cards[:idx], cards[idx+1:]...)
}

func HasCard(cards []int, card int) bool {
	for _, n := range  cards {
		if n == card {
			return true
		}
	}

	return false
}

/* we get the card from the deck and put it in our current hand */
func GetCardFromDeck() {
	value := GetValueFromStdin("Choose a card to add to your current hand")
	
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("error, ", err)	
		return
	}

	idx, _ := FindCardIndex(deck, parsedValue)	
	
	if idx == -1 {
		fmt.Println("Card not found")
		return	
	}

	userCards = append(userCards, parsedValue)
	deck = RemoveCardIndex(deck, idx)

	fmt.Printf("\nAdded %d to the hand", parsedValue)
}

func ExChangeCardFromDeck() (error) {
	strHandCard := GetValueFromStdin("Choose a card from your hand to change")
	handCard, err := strconv.Atoi(strHandCard)

	if err != nil {
		return err
	}

	if !HasCard(userCards, handCard) {
		return errors.New("you must select a card from the hand")
	}

	strDeckCard := GetValueFromStdin("Choose a card from the deck to swap")
	deckCard, err := strconv.Atoi(strDeckCard)

	if err != nil {
		return err
	}

	if !HasCard(deck, deckCard) {
		return errors.New("you must select a card from the deck")
	}

	deckIdx, _ := FindCardIndex(deck, deckCard)
	handIdx, _ := FindCardIndex(userCards, handCard)

	deck[deckIdx], userCards[handIdx] = userCards[handIdx], deck[deckIdx]

	return nil
}

func AddCardsOnTopOfHand() (error) {
	strCards := GetValueFromStdin("Choose cards from the deck to add to the top of the hand. Format 1,2,3,..")

	strValues := strings.Split(strCards, ",")

	for i:= len(strValues) - 1; i >= 0; i-- {
		value := strValues[i]
		intValue, err := strconv.Atoi(value)
		if err != nil {
			continue
		}

		deckIdx, err := FindCardIndex(deck, intValue)
		if err != nil {
			continue
		}

		deck = RemoveCardIndex(deck, deckIdx)

		handIdx, err := FindCardIndex(userCards, intValue)
		if err == nil {
			fmt.Println("Index already used ", handIdx)
			continue
		}

		userCards = append([]int {intValue}, userCards...)
	}

	return nil
}

func RemoveCardFromHand() (error) {
	strCard := GetValueFromStdin("Card to be removed from hand: ")
	value, err := strconv.Atoi(strCard)
	
	if	err != nil {
		return err
	}

	cardIdx, err := FindCardIndex(userCards, value)

	if	err != nil {
		return err
	}

	deck = append(deck, value)
	userCards = RemoveCardIndex(userCards, cardIdx)

	return nil
}

func GetValueFromStdin(askMessage string) string {
	var value string
	fmt.Println(askMessage)
	fmt.Scanln(&value)

	return value
}

func ClearScreen() {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
}
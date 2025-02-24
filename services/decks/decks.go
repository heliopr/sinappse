package decks

import (
	"encoding/json"
	"sinappsebackend/app"
)

type Deck struct {
	Id uint32 `db:"id" json:"id"`
	UserId uint32 `db:"user_id" json:"-"`
	Name string `db:"name" json:"name"`
	Cards []map[string]any `json:"cards"`
}

func GetFromUserId(userId uint32) (*[]Deck, error) {
	var decks []Deck
	err := app.DB.Select(&decks, "SELECT id,user_id,name FROM decks WHERE user_id=$1", userId)
	if err != nil {
		return nil, err
	}
	return &decks, nil
}

func DeckExistsFromUserByName(name string, userId uint32) (bool, error) {
	var count int
	err := app.DB.QueryRow("SELECT COUNT(*) FROM decks WHERE name=$1 AND user_id=$2 LIMIT 1", name, userId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count>0, nil
}

func IsDeckOwnedByUser(deckId uint32, userId uint32) (bool, error) {
	var count int
	err := app.DB.QueryRow("SELECT COUNT(*) FROM decks WHERE id=$1 AND user_id=$2 LIMIT 1", deckId, userId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count>0, nil
}

func GetDeck(deckId uint32) (*Deck, error) {
	var deck Deck
	err := app.DB.Get(&deck, "SELECT id,user_id,name FROM decks WHERE id=$1 LIMIT 1", deckId)
	if err != nil {
		return nil, err
	}

	var cardsStr string
	err = app.DB.QueryRow("SELECT cards FROM decks WHERE id=$1 LIMIT 1", deckId).Scan(&cardsStr)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(cardsStr), &deck.Cards)
	if err != nil {
		return nil, err
	}

	return &deck, nil
}

func UpdateCards(deckId uint32, cards *[]map[string]any) error {
	m, err := json.Marshal(cards)
	if err != nil {
		return err
	}
	_, err = app.DB.Exec("UPDATE decks SET cards=$1 WHERE id=$2", string(m), deckId)
	return err
}

func CreateDeck(name string, userId uint32) error {
	_, err := app.DB.Exec("INSERT INTO decks (user_id, name) VALUES ($1, $2)", userId, name)
	return err
}
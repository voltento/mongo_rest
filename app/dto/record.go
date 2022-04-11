package dto

type Record struct {
	Text   string  `bson:"text"`
	Number float64 `bson:",string"`
	Found  bool    `bson:",string"`
	Type   string  `bson:"type"`
}

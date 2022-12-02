package main

type Card struct {
	seqNum     int    // Порядковый номер в колоде для тасовки
	precedence int    // 6, 7, 8, 9, 10, 11, 12, 13, 14
	name       string // 6, 7, 8, 9, 10, Валет, Дама, Король, Туз
	suit       string // Масть: Буби, Крести, Черви, Пики
	suitPic    int32  // Картинка масти
	inGame     bool   // В игре или бита
	trump			 bool		// козырь-не козырь
}


// Включить в состав Card, переформатировав саму Card
// Все карты масти
type suitDeck struct {
	suitName string // название масти из массива suits
	suitPic int32		// из карты suitPics
	precedence int 	// старшинство (6, 7, 10, ...14)
	name string			// из карты names
	cardCount	int		// количество карт в масти

}

var precedences [9]int = [9]int{6, 7, 8, 9, 10, 11, 12, 13, 14}

var names map[int]string = map[int]string{
	6:  "6",  // six
	7:  "7",  // seven
	8:  "8",  // eigth
	9:  "9",  // nine
	10: "10", // ten
	11: "J",  // Jack
	12: "Q",  // Queen
	13: "K",  // King
	14: "A",  // Ace
}

var suits [4]string = [4]string{"spades", "hearts", "crosses", "diamonds"}

var suitPics map[string]int32 = map[string]int32{
	"spades":   '♤', // пики
	"hearts":   '♡', // черви
	"crosses":  '♧', // крести
	"diamonds": '♢', // буби
}

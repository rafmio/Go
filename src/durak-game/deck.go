package deck

type Card struct {
  SeqNum  int           // Порядковый номер в колоде для тасовки
  Priority int          // 6, 7, 8, 9, 10, 11, 12, 13, 14
  Name string           // 6, 7, 8, 9, 10, Valet, Dama, Korol, Tuz
  Suit string           // Bubi, Cresti, Chervi, Piki
  InGame bool           // В игре или бита
}

type Deck struct {
  GameCard Card
}

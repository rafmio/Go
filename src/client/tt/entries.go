package main

type Order struct {
  ID int
  LoadingDate string
  Shipper string
  Index int // "3ая" column
  CodeATI int
  DepatureCity string
  DestinationCity string
  Driver string
  Dispatcher string
  DispatcherPhoneNum string
  Price float64
  Note string
  PricePerKm string
  Truck string
  PaymentTerm string // "Дата оплаты"
  Payment string
}

var Entries = []Order {
  {1, "14.12.2022", "РБ Групп", 175,  155057, "Казань", "Новокузнецк", "Мингазеев",  "Артем", "89236-86-46-46", 0,      "НДС", "заявку", "Скания", "", ""},
  {2, "13.12.2022", "РБ Групп", 174,  155057,	"Казань",	"Сургут",      "Алиякберов", "Артем", "89236-86-46-46",	215000,	"НДС", "",       "МАН 2",  "", ""},
  {3, "01.12.2022", "РБ Групп", 162,	155057,	"Казань", "Сургут",      "Нурулин",    "Артем",	"89236-86-46-46",	215000, "НДС", "отпр",	 "ДАФ 2",  "", ""},
}

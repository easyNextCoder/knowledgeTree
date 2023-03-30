package rules

import (
	"fmt"
	"strconv"

	"sort"
)

const (
	ColorMask = 0xf0
	NumMask   = 0x0f
)
const (
	ColorNone   = 0x00 // 没有颜色，仅假OKEY牌能有这种颜色
	ColorBlue   = 0x10 //黑桃
	ColorRed    = 0x20 //红桃
	ColorBlack  = 0x30 //方块
	ColorYellow = 0x40 //梅花
)

var colorString = map[int]string{
	ColorBlue:   "\u2660",
	ColorYellow: "\u2666",
	ColorBlack:  "\u2663",
	ColorRed:    "\u2665",
}

const (
	CardNumA  = 1
	CardNum2  = 2
	CardNum3  = 3
	CardNum4  = 4
	CardNum5  = 5
	CardNum6  = 6
	CardNum7  = 7
	CardNum8  = 8
	CardNum9  = 9
	CardNum10 = 10
	CardNumJ  = 11
	CardNumQ  = 12
	CardNumK  = 13
)

const (
	CardNumStart = 1
	CardNumEnd   = 13
	// 用14代表OKEY牌，没有颜色
	CardOkey     = 14
	CardOkey2    = 15
	CardBigStart = 9
	CardSmallEnd = 7
)
const (
	// 玩家卡数量
	PlayerCardNum = 14
)

const OkeyCard = Card(ColorNone | CardOkey)
const OkeyCard2 = Card(ColorNone | CardOkey2)
const InvalidCard Card = 0 // 无效牌
const InvliadCard Card = 0 // 无效牌
var AllColorList = []int{ColorBlue, ColorYellow, ColorBlack, ColorRed}
var handScoreArr []int = []int{0, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
var gameOverHandCardScoreArr []int = []int{0, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

const GameOverJokerInHandScore int = 15

func IsOkCard(c Card) bool {
	return c == OkeyCard || c == OkeyCard2
}

type Card int
type GameMode int

func NewCard(color, num int) Card {
	return Card(color | num)
}

func NewCard16(color, num int16) Card {
	return Card(color | num)
}

func (c Card) Color() int {
	return int(c) & ColorMask
}

func (c Card) Num() int {
	return int(c) & NumMask
}

func (c Card) Next() (Card, bool) {
	num := c.Num()
	if num == CardNumEnd {
		return 0, false
	}
	return NewCard(c.Color(), num+1), true
}

// 上一张牌，用于抽取指示牌
func (c Card) Prev() (Card, bool) {
	color := c.Color()
	num := c.Num()
	if num == CardNumStart {
		return 0, false
	}
	return NewCard(color, num-1), true
}

func (c Card) Equal(card Card) bool {
	return c.Color() == card.Color() && c.Num() == card.Num()
}

func (c Card) Score() int {
	if c == OkeyCard || c == OkeyCard2 {
		return 101
	}
	if c.Num() > 0 && c.Num() < 14 {
		return handScoreArr[c.Num()]
	}
	return 0
}

// 手牌留在手里的分数
func (c Card) HandScore() int {
	if c == OkeyCard || c == OkeyCard2 {
		return 101
	}
	if c.Num() > 0 && c.Num() < 14 {
		return handScoreArr[c.Num()]
	}
	return 0
}

func (c Card) Valid() bool {
	return defaultCardMap[c]
}

func GetCardStr(suit int, week int) string {
	var out string
	switch suit {
	case ColorYellow:
		out = "\u2666"
	case ColorBlack:
		out = "\u2663"
	case ColorRed:
		out = "\u2665"
	case ColorBlue:
		out = "\u2660"
	}
	switch week {
	case 1:
		out += "A"
	case 11:
		out += "J"
	case 12:
		out += "Q"
	case 13:
		out += "K"
	default:
		out += strconv.Itoa(week)
	}
	return out
}

func (c Card) String() string {
	if c.Valid() {
		if c.Num() == 14 {
			return "Joker"
		} else if c.Num() == 15 {
			return "Joker2"
		} else {
			return GetCardStr(c.Color(), c.Num())
		}
	} else {
		return fmt.Sprintf("Invalid_%d", int(c))
	}
}

type CardsProject [][]Card

func (c CardsProject) Len() int      { return len(c) }
func (c CardsProject) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CardsProject) Less(i, j int) bool {
	return c[i][0] < c[j][0]
}

func (c CardsProject) ToCards() []Card {
	var res []Card
	for _, card := range c {
		res = append(res, card...)
	}
	return res
}

func SortCardsProject(list [][]Card) {
	sort.Sort(CardsProject(list))
}

func NewCardsWithoutShuffle() []Card {
	list := make([]Card, len(defaultCardList))
	copy(list, defaultCardList)
	return list
}

var defaultCardList []Card
var defaultCardMap = map[Card]bool{}
var FullCardMap = map[Card]int{}

func initFullCardMap() {
	for i := 0; i < 2; i++ {
		for _, color := range AllColorList {
			for num := CardNumStart; num <= CardNumEnd; num++ {
				card := NewCard(color, num)
				FullCardMap[card]++
			}
		}
	}
	FullCardMap[OkeyCard]++
	FullCardMap[OkeyCard2]++
}

func initDefaultCards() {
	var cardList = make([]Card, 0)
	var cardMap = make(map[Card]bool)
	cardList = append(cardList, OkeyCard, OkeyCard2)
	for _, color := range AllColorList {
		for num := CardNumStart; num <= CardNumEnd; num++ {
			card := NewCard(color, num)
			cardList = append(cardList, card, card)
			cardMap[card] = true
		}
	}
	cardMap[OkeyCard] = true
	cardMap[OkeyCard2] = true
	defaultCardList = cardList
	defaultCardMap = cardMap
}

func CopyCards(c []Card) []Card {
	list := make([]Card, len(c))
	copy(list, c)
	return list
}

func ToCardMap(cardList []Card) map[Card]int {
	cMap := make(map[Card]int)
	for _, c := range cardList {
		cMap[c] += 1
	}
	return cMap
}

func CardsEquals(a []Card, b []Card) bool {
	if len(a) != len(b) {
		return false
	}

	mapa := ToCardMap(a)
	mapb := ToCardMap(b)

	if len(mapa) != len(mapb) {
		return false
	}

	for k, v := range mapa {
		lv, ok := mapb[k]
		if !ok {
			return false
		}
		if lv != v {
			return false
		}
	}

	return true
}

func ToCardListWithOutSort(cardMap map[Card]int) []Card {
	cards := make([]Card, 0, len(cardMap))
	for c, count := range cardMap {
		for i := 0; i < count; i++ {
			cards = append(cards, c)
		}
	}

	return cards
}

func ToCardList(cardMap map[Card]int) []Card {
	cards := make([]Card, 0, len(cardMap))
	for c, count := range cardMap {
		for i := 0; i < count; i++ {
			cards = append(cards, c)
		}
	}
	sort.Sort(SortByColorNum(cards))
	return cards
}

func CardCount(cards map[Card]int) int {
	total := 0
	for _, count := range cards {
		if count > 0 {
			total += count
		}
	}
	return total
}

func IsOneCard(theCard Card, cards []Card) bool {
	for _, one := range cards {
		if one == theCard {
			return true
		}
	}
	return false
}

func init() {
	initDefaultCards()
	initFullCardMap()
}

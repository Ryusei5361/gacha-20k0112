package gacha

import (
	"math/rand"
	"time"
)

func init() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())
}

func DrawN(p *Player, n int, c *Character) ([]*Card, map[Rarity]int) {
	p.draw(n)

	results := make([]*Card, n)
	summary := make(map[Rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw(c)
		summary[results[i].Rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw(c *Character) *Card {
	rarity := drawrare()
	name := drawchara(c)
	return &Card{Rarity: rarity, Name: name}
}

func drawrare() Rarity {
	num := rand.Intn(100)
	switch {
	case num < 80:
		return RarityN
	case num < 95:
		return RarityR
	case num < 99:
		return RaritySR
	default:
		return RarityXR
	}
}

func drawchara(c *Character) string {
	chara := drawrare()
	switch chara {
	case RarityN:
		n := len(c.CharacterN)
		num := rand.Intn(n)
		return c.CharacterN[num]
	case RarityR:
		n := len(c.CharacterR)
		num := rand.Intn(n)
		return c.CharacterR[num]
	case RaritySR:
		n := len(c.CharacterSR)
		num := rand.Intn(n)
		return c.CharacterSR[num]
	default:
		n := len(c.CharacterXR)
		num := rand.Intn(n)
		return c.CharacterXR[num]
	}
}

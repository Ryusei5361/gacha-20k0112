package main

import (
	"fmt"

	"github.com/Ryusei5361/gacha-20k0112/gacha"
)

var chara = &gacha.Character{
	CharacterN:  []string{"スライム", "ゴブリン", "コボルト"},
	CharacterR:  []string{"オーク", "オーク2", "オーク3"},
	CharacterSR: []string{"ドラゴン", "フェンリル", "巨人"},
	CharacterXR: []string{"イフリート", "ゼウス", "アテナ"},
}

func main() {
	p := gacha.NewPlayer(10, 100)
	n := inputN(p)
	results, summary := gacha.DrawN(p, n, chara)

	for _, v := range results {
		fmt.Println(v)
	}
	fmt.Println(summary)
}

func inputN(p *gacha.Player) int {

	max := p.DrawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}

	return n
}

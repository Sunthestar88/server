//go:build !prod || full || novomatic

package dolphinspearl

import (
	"github.com/slotopol/server/game"
	"github.com/slotopol/server/util"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{Prov: "Novomatic", Name: "Dolphins Pearl"},
		{Prov: "Novomatic", Name: "Dolphins Pearl Deluxe"},
		{Prov: "Aristocrat", Name: "Dolphin Treasure"}, // See: https://freeslotshub.com/aristocrat/dolphin-treasure/
		{Prov: "Novomatic", Name: "Attila"},
		{Prov: "Novomatic", Name: "Banana Splash"},
		{Prov: "Novomatic", Name: "Dynasty Of Ming"},
		{Prov: "Novomatic", Name: "Gryphons Gold"},
		{Prov: "Novomatic", Name: "Joker Dolphin"},
		{Prov: "Novomatic", Name: "King Of Cards"},
		{Prov: "Novomatic", Name: "Lucky Lady's Charm"},
		{Prov: "Novomatic", Name: "Lucky Lady's Charm Deluxe"},
		{Prov: "Novomatic", Name: "Pharaon's Gold II"},
		{Prov: "Novomatic", Name: "Pharaon's Gold III"},
		{Prov: "Novomatic", Name: "Polar Fox"},
		{Prov: "Novomatic", Name: "Ramses II"},
		{Prov: "Novomatic", Name: "Royal Treasures"},
		{Prov: "Novomatic", Name: "Secret Forest"},
		{Prov: "Novomatic", Name: "The Money Game"},
		{Prov: "Novomatic", Name: "Unicorn Magic"},
	},
	GP: game.GPsel |
		game.GPretrig |
		game.GPfgmult |
		game.GPfgreel |
		game.GPscat |
		game.GPwild,
	SX:  5,
	SY:  3,
	SN:  len(LinePay),
	LN:  len(BetLines),
	BN:  0,
	RTP: game.MakeRtpList(ReelsMap),
}

func init() {
	game.GameList = append(game.GameList, &Info)
	for _, ga := range Info.Aliases {
		var aid = util.ToID(ga.Prov + "/" + ga.Name)
		game.ScanFactory[aid] = CalcStatReg
		game.GameFactory[aid] = func() any { return NewGame() }
	}
}

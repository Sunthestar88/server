//go:build !prod || full || megajack

package slotopol

import (
	"github.com/slotopol/server/game"
	"github.com/slotopol/server/util"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{Prov: "Megajack", Name: "Slotopol"},
	},
	GP: game.GPsel |
		game.GPfgno |
		game.GPscat |
		game.GPwild,
	SX:  5,
	SY:  3,
	SN:  len(LinePay),
	LN:  len(BetLines),
	BN:  2,
	RTP: game.MakeRtpList(ReelsMap),
}

func init() {
	game.GameList = append(game.GameList, &Info)
	for _, ga := range Info.Aliases {
		var aid = util.ToID(ga.Prov + "/" + ga.Name)
		game.ScanFactory[aid] = CalcStat
		game.GameFactory[aid] = func() any { return NewGame() }
	}
}

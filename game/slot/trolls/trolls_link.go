//go:build !prod || full || netent

package trolls

import (
	"context"

	game "github.com/slotopol/server/game"
	"github.com/spf13/pflag"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{ID: "trolls", Name: "Trolls"},
		{ID: "excalibur", Name: "Excalibur"},
		{ID: "pandorasbox", Name: "Pandora's Box"},
		{ID: "wildwitches", Name: "Wild Witches"},
	},
	Provider: "NetEnt",
	GP: game.GPsel |
		game.GPretrig |
		game.GPfgmult |
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
		game.ScanIters = append(game.ScanIters, func(flags *pflag.FlagSet, ctx context.Context) {
			if is, _ := flags.GetBool(ga.ID); is {
				var mrtp, _ = flags.GetFloat64("reels")
				CalcStat(ctx, mrtp)
			}
		})
		game.GameFactory[ga.ID] = func() any {
			return NewGame()
		}
	}
}
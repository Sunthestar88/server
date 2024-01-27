//go:build !prod || full || novomatic

package config

import (
	"context"

	"github.com/slotopol/server/game/katana"
	"github.com/spf13/pflag"
)

func init() {
	FlagsSetters = append(FlagsSetters, func(flags *pflag.FlagSet) {
		flags.Bool("katana", false, "'Katana' Novomatic 5x3 slots")
	})
	ScatIters = append(ScatIters, func(flags *pflag.FlagSet, ctx context.Context) {
		if is, _ := flags.GetBool("katana"); is {
			var rn, _ = flags.GetString("reels")
			if rn == "bon" {
				katana.CalcStatBon(ctx)
			} else {
				katana.CalcStatReg(ctx, rn)
			}
		}
	})

	for _, alias := range []string{
		"katana",
	} {
		GameAliases[alias] = "katana"
	}

	GameFactory["katana"] = func(name string) any {
		if _, ok := katana.ReelsMap[name]; ok {
			return katana.NewGame(name)
		}
		return nil
	}
}
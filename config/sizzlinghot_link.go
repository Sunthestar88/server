//go:build !prod || full || novomatic

package config

import (
	"context"

	"github.com/slotopol/server/game/sizzlinghot"
	"github.com/spf13/pflag"
)

func init() {
	FlagsSetters = append(FlagsSetters, func(flags *pflag.FlagSet) {
		flags.Bool("sizzlinghot", false, "'Sizzling Hot' Novomatic 5x3 slots")
	})
	ScatIters = append(ScatIters, func(flags *pflag.FlagSet, ctx context.Context) {
		if is, _ := flags.GetBool("sizzlinghot"); is {
			var rn, _ = flags.GetString("reels")
			sizzlinghot.CalcStat(ctx, rn)
		}
	})

	for _, alias := range []string{
		"sizzlinghot",
		"sizzlinghotdeluxe",
	} {
		GameAliases[alias] = "sizzlinghot"
	}

	GameFactory["sizzlinghot"] = func(name string) any {
		if _, ok := sizzlinghot.ReelsMap[name]; ok {
			return sizzlinghot.NewGame(name)
		}
		return nil
	}
}
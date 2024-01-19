package dolphinspearl

import (
	"context"
	"fmt"
	"time"

	"github.com/slotopol/server/game"
)

func CalcStatBon(ctx context.Context) float64 {
	var reels = &ReelsBon
	var g = NewGame("bon")
	var sbl = float64(g.SBL.Num())
	var s game.Stat

	var dur = func() time.Duration {
		var t0 = time.Now()
		var ctx2, cancel2 = context.WithCancel(ctx)
		defer cancel2()
		go s.Progress(ctx2, time.NewTicker(2*time.Second), sbl, float64(reels.Reshuffles()))
		s.BruteForce5x(ctx2, g, reels)
		return time.Since(t0)
	}()

	var n = float64(s.Reshuffles)
	var lrtp, srtp = float64(s.LinePay) / n / sbl * 100, float64(s.ScatPay) / n * 100
	var rtpsym = lrtp + srtp
	var q = float64(s.FreeCount) / n
	var sq = 1 / (1 - q)
	var rtp = sq * rtpsym * 3
	fmt.Printf("completed %.5g%%, selected %d lines, time spent %v\n", float64(s.Reshuffles)/float64(reels.Reshuffles())*100, g.SBL.Num(), dur)
	fmt.Printf("reels lengths [%d, %d, %d, %d, %d], total reshuffles %d\n",
		len(reels.Reel(1)), len(reels.Reel(2)), len(reels.Reel(3)), len(reels.Reel(4)), len(reels.Reel(5)), reels.Reshuffles())
	fmt.Printf("symbols: %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
	fmt.Printf("free games %d, q = %.5g, sq = 1/(1-q) = %.6f\n", s.FreeCount, q, sq)
	fmt.Printf("RTP = sq*rtp(sym)*3 = %.5g*%.5g*3 = %.6f%%\n", sq, rtpsym, rtp)
	return rtp
}

func CalcStatReg(ctx context.Context, rn string) float64 {
	fmt.Printf("*bonus reels calculations*\n")
	var rtpfs = CalcStatBon(ctx)
	fmt.Printf("*regular reels calculations*\n")
	var reels *game.Reels5x
	if rn != "" {
		var ok bool
		if reels, ok = ReelsMap[rn]; !ok {
			return 0
		}
	} else {
		rn, reels = "92", &ReelsReg92
	}
	var g = NewGame(rn)
	var sbl = float64(g.SBL.Num())
	var s game.Stat

	var dur = func() time.Duration {
		var t0 = time.Now()
		var ctx2, cancel2 = context.WithCancel(ctx)
		defer cancel2()
		go s.Progress(ctx2, time.NewTicker(2*time.Second), sbl, float64(reels.Reshuffles()))
		s.BruteForce5x(ctx2, g, reels)
		return time.Since(t0)
	}()

	var n = float64(s.Reshuffles)
	var lrtp, srtp = float64(s.LinePay) / n / sbl * 100, float64(s.ScatPay) / n * 100
	var rtpsym = lrtp + srtp
	var q = float64(s.FreeCount) / n
	var sq = 1 / (1 - q)
	var rtp = lrtp + srtp + q*rtpfs
	fmt.Printf("completed %.5g%%, selected %d lines, time spent %v\n", float64(s.Reshuffles)/float64(reels.Reshuffles())*100, g.SBL.Num(), dur)
	fmt.Printf("reels lengths [%d, %d, %d, %d, %d], total reshuffles %d\n",
		len(reels.Reel(1)), len(reels.Reel(2)), len(reels.Reel(3)), len(reels.Reel(4)), len(reels.Reel(5)), reels.Reshuffles())
	fmt.Printf("symbols: %.5g(lined) + %.5g(scatter) = %.6f%%\n", lrtp, srtp, rtpsym)
	fmt.Printf("free games %d, q = %.5g, sq = 1/(1-q) = %.6f\n", s.FreeCount, q, sq)
	fmt.Printf("RTP = %.5g(sym) + %.5g*%.5g(fg) = %.6f%%\n", rtpsym, sq, rtpfs, rtp)
	return rtp
}
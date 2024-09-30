package panthermoon

import (
	slot "github.com/slotopol/server/game/slot"
)

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 39.772(lined) + 11.288(scatter) = 51.059855%
// free spins 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 51.06(sym) + 0.057217*618.88(fg) = 86.470777%
var ReelsReg86 = slot.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 11, 4, 7, 9, 13, 10, 9, 3, 7, 5, 11, 9, 6, 11, 9, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 10, 8, 3, 10, 9, 2, 8, 10, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 5, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 7, 9, 2, 8, 6, 10, 4, 11, 12, 2, 7, 8, 9, 12, 13, 9, 11, 5, 11, 3, 10, 9, 4, 7, 5, 9, 11, 2, 7, 5, 12, 7, 3, 11, 9, 5, 7, 3},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 11, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 9, 6, 12, 5, 8, 7, 6, 10, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 12, 10, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 12, 6, 7, 2, 10, 8, 2, 12, 10, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 41.45(lined) + 11.288(scatter) = 52.738002%
// free spins 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 52.738(sym) + 0.057217*618.88(fg) = 88.148924%
var ReelsReg88 = slot.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 11, 4, 7, 9, 13, 10, 9, 3, 7, 5, 11, 9, 6, 11, 9, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 10, 8, 3, 10, 9, 2, 8, 10, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 11, 12, 4, 8, 10, 5, 7, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 7, 9, 2, 8, 6, 10, 4, 11, 12, 2, 7, 8, 9, 12, 13, 9, 11, 5, 11, 3, 10, 9, 4, 7, 5, 9, 11, 2, 7, 5, 12, 7, 3, 11, 10, 5, 7, 3},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 11, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 9, 6, 12, 5, 8, 7, 6, 10, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 12, 10, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 12, 6, 7, 2, 10, 8, 2, 12, 10, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 43.667(lined) + 11.288(scatter) = 54.954656%
// free spins 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 54.955(sym) + 0.057217*618.88(fg) = 90.365578%
var ReelsReg90 = slot.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 5, 8, 11, 4, 7, 9, 13, 10, 9, 3, 7, 5, 10, 8, 6, 11, 9, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 10, 8, 3, 10, 9, 2, 8, 10, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 7, 10, 6, 8, 7, 5, 12, 4, 9, 11},
	{11, 9, 1, 11, 4, 7, 9, 2, 8, 6, 10, 4, 10, 12, 2, 7, 8, 9, 12, 13, 9, 11, 5, 11, 3, 10, 9, 4, 7, 5, 10, 11, 2, 7, 6, 12, 8, 3, 11, 9, 6, 7, 3},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 11, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 9, 6, 12, 5, 8, 7, 6, 10, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 12, 10, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 12, 6, 7, 2, 10, 8, 2, 12, 10, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 45.775(lined) + 11.288(scatter) = 57.062817%
// free spins 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 57.063(sym) + 0.057217*618.88(fg) = 92.473739%
var ReelsReg92 = slot.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 7, 5, 11, 9, 10, 11, 9, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 8, 10, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 6, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 8, 9, 2, 8, 6, 10, 4, 11, 12, 2, 10, 8, 9, 12, 13, 9, 11, 5, 11, 3, 10, 9, 4, 7, 5, 9, 11, 2, 7, 6, 12, 7, 3, 11, 9, 6, 7, 3},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 8, 11, 1, 9, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 12, 10, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 12, 6, 7, 2, 10, 8, 2, 12, 10, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 47.537(lined) + 11.288(scatter) = 58.825038%
// free spins 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 58.825(sym) + 0.057217*618.88(fg) = 94.235960%
var ReelsReg94 = slot.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 7, 5, 11, 8, 10, 11, 5, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 8, 10, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 6, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 8, 9, 2, 8, 6, 10, 4, 7, 12, 2, 10, 8, 9, 12, 13, 9, 11, 5, 11, 3, 10, 9, 4, 7, 5, 9, 11, 2, 7, 6, 12, 7, 3, 11, 10, 6, 7, 3},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 8, 11, 1, 9, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 12, 10, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 12, 6, 7, 2, 10, 8, 2, 12, 10, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 48.704(lined) + 11.288(scatter) = 59.992083%
// free spins 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 59.992(sym) + 0.057217*618.88(fg) = 95.403005%
var ReelsReg95 = slot.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 10, 5, 11, 8, 10, 11, 5, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 7, 9, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 6, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 7, 2, 12, 6, 4, 10, 3, 12, 5, 11, 10, 2, 9, 12, 13, 9, 11, 3, 7, 5, 9, 6, 4, 8, 2, 10, 5, 7, 3, 12, 4, 9, 6, 7, 8, 5, 7, 6},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 8, 11, 1, 9, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 7, 9, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 9, 6, 7, 2, 8, 7, 2, 12, 9, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 49.703(lined) + 11.288(scatter) = 60.991139%
// free spins 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 60.991(sym) + 0.057217*618.88(fg) = 96.402061%
var ReelsReg96 = slot.Reels5x{
	{6, 11, 1, 9, 12, 8, 11, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 10, 5, 11, 8, 10, 11, 5, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 7, 9, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 6, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 7, 2, 8, 6, 4, 10, 3, 12, 5, 11, 10, 2, 9, 12, 13, 9, 11, 3, 7, 5, 9, 6, 4, 8, 2, 10, 5, 7, 3, 12, 4, 9, 6, 7, 8, 5, 7, 6},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 6, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 8, 11, 1, 9, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 7, 9, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 9, 6, 7, 2, 8, 7, 2, 12, 9, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 43, 43, 43, 43], total reshuffles 102564030
// symbols: 50.702(lined) + 11.288(scatter) = 61.990173%
// free spins 5868450, q = 0.057217, sq = 1/(1-q) = 1.060690
// free games frequency: 1/262.16
// RTP = 61.99(sym) + 0.057217*618.88(fg) = 97.401095%
var ReelsReg97 = slot.Reels5x{
	{6, 11, 1, 9, 12, 8, 6, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 10, 5, 11, 8, 10, 11, 5, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 7, 9, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 3, 10, 12, 4, 8, 10, 6, 8, 10, 5, 12, 4, 9, 7},
	{11, 9, 1, 11, 4, 7, 2, 8, 6, 4, 10, 3, 12, 5, 11, 10, 2, 9, 12, 13, 9, 11, 3, 7, 5, 9, 6, 4, 8, 2, 10, 5, 7, 3, 12, 4, 9, 6, 7, 8, 5, 7, 6},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 9, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 8, 11, 1, 9, 3, 6, 12, 4, 10, 5, 12, 10, 3},
	{8, 7, 1, 8, 4, 9, 5, 7, 9, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 9, 6, 7, 2, 8, 7, 2, 12, 9, 3, 8, 10, 4, 12, 5, 8, 12, 2},
}

// reels lengths [30, 32, 32, 32, 32], total reshuffles 31457280
// symbols: 54.789(lined) + 18.01(scatter) = 72.799053%
// free spins 3489480, q = 0.11093, sq = 1/(1-q) = 1.124768
// free games frequency: 1/135.22
// RTP = 72.799(sym) + 0.11093*618.88(fg) = 141.450303%
var ReelsReg141 = slot.Reels5x{
	{6, 11, 1, 9, 12, 8, 6, 7, 2, 11, 6, 8, 9, 4, 7, 9, 13, 10, 9, 3, 10, 5, 11, 8, 10, 11, 5, 10, 4, 7},
	{5, 8, 1, 7, 8, 11, 2, 9, 10, 12, 6, 11, 7, 3, 10, 9, 2, 7, 9, 5, 8, 11, 13, 12, 4, 10, 5, 8, 6, 10, 9, 7},
	{11, 9, 1, 11, 4, 7, 2, 8, 6, 4, 10, 3, 12, 5, 11, 10, 2, 9, 12, 13, 9, 11, 3, 7, 5, 9, 6, 4, 8, 2, 7, 4},
	{8, 11, 1, 9, 3, 7, 10, 4, 9, 2, 7, 6, 10, 5, 12, 10, 9, 12, 8, 5, 11, 13, 8, 6, 10, 2, 7, 6, 12, 5, 11, 9},
	{8, 7, 1, 8, 4, 9, 5, 7, 9, 6, 11, 3, 7, 8, 2, 9, 5, 10, 11, 2, 12, 13, 5, 10, 4, 8, 6, 7, 2, 9, 12, 10},
}

// reels lengths [33, 36, 36, 36, 36], total reshuffles 55427328
// symbols: 226.63(lined) + 144.15(scatter) = 370.783105%
// free spins 22219920, q = 0.40088, sq = 1/(1-q) = 1.669125
// free games frequency: 1/37.417
// RTP = sq*rtp(sym) = 1.6691*370.78 = 618.883497%
var ReelsBon = slot.Reels5x{
	{6, 11, 1, 9, 12, 5, 11, 8, 2, 7, 6, 10, 12, 4, 7, 9, 13, 10, 9, 3, 7, 5, 10, 9, 2, 8, 11, 10, 3, 8, 11, 4, 7},
	{5, 8, 1, 7, 8, 10, 2, 9, 4, 11, 5, 8, 1, 7, 8, 11, 13, 12, 4, 10, 7, 3, 11, 6, 10, 4, 6, 11, 13, 12, 5, 9, 6, 11, 9, 10},
	{11, 9, 1, 11, 4, 7, 8, 2, 6, 11, 9, 1, 11, 4, 8, 12, 13, 9, 8, 2, 12, 10, 5, 12, 3, 6, 7, 12, 13, 9, 10, 5, 3, 7, 5, 10},
	{8, 11, 1, 9, 3, 7, 10, 4, 7, 2, 8, 6, 10, 4, 5, 11, 13, 8, 6, 12, 11, 5, 7, 9, 10, 2, 5, 11, 13, 8, 6, 9, 3, 12, 7, 4},
	{8, 7, 1, 8, 4, 9, 5, 12, 7, 6, 11, 3, 7, 8, 2, 12, 13, 5, 10, 7, 8, 9, 6, 10, 4, 9, 2, 12, 13, 5, 10, 4, 11, 6, 3, 11},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	86.470777:  &ReelsReg86,
	88.148924:  &ReelsReg88,
	90.365578:  &ReelsReg90,
	92.473739:  &ReelsReg92,
	94.235960:  &ReelsReg94,
	95.403005:  &ReelsReg95,
	96.402061:  &ReelsReg96,
	97.401095:  &ReelsReg97,
	141.450303: &ReelsReg141,
}
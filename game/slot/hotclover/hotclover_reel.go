package hotclover

import (
	"github.com/slotopol/server/game/slot"
)

// reels lengths [73, 74, 74, 74, 73], total reshuffles 2159438696
// RTP = 79.372(lined) + 4.457(scatter) = 83.829264%
var Reels838 = slot.Reels5x{
	{4, 4, 4, 4, 9, 9, 2, 8, 8, 9, 9, 9, 9, 7, 7, 7, 7, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 10, 10, 3, 3, 3, 3, 6, 6, 6, 6, 7, 7, 5, 5, 5, 5, 10, 10, 10, 10, 9, 9, 9, 9, 5, 5, 5, 5, 4, 4, 4, 2, 6, 6, 8, 8, 8, 8, 7, 7, 7, 7, 8, 8, 8, 8},
	{4, 4, 4, 4, 8, 8, 8, 8, 2, 7, 7, 9, 9, 9, 9, 5, 5, 5, 5, 1, 4, 4, 4, 6, 6, 6, 6, 10, 10, 10, 10, 6, 6, 10, 10, 10, 10, 7, 7, 7, 7, 2, 8, 8, 8, 8, 6, 6, 6, 6, 7, 7, 7, 7, 3, 3, 3, 3, 9, 9, 8, 8, 5, 5, 5, 5, 9, 9, 9, 9, 10, 10, 10, 10},
	{4, 4, 4, 4, 8, 8, 8, 8, 2, 7, 7, 9, 9, 9, 9, 5, 5, 5, 5, 1, 4, 4, 4, 6, 6, 6, 6, 10, 10, 10, 10, 6, 6, 10, 10, 10, 10, 7, 7, 7, 7, 2, 8, 8, 8, 8, 6, 6, 6, 6, 7, 7, 7, 7, 3, 3, 3, 3, 9, 9, 8, 8, 5, 5, 5, 5, 9, 9, 9, 9, 10, 10, 10, 10},
	{4, 4, 4, 4, 8, 8, 8, 8, 2, 7, 7, 9, 9, 9, 9, 5, 5, 5, 5, 1, 4, 4, 4, 6, 6, 6, 6, 10, 10, 10, 10, 6, 6, 10, 10, 10, 10, 7, 7, 7, 7, 2, 8, 8, 8, 8, 6, 6, 6, 6, 7, 7, 7, 7, 3, 3, 3, 3, 9, 9, 8, 8, 5, 5, 5, 5, 9, 9, 9, 9, 10, 10, 10, 10},
	{4, 4, 4, 4, 9, 9, 2, 8, 8, 9, 9, 9, 9, 7, 7, 7, 7, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 10, 10, 3, 3, 3, 3, 6, 6, 6, 6, 7, 7, 5, 5, 5, 5, 10, 10, 10, 10, 9, 9, 9, 9, 5, 5, 5, 5, 4, 4, 4, 2, 6, 6, 8, 8, 8, 8, 7, 7, 7, 7, 8, 8, 8, 8},
}

// reels lengths [73, 74, 74, 74, 73], total reshuffles 2159438696
// RTP = 83.447(lined) + 4.457(scatter) = 87.903734%
var Reels879 = slot.Reels5x{
	{7, 7, 7, 7, 8, 8, 8, 8, 2, 7, 7, 4, 4, 4, 4, 9, 9, 6, 6, 3, 3, 3, 3, 6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 9, 9, 4, 4, 4, 10, 10, 6, 6, 6, 6, 5, 5, 5, 5, 8, 8, 8, 8, 10, 10, 10, 10, 8, 8, 8, 8, 5, 5, 5, 5, 10, 10, 10, 10, 2, 9, 9, 9, 9},
	{5, 5, 5, 5, 8, 8, 8, 8, 1, 9, 9, 7, 7, 10, 10, 10, 10, 8, 8, 8, 8, 2, 6, 6, 6, 6, 5, 5, 5, 5, 9, 9, 9, 9, 3, 3, 3, 3, 4, 4, 4, 6, 6, 7, 7, 7, 7, 6, 6, 6, 6, 10, 10, 4, 4, 4, 4, 10, 10, 10, 10, 7, 7, 7, 7, 2, 9, 9, 9, 9, 8, 8, 8, 8},
	{5, 5, 5, 5, 8, 8, 8, 8, 1, 9, 9, 7, 7, 10, 10, 10, 10, 8, 8, 8, 8, 2, 6, 6, 6, 6, 5, 5, 5, 5, 9, 9, 9, 9, 3, 3, 3, 3, 4, 4, 4, 6, 6, 7, 7, 7, 7, 6, 6, 6, 6, 10, 10, 4, 4, 4, 4, 10, 10, 10, 10, 7, 7, 7, 7, 2, 9, 9, 9, 9, 8, 8, 8, 8},
	{5, 5, 5, 5, 8, 8, 8, 8, 1, 9, 9, 7, 7, 10, 10, 10, 10, 8, 8, 8, 8, 2, 6, 6, 6, 6, 5, 5, 5, 5, 9, 9, 9, 9, 3, 3, 3, 3, 4, 4, 4, 6, 6, 7, 7, 7, 7, 6, 6, 6, 6, 10, 10, 4, 4, 4, 4, 10, 10, 10, 10, 7, 7, 7, 7, 2, 9, 9, 9, 9, 8, 8, 8, 8},
	{7, 7, 7, 7, 8, 8, 8, 8, 2, 7, 7, 4, 4, 4, 4, 9, 9, 6, 6, 3, 3, 3, 3, 6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 9, 9, 4, 4, 4, 10, 10, 6, 6, 6, 6, 5, 5, 5, 5, 8, 8, 8, 8, 10, 10, 10, 10, 8, 8, 8, 8, 5, 5, 5, 5, 10, 10, 10, 10, 2, 9, 9, 9, 9},
}

// reels lengths [71, 72, 72, 72, 71], total reshuffles 1881543168
// RTP = 84.666(lined) + 4.8607(scatter) = 89.526666%
var Reels895 = slot.Reels5x{
	{3, 3, 3, 3, 7, 7, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 8, 8, 8, 8, 10, 10, 10, 10, 8, 8, 8, 8, 9, 9, 9, 9, 10, 10, 6, 6, 4, 4, 4, 10, 10, 10, 10, 7, 7, 7, 7, 2, 9, 9, 5, 5, 5, 5, 8, 8, 6, 6, 6, 6, 5, 5, 5, 5, 9, 9, 9, 9, 2},
	{10, 10, 6, 6, 2, 7, 7, 7, 7, 8, 8, 7, 7, 1, 5, 5, 5, 5, 9, 9, 9, 9, 5, 5, 5, 5, 9, 9, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 6, 6, 6, 6, 9, 9, 9, 9, 3, 3, 3, 3, 10, 10, 10, 10, 7, 7, 7, 7, 8, 8, 8, 8, 2},
	{10, 10, 6, 6, 2, 7, 7, 7, 7, 8, 8, 7, 7, 1, 5, 5, 5, 5, 9, 9, 9, 9, 5, 5, 5, 5, 9, 9, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 6, 6, 6, 6, 9, 9, 9, 9, 3, 3, 3, 3, 10, 10, 10, 10, 7, 7, 7, 7, 8, 8, 8, 8, 2},
	{10, 10, 6, 6, 2, 7, 7, 7, 7, 8, 8, 7, 7, 1, 5, 5, 5, 5, 9, 9, 9, 9, 5, 5, 5, 5, 9, 9, 4, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 6, 6, 6, 6, 9, 9, 9, 9, 3, 3, 3, 3, 10, 10, 10, 10, 7, 7, 7, 7, 8, 8, 8, 8, 2},
	{3, 3, 3, 3, 7, 7, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 8, 8, 8, 8, 10, 10, 10, 10, 8, 8, 8, 8, 9, 9, 9, 9, 10, 10, 6, 6, 4, 4, 4, 10, 10, 10, 10, 7, 7, 7, 7, 2, 9, 9, 5, 5, 5, 5, 8, 8, 6, 6, 6, 6, 5, 5, 5, 5, 9, 9, 9, 9, 2},
}

// reels lengths [68, 69, 69, 69, 68], total reshuffles 1519025616
// RTP = 85.152(lined) + 5.5631(scatter) = 90.715490%
var Reels907 = slot.Reels5x{
	{9, 9, 9, 9, 7, 7, 7, 7, 9, 9, 9, 9, 8, 8, 8, 8, 2, 5, 5, 5, 7, 7, 5, 5, 5, 5, 6, 6, 6, 6, 10, 10, 10, 10, 3, 3, 3, 3, 2, 4, 4, 4, 10, 10, 4, 4, 4, 4, 9, 9, 8, 8, 7, 7, 7, 7, 6, 6, 6, 6, 10, 10, 10, 10, 8, 8, 8, 8},
	{10, 10, 9, 9, 3, 3, 3, 3, 2, 6, 6, 6, 6, 2, 4, 4, 4, 4, 10, 10, 10, 10, 9, 9, 9, 9, 8, 8, 7, 7, 6, 6, 6, 6, 1, 8, 8, 8, 8, 4, 4, 4, 5, 5, 5, 8, 8, 8, 8, 10, 10, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 7, 7, 7, 7, 9, 9, 9, 9},
	{10, 10, 9, 9, 3, 3, 3, 3, 2, 6, 6, 6, 6, 2, 4, 4, 4, 4, 10, 10, 10, 10, 9, 9, 9, 9, 8, 8, 7, 7, 6, 6, 6, 6, 1, 8, 8, 8, 8, 4, 4, 4, 5, 5, 5, 8, 8, 8, 8, 10, 10, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 7, 7, 7, 7, 9, 9, 9, 9},
	{10, 10, 9, 9, 3, 3, 3, 3, 2, 6, 6, 6, 6, 2, 4, 4, 4, 4, 10, 10, 10, 10, 9, 9, 9, 9, 8, 8, 7, 7, 6, 6, 6, 6, 1, 8, 8, 8, 8, 4, 4, 4, 5, 5, 5, 8, 8, 8, 8, 10, 10, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 7, 7, 7, 7, 9, 9, 9, 9},
	{9, 9, 9, 9, 7, 7, 7, 7, 9, 9, 9, 9, 8, 8, 8, 8, 2, 5, 5, 5, 7, 7, 5, 5, 5, 5, 6, 6, 6, 6, 10, 10, 10, 10, 3, 3, 3, 3, 2, 4, 4, 4, 10, 10, 4, 4, 4, 4, 9, 9, 8, 8, 7, 7, 7, 7, 6, 6, 6, 6, 10, 10, 10, 10, 8, 8, 8, 8},
}

// reels lengths [70, 71, 71, 71, 70], total reshuffles 1753763900
// RTP = 86.673(lined) + 5.081(scatter) = 91.754300%
var Reels917 = slot.Reels5x{
	{8, 8, 8, 8, 2, 7, 7, 10, 10, 10, 10, 2, 4, 4, 4, 4, 9, 9, 9, 9, 6, 6, 6, 6, 10, 10, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 9, 9, 5, 5, 5, 5, 8, 8, 9, 9, 9, 9, 7, 7, 7, 7, 8, 8, 8, 8, 6, 6, 6, 6, 10, 10, 10, 10, 5, 5, 5, 5},
	{5, 5, 5, 5, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 7, 7, 2, 6, 6, 6, 6, 8, 8, 7, 7, 7, 7, 9, 9, 9, 9, 8, 8, 8, 8, 5, 5, 5, 5, 4, 4, 4, 4, 2, 9, 9, 7, 7, 7, 7, 3, 3, 3, 3, 9, 9, 9, 9, 10, 10, 10, 10, 8, 8, 8, 8, 1, 4, 4, 4, 4},
	{5, 5, 5, 5, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 7, 7, 2, 6, 6, 6, 6, 8, 8, 7, 7, 7, 7, 9, 9, 9, 9, 8, 8, 8, 8, 5, 5, 5, 5, 4, 4, 4, 4, 2, 9, 9, 7, 7, 7, 7, 3, 3, 3, 3, 9, 9, 9, 9, 10, 10, 10, 10, 8, 8, 8, 8, 1, 4, 4, 4, 4},
	{5, 5, 5, 5, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 7, 7, 2, 6, 6, 6, 6, 8, 8, 7, 7, 7, 7, 9, 9, 9, 9, 8, 8, 8, 8, 5, 5, 5, 5, 4, 4, 4, 4, 2, 9, 9, 7, 7, 7, 7, 3, 3, 3, 3, 9, 9, 9, 9, 10, 10, 10, 10, 8, 8, 8, 8, 1, 4, 4, 4, 4},
	{8, 8, 8, 8, 2, 7, 7, 10, 10, 10, 10, 2, 4, 4, 4, 4, 9, 9, 9, 9, 6, 6, 6, 6, 10, 10, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 9, 9, 5, 5, 5, 5, 8, 8, 9, 9, 9, 9, 7, 7, 7, 7, 8, 8, 8, 8, 6, 6, 6, 6, 10, 10, 10, 10, 5, 5, 5, 5},
}

// reels lengths [68, 69, 69, 69, 68], total reshuffles 1519025616
// RTP = 86.675(lined) + 5.5631(scatter) = 92.237953%
var Reels922 = slot.Reels5x{
	{7, 7, 7, 7, 10, 10, 8, 8, 3, 3, 3, 3, 5, 5, 5, 5, 10, 10, 10, 10, 2, 9, 9, 9, 9, 7, 7, 7, 7, 8, 8, 8, 8, 4, 4, 4, 4, 6, 6, 6, 6, 9, 9, 9, 9, 5, 5, 5, 5, 2, 6, 6, 6, 6, 4, 4, 9, 9, 8, 8, 8, 8, 10, 10, 10, 10, 7, 7},
	{9, 9, 9, 9, 10, 10, 9, 9, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 4, 4, 4, 4, 9, 9, 9, 9, 2, 4, 4, 10, 10, 10, 10, 7, 7, 7, 7, 10, 10, 10, 10, 6, 6, 6, 6, 2, 7, 7, 8, 8, 8, 8, 3, 3, 3, 3, 8, 8, 8, 8, 5, 5, 5, 5, 8, 8, 1},
	{9, 9, 9, 9, 10, 10, 9, 9, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 4, 4, 4, 4, 9, 9, 9, 9, 2, 4, 4, 10, 10, 10, 10, 7, 7, 7, 7, 10, 10, 10, 10, 6, 6, 6, 6, 2, 7, 7, 8, 8, 8, 8, 3, 3, 3, 3, 8, 8, 8, 8, 5, 5, 5, 5, 8, 8, 1},
	{9, 9, 9, 9, 10, 10, 9, 9, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 4, 4, 4, 4, 9, 9, 9, 9, 2, 4, 4, 10, 10, 10, 10, 7, 7, 7, 7, 10, 10, 10, 10, 6, 6, 6, 6, 2, 7, 7, 8, 8, 8, 8, 3, 3, 3, 3, 8, 8, 8, 8, 5, 5, 5, 5, 8, 8, 1},
	{7, 7, 7, 7, 10, 10, 8, 8, 3, 3, 3, 3, 5, 5, 5, 5, 10, 10, 10, 10, 2, 9, 9, 9, 9, 7, 7, 7, 7, 8, 8, 8, 8, 4, 4, 4, 4, 6, 6, 6, 6, 9, 9, 9, 9, 5, 5, 5, 5, 2, 6, 6, 6, 6, 4, 4, 9, 9, 8, 8, 8, 8, 10, 10, 10, 10, 7, 7},
}

// reels lengths [72, 73, 73, 73, 72], total reshuffles 2016664128
// RTP = 87.99(lined) + 4.653(scatter) = 92.642577%
var Reels926 = slot.Reels5x{
	{10, 10, 6, 6, 5, 5, 5, 5, 7, 7, 7, 7, 9, 9, 9, 9, 5, 5, 5, 5, 7, 7, 3, 3, 3, 3, 10, 10, 10, 10, 2, 7, 7, 7, 7, 6, 6, 6, 6, 2, 8, 8, 4, 4, 4, 4, 8, 8, 8, 8, 4, 4, 4, 4, 8, 8, 8, 8, 9, 9, 9, 9, 6, 6, 6, 6, 10, 10, 10, 10, 9, 9},
	{10, 10, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 1, 9, 9, 4, 4, 4, 4, 8, 8, 8, 8, 7, 7, 9, 9, 9, 9, 3, 3, 3, 3, 10, 10, 2, 6, 6, 6, 6, 10, 10, 10, 10, 5, 5, 5, 5, 7, 7, 7, 7, 2, 9, 9, 9, 9, 8, 8, 8, 8, 4, 4, 4, 4, 6, 6},
	{10, 10, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 1, 9, 9, 4, 4, 4, 4, 8, 8, 8, 8, 7, 7, 9, 9, 9, 9, 3, 3, 3, 3, 10, 10, 2, 6, 6, 6, 6, 10, 10, 10, 10, 5, 5, 5, 5, 7, 7, 7, 7, 2, 9, 9, 9, 9, 8, 8, 8, 8, 4, 4, 4, 4, 6, 6},
	{10, 10, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 1, 9, 9, 4, 4, 4, 4, 8, 8, 8, 8, 7, 7, 9, 9, 9, 9, 3, 3, 3, 3, 10, 10, 2, 6, 6, 6, 6, 10, 10, 10, 10, 5, 5, 5, 5, 7, 7, 7, 7, 2, 9, 9, 9, 9, 8, 8, 8, 8, 4, 4, 4, 4, 6, 6},
	{10, 10, 6, 6, 5, 5, 5, 5, 7, 7, 7, 7, 9, 9, 9, 9, 5, 5, 5, 5, 7, 7, 3, 3, 3, 3, 10, 10, 10, 10, 2, 7, 7, 7, 7, 6, 6, 6, 6, 2, 8, 8, 4, 4, 4, 4, 8, 8, 8, 8, 4, 4, 4, 4, 8, 8, 8, 8, 9, 9, 9, 9, 6, 6, 6, 6, 10, 10, 10, 10, 9, 9},
}

// reels lengths [69, 70, 70, 70, 69], total reshuffles 1633023000
// RTP = 88.166(lined) + 5.3148(scatter) = 93.480408%
var Reels934 = slot.Reels5x{
	{5, 5, 5, 5, 9, 9, 8, 8, 7, 7, 7, 7, 6, 6, 6, 6, 10, 10, 10, 10, 7, 7, 7, 7, 10, 10, 9, 9, 9, 9, 8, 8, 8, 8, 7, 7, 4, 4, 4, 4, 9, 9, 9, 9, 8, 8, 8, 8, 2, 6, 6, 6, 6, 5, 5, 5, 5, 2, 4, 4, 4, 10, 10, 10, 10, 3, 3, 3, 3},
	{6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 9, 9, 10, 10, 10, 10, 5, 5, 5, 5, 2, 4, 4, 4, 4, 3, 3, 3, 3, 7, 7, 6, 6, 6, 6, 7, 7, 7, 7, 10, 10, 9, 9, 1, 9, 9, 9, 9, 4, 4, 4, 8, 8, 10, 10, 10, 10, 8, 8, 8, 8, 2, 8, 8, 8, 8, 5, 5, 5, 5},
	{6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 9, 9, 10, 10, 10, 10, 5, 5, 5, 5, 2, 4, 4, 4, 4, 3, 3, 3, 3, 7, 7, 6, 6, 6, 6, 7, 7, 7, 7, 10, 10, 9, 9, 1, 9, 9, 9, 9, 4, 4, 4, 8, 8, 10, 10, 10, 10, 8, 8, 8, 8, 2, 8, 8, 8, 8, 5, 5, 5, 5},
	{6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 9, 9, 10, 10, 10, 10, 5, 5, 5, 5, 2, 4, 4, 4, 4, 3, 3, 3, 3, 7, 7, 6, 6, 6, 6, 7, 7, 7, 7, 10, 10, 9, 9, 1, 9, 9, 9, 9, 4, 4, 4, 8, 8, 10, 10, 10, 10, 8, 8, 8, 8, 2, 8, 8, 8, 8, 5, 5, 5, 5},
	{5, 5, 5, 5, 9, 9, 8, 8, 7, 7, 7, 7, 6, 6, 6, 6, 10, 10, 10, 10, 7, 7, 7, 7, 10, 10, 9, 9, 9, 9, 8, 8, 8, 8, 7, 7, 4, 4, 4, 4, 9, 9, 9, 9, 8, 8, 8, 8, 2, 6, 6, 6, 6, 5, 5, 5, 5, 2, 4, 4, 4, 10, 10, 10, 10, 3, 3, 3, 3},
}

// reels lengths [67, 68, 68, 68, 67], total reshuffles 1411485248
// RTP = 88.691(lined) + 5.8271(scatter) = 94.518008%
var Reels945 = slot.Reels5x{
	{10, 10, 10, 10, 8, 8, 8, 8, 9, 9, 9, 9, 4, 4, 4, 4, 2, 6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 10, 10, 10, 10, 6, 6, 6, 6, 5, 5, 5, 5, 9, 9, 9, 9, 4, 4, 4, 10, 10, 5, 5, 5, 5, 8, 8, 8, 8, 7, 7, 7, 7, 8, 8, 3, 3, 3, 3, 2},
	{10, 10, 6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 10, 10, 10, 10, 8, 8, 8, 8, 4, 4, 4, 8, 8, 1, 5, 5, 5, 5, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 2, 5, 5, 5, 5, 10, 10, 10, 10, 9, 9, 9, 9, 3, 3, 3, 3, 2, 8, 8, 8, 8, 9, 9, 9, 9},
	{10, 10, 6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 10, 10, 10, 10, 8, 8, 8, 8, 4, 4, 4, 8, 8, 1, 5, 5, 5, 5, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 2, 5, 5, 5, 5, 10, 10, 10, 10, 9, 9, 9, 9, 3, 3, 3, 3, 2, 8, 8, 8, 8, 9, 9, 9, 9},
	{10, 10, 6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 10, 10, 10, 10, 8, 8, 8, 8, 4, 4, 4, 8, 8, 1, 5, 5, 5, 5, 6, 6, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 2, 5, 5, 5, 5, 10, 10, 10, 10, 9, 9, 9, 9, 3, 3, 3, 3, 2, 8, 8, 8, 8, 9, 9, 9, 9},
	{10, 10, 10, 10, 8, 8, 8, 8, 9, 9, 9, 9, 4, 4, 4, 4, 2, 6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 10, 10, 10, 10, 6, 6, 6, 6, 5, 5, 5, 5, 9, 9, 9, 9, 4, 4, 4, 10, 10, 5, 5, 5, 5, 8, 8, 8, 8, 7, 7, 7, 7, 8, 8, 3, 3, 3, 3, 2},
}

// reels lengths [69, 70, 70, 70, 69], total reshuffles 1633023000
// RTP = 90.66(lined) + 5.3148(scatter) = 95.974622%
var Reels959 = slot.Reels5x{
	{9, 9, 9, 9, 2, 9, 9, 8, 8, 8, 8, 5, 5, 5, 5, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 10, 10, 10, 10, 9, 9, 9, 9, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 10, 10, 10, 10, 6, 6, 6, 6, 2, 3, 3, 3, 3, 8, 8},
	{9, 9, 4, 4, 4, 5, 5, 5, 5, 10, 10, 10, 10, 7, 7, 7, 7, 9, 9, 9, 9, 6, 6, 10, 10, 10, 10, 2, 8, 8, 9, 9, 9, 9, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 10, 10, 2, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 8, 8, 6, 6, 6, 6, 8, 8, 8, 8, 1},
	{9, 9, 4, 4, 4, 5, 5, 5, 5, 10, 10, 10, 10, 7, 7, 7, 7, 9, 9, 9, 9, 6, 6, 10, 10, 10, 10, 2, 8, 8, 9, 9, 9, 9, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 10, 10, 2, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 8, 8, 6, 6, 6, 6, 8, 8, 8, 8, 1},
	{9, 9, 4, 4, 4, 5, 5, 5, 5, 10, 10, 10, 10, 7, 7, 7, 7, 9, 9, 9, 9, 6, 6, 10, 10, 10, 10, 2, 8, 8, 9, 9, 9, 9, 4, 4, 4, 4, 7, 7, 7, 7, 3, 3, 3, 3, 10, 10, 2, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 8, 8, 6, 6, 6, 6, 8, 8, 8, 8, 1},
	{9, 9, 9, 9, 2, 9, 9, 8, 8, 8, 8, 5, 5, 5, 5, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 4, 4, 4, 8, 8, 8, 8, 6, 6, 6, 6, 10, 10, 10, 10, 9, 9, 9, 9, 6, 6, 4, 4, 4, 4, 7, 7, 7, 7, 10, 10, 10, 10, 6, 6, 6, 6, 2, 3, 3, 3, 3, 8, 8},
}

// reels lengths [71, 72, 72, 72, 71], total reshuffles 1881543168
// RTP = 91.654(lined) + 4.8607(scatter) = 96.514965%
var Reels965 = slot.Reels5x{
	{6, 6, 6, 6, 9, 9, 9, 9, 10, 10, 10, 10, 5, 5, 5, 5, 3, 3, 3, 3, 4, 4, 4, 4, 7, 7, 10, 10, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 2, 8, 8, 5, 5, 5, 5, 2, 10, 10, 10, 10, 7, 7, 7, 7, 4, 4, 4, 4, 9, 9, 9, 9, 5, 7, 7, 7, 7, 8, 8, 8, 8},
	{8, 8, 8, 8, 3, 3, 3, 3, 6, 6, 6, 6, 9, 9, 4, 4, 4, 4, 5, 7, 7, 7, 7, 2, 5, 5, 5, 5, 7, 7, 7, 7, 9, 9, 9, 9, 4, 4, 4, 4, 9, 9, 9, 9, 1, 8, 8, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 10, 10, 5, 5, 5, 5, 10, 10, 2, 7, 7, 8, 8, 8, 8},
	{8, 8, 8, 8, 3, 3, 3, 3, 6, 6, 6, 6, 9, 9, 4, 4, 4, 4, 5, 7, 7, 7, 7, 2, 5, 5, 5, 5, 7, 7, 7, 7, 9, 9, 9, 9, 4, 4, 4, 4, 9, 9, 9, 9, 1, 8, 8, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 10, 10, 5, 5, 5, 5, 10, 10, 2, 7, 7, 8, 8, 8, 8},
	{8, 8, 8, 8, 3, 3, 3, 3, 6, 6, 6, 6, 9, 9, 4, 4, 4, 4, 5, 7, 7, 7, 7, 2, 5, 5, 5, 5, 7, 7, 7, 7, 9, 9, 9, 9, 4, 4, 4, 4, 9, 9, 9, 9, 1, 8, 8, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 10, 10, 5, 5, 5, 5, 10, 10, 2, 7, 7, 8, 8, 8, 8},
	{6, 6, 6, 6, 9, 9, 9, 9, 10, 10, 10, 10, 5, 5, 5, 5, 3, 3, 3, 3, 4, 4, 4, 4, 7, 7, 10, 10, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 2, 8, 8, 5, 5, 5, 5, 2, 10, 10, 10, 10, 7, 7, 7, 7, 4, 4, 4, 4, 9, 9, 9, 9, 5, 7, 7, 7, 7, 8, 8, 8, 8},
}

// reels lengths [64, 65, 65, 65, 64], total reshuffles 1124864000
// RTP = 90.301(lined) + 6.7267(scatter) = 97.027255%
var Reels970 = slot.Reels5x{
	{8, 8, 8, 8, 6, 6, 6, 6, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 10, 10, 10, 10, 5, 5, 5, 10, 10, 2, 8, 8, 8, 8, 7, 7, 7, 7, 9, 9, 9, 9, 10, 10, 10, 10, 9, 9, 7, 7, 7, 7, 2, 4, 4, 4, 4, 6, 6, 6, 6, 9, 9, 9, 9},
	{9, 9, 9, 9, 5, 5, 5, 6, 6, 6, 6, 2, 10, 10, 4, 4, 4, 1, 8, 8, 8, 8, 2, 5, 5, 5, 5, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 4, 8, 8, 8, 8, 7, 7, 7, 7, 9, 9, 9, 9, 3, 3, 3, 3, 9, 9, 10, 10, 10, 10, 7, 7, 7, 7},
	{9, 9, 9, 9, 5, 5, 5, 6, 6, 6, 6, 2, 10, 10, 4, 4, 4, 1, 8, 8, 8, 8, 2, 5, 5, 5, 5, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 4, 8, 8, 8, 8, 7, 7, 7, 7, 9, 9, 9, 9, 3, 3, 3, 3, 9, 9, 10, 10, 10, 10, 7, 7, 7, 7},
	{9, 9, 9, 9, 5, 5, 5, 6, 6, 6, 6, 2, 10, 10, 4, 4, 4, 1, 8, 8, 8, 8, 2, 5, 5, 5, 5, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 4, 8, 8, 8, 8, 7, 7, 7, 7, 9, 9, 9, 9, 3, 3, 3, 3, 9, 9, 10, 10, 10, 10, 7, 7, 7, 7},
	{8, 8, 8, 8, 6, 6, 6, 6, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 10, 10, 10, 10, 5, 5, 5, 10, 10, 2, 8, 8, 8, 8, 7, 7, 7, 7, 9, 9, 9, 9, 10, 10, 10, 10, 9, 9, 7, 7, 7, 7, 2, 4, 4, 4, 4, 6, 6, 6, 6, 9, 9, 9, 9},
}

// reels lengths [64, 65, 65, 65, 64], total reshuffles 1124864000
// RTP = 91.174(lined) + 6.7267(scatter) = 97.900896%
var Reels979 = slot.Reels5x{
	{10, 10, 4, 4, 4, 10, 10, 10, 10, 9, 9, 9, 9, 3, 3, 3, 3, 9, 9, 9, 9, 8, 8, 8, 8, 2, 8, 8, 8, 8, 7, 7, 7, 7, 2, 4, 4, 4, 4, 7, 7, 7, 7, 9, 9, 6, 6, 6, 6, 10, 10, 10, 10, 5, 5, 5, 6, 6, 6, 6, 5, 5, 5, 5},
	{5, 5, 5, 5, 10, 10, 10, 10, 8, 8, 8, 8, 5, 5, 5, 10, 10, 10, 10, 4, 4, 4, 4, 8, 8, 8, 8, 2, 9, 9, 10, 10, 6, 6, 6, 6, 4, 4, 4, 6, 6, 6, 6, 1, 3, 3, 3, 3, 9, 9, 9, 9, 7, 7, 7, 7, 9, 9, 9, 9, 2, 7, 7, 7, 7},
	{5, 5, 5, 5, 10, 10, 10, 10, 8, 8, 8, 8, 5, 5, 5, 10, 10, 10, 10, 4, 4, 4, 4, 8, 8, 8, 8, 2, 9, 9, 10, 10, 6, 6, 6, 6, 4, 4, 4, 6, 6, 6, 6, 1, 3, 3, 3, 3, 9, 9, 9, 9, 7, 7, 7, 7, 9, 9, 9, 9, 2, 7, 7, 7, 7},
	{5, 5, 5, 5, 10, 10, 10, 10, 8, 8, 8, 8, 5, 5, 5, 10, 10, 10, 10, 4, 4, 4, 4, 8, 8, 8, 8, 2, 9, 9, 10, 10, 6, 6, 6, 6, 4, 4, 4, 6, 6, 6, 6, 1, 3, 3, 3, 3, 9, 9, 9, 9, 7, 7, 7, 7, 9, 9, 9, 9, 2, 7, 7, 7, 7},
	{10, 10, 4, 4, 4, 10, 10, 10, 10, 9, 9, 9, 9, 3, 3, 3, 3, 9, 9, 9, 9, 8, 8, 8, 8, 2, 8, 8, 8, 8, 7, 7, 7, 7, 2, 4, 4, 4, 4, 7, 7, 7, 7, 9, 9, 6, 6, 6, 6, 10, 10, 10, 10, 5, 5, 5, 6, 6, 6, 6, 5, 5, 5, 5},
}

// reels lengths [68, 69, 69, 69, 68], total reshuffles 1519025616
// RTP = 92.735(lined) + 5.5631(scatter) = 98.297931%
var Reels982 = slot.Reels5x{
	{7, 7, 7, 7, 10, 10, 10, 10, 2, 5, 5, 5, 5, 8, 8, 9, 9, 9, 9, 4, 4, 4, 4, 7, 7, 7, 7, 2, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 3, 3, 3, 3, 10, 10, 4, 4, 4, 4, 10, 10, 10, 10, 9, 9, 9, 9, 8, 8, 8, 8},
	{8, 8, 8, 8, 4, 4, 4, 4, 9, 9, 9, 9, 10, 10, 2, 9, 9, 7, 7, 7, 7, 9, 9, 9, 9, 10, 10, 10, 10, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 1, 6, 6, 6, 6, 7, 7, 7, 7, 2, 5, 5, 5, 5, 8, 8, 8, 8, 10, 10, 10, 10},
	{8, 8, 8, 8, 4, 4, 4, 4, 9, 9, 9, 9, 10, 10, 2, 9, 9, 7, 7, 7, 7, 9, 9, 9, 9, 10, 10, 10, 10, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 1, 6, 6, 6, 6, 7, 7, 7, 7, 2, 5, 5, 5, 5, 8, 8, 8, 8, 10, 10, 10, 10},
	{8, 8, 8, 8, 4, 4, 4, 4, 9, 9, 9, 9, 10, 10, 2, 9, 9, 7, 7, 7, 7, 9, 9, 9, 9, 10, 10, 10, 10, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 1, 6, 6, 6, 6, 7, 7, 7, 7, 2, 5, 5, 5, 5, 8, 8, 8, 8, 10, 10, 10, 10},
	{7, 7, 7, 7, 10, 10, 10, 10, 2, 5, 5, 5, 5, 8, 8, 9, 9, 9, 9, 4, 4, 4, 4, 7, 7, 7, 7, 2, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 3, 3, 3, 3, 10, 10, 4, 4, 4, 4, 10, 10, 10, 10, 9, 9, 9, 9, 8, 8, 8, 8},
}

// reels lengths [65, 66, 66, 66, 65], total reshuffles 1214670600
// RTP = 92.313(lined) + 6.4075(scatter) = 98.720245%
var Reels987 = slot.Reels5x{
	{10, 10, 10, 10, 8, 8, 8, 8, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 8, 8, 2, 10, 10, 9, 9, 10, 10, 10, 10, 9, 9, 9, 9, 6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 9, 9, 5, 5, 5, 5, 4, 4, 4, 2, 7, 7, 7, 7, 3, 3, 3, 3, 4, 4, 4, 4},
	{6, 6, 6, 6, 5, 5, 5, 5, 7, 7, 7, 7, 2, 7, 7, 7, 7, 3, 3, 3, 3, 10, 10, 10, 10, 8, 8, 8, 8, 5, 5, 5, 5, 8, 8, 8, 8, 9, 9, 9, 9, 4, 4, 4, 1, 10, 10, 10, 10, 2, 4, 4, 4, 4, 9, 9, 9, 9, 10, 10, 9, 9, 6, 6, 6, 6},
	{6, 6, 6, 6, 5, 5, 5, 5, 7, 7, 7, 7, 2, 7, 7, 7, 7, 3, 3, 3, 3, 10, 10, 10, 10, 8, 8, 8, 8, 5, 5, 5, 5, 8, 8, 8, 8, 9, 9, 9, 9, 4, 4, 4, 1, 10, 10, 10, 10, 2, 4, 4, 4, 4, 9, 9, 9, 9, 10, 10, 9, 9, 6, 6, 6, 6},
	{6, 6, 6, 6, 5, 5, 5, 5, 7, 7, 7, 7, 2, 7, 7, 7, 7, 3, 3, 3, 3, 10, 10, 10, 10, 8, 8, 8, 8, 5, 5, 5, 5, 8, 8, 8, 8, 9, 9, 9, 9, 4, 4, 4, 1, 10, 10, 10, 10, 2, 4, 4, 4, 4, 9, 9, 9, 9, 10, 10, 9, 9, 6, 6, 6, 6},
	{10, 10, 10, 10, 8, 8, 8, 8, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 8, 8, 2, 10, 10, 9, 9, 10, 10, 10, 10, 9, 9, 9, 9, 6, 6, 6, 6, 7, 7, 7, 7, 9, 9, 9, 9, 5, 5, 5, 5, 4, 4, 4, 2, 7, 7, 7, 7, 3, 3, 3, 3, 4, 4, 4, 4},
}

// reels lengths [64, 65, 65, 65, 64], total reshuffles 1124864000
// RTP = 93.775(lined) + 6.7267(scatter) = 100.501424%
var Reels100 = slot.Reels5x{
	{5, 5, 5, 5, 7, 7, 7, 7, 9, 9, 9, 9, 2, 4, 4, 8, 8, 8, 8, 7, 7, 7, 7, 10, 10, 9, 9, 10, 10, 10, 10, 6, 6, 6, 6, 9, 9, 9, 9, 10, 10, 10, 10, 6, 6, 6, 6, 2, 3, 3, 3, 3, 5, 5, 5, 5, 8, 8, 8, 8, 4, 4, 4, 4},
	{4, 4, 4, 4, 5, 5, 5, 5, 2, 6, 6, 6, 6, 1, 10, 10, 10, 10, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 10, 10, 10, 10, 7, 7, 7, 7, 9, 9, 9, 9, 4, 4, 3, 3, 3, 3, 9, 9, 9, 9, 5, 5, 5, 5, 10, 10, 2, 8, 8, 8, 8, 7, 7, 7, 7},
	{4, 4, 4, 4, 5, 5, 5, 5, 2, 6, 6, 6, 6, 1, 10, 10, 10, 10, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 10, 10, 10, 10, 7, 7, 7, 7, 9, 9, 9, 9, 4, 4, 3, 3, 3, 3, 9, 9, 9, 9, 5, 5, 5, 5, 10, 10, 2, 8, 8, 8, 8, 7, 7, 7, 7},
	{4, 4, 4, 4, 5, 5, 5, 5, 2, 6, 6, 6, 6, 1, 10, 10, 10, 10, 8, 8, 8, 8, 6, 6, 6, 6, 9, 9, 10, 10, 10, 10, 7, 7, 7, 7, 9, 9, 9, 9, 4, 4, 3, 3, 3, 3, 9, 9, 9, 9, 5, 5, 5, 5, 10, 10, 2, 8, 8, 8, 8, 7, 7, 7, 7},
	{5, 5, 5, 5, 7, 7, 7, 7, 9, 9, 9, 9, 2, 4, 4, 8, 8, 8, 8, 7, 7, 7, 7, 10, 10, 9, 9, 10, 10, 10, 10, 6, 6, 6, 6, 9, 9, 9, 9, 10, 10, 10, 10, 6, 6, 6, 6, 2, 3, 3, 3, 3, 5, 5, 5, 5, 8, 8, 8, 8, 4, 4, 4, 4},
}

// reels lengths [65, 66, 66, 66, 65], total reshuffles 1214670600
// RTP = 94.609(lined) + 6.4075(scatter) = 101.016005%
var Reels101 = slot.Reels5x{
	{5, 5, 5, 5, 7, 7, 7, 7, 10, 10, 9, 9, 9, 9, 10, 10, 10, 10, 6, 6, 6, 6, 5, 5, 5, 5, 4, 4, 4, 4, 2, 9, 9, 9, 9, 8, 8, 8, 8, 9, 9, 4, 4, 4, 7, 7, 7, 7, 8, 8, 8, 8, 2, 6, 6, 6, 6, 10, 10, 10, 10, 3, 3, 3, 3},
	{8, 8, 8, 8, 9, 9, 9, 9, 3, 3, 3, 3, 8, 8, 8, 8, 9, 9, 5, 5, 5, 5, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 6, 6, 6, 6, 2, 9, 9, 9, 9, 2, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 7, 7, 7, 7, 1, 4, 4, 4, 4},
	{8, 8, 8, 8, 9, 9, 9, 9, 3, 3, 3, 3, 8, 8, 8, 8, 9, 9, 5, 5, 5, 5, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 6, 6, 6, 6, 2, 9, 9, 9, 9, 2, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 7, 7, 7, 7, 1, 4, 4, 4, 4},
	{8, 8, 8, 8, 9, 9, 9, 9, 3, 3, 3, 3, 8, 8, 8, 8, 9, 9, 5, 5, 5, 5, 10, 10, 10, 10, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 6, 6, 6, 6, 2, 9, 9, 9, 9, 2, 10, 10, 7, 7, 7, 7, 5, 5, 5, 5, 7, 7, 7, 7, 1, 4, 4, 4, 4},
	{5, 5, 5, 5, 7, 7, 7, 7, 10, 10, 9, 9, 9, 9, 10, 10, 10, 10, 6, 6, 6, 6, 5, 5, 5, 5, 4, 4, 4, 4, 2, 9, 9, 9, 9, 8, 8, 8, 8, 9, 9, 4, 4, 4, 7, 7, 7, 7, 8, 8, 8, 8, 2, 6, 6, 6, 6, 10, 10, 10, 10, 3, 3, 3, 3},
}

// reels lengths [66, 67, 67, 67, 66], total reshuffles 1310123628
// RTP = 98.324(lined) + 6.1081(scatter) = 104.432531%
var Reels104 = slot.Reels5x{
	{4, 4, 4, 4, 9, 9, 9, 9, 8, 8, 8, 8, 6, 6, 6, 6, 10, 10, 10, 10, 2, 9, 9, 9, 9, 4, 4, 4, 4, 2, 8, 8, 8, 8, 10, 10, 6, 6, 6, 6, 7, 7, 7, 7, 3, 3, 3, 3, 5, 5, 5, 5, 7, 7, 7, 7, 10, 10, 10, 10, 9, 9, 5, 5, 5, 5},
	{7, 7, 7, 7, 4, 4, 4, 4, 10, 10, 10, 10, 9, 9, 9, 9, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 8, 8, 9, 9, 2, 4, 4, 4, 4, 6, 6, 6, 6, 5, 5, 5, 5, 10, 10, 10, 10, 1, 9, 9, 9, 9, 2, 8, 8, 8, 8, 7, 7, 7, 7, 10, 10, 3, 3, 3, 3},
	{7, 7, 7, 7, 4, 4, 4, 4, 10, 10, 10, 10, 9, 9, 9, 9, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 8, 8, 9, 9, 2, 4, 4, 4, 4, 6, 6, 6, 6, 5, 5, 5, 5, 10, 10, 10, 10, 1, 9, 9, 9, 9, 2, 8, 8, 8, 8, 7, 7, 7, 7, 10, 10, 3, 3, 3, 3},
	{7, 7, 7, 7, 4, 4, 4, 4, 10, 10, 10, 10, 9, 9, 9, 9, 5, 5, 5, 5, 6, 6, 6, 6, 8, 8, 8, 8, 9, 9, 2, 4, 4, 4, 4, 6, 6, 6, 6, 5, 5, 5, 5, 10, 10, 10, 10, 1, 9, 9, 9, 9, 2, 8, 8, 8, 8, 7, 7, 7, 7, 10, 10, 3, 3, 3, 3},
	{4, 4, 4, 4, 9, 9, 9, 9, 8, 8, 8, 8, 6, 6, 6, 6, 10, 10, 10, 10, 2, 9, 9, 9, 9, 4, 4, 4, 4, 2, 8, 8, 8, 8, 10, 10, 6, 6, 6, 6, 7, 7, 7, 7, 3, 3, 3, 3, 5, 5, 5, 5, 7, 7, 7, 7, 10, 10, 10, 10, 9, 9, 5, 5, 5, 5},
}

// reels lengths [64, 65, 65, 65, 64], total reshuffles 1124864000
// RTP = 105.67(lined) + 6.7267(scatter) = 112.396472%
var Reels112 = slot.Reels5x{
	{5, 5, 5, 5, 2, 9, 9, 9, 9, 10, 10, 9, 9, 9, 9, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 4, 8, 8, 8, 8, 10, 10, 10, 10, 7, 7, 7, 7, 2, 4, 4, 4, 4, 7, 7, 7, 7, 8, 8, 8, 8, 6, 6, 6, 6, 5, 5, 5, 5, 3, 3, 3, 3},
	{8, 8, 8, 8, 1, 8, 8, 8, 8, 5, 5, 5, 5, 7, 7, 7, 7, 6, 6, 6, 6, 7, 7, 7, 7, 10, 10, 10, 10, 6, 6, 6, 6, 5, 5, 5, 5, 2, 9, 9, 9, 9, 4, 4, 4, 4, 10, 10, 10, 10, 2, 3, 3, 3, 3, 10, 10, 4, 4, 4, 4, 9, 9, 9, 9},
	{8, 8, 8, 8, 1, 8, 8, 8, 8, 5, 5, 5, 5, 7, 7, 7, 7, 6, 6, 6, 6, 7, 7, 7, 7, 10, 10, 10, 10, 6, 6, 6, 6, 5, 5, 5, 5, 2, 9, 9, 9, 9, 4, 4, 4, 4, 10, 10, 10, 10, 2, 3, 3, 3, 3, 10, 10, 4, 4, 4, 4, 9, 9, 9, 9},
	{8, 8, 8, 8, 1, 8, 8, 8, 8, 5, 5, 5, 5, 7, 7, 7, 7, 6, 6, 6, 6, 7, 7, 7, 7, 10, 10, 10, 10, 6, 6, 6, 6, 5, 5, 5, 5, 2, 9, 9, 9, 9, 4, 4, 4, 4, 10, 10, 10, 10, 2, 3, 3, 3, 3, 10, 10, 4, 4, 4, 4, 9, 9, 9, 9},
	{5, 5, 5, 5, 2, 9, 9, 9, 9, 10, 10, 9, 9, 9, 9, 6, 6, 6, 6, 10, 10, 10, 10, 4, 4, 4, 4, 8, 8, 8, 8, 10, 10, 10, 10, 7, 7, 7, 7, 2, 4, 4, 4, 4, 7, 7, 7, 7, 8, 8, 8, 8, 6, 6, 6, 6, 5, 5, 5, 5, 3, 3, 3, 3},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	83.829264:  &Reels838,
	87.903734:  &Reels879,
	89.526666:  &Reels895,
	90.715490:  &Reels907,
	91.754300:  &Reels917,
	92.237953:  &Reels922,
	92.642577:  &Reels926,
	93.480408:  &Reels934,
	94.518008:  &Reels945,
	95.974622:  &Reels959,
	96.514965:  &Reels965,
	97.027255:  &Reels970,
	97.900896:  &Reels979,
	98.297931:  &Reels982,
	98.720245:  &Reels987,
	100.501424: &Reels100,
	101.016005: &Reels101,
	104.432531: &Reels104,
	112.396472: &Reels112,
}

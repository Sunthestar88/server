package sevenhot

import (
	"github.com/slotopol/server/game/slot"
)

// reels lengths [25, 25, 26, 25, 25], total reshuffles 10156250
// RTP = 88.74(lined) + 5.8741(scatter) = 94.614174%
var Reels946 = slot.Reels5x{
	{4, 7, 6, 6, 6, 6, 7, 7, 7, 5, 1, 2, 2, 2, 1, 4, 4, 4, 8, 5, 5, 5, 3, 3, 3},
	{4, 7, 6, 6, 6, 6, 7, 7, 7, 5, 1, 2, 2, 2, 1, 4, 4, 4, 8, 5, 5, 5, 3, 3, 3},
	{5, 5, 5, 8, 7, 7, 7, 4, 4, 4, 7, 6, 6, 6, 1, 3, 3, 3, 8, 2, 2, 2, 4, 5, 1, 6},
	{4, 7, 6, 6, 6, 6, 7, 7, 7, 5, 1, 2, 2, 2, 1, 4, 4, 4, 8, 5, 5, 5, 3, 3, 3},
	{4, 7, 6, 6, 6, 6, 7, 7, 7, 5, 1, 2, 2, 2, 1, 4, 4, 4, 8, 5, 5, 5, 3, 3, 3},
}

// reels lengths [28, 29, 28, 29, 28], total reshuffles 18461632
// RTP = 89.17(lined) + 6.4025(scatter) = 95.572677%
var Reels955 = slot.Reels5x{
	{5, 5, 5, 5, 5, 6, 6, 6, 7, 7, 7, 1, 4, 6, 6, 1, 2, 2, 2, 8, 7, 7, 3, 3, 3, 4, 4, 4},
	{4, 4, 4, 7, 7, 1, 5, 5, 5, 6, 6, 6, 8, 3, 3, 3, 8, 2, 2, 2, 5, 5, 4, 7, 7, 7, 1, 6, 6},
	{5, 5, 5, 5, 5, 6, 6, 6, 7, 7, 7, 1, 4, 6, 6, 1, 2, 2, 2, 8, 7, 7, 3, 3, 3, 4, 4, 4},
	{4, 4, 4, 7, 7, 1, 5, 5, 5, 6, 6, 6, 8, 3, 3, 3, 8, 2, 2, 2, 5, 5, 4, 7, 7, 7, 1, 6, 6},
	{5, 5, 5, 5, 5, 6, 6, 6, 7, 7, 7, 1, 4, 6, 6, 1, 2, 2, 2, 8, 7, 7, 3, 3, 3, 4, 4, 4},
}

// reels lengths [27, 27, 28, 27, 27], total reshuffles 14880348
// RTP = 91.609(lined) + 4.6356(scatter) = 96.244799%
var Reels962 = slot.Reels5x{
	{7, 7, 7, 3, 3, 3, 4, 1, 7, 7, 4, 4, 4, 5, 1, 6, 6, 6, 2, 2, 2, 8, 5, 5, 5, 6, 6},
	{7, 7, 7, 3, 3, 3, 4, 1, 7, 7, 4, 4, 4, 5, 1, 6, 6, 6, 2, 2, 2, 8, 5, 5, 5, 6, 6},
	{7, 7, 7, 1, 6, 6, 6, 5, 5, 5, 4, 7, 7, 8, 6, 6, 5, 3, 3, 3, 4, 4, 4, 1, 2, 2, 2, 8},
	{7, 7, 7, 3, 3, 3, 4, 1, 7, 7, 4, 4, 4, 5, 1, 6, 6, 6, 2, 2, 2, 8, 5, 5, 5, 6, 6},
	{7, 7, 7, 3, 3, 3, 4, 1, 7, 7, 4, 4, 4, 5, 1, 6, 6, 6, 2, 2, 2, 8, 5, 5, 5, 6, 6},
}

// reels lengths [27, 28, 27, 28, 27], total reshuffles 15431472
// RTP = 90.11(lined) + 7.1569(scatter) = 97.266696%
var Reels972 = slot.Reels5x{
	{7, 7, 7, 3, 3, 3, 4, 1, 7, 7, 4, 4, 4, 5, 1, 6, 6, 6, 2, 2, 2, 8, 5, 5, 5, 6, 6},
	{7, 7, 7, 1, 6, 6, 6, 5, 5, 5, 4, 7, 7, 8, 6, 6, 5, 3, 3, 3, 4, 4, 4, 1, 2, 2, 2, 8},
	{7, 7, 7, 3, 3, 3, 4, 1, 7, 7, 4, 4, 4, 5, 1, 6, 6, 6, 2, 2, 2, 8, 5, 5, 5, 6, 6},
	{7, 7, 7, 1, 6, 6, 6, 5, 5, 5, 4, 7, 7, 8, 6, 6, 5, 3, 3, 3, 4, 4, 4, 1, 2, 2, 2, 8},
	{7, 7, 7, 3, 3, 3, 4, 1, 7, 7, 4, 4, 4, 5, 1, 6, 6, 6, 2, 2, 2, 8, 5, 5, 5, 6, 6},
}

// reels lengths [24, 24, 25, 24, 24], total reshuffles 8294400
// RTP = 91.463(lined) + 6.6621(scatter) = 98.125506%
var Reels981 = slot.Reels5x{
	{4, 4, 4, 6, 6, 6, 8, 2, 2, 2, 1, 7, 5, 6, 5, 5, 5, 7, 7, 7, 1, 3, 3, 3},
	{4, 4, 4, 6, 6, 6, 8, 2, 2, 2, 1, 7, 5, 6, 5, 5, 5, 7, 7, 7, 1, 3, 3, 3},
	{7, 7, 7, 5, 5, 5, 7, 1, 2, 2, 2, 8, 5, 4, 4, 4, 6, 1, 3, 3, 3, 6, 6, 6, 8},
	{4, 4, 4, 6, 6, 6, 8, 2, 2, 2, 1, 7, 5, 6, 5, 5, 5, 7, 7, 7, 1, 3, 3, 3},
	{4, 4, 4, 6, 6, 6, 8, 2, 2, 2, 1, 7, 5, 6, 5, 5, 5, 7, 7, 7, 1, 3, 3, 3},
}

// reels lengths [26, 27, 26, 27, 26], total reshuffles 12812904
// RTP = 91.02(lined) + 8.0349(scatter) = 99.055140%
var Reels990 = slot.Reels5x{
	{2, 2, 2, 7, 7, 7, 1, 3, 3, 3, 8, 4, 6, 5, 4, 4, 4, 1, 5, 5, 5, 6, 6, 6, 7, 7},
	{4, 4, 4, 3, 3, 3, 7, 7, 8, 5, 6, 6, 6, 4, 2, 2, 2, 1, 7, 7, 7, 1, 5, 5, 5, 8, 6},
	{2, 2, 2, 7, 7, 7, 1, 3, 3, 3, 8, 4, 6, 5, 4, 4, 4, 1, 5, 5, 5, 6, 6, 6, 7, 7},
	{4, 4, 4, 3, 3, 3, 7, 7, 8, 5, 6, 6, 6, 4, 2, 2, 2, 1, 7, 7, 7, 1, 5, 5, 5, 8, 6},
	{2, 2, 2, 7, 7, 7, 1, 3, 3, 3, 8, 4, 6, 5, 4, 4, 4, 1, 5, 5, 5, 6, 6, 6, 7, 7},
}

// reels lengths [27, 30, 27, 30, 27], total reshuffles 17714700
// RTP = 93.188(lined) + 6.4746(scatter) = 99.662314%
var Reels996 = slot.Reels5x{
	{4, 3, 3, 3, 1, 2, 2, 2, 6, 6, 5, 5, 5, 8, 6, 6, 6, 1, 4, 4, 4, 7, 7, 7, 7, 7, 5},
	{5, 5, 5, 1, 3, 3, 3, 5, 1, 4, 6, 6, 6, 8, 4, 4, 4, 7, 7, 7, 8, 7, 7, 7, 6, 6, 6, 2, 2, 2},
	{4, 3, 3, 3, 1, 2, 2, 2, 6, 6, 5, 5, 5, 8, 6, 6, 6, 1, 4, 4, 4, 7, 7, 7, 7, 7, 5},
	{5, 5, 5, 1, 3, 3, 3, 5, 1, 4, 6, 6, 6, 8, 4, 4, 4, 7, 7, 7, 8, 7, 7, 7, 6, 6, 6, 2, 2, 2},
	{4, 3, 3, 3, 1, 2, 2, 2, 6, 6, 5, 5, 5, 8, 6, 6, 6, 1, 4, 4, 4, 7, 7, 7, 7, 7, 5},
}

// reels lengths [27, 28, 27, 28, 27], total reshuffles 15431472
// RTP = 93.296(lined) + 7.1569(scatter) = 100.452458%
var Reels100 = slot.Reels5x{
	{4, 4, 4, 6, 6, 1, 5, 5, 5, 7, 7, 6, 6, 6, 4, 8, 2, 2, 7, 7, 7, 1, 5, 5, 3, 3, 3},
	{4, 4, 4, 7, 7, 1, 7, 7, 7, 8, 5, 5, 5, 8, 3, 3, 3, 2, 2, 5, 5, 6, 6, 4, 6, 6, 6, 1},
	{4, 4, 4, 6, 6, 1, 5, 5, 5, 7, 7, 6, 6, 6, 4, 8, 2, 2, 7, 7, 7, 1, 5, 5, 3, 3, 3},
	{4, 4, 4, 7, 7, 1, 7, 7, 7, 8, 5, 5, 5, 8, 3, 3, 3, 2, 2, 5, 5, 6, 6, 4, 6, 6, 6, 1},
	{4, 4, 4, 6, 6, 1, 5, 5, 5, 7, 7, 6, 6, 6, 4, 8, 2, 2, 7, 7, 7, 1, 5, 5, 3, 3, 3},
}

// reels lengths [27, 28, 27, 28, 27], total reshuffles 15431472
// RTP = 98.318(lined) + 7.1569(scatter) = 105.474371%
var Reels105 = slot.Reels5x{
	{3, 3, 3, 2, 5, 6, 7, 5, 5, 5, 3, 4, 4, 4, 2, 2, 2, 1, 4, 6, 6, 6, 1, 7, 7, 7, 8},
	{7, 7, 7, 5, 2, 2, 2, 7, 4, 8, 6, 4, 4, 4, 6, 6, 6, 1, 3, 3, 3, 3, 2, 8, 5, 5, 5, 1},
	{3, 3, 3, 2, 5, 6, 7, 5, 5, 5, 3, 4, 4, 4, 2, 2, 2, 1, 4, 6, 6, 6, 1, 7, 7, 7, 8},
	{7, 7, 7, 5, 2, 2, 2, 7, 4, 8, 6, 4, 4, 4, 6, 6, 6, 1, 3, 3, 3, 3, 2, 8, 5, 5, 5, 1},
	{3, 3, 3, 2, 5, 6, 7, 5, 5, 5, 3, 4, 4, 4, 2, 2, 2, 1, 4, 6, 6, 6, 1, 7, 7, 7, 8},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	94.614174:  &Reels946,
	95.572677:  &Reels955,
	96.244799:  &Reels962,
	97.266696:  &Reels972,
	98.125506:  &Reels981,
	99.055140:  &Reels990,
	99.662314:  &Reels996,
	100.452458: &Reels100,
	105.474371: &Reels105,
}

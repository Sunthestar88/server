package deserttreasure

import (
	"github.com/slotopol/server/game"
)

// reels lengths [31, 31, 31, 31, 31], total reshuffles 28629151
// symbols: 51.213(lined) + 4.8815(scatter) = 56.094615%
// free spins 2412450, q = 0.084266, sq = 1/(1-q) = 1.092020
// free games frequency: 1/128.23
// RTP = 56.095(sym) + 0.084266*333.51(fg) = 84.197930%
var ReelsReg84 = game.Reels5x{
	{6, 10, 2, 3, 8, 7, 6, 9, 8, 10, 4, 8, 9, 10, 1, 9, 6, 8, 9, 7, 3, 4, 5, 8, 10, 7, 6, 9, 7, 10, 5},
	{7, 9, 5, 1, 6, 10, 8, 5, 9, 10, 4, 9, 8, 4, 3, 10, 7, 8, 6, 9, 10, 7, 8, 6, 9, 2, 7, 10, 8, 3, 6},
	{8, 7, 3, 9, 1, 7, 8, 3, 6, 9, 7, 10, 5, 9, 10, 5, 8, 7, 10, 6, 8, 10, 6, 4, 9, 6, 10, 4, 8, 9, 2},
	{6, 9, 4, 6, 9, 10, 7, 3, 2, 8, 3, 10, 6, 8, 1, 10, 7, 9, 8, 5, 4, 7, 9, 8, 7, 6, 10, 8, 9, 10, 5},
	{6, 9, 8, 7, 3, 10, 9, 8, 5, 7, 9, 4, 6, 10, 8, 3, 9, 8, 10, 7, 6, 10, 8, 5, 7, 6, 2, 4, 9, 10, 1},
}

// reels lengths [31, 31, 31, 31, 31], total reshuffles 28629151
// symbols: 53.022(lined) + 4.8815(scatter) = 57.903051%
// free spins 2412450, q = 0.084266, sq = 1/(1-q) = 1.092020
// free games frequency: 1/128.23
// RTP = 57.903(sym) + 0.084266*333.51(fg) = 86.006366%
var ReelsReg86 = game.Reels5x{
	{5, 6, 4, 9, 10, 8, 7, 9, 8, 7, 6, 4, 3, 6, 8, 5, 2, 9, 6, 10, 7, 5, 3, 4, 5, 1, 10, 8, 9, 7, 10},
	{9, 7, 6, 4, 8, 9, 5, 3, 10, 8, 6, 5, 2, 10, 8, 3, 7, 8, 4, 9, 7, 1, 9, 6, 5, 7, 6, 10, 5, 4, 10},
	{10, 8, 6, 10, 4, 9, 8, 7, 9, 6, 4, 3, 8, 7, 6, 5, 10, 4, 7, 5, 1, 3, 9, 8, 10, 7, 5, 6, 9, 2, 5},
	{1, 5, 6, 7, 5, 3, 2, 8, 5, 3, 9, 6, 4, 8, 7, 9, 10, 8, 4, 10, 7, 4, 6, 10, 5, 9, 6, 10, 7, 9, 8},
	{8, 4, 7, 6, 9, 4, 5, 9, 8, 5, 10, 8, 1, 3, 5, 6, 9, 7, 10, 4, 7, 10, 6, 8, 10, 3, 9, 5, 6, 2, 7},
}

// reels lengths [33, 33, 33, 33, 33], total reshuffles 39135393
// symbols: 62.096(lined) + 3.9739(scatter) = 66.069558%
// free spins 2745900, q = 0.070164, sq = 1/(1-q) = 1.075459
// free games frequency: 1/153.24
// RTP = 66.07(sym) + 0.070164*333.51(fg) = 89.469927%
var ReelsReg89 = game.Reels5x{
	{10, 9, 3, 10, 4, 6, 8, 10, 7, 9, 6, 5, 9, 10, 5, 6, 4, 8, 7, 5, 1, 3, 6, 8, 4, 7, 8, 3, 4, 7, 9, 5, 2},
	{8, 3, 5, 1, 10, 4, 9, 3, 6, 9, 7, 8, 4, 6, 5, 4, 10, 6, 7, 10, 9, 8, 4, 7, 8, 5, 3, 6, 7, 2, 10, 5, 9},
	{1, 9, 4, 8, 9, 7, 6, 9, 10, 6, 3, 5, 8, 6, 10, 7, 3, 10, 2, 8, 4, 10, 9, 5, 7, 8, 3, 5, 4, 6, 5, 4, 7},
	{5, 10, 6, 1, 8, 4, 3, 9, 6, 10, 4, 7, 6, 10, 9, 3, 5, 8, 3, 4, 9, 10, 4, 5, 7, 8, 2, 7, 9, 5, 8, 6, 7},
	{10, 4, 3, 7, 10, 9, 3, 5, 8, 7, 4, 10, 5, 4, 6, 1, 8, 6, 9, 8, 10, 6, 4, 9, 5, 7, 9, 8, 5, 3, 7, 2, 6},
}

// reels lengths [32, 32, 32, 32, 32], total reshuffles 33554432
// symbols: 60.44(lined) + 4.3967(scatter) = 64.836717%
// free spins 2576475, q = 0.076785, sq = 1/(1-q) = 1.083171
// free games frequency: 1/140.36
// RTP = 64.837(sym) + 0.076785*333.51(fg) = 90.445193%
var ReelsReg90 = game.Reels5x{
	{7, 9, 5, 10, 9, 4, 1, 7, 5, 10, 3, 6, 8, 3, 9, 10, 6, 8, 4, 3, 7, 8, 9, 10, 7, 6, 8, 5, 2, 4, 5, 6},
	{4, 5, 9, 8, 10, 4, 9, 10, 7, 9, 10, 5, 8, 7, 6, 2, 5, 7, 3, 8, 9, 5, 3, 6, 1, 4, 6, 3, 7, 8, 10, 6},
	{3, 1, 8, 3, 5, 7, 6, 5, 8, 10, 4, 5, 10, 2, 9, 8, 7, 9, 4, 5, 6, 3, 9, 7, 6, 10, 9, 6, 10, 8, 4, 7},
	{4, 7, 8, 10, 4, 9, 7, 3, 6, 8, 3, 10, 4, 5, 10, 1, 9, 3, 7, 6, 9, 10, 6, 2, 8, 5, 6, 9, 5, 7, 8, 5},
	{6, 5, 10, 7, 2, 6, 8, 9, 7, 1, 5, 4, 3, 10, 8, 5, 9, 4, 8, 10, 3, 7, 10, 5, 7, 6, 3, 8, 9, 6, 4, 9},
}

// reels lengths [31, 31, 31, 31, 31], total reshuffles 28629151
// symbols: 58.051(lined) + 4.8815(scatter) = 62.932635%
// free spins 2412450, q = 0.084266, sq = 1/(1-q) = 1.092020
// free games frequency: 1/128.23
// RTP = 62.933(sym) + 0.084266*333.51(fg) = 91.035950%
var ReelsReg91 = game.Reels5x{
	{8, 3, 4, 9, 6, 10, 9, 5, 10, 6, 8, 7, 10, 4, 3, 5, 7, 10, 1, 8, 10, 7, 9, 8, 5, 9, 2, 5, 3, 6, 4},
	{1, 8, 9, 3, 8, 4, 5, 10, 3, 8, 5, 9, 6, 10, 8, 6, 9, 5, 10, 7, 2, 5, 7, 10, 4, 3, 10, 9, 7, 6, 4},
	{9, 8, 4, 9, 5, 3, 4, 9, 10, 3, 2, 7, 10, 8, 7, 6, 10, 4, 6, 5, 9, 1, 8, 3, 5, 10, 6, 8, 5, 7, 10},
	{9, 4, 8, 9, 5, 8, 10, 6, 3, 1, 10, 4, 2, 10, 3, 7, 5, 4, 8, 5, 9, 6, 7, 10, 9, 5, 3, 7, 6, 8, 10},
	{5, 10, 4, 1, 9, 10, 5, 9, 7, 5, 9, 4, 7, 10, 3, 8, 10, 5, 8, 7, 3, 4, 8, 3, 6, 2, 9, 6, 10, 8, 6},
}

// reels lengths [31, 31, 30, 31, 31], total reshuffles 27705630
// symbols: 58.344(lined) + 4.9888(scatter) = 63.333012%
// free spins 2380185, q = 0.08591, sq = 1/(1-q) = 1.093984
// free games frequency: 1/125.84
// RTP = 63.333(sym) + 0.08591*333.51(fg) = 91.984711%
var ReelsReg92 = game.Reels5x{
	{8, 3, 4, 9, 6, 10, 9, 5, 10, 6, 8, 7, 10, 4, 3, 5, 7, 10, 1, 8, 10, 7, 9, 8, 5, 9, 2, 5, 3, 6, 4},
	{1, 8, 9, 3, 8, 4, 5, 10, 3, 8, 5, 9, 6, 10, 8, 6, 9, 5, 10, 7, 2, 5, 7, 10, 4, 3, 10, 9, 7, 6, 4},
	{7, 8, 3, 6, 9, 4, 2, 5, 10, 9, 8, 10, 4, 8, 7, 10, 6, 1, 5, 10, 9, 3, 6, 9, 8, 7, 5, 4, 10, 3},
	{9, 4, 8, 9, 5, 8, 10, 6, 3, 1, 10, 4, 2, 10, 3, 7, 5, 4, 8, 5, 9, 6, 7, 10, 9, 5, 3, 7, 6, 8, 10},
	{5, 10, 4, 1, 9, 10, 5, 9, 7, 5, 9, 4, 7, 10, 3, 8, 10, 5, 8, 7, 3, 4, 8, 3, 6, 2, 9, 6, 10, 8, 6},
}

// reels lengths [30, 31, 31, 31, 30], total reshuffles 26811900
// symbols: 58.78(lined) + 5.0983(scatter) = 63.878673%
// free spins 2348190, q = 0.08758, sq = 1/(1-q) = 1.095987
// free games frequency: 1/123.51
// RTP = 63.879(sym) + 0.08758*333.51(fg) = 93.087448%
var ReelsReg93 = game.Reels5x{
	{4, 5, 9, 3, 8, 10, 3, 8, 9, 10, 8, 7, 5, 4, 1, 6, 10, 4, 9, 10, 8, 7, 6, 5, 7, 9, 6, 2, 3, 10},
	{1, 8, 9, 3, 8, 4, 5, 10, 3, 8, 5, 9, 6, 10, 8, 6, 9, 5, 10, 7, 2, 5, 7, 10, 4, 3, 10, 9, 7, 6, 4},
	{9, 8, 4, 9, 5, 3, 4, 9, 10, 3, 2, 7, 10, 8, 7, 6, 10, 4, 6, 5, 9, 1, 8, 3, 5, 10, 6, 8, 5, 7, 10},
	{9, 4, 8, 9, 5, 8, 10, 6, 3, 1, 10, 4, 2, 10, 3, 7, 5, 4, 8, 5, 9, 6, 7, 10, 9, 5, 3, 7, 6, 8, 10},
	{6, 5, 10, 4, 5, 9, 6, 7, 3, 10, 9, 7, 8, 3, 6, 5, 8, 2, 9, 4, 7, 10, 8, 4, 9, 10, 3, 8, 10, 1},
}

// reels lengths [29, 29, 29, 29, 29], total reshuffles 20511149
// symbols: 53.291(lined) + 6.0868(scatter) = 59.378127%
// free spins 2100600, q = 0.10241, sq = 1/(1-q) = 1.114098
// free games frequency: 1/106.11
// RTP = 59.378(sym) + 0.10241*333.51(fg) = 93.533660%
var ReelsReg94 = game.Reels5x{
	{7, 4, 5, 6, 2, 9, 3, 8, 10, 9, 4, 1, 8, 5, 7, 9, 5, 4, 6, 10, 8, 7, 10, 5, 9, 6, 3, 10, 8},
	{3, 10, 8, 5, 9, 4, 3, 8, 7, 9, 6, 10, 8, 9, 4, 10, 9, 6, 5, 7, 1, 8, 7, 5, 2, 6, 10, 5, 4},
	{4, 10, 8, 9, 4, 3, 2, 5, 9, 1, 10, 7, 8, 4, 6, 8, 10, 5, 8, 9, 6, 3, 10, 9, 7, 5, 6, 7, 5},
	{10, 3, 5, 9, 8, 1, 9, 10, 2, 7, 8, 9, 10, 7, 6, 9, 4, 3, 5, 6, 7, 10, 8, 6, 5, 4, 8, 5, 4},
	{7, 10, 5, 7, 4, 2, 5, 10, 6, 9, 4, 8, 7, 10, 4, 8, 3, 9, 5, 6, 8, 10, 5, 1, 9, 8, 6, 3, 9},
}

// reels lengths [31, 31, 31, 31, 31], total reshuffles 28629151
// symbols: 61.853(lined) + 4.8815(scatter) = 66.734284%
// free spins 2412450, q = 0.084266, sq = 1/(1-q) = 1.092020
// free games frequency: 1/128.23
// RTP = 66.734(sym) + 0.084266*333.51(fg) = 94.837599%
var ReelsReg95 = game.Reels5x{
	{9, 10, 4, 7, 1, 3, 7, 10, 6, 4, 5, 8, 9, 3, 6, 10, 2, 8, 7, 10, 8, 7, 5, 9, 3, 5, 9, 4, 8, 6, 5},
	{6, 7, 8, 6, 5, 1, 10, 9, 4, 10, 8, 9, 7, 8, 10, 3, 5, 10, 9, 8, 4, 5, 6, 4, 7, 5, 3, 7, 2, 3, 9},
	{7, 3, 5, 9, 10, 8, 4, 7, 9, 10, 2, 7, 6, 9, 7, 4, 3, 10, 8, 6, 1, 5, 10, 6, 5, 8, 3, 9, 8, 5, 4},
	{10, 9, 8, 6, 7, 4, 1, 8, 7, 3, 8, 10, 9, 7, 8, 6, 9, 4, 5, 10, 7, 5, 3, 2, 10, 5, 6, 9, 3, 5, 4},
	{8, 10, 5, 7, 6, 10, 5, 3, 4, 8, 9, 7, 5, 3, 9, 4, 1, 7, 9, 8, 7, 6, 8, 5, 9, 3, 10, 2, 4, 10, 6},
}

// reels lengths [32, 32, 32, 32, 32], total reshuffles 33554432
// symbols: 65.729(lined) + 4.3967(scatter) = 70.125782%
// free spins 2576475, q = 0.076785, sq = 1/(1-q) = 1.083171
// free games frequency: 1/140.36
// RTP = 70.126(sym) + 0.076785*333.51(fg) = 95.734259%
var ReelsReg96 = game.Reels5x{
	{3, 1, 4, 10, 9, 8, 5, 7, 3, 6, 4, 8, 7, 6, 9, 7, 10, 3, 9, 2, 6, 8, 4, 5, 7, 6, 9, 10, 4, 5, 10, 8},
	{9, 8, 3, 9, 7, 4, 9, 2, 4, 5, 8, 7, 10, 6, 8, 7, 3, 5, 6, 3, 10, 8, 1, 4, 10, 9, 6, 4, 7, 6, 10, 5},
	{10, 4, 3, 9, 8, 6, 5, 7, 8, 6, 7, 3, 8, 4, 10, 5, 7, 2, 8, 6, 10, 9, 6, 4, 3, 9, 10, 1, 7, 9, 4, 5},
	{8, 5, 9, 3, 6, 10, 3, 7, 6, 1, 9, 4, 5, 7, 8, 4, 6, 10, 9, 5, 7, 8, 9, 4, 10, 2, 3, 8, 6, 4, 7, 10},
	{7, 3, 5, 8, 9, 2, 8, 7, 10, 8, 1, 5, 10, 6, 4, 3, 7, 4, 10, 6, 9, 8, 4, 6, 3, 9, 10, 6, 5, 7, 9, 4},
}

// reels lengths [30, 30, 30, 30, 30], total reshuffles 24300000
// symbols: 60.784(lined) + 5.44(scatter) = 66.223951%
// free spins 2253825, q = 0.09275, sq = 1/(1-q) = 1.102232
// free games frequency: 1/116.82
// RTP = 66.224(sym) + 0.09275*333.51(fg) = 97.156921%
var ReelsReg97 = game.Reels5x{
	{4, 5, 9, 3, 8, 10, 3, 8, 9, 10, 8, 7, 5, 4, 1, 6, 10, 4, 9, 10, 8, 7, 6, 5, 7, 9, 6, 2, 3, 10},
	{8, 4, 10, 8, 5, 10, 3, 9, 10, 1, 9, 4, 2, 10, 8, 9, 7, 10, 6, 4, 8, 5, 6, 3, 9, 7, 3, 5, 7, 6},
	{7, 8, 3, 6, 9, 4, 2, 5, 10, 9, 8, 10, 4, 8, 7, 10, 6, 1, 5, 10, 9, 3, 6, 9, 8, 7, 5, 4, 10, 3},
	{7, 9, 3, 7, 8, 9, 10, 3, 6, 7, 3, 5, 6, 10, 9, 8, 4, 10, 9, 4, 10, 1, 8, 5, 10, 4, 6, 8, 5, 2},
	{6, 5, 10, 4, 5, 9, 6, 7, 3, 10, 9, 7, 8, 3, 6, 5, 8, 2, 9, 4, 7, 10, 8, 4, 9, 10, 3, 8, 10, 1},
}

// reels lengths [31, 31, 31, 31, 31], total reshuffles 28629151
// symbols: 66.063(lined) + 4.8815(scatter) = 70.944790%
// free spins 2412450, q = 0.084266, sq = 1/(1-q) = 1.092020
// free games frequency: 1/128.23
// RTP = 70.945(sym) + 0.084266*333.51(fg) = 99.048105%
var ReelsReg99 = game.Reels5x{
	{3, 9, 4, 3, 8, 4, 10, 1, 9, 5, 3, 4, 10, 2, 6, 7, 10, 6, 8, 5, 10, 9, 7, 8, 4, 5, 9, 8, 5, 6, 7},
	{10, 5, 8, 3, 6, 5, 10, 4, 9, 7, 3, 1, 8, 10, 3, 4, 6, 5, 4, 8, 6, 7, 8, 4, 9, 10, 5, 9, 2, 7, 9},
	{8, 5, 3, 8, 7, 10, 8, 5, 9, 6, 4, 5, 6, 4, 10, 3, 9, 4, 1, 9, 6, 7, 8, 10, 9, 7, 4, 3, 10, 5, 2},
	{10, 9, 5, 10, 4, 8, 7, 9, 6, 8, 7, 10, 3, 9, 5, 4, 6, 3, 9, 5, 7, 4, 5, 2, 4, 8, 10, 3, 6, 8, 1},
	{3, 6, 5, 10, 4, 5, 2, 7, 8, 4, 3, 10, 8, 9, 4, 1, 9, 5, 7, 3, 8, 9, 7, 6, 4, 8, 6, 10, 9, 5, 10},
}

// reels lengths [30, 30, 30, 30, 30], total reshuffles 24300000
// symbols: 63.194(lined) + 5.44(scatter) = 68.633992%
// free spins 2253825, q = 0.09275, sq = 1/(1-q) = 1.102232
// free games frequency: 1/116.82
// RTP = 68.634(sym) + 0.09275*333.51(fg) = 99.566962%
var ReelsReg100 = game.Reels5x{
	{5, 9, 8, 7, 6, 5, 10, 3, 5, 4, 8, 3, 2, 10, 8, 3, 9, 7, 10, 9, 6, 7, 10, 5, 4, 8, 6, 9, 4, 1},
	{6, 4, 2, 10, 5, 7, 3, 9, 5, 8, 9, 3, 6, 5, 7, 10, 8, 1, 6, 9, 10, 5, 4, 8, 10, 3, 4, 8, 9, 7},
	{7, 6, 8, 9, 5, 2, 6, 7, 10, 8, 3, 10, 5, 4, 8, 9, 4, 3, 10, 5, 7, 8, 10, 9, 5, 3, 9, 1, 6, 4},
	{6, 8, 4, 5, 6, 3, 5, 7, 9, 8, 10, 1, 5, 10, 9, 5, 8, 9, 10, 6, 3, 4, 8, 7, 9, 4, 10, 3, 2, 7},
	{9, 3, 8, 5, 7, 8, 9, 4, 10, 5, 9, 10, 2, 8, 4, 6, 3, 5, 1, 10, 3, 7, 6, 4, 10, 9, 8, 7, 5, 6},
}

// reels lengths [27, 27, 27, 27, 27], total reshuffles 14348907
// symbols: 61.778(lined) + 7.7224(scatter) = 69.500318%
// free spins 1810350, q = 0.12617, sq = 1/(1-q) = 1.144383
// free games frequency: 1/86.709
// RTP = 69.5(sym) + 0.12617*333.51(fg) = 111.577963%
var ReelsReg112 = game.Reels5x{
	{9, 7, 10, 5, 8, 6, 9, 8, 10, 1, 3, 5, 4, 7, 10, 4, 2, 8, 6, 10, 9, 6, 5, 4, 3, 7, 5},
	{7, 10, 6, 8, 10, 5, 8, 9, 6, 5, 4, 1, 7, 8, 10, 5, 3, 9, 4, 10, 6, 9, 7, 2, 3, 4, 5},
	{10, 9, 4, 10, 9, 4, 8, 3, 1, 5, 6, 7, 2, 8, 5, 10, 8, 7, 6, 4, 10, 7, 5, 9, 6, 3, 5},
	{10, 1, 4, 5, 6, 3, 8, 5, 6, 7, 10, 5, 3, 10, 6, 9, 8, 7, 9, 5, 7, 2, 4, 9, 10, 8, 4},
	{4, 10, 5, 3, 10, 4, 9, 5, 2, 8, 6, 5, 8, 1, 4, 9, 3, 7, 5, 6, 9, 7, 10, 8, 7, 10, 6},
}

// reels lengths [26, 26, 26, 26, 26], total reshuffles 11881376
// symbols: 260.25(lined) + 26.288(scatter) = 286.539034%
// free spins 1673325, q = 0.14084, sq = 1/(1-q) = 1.163922
// free games frequency: 1/77.968
// RTP = sq*rtp(sym) = 1.1639*286.54 = 333.509110%
var ReelsBon = game.Reels5x{
	{8, 2, 4, 5, 7, 10, 3, 9, 10, 6, 8, 4, 6, 9, 3, 7, 1, 5, 10, 9, 7, 4, 8, 6, 5, 3},
	{2, 10, 4, 5, 8, 7, 9, 3, 5, 1, 10, 3, 7, 4, 9, 8, 6, 4, 10, 5, 7, 6, 3, 9, 8, 6},
	{3, 7, 1, 3, 4, 5, 6, 3, 9, 10, 6, 2, 10, 5, 4, 7, 8, 4, 9, 8, 5, 9, 6, 10, 7, 8},
	{9, 4, 10, 7, 3, 4, 10, 6, 2, 8, 5, 3, 7, 6, 8, 1, 9, 4, 6, 8, 3, 5, 9, 7, 10, 5},
	{3, 9, 10, 6, 9, 3, 6, 9, 10, 7, 5, 8, 7, 1, 4, 7, 6, 2, 3, 5, 8, 4, 5, 8, 4, 10},
}

// Map with available reels.
var ReelsMap = map[string]*game.Reels5x{
	"84":  &ReelsReg84,
	"86":  &ReelsReg86,
	"89":  &ReelsReg89,
	"90":  &ReelsReg90,
	"91":  &ReelsReg91,
	"92":  &ReelsReg92,
	"93":  &ReelsReg93,
	"94":  &ReelsReg94,
	"95":  &ReelsReg95,
	"96":  &ReelsReg96,
	"97":  &ReelsReg97,
	"99":  &ReelsReg99,
	"100": &ReelsReg100,
	"112": &ReelsReg112,
	"bon": &ReelsBon,
}

// Lined payment.
var LinePay = [10][5]float64{
	{0, 8, 80, 800, 8000}, //  1 wild
	{0, 0, 0, 0, 0},       //  2 scatter
	{0, 5, 40, 250, 1000}, //  3 shield
	{0, 0, 20, 75, 500},   //  4 swords
	{0, 0, 0, 50, 250},    //  5 lamp
	{0, 2, 10, 30, 150},   //  6 ligature1
	{0, 2, 10, 30, 150},   //  7 ligature2
	{0, 0, 5, 15, 75},     //  8 ligature3
	{0, 0, 5, 15, 75},     //  9 ligature4
	{0, 0, 0, 10, 50},     // 10 ligature5
}

// Scatters payment.
var ScatPay = [5]float64{0, 0, 4, 40, 400} // scatter

// Scatter freespins table
var ScatFreespin = [5]int{0, 0, 10, 25, 50} // scatter

type Game struct {
	game.Slot5x3 `yaml:",inline"`
	// free spin number
	FS int `json:"fs,omitempty" yaml:"fs,omitempty" xml:"fs,omitempty"`
}

func NewGame(rd string) *Game {
	return &Game{
		Slot5x3: game.Slot5x3{
			RD:  rd,
			SBL: game.MakeBitNum(9),
			Bet: 1,
		},
		FS: 0,
	}
}

const wild, scat = 1, 2

var bl = game.Lineset5x{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{1, 1, 2, 3, 3}, // 6
	{3, 3, 2, 1, 1}, // 7
	{2, 1, 2, 3, 2}, // 8
	{2, 3, 2, 1, 2}, // 9
}

func (g *Game) Scanner(screen game.Screen, ws *game.WinScan) {
	g.ScanLined(screen, ws)
	g.ScanScatters(screen, ws)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen game.Screen, ws *game.WinScan) {
	for li := g.SBL.Next(0); li != 0; li = g.SBL.Next(li) {
		var line = bl.Line(li)

		var numw, numl = 0, 5
		var syml game.Sym
		for x := 1; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx == wild {
				if syml == 0 {
					numw = x
				}
			} else if syml == 0 && sx != scat {
				syml = sx
			} else if sx != syml {
				numl = x - 1
				break
			}
		}

		var payw, payl float64
		if numw > 0 {
			payw = LinePay[wild-1][numw-1]
		}
		if numl > 0 && syml > 0 {
			payl = LinePay[syml-1][numl-1]
		}
		if payl > payw {
			var mm float64 = 1 // mult mode
			if g.FS > 0 {
				mm = 3
			}
			ws.Wins = append(ws.Wins, game.WinItem{
				Pay:  g.Bet * payl,
				Mult: mm,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		} else if payw > 0 {
			var mm float64 = 1 // mult mode
			if g.FS > 0 {
				mm = 3
			}
			ws.Wins = append(ws.Wins, game.WinItem{
				Pay:  g.Bet * payw,
				Mult: mm,
				Sym:  wild,
				Num:  numw,
				Line: li,
				XY:   line.CopyL(numw),
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen game.Screen, ws *game.WinScan) {
	if count := screen.ScatNum(scat); count >= 2 {
		var mm float64 = 1 // mult mode
		if g.FS > 0 {
			mm = 3
		}
		var pay, fs = ScatPay[count-1], ScatFreespin[count-1]
		ws.Wins = append(ws.Wins, game.WinItem{
			Pay:  g.Bet * float64(g.SBL.Num()) * pay,
			Mult: mm,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
			Free: fs,
		})
	}
}

func (g *Game) Spin(screen game.Screen) {
	if g.FS == 0 {
		screen.Spin(ReelsMap[g.RD])
	} else {
		screen.Spin(&ReelsBon)
	}
}

func (g *Game) Apply(screen game.Screen, sw *game.WinScan) {
	if g.FS > 0 {
		g.Gain += sw.Gain()
	} else {
		g.Gain = sw.Gain()
	}

	if g.FS > 0 {
		g.FS--
	}
	for _, wi := range sw.Wins {
		if wi.Free > 0 {
			g.FS += wi.Free
		}
	}
}

func (g *Game) FreeSpins() int {
	return g.FS
}

func (g *Game) SetLines(sbl game.Bitset) error {
	var mask game.Bitset = (1<<len(bl) - 1) << 1
	if sbl == 0 {
		return game.ErrNoLineset
	}
	if sbl&^mask != 0 {
		return game.ErrLinesetOut
	}
	if g.FreeSpins() > 0 {
		return game.ErrNoFeature
	}
	g.SBL = sbl
	return nil
}

func (g *Game) SetReels(rd string) error {
	if _, ok := ReelsMap[rd]; !ok {
		return game.ErrNoReels
	}
	g.RD = rd
	return nil
}

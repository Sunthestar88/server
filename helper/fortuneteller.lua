local path = arg[0]:match("(.*[/\\])")
dofile(path.."reelgen.lua")

local symset = {
	1, --  1 wild
	1, --  2 cat
	1, --  3 bonus
	3, --  4 girl
	3, --  5 hand
	3, --  6 zodiac
	3, --  7 candle
	3, --  8 ace
	3, --  9 king
	3, -- 10 queen
	4, -- 11 jack
	4, -- 12 ten
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,
	{ 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, }, --  1
	{ 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, }, --  2
	{ 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, }, --  3
	{ 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, }, --  4
	{ 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, }, --  5
	{ 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, }, --  6
	{ 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, }, --  7
	{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, }, --  8
	{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, }, --  9
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, }, -- 10
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, }, -- 11
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, }, -- 12
}

math.randomseed(os.time())
local reel = MakeReel(symset)
print("reel length: " .. #reel)
ShuffleReel(reel)
local iter = CorrectReel(reel, neighbours)
RrintReel(reel, iter)

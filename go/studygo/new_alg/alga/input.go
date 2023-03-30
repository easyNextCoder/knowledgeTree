package alga

func (self *AlgX) Input() {

	//案例1
	//self.hand_cards = Deck{Card{1, 1}, Card{1, 1}, Card{2, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{7, 1},
	//	Card{8, 1}, Card{9, 1}, Card{10, 1}, Card{11, 1}, Card{12, 1}, Card{13, 1}}
	//self.jokerLeft = 2

	//案例2
	//self.hand_cards = Deck{Card{2, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{7, 1}, Card{6, 1}, Card{7, 1}, Card{8, 1}, Card{10, 1}, Card{11, 1},
	//	Card{4, 2}, Card{5, 2}, Card{6, 2},
	//	Card{4, 3}, Card{6, 3},
	//	Card{6, 4}}
	//self.jokerLeft = 2 //最大分值114分
	// 2  3  4  5  6 6 7 7  8  10 11
	//       4  5  6
	//       4     6
	//             6

	//案例3
	//self.hand_cards = Deck{Card{1, 1}, Card{2, 1}, Card{3, 1},
	//	Card{3, 1}, Card{4, 1}, Card{5, 1}}
	//self.jokerLeft = 2
	//1 2 3 3 4 5//目前跑不过老算法 26us -> 24us

	//案例4
	//self.hand_cards = Deck{Card{2, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{7, 1}, {6, 1}, {7, 1}, Card{8, 1},
	//	Card{4, 2}, Card{5, 2}, Card{6, 2},
	//	Card{4, 3}, Card{6, 3},
	//	Card{6, 4}}
	//self.jokerLeft = 2
	// 2  3  4  5  6 6 7 7  8
	//       4  5  6
	//       4     6
	//             6

	//案例2
	//self.hand_cards = Deck{Card{2, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{5, 1}, Card{6, 1}, Card{7, 1}, Card{8, 1},
	//	Card{4, 2}, Card{5, 2}, Card{6, 2},
	//	Card{4, 3}, Card{6, 3},
	//	Card{6, 4}}

	//案例1: 24张牌 最大复杂度
	//self.hand_cards = Deck{Card{2, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{7, 1},
	//	Card{2, 2}, Card{3, 2}, Card{4, 2}, Card{5, 2}, Card{6, 2}, Card{7, 2},
	//	Card{2, 3}, Card{3, 3}, Card{4, 3}, Card{5, 3}, Card{6, 3}, Card{7, 3},
	//	Card{2, 4}, Card{3, 4}, Card{4, 4}, Card{5, 4}, Card{6, 4}, Card{7, 4},
	//} //2joker 用时12s
	//self.jokerLeft = 2
	//"[18,19,20,21,22,23, 34,35,36,37,38,39, 50,51,52,53,54,55, 66,67,68,69,70,71]\n",

	//案例2
	//self.hand_cards = Deck{Card{2, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{7, 1}}
	//self.jokerLeft = 2

	//案例3
	//self.hand_cards = Deck{Card{1, 1}, Card{1, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{7, 1},
	//	Card{8, 1}, Card{9, 1}, Card{10, 1}, Card{11, 1}, Card{12, 1}, Card{13, 1}, Card{12, 1}, Card{13, 1}}
	//
	//self.jokerLeft = 1
	////案例3 用于验证同花色有双A的情况下能否正确求解
	//self.hand_cards = Deck{Card{1, 1}, Card{1, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{7, 1},
	//	Card{8, 1}, Card{9, 1}, Card{10, 1}, Card{11, 1}, Card{12, 1}, Card{13, 1},
	//	Card{4, 1}, Card{5, 1}, Card{6, 1},
	//	Card{12, 1}, Card{13, 1},
	//}

	//
	//self.hand_cards = Deck{Card{1, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{7, 1},
	//	Card{8, 1}, Card{9, 1}, Card{10, 1}, Card{11, 1}, Card{12, 1}, Card{13, 1},
	//	Card{4, 1}, Card{5, 1}, Card{6, 1},
	//	Card{12, 1}, Card{13, 1},
	//} //这个案例会优先组成Q K A 这样是符合最高分的要求的
	////
	//self.hand_cards = Deck{Card{1, 1}, Card{2, 1}, Card{3, 1}, Card{4, 1},
	//	Card{1, 1}, Card{2, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}}
	//
	//self.hand_cards = Deck{Card{1, 1}, Card{2, 1}, Card{3, 1},
	//	Card{2, 1}, Card{3, 1}, Card{4, 1}}
	//self.jokerLeft = 2
	//
	//self.hand_cards = Deck{Card{1, 1}, Card{2, 1}, Card{3, 1},
	//	Card{2, 1}, Card{3, 1}, Card{4, 1}}
	//self.jokerLeft = 2
	//

	//self.hand_cards = Deck{Card{1, 1}, Card{4, 1}, Card{6, 1}, Card{9, 1},
	//	Card{6, 2}, Card{9, 2},
	//}
	//self.jokerLeft = 2
	//1 4 5   9
	//      6 9
	//joker joker
	//结果:{6 1} {6 2} {6 3}] [{9 1} {9 2} {9 3}]] 45}

	//self.hand_cards = Deck{Card{1, 1}, Card{2, 1}, Card{3, 1},
	//	Card{1, 2}, Card{2, 2}, Card{3, 2}, Card{4, 2}}
	//self.jokerLeft = 2
	//1 2 3
	//1 2 3 4
	//joker joker

	//案例
	//self.hand_cards = Deck{Card{1, 1}, Card{10, 1}, Card{11, 1}, Card{12, 1}, Card{13, 1},
	//	Card{1, 1}, Card{10, 1}, Card{11, 1}, Card{12, 1}, Card{13, 1}}
	//self.jokerLeft = 2 //{[[{5 1} {6 1} {7 1} {8 1} {9 1} {10 1} {11 1} {12 1} {13 1}] [{5 1} {6 1} {7 1} {8 1} {9 1} {10 1} {11 1} {12 1} {13 1} {14 1}] [{1 1} {1 2} {1 3}]] 190} costTime 281us
	//"[21,22,23,24,25,26,27,28,29,17, 37,38,39,40,41,42,43,44,45,33, 14, 14]\n", //老算法需要1.6s

	//案例4
	//self.hand_cards = Deck{Card{1, 1}, Card{2, 1}, Card{2, 2}, Card{2, 3}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{7, 1}}
	//self.jokerLeft = 2

	//案例5
	//self.hand_cards = Deck{Card{1, 1}, Card{2, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{6, 2}, Card{6, 3}, Card{7, 1}}
	//self.jokerLeft = 2
	//1 2 3 4 5 6 7
	//          6
	//          6
	//输出:&{[[{6 2} {6 3} {6 1}] [{1 1} {2 1} {3 1} {4 1} {5 1}] [{7 1} {8 1} {9 1}]] 66} 88us

	//案例6
	//self.hand_cards = Deck{Card{1, 1}, Card{2, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{6, 1}, Card{6, 2}, Card{6, 3}, Card{6, 2}, Card{6, 3}, Card{7, 1}} //{[[{6 2} {6 3} {6 1}] [{6 1} {6 2} {6 3}] [{1 1} {2 1} {3 1} {4 1} {5 1} {6 1} {7 1}]] 73}
	//start := time.Now()

	//案例
	//♠5 ♠7 ♠9 ♠9 ♥5 ♥8 ♥9 ♥9 ♣7 ♣9 ♣9 ♣10 ♦5 ♦6 ♦K
	self.hand_cards = Deck{Card{5, 1}, Card{7, 1}, Card{9, 1}, Card{9, 1},
		Card{5, 2}, Card{8, 2}, Card{9, 2}, Card{9, 2},
		Card{7, 3}, Card{9, 3}, Card{9, 3}, Card{10, 3},
		Card{5, 4}, Card{6, 4}, Card{13, 4}}

	//self.jokerLeft = 2
	//5   7   9 9
	//5     8 9 9
	//    7   9 9 10
	//5 6                13
	//结果:{[[{5 1} {5 2} {5 4}] [{9 3} {9 1} {9 2}] [{9 3} {9 1} {9 2}] [{7 3} {8 3} {9 3} {10 3}]] 103} costTime: 4458

	//案例：用来测试最长和最大分
	//self.hand_cards = Deck{Card{2, 1}, Card{3, 1}, Card{4, 1}, Card{5, 1}, Card{5, 2}, Card{5, 3}}
	//self.jokerLeft = 0

	//self.hand_cards = Deck{Card{2, 1}, Card{3, 1}, Card{5, 1}, Card{6, 1}, Card{10, 3}, Card{11, 3}, Card{13, 3}}
	//self.jokerLeft = 1
}

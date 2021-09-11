package problem1_100

type hash map[byte]bool

func isValidSudoku(board [][]byte) bool {
	var rowHash = make([]hash, 9)
	var colHash = make([]hash, 9)
	var boxHash = make([]hash, 9)
	// 初始化
	for i := 0; i < 9; i++ {
		rowHash[i] = make(map[byte]bool)
		colHash[i] = make(map[byte]bool)
		boxHash[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var ch = board[i][j]

			if ch == '.' {
				continue
			}

			if _, ok := rowHash[i][ch]; ok {
				return false
			} else {
				rowHash[i][ch] = true
			}

			if _, ok := colHash[j][ch]; ok {
				return false
			} else {
				colHash[j][ch] = true
			}

			var boxIndex = 3*(i/3) + j/3
			if _, ok := boxHash[boxIndex][ch]; ok {
				return false
			} else {
				boxHash[boxIndex][ch] = true
			}
		}
	}

	return true
}

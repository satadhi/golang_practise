package chessboard

// File stores whether a square is occupied.
type File []bool

// Chessboard stores files A-H.
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0

	// Check if file exists
	squares, ok := cb[file]
	if !ok {
		return 0
	}

	for _, occupied := range squares {
		if occupied {
			count++
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	// Valid ranks are 1-8
	if rank < 1 || rank > 8 {
		return 0
	}

	count := 0

	// Convert to slice index (0-7)
	index := rank - 1

	for _, file := range cb {
		if file[index] {
			count++
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0

	for range cb {
		count += 8
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

	for _, file := range cb {
		for _, occupied := range file {
			if occupied {
				count++
			}
		}
	}

	return count
}
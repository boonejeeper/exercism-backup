package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    f := cb[file]
    var count int
	for _, x := range f {
        if x {
        	count++
		}
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    var count int
	for _, file := range cb {
        for r, occupied := range file {
            if rank - 1 == r && occupied {
                count++
            }
        }
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count int
    for _, file := range cb {
        count += len(file)
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count int
	for _, file := range cb {
        for _, occupied := range file {
            if occupied {
                count++
            }
        }
    }
    return count
}

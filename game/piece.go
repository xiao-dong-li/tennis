package game

// Pieces is the set of all the possible pieces.
var Pieces map[BlockType]*Piece

type Piece struct {
	blockType BlockType
	blocks    [][]bool
}

func init() {
	const (
		t = true
		f = false
	)

	Pieces = map[BlockType]*Piece{
		BlockTypeI: {
			blockType: BlockTypeI,
			blocks: [][]bool{
				{f, f, f, f},
				{t, t, t, t},
				{f, f, f, f},
				{f, f, f, f},
			},
		},
		BlockTypeJ: {
			blockType: BlockTypeJ,
			blocks: [][]bool{
				{f, t, f},
				{f, t, f},
				{t, t, f},
			},
		},
		BlockTypeL: {
			blockType: BlockTypeL,
			blocks: [][]bool{
				{f, t, f},
				{f, t, f},
				{f, t, t},
			},
		},
		BlockTypeO: {
			blockType: BlockTypeO,
			blocks: [][]bool{
				{f, f},
				{f, f},
			},
		},
		BlockTypeS: {
			blockType: BlockTypeS,
			blocks: [][]bool{
				{f, t, t},
				{t, t, f},
				{f, f, f},
			},
		},
		BlockTypeT: {
			blockType: BlockTypeT,
			blocks: [][]bool{
				{f, t, f},
				{t, t, t},
				{f, f, f},
			},
		},
		BlockTypeZ: {
			blockType: BlockTypeZ,
			blocks: [][]bool{
				{t, t, f},
				{f, t, t},
				{f, f, f},
			},
		},
	}
}

func rotate(matrix [][]bool, clockwise bool) {
	n := len(matrix)
	tmp := make([][]bool, n)
	for i := range tmp {
		tmp[i] = make([]bool, n)
	}

	for i, row := range matrix {
		for j, v := range row {
			if clockwise {
				tmp[j][n-1-i] = v
			} else {
				tmp[n-1-j][i] = v
			}
		}
	}

	copy(matrix, tmp)
}

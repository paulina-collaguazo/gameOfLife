package patterns

func GetPatternFromOption(patternOption int) [][]int {
	switch patternOption {
	case 1:
		patternBlock := [][]int{
			{2, 1, 1}, {2, 2, 1},
			{3, 1, 1}, {3, 2, 1},
		}
		return patternBlock

	case 2:
		patternBeehive := [][]int{
			{1, 2, 1}, {1, 3, 1}, {2, 1, 1},
			{2, 4, 1}, {3, 2, 1}, {3, 3, 1},
		}
		return patternBeehive

	case 3:
		patternLoaf := [][]int{
			{1, 2, 1}, {1, 3, 1}, {2, 1, 1}, {2, 4, 1},
			{3, 2, 1}, {3, 4, 1}, {4, 3, 1},
		}
		return patternLoaf

	case 4:
		patternBoat := [][]int{
			{2, 1, 1}, {2, 2, 1},
			{3, 1, 1}, {3, 2, 1},
		}
		return patternBoat
	case 5:
		patternTub := [][]int{
			{1, 2, 1}, {2, 1, 1},
			{2, 3, 1}, {3, 2, 1},
		}
		return patternTub
	default:
		return [][]int{}
	}

}

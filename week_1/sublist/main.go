package sublist

// Relation is the comparison between lists
type Relation string

// Possible relations
const (
	RelationEqual     Relation = "equal"
	RelationSublist   Relation = "sublist"
	RelationSuperlist Relation = "superlist"
	RelationUnequal   Relation = "unequal"
)

// Sublist checks difference of two lists and
// returns equal, sublist, superlist or unequal according
// to their relation to each other.
func Sublist(l1, l2 []int) Relation {
	if isEqual(l1, l2) {
		return RelationEqual
	} else if contains(l1, l2) {
		return RelationSuperlist
	} else if contains(l2, l1) {
		return RelationSublist
	}
	return RelationUnequal
}

func isEqual(l1, l2 []int) bool {
	// Тут має бути рішення
	// написавши код - необхідно запустити тести
	// Ці коментарі можна видаляти
	// !ВАЖЛИВО - не забудьте виправити return
	return false
}

func contains(l1, l2 []int) bool {
	// Тут має бути рішення
	// написавши код - необхідно запустити тести
	// Ці коментарі можна видаляти
	// !ВАЖЛИВО - не забудьте виправити return
	return false
}

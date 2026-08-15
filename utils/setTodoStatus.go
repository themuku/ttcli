package utils

func SetTodoStatus(s int) string {
	switch s {
	case 0:
		return "Todo"
	case 1:
		return "In progress"
	case 2:
		return "Done"
	default:
		panic("Invalid status")
	}
}

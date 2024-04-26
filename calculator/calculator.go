package calculator



type Operation interface {
	Operate(a, b int) int
}

type AdditionOperation struct{}

func (op AdditionOperation) Operate(a, b int) int {
	return a + b
}

type MultiplicationOperation struct{}

func (op MultiplicationOperation) Operate(a, b int) int {
	return a * b
}

func GetOperation(opType string) Operation {
	switch opType {
	case "addition":
		return AdditionOperation{}
	case "multiplication":
		return MultiplicationOperation{}
	default:
		panic("Unknown operation type")
	}
}


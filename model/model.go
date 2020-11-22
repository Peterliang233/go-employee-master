package model

type WorkMan struct {
	Number     string
	Name       string
	Profession string
	Task       string
}

var Worker []WorkMan

var Num int

func (w WorkMan) AddEmployee() {
	Worker=append(Worker, w)
	Num++
}

func DeleteEmployee(i int) {
	if i < len(Worker) - 1 {
		Worker = append(Worker[:i], Worker[i+1:]...)
	}else{
		Worker= Worker[:len(Worker)-1]
	}
	Num--
}
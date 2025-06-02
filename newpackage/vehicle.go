package newpackage

type Vehicle struct {
	rno int
}

func (v *Vehicle) GetRno() int {
	return v.rno
}
func (v *Vehicle) SetRno(rno int) {
	v.rno = rno
}

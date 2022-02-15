package pkg

import "fmt"

type Medical struct {
	Next Departament
}
func (m *Medical) Execute(p *Patient){
	if p.MedicalDone {
		fmt.Println("Medicine already given to patient")
		m.Next.Execute(p)
		return
	}
	fmt.Println("Medicine giving medicine to patient")
	p.MedicalDone = true
	m.Next.Execute(p)
}

func (m *Medical) SetNext(next Departament){
	m.Next = next
}

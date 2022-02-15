package pkg

import "fmt"

type Reception struct {
	Next Departament
}

func (r *Reception) Execute(p *Patient){
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.Next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.Next.Execute(p)
}

func (r *Reception) SetNext(next Departament){
	r.Next = next
}
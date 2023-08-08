package main

import(
	"fmt"
)
type elec struct {
	u_stress int
	i_flow int
}

func (e *elec) power() int {

	return e.u_stress * e.i_flow

}
func (e *elec) work(time int) int {

	return e.u_stress * e.i_flow * time

}

func main (){

	//this is method of golang,

	e := elec{u_stress:6,i_flow:20}
     fmt.Println("this is method of golang  to tech elec is   and elec work is ",e.power(),e.work(5))  
	

}
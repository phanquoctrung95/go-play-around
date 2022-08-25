/**
 * @author trung
 * @created 25/08/2022
 * @project go-play-around
 */
package main

import (
	"os"
	"text/template"
)

const msg = `The musketeers are:
 {{range .}}{{print .}} {{end}}`

func main() {
	musketeers := []string{"Athos", "Porthos", "Aramis", "D’Artagnan"}

	t := template.Must(template.New("msg").Parse(msg))

	if err := t.Execute(os.Stdout, musketeers); err != nil {
		panic(err)
	}

	u := User{"John", "John33", "john@gmail.com"}
	t2 := template.Must(template.New("msg").Parse(Email))
	if err := t2.Execute(os.Stdout, u); err != nil {
		panic(err)
	}

}

type User struct {
	Name   string
	UserId string
	Email  string
}

const Email = `Dear {{.Name}},
 You were registered with id {{.UserId}}
 and e-mail {{.Email}}.`

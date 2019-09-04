// GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause License

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	defer func ()  {
		err := recover()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}()
	save(down())
}

func down() *http.Response {
	req, err := http.NewRequest("GET","http://chronos.iut-velizy.uvsq.fr/EDTISTY/g68673.pdf",nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("etuisty", "isty")

	rep, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	return rep
}

// Save the Body of a response
func save(rep *http.Response) {
	agenda,err := os.OpenFile("agenda.pdf", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0664)
	if err != nil {
		panic(err)
	}
	defer agenda.Close()
	io.Copy(agenda, rep.Body)
	rep.Body.Close()
}

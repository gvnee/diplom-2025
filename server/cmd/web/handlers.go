package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

type Body struct {
	Input  string `json:"input"`
	Output string `json:"output"`
	Code   string `json:"code"`
}

func test(w http.ResponseWriter, r *http.Request) {
	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("code.cpp", []byte(body.Code), 0o777)
	if err != nil {
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	err = os.WriteFile("input.in", []byte(body.Input), 0o755)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = os.WriteFile("correctOutput.out", []byte(body.Output), 0o755)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	outputFile, err := os.Create("output.out")
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	defer outputFile.Close()

	app := "gcc"
	arg0 := "code.cpp"
	cmd := exec.Command(app, arg0)
	stdout, err1 := cmd.Output()
	if err1 != nil {
		fmt.Println("error1: ", err1.Error())
	}
	fmt.Println(string(stdout))
	arg1 := "< input.in"
	arg2 := "> output.out"
	cmd2 := exec.Command("./a.out", arg1, arg2)
	stdout2, err2 := cmd2.Output()

	if err2 != nil {
		fmt.Println(err2.Error())
	}

	fmt.Println(string(stdout2))
}

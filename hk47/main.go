package main

import (
	"fmt"
	"os/exec"
	"time"
)

func encode(s string) []int {
	t := []int{}
	for _, r := range s {
		t = append(t, int(r-97))
	}
	return t
}

func main() {

	s := "frankensteins monster"

	tunes := encode(s)
	fmt.Println(tunes)
	for _, tune := range tunes {
		if tune == -65 {
			time.Sleep(250 * time.Millisecond)
			_, tunes = tunes[0], tunes[1:]
			continue
		}
		if tune > 24 {
			continue
		}
		cmd := exec.Command("play", fmt.Sprintf("sounds/%d.wav", tune+1), "tempo", "1.5")
		cmd.Start()
		cmd.Wait()
		_, tunes = tunes[0], tunes[1:]
	}
}

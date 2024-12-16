package main

import (
	"fmt"
	"os"
)

type Keys struct {
	public1 int
	public2 int
	private int
	full    *int
}

func (K *Keys) gen_part() int {
	return (K.public1 * K.private) % K.public2
}
func (K *Keys) gen_full(second_part int) int {
	full := (second_part * K.private) % K.public2
	K.full = &full
	return full
}
func (K *Keys) encode(msg string) string {
	key := *K.full
	message := ""
	for _, ch := range msg {
		message += string(ch + rune(key))
	}
	return message
}
func (K *Keys) decode(msg string) string {
	key := *K.full
	message := ""
	for _, ch := range msg {
		message += string(ch - rune(key))
	}
	return message
}
func main() {
	user1_public := 5
	user1_private := 6
	user2_public := 13
	user2_private := 8

	user1 := Keys{public1: user1_public, public2: user2_public, private: user1_private}
	user2 := Keys{public1: user1_public, public2: user2_public, private: user2_private}

	user1.gen_full(user2.gen_part())
	user2.gen_full(user1.gen_part())
	queue := -1
	for {
		queue = queue * (-1)
		var message string
		if queue == 1 {
			fmt.Fscanln(os.Stdin, &message)

			if message == "" {
				break
			}
			user1_encoded := user1.encode(message)
			fmt.Println(user1_encoded)
			user2_decoded := user2.decode(user1_encoded)
			fmt.Println(user2_decoded)
		} else {
			fmt.Fscanln(os.Stdin, &message)

			if message == "" {
				break
			}

			user2_encoded := user2.encode(message)
			fmt.Println(user2_encoded)
			user1_decoded := user1.decode(user2_encoded)
			fmt.Println(user1_decoded)
		}
	}
}

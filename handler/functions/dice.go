package functions

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type dice struct {
	diceCount int
	diceFaces int
}

func DiceExecute(session *discordgo.Session, event *discordgo.MessageCreate) {
	message := event.Content
	dividedMessage := strings.Split(message, " ")

	if len(dividedMessage) == 1 {
		return
	}

	rawDice := dividedMessage[1]

	d, err := createDice(rawDice)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := rollDice(*d)
	session.ChannelMessageSend(event.ChannelID, strconv.Itoa(result))
}

func rollDice(useDice dice) int {
	result := 0
	rand.Seed(time.Now().UnixNano())
	for range make([]int, useDice.diceCount) {
		result += rand.Intn(useDice.diceFaces + 1)
	}
	return result
}

func createDice(message string) (*dice, error) {
	if strings.Count(message, "D") != 1 {
		return nil, errors.New("Message do not match dice format")
	}

	var commands []string = strings.Split(message, "D")

	diceCount, err := strconv.Atoi(commands[0])
	if err != nil {
		return nil, errors.New("diceCount should can cast to int32")
	}
	diceFaces, err := strconv.Atoi(commands[1])

	if err != nil {
		return nil, errors.New("diceCount should can cast to int32")
	}

	result := dice{diceCount, diceFaces}

	return &result, nil
}

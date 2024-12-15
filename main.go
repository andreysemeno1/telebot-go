package main

import (
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	b, err := tele.NewBot(tele.Settings{
		Token:  "7518350168:AAHiWf8ACgkUZgVP7lvC8KRABaX9CabZyXQ",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello user!")

import (
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	b, err := tele.NewBot(tele.Settings{
		Token:  "7518350168:AAHiWf8ACgkUZgVP7lvC8KRABaX9CabZyXQ",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello user!")
	})

	b.Start()
}

	})

	b.Start()
}

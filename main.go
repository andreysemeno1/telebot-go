package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	b, err := tele.NewBot(tele.Settings{
		Token:  "7518350168:AAHiWf8ACgkUZgVP7lvC8KRABaX9CabZyXQ",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		// Ignore response from Send method to avoid errors if there's no connection to Telegram
		return c.Send("Hello user!")
	})

	b.Start()

	b.Handle("/stop", func(c tele.Context) error {
		err := c.Send(b.Close())
		if err != nil {
			b.Close()
		}
		return nil
	})

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt,syscall.SIGTERM ,syscall.SIGINT)

	
		
}
package t4

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

// NumbersChanPrinter Сначала запишем числа в канал, затем обработаем ввод пользователя
func NumbersChanPrinter(numOfWorkers, numOfNums int) {
	stopped := false
	cp := ChannelPrinter{Num: numOfWorkers}
	cp.Launch()
	channel := cp.GetChan()
	for i := 0; i < numOfNums; i++ {
		*channel <- i
	}
	// При получении сигнала CTRL C происходит вызов функции Stop,
	// после этого программа завершается со статусом 0
	go func() {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt)
		<-signalChannel
		stopped = true
		cp.Stop()
		os.Exit(0)
	}()
	for !stopped {
		var str string
		_, err := fmt.Scanf("%s\n", &str)
		if err != nil {
			log.Println(err)

		}
		*channel <- str
	}
}

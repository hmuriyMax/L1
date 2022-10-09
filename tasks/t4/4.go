package t4

import (
	"fmt"
	"log"
)

// NumbersChanPrinter Сначала запишемчисла в канал, затем обработаем ввод пользователя
func NumbersChanPrinter(numOfWorkers, numOfNums int) {
	cp := ChannelPrinter{Num: numOfWorkers}
	cp.Launch()
	defer cp.Stop()
	channel := cp.GetChan()
	for i := 0; i < numOfNums; i++ {
		*channel <- i
	}

	for {
		var str string
		_, err := fmt.Scanln(&str)
		if err != nil {
			log.Fatalln(err)
		}
		*channel <- str
	}
}

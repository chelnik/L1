package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

// структура worker
type worker struct {
	// канкал по которому будут поступать данные для работы воркеров
	workChan chan interface{}
	// идентификатор worker
	workerId int
	// канал с сигналом о прекращении работы worker
	chanExit chan bool
	// канал для сигнала о готовности всей работы
	readyWork chan bool
}

// job - метод выполнения работы для воркеров
func (work *worker) job() {
	for {
		time.Sleep(time.Millisecond * 100)
		select {
		// если данные есть то печатаем их
		case data, ok := <-work.workChan:
			if !ok {
				continue
			}
			fmt.Println(work.workerId, "-", data)
		case <-work.chanExit:
			// как только канал прислал сообщение что работа завершается, считываем остатки по каналу
			for i := range work.workChan {
				fmt.Println(work.workerId, "-", i)
			}
			// сообщение о том что работа завершилась
			work.readyWork <- true
			//break
		}
	}
	fmt.Println("вышел")

}

func main() {
	// Получение данных из аргументов
	n := flag.Int("n", 2, "number of worker")
	flag.Parse()

	// Инициализируем slise work`еров
	workers := make([]worker, *n)
	// инициализируем каналы
	// буферизированный канал для получения данных воркерами
	workChan := make(chan interface{}, 10)
	// канал для завервения работы воркеров
	exitChan := make(chan bool, 1)
	//канал для получения данных о готовности работы
	readyChan := make(chan bool, 1)
	// канал для сигналов
	c := make(chan os.Signal)
	// Указываем какие сигналы ожидаем
	signal.Notify(c, os.Interrupt, os.Kill)

	// инициализируем work`ов
	for i := 0; i < *n; i++ {
		workers[i].workerId = i + 1
		workers[i].workChan = workChan
		workers[i].chanExit = exitChan
		workers[i].readyWork = readyChan
	}
	// Запускаем воркеров в работу
	for i := 0; i < *n; i++ {
		go workers[i].job()
	}
	// Функция отправления бесконечного потока данных в канал
	go func() {
		for {
			select {
			case <-c:
				return
			default:
				time.Sleep(time.Millisecond * 10)
				workChan <- rand.Intn(100)
			}
		}
	}()
	// Ожидаем сигнал
	<-c
	fmt.Println("the job closes...")
	c <- os.Kill
	close(workChan)
	exitChan <- true
	<-readyChan
}

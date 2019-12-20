package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

var (
	totalSum  float64
	failNum   int
	sumLocker = new(sync.RWMutex)
)

func main() {
	if len(os.Args) > 3 {

		goNum := os.Args[2]
		testNum := os.Args[3]

		goNumInt, _ := strconv.Atoi(goNum)
		testNumInt, _ := strconv.Atoi(testNum)

		var (
			wg    = new(sync.WaitGroup)
			wChan = make(chan struct{}, 1000)
		)

		// 開goroutine
		openGoroutineWorker(wChan, wg, goNumInt)

		// 打測試
		for i := 0; i < testNumInt; i++ {
			wg.Add(1)
			wChan <- struct{}{}
		}
		wg.Wait()

		log.Println("=======================")
		log.Println("平均時間:", int(totalSum)/(testNumInt-failNum), "ms,成功測試次數:", (testNumInt - failNum))

	} else {
		log.Println("請輸入以下參數 [網址] [goroutine個數] [測試次數]")
	}
}

func openGoroutineWorker(wChan chan struct{}, wg *sync.WaitGroup, num int) {
	for i := 0; i < num; i++ {
		go workWorrior(wChan, wg)
	}
}

func workWorrior(jobs <-chan struct{}, wg *sync.WaitGroup) {
	for job := range jobs {

		commandPing(job)
		wg.Done()
	}
}

func commandPing(s struct{}) {
	Command := fmt.Sprintf("ping -c 1 " + os.Args[1])
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	if err != nil {
		log.Println("失敗！err is : ", err)
		failNum++
		return
	}
	fmt.Print(string(output))

	out := string(output)

	a := strings.Split(out, "time")
	if len(a) >= 1 {
		b := strings.Split(a[1], "=")
		c := strings.Split(b[1], "ms")
		d := strings.Trim(c[0], " ")

		e, _ := strconv.ParseFloat(d, 64)

		sumLocker.Lock()
		totalSum += e
		sumLocker.Unlock()
	} else {
		log.Println("失敗，拿到info是", out)
		failNum++
	}
}

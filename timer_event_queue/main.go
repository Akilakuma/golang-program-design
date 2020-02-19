package main

import (
	"log"
	"time"
)

func main() {
	em := NewEM()

	event1 := &Event{
		Name:     "e1",
		Period:   5 * time.Second,
		IsRepeat: true,
		Action:   action1,
	}

	event2 := &Event{
		Name:     "e2",
		Period:   7 * time.Second,
		IsRepeat: true,
		Action:   action2,
	}

	event3 := &Event{
		Name:     "e3",
		Period:   10 * time.Second,
		IsRepeat: true,
		Action:   action3,
	}

	event4 := &Event{
		Name:     "e4",
		Period:   2 * time.Second,
		IsRepeat: false,
		Action:   action4,
	}

	em.PushEvent(event1)
	em.PushEvent(event2)
	em.PushEvent(event3)
	em.PushEvent(event4)

	go em.Running()

	for {
		time.Sleep(1 * time.Second)
		log.Println(em.countTime)
		// go em.Close()
	}
}

func action1() error {
	log.Println("我是一隻貓，喵喵！")
	time.Sleep(1 * time.Second)
	return nil
}

func action2() error {
	log.Println("浣熊神敎，只能覺得浣熊最可愛")
	time.Sleep(2 * time.Second)
	return nil
}

func action3() error {
	log.Println("今天天氣真好")
	time.Sleep(3 * time.Second)
	return nil
}

func action4() error {
	log.Println("大家好")
	time.Sleep(1 * time.Second)
	return nil
}

// EventManager 事件管理
type EventManager struct {
	jobQueue        chan *Event
	exitMsg         chan struct{}
	isNeedCountDown bool // 是否需要知道倒數
	countTime       int
	isRunning       bool // 是否正在執行
}

// Event 事件資料
type Event struct {
	Name     string
	Period   time.Duration
	IsRepeat bool
	Action   func() error
}

// NewEM 新的管理器實體
func NewEM() *EventManager {
	return &EventManager{
		jobQueue:        make(chan *Event, 100),
		exitMsg:         make(chan struct{}),
		isNeedCountDown: true,
	}
}

// PushEvent 加入event
func (em *EventManager) PushEvent(event *Event) {
	em.jobQueue <- event
}

// PopEvent 拿出event
func (em *EventManager) PopEvent() *Event {
	event, ok := <-em.jobQueue
	if ok {
		return event
	}

	return nil
}

// Running 啟動執行
func (em *EventManager) Running() {

	log.Println("start running")

	// 首先拿出事件
	// 查看事件的設定時間
	// 根據這個時間設定一個timer
	event := em.PopEvent()
	t := event.Period
	eventTimer := time.NewTimer(t)

	// 下一輪開始倒數
	if em.isNeedCountDown {
		go em.CountDown(t)
	}

	for {
		select {
		// timer 到了
		case <-eventTimer.C:

			log.Println("time arrive")

			eventTimer.Stop()

			// 執行event
			go event.Action()

			// 如果需要重複的話，塞回channel末端
			if event.IsRepeat {
				em.PushEvent(event)
			}

			// 下一輪開始
			event = em.PopEvent()
			t = event.Period
			eventTimer = time.NewTimer(t)

			// 下一輪開始倒數
			if em.isNeedCountDown {
				go em.CountDown(t)
			}

		case <-em.exitMsg:
			eventTimer.Stop()
			log.Println("leave")
			return
		}
	}
}

// Close 啟動執行
func (em *EventManager) Close() {
	if em.isRunning {
		em.exitMsg <- struct{}{}
	}
}

// CountDown 倒數計時
func (em *EventManager) CountDown(Period time.Duration) {

	t := Period
	countDownPeriod := time.NewTicker(t)

	secondCount := time.NewTicker(1 * time.Second)

	defer countDownPeriod.Stop()
	defer secondCount.Stop()

	for {
		select {
		case <-countDownPeriod.C:
			em.countTime = 0
			return
		case <-secondCount.C:
			em.countTime++
		}
	}
}

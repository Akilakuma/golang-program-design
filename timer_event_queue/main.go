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

	em.Running()
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
	jobQueue chan *Event
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
		jobQueue: make(chan *Event, 100),
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
	// 首先拿出事件
	// 查看事件的設定時間
	// 根據這個時間設定一個timer
	event := em.PopEvent()
	t := event.Period
	eventTimer := time.NewTimer(t)

	for {
		// timer 到了
		<-eventTimer.C
		eventTimer.Stop()

		// 執行event
		go event.Action()

		// 如果需要重複的話，塞回channel末端
		if event.IsRepeat {
			em.PushEvent(event)
		}

		event = em.PopEvent()
		t = event.Period
		eventTimer = time.NewTimer(t)
	}

}

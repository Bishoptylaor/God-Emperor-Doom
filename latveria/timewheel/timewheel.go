package timewheel

import (
	"container/list"
	"sync"
	"time"
)

/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2024/7/29 -- 17:34
 @Author  : bishop ❤️ MONEY
 @Description: 时间轮
*/

const defaultSlotNum = int(60)

type TimeWheel struct {
	sync.Once
	interval    time.Duration            // time interval to check and run task
	ticker      *time.Ticker             // ticker to run
	stopChan    chan struct{}            // channel to stop current time wheel
	addTaskChan chan *taskElem           // add task
	delTaskChan chan string              // del task
	slots       []*list.List             // task list
	curSlot     int                      // index of current executing slot
	taskMap     map[string]*list.Element // map of key to taskElm
}

type taskElem struct {
	task  func()
	pos   int    // index in slots; eg:60s as 60 slots,pos is in [0-59],indicate which slot list current taskElem is in
	cycle int    // how many cycle to wait before the task runtime; eg: wait cycle*60 + pos after curSlot to run current taskElem
	key   string // task id
}

func NewTimeWheel(slotNum int, interval time.Duration) *TimeWheel {
	if slotNum <= 0 {
		slotNum = defaultSlotNum
	}
	if interval <= time.Second {
		interval = time.Second
	}
	tw := &TimeWheel{
		Once:        sync.Once{},
		interval:    interval,
		ticker:      time.NewTicker(interval),
		stopChan:    make(chan struct{}),
		addTaskChan: make(chan *taskElem),
		delTaskChan: make(chan string),
		slots:       make([]*list.List, slotNum),
		taskMap:     make(map[string]*list.Element),
	}
	for i := 0; i < slotNum; i++ {
		tw.slots = append(tw.slots, list.New())
	}

	go tw.run()
	return tw
}

func (tw *TimeWheel) run() {
	defer func() {
		if err := recover(); err != nil {
			// log
		}
	}()

	for {
		select {
		case <-tw.stopChan:
			return
		case <-tw.ticker.C:
			tw.tick()
		case taskE := <-tw.addTaskChan:
			tw.addTask(taskE)
		case key := <-tw.delTaskChan:
			tw.delTask(key)
		}
	}
}

func (tw *TimeWheel) Stop() {
	// channel can close only once
	// so sync.once
	tw.Once.Do(func() {
		tw.ticker.Stop()
		close(tw.stopChan)
	})
}

func (tw *TimeWheel) AddTask(key string, task func(), doAt time.Time) {
	pos, cycle := tw.calPosAndCycle(doAt)
	tw.addTaskChan <- &taskElem{
		task:  task,
		pos:   pos,
		cycle: cycle,
		key:   key,
	}
}

func (tw *TimeWheel) calPosAndCycle(doAt time.Time) (int, int) {
	delay := time.Until(doAt)
	cycle := int(delay) / (len(tw.slots) * int(tw.interval))
	pos := (tw.curSlot + int(delay/tw.interval)) % len(tw.slots)
	return pos, cycle
}

func (tw *TimeWheel) addTask(taskE *taskElem) {
	curSlot := tw.slots[taskE.pos]
	if _, ok := tw.taskMap[taskE.key]; ok {
		tw.delTask(taskE.key)
	}
	newE := curSlot.PushBack(taskE)
	// element in list.list
	tw.taskMap[taskE.key] = newE
}

func (tw *TimeWheel) DelTask(key string) {
	tw.delTaskChan <- key
}

func (tw *TimeWheel) delTask(key string) {
	taskListE, ok := tw.taskMap[key]
	if !ok {
		return
	}
	delete(tw.taskMap, key)
	taskE, _ := taskListE.Value.(*taskElem)
	_ = tw.slots[taskE.pos].Remove(taskListE)
}

func (tw *TimeWheel) tick() {
	slot := tw.slots[tw.curSlot]
	defer tw.reCalCurSlot()
	tw.do(slot)
}

func (tw *TimeWheel) do(slot *list.List) {
	for ele := slot.Front(); ele != nil; {
		taskE, _ := ele.Value.(*taskElem)
		// not ready
		if taskE.cycle > 0 {
			taskE.cycle--
			ele = ele.Next()
			continue
		}

		go func() {
			defer func() {
				if err := recover(); err != nil {
					// log
				}
			}()
			taskE.task()
		}()

		// just like linkedlist
		next := ele.Next()
		slot.Remove(ele)
		delete(tw.taskMap, taskE.key)
		ele = next
	}
}

func (tw *TimeWheel) reCalCurSlot() {
	// go to next slot
	tw.curSlot = (tw.curSlot + 1) % len(tw.slots)
}

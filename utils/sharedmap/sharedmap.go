package sharedmap

type SharedMap[K comparable, V interface{}] struct {
	m map[K]V            // 실제 값이 저장될 맵
	c chan command[K, V] // ShardMap에 명령을 전달하기 위한 채널
}

type command[K comparable, V interface{}] struct {
	key    K                  // 키
	value  V                  // 값
	action int                // 수행할 액션
	result chan<- interface{} // 액션 처리 결과
}

const (
	set = iota
	get
	remove
	count
)

func (sm SharedMap[K, V]) Set(k K, v V) {
	sm.c <- command[K, V]{action: set, key: k, value: v}
}

func (sm SharedMap[K, V]) Get(k K) (interface{}, bool) {
	callback := make(chan interface{})
	sm.c <- command[K, V]{action: get, key: k, result: callback}
	result := (<-callback).([2]interface{})
	return result[0], result[1].(bool)
}

func (sm SharedMap[K, V]) Remove(k K) {
	sm.c <- command[K, V]{action: remove, key: k}
}

func (sm SharedMap[K, V]) Count() int {
	callback := make(chan interface{})
	sm.c <- command[K, V]{action: count, result: callback}
	return (<-callback).(int)
}

func (sm SharedMap[K, V]) run() {
	for cmd := range sm.c {
		switch cmd.action {
		case set:
			sm.m[cmd.key] = cmd.value
		case get:
			v, ok := sm.m[cmd.key]
			cmd.result <- [2]interface{}{v, ok}
		case remove:
			delete(sm.m, cmd.key)
		case count:
			cmd.result <- len(sm.m)
		}
	}
}

func New[K comparable, V interface{}]() SharedMap[K, V] {
	sm := SharedMap[K, V]{
		m: make(map[K]V),
		c: make(chan command[K, V]),
	}
	go sm.run()
	return sm
}

package wallclockticker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTicker_DoubleDigitHours_Success(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(3, 23, 52),
			createTestTime(11, 23, 52),
			createTestTime(23, 23, 52),
		},
	}

	testNum := 3

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Hour * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	tickNum := 0
	for {
		if tickNum == len(mockTick.times)-1 {
			return
		}
		select {
		case tm := <-wcTick.C:
			assert.True(t, tm.Hour() == testNum || tm.Hour()%10 == testNum)
			tickNum++

		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}
}

func TestTicker_DoubleDigitHours_Fail(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(22, 23, 52),
			createTestTime(11, 23, 52),
			createTestTime(21, 23, 52),
		},
	}

	testNum := 3

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Hour * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	for {
		select {
		case <-wcTick.C:
			t.Fatal("must fail")
		case <-time.After(time.Second):
			return
		}
	}
}

func TestTicker_SingleDigitHours_Success(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(3, 23, 52),
			createTestTime(11, 23, 52),
			createTestTime(23, 23, 52),
		},
	}

	testNum := 3

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Hour * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	tickNum := 0
	for {
		if tickNum == len(mockTick.times)-1 {
			return
		}
		select {
		case tm := <-wcTick.C:
			assert.True(t, tm.Hour() == testNum || tm.Hour()%10 == testNum)
			tickNum++

		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}
}

func TestTicker_SingleDigitHours_Fail(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(4, 23, 52),
			createTestTime(5, 23, 52),
			createTestTime(1, 23, 52),
		},
	}

	testNum := 3

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Hour * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	for {
		select {
		case <-wcTick.C:
			t.Fatal("must fail")
		case <-time.After(time.Second):
			return
		}
	}
}

func TestTicker_DoubleDigitMinutes_Success(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(3, 44, 52),
			createTestTime(11, 22, 52),
			createTestTime(23, 24, 52),
		},
	}

	testNum := 4

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Minute * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	tickNum := 0
	for {
		if tickNum == len(mockTick.times)-1 {
			return
		}
		select {
		case tm := <-wcTick.C:
			assert.True(t, tm.Minute() == testNum || tm.Minute()%10 == testNum)
			tickNum++

		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}
}

func TestTicker_DoubleDigitMinutes_Fail(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(22, 23, 52),
			createTestTime(11, 11, 52),
			createTestTime(21, 55, 52),
		},
	}

	testNum := 4

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Minute * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	for {
		select {
		case <-wcTick.C:
			t.Fatal("must fail")
		case <-time.After(time.Second):
			return
		}
	}
}

func TestTicker_SingleDigitMinutes_Success(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(3, 4, 52),
			createTestTime(11, 2, 52),
			createTestTime(23, 04, 52),
		},
	}

	testNum := 4

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Minute * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	tickNum := 0
	for {
		if tickNum == len(mockTick.times)-1 {
			return
		}
		select {
		case tm := <-wcTick.C:
			assert.True(t, tm.Minute() == testNum || tm.Minute()%10 == testNum)
			tickNum++

		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}
}

func TestTicker_SingleDigitMinutes_Fail(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(22, 3, 52),
			createTestTime(11, 2, 52),
			createTestTime(21, 01, 52),
		},
	}

	testNum := 4

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Minute * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	for {
		select {
		case <-wcTick.C:
			t.Fatal("must fail")
		case <-time.After(time.Second):
			return
		}
	}
}

func TestTicker_DoubleDigitSeconds_Success(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(3, 44, 52),
			createTestTime(11, 22, 33),
			createTestTime(23, 24, 22),
		},
	}

	testNum := 2

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Second * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	tickNum := 0
	for {
		if tickNum == len(mockTick.times)-1 {
			return
		}
		select {
		case tm := <-wcTick.C:
			assert.True(t, tm.Second() == testNum || tm.Second()%10 == testNum)
			tickNum++

		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}
}

func TestTicker_DoubleDigitSeconds_Fail(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(22, 23, 22),
			createTestTime(14, 11, 33),
			createTestTime(21, 55, 11),
		},
	}

	testNum := 4

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Second * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	for {
		select {
		case <-wcTick.C:
			t.Fatal("must fail")
		case <-time.After(time.Second):
			return
		}
	}
}

func TestTicker_SingleDigitSeconds_Success(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(3, 44, 2),
			createTestTime(11, 22, 33),
			createTestTime(23, 24, 02),
		},
	}

	testNum := 2

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Second * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	tickNum := 0
	for {
		if tickNum == len(mockTick.times)-1 {
			return
		}
		select {
		case tm := <-wcTick.C:
			assert.True(t, tm.Second() == testNum || tm.Second()%10 == testNum)
			tickNum++

		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}
}

func TestTicker_SingleDigitSeconds_Fail(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(22, 23, 1),
			createTestTime(14, 11, 5),
			createTestTime(21, 55, 8),
		},
	}

	testNum := 4

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Second * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	for {
		select {
		case <-wcTick.C:
			t.Fatal("must fail")
		case <-time.After(time.Second):
			return
		}
	}
}

func TestTicker_BigMinutesDuration_Success(t *testing.T) {
	mockTick := mockTicker{
		c:    make(chan time.Time, 1),
		done: make(chan struct{}),
		times: []time.Time{
			createTestTime(1, 9, 2),
			createTestTime(11, 22, 33),
			createTestTime(21, 29, 02),
		},
	}

	testNum := 69 // 1h 09m

	c := make(chan time.Time, 1)
	wcTick := WCTicker{
		C:           c,
		pos:         time.Time{}.Add(time.Minute * time.Duration(testNum)),
		accuracy:    time.Millisecond,
		done:        make(chan struct{}),
		innerTicker: &mockTick,
	}

	go mockTick.run()
	go wcTick.tick(c)

	tickNum := 0
	for {
		if tickNum == len(mockTick.times)-1 {
			return
		}
		select {
		case tm := <-wcTick.C:
			assert.True(t, tm.Minute() == testNum || tm.Minute()%10 == testNum%10)
			assert.True(t, tm.Hour() == int(testNum/60) || tm.Hour()%10 == int(testNum/60))

			tickNum++

		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}
}

func TestTicker_Close_Success(t *testing.T) {
	wcTick := NewCWTicker(time.Second, time.Millisecond)

	go wcTick.Stop()
	for now := range wcTick.C {
		t.Fatal(now)
	}
}

func createTestTime(h, m, s int) time.Time {
	return time.Date(2020, 3, 24, h, m, s, 10, time.Local)
}

type mockTicker struct {
	c     chan time.Time
	done  chan struct{}
	times []time.Time
}

func (t mockTicker) run() {
	for _, tm := range t.times {
		t.c <- tm
	}
}

func (t mockTicker) Ch() <-chan time.Time {
	return t.c
}

func (t mockTicker) Stop() {
	close(t.done)
}

package wallclockticker

import (
	"time"
)

type iTicker interface {
	Ch() <-chan time.Time
	Stop()
}

type innerTicker struct {
	*time.Ticker
}

func newInnerTicker(d time.Duration) iTicker {
	return &innerTicker{Ticker: time.NewTicker(d)}
}

func (t *innerTicker) Ch() <-chan time.Time {
	return t.C
}

type WCTicker struct {
	C           <-chan time.Time
	pos         time.Time
	accuracy    time.Duration
	done        chan struct{}
	innerTicker iTicker
}

func NewCWTicker(clockArrowsPosition, accuracy time.Duration) *WCTicker {
	if clockArrowsPosition <= accuracy {
		panic("accuracy check duration is not less than wall clock duration")
	}
	if accuracy < time.Millisecond {
		panic(
			"there is no point to set accuracy check duration less than microsecond" +
				" - regular wall clock have only seconds precision",
		)
	}

	c := make(chan time.Time, 1)
	t := &WCTicker{
		C:           c,
		pos:         time.Time{}.Add(clockArrowsPosition),
		accuracy:    accuracy,
		done:        make(chan struct{}),
		innerTicker: newInnerTicker(accuracy),
	}
	go t.tick(c)

	return t
}

func (t *WCTicker) Stop() {
	if t.innerTicker != nil {
		t.innerTicker.Stop()
	}
	close(t.done)
}

func (t *WCTicker) tick(c chan time.Time) {
	var isTicked bool
	for {
		select {
		case now := <-t.innerTicker.Ch():
			if !t.isEqualDuration(now) {
				isTicked = false
				continue
			}
			if isTicked {
				continue
			}
			isTicked = true

			select {
			case c <- now:
			default:
			}
		case <-t.done:
			close(c)
			return
		}
	}

}

func (t *WCTicker) isEqualDuration(now time.Time) bool {
	posH, posM, posS := t.pos.Clock()
	nowH, nowM, nowS := now.Clock()

	isHourEqual := posH == 0 || posH == nowH || posH == nowH%10
	isMinuteEqual := posM == 0 || posM == nowM || posM == nowM%10
	isSecondEqual := posS == 0 || posS == nowS || posS == nowS%10

	return isHourEqual && isMinuteEqual && isSecondEqual

}

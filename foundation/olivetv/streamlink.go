package olivetv

import (
	"os/exec"
	"time"
	"math/rand"
)

func init() {
	registerSite("streamlink", &streamlink{})
}

type streamlink struct {
	base
}

func (this *streamlink) Name() string {
	return "streamlink"
}

func (this *streamlink) Snap(tv *TV) error {
	tv.Info = &Info{
		Timestamp: time.Now().Unix(),
	}
	return this.set(tv)
}

func delay() {
        rand.Seed(time.Now().UnixNano()) // 设置随机数种子
        n := rand.Intn(30) + 5           // 生成5到30之间的随机数
        time.Sleep(time.Duration(n) * time.Second) // 延迟n秒
}

func (this *streamlink) set(tv *TV) error {
	defer delay() // 延迟调用delay
	cmd := exec.Command(
		"streamlink",
		tv.RoomID,
	)
	if err := cmd.Run(); err != nil {
		return nil
	}

	tv.roomOn = true
	tv.streamURL = tv.RoomID

	return nil
}

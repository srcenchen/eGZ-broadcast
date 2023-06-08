package mediaService

/**
 * 媒体控制器
 */
import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
	"os"
	"path"
	"sync"
	"time"
)

var sr beep.SampleRate = 44100                                            // 默认采样率
var Format beep.Format                                                    // Format
var Streamer beep.StreamSeekCloser                                        // 流
var Ctrl = &beep.Ctrl{Streamer: beep.Loop(Loop, Streamer), Paused: false} // 控制器
var Loop = 1                                                              // 是否循环

var once sync.Once

// Play 播放音乐
func Play(fileName string, sync bool, loop bool) {
	musicFile, _ := os.Open("./resource/music/" + fileName)
	speaker.Clear() // 清空speaker
	switch path.Ext(fileName) {
	case ".mp3":
		Streamer, Format, _ = mp3.Decode(musicFile)
		break
	case ".wav":
		Streamer, Format, _ = wav.Decode(musicFile)
		break
	case ".flac":
		Streamer, Format, _ = flac.Decode(musicFile)
		break
	case ".ogg":
		Streamer, Format, _ = vorbis.Decode(musicFile)
		break
	default:
		break
	}
	// 首次初始化
	once.Do(func() {
		_ = speaker.Init(Format.SampleRate, Format.SampleRate.N(time.Second/10)) // 初始化speaker
		sr = Format.SampleRate
	})
	if loop {
		Loop = -1
	} else {
		Loop = 1
	}
	Ctrl = &beep.Ctrl{Streamer: beep.Loop(Loop, Streamer), Paused: false} // 更新重置控制器
	resampled := beep.Resample(4, Format.SampleRate, sr, Ctrl)            // 重新获取采样率

	if sync {
		playSpeakerSync(resampled)
	} else {
		playSpeaker(resampled)
	}
}

var done = make(chan bool)

// playSpeaker 播放音乐 异步
func playSpeaker(streamer beep.Streamer) {
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		_ = Streamer.Seek(0) // 重置播放位置
		MusicName = "未在播放"
	}))) // 播放 完成回调
}

// playSpeakerSync 播放音乐 同步
func playSpeakerSync(streamer beep.Streamer) {
	done = make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		_ = Streamer.Seek(0) // 重置播放位置
		MusicName = "未在播放"
		done <- true
	}))) // 播放 完成回调
	<-done
}

// Pause 暂停播放音乐
func Pause() {
	speaker.Lock()
	Ctrl.Paused = true
	speaker.Unlock()
}

// Resume 恢复播放音乐
func Resume() {
	speaker.Lock()
	Ctrl.Paused = false
	speaker.Unlock()
}

// Stop 停止播放音乐
func Stop() {
	_ = Streamer.Seek(0) // 重置播放位置
	MusicName = "未在播放"
	Loop = 1
	speaker.Clear()
}

// EnableLoop 开启循环播放
func EnableLoop() {
	speaker.Lock()
	Loop = -1 // 循环播放
	Ctrl.Streamer = beep.Loop(Loop, Streamer)
	speaker.Unlock()
}

// DisableLoop 关闭循环播放
func DisableLoop() {
	speaker.Lock()
	Loop = 1 // 不循环播放
	Ctrl.Streamer = beep.Loop(Loop, Streamer)
	speaker.Unlock()
}

// NextSong 下一首
func NextSong() {
	done <- true
}

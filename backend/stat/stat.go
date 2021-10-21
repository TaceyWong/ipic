package stat

import (
	"encoding/json"

	"github.com/TaceyWong/ipic/backend/localstore"
	"github.com/wailsapp/wails"
)

const filename = "stats.json"

// Stat represents application statistics.
type Stat struct {
	ByteCount  int64 `json:"byteCount"`
	ImageCount int   `json:"imageCount"`
	TimeCount  int64 `json:"timeCount"`

	Runtime *wails.Runtime
	Logger  *wails.CustomLogger

	localStore *localstore.LocalStore
}

// NewStat returns a new Stat instance.
func NewStat() *Stat {
	s := &Stat{
		localStore: localstore.NewLocalStore(),
	}

	d, _ := s.localStore.Load(filename)
	_ = json.Unmarshal(d, &s)
	return s
}

// WailsInit performs setup when Wails is ready.
func (s *Stat) WailsInit(runtime *wails.Runtime) error {
	s.Runtime = runtime
	s.Logger = s.Runtime.Log.New("Stat")
	s.Logger.Info("Stat initialized...")
	return nil
}

// GetStats returns the application stats.
func (s *Stat) GetStats() map[string]interface{} {
	return map[string]interface{}{
		"byteCount":  s.ByteCount,
		"imageCount": s.ImageCount,
		"timeCount":  s.TimeCount,
	}
}

// SetByteCount adds and persists the given byte count to the app stats.
func (s *Stat) SetByteCount(b int64) {
	if b <= 0 {
		return
	}
	s.ByteCount += b
	if err := s.store(); err != nil {
		s.Logger.Errorf("failed to store stats: %v", err)
	}
}

// SetImageCount adds and persists the given image count to the app stats.
func (s *Stat) SetImageCount(i int) {
	if i <= 0 {
		return
	}
	s.ImageCount += i
	if err := s.store(); err != nil {
		s.Logger.Errorf("failed to store stats: %v", err)
	}
}

// SetTimeCount adds and persists the given time count to the app stats.
func (s *Stat) SetTimeCount(t int64) {
	if t < 0 {
		return
	}
	s.TimeCount += t
	if err := s.store(); err != nil {
		s.Logger.Errorf("failed to store stats: %v", err)
	}
}

// store stores the app stats to the file system.
func (s *Stat) store() error {
	js, err := json.Marshal(s.GetStats())
	if err != nil {
		return err
	}
	if err = s.localStore.Store(js, filename); err != nil {
		return err
	}
	return nil
}

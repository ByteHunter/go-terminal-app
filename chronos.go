package main

import "time"

type Chronos struct {
  ms, fps float64
  avg_fps float64
  avg_size int
  start time.Time
}

func NewChronos() *Chronos {
  return &Chronos{}
}

func (chronos *Chronos) get_millis() (float64) {
  return chronos.ms
}

func (chronos *Chronos) get_fps() (float64) {
  return chronos.fps
}

func (chronos *Chronos) begin() {
  chronos.start = time.Now()
}

func (chronos *Chronos) measure() {
	end := time.Now()
	elapsed := end.Sub(chronos.start)
	ms := float64(elapsed) / 1_000_000
	fps := 1 / (ms / 1_000)

  avg, size := chronos.new_average(chronos.avg_fps, chronos.avg_size, fps)
  chronos.ms = ms
  chronos.fps = fps
  chronos.avg_fps = avg
  chronos.avg_size = size
}

func (chrono *Chronos) new_average(old float64, size int, value float64) (float64, int) {
	new_size := size + 1
	new_average := old + (value-old)/float64(new_size)
	return new_average, new_size
}

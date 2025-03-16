package core

import (
	"github.com/notfrancois/omp-streamer/internal/spatial"
	"github.com/notfrancois/omp-streamer/pkg/openmp"
)

type Core struct {
	grid     *spatial.Grid
	data     *Data
	streamer *Streamer
	openmp   *openmp.API
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) GetData() *Data {
	return c.data
}

func (c *Core) GetGrid() *spatial.Grid {
	return c.grid
}

func (c *Core) GetStreamer() *Streamer {
	return c.streamer
}

func (c *Core) GameSdk() *openmp.API {
	return c.openmp
}

var GlobalCore Service

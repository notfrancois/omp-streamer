package core

import (
	"github.com/notfrancois/omp-streamer/internal/spatial"
	"github.com/notfrancois/omp-streamer/pkg/openmp"
)

type Service interface {
	GetData() *Data
	GetGrid() *spatial.Grid
	GetStreamer() *Streamer
	GameSdk() *openmp.API
}

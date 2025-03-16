package cell_test

import (
	"testing"

	spatial "github.com/notfrancois/omp-streamer/pkg/spatial/cell"
	"github.com/notfrancois/omp-streamer/pkg/types/common"
)

func TestCellImplementsInterface(t *testing.T) {
	var cell common.Cell = &spatial.Cell{}

	// check if the cell implements the common.Cell interface
	if cell == nil {
		t.Errorf("cell is nil")
	}
}

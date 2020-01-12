package mat

import (
	"github.com/mogaika/go-collada"

	"github.com/mogaika/god_of_war_browser/pack/wad"
)

type ColladaExportContext struct {
	Prefix string

	m *Material
}

func (m *Material) ColladaExport(wrsrc *wad.WadNodeRsrc, c *collada.Collada) (*ColladaExportContext, error) {
	cec := &ColladaExportContext{
		Prefix: wrsrc.Name(),
	}
	return cec, nil
}

/*
import (
	"fmt"
	"log"

	"github.com/mogaika/god_of_war_browser/utils"

	"github.com/mogaika/go-collada"

	"github.com/mogaika/god_of_war_browser/pack/wad"
	"github.com/mogaika/god_of_war_browser/utils/colexpu"
)


type ColladaExportContext struct {
	Prefix  string

	Objects []*ColladaExportObject
	m       *Mesh
}
*/

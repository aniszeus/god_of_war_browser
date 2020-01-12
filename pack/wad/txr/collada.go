package txr

import (
	"fmt"
	"log"

	"github.com/mogaika/god_of_war_browser/utils"

	"github.com/mogaika/go-collada"

	"github.com/mogaika/god_of_war_browser/pack/wad"
	"github.com/mogaika/god_of_war_browser/utils/colexpu"
)

/*
import (
	"fmt"
	"log"
	"github.com/mogaika/god_of_war_browser/utils"
	"github.com/mogaika/go-collada"
	"github.com/mogaika/god_of_war_browser/pack/wad"
	"github.com/mogaika/god_of_war_browser/utils/colexpu"
)
*/

type ColladaExportContext struct {
	ImageId string
}

func (t *Texture) ColladaExport(wrsrc *wad.WadNodeRsrc, c *collada.Collada) (*ColladaExportContext, error) {
	cec := &ColladaExportContext{
		ImageId: fmt.Sprintf("%d_%s", wrsrc.Tag.Id, wrsrc.Tag.Name),
	}

}

package mesh

import (
	"bytes"
	"log"
	"net/http"

	"github.com/mogaika/god_of_war_browser/pack/wad"
	"github.com/mogaika/god_of_war_browser/webutils"
)

func (mesh *Mesh) HttpAction(wrsrc *wad.WadNodeRsrc, w http.ResponseWriter, r *http.Request, action string) {
	switch action {
	case "obj":
		var buf bytes.Buffer
		log.Printf("Error when exporting mesh: %v", mesh.ExportObj(&buf, nil, nil))
		webutils.WriteFile(w, bytes.NewReader(buf.Bytes()), wrsrc.Tag.Name+".obj")
	case "collada":
		c, err := mesh.ColladaExportDefault(wrsrc)
		if err != nil {
			webutils.WriteError(w, err)
		} else {
			var b bytes.Buffer
			if err := c.ExportToWriter(&b); err != nil {
				webutils.WriteError(w, err)
			} else {
				webutils.WriteFile(w, &b, wrsrc.Tag.Name+".dae")
				//webutils.WriteResult(w, b.Bytes())
			}
		}
	}
}

package colexpu

import (
	"time"

	"github.com/mogaika/go-collada"
)

func GetDefaultCollada() *collada.Collada {
	c := &collada.Collada{}
	c.Xmlns = `http://www.collada.org/2005/11/COLLADASchema`
	c.Version = "1.4.1"
	c.Asset = &collada.Asset{}
	c.Asset.Unit = &collada.Unit{}
	c.Asset.Unit.Meter = 0.01
	c.Asset.Unit.Name = "centimeter"
	c.Asset.Created = time.Now().Format(time.RFC3339)
	c.Asset.Modified = time.Now().Format(time.RFC3339)
	c.Asset.Revision = "1.0"
	return c
}

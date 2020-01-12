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

	c.LibraryGeometries = []*collada.LibraryGeometries{
		&collada.LibraryGeometries{
			Geometry: make([]*collada.Geometry, 0),
		},
	}
	c.LibraryVisualScenes = []*collada.LibraryVisualScenes{
		&collada.LibraryVisualScenes{
			VisualScene: make([]*collada.VisualScene, 0),
		},
	}
	c.LibraryEffects = []*collada.LibraryEffects{
		&collada.LibraryEffects{},
	}
	c.LibraryImages = []*collada.LibraryImages{
		&collada.LibraryImages{},
	}
	c.LibraryAnimations = []*collada.LibraryAnimations{
		&collada.LibraryAnimations{},
	}
	c.LibraryAnimationClips = []*collada.LibraryAnimationClips{
		&collada.LibraryAnimationClips{},
	}
	c.LibraryControllers = []*collada.LibraryControllers{
		&collada.LibraryControllers{},
	}

	return c
}

func GetDefaultColladaWithVisualScene() (c *collada.Collada, scene *collada.VisualScene) {
	c = GetDefaultCollada()

	scene = &collada.VisualScene{}
	scene.Id = "default-scene"
	scene.Name = "Default scene"

	c.LibraryVisualScenes[0].VisualScene = append(c.LibraryVisualScenes[0].VisualScene, scene)

	c.Scene = &collada.Scene{}
	c.Scene.InstanceVisualScene = &collada.InstanceVisualScene{}
	c.Scene.InstanceVisualScene.Url = collada.Uri("#" + scene.Id)

	return
}

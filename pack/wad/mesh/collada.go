package mesh

import (
	"fmt"
	"log"

	"github.com/mogaika/god_of_war_browser/utils"

	"github.com/mogaika/go-collada"

	"github.com/mogaika/god_of_war_browser/pack/wad"
	"github.com/mogaika/god_of_war_browser/utils/colexpu"
)

type ColladaExportObject struct {
	GeometryId      string
	ColladaGeometry *collada.Geometry
	Part            int
	Group           int
	Object          int
	MaterialId      int
	InstanceId      int
}

type ColladaExportContext struct {
	Prefix  string
	Objects []*ColladaExportObject
	m       *Mesh
}

func uint8ColToF32(col uint16) float32 {
	return float32(col) / 128.0
}

func uint8AlpfaToF32(col uint16) float32 {
	return float32(col) / 255.0
}

func (cec *ColladaExportContext) exportObject(ceco *ColladaExportObject) error {
	o := &cec.m.Parts[ceco.Part].Groups[ceco.Group].Objects[ceco.Object]

	vertices := make([]float32, 0)
	indexes := make([]int, 0)
	rgba := make([]float32, 0)
	normals := make([]float32, 0)

	// first extract pos, color, norm
	for iPacket := range o.Packets[0] {
		packet := o.Packets[0][iPacket]

		haveNorm := packet.Norms.X != nil
		haveRgba := packet.Blend.R != nil

		for iVertex := range packet.Trias.X {
			if !packet.Trias.Skip[iVertex] {
				curIndex := len(vertices) / 3
				indexes = append(indexes, curIndex-2, curIndex-1, curIndex)
			}

			vertices = append(vertices,
				packet.Trias.X[iVertex], packet.Trias.Y[iVertex], packet.Trias.Z[iVertex])
			if haveNorm {
				normals = append(normals,
					packet.Norms.X[iVertex], packet.Norms.Y[iVertex], packet.Norms.Z[iVertex])
			}
			if haveRgba {
				rgba = append(rgba,
					uint8ColToF32(packet.Blend.R[iVertex]),
					uint8ColToF32(packet.Blend.G[iVertex]),
					uint8ColToF32(packet.Blend.B[iVertex]))
				// no alpha because of collada TODO
				// TODO: implement alpha export
				// uint8AlpfaToF32(packet.Blend.A[iVertex]))
			}

		}
	}

	// extract texture layers

	cm := &collada.Mesh{}

	sourcePosition := &collada.Source{}
	sourcePosition.Id = collada.Id(ceco.GeometryId + "-positions")
	sourcePosition.FloatArray = &collada.FloatArray{}
	sourcePosition.FloatArray.V = colexpu.Floats32ToString(vertices)
	sourcePosition.FloatArray.Count = len(vertices)
	sourcePosition.FloatArray.Id = sourcePosition.Id + "-array"
	sourcePosition.TechniqueCommon.XML = colexpu.CreateAccessor(
		len(vertices)/3, 0, "#"+string(sourcePosition.FloatArray.Id), 3,
		"X", "float", "Y", "float", "Z", "float")

	cm.Source = append(cm.Source, sourcePosition)

	cm.Vertices.Id = collada.Id(ceco.GeometryId + "-vertices")
	cm.Vertices.Input = make([]*collada.InputUnshared, 1)
	cm.Vertices.Input[0] = &collada.InputUnshared{
		Semantic: "POSITION",
		Source:   collada.Uri("#" + sourcePosition.Id),
	}

	cmTriangles := &collada.Triangles{}
	cmTriangles.Count = len(indexes) / 3
	cmTriangles.Input = make([]*collada.InputShared, 1)
	cmTriangles.Input[0] = &collada.InputShared{
		Semantic: "VERTEX",
		Source:   collada.Uri("#" + cm.Vertices.Id),
	}
	cmTriangles.P = &collada.P{}
	cmTriangles.P.V = colexpu.IntsToString(indexes)

	cm.Triangles = []*collada.Triangles{cmTriangles}

	ceco.ColladaGeometry = &collada.Geometry{}
	ceco.ColladaGeometry.Mesh = cm
	ceco.ColladaGeometry.Id = collada.Id(ceco.GeometryId)

	return nil
}

func (m *Mesh) ColladaExport(wrsrc *wad.WadNodeRsrc, c *collada.Collada) (*ColladaExportContext, error) {
	cec := &ColladaExportContext{
		Prefix: wrsrc.Name(),
		m:      m,
	}

	for iPart := range m.Parts {
		part := &m.Parts[iPart]
		for iGroup := range part.Groups {
			group := &part.Groups[iGroup]
			for iObject := range group.Objects {
				object := &group.Objects[iObject]

				for iInstance := uint32(0); iInstance < object.InstancesCount; iInstance++ {
					objecId := fmt.Sprintf("%s_p%v_g%v_o%v_i%v", cec.Prefix, iPart, iGroup, iObject, iInstance)
					log.Printf("parsing %v", objecId)

					ceco := &ColladaExportObject{
						GeometryId: objecId,
						Part:       iPart,
						Group:      iGroup,
						Object:     iObject,
						InstanceId: int(iInstance),
					}
					if err := cec.exportObject(ceco); err != nil {
						return nil, fmt.Errorf("Error parsing %v: %v", objecId, err)
					}
					cec.Objects = append(cec.Objects, ceco)
				}
			}
		}
	}
	cec.m = nil // free memory

	return cec, nil
}

func (m *Mesh) ColladaExportDefault(wrsrc *wad.WadNodeRsrc) (*collada.Collada, error) {
	c := colexpu.GetDefaultCollada()

	cec, err := m.ColladaExport(wrsrc, c)
	if err != nil {
		return nil, err
	}

	node := &collada.Node{}
	node.Id = collada.Id(wrsrc.Name())
	node.Name = wrsrc.Name()

	/*
		node.Translate = []*collada.Translate{&collada.Translate{}}
		node.Translate[0].V = "0 0 0"

		node.Rotate = []*collada.Rotate{&collada.Rotate{}, &collada.Rotate{}, &collada.Rotate{}}
		node.Rotate[0].V = "0 0 1 0"
		node.Rotate[1].V = "0 1 0 0"
		node.Rotate[2].V = "1 0 0 0"

		node.Scale = []*collada.Scale{&collada.Scale{}}
		node.Scale[0].V = "1 1 1"
	*/

	geoms := make([]*collada.Geometry, 0)

	node.InstanceGeometry = make([]*collada.InstanceGeometry, len(cec.Objects))
	for iObject, object := range cec.Objects {
		instance := &collada.InstanceGeometry{}
		node.InstanceGeometry[iObject] = instance

		instance.Url = collada.Uri("#" + object.GeometryId)

		geoms = append(geoms, object.ColladaGeometry)
		log.Printf("inserting %v", object.GeometryId)
	}

	scene := &collada.VisualScene{}
	scene.Id = "default-scene"
	scene.Name = "Default scene"
	scene.Node = []*collada.Node{node}

	c.LibraryGeometries = []*collada.LibraryGeometries{
		&collada.LibraryGeometries{Geometry: geoms},
	}
	c.LibraryVisualScenes = []*collada.LibraryVisualScenes{
		&collada.LibraryVisualScenes{
			VisualScene: []*collada.VisualScene{scene},
		},
	}
	c.Scene = &collada.Scene{}
	c.Scene.InstanceVisualScene = &collada.InstanceVisualScene{}
	c.Scene.InstanceVisualScene.Url = collada.Uri("#" + scene.Id)

	utils.LogDump(c)

	return c, nil
}

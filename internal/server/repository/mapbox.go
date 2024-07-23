package repository

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"maps-service/config"
	"maps-service/internal/models"
	pb "protos/maps"
	"strings"
)

func (r *Repository) StyledMap(mapID string, cfg *config.Config) (*pb.StyledMap, error) {
	log := hclog.Default()

	mapModel := models.Map{}
	labelIndex := make([]uuid.UUID, 0)
	var styled []*models.PgStyledMapModel

	result := r.DB.Where("id = ?", mapID).First(&mapModel)
	if result.Error != nil {
		log.Error("[repository.StyledMap] r.DB.Find failed", "error", result.Error)
		return nil, result.Error
	}

	result = r.DB.Raw(styledMap, mapID).Scan(&styled)
	if result.Error != nil {
		log.Error("[repository.StyledMap] r.DB.Raw failed", "error", result.Error)
		return nil, result.Error
	}

	var response pb.StyledMap

	response.Id = mapModel.ID
	response.Version = 8
	response.Name = mapModel.Name
	response.Sprite = strings.Join(
		[]string{"http://", cfg.ParentServerHost, "V0.0.0", "maps", "pattern", "sprite"},
		"/",
	)
	response.Glyphs = strings.Join(
		[]string{"http://", cfg.ParentServerHost, "V0.1.0", "glyphs", "fonts", "{fontstack}", "{range}.pbf"},
		"/",
	)

	var sources string
	for _, v := range styled {
		if v.LayerType == "vector" && v.LayerTableID != "" {
			sources = sources + v.LayerTableID + ","
		}
	}

	if len(sources) < 1 {
		sources = "*"
	}

	Sources := make(map[string]*pb.Source)
	Sources["sources"] = &pb.Source{
		Type: "vector",
		Tiles: []string{
			strings.Join([]string{
				"http://",
				cfg.ParentServerHost,
				"V0.0.0",
				"gis",
				"tiles",
				sources[:len(sources)-1],
				"{z}/{x}/{y}.pbf",
			}, "/"),
		},
	}

	response.Sources = Sources

	for _, v := range styled {
		index := uuid.New()
		label := uuid.New()
		labelIndex = append(labelIndex, label)

		P, L := getPallet(v)
		ptrLabel := label.String()
		ptrIndex := index.String()
		response.Layers = append(response.Layers, &pb.LayerMapbox{
			Id:          &ptrIndex,
			Type:        v.StyleType,
			MaxZoom:     &v.StyleMaxZoom,
			MinZoom:     &v.StyleMinZoom,
			Paint:       &P,
			Layout:      &L,
			SourceLayer: v.LayerTableID,
			Source:      getSource(v),
			Label:       &ptrLabel,
			LayerType:   "layer",
		})
	}

	for i, v := range styled {
		visibility := "none"
		if v.StyleLabel {
			visibility = "visible"
		} else {
			visibility = "none"
		}

		ptrIndex := labelIndex[i].String()

		response.Labels = append(response.Labels, &pb.LabelMapbox{
			Id:        &ptrIndex,
			Type:      "symbol",
			LayerType: "label",
			Active:    v.StyleLabel,
			MaxZoom:   &v.StyleMaxZoom,
			MinZoom:   &v.StyleMinZoom,
			Paint: &pb.PaintLB{
				TextColor:     &v.LabelTextColor,
				TextHaloWidth: &v.LabelTextHaloWidth,
				TextHaloBlur:  &v.LabelTextHaloBlur,
				TextHaloColor: &v.LabelTextHaloColor,
				TextOpacity:   &v.LabelTextOpacity,
			},
			Layout: &pb.LayoutLB{
				Visibility: visibility,
				TextField:  &v.LabelTextField,
				TextFont:   &v.LabelTextFont,
				TextRotate: &v.LabelTextRotate,
				TextSize:   &v.LabelTextSize,
			},
			SourceLayer: v.LayerTableID,
			Source:      getSource(v),
		})
	}

	return &response, nil
}

func getPallet(v *models.PgStyledMapModel) (pb.Paint, pb.Layout) {
	switch v.LayerType {
	case "vector":
		switch v.StyleType {
		case "fill":
			return pb.Paint{
					FillAntialias:    true,
					FillColor:        &v.FillColor,
					FillOpacity:      &v.FillOpacity,
					FillOutlineColor: &v.FillOutlineColor,
					FillPattern:      &v.FillPattern,
				}, pb.Layout{
					Visibility:  v.FillVisibility,
					FillSortKey: nil,
				}
		case "line":
			return pb.Paint{
					//LineBlur:      v.LineB,
					LineColor:     &v.LineColor,
					LineGapWidth:  &v.LineGapWidth,
					LineOpacity:   &v.LineOpacity,
					LineWidth:     &v.LineWidth,
					LinePattern:   &v.LinePattern,
					LineDasharray: &v.LineDasharray,
				}, pb.Layout{
					Visibility: v.LineVisibility,
					LineCap:    &v.LineCap,
					LineJoin:   &v.LineJoin,
				}
		case "symbol":
			return pb.Paint{
					TextColor:     &v.SymbolTextColor,
					TextHaloWidth: &v.SymbolTextHaloWidth,
					TextHaloBlur:  &v.SymbolTextHaloBlur,
					TextHaloColor: &v.SymbolTextHaloColor,
					//TextOpacity:   v.SymbolTextOpacity,
					//TextTranslate:       v.SymbolText,
					//TextTranslateAnchor: v.SymbolTe,
				}, pb.Layout{
					Visibility:       "visible",
					TextAllowOverlap: &v.SymbolTextAllowOverlap,
					TextField:        &v.SymbolTextField,
					TextFont:         &v.SymbolTextFont,
					TextJustify:      &v.SymbolTextJustify,
					//TextMaxAngle:        v.SymbolTextM,
					//TextMaxWidth:        v.SymbolTextMaxWidth,
					TextOffset:          &v.SymbolTextOffset,
					TextIgnorePlacement: &v.SymbolTextIgnorePlacement,
					//TextPadding:         v.SymbolTextPadding,
					TextRotate: &v.SymbolTextRotate,
					TextSize:   &v.SymbolTextSize,
					//TextTransform:       v.SymbolTe,
				}
		default:
			return pb.Paint{}, pb.Layout{}
		}
	default:
		return pb.Paint{}, pb.Layout{}
	}
}

func getSource(v *models.PgStyledMapModel) string {
	switch v.LayerType {
	case "vector":
		return "sources"
	case "raster":
		return v.LayerName
	default:
		return "sources"
	}
}

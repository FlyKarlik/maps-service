package models

import pb "protos/maps"

type proj struct {
	Name      string     `json:"name"`
	Center    [2]float64 `json:"center"`
	Parallels [2]float64 `json:"parallels"`
}

type MapSource struct {
	Type             string      `json:"type"`
	Tiles            []string    `json:"tiles,omitempty"`
	Attribution      string      `json:"attribution,omitempty"`
	Bounds           []float64   `json:"bounds,omitempty"`
	MaxZoom          float64     `json:"maxzoom,omitempty"`
	MinZoom          float64     `json:"minzoom,omitempty"`
	Scheme           string      `json:"scheme,omitempty"`
	Url              string      `json:"url,omitempty"`
	Volatile         bool        `json:"volatile,omitempty"`
	Encoding         string      `json:"encoding,omitempty"`
	TileSize         int         `json:"tileSize,omitempty"`
	Buffer           int         `json:"buffer,omitempty"`
	Cluster          bool        `json:"cluster,omitempty"`
	ClusterMaxZoom   int         `json:"clusterMaxZoom,omitempty"`
	ClusterMinPoints int         `json:"clusterMinPoints,omitempty"`
	ClusterRadius    int         `json:"clusterRadius,omitempty"`
	Data             interface{} `json:"data,omitempty"`
	//Coordinates      [4][2]float64 `json:"coordinates,omitempty"`
	Urls []string `json:"urls,omitempty"`
}

type Paint struct {
	BackgroundColor     string      `json:"background-color,omitempty"`
	BackgroundOpacity   float64     `json:"background-opacity,omitempty"`
	BackgroundPattern   string      `json:"background-pattern,omitempty"`
	FillAntialias       interface{} `json:"fill-antialias,omitempty"`
	FillColor           interface{} `json:"fill-color,omitempty"`
	FillOpacity         interface{} `json:"fill-opacity,omitempty"`
	FillOutlineColor    interface{} `json:"fill-outline-color,omitempty"`
	FillPattern         interface{} `json:"fill-pattern,omitempty"`
	FillTranslate       interface{} `json:"fill-translate,omitempty"`
	FillTranslateAnchor interface{} `json:"fill-translate-anchor,omitempty"`
	LineBlur            interface{} `json:"line-blur,omitempty"`
	LineColor           interface{} `json:"line-color,omitempty"`
	LineGapWidth        interface{} `json:"line-gap-width,omitempty"`
	LineOpacity         interface{} `json:"line-opacity,omitempty"`
	LineWidth           interface{} `json:"line-width,omitempty"`
	LinePattern         interface{} `json:"line-pattern,omitempty"`
	LineDasharray       interface{} `json:"line-dasharray,omitempty"`
	CircleColor         interface{} `json:"circle-color,omitempty"`
	CircleOpacity       interface{} `json:"circle-opacity,omitempty"`
	CircleRadius        interface{} `json:"circle-radius,omitempty"`
	TextColor           interface{} `json:"text-color,omitempty"`
	TextHaloWidth       interface{} `json:"text-halo-width,omitempty"`
	TextHaloBlur        interface{} `json:"text-halo-blur,omitempty"`
	TextHaloColor       interface{} `json:"text-halo-color,omitempty"`
	TextOpacity         interface{} `json:"text-opacity,omitempty"`
	TextTranslate       interface{} `json:"text-translate,omitempty"`
	TextTranslateAnchor interface{} `json:"text-translate-anchor,omitempty"`
}

type Layout struct {
	Visibility          interface{} `json:"visibility,omitempty"`
	FillSortKey         interface{} `json:"fill-sort-key,omitempty"`
	LineCap             interface{} `json:"line-cap,omitempty"`
	LineJoin            interface{} `json:"line-join,omitempty"`
	TextAllowOverlap    interface{} `json:"text-allow-overlap,omitempty"`
	TextField           interface{} `json:"text-field,omitempty"`
	TextFont            interface{} `json:"text-font,omitempty"`
	TextJustify         interface{} `json:"text-justify,omitempty"`
	TextLineHeight      interface{} `json:"text-line-height,omitempty"`
	TextMaxAngle        interface{} `json:"text-max-angle,omitempty"`
	TextMaxWidth        interface{} `json:"text-max-width,omitempty"`
	TextOffset          interface{} `json:"text-offset,omitempty"`
	TextIgnorePlacement interface{} `json:"text-ignore-placement,omitempty"`
	TextPadding         interface{} `json:"text-padding,omitempty"`
	TextRotate          interface{} `json:"text-rotate,omitempty"`
	TextSize            interface{} `json:"text-size,omitempty"`
	TextTransform       string      `json:"text-transform,omitempty"`
}

type MapLayer struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Filter      interface{}            `json:"filter,omitempty"`
	Layout      Layout                 `json:"layout"`
	Paint       Paint                  `json:"paint"`
	Source      string                 `json:"source"`
	SourceLayer string                 `json:"source-layer,omitempty"`
	MaxZoom     interface{}            `json:"maxzoom,omitempty"`
	MinZoom     interface{}            `json:"minzoom,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Label       string                 `json:"label,omitempty"`
}

type StyledMap struct {
	ID       string                `json:"id"`
	Version  int32                 `json:"version"`
	Name     string                `json:"name"`
	Sprite   string                `json:"sprite"`
	Glyphs   string                `json:"glyphs"`
	Sources  map[string]*pb.Source `json:"sources"`
	Layers   []*pb.LayerMapbox     `json:"layers"`
	Bearing  *int32                `json:"bearing,omitempty"`
	Zoom     *float32              `json:"zoom,omitempty"`
	Metadata map[string]string     `json:"metadata,omitempty"`
	Pitch    *int32                `json:"pitch,omitempty"`
}

type PgStyledMapModel struct {
	MapID                     string  `db:"map_id"`
	MapName                   string  `db:"map_name"`
	GroupOrder                int32   `db:"group_order"`
	GroupID                   string  `db:"group_id"`
	GroupName                 string  `db:"group_name"`
	LayerID                   string  `db:"layer_id"`
	LayerName                 string  `db:"layer_name"`
	LayerType                 string  `db:"layer_type"`
	LayerTableID              string  `db:"layer_table_id"`
	LayerTableName            string  `db:"layer_table_name"`
	LayerOrder                int32   `db:"layer_order"`
	StyleID                   string  `db:"style_id"`
	StyleName                 string  `db:"style_name"`
	StyleType                 string  `db:"StyleType"`
	StyleSourceLayer          string  `db:"style_source_layer"`
	StyleFilter1              bool    `db:"style_filter1"`
	StyleFilter2              bool    `db:"style_filter2"`
	StyleFilterField1         string  `db:"style_filter_field1"`
	StyleFilterField2         string  `db:"style_filter_field2"`
	StyleFilterValues         string  `db:"style_filter_values"`
	StyleMaxZoom              float32 `db:"style_max_zoom"`
	StyleMinZoom              float32 `db:"style_min_zoom"`
	StyleLabel                bool    `db:"style_label"`
	LabelTextColor            string  `db:"label_text_color"`
	LabelTextHaloWidth        int32   `db:"label_text_halo_width"`
	LabelTextHaloBlur         int32   `db:"label_text_halo_blur"`
	LabelTextHaloColor        string  `db:"label_text_halo_color"`
	LabelTextField            string  `db:"label_text_field"`
	LabelTextFont             string  `db:"label_text_font"`
	LabelTextOffset           string  `db:"label_text_offset"`
	LabelTextOpacity          string  `db:"label_text_opacity"`
	LabelTextJustify          string  `db:"label_text_justify"`
	LabelTextLineHeight       float32 `db:"label_text_line_height"`
	LabelTextIgnorePlacement  bool    `db:"label_text_ignore_placement"`
	LabelTextPadding          int32   `db:"label_text_padding"`
	LabelTextRotate           float32 `db:"label_text_rotate"`
	LabelTextSize             float32 `db:"label_text_size"`
	LabelTextTransform        string  `db:"label_text_transform"`
	FillAntialias             string  `db:"fill_antialias"`
	FillColor                 string  `db:"fill_color"`
	FillOpacity               string  `db:"fill_opacity"`
	FillOutlineColor          string  `db:"fill_outline_color"`
	FillPattern               string  `db:"fill_pattern"`
	FillVisibility            string  `db:"fill_visibility"`
	LineBlur                  string  `db:"line_blur"`
	LineColor                 string  `db:"line_color"`
	LineGapWidth              string  `db:"line_gap_width"`
	LineOpacity               string  `db:"line_opacity"`
	LineWidth                 string  `db:"line_width"`
	LinePattern               string  `db:"line_pattern"`
	LineDasharray             string  `db:"line_dasharray"`
	LineCap                   string  `db:"line_cap"`
	LineJoin                  string  `db:"line_join"`
	LineVisibility            string  `db:"line_visibility"`
	SymbolTextAllowOverlap    string  `db:"symbol_text_allow_overlap"`
	SymbolTextColor           string  `db:"symbol_text_color"`
	SymbolTextField           string  `db:"symbol_text_field"`
	SymbolTextFont            string  `db:"symbol_text_font"`
	SymbolTextHaloBlur        string  `db:"symbol_text_halo_blur"`
	SymbolTextHaloColor       string  `db:"symbol_text_halo_color"`
	SymbolTextHaloWidth       string  `db:"symbol_text_halo_width"`
	SymbolTextIgnorePlacement string  `db:"symbol_text_ignore_placement"`
	SymbolTextJustify         string  `db:"symbol_text_justify"`
	SymbolTextRotate          string  `db:"symbol_text_rotate"`
	SymbolTextSize            string  `db:"symbol_text_size"`
	SymbolTextOffset          string  `db:"symbol_text_offset"`
	SymbolTextOpacity         string  `db:"symbol_text_opacity"`
}

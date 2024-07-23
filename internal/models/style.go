package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	pb "protos/maps"
	"time"
)

// Style model for style
type Style struct {
	ID string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`

	StyleName                 string  `gorm:"not null" json:"style_name"`
	StyleType                 string  `gorm:"not null" json:"style_type"`
	StyleSourceLayer          string  `gorm:"not null" json:"style_source_layer"`
	StyleFilter1              bool    `gorm:"not null" json:"style_filter_1"`
	StyleFilterField1         string  `json:"style_filter_field_1"`
	StyleFilter2              bool    `gorm:"not null" json:"style_filter_2"`
	StyleFilterField2         string  `json:"style_filter_field_2"`
	StyleFilterValues         string  `gorm:"type:text" json:"style_filter_values"`
	StyleMaxZoom              float32 `gorm:"not null;default:24" json:"style_max_zoom"`
	StyleMinZoom              float32 `gorm:"not null;default:0" json:"style_min_zoom"`
	StyleLabel                bool    `json:"style_label"`
	LabelTextColor            string  `json:"label_text_color"`
	LabelTextHaloWidth        float32 `json:"label_text_halo_width"`
	LabelTextHaloBlur         float32 `json:"label_text_halo_blur"`
	LabelTextHaloColor        string  `json:"label_text_halo_color"`
	LabelTextField            string  `json:"label_text_field"`
	LabelTextFont             string  `json:"label_text_font"`
	LabelTextOffset           string  `gorm:"type:text" json:"label_text_offset"`
	LabelTextOpacity          string  `gorm:"type:text" json:"label_text_opacity"`
	LabelTextJustify          string  `json:"label_text_justify"`
	LabelTextLineHeight       float32 `json:"label_text_line_height"`
	LabelTextIgnorePlacement  bool    `json:"label_text_ignore_placement"`
	LabelTextPadding          int32   `json:"label_text_padding"`
	LabelTextRotate           float32 `json:"label_text_rotate"`
	LabelTextSize             float32 `json:"label_text_size"`
	LabelTextTransform        string  `json:"label_text_transform"`
	FillAntialias             string  `gorm:"type:text;default:'[]'" json:"fill_antialias"`
	FillColor                 string  `gorm:"type:text;default:'[]'" json:"fill_color"`
	FillOpacity               string  `gorm:"type:text;default:'[]'" json:"fill_opacity"`
	FillOutlineColor          string  `gorm:"type:text;default:'[]'" json:"fill_outline_color"`
	FillPattern               string  `gorm:"type:text;default:'[]'" json:"fill_pattern"`
	FillVisibility            string  `gorm:"type:text;default:'[]'" json:"fill_visibility"`
	LineBlur                  string  `gorm:"type:text;default:'[]'" json:"line_blur"`
	LineColor                 string  `gorm:"type:text;default:'[]'" json:"line_color"`
	LineGapWidth              string  `gorm:"type:text;default:'[]'" json:"line_gap_width"`
	LineOpacity               string  `gorm:"type:text;default:'[]'" json:"line_opacity"`
	LineWidth                 string  `gorm:"type:text;default:'[]'" json:"line_width"`
	LinePattern               string  `gorm:"type:text;default:'[]'" json:"line_pattern"`
	LineDasharray             string  `gorm:"type:text;default:'[]'" json:"line_dasharray"`
	LineCap                   string  `gorm:"type:text;default:'[]'" json:"line_cap"`
	LineJoin                  string  `gorm:"type:text;default:'[]'"  json:"line_join"`
	LineVisibility            string  `gorm:"type:text;default:'[]'" json:"line_visibility"`
	SymbolTextAllowOverlap    string  `gorm:"type:text;default:'[]'" json:"symbol_text_allow_overlap"`
	SymbolTextColor           string  `gorm:"type:text;default:'[]'" json:"symbol_text_color"`
	SymbolTextField           string  `gorm:"type:text;default:'[]'" json:"symbol_text_field"`
	SymbolTextFont            string  `gorm:"type:text;default:'[]'" json:"symbol_text_font"`
	SymbolTextHaloBlur        string  `gorm:"type:text;default:'[]'" json:"symbol_text_halo_blur"`
	SymbolTextHaloColor       string  `gorm:"type:text;default:'[]'" json:"symbol_text_halo_color"`
	SymbolTextHaloWidth       string  `gorm:"type:text;default:'[]'" json:"symbol_text_halo_width"`
	SymbolTextIgnorePlacement string  `gorm:"type:text;default:'[]'" json:"symbol_text_ignore_placement"`
	SymbolTextJustify         string  `gorm:"type:text;default:'[]'" json:"symbol_text_justify"`
	SymbolTextRotate          string  `gorm:"type:text;default:'[]'" json:"symbol_text_rotate"`
	SymbolTextSize            string  `gorm:"type:text;default:'[]'" json:"symbol_text_size"`
	SymbolTextOffset          string  `gorm:"type:text;default:'[]'" json:"symbol_text_offset"`
	SymbolTextOpacity         string  `gorm:"type:text;default:'[]'" json:"symbol_text_opacity"`

	CreateUserIP string `json:"create_user_ip"`
	UpdateUserIP string `json:"update_user_ip"`
	CreateUserID string `json:"create_user_id"`
	UpdateUserID string `json:"update_user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// BeforeCreate - create new uuid
func (style *Style) BeforeCreate(tx *gorm.DB) (err error) {
	style.ID = uuid.NewString()

	return
}

type StyleAdapter struct {
	id                        string
	styleFilter1              bool
	styleFilter2              bool
	styleLabel                bool
	labelTextIgnorePlacement  bool
	styleMaxZoom              float32
	styleMinZoom              float32
	labelTextHaloWidth        float32
	labelTextHaloBlur         float32
	labelTextLineHeight       float32
	labelTextRotate           float32
	labelTextSize             float32
	labelTextPadding          int32
	styleName                 string
	styleType                 string
	styleSourceLayer          string
	styleFilterField1         string
	styleFilterField2         string
	styleFilterValues         string
	labelTextColor            string
	labelTextHaloColor        string
	labelTextField            string
	labelTextFont             string
	labelTextOffset           string
	labelTextOpacity          string
	labelTextJustify          string
	labelTextTransform        string
	fillAntialias             string
	fillColor                 string
	fillOpacity               string
	fillOutlineColor          string
	fillPattern               string
	fillVisibility            string
	lineBlur                  string
	lineColor                 string
	lineGapWidth              string
	lineOpacity               string
	lineWidth                 string
	linePattern               string
	lineDasharray             string
	lineCap                   string
	lineJoin                  string
	lineVisibility            string
	symbolTextAllowOverlap    string
	symbolTextColor           string
	symbolTextField           string
	symbolTextFont            string
	symbolTextHaloBlur        string
	symbolTextHaloColor       string
	symbolTextHaloWidth       string
	symbolTextIgnorePlacement string
	symbolTextJustify         string
	symbolTextRotate          string
	symbolTextSize            string
	symbolTextOffset          string
	symbolTextOpacity         string
	createUserIp              string
	createUserID              string
	updateUserIp              string
	updateUserID              string
}

func FromMStyle(ms *pb.MStyle) Style {
	return Style{
		ID:                        ms.Id,
		StyleName:                 ms.StyleName,
		StyleType:                 ms.StyleType,
		StyleSourceLayer:          ms.StyleSourceLayer,
		StyleFilter1:              ms.StyleFilter_1,
		StyleFilterField1:         ms.StyleFilterField_1,
		StyleFilter2:              ms.StyleFilter_2,
		StyleFilterField2:         ms.StyleFilterField_2,
		StyleFilterValues:         ms.StyleFilterValues,
		StyleMaxZoom:              ms.StyleMaxZoom,
		StyleMinZoom:              ms.StyleMinZoom,
		StyleLabel:                ms.StyleLabel,
		LabelTextColor:            ms.LabelTextColor,
		LabelTextHaloWidth:        ms.LabelTextHaloWidth,
		LabelTextHaloBlur:         ms.LabelTextHaloBlur,
		LabelTextHaloColor:        ms.LabelTextHaloColor,
		LabelTextField:            ms.LabelTextField,
		LabelTextFont:             ms.LabelTextFont,
		LabelTextOffset:           ms.LabelTextOffset,
		LabelTextOpacity:          ms.LabelTextOpacity,
		LabelTextJustify:          ms.LabelTextJustify,
		LabelTextLineHeight:       ms.LabelTextLineHeight,
		LabelTextIgnorePlacement:  ms.LabelTextIgnorePlacement,
		LabelTextPadding:          ms.LabelTextPadding,
		LabelTextRotate:           ms.LabelTextRotate,
		LabelTextSize:             ms.LabelTextSize,
		LabelTextTransform:        ms.LabelTextTransform,
		FillAntialias:             ms.FillAntialias,
		FillColor:                 ms.FillColor,
		FillOpacity:               ms.FillOpacity,
		FillOutlineColor:          ms.FillOutlineColor,
		FillPattern:               ms.FillPattern,
		FillVisibility:            ms.FillVisibility,
		LineBlur:                  ms.LineBlur,
		LineColor:                 ms.LineColor,
		LineGapWidth:              ms.LineGapWidth,
		LineOpacity:               ms.LineOpacity,
		LineWidth:                 ms.LineWidth,
		LinePattern:               ms.LinePattern,
		LineDasharray:             ms.LineDasharray,
		LineCap:                   ms.LineCap,
		LineJoin:                  ms.LineJoin,
		LineVisibility:            ms.LineVisibility,
		SymbolTextAllowOverlap:    ms.SymbolTextAllowOverlap,
		SymbolTextColor:           ms.SymbolTextColor,
		SymbolTextField:           ms.SymbolTextField,
		SymbolTextFont:            ms.SymbolTextFont,
		SymbolTextHaloBlur:        ms.SymbolTextHaloBlur,
		SymbolTextHaloColor:       ms.SymbolTextHaloColor,
		SymbolTextHaloWidth:       ms.SymbolTextHaloWidth,
		SymbolTextIgnorePlacement: ms.SymbolTextIgnorePlacement,
		SymbolTextJustify:         ms.SymbolTextJustify,
		SymbolTextRotate:          ms.SymbolTextRotate,
		SymbolTextSize:            ms.SymbolTextSize,
		SymbolTextOffset:          ms.SymbolTextOffset,
		SymbolTextOpacity:         ms.SymbolTextOpacity,
		CreateUserIP:              ms.CreateUserIp,
		CreateUserID:              ms.CreateUserId,
		UpdateUserID:              ms.UpdateUserId,
		UpdateUserIP:              ms.UpdateUserIp,
	}
}

// Protobuf2StyleAdapter convert protobuf style message to models.StyleAdapter
func Protobuf2StyleAdapter(r2 *pb.MStyle) *StyleAdapter {
	return &StyleAdapter{
		id:                        r2.Id,
		styleFilter1:              r2.StyleFilter_1,
		styleFilter2:              r2.StyleFilter_2,
		styleLabel:                r2.StyleLabel,
		labelTextIgnorePlacement:  r2.LabelTextIgnorePlacement,
		styleMaxZoom:              r2.StyleMaxZoom,
		styleMinZoom:              r2.StyleMinZoom,
		labelTextHaloWidth:        r2.LabelTextHaloWidth,
		labelTextHaloBlur:         r2.LabelTextHaloBlur,
		labelTextLineHeight:       r2.LabelTextSize,
		labelTextRotate:           r2.LabelTextRotate,
		labelTextSize:             r2.LabelTextSize,
		labelTextPadding:          r2.LabelTextPadding,
		styleName:                 r2.StyleName,
		styleType:                 r2.StyleType,
		styleSourceLayer:          r2.StyleSourceLayer,
		styleFilterField1:         r2.StyleFilterField_1,
		styleFilterField2:         r2.StyleFilterField_2,
		styleFilterValues:         r2.StyleFilterValues,
		labelTextColor:            r2.LabelTextColor,
		labelTextHaloColor:        r2.LabelTextHaloColor,
		labelTextField:            r2.LabelTextField,
		labelTextFont:             r2.LabelTextFont,
		labelTextOffset:           r2.LabelTextOffset,
		labelTextOpacity:          r2.LabelTextOpacity,
		labelTextJustify:          r2.LabelTextJustify,
		labelTextTransform:        r2.LabelTextTransform,
		fillAntialias:             r2.FillAntialias,
		fillColor:                 r2.FillColor,
		fillOpacity:               r2.FillOpacity,
		fillOutlineColor:          r2.FillOutlineColor,
		fillPattern:               r2.FillPattern,
		fillVisibility:            r2.FillVisibility,
		lineBlur:                  r2.LineBlur,
		lineColor:                 r2.LineColor,
		lineGapWidth:              r2.LineGapWidth,
		lineOpacity:               r2.LineOpacity,
		lineWidth:                 r2.LineWidth,
		linePattern:               r2.LinePattern,
		lineDasharray:             r2.LineDasharray,
		lineCap:                   r2.LineCap,
		lineJoin:                  r2.LineJoin,
		lineVisibility:            r2.LineVisibility,
		symbolTextAllowOverlap:    r2.SymbolTextAllowOverlap,
		symbolTextColor:           r2.SymbolTextColor,
		symbolTextField:           r2.SymbolTextField,
		symbolTextFont:            r2.SymbolTextFont,
		symbolTextHaloBlur:        r2.SymbolTextHaloBlur,
		symbolTextHaloColor:       r2.SymbolTextHaloColor,
		symbolTextHaloWidth:       r2.SymbolTextHaloWidth,
		symbolTextIgnorePlacement: r2.SymbolTextIgnorePlacement,
		symbolTextJustify:         r2.SymbolTextJustify,
		symbolTextRotate:          r2.SymbolTextRotate,
		symbolTextSize:            r2.SymbolTextSize,
		symbolTextOffset:          r2.SymbolTextOffset,
		symbolTextOpacity:         r2.SymbolTextOpacity,
		updateUserIp:              r2.UpdateUserIp,
		updateUserID:              r2.UpdateUserId,
		createUserID:              r2.CreateUserId,
		createUserIp:              r2.CreateUserIp,
	}
}

func Style2Protobuf(s Style) *pb.MStyle {
	return &pb.MStyle{
		Id:                        s.ID,
		StyleName:                 s.StyleName,
		StyleType:                 s.StyleType,
		StyleSourceLayer:          s.StyleSourceLayer,
		StyleFilter_1:             s.StyleFilter1,
		StyleFilterField_1:        s.StyleFilterField1,
		StyleFilter_2:             s.StyleFilter2,
		StyleFilterField_2:        s.StyleFilterField2,
		StyleFilterValues:         s.StyleFilterValues,
		StyleMaxZoom:              s.StyleMaxZoom,
		StyleMinZoom:              s.StyleMinZoom,
		StyleLabel:                s.StyleLabel,
		LabelTextColor:            s.LabelTextColor,
		LabelTextHaloWidth:        s.LabelTextHaloWidth,
		LabelTextHaloBlur:         s.LabelTextHaloBlur,
		LabelTextHaloColor:        s.LabelTextHaloColor,
		LabelTextField:            s.LabelTextField,
		LabelTextFont:             s.LabelTextFont,
		LabelTextOffset:           s.LabelTextOffset,
		LabelTextOpacity:          s.LabelTextOpacity,
		LabelTextJustify:          s.LabelTextJustify,
		LabelTextLineHeight:       s.LabelTextLineHeight,
		LabelTextIgnorePlacement:  s.LabelTextIgnorePlacement,
		LabelTextPadding:          s.LabelTextPadding,
		LabelTextRotate:           s.LabelTextRotate,
		LabelTextSize:             s.LabelTextSize,
		LabelTextTransform:        s.LabelTextTransform,
		FillAntialias:             s.FillAntialias,
		FillColor:                 s.FillColor,
		FillOpacity:               s.FillOpacity,
		FillOutlineColor:          s.FillOutlineColor,
		FillPattern:               s.FillPattern,
		FillVisibility:            s.FillVisibility,
		LineBlur:                  s.LineBlur,
		LineColor:                 s.LineColor,
		LineGapWidth:              s.LineGapWidth,
		LineOpacity:               s.LineOpacity,
		LineWidth:                 s.LineWidth,
		LinePattern:               s.LinePattern,
		LineDasharray:             s.LineDasharray,
		LineCap:                   s.LineCap,
		LineJoin:                  s.LineJoin,
		LineVisibility:            s.LineVisibility,
		SymbolTextAllowOverlap:    s.SymbolTextAllowOverlap,
		SymbolTextColor:           s.SymbolTextColor,
		SymbolTextField:           s.SymbolTextField,
		SymbolTextFont:            s.SymbolTextFont,
		SymbolTextHaloBlur:        s.SymbolTextHaloBlur,
		SymbolTextHaloColor:       s.SymbolTextHaloColor,
		SymbolTextHaloWidth:       s.SymbolTextHaloWidth,
		SymbolTextIgnorePlacement: s.SymbolTextIgnorePlacement,
		SymbolTextJustify:         s.SymbolTextJustify,
		SymbolTextRotate:          s.SymbolTextRotate,
		SymbolTextSize:            s.SymbolTextSize,
		SymbolTextOffset:          s.SymbolTextOffset,
		SymbolTextOpacity:         s.SymbolTextOpacity,
		CreateUserId:              s.CreateUserID,
		UpdateUserId:              s.UpdateUserID,
		CreateUserIp:              s.CreateUserIP,
		UpdateUserIp:              s.UpdateUserIP,
		CreatedAt:                 s.CreatedAt.String(),
		UpdatedAt:                 s.UpdatedAt.String(),
	}
}

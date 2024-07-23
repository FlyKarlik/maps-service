package repository

import (
	"fmt"
	"strings"
)

const (
	layerGroupRelationJoins = `	SELECT groups.* FROM group_layer_relations
								LEFT JOIN groups ON group_layer_relations.group_id = groups.id
								WHERE layer_id = ? AND groups.id IS NOT NULL`
	groupRelationLayersJoinLayer = `	SELECT layers.* FROM group_layer_relations 
										LEFT JOIN layers ON group_layer_relations.layer_id = layers.id
										WHERE group_id = ? AND layers.id IS NOT NULL`
	mapRelationGroupsJoinGroup = `	SELECT groups.* from map_group_relations
									LEFT JOIN groups on map_group_relations.group_id = groups.id 
	                            	WHERE map_id = ? AND groups.id IS NOT NULL`
	groupRelationMapsJoinMap = `SELECT maps.* from map_group_relations
								LEFT JOIN maps on map_group_relations.map_id = maps.id 
								WHERE group_id = ? AND maps.id IS NOT NULL`
	layerRelationStylesJoinStyle = `SELECT styles.* FROM layer_style_relations
									LEFT JOIN styles ON layer_style_relations.style_id = styles.id
									WHERE layer_id = ? AND styles.id IS NOT NULL LIMIT 1`
	styleRelationLayersJoinLayer = `SELECT layers.* from layer_style_relations
 									LEFT JOIN layers ON layer_style_relations.layer_id = layers.id
 									WHERE style_id = ? AND layers.id IS NOT NULL`

	styledMap = `
	SELECT maps.id AS map_id,
    maps.name AS map_name,
    map_group_relations.group_order,
    groups.id AS group_id,
    groups.name AS group_name,
    layers.id AS layer_id,
    layers.name AS layer_name,
    layers.layer_type,
    layers.table_id AS layer_table_id,
    tables.name AS layer_table_name,
    group_layer_relations.layer_order,
    styles.id AS style_id,
    styles.style_name,
    styles.style_type,
    styles.style_source_layer,
    styles.style_filter1,
    styles.style_filter_field1,
    styles.style_filter2,
    styles.style_filter_field2,
    styles.style_filter_values,
    styles.style_max_zoom,
    styles.style_min_zoom,
    styles.style_label,
    styles.label_text_color,
    styles.label_text_halo_width,
    styles.label_text_halo_blur,
    styles.label_text_halo_color,
    styles.label_text_field,
    styles.label_text_justify,
    styles.label_text_line_height,
    styles.label_text_ignore_placement,
    styles.label_text_padding,
    styles.label_text_rotate,
    styles.label_text_size,
    styles.label_text_transform,
    styles.fill_antialias,
    styles.fill_color,
    styles.fill_opacity,
    styles.fill_outline_color,
    styles.fill_pattern,
    styles.fill_visibility,
    styles.line_blur,
    styles.line_color,
    styles.line_gap_width,
    styles.line_opacity,
    styles.line_width,
    styles.line_pattern,
    styles.line_dasharray,
    styles.line_cap,
    styles.line_join,
    styles.line_visibility,
    styles.symbol_text_allow_overlap,
    styles.symbol_text_color,
    styles.symbol_text_field,
    styles.symbol_text_font,
    styles.symbol_text_halo_blur,
    styles.symbol_text_halo_color,
    styles.symbol_text_halo_width,
    styles.symbol_text_ignore_placement,
    styles.symbol_text_justify,
    styles.symbol_text_rotate,
    styles.symbol_text_size,
    styles.symbol_text_offset,
    styles.symbol_text_opacity,
    styles.label_text_font,
    styles.label_text_offset,
    styles.label_text_opacity
   FROM maps
     LEFT JOIN map_group_relations ON maps.id = map_group_relations.map_id
     LEFT JOIN groups ON groups.id = map_group_relations.group_id
     LEFT JOIN group_layer_relations ON group_layer_relations.group_id = groups.id
     LEFT JOIN layers ON layers.id = group_layer_relations.layer_id
     LEFT JOIN tables ON tables.id = layers.table_id
     LEFT JOIN layer_style_relations ON layer_style_relations.layer_id = layers.id
     LEFT JOIN styles ON styles.id = layer_style_relations.style_id
  WHERE maps.id = ?;
`
)

func GetTableFeatureIntersectWithPolygon(tableName string, xMin, yMin, xMax, yMax float32) (string, error) {
	q := fmt.Sprintf("SELECT (to_jsonb(%s)- 'geom'::text) #>> '{}' AS attributes, to_jsonb(ST_Transform(%s.geom, 4326)) #>> '{}' AS geometry FROM ippd.%s WHERE ST_Intersects(ST_Transform(geom, 4326), ST_MakeEnvelope(%f,%f,%f,%f, 4326)) limit 1", tableName, tableName, tableName, xMin, yMin, xMax, yMax)
	if strings.Contains(q, "DROP") || strings.Contains(q, "DELETE") || strings.Contains(q, "TRUNCATE") || strings.Contains(q, "INSERT") || strings.Contains(q, "UPDATE") {
		return "", fmt.Errorf("statement error")
	}

	return q, nil
}

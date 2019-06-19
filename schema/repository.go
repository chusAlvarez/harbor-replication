package schema

type Repository struct {
	Id           int32    `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	ProjectID    int32    `json:"project_id,omitempty"`
	Description  string   `json:"description,omitempty"`
	pullCount    int32    `json:"pull_count,omitempty"`
	starCount    int32    `json:"star_count,omitempty"`
	tagsCount    int32    `json:"tags_count,omitempty"`
	labels       []string `json:"labels,omitempty"`
	creationTime string   `json:"creation_time,omitempty"`
	updateTime   string   `json:"update_time,omitempty"`
}

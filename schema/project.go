package schema

type Chart struct {
}

type Search struct {
	Projects     []Project    `json:"project,omitempty"`
	Repositories []Repository `json:"repository,omitempty"`
	Chartts      []Chart      `json:"chart,omitempty"`
}

type Project struct {
	// Project ID
	ProjectId int32 `json:"project_id,omitempty"`
	// The owner ID of the project always means the creator of the project.
	OwnerId int32 `json:"owner_id,omitempty"`
	// The name of the project.
	Name string `json:"name,omitempty"`
	// The creation time of the project.
	CreationTime string `json:"creation_time,omitempty"`
	// The update time of the project.
	UpdateTime string `json:"update_time,omitempty"`
	// A deletion mark of the project.
	Deleted bool `json:"deleted,omitempty"`
	// The owner name of the project.
	OwnerName string `json:"owner_name,omitempty"`
	// Correspond to the UI about whether the project's publicity is  updatable (for UI)
	Togglable bool `json:"togglable,omitempty"`
	// The role ID of the current user who triggered the API (for UI)
	CurrentUserRoleId int32 `json:"current_user_role_id,omitempty"`
	// The number of the repositories under this project.
	RepoCount int32 `json:"repo_count,omitempty"`
	// The total number of charts under this project.
	ChartCount int32 `json:"chart_count,omitempty"`
	// The metadata of the project.
	Metadata *ProjectMetadata `json:"metadata,omitempty"`
}

type ProjectMetadata struct {
	// The public status of the project. The valid values are \"true\", \"false\".
	Public string `json:"public,omitempty"`
	// Whether content trust is enabled or not. If it is enabled, user cann't pull unsigned images from this project. The valid values are \"true\", \"false\".
	EnableContentTrust string `json:"enable_content_trust,omitempty"`
	// Whether prevent the vulnerable images from running. The valid values are \"true\", \"false\".
	PreventVul string `json:"prevent_vul,omitempty"`
	// If the vulnerability is high than severity defined here, the images cann't be pulled. The valid values are \"negligible\", \"low\", \"medium\", \"high\", \"critical\".
	Severity string `json:"severity,omitempty"`
	// Whether scan images automatically when pushing. The valid values are \"true\", \"false\".
	AutoScan string `json:"auto_scan,omitempty"`
}

type ProjectReq struct {
	// The name of the project.
	ProjectName string `json:"project_name,omitempty"`
	// The metadata of the project.
	Metadata *ProjectMetadata `json:"metadata,omitempty"`
}

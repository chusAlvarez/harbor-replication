package schema

type Tag struct {
	Digest        string          `json:"digest,omitempty"`
	Name          string          `json:"name,omitempty"`
	Size          int32           `json:"size,omitempty"`
	Architecture  string          `json:"architecture,omitempty"`
	Os            string          `json:"os,omitempty"`
	DockerVersion string          `json:"docker_version,omitempty"`
	Author        string          `json:"author,omitempty"`
	Created       string          `json:"created,omitempty"`
	Config        TagConfig       `json:"config,omitempty"`
	Signature     TagSignature    `json:"signature,omitempty"`
	ScanOverview  TagScanOverview `json:"scan_overview,omitempty"`
	Labels        []TagLabel      `json:"labels,omitempty"`
}

type TagSignature struct {
}

type TagConfigLabels struct {
}

type TagConfig struct {
	Labels TagConfigLabels `json:"labels,omitempty"`
}

type TagScanOverview struct {
	ImageDigest  string         `json:"image_digest,omitempty"`
	ScanStatus   string         `json:"scan_status,omitempty"`
	JobID        int32          `json:"job_id,omitempty"`
	Severity     int32          `json:"severity,omitempty"`
	Components   ScanComponents `json:"components,omitempty"`
	DetailsKey   string         `json:"details_key,omitempty"`
	CreationTime string         `json:"creation_time,omitempty"`
	UpdateTime   string         `json:"update_time,omitempty"`
}

type ScanComponents struct {
	Total   int32         `json:"total,omitempty"`
	Summary []ScanSummary `json:"summary,omitempty"`
}

type ScanSummary struct {
	Severity int32 `json:"severity,omitempty"`
	Count    int32 `json:"count,omitempty"`
}

type TagLabel struct {
	ID           int32  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Color        string `json:"color,omitempty"`
	Scope        string `json:"scope,omitempty"`
	ProjectID    int32  `json:"project_id,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
	UpdateTime   string `json:"update_time,omitempty"`
	Deleted      bool   `json:"deleted,omitempty"`
}

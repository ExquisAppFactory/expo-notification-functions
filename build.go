package expo-notification-functions

import "time"

type BuildDetails struct {
	ID                    string    `json:"id"`
	AccountName           string    `json:"accountName"`
	ProjectName           string    `json:"projectName"`
	BuildDetailsPageUrl   string    `json:"buildDetailsPageUrl"`
	ParentBuildId         string    `json:"parentBuildId"`
	AppID                 string    `json:"appId"`
	InitiatingUserId      string    `json:"initiatingUserId"`
	CancelingUserId       *string   `json:"cancelingUserId"` // Using pointer to represent null
	Platform              string    `json:"platform"`
	Status                string    `json:"status"`
	Artifacts             Artifacts `json:"artifacts"`
	Metadata              Metadata  `json:"metadata"`
	Metrics               Metrics   `json:"metrics"`
	Error                 Error     `json:"error"`
	CreatedAt             time.Time `json:"createdAt"`
	EnqueuedAt            time.Time `json:"enqueuedAt"`
	ProvisioningStartedAt time.Time `json:"provisioningStartedAt"`
	WorkerStartedAt       time.Time `json:"workerStartedAt"`
	CompletedAt           time.Time `json:"completedAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	ExpirationDate        time.Time `json:"expirationDate"`
	Priority              string    `json:"priority"`
	ResourceClass         string    `json:"resourceClass"`
	ActualResourceClass   string    `json:"actualResourceClass"`
	MaxRetryTimeMinutes   int       `json:"maxRetryTimeMinutes"`
}

type Artifacts struct {
	BuildUrl        string `json:"buildUrl"`
	LogsS3KeyPrefix string `json:"logsS3KeyPrefix"`
}

type TrackingContext struct {
	Platform         string `json:"platform"`
	AccountID        string `json:"account_id"`
	DevClient        bool   `json:"dev_client"`
	ProjectID        string `json:"project_id"`
	TrackingID       string `json:"tracking_id"`
	ProjectType      string `json:"project_type"`
	DevClientVersion string `json:"dev_client_version"`
}

type Metadata struct {
	AppName               string          `json:"appName"`
	Username              string          `json:"username"`
	Workflow              string          `json:"workflow"`
	AppVersion            string          `json:"appVersion"`
	AppBuildVersion       string          `json:"appBuildVersion"`
	CliVersion            string          `json:"cliVersion"`
	SdkVersion            string          `json:"sdkVersion"`
	BuildProfile          string          `json:"buildProfile"`
	Distribution          string          `json:"distribution"`
	AppIdentifier         string          `json:"appIdentifier"`
	GitCommitHash         string          `json:"gitCommitHash"`
	GitCommitMessage      string          `json:"gitCommitMessage"`
	RuntimeVersion        string          `json:"runtimeVersion"`
	Channel               string          `json:"channel"`
	ReleaseChannel        string          `json:"releaseChannel"`
	ReactNativeVersion    string          `json:"reactNativeVersion"`
	TrackingContext       TrackingContext `json:"trackingContext"`
	CredentialsSource     string          `json:"credentialsSource"`
	IsGitWorkingTreeDirty bool            `json:"isGitWorkingTreeDirty"`
	Message               string          `json:"message"`
	RunFromCI             bool            `json:"runFromCI"`
}

type Metrics struct {
	Memory                   int64   `json:"memory"`
	BuildEndTimestamp        int64   `json:"buildEndTimestamp"`
	TotalDiskReadBytes       int     `json:"totalDiskReadBytes"`
	BuildStartTimestamp      int64   `json:"buildStartTimestamp"`
	TotalDiskWriteBytes      int     `json:"totalDiskWriteBytes"`
	CpuActiveMilliseconds    float64 `json:"cpuActiveMilliseconds"`
	BuildEnqueuedTimestamp   int64   `json:"buildEnqueuedTimestamp"`
	TotalNetworkEgressBytes  int     `json:"totalNetworkEgressBytes"`
	TotalNetworkIngressBytes int64   `json:"totalNetworkIngressBytes"`
}

type Error struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

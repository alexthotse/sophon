package shared

import (
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/shopspring/decimal"
)

type Org struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	IsTrial            bool   `json:"isTrial"`
	AutoAddDomainUsers bool   `json:"autoAddDomainUsers"`

	IntegratedModelsMode bool `json:"integratedModelsMode,omitempty"`
}

type User struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	NumNonDraftPlans int    `json:"numNonDraftPlans"`

	DefaultPlanConfig *PlanConfig `json:"defaultPlanConfig,omitempty"`
}

type OrgUser struct {
	OrgId     string `json:"orgId"`
	UserId    string `json:"userId"`
	OrgRoleId string `json:"orgRoleId"`

	Config *OrgUserConfig `json:"config,omitempty"`
}

type Invite struct {
	Id         string     `json:"id"`
	OrgId      string     `json:"orgId"`
	Email      string     `json:"email"`
	Name       string     `json:"name"`
	OrgRoleId  string     `json:"orgRoleId"`
	InviterId  string     `json:"inviterId"`
	InviteeId  *string    `json:"inviteeId"`
	AcceptedAt *time.Time `json:"acceptedAt"`
	CreatedAt  time.Time  `json:"createdAt"`
}

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Plan struct {
	Id              string      `json:"id"`
	OwnerId         string      `json:"ownerId"`
	ProjectId       string      `json:"projectId"`
	Name            string      `json:"name"`
	SharedWithOrgAt *time.Time  `json:"sharedWithOrgAt,omitempty"`
	TotalReplies    int         `json:"totalReplies"`
	ActiveBranches  int         `json:"activeBranches"`
	PlanConfig      *PlanConfig `json:"planConfig,omitempty"`
	ArchivedAt      *time.Time  `json:"archivedAt,omitempty"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
}

type Branch struct {
	Id              string     `json:"id"`
	PlanId          string     `json:"planId"`
	OwnerId         string     `json:"ownerId"`
	ParentBranchId  *string    `json:"parentBranchId"`
	Name            string     `json:"name"`
	Status          PlanStatus `json:"status"`
	ContextTokens   int        `json:"contextTokens"`
	ConvoTokens     int        `json:"convoTokens"`
	SharedWithOrgAt *time.Time `json:"sharedWithOrgAt,omitempty"`
	ArchivedAt      *time.Time `json:"archivedAt,omitempty"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
}

type ContextType string

const (
	ContextFileType          ContextType = "file"
	ContextURLType           ContextType = "url"
	ContextNoteType          ContextType = "note"
	ContextDirectoryTreeType ContextType = "directory tree"
	ContextPipedDataType     ContextType = "piped data"
	ContextImageType         ContextType = "image"
	ContextMapType           ContextType = "map"
)

type FileMapBodies map[string]string

type Context struct {
	Id              string                `json:"id"`
	OwnerId         string                `json:"ownerId"`
	ContextType     ContextType           `json:"contextType"`
	Name            string                `json:"name"`
	Url             string                `json:"url"`
	FilePath        string                `json:"file_path"`
	Sha             string                `json:"sha"`
	NumTokens       int                   `json:"numTokens"`
	Body            string                `json:"body,omitempty"`
	BodySize        int64                 `json:"bodySize,omitempty"`
	ForceSkipIgnore bool                  `json:"forceSkipIgnore"`
	ImageDetail     openai.ImageURLDetail `json:"imageDetail,omitempty"`
	MapParts        FileMapBodies         `json:"mapParts,omitempty"`
	MapShas         map[string]string     `json:"mapShas,omitempty"`
	MapTokens       map[string]int        `json:"mapTokens,omitempty"`
	MapSizes        map[string]int64      `json:"mapSizes,omitempty"`
	AutoLoaded      bool                  `json:"autoLoaded"`
	CreatedAt       time.Time             `json:"createdAt"`
	UpdatedAt       time.Time             `json:"updatedAt"`
}

type TellStage string

const (
	TellStagePlanning       TellStage = "planning"
	TellStageImplementation TellStage = "implementation"
)

type PlanningPhase string

const (
	PlanningPhaseContext PlanningPhase = "context"
	PlanningPhaseTasks   PlanningPhase = "tasks"
)

type CurrentStage struct {
	TellStage     TellStage
	PlanningPhase PlanningPhase
}

type ConvoMessageFlags struct {
	DidMakePlan           bool `json:"didMakePlan"`
	DidRemoveTasks        bool `json:"didRemoveTasks"`
	DidMakeDebuggingPlan  bool `json:"didMakeDebuggingPlan"`
	DidLoadContext        bool `json:"didLoadContext"`
	CurrentStage          CurrentStage
	IsChat                bool `json:"isChat"`
	DidWriteCode          bool `json:"didWriteCode"`
	DidCompleteTask       bool `json:"didCompleteTask"`
	DidCompletePlan       bool `json:"didCompletePlan"`
	HasUnfinishedSubtasks bool `json:"hasUnfinishedSubtasks"`
	IsApplyDebug          bool `json:"isApplyDebug"`
	IsUserDebug           bool `json:"isUserDebug"`
	HasError              bool `json:"hasError"`
}

type Subtask struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	UsesFiles   []string `json:"usesFiles"`
	IsFinished  bool     `json:"isFinished"`
}

type ConvoMessage struct {
	Id               string            `json:"id"`
	UserId           string            `json:"userId"`
	Role             string            `json:"role"`
	Tokens           int               `json:"tokens"`
	Num              int               `json:"num"`
	Message          string            `json:"message"`
	Stopped          bool              `json:"stopped"`
	Flags            ConvoMessageFlags `json:"flags"`
	Subtask          *Subtask          `json:"subtask,omitempty"`
	AddedSubtasks    []*Subtask        `json:"addedSubtasks,omitempty"`
	RemovedSubtasks  []string          `json:"removedSubtasks,omitempty"`
	ActiveContextIds []string          `json:"activeContextIds"`
	CreatedAt        time.Time         `json:"createdAt"`
}

type ConvoSummary struct {
	Id                          string    `json:"id"`
	LatestConvoMessageCreatedAt time.Time `json:"latestConvoMessageCreatedAt"`
	LatestConvoMessageId        string    `json:"lastestConvoMessageId"`
	Summary                     string    `json:"summary"`
	Tokens                      int       `json:"tokens"`
	NumMessages                 int       `json:"numMessages"`
	CreatedAt                   time.Time `json:"createdAt"`
}

type OperationType string

const (
	OperationTypeFile   OperationType = "file"
	OperationTypeMove   OperationType = "move"
	OperationTypeRemove OperationType = "remove"
	OperationTypeReset  OperationType = "reset"
)

type Operation struct {
	Type        OperationType
	Path        string
	Destination string
	Content     string
	Description string
	ReplyBefore string
	NumTokens   int
}

func (o *Operation) Name() string {
	res := string(o.Type) + " | " + o.Path
	if o.Destination != "" {
		res += " → " + o.Destination
	}
	return res
}

type ConvoMessageDescription struct {
	Id                    string `json:"id"`
	ConvoMessageId        string `json:"convoMessageId"`
	SummarizedToMessageId string `json:"summarizedToMessageId"`
	WroteFiles            bool   `json:"wroteFiles"`
	CommitMsg             string `json:"commitMsg"`
	// Files                 []string        `json:"files"`
	Operations            []*Operation    `json:"operations"`
	DidBuild              bool            `json:"didBuild"`
	BuildPathsInvalidated map[string]bool `json:"buildPathsInvalidated"`
	Error                 string          `json:"error"`
	AppliedAt             *time.Time      `json:"appliedAt,omitempty"`
	CreatedAt             time.Time       `json:"createdAt"`
	UpdatedAt             time.Time       `json:"updatedAt"`
}

type PlanBuild struct {
	Id             string    `json:"id"`
	ConvoMessageId string    `json:"convoMessageId"`
	FilePath       string    `json:"filePath"`
	Error          string    `json:"error"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type Replacement struct {
	Id             string                      `json:"id"`
	Old            string                      `json:"old"`
	Summary        string                      `json:"summary"`
	EntireFile     bool                        `json:"entireFile"`
	New            string                      `json:"new"`
	Failed         bool                        `json:"failed"`
	RejectedAt     *time.Time                  `json:"rejectedAt,omitempty"`
	StreamedChange *StreamedChangeWithLineNums `json:"streamedChange"`
}

type PlanFileResult struct {
	Id                  string         `json:"id"`
	TypeVersion         int            `json:"typeVersion"`
	ReplaceWithLineNums bool           `json:"replaceWithLineNums"`
	ConvoMessageId      string         `json:"convoMessageId"`
	PlanBuildId         string         `json:"planBuildId"`
	Path                string         `json:"path"`
	Content             string         `json:"content"`
	AnyFailed           bool           `json:"anyFailed"`
	AppliedAt           *time.Time     `json:"appliedAt,omitempty"`
	RejectedAt          *time.Time     `json:"rejectedAt,omitempty"`
	Replacements        []*Replacement `json:"replacements"`

	RemovedFile bool `json:"removedFile"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CurrentPlanFiles struct {
	Files           map[string]string    `json:"files"`
	Removed         map[string]bool      `json:"removedByPath"`
	UpdatedAtByPath map[string]time.Time `json:"updatedAtByPath"`
}

type PlanFileResultsByPath map[string][]*PlanFileResult
type PlanResult struct {
	SortedPaths        []string                  `json:"sortedPaths"`
	FileResultsByPath  PlanFileResultsByPath     `json:"fileResultsByPath"`
	Results            []*PlanFileResult         `json:"results"`
	ReplacementsByPath map[string][]*Replacement `json:"replacementsByPath"`
}

type PlanApply struct {
	Id                         string    `json:"id"`
	UserId                     string    `json:"userId"`
	ConvoMessageIds            []string  `json:"convoMessageIds"`
	ConvoMessageDescriptionIds []string  `json:"convoMessageDescriptionIds"`
	PlanFileResultIds          []string  `json:"planFileResultIds"`
	CommitMsg                  string    `json:"commitMsg"`
	CreatedAt                  time.Time `json:"createdAt"`
}

type CurrentPlanState struct {
	PlanResult               *PlanResult                `json:"planResult"`
	CurrentPlanFiles         *CurrentPlanFiles          `json:"currentPlanFiles"`
	ConvoMessageDescriptions []*ConvoMessageDescription `json:"convoMessageDescriptions"`
	PlanApplies              []*PlanApply               `json:"planApplies"`
	ContextsByPath           map[string]*Context        `json:"contextsByPath"`
}

type OrgRole struct {
	Id          string `json:"id"`
	IsDefault   bool   `json:"isDefault"`
	Label       string `json:"label"`
	Description string `json:"description"`
}


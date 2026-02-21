package plan

import (
	"sophon-server/db"
	"sophon-server/hooks"
	"sophon-server/model"
	"sophon-server/types"

	shared "sophon-shared"

	sitter "github.com/smacker/go-tree-sitter"
)

const MaxBuildErrorRetries = 3 // uses semi-exponential backoff so be careful with this

type activeBuildStreamState struct {
	modelStreamId string
	clients       map[string]model.ClientInfo
	authVars      map[string]string
	auth          *types.ServerAuth
	currentOrgId  string
	currentUserId string
	orgUserConfig *shared.OrgUserConfig
	plan          *db.Plan
	branch        string
	settings      *shared.PlanSettings
	modelContext  []*db.Context
	convo         []*db.ConvoMessage
}

type activeBuildStreamFileState struct {
	*activeBuildStreamState
	filePath                   string
	convoMessageId             string
	build                      *db.PlanBuild
	currentPlanState           *shared.CurrentPlanState
	activeBuild                *types.ActiveBuild
	preBuildState              string
	parser                     *sitter.Parser
	language                   shared.Language
	syntaxCheckTimedOut        bool
	preBuildStateSyntaxInvalid bool
	validationNumRetry         int
	wholeFileNumRetry          int
	isNewFile                  bool
	contextPart                *db.Context

	builderRun hooks.DidFinishBuilderRunParams
}

package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sophon-server/handlers"
	"sophon-server/hooks"

	"github.com/gorilla/mux"
)

type SophonHandler func(w http.ResponseWriter, r *http.Request)
type HandleSophon func(router *mux.Router, path string, isStreaming bool, handler SophonHandler) *mux.Route

var HandleSophonFn HandleSophon

func RegisterHandleSophon(fn HandleSophon) {
	HandleSophonFn = fn
}

func EnsureHandleSophon() {
	if HandleSophonFn == nil {
		panic("handleSophonFn is not set")
	}
}

func AddHealthRoutes(r *mux.Router) {
	EnsureHandleSophon()

	HandleSophonFn(r, "/health", false, func(w http.ResponseWriter, r *http.Request) {
		_, apiErr := hooks.ExecHook(hooks.HealthCheck, hooks.HookParams{})
		if apiErr != nil {
			log.Printf("Error in health check hook: %v\n", apiErr)
			http.Error(w, apiErr.Msg, apiErr.Status)
			return
		}
		fmt.Fprint(w, "OK")
	})

	HandleSophonFn(r, "/version", false, func(w http.ResponseWriter, r *http.Request) {
		// Log the host
		host := r.Host
		log.Printf("Host header: %s", host)

		execPath, err := os.Executable()
		if err != nil {
			log.Fatal("Error getting current directory: ", err)
		}
		currentDir := filepath.Dir(execPath)

		// get version from version.txt
		var path string
		if os.Getenv("IS_CLOUD") != "" {
			path = filepath.Join(currentDir, "..", "version.txt")
		} else {
			path = filepath.Join(currentDir, "version.txt")
		}

		bytes, err := os.ReadFile(path)

		if err != nil {
			http.Error(w, "Error getting version", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(bytes))
	})
}

func AddApiRoutes(r *mux.Router) {
	addApiRoutes(r, "")
}

func AddApiRoutesWithPrefix(r *mux.Router, prefix string) {
	addApiRoutes(r, prefix)
}

func AddProxyableApiRoutes(r *mux.Router) {
	addProxyableApiRoutes(r, "")
}

func AddProxyableApiRoutesWithPrefix(r *mux.Router, prefix string) {
	addProxyableApiRoutes(r, prefix)
}

func addApiRoutes(r *mux.Router, prefix string) {
	EnsureHandleSophon()

	HandleSophonFn(r, prefix+"/accounts/email_verifications", false, handlers.CreateEmailVerificationHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/accounts/email_verifications/check_pin", false, handlers.CheckEmailPinHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/accounts/sign_in_codes", false, handlers.CreateSignInCodeHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/accounts/sign_in", false, handlers.SignInHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/accounts/sign_out", false, handlers.SignOutHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/accounts", false, handlers.CreateAccountHandler).Methods("POST")

	HandleSophonFn(r, prefix+"/orgs/session", false, handlers.GetOrgSessionHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/orgs", false, handlers.ListOrgsHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/orgs", false, handlers.CreateOrgHandler).Methods("POST")

	HandleSophonFn(r, prefix+"/users", false, handlers.ListUsersHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/orgs/users/{userId}", false, handlers.DeleteOrgUserHandler).Methods("DELETE")
	HandleSophonFn(r, prefix+"/orgs/roles", false, handlers.ListOrgRolesHandler).Methods("GET")

	HandleSophonFn(r, prefix+"/invites", false, handlers.InviteUserHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/invites/pending", false, handlers.ListPendingInvitesHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/invites/accepted", false, handlers.ListAcceptedInvitesHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/invites/all", false, handlers.ListAllInvitesHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/invites/{inviteId}", false, handlers.DeleteInviteHandler).Methods("DELETE")

	HandleSophonFn(r, prefix+"/projects", false, handlers.CreateProjectHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/projects", false, handlers.ListProjectsHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/projects/{projectId}/set_plan", false, handlers.ProjectSetPlanHandler).Methods("PUT")
	HandleSophonFn(r, prefix+"/projects/{projectId}/rename", false, handlers.RenameProjectHandler).Methods("PUT")

	HandleSophonFn(r, prefix+"/projects/{projectId}/plans/current_branches", false, handlers.GetCurrentBranchByPlanIdHandler).Methods("POST")

	HandleSophonFn(r, prefix+"/plans", false, handlers.ListPlansHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/archive", false, handlers.ListArchivedPlansHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/ps", false, handlers.ListPlansRunningHandler).Methods("GET")

	HandleSophonFn(r, prefix+"/projects/{projectId}/plans", false, handlers.CreatePlanHandler).Methods("POST")

	HandleSophonFn(r, prefix+"/projects/{projectId}/plans", false, handlers.DeleteAllPlansHandler).Methods("DELETE")

	HandleSophonFn(r, prefix+"/plans/{planId}", false, handlers.GetPlanHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/{planId}", false, handlers.DeletePlanHandler).Methods("DELETE")

	HandleSophonFn(r, prefix+"/plans/{planId}/current_plan/{sha}", false, handlers.CurrentPlanHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/current_plan", false, handlers.CurrentPlanHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/apply", false, handlers.ApplyPlanHandler).Methods("PATCH")
	HandleSophonFn(r, prefix+"/plans/{planId}/archive", false, handlers.ArchivePlanHandler).Methods("PATCH")
	HandleSophonFn(r, prefix+"/plans/{planId}/unarchive", false, handlers.UnarchivePlanHandler).Methods("PATCH")

	HandleSophonFn(r, prefix+"/plans/{planId}/rename", false, handlers.RenamePlanHandler).Methods("PATCH")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/reject_all", false, handlers.RejectAllChangesHandler).Methods("PATCH")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/reject_file", false, handlers.RejectFileHandler).Methods("PATCH")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/reject_files", false, handlers.RejectFilesHandler).Methods("PATCH")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/diffs", false, handlers.GetPlanDiffsHandler).Methods("GET")

	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/context", false, handlers.ListContextHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/context", false, handlers.LoadContextHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/context/{contextId}/body", false, handlers.GetContextBodyHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/context", false, handlers.UpdateContextHandler).Methods("PUT")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/context", false, handlers.DeleteContextHandler).Methods("DELETE")

	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/convo", false, handlers.ListConvoHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/rewind", false, handlers.RewindPlanHandler).Methods("PATCH")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/logs", false, handlers.ListLogsHandler).Methods("GET")

	HandleSophonFn(r, prefix+"/plans/{planId}/branches", false, handlers.ListBranchesHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/{planId}/branches/{branch}", false, handlers.DeleteBranchHandler).Methods("DELETE")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/branches", false, handlers.CreateBranchHandler).Methods("POST")

	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/settings", false, handlers.GetSettingsHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/settings", false, handlers.UpdateSettingsHandler).Methods("PUT")

	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/status", false, handlers.GetPlanStatusHandler).Methods("GET")

	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/tell", true, handlers.TellPlanHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/build", true, handlers.BuildPlanHandler).Methods("PATCH")

	HandleSophonFn(r, prefix+"/custom_models", false, handlers.ListCustomModelsHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/custom_models", false, handlers.UpsertCustomModelsHandler).Methods("POST")

	HandleSophonFn(r, prefix+"/custom_models/{modelId}", false, handlers.GetCustomModelHandler).Methods("GET")

	HandleSophonFn(r, prefix+"/custom_providers", false, handlers.ListCustomProvidersHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/custom_providers/{providerId}", false, handlers.GetCustomProviderHandler).Methods("GET")

	HandleSophonFn(r, prefix+"/model_sets", false, handlers.ListModelPacksHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/model_sets", false, handlers.CreateModelPackHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/model_sets/{setId}", false, handlers.UpdateModelPackHandler).Methods("PUT")
	HandleSophonFn(r, prefix+"/default_settings", false, handlers.GetDefaultSettingsHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/default_settings", false, handlers.UpdateDefaultSettingsHandler).Methods("PUT")

	HandleSophonFn(r, prefix+"/file_map", false, handlers.GetFileMapHandler).Methods("POST")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/load_cached_file_map", false, handlers.LoadCachedFileMapHandler).Methods("POST")

	HandleSophonFn(r, prefix+"/plans/{planId}/config", false, handlers.GetPlanConfigHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/plans/{planId}/config", false, handlers.UpdatePlanConfigHandler).Methods("PUT")

	HandleSophonFn(r, prefix+"/default_plan_config", false, handlers.GetDefaultPlanConfigHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/default_plan_config", false, handlers.UpdateDefaultPlanConfigHandler).Methods("PUT")

	HandleSophonFn(r, prefix+"/org_user_config", false, handlers.GetOrgUserConfigHandler).Methods("GET")
	HandleSophonFn(r, prefix+"/org_user_config", false, handlers.UpdateOrgUserConfigHandler).Methods("PUT")
}

func addProxyableApiRoutes(r *mux.Router, prefix string) {
	EnsureHandleSophon()

	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/connect", true, handlers.ConnectPlanHandler).Methods("PATCH")
	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/stop", false, handlers.StopPlanHandler).Methods("DELETE")

	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/respond_missing_file", false, handlers.RespondMissingFileHandler).Methods("POST")

	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/auto_load_context", false, handlers.AutoLoadContextHandler).Methods("POST")

	HandleSophonFn(r, prefix+"/plans/{planId}/{branch}/build_status", false, handlers.GetBuildStatusHandler).Methods("GET")
}

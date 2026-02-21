sdx-1: package cmd
sdx-2:
sdx-3: import (
sdx-4: 	"fmt"
sdx-5: 	"path/filepath"
sdx-6: 	"sophon/api"
sdx-7: 	"sophon/auth"
sdx-8: 	"sophon/lib"
sdx-9: 	"sophon/term"
sdx-10:
sdx-11: 	"sophon-shared"
sdx-12: 	"github.com/spf13/cobra"
sdx-13: )
sdx-14:
sdx-15: var contextRmCmd = &cobra.Command{
sdx-16: 	Use:     "rm",
sdx-17: 	Aliases: []string{"remove", "unload"},
sdx-18: 	Short:   "Remove context",
sdx-19: 	Long:    `Remove context by index, name, or glob.`,
sdx-20: 	Args:    cobra.MinimumNArgs(1),
sdx-21: 	Run:     contextRm,
sdx-22: }
sdx-23:
sdx-24: func contextRm(cmd *cobra.Command, args []string) {
sdx-25: 	auth.MustResolveAuthWithOrg()
sdx-26: 	lib.MustResolveProject()
sdx-27:
sdx-28: 	if lib.CurrentPlanId == "" {
sdx-29: 		fmt.Println("🤷‍♂️ No current plan")
sdx-30: 		return
sdx-31: 	}
sdx-32:
sdx-33: 	term.StartSpinner("")
sdx-34: 	contexts, err := api.Client.ListContext(lib.CurrentPlanId, lib.CurrentBranch)
sdx-35:
sdx-36: 	if err != nil {
sdx-37: 		term.OutputErrorAndExit("Error retrieving context: %v", err)
sdx-38: 	}
sdx-39:
sdx-40: 	deleteIds := map[string]bool{}
sdx-41:
sdx-42: 	for i, context := range contexts {
sdx-43: 		for _, id := range args {
sdx-44: 			if fmt.Sprintf("%d", i+1) == id || context.Name == id || context.FilePath == id || context.Url == id {
sdx-45: 				deleteIds[context.Id] = true
sdx-46: 				break
sdx-47: 			} else if context.FilePath != "" {
sdx-48: 				// Check if id is a glob pattern
sdx-49: 				matched, err := filepath.Match(id, context.FilePath)
sdx-50: 				if err != nil {
sdx-51: 					term.OutputErrorAndExit("Error matching glob pattern: %v", err)
sdx-52: 				}
sdx-53: 				if matched {
sdx-54: 					deleteIds[context.Id] = true
sdx-55: 					break
sdx-56: 				}
sdx-57:
sdx-58: 				// Check if id is a parent directory
sdx-59: 				parentDir := context.FilePath
sdx-60: 				for parentDir != "." && parentDir != "/" && parentDir != "" {
sdx-61: 					if parentDir == id {
sdx-62: 						deleteIds[context.Id] = true
sdx-63: 						break
sdx-64: 					}
sdx-65: 					parentDir = filepath.Dir(parentDir) // Move up one directory
sdx-66: 				}
sdx-67:
sdx-68: 			}
sdx-69: 		}
sdx-70: 	}
sdx-71:
sdx-72: 	if len(deleteIds) > 0 {
sdx-73: 		res, err := api.Client.DeleteContext(lib.CurrentPlanId, lib.CurrentBranch, shared.DeleteContextRequest{
sdx-74: 			Ids: deleteIds,
sdx-75: 		})
sdx-76: 		term.StopSpinner()
sdx-77:
sdx-78: 		if err != nil {
sdx-79: 			term.OutputErrorAndExit("Error deleting context: %v", err)
sdx-80: 		}
sdx-81:
sdx-82: 		fmt.Println("✅ " + res.Msg)
sdx-83: 	} else {
sdx-84: 		term.StopSpinner()
sdx-85: 		fmt.Println("🤷‍♂️ No context removed")
sdx-86: 	}
sdx-87: }
sdx-88:
sdx-89: func init() {
sdx-90: 	RootCmd.AddCommand(contextRmCmd)
sdx-91: }
sdx-92:
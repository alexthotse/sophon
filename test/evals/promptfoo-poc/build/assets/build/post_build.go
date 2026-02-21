sdx-1: package cmd
sdx-2:
sdx-3: import (
sdx-4: 	"fmt"
sdx-5: 	"path/filepath"
sdx-6: 	"sophon/api"
sdx-7: 	"sophon/auth"
sdx-8: 	"sophon/lib"
sdx-9: 	"sophon/term"
sdx-10: 	"strconv"
sdx-11: 	"strings"
sdx-12:
sdx-13: 	"sophon-shared"
sdx-14: 	"github.com/spf13/cobra"
sdx-15: )
sdx-16:
sdx-17: func parseRange(arg string) ([]int, error) {
sdx-18: 	var indices []int
sdx-19: 	parts := strings.Split(arg, "-")
sdx-20: 	if len(parts) == 2 {
sdx-21: 		start, err := strconv.Atoi(parts[0])
sdx-22: 		if err != nil {
sdx-23: 			return nil, err
sdx-24: 		}
sdx-25: 		end, err := strconv.Atoi(parts[1])
sdx-26: 		if err != nil {
sdx-27: 			return nil, err
sdx-28: 		}
sdx-29: 		for i := start; i <= end; i++ {
sdx-30: 			indices = append(indices, i)
sdx-31: 		}
sdx-32: 	} else {
sdx-33: 		index, err := strconv.Atoi(arg)
sdx-34: 		if err != nil {
sdx-35: 			return nil, err
sdx-36: 		}
sdx-37: 		indices = append(indices, index)
sdx-38: 	}
sdx-39: 	return indices, nil
sdx-40: }
sdx-41:
sdx-42: func contextRm(cmd *cobra.Command, args []string) {
sdx-43: 	auth.MustResolveAuthWithOrg()
sdx-44: 	lib.MustResolveProject()
sdx-45:
sdx-46: 	if lib.CurrentPlanId == "" {
sdx-47: 		fmt.Println("🤷‍♂️ No current plan")
sdx-48: 		return
sdx-49: 	}
sdx-50:
sdx-51: 	term.StartSpinner("")
sdx-52: 	contexts, err := api.Client.ListContext(lib.CurrentPlanId, lib.CurrentBranch)
sdx-53:
sdx-54: 	if err != nil {
sdx-55: 		term.OutputErrorAndExit("Error retrieving context: %v", err)
sdx-56: 	}
sdx-57:
sdx-58: 	deleteIds := map[string]bool{}
sdx-59:
sdx-60: 	for _, arg := range args {
sdx-61: 		indices, err := parseRange(arg)
sdx-62: 		if err != nil {
sdx-63: 			term.OutputErrorAndExit("Error parsing range: %v", err)
sdx-64: 		}
sdx-65:
sdx-66: 		for _, index := range indices {
sdx-67: 			if index > 0 && index <= len(contexts) {
sdx-68: 				context := contexts[index-1]
sdx-69: 				deleteIds[context.Id] = true
sdx-70: 			}
sdx-71: 		}
sdx-72: 	}
sdx-73:
sdx-74: 	for i, context := range contexts {
sdx-75: 		for _, id := range args {
sdx-76: 			if fmt.Sprintf("%d", i+1) == id || context.Name == id || context.FilePath == id || context.Url == id {
sdx-77: 				deleteIds[context.Id] = true
sdx-78: 				break
sdx-79: 			} else if context.FilePath != "" {
sdx-80: 				// Check if id is a glob pattern
sdx-81: 				matched, err := filepath.Match(id, context.FilePath)
sdx-82: 				if err != nil {
sdx-83: 					term.OutputErrorAndExit("Error matching glob pattern: %v", err)
sdx-84: 				}
sdx-85: 				if matched {
sdx-86: 					deleteIds[context.Id] = true
sdx-87: 					break
sdx-88: 				}
sdx-89:
sdx-90: 				// Check if id is a parent directory
sdx-91: 				parentDir := context.FilePath
sdx-92: 				for parentDir != "." && parentDir != "/" && parentDir != "" {
sdx-93: 					if parentDir == id {
sdx-94: 						deleteIds[context.Id] = true
sdx-95: 						break
sdx-96: 					}
sdx-97: 					parentDir = filepath.Dir(parentDir) // Move up one directory
sdx-98: 				}
sdx-99: 			}
sdx-100: 		}
sdx-101: 	}
sdx-102:
sdx-103: 	if len(deleteIds) > 0 {
sdx-104: 		res, err := api.Client.DeleteContext(lib.CurrentPlanId, lib.CurrentBranch, shared.DeleteContextRequest{
sdx-105: 			Ids: deleteIds,
sdx-106: 		})
sdx-107: 		term.StopSpinner()
sdx-108:
sdx-109: 		if err != nil {
sdx-110: 			term.OutputErrorAndExit("Error deleting context: %v", err)
sdx-111: 		}
sdx-112:
sdx-113: 		fmt.Println("✅ " + res.Msg)
sdx-114: 	} else {
sdx-115: 		term.StopSpinner()
sdx-116: 		fmt.Println("🤷‍♂️ No context removed")
sdx-117: 	}
sdx-118: }
sdx-119:
sdx-120: func init() {
sdx-121: 	RootCmd.AddCommand(contextRmCmd)
sdx-122: }
sdx-123:
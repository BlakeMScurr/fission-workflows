package dynamic

import (
	"github.com/fission/fission-workflows/pkg/api/invocation"
	"github.com/fission/fission-workflows/pkg/api/workflow"
	"github.com/fission/fission-workflows/pkg/api/workflow/parse"
	"github.com/fission/fission-workflows/pkg/fnenv/workflows"
	"github.com/fission/fission-workflows/pkg/types"
	"github.com/fission/fission-workflows/pkg/types/typedvalues"
	"github.com/fission/fission-workflows/pkg/types/validate"
	"github.com/gogo/protobuf/proto"
)

// Api that servers mainly as a function.Runtime wrapper that deals with the higher-level logic workflow-related logic.
type Api struct {
	wfApi    *workflow.Api
	wfiApi   *invocation.Api
	resolver *parse.Resolver
}

func NewApi(wfApi *workflow.Api, wfiApi *invocation.Api) *Api {
	return &Api{
		wfApi:    wfApi,
		wfiApi:   wfiApi,
		resolver: wfApi.Resolver,
	}
}

func (ap *Api) AddDynamicTask(invocationId string, parentId string, taskSpec *types.TaskSpec) error {

	// Transform TaskSpec into WorkflowSpec
	// TODO dedup workflows
	// TODO indicate relation with workflow somehow?
	wfSpec := &types.WorkflowSpec{
		OutputTask: "main",
		Tasks: map[string]*types.TaskSpec{
			"main": taskSpec,
		},
		Internal:   true, // TODO take into account
		ApiVersion: types.WorkflowApiVersion,
	}

	return ap.addDynamicWorkflow(invocationId, parentId, wfSpec, taskSpec)
}

func (ap *Api) AddDynamicWorkflow(invocationId string, parentId string, workflowSpec *types.WorkflowSpec) error {
	taskSpec := types.NewTaskSpec()
	// TODO add inputs to WorkflowSpec
	return ap.addDynamicWorkflow(invocationId, parentId, workflowSpec, taskSpec)
}

func (ap *Api) addDynamicWorkflow(invocationId string, parentId string, wfSpec *types.WorkflowSpec,
	stubTask *types.TaskSpec) error {

	// Clean-up WorkflowSpec and submit
	sanitizeWorkflow(wfSpec)
	err := validate.WorkflowSpec(wfSpec)
	if err != nil {
		return err
	}
	wfId, err := ap.wfApi.Create(wfSpec)
	if err != nil {
		return err
	}

	// Generate Proxy Task
	proxyTaskSpec := proto.Clone(stubTask).(*types.TaskSpec)
	proxyTaskSpec.FunctionRef = wfId
	proxyTaskSpec.AddInput("_parent", typedvalues.ParseString(invocationId))
	proxyTaskId := parentId + "_child"
	proxyTask := types.NewTask(proxyTaskId)
	proxyTask.Spec = proxyTaskSpec
	proxyTask.Status.Status = types.TaskStatus_READY
	proxyTask.Status.Resolved = workflows.ResolvedWorkflow(wfId)

	// Ensure that the only link of the dynamic task is with its parent
	proxyTaskSpec.Requires = map[string]*types.TaskDependencyParameters{
		parentId: {
			Type: types.TaskDependencyParameters_DYNAMIC_OUTPUT,
		},
	}

	err = validate.TaskSpec(proxyTaskSpec)
	if err != nil {
		return err
	}

	// Submit added task to workflow invocation
	return ap.wfiApi.AddTask(invocationId, proxyTask)

}

func sanitizeWorkflow(v *types.WorkflowSpec) {
	if len(v.ApiVersion) == 0 {
		v.ApiVersion = types.WorkflowApiVersion
	}

	// ForceID is not supported for internal workflows
	v.ForceId = ""
}
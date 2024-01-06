/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Response structure for NodeExecutionGetDataRequest which contains inputs and outputs for a node execution.
type AdminNodeExecutionGetDataResponse struct {
	// Signed url to fetch a core.LiteralMap of node execution inputs. Deprecated: Please use full_inputs instead.
	Inputs *AdminUrlBlob `json:"inputs,omitempty"`
	// Signed url to fetch a core.LiteralMap of node execution outputs. Deprecated: Please use full_outputs instead.
	Outputs *AdminUrlBlob `json:"outputs,omitempty"`
	// Full_inputs will only be populated if they are under a configured size threshold. Deprecated: Please use input_data instead.
	FullInputs *CoreLiteralMap `json:"full_inputs,omitempty"`
	// Full_outputs will only be populated if they are under a configured size threshold. Deprecated: Please use output_data instead.
	FullOutputs *CoreLiteralMap `json:"full_outputs,omitempty"`
	// InputData will only be populated if they are under a configured size threshold.
	InputData *CoreInputData `json:"input_data,omitempty"`
	// OutputData will only be populated if they are under a configured size threshold.
	OutputData *CoreOutputData `json:"output_data,omitempty"`
	// Optional Workflow closure for a dynamically generated workflow, in the case this node yields a dynamic workflow we return its structure here.
	DynamicWorkflow *FlyteidladminDynamicWorkflowNodeMetadata `json:"dynamic_workflow,omitempty"`
	FlyteUrls       *AdminFlyteUrLs                           `json:"flyte_urls,omitempty"`
}

package sdk

import "time"

const (
	// Global Role
	GlobalRoleManagePermission   = "manage-permission"
	GlobalRoleManageOrganization = "manage-organization"
	GlobalRoleManageRegion       = "manage-region"
	GlobalRoleManageHatchery     = "manage-hatchery"
	GlobalRoleManageUser         = "manage-user"
	GlobalRoleManageGroup        = "manage-group"
	GlobalRoleManagePlugin       = "manage-plugin"
	GlobalRoleProjectCreate      = "create-project"

	// Project Role
	ProjectRoleRead                   = "read"
	ProjectRoleManage                 = "manage"
	ProjectRoleManageNotification     = "manage-notification"
	ProjectRoleManageWorkerModel      = "manage-worker-model"
	ProjectRoleManageAction           = "manage-action"
	ProjectRoleManageWorkflow         = "manage-workflow"
	ProjectRoleManageWorkflowTemplate = "manage-workflow-template"
	ProjectRoleManageVariableSet      = "manage-variableset"

	// Hatchery Role
	HatcheryRoleSpawn = "start-worker"

	// Region Role
	RegionRoleList    = "list"
	RegionRoleExecute = "execute"
	RegionRoleManage  = "manage"
)

type RBAC struct {
	ID             string              `json:"id" db:"id"`
	Name           string              `json:"name" db:"name" cli:"name"`
	Created        time.Time           `json:"created" db:"created"`
	LastModified   time.Time           `json:"last_modified" db:"last_modified" cli:"last_modified"`
	Global         []RBACGlobal        `json:"global,omitempty" db:"-"`
	Projects       []RBACProject       `json:"projects,omitempty" db:"-"`
	Regions        []RBACRegion        `json:"regions,omitempty" db:"-"`
	Hatcheries     []RBACHatchery      `json:"hatcheries,omitempty" db:"-"`
	Workflows      []RBACWorkflow      `json:"workflows,omitempty" db:"-"`
	VariableSets   []RBACVariableSet   `json:"variablesets,omitempty" db:"-"`
	RegionProjects []RBACRegionProject `json:"region_projects,omitempty" db:"-"`
}

func (rbac *RBAC) IsEmpty() bool {
	return len(rbac.Projects) == 0 && len(rbac.Hatcheries) == 0 && len(rbac.Global) == 0 && len(rbac.Regions) == 0 && len(rbac.VariableSets) == 0 && len(rbac.Workflows) == 0
}

package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

/*

- Organization is the top level entity in the hierarchy.
- Users join the organization through OrganizationMembership and are assigned a role, either owner or member.
- Owners can create teams within organization.
- Teams can have multiple members and multiple roles (roles provide permissions to resources)
- Members of a team example:
	1. User1 has Read role - can see and access most of the resources
	2. User2 has Write role - can see and access most of the resources, update and delete apps
	3. User3 has Admin role - can see and access all resources, invite new members

- Users grant access to Apps using ResourceAccessBinding. You can create many instances of ResourceAccessBinding for multiple
  users and teams. Each instance can have different roles.
*/

type Organization struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	Owner     string         `json:"owner"` // Who created the org
}

// OrganizationMembership - organization membership is simple, once added, the user is either an owner or a member
type OrganizationMembership struct {
	UserID string `json:"user_id" yaml:"user_id" gorm:"primaryKey"` // composite key
	OrgID  string `json:"org_id" yaml:"org_id" gorm:"primaryKey"`

	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" yaml:"updated_at"`

	Role OrganizationRole `json:"role,omitempty" yaml:"role,omitempty"`
}

type OrganizationRole string

const (
	OrganizationRoleOwner  OrganizationRole = "owner"  // Has full administrative access to the entire organization.
	OrganizationRoleMember OrganizationRole = "member" // Can see every member and team in the organization and can create new apps
)

type Team struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `json:"name"`
}

// Role - a role is a collection of permissions that can be assigned to a user or team.
// Roles are defined within an organization and can be used across teams.
type Role struct {
	ID          string    `json:"id" yaml:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" yaml:"updated_at"`
	OrgID       string    `json:"org_id" yaml:"org_id" gorm:"index"`
	Name        string    `json:"name" yaml:"name"`
	Description string    `json:"description" yaml:"description"`
	Config      Config    `json:"config" yaml:"config"`
}

type Membership struct {
	UserID string `json:"user_id" yaml:"user_id" gorm:"primaryKey"` // composite key
	TeamID string `json:"team_id" yaml:"team_id" gorm:"primaryKey"`

	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" yaml:"updated_at"`

	// extra data fields (optional)
	User  User   `json:"user,omitempty" yaml:"user,omitempty" gorm:"-"`
	Team  Team   `json:"team,omitempty" yaml:"team,omitempty" gorm:"-"`
	Roles []Role `json:"roles,omitempty" yaml:"roles,omitempty" gorm:"-"`
}

type MembershipRoleBindings []MembershipRoleBinding

type MembershipRoleBinding struct {
	UserID    string    `json:"user_id" yaml:"user_id" gorm:"primaryKey"`
	RoleID    string    `json:"role_id" yaml:"role_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" yaml:"updated_at"`
	TeamID    string    `json:"team_id" yaml:"team_id" gorm:"index"`

	// extra data fields (optional)
	User User `json:"user,omitempty" yaml:"user,omitempty" gorm:"-"`
	Role Role `json:"role,omitempty" yaml:"role,omitempty" gorm:"-"`
	Team Team `json:"team,omitempty" yaml:"team,omitempty" gorm:"-"`
}

// ResourceAccessBinding grant access to a resource for a team or user. This allows users
// to share their application, knowledge, provider endpoint, etc with other users or teams.
type ResourceAccessBinding struct {
	ID         string   `json:"id" yaml:"id" gorm:"primaryKey"`
	Resource   Resource `json:"resource" yaml:"resource"`       // Kind of resource
	ResourceID string   `json:"resource_id" yaml:"resource_id"` // App ID, Knowledge ID, etc
	TeamID     string   `json:"team_id" yaml:"team_id"`         // If granted to a team
	UserID     string   `json:"user_id" yaml:"user_id"`         // If granted to a user
	Roles      []Role   `json:"roles,omitempty" yaml:"roles,omitempty" gorm:"-"`
}

// ResourceAccessRoleBinding grants a role to the resource access binding
type ResourceAccessRoleBinding struct {
	ID     string `json:"id" yaml:"id" gorm:"primaryKey"`
	RoleID string `json:"role_id" yaml:"role_id" gorm:"primaryKey"`

	OrgID string `json:"org_id" yaml:"org_id" gorm:"index"`

	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" yaml:"updated_at"`
}

// this lives in the database
// the ID is the keycloak user ID
// there might not be a record for every user
type UserMeta struct {
	ID     string     `json:"id"`
	Config UserConfig `json:"config" gorm:"type:json"`
}

type UserConfig struct {
	StripeSubscriptionActive bool   `json:"stripe_subscription_active"`
	StripeCustomerID         string `json:"stripe_customer_id"`
	StripeSubscriptionID     string `json:"stripe_subscription_id"`
}

func (u UserConfig) Value() (driver.Value, error) {
	j, err := json.Marshal(u)
	return j, err
}

func (u *UserConfig) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}
	var result UserConfig
	if err := json.Unmarshal(source, &result); err != nil {
		return err
	}
	*u = result
	return nil
}

func (UserConfig) GormDataType() string {
	return "json"
}

type Config struct {
	Rules []Rule `json:"rules,omitempty" yaml:"rules,omitempty"`
}

type Rule struct {
	Resources []Resource `json:"resource,omitempty" yaml:"resource,omitempty"`
	Actions   []Action   `json:"actions,omitempty" yaml:"actions,omitempty"`
	Effect    Effect     `json:"effect,omitempty" yaml:"effect,omitempty"`
}

type Effect string

const (
	EffectAllow = Effect("allow")
	EffectDeny  = Effect("deny")
)

type Resource string

const (
	ResourceTeam                  Resource = "Team"
	ResourceOrganization          Resource = "Organization"
	ResourceRole                  Resource = "Role"
	ResourceMembership            Resource = "Membership"
	ResourceMembershipRoleBinding Resource = "MembershipRoleBinding"
	ResourceApplication           Resource = "Application"
	ResourceKnowledge             Resource = "Knowledge"
	ResourceUser                  Resource = "User"
	ResourceAny                   Resource = "*"
)

type Action string

const (
	ActionGet       Action = "Get"
	ActionList      Action = "List"
	ActionDelete    Action = "Delete"
	ActionUpdate    Action = "Update"
	ActionCreate    Action = "Create"
	ActionUseAction Action = "UseAction" // For example "use app"
)

var AvailableActions = map[Action]bool{
	ActionGet:       true,
	ActionList:      true,
	ActionCreate:    true,
	ActionDelete:    true,
	ActionUpdate:    true,
	ActionUseAction: true,
}

func (a Action) String() string {
	return string(a)
}

func ParseActions(actions []string) ([]Action, error) {
	var result []Action
	for _, action := range actions {
		a, err := ParseAction(action)
		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}
	return result, nil

}

func ParseAction(a string) (Action, error) {
	_, ok := AvailableActions[Action(cases.Title(language.English).String(a))]
	if !ok {
		return Action(""), fmt.Errorf("action %s not found", a)
	}
	return Action(cases.Title(language.English).String(a)), nil
}

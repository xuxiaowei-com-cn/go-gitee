package gitee

import (
	"net/http"
)

type SortValue string

const (
	Created  SortValue = "created"   // 创建时间
	Updated  SortValue = "updated"   // 更新时间
	Pushed   SortValue = "pushed"    // 最后推送时间
	FullName SortValue = "full_name" // 仓库所属与名称
)

type VisibilityValue string

const (
	VisibilityPublic  VisibilityValue = "public"  // 公开
	VisibilityPrivate VisibilityValue = "private" // 私有
	VisibilityAll     VisibilityValue = "all"     // 所有
)

type AffiliationValue string

const (
	AffiliationOwner              AffiliationValue = "owner"               // 授权用户拥有的仓库
	AffiliationCollaborator       AffiliationValue = "collaborator"        // 授权用户为仓库成员
	AffiliationOrganizationMember AffiliationValue = "organization_member" // 授权用户为仓库所在组织并有访问仓库权限
	AffiliationEnterpriseMember   AffiliationValue = "enterprise_member"   // 授权用户所在企业并有访问仓库权限
	AffiliationAdmin              AffiliationValue = "admin"               // 所有有权限的，包括所管理的组织中所有仓库、所管理的企业的所有仓库
)

type ReposType string

const (
	ReposOwner    ReposType = "owner"    // 创建者
	ReposPersonal ReposType = "personal" // 个人
	ReposMember   ReposType = "member"   // 其为成员
	ReposPublic   ReposType = "public"   // 公开
	ReposPrivate  ReposType = "private"  // 私有
)

type Direction string

const (
	DirectionAsc  Direction = "asc"
	DirectionDesc Direction = "desc"
)

type GetV5UserReposRequest struct {
	Visibility  VisibilityValue    `url:"visibility,omitempty" json:"visibility,omitempty"`   // 公开(public)、私有(private)或者所有(all)，默认: 所有(all)
	Affiliation []AffiliationValue `url:"affiliation,omitempty" json:"affiliation,omitempty"` // owner(授权用户拥有的仓库)、collaborator(授权用户为仓库成员)、organization_member(授权用户为仓库所在组织并有访问仓库权限)、enterprise_member(授权用户所在企业并有访问仓库权限)、admin(所有有权限的，包括所管理的组织中所有仓库、所管理的企业的所有仓库)。 可以用逗号分隔符组合。如: owner, organization_member 或 owner, collaborator, organization_member
	Type        ReposType          `url:"type,omitempty" json:"type,omitempty"`               // 筛选用户仓库: 其创建(owner)、个人(personal)、其为成员(member)、公开(public)、私有(private)，不能与 visibility 或 affiliation 参数一并使用，否则会报 422 错误
	Sort        SortValue          `url:"sort,omitempty" json:"sort,omitempty"`               // 排序方式: 创建时间(created)，更新时间(updated)，最后推送时间(pushed)，仓库所属与名称(full_name)。默认: full_name
	Direction   Direction          `url:"direction,omitempty" json:"direction,omitempty"`     // 如果sort参数为full_name，用升序(asc)。否则降序(desc)
	Q           string             `url:"q,omitempty" json:"q,omitempty"`                     // 搜索关键字
	ListOptions
}

type Assignee struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Name              string `json:"name"`
	AvatarURL         string `json:"avatar_url"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	Remark            string `json:"remark"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
}

type Namespace struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	HTMLURL string `json:"html_url"`
}

type Owner struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Name              string `json:"name"`
	AvatarURL         string `json:"avatar_url"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	Remark            string `json:"remark"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
}

type Repository struct {
	ID                  int         `json:"id"`
	FullName            string      `json:"full_name"`
	HumanName           string      `json:"human_name"`
	URL                 string      `json:"url"`
	Namespace           Namespace   `json:"namespace"`
	Path                string      `json:"path"`
	Name                string      `json:"name"`
	Owner               Owner       `json:"owner"`
	Assigner            Assignee    `json:"assigner"`
	Description         string      `json:"description"`
	Private             bool        `json:"private"`
	Public              bool        `json:"public"`
	Internal            bool        `json:"internal"`
	Fork                bool        `json:"fork"`
	HTMLURL             string      `json:"html_url"`
	SSHURL              string      `json:"ssh_url"`
	ForksURL            string      `json:"forks_url"`
	KeysURL             string      `json:"keys_url"`
	CollaboratorsURL    string      `json:"collaborators_url"`
	HooksURL            string      `json:"hooks_url"`
	BranchesURL         string      `json:"branches_url"`
	TagsURL             string      `json:"tags_url"`
	BlobsURL            string      `json:"blobs_url"`
	StargazersURL       string      `json:"stargazers_url"`
	ContributorsURL     string      `json:"contributors_url"`
	CommitsURL          string      `json:"commits_url"`
	CommentsURL         string      `json:"comments_url"`
	IssueCommentURL     string      `json:"issue_comment_url"`
	IssuesURL           string      `json:"issues_url"`
	PullsURL            string      `json:"pulls_url"`
	MilestonesURL       string      `json:"milestones_url"`
	NotificationsURL    string      `json:"notifications_url"`
	LabelsURL           string      `json:"labels_url"`
	ReleasesURL         string      `json:"releases_url"`
	Recommend           bool        `json:"recommend"`
	GVP                 bool        `json:"gvp"`
	Homepage            interface{} `json:"homepage"`
	Language            interface{} `json:"language"`
	ForksCount          int         `json:"forks_count"`
	StargazersCount     int         `json:"stargazers_count"`
	WatchersCount       int         `json:"watchers_count"`
	DefaultBranch       string      `json:"default_branch"`
	OpenIssuesCount     int         `json:"open_issues_count"`
	HasIssues           bool        `json:"has_issues"`
	HasWiki             bool        `json:"has_wiki"`
	IssueComment        bool        `json:"issue_comment"`
	CanComment          bool        `json:"can_comment"`
	PullRequestsEnabled bool        `json:"pull_requests_enabled"`
	HasPage             bool        `json:"has_page"`
	License             interface{} `json:"license"`
	Outsourced          bool        `json:"outsourced"`
	ProjectCreator      string      `json:"project_creator"`
	Members             []string    `json:"members"`
	PushedAt            string      `json:"pushed_at"`
	CreatedAt           string      `json:"created_at"`
	UpdatedAt           string      `json:"updated_at"`
	Parent              interface{} `json:"parent"`
	PaaS                interface{} `json:"paas"`
	Stared              bool        `json:"stared"`
	Watched             bool        `json:"watched"`
	Permission          struct {
		Pull  bool `json:"pull"`
		Push  bool `json:"push"`
		Admin bool `json:"admin"`
	} `json:"permission"`
	Relation            string        `json:"relation"`
	AssigneesNumber     int           `json:"assignees_number"`
	TestersNumber       int           `json:"testers_number"`
	Assignee            []Assignee    `json:"assignee"`
	Testers             []Assignee    `json:"testers"`
	Status              string        `json:"status"`
	Programs            []interface{} `json:"programs"`
	Enterprise          interface{}   `json:"enterprise"`
	ProjectLabels       []interface{} `json:"project_labels"`
	IssueTemplateSource string        `json:"issue_template_source"`
}

type GetV5UserReposService struct {
	client *Client
}

// GetV5UserRepos 列出授权用户的所有仓库 https://gitee.com/api/v5/swagger#/getV5UserRepos
func (s *GetV5UserReposService) GetV5UserRepos(request *GetV5UserReposRequest) ([]*Repository, *Response, error) {

	u := "user/repos"

	req, err := s.client.NewRequest(http.MethodGet, u, request)
	if err != nil {
		return nil, nil, err
	}

	var repositories []*Repository
	resp, err := s.client.Do(req, &repositories)
	if err != nil {
		return nil, resp, err
	}

	return repositories, resp, nil
}

package gitee

import (
	"os"
	"testing"
)

// 列出授权用户的所有仓库
func TestGetV5UserRepos(t *testing.T) {

	var token = os.Getenv("GO_GITEE_TOKEN")

	gitClient, err := NewClient(token)
	if err != nil {
		t.Fatalf("创建客户端异常：%s", err)
	}

	request := &GetV5UserReposRequest{
		Visibility:  VisibilityAll,
		Affiliation: []AffiliationValue{AffiliationOwner, AffiliationCollaborator},
		//Type:        ReposOwner,
		Sort:      Created,
		Direction: DirectionDesc,
		Q:         "xuxiaowei",
		ListOptions: ListOptions{
			Page:    1,
			PerPage: 12,
		},
	}

	repos, response, err := gitClient.GetV5UserRepos.GetV5UserRepos(request)
	if err != nil {
		t.Fatalf("列出授权用户的所有仓库 异常：%s", err)
	}

	t.Log(response.Status)
	t.Log(len(repos))
	t.Log(repos[0].Name)

	//jsonData, err := json.Marshal(repos)
	//if err != nil {
	//	t.Log("转换失败：", err)
	//	return
	//}
	//t.Log(string(jsonData))
}

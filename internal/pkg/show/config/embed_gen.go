// Code generated by gen_embed_var.go; DO NOT EDIT.
package config

import _ "embed"

//go:embed default.yaml
var DefaultConfig string

// plugin default config
var (

	//go:embed plugins/argocd.yaml
	ArgocdDefaultConfig string

	//go:embed plugins/argocdapp.yaml
	ArgocdappDefaultConfig string

	//go:embed plugins/devlake.yaml
	DevlakeDefaultConfig string

	//go:embed plugins/github-repo-scaffolding-golang.yaml
	GithubRepoScaffoldingGolangDefaultConfig string

	//go:embed plugins/githubactions-golang.yaml
	GithubactionsGolangDefaultConfig string

	//go:embed plugins/githubactions-nodejs.yaml
	GithubactionsNodejsDefaultConfig string

	//go:embed plugins/githubactions-python.yaml
	GithubactionsPythonDefaultConfig string

	//go:embed plugins/gitlab-ce-docker.yaml
	GitlabCeDockerDefaultConfig string

	//go:embed plugins/gitlab-repo-scaffolding-golang.yaml
	GitlabRepoScaffoldingGolangDefaultConfig string

	//go:embed plugins/gitlabci-generic.yaml
	GitlabciGenericDefaultConfig string

	//go:embed plugins/gitlabci-golang.yaml
	GitlabciGolangDefaultConfig string

	//go:embed plugins/harbor.yaml
	HarborDefaultConfig string

	//go:embed plugins/hashicorp-vault.yaml
	HashicorpVaultDefaultConfig string

	//go:embed plugins/helm-generic.yaml
	HelmGenericDefaultConfig string

	//go:embed plugins/jenkins.yaml
	JenkinsDefaultConfig string

	//go:embed plugins/jira-github-integ.yaml
	JiraGithubIntegDefaultConfig string

	//go:embed plugins/kube-prometheus.yaml
	KubePrometheusDefaultConfig string

	//go:embed plugins/openldap.yaml
	OpenldapDefaultConfig string

	//go:embed plugins/tekton.yaml
	TektonDefaultConfig string

	//go:embed plugins/trello-github-integ.yaml
	TrelloGithubIntegDefaultConfig string

	//go:embed plugins/trello.yaml
	TrelloDefaultConfig string

	//go:embed plugins/zentao.yaml
	ZentaoDefaultConfig string
)

var pluginDefaultConfigs = map[string]string{
	"argocd":                         ArgocdDefaultConfig,
	"argocdapp":                      ArgocdappDefaultConfig,
	"devlake":                        DevlakeDefaultConfig,
	"github-repo-scaffolding-golang": GithubRepoScaffoldingGolangDefaultConfig,
	"githubactions-golang":           GithubactionsGolangDefaultConfig,
	"githubactions-nodejs":           GithubactionsNodejsDefaultConfig,
	"githubactions-python":           GithubactionsPythonDefaultConfig,
	"gitlab-ce-docker":               GitlabCeDockerDefaultConfig,
	"gitlab-repo-scaffolding-golang": GitlabRepoScaffoldingGolangDefaultConfig,
	"gitlabci-generic":               GitlabciGenericDefaultConfig,
	"gitlabci-golang":                GitlabciGolangDefaultConfig,
	"harbor":                         HarborDefaultConfig,
	"hashicorp-vault":                HashicorpVaultDefaultConfig,
	"helm-generic":                   HelmGenericDefaultConfig,
	"jenkins":                        JenkinsDefaultConfig,
	"jira-github-integ":              JiraGithubIntegDefaultConfig,
	"kube-prometheus":                KubePrometheusDefaultConfig,
	"openldap":                       OpenldapDefaultConfig,
	"tekton":                         TektonDefaultConfig,
	"trello-github-integ":            TrelloGithubIntegDefaultConfig,
	"trello":                         TrelloDefaultConfig,
	"zentao":                         ZentaoDefaultConfig,
}

//go:embed quickstart.yaml
var QuickStart string

//go:embed gitops.yaml
var GitOps string

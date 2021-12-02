package internal

import (
	"regexp"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	viperInstance := viper.New()
	viperInstance.SetConfigName("config_test.yaml")
	viperInstance.SetConfigType("yaml")
	viperInstance.AddConfigPath(".")
	err := viperInstance.ReadInConfig()
	assert.NoError(t, err)

	c1, err := LoadConfig(*viperInstance)
	assert.NoError(t, err)

	c2 := Config{
		Files: ConfigFiles{
			Includes: []regexp.Regexp{*regexp.MustCompile(`\.yaml$`)},
			Excludes: []regexp.Regexp{*regexp.MustCompile(`\.generated\.yaml$`)},
		},
		Registries: map[string]Registry{
			"docker": DockerRegistry{
				Interval: time.Duration(60000000000),
				Url:      "https://registry-1.docker.io",
				Credentials: HttpBasicCredentials{
					Username: "user",
					Password: "pass",
				},
			},
			"helm": HelmRegistry{
				Interval: time.Duration(3600000000000),
				Url:      "https://charts.helm.sh/stable",
				Credentials: HttpBasicCredentials{
					Username: "user",
					Password: "pass",
				},
			},
			"git-hub": GitHubTagRegistry{
				Interval: time.Duration(3600000000000),
				Credentials: HttpBasicCredentials{
					Username: "user",
					Password: "pass",
				},
			},
		},
		Policies: map[string]Policy{
			"lexicographic": {
				Pattern: regexp.MustCompile(`^(?P<all>.*)$`),
				Extracts: []Extract{
					{
						Value: "<all>",
						Strategy: LexicographicExtractStrategy{
							Pin: true,
						},
					},
				},
			},
			"numeric": {
				Pattern: regexp.MustCompile(`^(?P<all>.*)$`),
				Extracts: []Extract{
					{
						Value: "<all>",
						Strategy: NumericExtractStrategy{
							Pin: true,
						},
					},
				},
			},
			"semver": {
				Pattern: regexp.MustCompile(`^(?P<all>.*)$`),
				Extracts: []Extract{
					{
						Value: "<all>",
						Strategy: SemverExtractStrategy{
							PinMajor:         true,
							PinMinor:         true,
							PinPatch:         true,
							AllowPrereleases: true,
						},
					},
				},
			},
		},
		Git: Git{
			Provider: GitHubGitProvider{
				Author: GitAuthor{
					Name:  "name",
					Email: "email",
				},
				AccessToken: "access_token",
			},
		},
	}

	assert.Equal(t, c2.Files, c1.Files)
	assert.Equal(t, c2.Registries, c1.Registries)
	assert.Equal(t, c2.Policies, c1.Policies)
	assert.Equal(t, c2.Git, c1.Git)
}

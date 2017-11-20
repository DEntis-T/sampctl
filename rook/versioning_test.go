package rook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDependency_Validate(t *testing.T) {
	tests := []struct {
		name      string
		d         DependencyString
		wantValid bool
		wantErr   bool
	}{
		// Unversioned
		{"v u https url", DependencyString("https://github.com/user/repo_name"), true, false},
		{"v u http url", DependencyString("http://github.com/user/repo_name"), true, false},
		{"v u naked url", DependencyString("github.com/user/repo_name"), true, false},
		{"v u user/repo", DependencyString("user/repo_name"), true, false},
		{"v u https url path", DependencyString("https://github.com/user/repo_name/inc"), true, false},
		{"v u http url path", DependencyString("http://github.com/user/repo_name/inc"), true, false},
		{"v u naked url path", DependencyString("github.com/user/repo_name/inc"), true, false},
		{"v u user/repo path", DependencyString("user/repo_name/inc"), true, false},

		// Versioned - semver
		{"v v https url", DependencyString("https://github.com/user/repo:1.2.3"), true, false},
		{"v v http url", DependencyString("http://github.com/user/repo:1.2.3"), true, false},
		{"v v naked url", DependencyString("github.com/user/repo:1.2.3"), true, false},
		{"v v user/repo", DependencyString("user/repo:1.2.3"), true, false},
		{"v v user/repo", DependencyString("user/repo:^1.2.3"), true, false},
		{"v v user/repo", DependencyString("user/repo:^2.0"), true, false},
		{"v v user/repo", DependencyString("user/repo:2.1.x"), true, false},
		{"v v user/repo", DependencyString("user/repo:~1"), true, false},
		{"v v user/repo", DependencyString("user/repo:~2.x"), true, false},
		{"v v https url path", DependencyString("https://github.com/user/repo/inc:1.2.3"), true, false},
		{"v v http url path", DependencyString("http://github.com/user/repo/inc:1.2.3"), true, false},
		{"v v naked url path", DependencyString("github.com/user/repo/inc:1.2.3"), true, false},
		{"v v user/repo path", DependencyString("user/repo/inc:1.2.3"), true, false},
		{"v v user/repo path", DependencyString("user/repo/inc:^1.2.3"), true, false},
		{"v v user/repo path", DependencyString("user/repo/inc:^2.0"), true, false},
		{"v v user/repo path", DependencyString("user/repo/inc:2.1.x"), true, false},
		{"v v user/repo path", DependencyString("user/repo/inc:~1"), true, false},
		{"v v user/repo path", DependencyString("user/repo/inc:~2.x"), true, false},

		// Versioned - custom
		{"v c https url", DependencyString("https://github.com/user/repo:stable-release-3"), true, true},
		{"v c http url", DependencyString("http://github.com/user/repo:stable-release-3"), true, true},
		{"v c naked url", DependencyString("github.com/user/repo:stable-release-3"), true, true},
		{"v c user/repo", DependencyString("user/repo:stable-release-3"), true, true},
		{"v c https url path", DependencyString("https://github.com/user/repo/inc:stable-release-3"), true, true},
		{"v c http url path", DependencyString("http://github.com/user/repo/inc:stable-release-3"), true, true},
		{"v c naked url path", DependencyString("github.com/user/repo/inc:stable-release-3"), true, true},
		{"v c user/repo path", DependencyString("user/repo/inc:stable-release-3"), true, true},

		// Unversioned - Invalid
		{"i u www", DependencyString("www.github.com/user/repo"), false, true},
		{"i u user", DependencyString("http://github.com/repo"), false, true},
		{"i u project", DependencyString("project"), false, true},
		{"i u user:repo", DependencyString("user:repo"), false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, err := tt.d.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.wantValid, gotValid)
		})
	}
}

func TestDependencyString_Explode(t *testing.T) {
	tests := []struct {
		name    string
		d       DependencyString
		wantDep Dependency
		wantErr bool
	}{
		// Unversioned
		{"v u https url", DependencyString("https://github.com/user/repo_name"), Dependency{"user", "repo_name", "", ""}, false},
		{"v u http url", DependencyString("http://github.com/user/repo_name"), Dependency{"user", "repo_name", "", ""}, false},
		{"v u naked url", DependencyString("github.com/user/repo_name"), Dependency{"user", "repo_name", "", ""}, false},
		{"v u user/repo", DependencyString("user/repo_name"), Dependency{"user", "repo_name", "", ""}, false},
		{"v u https url path", DependencyString("https://github.com/user/repo_name/inc"), Dependency{"user", "repo_name", "inc", ""}, false},
		{"v u http url path", DependencyString("http://github.com/user/repo_name/inc"), Dependency{"user", "repo_name", "inc", ""}, false},
		{"v u naked url path", DependencyString("github.com/user/repo_name/inc"), Dependency{"user", "repo_name", "inc", ""}, false},
		{"v u user/repo path", DependencyString("user/repo_name/inc"), Dependency{"user", "repo_name", "inc", ""}, false},

		// Versioned - semver
		{"v v https url", DependencyString("https://github.com/user/repo:1.2.3"), Dependency{"user", "repo", "", "1.2.3"}, false},
		{"v v http url", DependencyString("http://github.com/user/repo:1.2.3"), Dependency{"user", "repo", "", "1.2.3"}, false},
		{"v v naked url", DependencyString("github.com/user/repo:1.2.3"), Dependency{"user", "repo", "", "1.2.3"}, false},
		{"v v user/repo", DependencyString("user/repo:1.2.3"), Dependency{"user", "repo", "", "1.2.3"}, false},
		{"v v user/repo", DependencyString("user/repo:^1.2.3"), Dependency{"user", "repo", "", "^1.2.3"}, false},
		{"v v user/repo", DependencyString("user/repo:^2.0"), Dependency{"user", "repo", "", "^2.0"}, false},
		{"v v user/repo", DependencyString("user/repo:2.1.x"), Dependency{"user", "repo", "", "2.1.x"}, false},
		{"v v user/repo", DependencyString("user/repo:~1"), Dependency{"user", "repo", "", "~1"}, false},
		{"v v user/repo", DependencyString("user/repo:~2.x"), Dependency{"user", "repo", "", "~2.x"}, false},
		{"v v https url path", DependencyString("https://github.com/user/repo/inc:1.2.3"), Dependency{"user", "repo", "inc", "1.2.3"}, false},
		{"v v http url path", DependencyString("http://github.com/user/repo/inc:1.2.3"), Dependency{"user", "repo", "inc", "1.2.3"}, false},
		{"v v naked url path", DependencyString("github.com/user/repo/inc:1.2.3"), Dependency{"user", "repo", "inc", "1.2.3"}, false},
		{"v v user/repo path", DependencyString("user/repo/inc:1.2.3"), Dependency{"user", "repo", "inc", "1.2.3"}, false},
		{"v v user/repo path", DependencyString("user/repo/inc:^1.2.3"), Dependency{"user", "repo", "inc", "^1.2.3"}, false},
		{"v v user/repo path", DependencyString("user/repo/inc:^2.0"), Dependency{"user", "repo", "inc", "^2.0"}, false},
		{"v v user/repo path", DependencyString("user/repo/inc:2.1.x"), Dependency{"user", "repo", "inc", "2.1.x"}, false},
		{"v v user/repo path", DependencyString("user/repo/inc:~1"), Dependency{"user", "repo", "inc", "~1"}, false},
		{"v v user/repo path", DependencyString("user/repo/inc:~2.x"), Dependency{"user", "repo", "inc", "~2.x"}, false},

		// Versioned - custom
		{"v c https url", DependencyString("https://github.com/user/repo:stable-release-3"), Dependency{"user", "repo", "", "stable-release-3"}, false},
		{"v c http url", DependencyString("http://github.com/user/repo:stable-release-3"), Dependency{"user", "repo", "", "stable-release-3"}, false},
		{"v c naked url", DependencyString("github.com/user/repo:stable-release-3"), Dependency{"user", "repo", "", "stable-release-3"}, false},
		{"v c user/repo", DependencyString("user/repo:stable-release-3"), Dependency{"user", "repo", "", "stable-release-3"}, false},
		{"v c https url path", DependencyString("https://github.com/user/repo/inc:stable-release-3"), Dependency{"user", "repo", "inc", "stable-release-3"}, false},
		{"v c http url path", DependencyString("http://github.com/user/repo/inc:stable-release-3"), Dependency{"user", "repo", "inc", "stable-release-3"}, false},
		{"v c naked url path", DependencyString("github.com/user/repo/inc:stable-release-3"), Dependency{"user", "repo", "inc", "stable-release-3"}, false},
		{"v c user/repo path", DependencyString("user/repo/inc:stable-release-3"), Dependency{"user", "repo", "inc", "stable-release-3"}, false},

		// Unversioned - Invalid
		{"i u www", DependencyString("www.github.com/user/repo"), Dependency{"", "", "", ""}, true},
		{"i u user", DependencyString("http://github.com/repo"), Dependency{"", "", "", ""}, true},
		{"i u project", DependencyString("project"), Dependency{"", "", "", ""}, true},
		{"i u user:repo", DependencyString("user:repo"), Dependency{"", "", "", ""}, true}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDep, err := tt.d.Explode()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.wantDep, gotDep)
		})
	}
}

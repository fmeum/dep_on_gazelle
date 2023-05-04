package dep_on_gazelle

import "github.com/bazelbuild/bazel-gazelle/label"

func MakeLabel(repo, pkg, name string) label.Label {
	return label.New(repo, pkg, name)
}

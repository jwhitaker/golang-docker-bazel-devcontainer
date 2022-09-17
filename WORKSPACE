workspace(name = "com_github_thewhitakers_echoer")

load("//:bazel/rules/deps.bzl", "bazel_dependencies")

bazel_dependencies()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:bazel/go/deps.bzl", "go_dependencies")


# gazelle:repository_macro bazel/go/deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.19.1")

gazelle_dependencies()

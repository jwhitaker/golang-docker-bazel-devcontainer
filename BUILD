load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/thewhitakers/echoer
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=bazel/go/deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

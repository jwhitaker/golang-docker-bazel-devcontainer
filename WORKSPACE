workspace(name = "com_github_thewhitakers_echoer")

load("//:bazel/rules/deps.bzl", "bazel_dependencies")

bazel_dependencies()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:bazel/go/deps.bzl", "go_dependencies")
load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories"
)


# gazelle:repository_macro bazel/go/deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.19.1")

gazelle_dependencies()


container_repositories()

_go_image_repos()

load("@io_bazel_rules_docker//container:pull.bzl", "container_pull")

container_pull(
    name = "ubuntu_jammy_amd64",
    registry = "index.docker.io",
    repository = "ubuntu",
    tag = "jammy",
)

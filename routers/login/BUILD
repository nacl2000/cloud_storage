load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "login",
    srcs = ["login.go"],
    data = ["//frontend/login:login_html"],
    importpath = "github.com/nacl2000/clould_storage/routers/login",
    visibility = ["//visibility:public"],
    deps = ["@com_github_gin_gonic_gin//:go_default_library"],
)

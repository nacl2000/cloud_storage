load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "home",
    srcs = ["home.go"],
    data = ["//frontend/home:home_ui"],
    importpath = "github.com/nacl2000/clould_storage/routers/home",
    visibility = ["//service:__pkg__"],
    deps = [
        "//tools/path",
        "@com_github_gin_gonic_contrib//sessions",
        "@com_github_gin_gonic_gin//:gin",
    ],
)

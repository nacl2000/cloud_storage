load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "main",
    embed = [":service_lib"],
    importpath = "github.com/nacl2000/clould_storage/service",
    visibility = ["//visibility:private"],
)

go_library(
    name = "service_lib",
    srcs = ["main.go"],
    importpath = "github.com/nacl2000/clould_storage/service",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_gin_gonic_contrib//sessions",
        "@com_github_gin_gonic_gin//:gin",
        "//routers/login",
        "//tools/path",
    ],
)

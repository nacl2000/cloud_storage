load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "main",
    srcs = ["main.go"],
    importpath = "github.com/nacl2000/clould_storage/service",
    visibility = ["//visibility:private"],
    deps = [
        "//routers/login",
        "//tools/path",
        "@com_github_gin_gonic_gin//:go_default_library",
    ],
)

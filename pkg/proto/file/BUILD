load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "file_proto",
    srcs = ["file.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "file_go_proto",
    importpath = "github.com/nacl2000/cloud_storage/pkg/proto/file",
    proto = ":file_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "file",
    embed = [":file_go_proto"],
    importpath = "github.com/nacl2000/cloud_storage/pkg/proto/file",
    visibility = ["//visibility:public"],
)

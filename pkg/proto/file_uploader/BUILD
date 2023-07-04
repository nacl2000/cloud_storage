load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "file_uploader_proto",
    srcs = ["file_uploader.proto"],
    visibility = ["//visibility:public"],
    deps = ["//pkg/proto/file:file_proto"],
)

go_proto_library(
    name = "file_uploader_go_grpc",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/nacl2000/cloud_storage/pkg/proto/file_uploader",
    proto = ":file_uploader_proto",
    visibility = ["//visibility:public"],
    deps = ["//pkg/proto/file:file_go_proto"],
)

go_library(
    name = "file_uploader",
    embed = [":file_uploader_go_grpc"],
    importpath = "github.com/nacl2000/cloud_storage/pkg/proto/file_uploader",
    visibility = ["//visibility:public"],
)

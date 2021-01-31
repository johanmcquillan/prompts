package(default_visibility = ["PUBLIC"])

go_library(
    name = "colors",
    srcs = [
        "ansi.go",
        "colors.go",
        "shell.go",
        "zsh.go",
    ],
    deps = [
        "//src/prompts/env",
    ],
)

go_test(
    name = "test",
    srcs = [
        "ansi_test.go",
    ],
    deps = [
        ":colors",
        "//deps:testify",
        "//src/prompts",
    ],
)

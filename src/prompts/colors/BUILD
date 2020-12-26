package(default_visibility = ["PUBLIC"])

go_library(
    name = "colors",
    srcs = [
        "ansi.go",
        "colors.go",
        "zsh.go",
        "shell.go",
    ],
    deps = [
        "//src/prompts",
    ],
)

go_test(
    name = "test",
    srcs = [
        "ansi_test.go",
    ],
    deps = [
        ":colors",
        "//src/prompts",
        "//deps:testify",
    ],
)

package(default_visibility = ["PUBLIC"])

go_library(
    name = "env",
    srcs = [
        "env.go",
        "pwd.go",
        "shell.go",
    ],
    deps = [
        "//src/prompts",
    ],
)

go_test(
    name = "test",
    srcs = [
        "env_test.go",
    ],
    deps = [
        ":env",
        "//third_party:testify",
        "//src/prompts",
    ],
)

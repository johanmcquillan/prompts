package(default_visibility = ['PUBLIC'])

go_library(
    name = 'env',
    srcs = [
        'env.go',
        'pwd.go',
    ],
    deps = [
        '//src/prompts',
    ],
)

go_test(
    name = 'test',
    srcs = [
        'env_test.go',
    ],
    deps = [
        ":env",
        "//src/prompts",
        "//deps:testify",
    ],
)

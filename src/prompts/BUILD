go_library(
    name = 'prompts',
    srcs = glob(
        ['*.go'],
        exclude = ['*_test.go'],
    ),
    deps = [
        '//deps:flags',
    ],
    visibility = ['PUBLIC'],
)

go_test(
    name = 'test',
    srcs = glob(
        ['*_test.go'],
    ),
    deps = [
        ":prompts",
        "//deps:testify",
    ],
    visibility = ['PUBLIC'],
)

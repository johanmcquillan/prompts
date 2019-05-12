go_library(
    name = 'prompts',
    srcs = glob(
        ['*.go'],
        exclude = ['*_test.go'],
    ),
    deps = [
        '//third_party:flags',
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
        "//third_party:testify",
    ],
    visibility = ['PUBLIC'],
)

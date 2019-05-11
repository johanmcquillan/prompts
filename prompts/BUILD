go_library(
    name = 'prompts',
    srcs = glob(
        ['*.go'],
        exclude = ['*_test.go'],
    ),
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

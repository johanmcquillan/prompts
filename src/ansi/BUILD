package(default_visibility = ['PUBLIC'])

go_library(
    name = 'ansi',
    srcs = ['ansi.go'],
    deps = [
        '//src/prompts',
    ],
)

go_test(
    name = 'test',
    srcs = glob(
        ['*_test.go'],
    ),
    deps = [
        ":ansi",
        "//src/prompts",
        "//third_party:testify",
    ],
    visibility = ['PUBLIC'],
)

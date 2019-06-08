go_binary(
    name = 'my-prompt-bash',
    srcs = ['my-prompt-bash.go'],
    deps = [
        '//src/ansi',
        '//src/git',
        '//src/k8s',
        '//src/prompts',
        '//third_party:flags',
    ],
)

go_binary(
    name = 'my-prompt-zsh',
    srcs = ['my-prompt-zsh.go'],
    deps = [
        '//src/git',
        '//src/k8s',
        '//src/prompts',
        '//src/zsh',
        '//third_party:flags',
    ],
)

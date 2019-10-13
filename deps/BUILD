package(default_visibility = ["PUBLIC"])

go_get(
    name = "testify",
    get = "github.com/stretchr/testify",
    install = [
        "assert",
        "require",
        "vendor/github.com/davecgh/go-spew/spew",
        "vendor/github.com/pmezard/go-difflib/difflib",
    ],
    deps = [
        ":yaml",
    ],
)

go_get(
    name = "flags",
    get = "github.com/jessevdk/go-flags",
)

go_get(
    name = "yaml",
    get = "gopkg.in/yaml.v2",
)

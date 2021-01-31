package(default_visibility = ["PUBLIC"])

go_get(
    name = "testify",
    get = "github.com/stretchr/testify",
    install = [
        "assert",
        "require",
    ],
    revision = "v1.6.1",
    deps = [
        ":difflib",
        ":spew",
        ":yaml",
    ],
)

go_get(
    name = "spew",
    get = "github.com/davecgh/go-spew/spew",
    revision = "v1.1.1",
)

go_get(
    name = "yaml",
    get = "gopkg.in/yaml.v3",
)

go_get(
    name = "difflib",
    get = "github.com/pmezard/go-difflib/difflib",
    revision = "v1.0.0",
)

go_get(
    name = "flags",
    get = "github.com/jessevdk/go-flags",
    revision = "v1.4.0",
    deps = [
        ":x-sys",
    ],
)

go_get(
    name = "x-sys",
    get = "golang.org/x/sys",
    install = [
        "unix",
    ],
)
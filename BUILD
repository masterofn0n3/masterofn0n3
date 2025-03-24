"""
Primary BUILD file for the project.
"""

load("@bazel_gazelle//:def.bzl", "gazelle")

# Configure Gazelle to generate BUILD files
gazelle(
    name = "gazelle",
    prefix = "github.com/masterofn0n3/masterofn0n3",  # Your module path
)

# Update Gazelle with dependencies from go.mod
gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=go_deps.bzl%go_dependencies",
        "-bzlmod",
        "-prune",
    ],
    command = "update-repos",
)

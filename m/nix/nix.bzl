"""
"""

def nix_flake(name, visibility = None, flakes = []):
    """ something

    Args:
        name: something
        visibility: something
        flakes: something
    """
    dir = native.package_name()

    native.filegroup(
        name = "{}_config".format(name),
        srcs = [
            "flake.json",
            "flake.lock",
            "flake.nix",
        ],
    )

    flake_store_script = Label(":flake_store_script")

    native.genrule(
        name = "{}_archives".format(name),
        srcs = [
            flake_store_script,
            "{}_config".format(name),
        ],
        outs = ["{}_nix_stores".format(name)],
        cmd = "$(location //{}:{}) {} $@".format(flake_store_script.package, flake_store_script.name, dir),
        visibility = visibility,
    )

    flake_path_script = Label(":flake_path_script".format(name))

    for f in flakes:
        native.genrule(
            name = "{}_{}".format(name, f),
            srcs = [
                flake_path_script,
                "{}_config".format(name),
            ],
            outs = ["{}_{}_bin".format(name, f)],
            cmd = "$(location //{}:{}) {} $@ which {}".format(flake_path_script.package, flake_path_script.name, dir, f),
            visibility = visibility,
        )

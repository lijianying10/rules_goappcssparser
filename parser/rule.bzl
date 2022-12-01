"modle: imple the metaprogramming parser"
def _goapp_css_parser(ctx):
    "imple the rule"
    #print("debug: out path",ctx.outputs.out.path)
    #print("debug: out short_path",ctx.outputs.out.short_path)
    #print("debug: out name",ctx.outputs.out.basename)
    input_list = []
    args_list = []
    for i in ctx.attr.srcs:
        for f in i.files.to_list():
            args_list.append("-i")
            args_list.append(f.short_path)
            input_list.append(f)

    args_list.append("-o") 
    args_list.append(ctx.outputs.out.path)

    #ctx.actions.declare_file(ctx.outputs.out.basename)

    ctx.actions.run(
        outputs=[ctx.outputs.out], 
        inputs=input_list, 
        arguments=args_list, 
        mnemonic="goappcssparser", 
        executable = ctx.executable._runner, 
        progress_message=None, 
        use_default_shell_env=False,
        execution_requirements=None, 
        input_manifests=None, 
    )

    pass

goapp_css_parser = rule(
    implementation = _goapp_css_parser,
    attrs = {
        "srcs":attr.label_list(
            allow_files = [".go"],
        ),
        "out": attr.output(mandatory = True),

        "_runner":attr.label(
            default = Label("@rules_goappcssparser//parser:goappcssparser"),
            executable = True,
            cfg = "exec"
        ),
    }
)
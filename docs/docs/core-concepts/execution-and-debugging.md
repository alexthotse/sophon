---
sidebar_position: 6
sidebar_label: Execution and Debugging
---

# Execution and Debugging

Sophon includes command execution and automated debugging capabilities that aim to balance power, control, and safety.

## Command Execution

During a plan, apart from making changes to files, Sophon can write to a special path, `_apply.sh`, with any commands required to complete the plan. This commonly includes installing dependencies, running tests, building and running code, starting servers, and so on.

Commands accumulate in the sandbox just like [pending changes to files](./reviewing-changes.md). If execution fails, you can roll back changes and optionally send the output to the model for automated debugging and retries.

While Sophon will attempt to automatically infer relevant commands to run, it also tries not to overdo it. It generally won't, for example, run a test suite unless you've specifically asked it to, since it may not be desirable after every change. It tries to only run what's strictly necessary, to make local project-level changes instead of global system-wide changes, to check existing dependencies before installing, to write idempotent commands, to avoid hiding output or asking for user input, and to recover gracefully from failures.

If you want specific commands to run and Sophon isn't including automatically them because of it's somewhat conservative approach, you can either mention them in the prompt, or you can use the `sophon debug` command to automatically debug based on the output of any command you choose.

### Debugging Browser Applications

If Chrome is installed, Sophon can automatically debug browser applications by catching errors and reading the console logs.

If a plan calls for starting a browser, Sophon will automatically include a `sophon browser` call to start Chrome and monitor for errors.
   
### Execution Config

Control whether Sophon can execute commands:

```bash
sophon set-config can-exec true  # Allow command execution (default)
sophon set-config can-exec false # Disable command execution
```

If you toggle `can-exec` to `false`, Sophon will completely skip writing any commands to `_apply.sh`.

Control whether commands are executed automatically after applying changes (be careful with this):

```bash
sophon set-config auto-exec true  # Auto-execute commands
sophon set-config auto-exec false # Prompt before executing (default)
```

## Automated Debugging

The `sophon debug` command repeatedly runs a terminal command, making fixes until it succeeds:

```bash
sophon debug 'npm test'  # Try up to 5 times (default)
sophon debug 10 'npm test'  # Try up to 10 times
```

This will:

1. Run the command and check for success/failure
2. If it fails, send the output to the LLM
3. Tentatively apply suggested fixes to your project files
4. If command is succesful after fixes, commit changes (if auto-commit is enabled). Otherwise, roll back changes and return to step 2.
5. Repeat until success or max tries reached

You can configure the default number of tries:

```bash
sophon set-config auto-debug-tries 10  # Set default to 10 tries
```

## Common Debugging Workflows

### Fixing Failing Tests

```bash
sophon debug 'npm test'
sophon debug 'go test ./...'
sophon debug 'pytest'
```

### Fixing Build Errors

```bash
sophon debug 'npm run build'
sophon debug 'go build'
sophon debug 'cargo build'
```

### Fixing Linting Errors

```bash
sophon debug 'npm run lint'
sophon debug 'golangci-lint run'
```

### Fixing Type Errors

```bash
sophon debug 'npm run typecheck'
sophon debug 'tsc --noEmit'
```

## A Manual Alternative

For a less automated approach, you can pipe the output of a command into `sophon chat` or `sophon tell`.

```bash
npm test | sophon tell 'npm test output'
go build | sophon chat 'what could be causing these type errors?'
```

This works similarly to `sophon debug` but without automatically applying changes and retrying.

Note that piping output into a prompt requires using the CLI directly in the terminal. You can't do it from inside the [REPL](../repl.md).

## Autonomy Matrix

Execution and debugging behavior is affected by your [autonomy level](./autonomy.md):

| Setting      | None | Basic | Plus | Semi | Full |
| ------------ | ---- | ----- | ---- | ---- | ---- |
| `can-exec`   | ❌   | ❌    | ✅   | ✅   | ✅   |
| `auto-exec`  | ❌   | ❌    | ❌   | ❌   | ✅   |
| `auto-debug` | ❌   | ❌    | ❌   | ❌   | ✅   |


With `full` autonomy, commands are automatically executed and debugged after changes are applied. For other levels, you'll be prompted to approve execution and debugging steps.

## Safety

Needless to say, you should be extremely careful when using full auto mode, `auto-exec`, `auto-debug`, and the `debug` command. They can make many changes quickly without any prompting or review, and can run commands that could potentially be destructive to your system. While the best LLMs are quite trustworthy when it comes to running commands and are unlikely to cause harm, it still pays to be cautious.

It's a good idea to make sure your git state is clean, and to check out an isolated branch before using these features.

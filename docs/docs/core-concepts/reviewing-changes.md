---
sidebar_position: 5
sidebar_label: Pending Changes
---

# Pending Changes

When you give Sophon a task, by default the changes aren't applied directly to your project files. Instead, they are accumulated in Sophon's version-controlled **sandbox** so that you can review them first.

## Review Menu

Once Sophon has finished with a task, you'll see a review menu with several hotkey options. These hotkeys act as shortcuts for the commands described below.

## Viewing Changes

### `sophon diff` / `sophon diff --ui`

When Sophon has finished with your task, you can review the proposed changes with the `sophon diff` command, which shows them in `git diff` format:

```bash
sophon diff
```

`--plain/-p`: Outputs the diff in plain text with no ANSI codes.

You can also view the changes in a local browser UI with the `sophon diff --ui` command:

```bash
sophon diff --ui
```

The UI view offers additional options:

- `--side-by-side/-s`: Show diffs in side-by-side view
- `--line-by-line/-l`: Show diffs in line-by-line view (default)

## Rejecting Files

If the plan's changes were applied incorrectly to a file, or you don't want to apply them for another reason, you can either [apply the changes](#applying-changes) and then fix the problems manually, _or_ you can reject the updates to that file and then make the proposed changes yourself manually.

To reject changes to a file (or multiple files), you can run `sophon reject`. You'll be prompted to select which files to reject.

```bash
sophon reject # select files to reject
```

You can reject _all_ currently pending files by passing the `--all` flag to the reject command, or you can pass a list of specific files to reject:

```bash
sophon reject --all
sophon reject file1.ts file2.ts
```

If you rejected a file due to the changes being applied incorrectly, but you still want to use the code, either scroll up and copy the changes from the plan's output or run `sophon convo` to output the full conversation and copy from there. Then apply the updates to that file yourself.

## Applying Changes

Once you're happy with the plan's changes, you can apply them to your project files with `sophon apply`:

```bash
sophon apply
```

### Apply Flags & Config

Sophon v2 introduced several [new config settings and flags](./configuration.md) for the `apply` command that give you control over what happens after changes are applied.

### Command Execution & Debugging

After applying changes, Sophon can automatically execute pending commands. This is useful for running tests, starting servers, or performing other actions that verify the changes work as expected.

If commands fail, the changes are rolled back. Depending on the autonomy level and config, Sophon will then either attempt to debug automatically or prompt you with debugging options.

## Auto-Applying Changes

When `auto-apply` is enabled, Sophon will automatically apply changes after a plan is complete without prompting or review. This is enabled at the `full` [autonomy level](./autonomy.md), and also during auto-debugging.

## Apply + Full Auto

You can also apply changes and debug in full auto mode with the `--full` flag:

```bash
sophon apply --full
```

## Autonomy Matrix

| Setting       | None | Basic | Plus | Semi | Full |
| ------------- | ---- | ----- | ---- | ---- | ---- |
| `auto-apply`  | ❌   | ❌    | ❌   | ❌   | ✅   |
| `auto-exec`   | ❌   | ❌    | ❌   | ❌   | ✅   |
| `auto-debug`  | ❌   | ❌    | ❌   | ❌   | ✅   |
| `auto-commit` | ❌   | ❌    | ✅   | ✅   | ✅   |

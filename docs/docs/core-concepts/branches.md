---
sidebar_position: 8
sidebar_label: Branches
---

# Branches

Branches in Sophon allow you to easily try out multiple approaches to a task and see which gives you the best results. They work in conjunction with [version control](./version-control.md). Use cases include:

- Comparing different prompting strategies.
- Comparing results with different files in context.
- Comparing results with different models or model-settings.
- Using `sophon rewind` without losing history (first check out a new branch, then rewind).

## Creating a Branch

To create a new branch, use the `sophon checkout` command:

```bash
sophon checkout new-branch
sdxd new-branch # alias
```

## Switching Branches

To switch to a different branch, also use the `sophon checkout` command:

```bash
sophon checkout existing-branch
```

## Listing Branches

To list all branches, use the `sophon branches` command:

```bash
sophon branches
```

## Deleting a Branch

To delete a branch, use the `sophon delete-branch` command:

```bash
sophon delete-branch branch-name
```

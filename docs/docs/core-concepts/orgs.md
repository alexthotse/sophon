---
sidebar_position: 13
sidebar_label: Collaboration / Orgs
---

# Collaboration and Orgs

While so far Sophon is mainly focused on a single-user experience, we plan to add features for sharing, collaboration, and team management in the future, and some groundwork has already been done. **Orgs** are the basis for collaboration in Sophon.

## Multiple Users

Orgs are helpful already if you have multiple users using Sophon in the same project. Because Sophon outputs a `.sophon` file containing a bit of non-sensitive config data in each directory a plan is created in, you'll have problems with multiple users unless you either get each user into the same org or put `.sophon` in your `.gitignore` file. Otherwise, each user will overwrite other users' `.sophon` files on every push, and no one will be happy.

## Domain Access

When starting out with Sophon and creating a new org, you have the option of automatically granting access to anyone with an email address on your domain.

## Invitations

If you choose not to grant access to your whole domain, or you want to invite someone from outside your email domain, you can use `sophon invite`:

```bash
sophon invite
```

## Joining an Org

To join an org you've been invited to, use `sophon sign-in`:

```bash
sophon sign-in
```

## Listing Users and Invites

To list users and pending invites, use `sophon users`:

```bash
sophon users
```

## Revoking Users and Invites

To revoke an invite or remove a user, use `sophon revoke`:

```bash
sophon revoke
```

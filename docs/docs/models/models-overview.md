---
sidebar_position: 1
sidebar_label: Overview
---

# Models Overview

By default, Sophon uses a mix of Anthropic, OpenAI, and Google models. While this is a good starting point, and is the recommended way to use Sophon for most users, full customization of models and providers is also supported.

## Roles and Model Packs

Sophon has multiple [roles](./roles.md) which are responsible for different aspects of planning, coding, and applying changes. Each role can be assigned a different model. A **model pack** is a mapping of roles to specific models.

## Built-in Models and Model Packs

Sophon provides a curated set of built-in models and model packs.

You can see the list of available model packs with:

```bash
\model-packs # REPL
sophon model-packs # CLI
```

You can see the details of a specific model pack with:

```bash
\model-packs show model-pack-name # REPL
sophon model-packs show model-pack-name # CLI
```

You can see the list of available models with:

```bash
\models available # REPL
sophon models available # CLI
```

## Model Settings

You can see the model settings for the current plan with:

```bash
\models # REPL
sophon models # CLI
```

And you can see the default model settings for new plans with:

```bash
\models default # REPL
sophon models default # CLI
```

You can change the model settings for the current plan with:

```bash
\set-model # REPL
sophon set-model # CLI
```

And you can set the default model settings for new plans with:

```bash
\set-model default # REPL
sophon set-model default # CLI
```

[More details on model settings](./model-settings.md)

## Providers

Sophon offers flexibility on the providers you can use to serve models.

If you use [Sophon Cloud](../hosting/cloud.md) in **Integrated Models Mode**, you can use Sophon credits to pay for AI models. In that case, you won't need to worry about providers, provider accounts, or API keys.

If instead you use **BYO API Key Mode** with Sophon Cloud, or if you [self-host](../hosting/self-hosting/local-mode-quickstart.md) Sophon, you'll need to set API keys (or other credentials) for the providers you want to use. Multiple built-in providers are supported.

If you're self-hosting, you can also configure custom providers—they just need to be OpenAI-compatible.

[More details on providers](./model-providers.md)

## Custom Models, Providers, and Model Packs

You can configure custom models, providers, and model packs with a dev-friendly JSON config file:

```bash
\models custom # REPL
sophon models custom # CLI
```

[More details on custom models, providers, and model packs](./custom-models.md)

## Model Performance

While you can use Sophon with many different providers and models as described above, Sophon's prompts have mainly been written and tested against the built-in models and model packs, so you should expect them to give the best results.

## Local Models

Sophon supports local models via [Ollama](https://ollama.com/). For more details, see the [Ollama Quickstart](./ollama.md).



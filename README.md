# Go Mine Lang

[![tests](https://github.com/jfelipearaujo/gominelang/actions/workflows/tests.yml/badge.svg)](https://github.com/jfelipearaujo/gominelang/actions/workflows/tests.yml)
[![version](https://img.shields.io/github/v/release/jfelipearaujo/gominelang.svg)](https://github.com/jfelipearaujo/gominelang/releases/latest)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/jfelipearaujo/gominelang#section-readme)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jfelipearaujo/gominelang/blob/main/LICENSE)

This project was created to help the translation of Mod's content from a language to another one using Google Translator. Always review the translation after use this tool to avoid localization issues or undesired changes..

## How to use

To use this tool, download the latest release according to your operating system and architecture from the [releases page](https://github.com/jfelipearaujo/gominelang/releases/latest).

After downloaded, you can check the version by running:

```bash
gominelang version
```

The version will be printed on the screen like this:

```bash
Version: vX.X.X
```

Now create the configuration file `.gominelang.yaml` and add your mod configuration.

See [this](./.gominelang.example.yaml) file as an example.

Then, you can run it by using the following command:

```bash
gominelang
```

# Translation Engines

### Google Translate

To use Google Translate, you do not need to do anything. Just enable it in your `.gominelang.yaml` file:

```yaml
engine:
  google_translate:
    enabled: true
```

### OpenAI

To use OpenAI, you need to create an account and get an API key.

Then you need to add the following configuration to your `.gominelang.yaml` file:

```yaml
engine:
  open_ai:
    enabled: true
    api_key: your-api-key
```

ATTENTION: The API key is sensitive information, so do not share it with anyone.

## LLM Costs

You don't need to worry about the costs of the translations, because this can check if the translation must be created/updated or not. In the first execution, a `.gominelang.db` file will be created in the same folder of the `.gominelang.yaml` file. This database will be used to store the hashs of the files that have already been translated. If the hash of the file is the same as the one in the database, the file will not be translated again.

## Contributing

If you want to contribute to this project, you can do it by opening an issue or a pull request.

## LICENSE

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

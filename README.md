# Go Mine Lang

[![tests](https://github.com/jfelipearaujo/gominelang/actions/workflows/tests.yml/badge.svg)](https://github.com/jfelipearaujo/gominelang/actions/workflows/tests.yml)
[![version](https://img.shields.io/github/v/release/jfelipearaujo/gominelang.svg)](https://github.com/jfelipearaujo/gominelang/releases/latest)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/jfelipearaujo/gominelang#section-readme)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jfelipearaujo/gominelang/blob/main/LICENSE)

This project was created to help the translation of Mod's content from a language to another one using Google Translator. Always review the translation after use this tool to avoid localization issues or undesired changes..

## How to use

To use this CLI, you need to have GoLang installed and then run the following command:

```bash
go install github.com/jfelipearaujo/gominelang@latest
```

After installed, you can check the version by running:

```bash
gominelang version
```

The version will be printed on the screen like this:

```bash
GoMineLang Version: vX.X.X
```

Now create the configuration file `.gominelang.yaml` and add your mod configuration.

See [this](./.gominelang.yaml) file as an example.

Then, you can run it by using the following command:

```bash
gominelang
```

## Contributing

If you want to contribute to this project, you can do it by opening an issue or a pull request.

## LICENSE

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

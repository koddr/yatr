# yatr – A modern alternative to `Makefile` with a few little bells and whistles

[![Go version][go_version_img]][yatr_go_dev_url]
[![Go report][go_report_img]][go_report_url]
![Code coverage][go_code_coverage_img]
[![Wiki][wiki_img]][wiki_url]
[![License][license_img]][license_url]

**Y**et **A**nother **T**ask **R**unner (or `yatr` for a short) allows 
you to organize and automate your routine operations that you normally do in 
`Makefile` (or else) for each project.

Features:

- 100% **free** and **open source**
- can run any task in **async** mode (in separated goroutines)
- supports **popular** formats of the task file (JSON, YAML, TOML, HCL)
- has a **beautiful** appearance and **understandable** output

<img src="https://github.com/koddr/yatr/assets/11155743/28c87af5-9a04-4761-a8d8-21db42bb62fd" width="960" alt="img"/>

## ⚡️ Quick start

First, [download][go_download] and install **Go**. Version `1.20` (or higher)
is required.

Installation is done by using the [`go install`][go_install] command:

```console
go install github.com/koddr/yatr@latest
```

Prepare your task file in one of the supported formats:

- `json` – JSON ([example][json_example_file])
- `yaml` / `yml` – YAML ([example][yaml_example_file])
- `toml` – TOML ([example][toml_example_file])
- `tf` – Terraform, HCL ([example][tf_example_file])

> 💡 Note: See the repository's [Wiki page][wiki_url] to understand 
> structure of the task file.

Next, run `yatr` with (_or without_) options:

```console
yatr -p ./path/to/tasks-file.json
```

Done! 🎉 Your tasks have been executed.

## 🧩 Options

- `-p [path]` set a **path** to the file with tasks in any of the supporting 
  format (see [Wiki][wiki_tasks_file_url]);

## 🏆 A win-win cooperation

And now, I invite you to participate in this project! Let's work **together** to
create the **most useful** tool for developers on the web today.

- [Issues][yatr_issues_url]: ask questions and submit your features.
- [Pull requests][yatr_pull_request_url]: send your improvements to the current.

Your PRs & issues are welcome! Thank you 😘

## ⚠️ License

[`yatr`][yatr_url] is free and open-source software licensed under the
[Apache 2.0 License][license_url], created and supported with 🩵 for people and
robots by [Vic Shóstak][author].

[go_download]: https://golang.org/dl/
[go_install]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_version_img]: https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none
[go_report_url]: https://goreportcard.com/report/github.com/koddr/yatr
[go_code_coverage_img]: https://img.shields.io/badge/code_coverage-79%25-success?style=for-the-badge&logo=none
[wiki_img]: https://img.shields.io/badge/docs-wiki_page-blue?style=for-the-badge&logo=none
[wiki_url]: https://github.com/koddr/yatr/wiki
[wiki_tasks_file_url]: https://github.com/koddr/yatr/wiki#file-with-tasks
[json_example_file]: https://github.com/koddr/yatr/blob/main/examples/tasks.json
[yaml_example_file]: https://github.com/koddr/yatr/blob/main/examples/tasks.yaml
[toml_example_file]: https://github.com/koddr/yatr/blob/main/examples/tasks.toml
[tf_example_file]: https://github.com/koddr/yatr/blob/main/examples/tasks.tf
[license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[license_url]: https://github.com/koddr/yatr/blob/main/LICENSE
[yatr_url]: https://github.com/koddr/yatr
[yatr_go_dev_url]: https://pkg.go.dev/github.com/koddr/yatr
[yatr_issues_url]: https://github.com/koddr/yatr/issues
[yatr_pull_request_url]: https://github.com/koddr/yatr/pulls
[author]: https://github.com/koddr

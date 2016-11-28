# mcedit-syntax-terraform

This package can be used to generate mcedit (Midnight Commander) syntax file for terraform. It extracts terraform functions list, AWS resources and AWS data sources list and generates syntax file. Colors scheme can be customized as arguments or in template file

## Syntax file generation

    go get github.com/envimate/mcedit-syntax-terraform
    mcedit-syntax-terraform > tf.syntax

You can specify different colors

## Usage
```
Usage of colors:
  -data-source string
    	color of data sources (default "brightcyan")
  -func string
    	color of functions (default "yellow")
  -keyword string
    	color of keywords (default "white")
  -provider string
    	color of providers (default "red")
  -resource string
    	color of resources (default "brightcyan")
  -var string
    	color of variables (default "yellow")
```

For more information about syntax file format and allowed colors list please refer to [mcedit manual](http://linuxcommand.org/man_pages/mcedit1.html).

## Adding into mcedit
You can generate new `tf.syntax` with different color scheme, or simply download it from releases page.

Copy the tf.syntax into your syntax folder `/usr/share/mc/syntax/`, or if you are using user-specific configuration, than `$HOME/.local/share/mc/mcedit/`.

You need to register new file type in your `Syntax` file, by adding following lines before `unknown` file type

    file ..\*\\.(tf|tf)$ Terraform
    include tf.syntax

`Syntax` file usually can be found in `/usr/share/mc/syntax/Syntax` or `$HOME/.config/mc/mcedit/Syntax`

# Known issues
There is a formating issue with lines like

    source_code_hash = "${base64sha256(file("email-function.zip"))}"

Terraform supports quotes in quotes, but mcedit cannot parse that propely, and inner part is being shown with color scheme of default context.

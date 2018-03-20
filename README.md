# mcedit-syntax-terraform

This package can be used to generate mcedit (Midnight Commander) syntax file for terraform (.tf) file format highlighting.
It extracts terraform functions list, AWS resources and AWS data sources list and generates syntax file.
Colors scheme can be customized as arguments or in template file

## Adding into mcedit
You can generate new `tf.syntax` with different color scheme, or simply download latest [release](../../releases/download/v0.2/tf.syntax).

Copy the tf.syntax into your syntax folder `/usr/share/mc/syntax/`, or if you are using user-specific configuration, than `$HOME/.local/share/mc/mcedit/`.

You need to register new file type in your `Syntax` file, by adding following lines before `unknown` file type

    file ..\*\\.(tf|tfvars)$ Terraform
    include tf.syntax

`Syntax` file usually can be found in `/usr/share/mc/syntax/Syntax` or `$HOME/.config/mc/mcedit/Syntax`

## Preview

![Terraform highlight Preview](https://gist.githubusercontent.com/spirius/8cd532841e612c55deae244c7eef9180/raw/f8fbf7cf12d81a3daea555b339f8f40d3765ec1d/terraform-syntax-preview.png)

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

# Known issues
There is a formating issue with lines like

    source_code_hash = "${base64sha256(file("email-function.zip"))}"

Terraform supports quotes in quotes, but mcedit cannot parse that propely, and inner part is being shown with color scheme of default context.

# Other
If you want to get default black background color, edit `$HOME/.config/mc/ini` and change `basecolor` to

    base_color=editnormal=default,black:editwhitespace=default,red

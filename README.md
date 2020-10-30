# Empty DCI TT

A simple tool enables a user to create a minimal ST 428-7 Subtitle XML Document in accordance with ISDCF Technicial Doc 16 - SMPTE ST 428-7 D-Cinema Distribution Master Subtitle - Minimal Empty Document Requirements as per said requirements stipulated in RDD 52 - SMPTE DCP Bv2.1 Application Profile available at [https://doi.org/10.5594/SMPTE.RDD52.2020](https://doi.org/10.5594/SMPTE.RDD52.2020).



## Installation & Build

Empty TT is a multi-platform tool and has been built under Windows/x86_64, Darwin/x86_64 (macOS 10.12 or higher) and linux/x86_64. It relies on ASDCPlib if you want to wrap the generated XML and anciliary resources into D-Cinema MXF track files. ASDCPlib can be found [here](https://github.com/cinecert/asdcplib).



The tt library can be used in two ways;

- With the command line wrapper ''empty-tt'

- Or with the API directly. See /tt/lib.go or the GoDoc reference for details.



Pre-compiled binaries are provided for convenience and can be found under [Releases](https://github.com/jack-watts/empty-dci-tt/releases).



To build from source using Go

```go
go get -u
sh build.sh
```



## Usage

  -T                           - write MXF trackfile, requires '-d'
  -d <int>              - set the duration of the track file. (default 24)
  -e                           - encrypt trackfile
  -image                  - Inidcate that image profile is to be used.
  -l <string>         - set the RFC 5646 Language subtag (default "en")
  -m <int>             - set the DisplayType.'0'=MainSubtitle,'1'=ClosedCaption. (default "0")
  -o <string>       - set the output path, Default is StdOut
  -p <string>       - set the frame rate of the track file. (default "24")
  -r <int>               - set the ReelNumber (default 1)
  -t <string>        - set the ContentTitleText value. (default "No Title")
  -text                      - Inidcate that text profile is to be used. (default true)
  -x <string>       - path to 428-7 XML to use as template

### Examples

The follwoing examples showcase the different command expressions that can be used.

**1. Create a MainSubtitle Text profile document**

```bash
$ empty-tt -text -p 24 -m 0 -r 1 -t "MyTitle"
```

**2. Create an MainSubtitle Image profile document**

```bash
$ empty-tt -image -p 24 -m 0 -r 1 -t "MyTitle"
```

**3. Create a MainSubtitle Text profile document and MXF track file**

```bash
$ empty-tt -text -T -p 24 -m 0 -r 1 -t "MyTitle" -d 48 -o <path-to-dir>
```

**4. Create a ClosedCaption Text profile document from XML template and encrypted MXF track file**

```bash
$ empty-tt -text -T -e -p 24 -m -r 1 -t "MyTitle" -d 48 -x <path-to-xml-file> -o <path-to-dir>
```

#### Notes

1. Writes to StdOut if no output path is specified.

2. When writing to StdOut, no anciliary resources are generated.

3. Using the MXF option "-T" also requires "-o".

4. Generates a unique PNG image every execution.

5. Any integer value greater than 1 for option "-m" results in a ClosedCaption DisplayType.

6. An template document with a DCST 2007 namespace will be rejected.

7. Documents are written with a DCST 2014 namespace by default.

8. The shipped bundled default Font is Arial Unicode with no glyph structures present. It has a file name of constant = "232c45d8-fde8-4e5e-86b9-86e96354daf3".

9. Available flags can be invoked out of order.

10. file prefix and suffix are handled autmatically and are structed as follows:
    
    1. XML: <uuid>_<reelNo>.xml
    
    2. MXF: <uuid>_<reelNo>_sub.mxf | <uuid>_<reelNo>_cap.mxf

## License

See [GPL-3.0 License](https://github.com/jack-watts/empty-dci-tt/blob/main/LICENSE).

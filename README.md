# dd-log-cat

Extracts the `Message` attribute for Datadog log exports (CSV).

```sh
# From stdin
$ cat extract-2025-04-22T10_59_52.639Z.csv | dd-log-cat

# From a file
$ dd-log-cat extract-2025-04-22T10_59_52.639Z.csv
```

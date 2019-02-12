# gohaste-client

Simple command line client for [hastebin servers](https://github.com/seejohnrun/haste-server) made in **Go**

This should work on all platforms where Go is supported 🤗

## Usage

You can pipe content into the application stdin or choose a file to upload

```
ps aux | gohaste-client
```

``` 
gohaste-client my_custom_file.txt
```

By default all the content is uploaded to [https://hastebin.com] you can however specify your custom hastebin server URL using the `-server` flag.
You can also use the `-raw` flag to get the plaintext URL

## License

`gohaste-client` is licensed under the **MIT** license

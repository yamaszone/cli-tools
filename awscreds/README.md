# Usage
Simple CLI tool to extract key/secret from an AWS credentials file (`~/.aws/credentials`).
```
$ awscreds --help
Usage of awscreds:
  -field string
    	Credentials field {key|secret|both} to extract. (default "key")
  -profile string
    	AWS profile to be used to parse key/secret. (default "default")
```

# Installation
- Linux
  - [CoreOS] `curl -L https://github.com/yamaszone/cli-tools/releases/download/v1.1.0/awscreds-linux-amd64 -o ~/awscreds && chmod +x ~/awscreds && sudo mv ~/awscreds /opt/bin/awscreds`
  - [Ubuntu|Fedora] `curl -L https://github.com/yamaszone/cli-tools/releases/download/v1.1.0/awscreds-linux-amd64 -o /usr/local/bin/awscreds && chmod +x /usr/local/bin/awscreds`
- macOS
  - `curl -L https://github.com/yamaszone/cli-tools/releases/download/v1.1.0/awscreds-darwin-amd64 -o /usr/local/bin/awscreds && chmod +x /usr/local/bin/awscreds`

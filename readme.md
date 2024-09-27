# read_discord_go

Extract discord's package.zip messages to a readable format. (currently outputs .txt files)

## Usage
```bash
go run . --input package.zip

go build .
./read_discord_go --input package.zip

# use air (https://github.com/air-verse/air) for hot reloading
air
```

## Output
```
out/
    - 1234567890.txt
    - 1234567891.txt
    - 1234567892.txt
```

## File Format
```
Channel: Channel Name
Channel Id: [ID]

Messages:
[ID] Message content
```

## To Do
- [ ] Add support for attachments
- [ ] Output an interface for the messages
- [ ] Parse user information
- [ ] Add option for output path
- [ ] Prettify output file format

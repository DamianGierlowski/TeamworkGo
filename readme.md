

# CSV customer reader - created for interview Teamwork

## Example CSV
```csv
first_name,last_name,email,gender,ip_address
Mildred,Hernandez,mhernandez0@github.io,Female,38.194.51.128
Bonnie,Ortiz,bortiz1@cyberchimps.com,Female,197.54.209.129
Dennis,Henry,dhenry2@hubpages.com,Male,155.75.186.217
```

## How to use

## Pull
```bash
git clone https://github.com/DamianGierlowski/TeamworkGoTests.git
```

### Start
```bash
go run cmd/main.go -input customers.csv -output data.json
```

### Tests

```bash
go test ./internal/customerimporter
```

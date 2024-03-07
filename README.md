# JWTSecretGen

JWTSecretGen is a tool written in Go for generating JSON Web Token (JWT) secrets securely.

## Usage

To generate a JWT secret, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/nakinami/JWTSecretGen.git

2. Navigate to the directory containing the Go script:

   ```bash
   cd JWTSecretGen

3. Run the Go script:
 
   ```bash
    go run jwt.go

## Notes

- The generated JWT secret is encoded using base64.URLEncoding. Make sure to store it securely and never expose it in plaintext.

- You can adjust the length of the secret by modifying the `secretLength` variable in the `generateJWTSecret` function. The default length is 64 bytes.

- Ensure that you have Go installed on your machine before running the script. You can download and install Go from the [official website](https://golang.org/).

module github.com/fakovacic/amadeus-golang/examples

go 1.14

replace (
	github.com/fakovacic/amadeus-golang/amadeus => ./../amadeus
)


require (
	github.com/fakovacic/amadeus-golang/amadeus v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.3.0
)

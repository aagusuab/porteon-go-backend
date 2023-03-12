package controllers

var (
	TestDriverId              string
	TestTruckId               string
	TestOriginLocationId      string
	TestDestinationLocationId string
	TestLocationId            string
)

func BeforeEach() {
	TestDriverId = randomString(16)
	TestTruckId = randomString(16)
	TestOriginLocationId = randomString(16)
	TestDestinationLocationId = randomString(16)
	TestLocationId = randomString(16)
}

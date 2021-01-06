package actions

type AWSSetting struct {
	Region string
}

func GetAWSSetting() AWSSetting {
	return AWSSetting{
		Region: "eu-north-1",
	}
}

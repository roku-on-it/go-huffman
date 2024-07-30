package main

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must2(_ any, err error) {
	if err != nil {
		panic(err)
	}
}

func Must3(_ any, _ any, err error) {
	if err != nil {
		panic(err)
	}
}

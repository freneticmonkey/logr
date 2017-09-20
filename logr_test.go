package logr

import "testing"

func TestCreate(t *testing.T) {

	writer := NewWriter(INFO, "")

	writer.Info("Hello World")
	writer.Ok("Good to go!")
	writer.Attention("This just happened")
	writer.Warn("Hmm. Something isn't ok")
	writer.Error("Everything is broken")

	return
}

func TestDump(t *testing.T) {

	testObject := struct {
		Name      string
		Age       int
		Addresses []string
	}{
		Name: "Joe B.",
		Age:  21,
		Addresses: []string{
			"1 Hope St, Hopeville",
			"PO Box 1, Hopeville",
		},
	}

	Dump(testObject)

	return
}

func TestDiffs(t *testing.T) {

	testObjectA := struct {
		Name      string
		Age       int
		Addresses []string
	}{
		Name: "Joe B.",
		Age:  21,
		Addresses: []string{
			"1 Hope St, Hopeville",
			"PO Box 1, Hopeville",
		},
	}

	testObjectB := struct {
		Name      string
		Age       int
		Addresses []string
	}{
		Name: "Joe C.",
		Age:  21,
		Addresses: []string{
			"1 Hope St, Hopeville",
			"PO Box 1, Hopeville",
		},
	}

	DumpDiff(testObjectA, testObjectB)

	return
}

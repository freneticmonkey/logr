# Logr

A basic logging library for generating logs in a standardised format.  Also supports:
- Colourised output using github.com/fatih/color.
- Object dumping and diffing using github.com/davecgh/go-spew

## Creating a Writer and generating a log
Logs are generated using a Writer instance.

    writer := logr.NewWriter(INFO, "")

    writer.Info("Hello World")
    writer.Ok("Good to go!")
    writer.Attention("This just happened")
    writer.Warn("Hmm. Something isn't ok")
    writer.Error("Everything is broken")

Will generate the following output

<pre>
    2017/09/20 20:59:21 INFO:  Hello World
    <span style="color:green">2017/09/20 20:59:21 OK:    Good to go!</span>
    <span style="color:yellow">2017/09/20 20:59:21 ATTN:  This just happened</span>
    <span style="color:magenta">2017/09/20 20:59:21 WARN:  Hmm. Something isn't ok</span>
    <span style="color:red">2017/09/20 20:59:21 ERROR: Everything is broken</span>
</pre>

## Dumping Objects
Logr also support pretty printing objects for easy debugging.  


A simple example:

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

    logr.Dump(testObject)

will output:

    (struct { Name string; Age int; Addresses []string }) {
     Name: (string) (len=6) "Joe B.",
     Age: (int) 21,
     Addresses: ([]string) (len=2 cap=2) {
      (string) (len=20) "1 Hope St, Hopeville",
      (string) (len=19) "PO Box 1, Hopeville"
     }
    }

## Diffing Objects
Another useful tool for debugging or displaying errors during tests is diffing objects.  Logr calculates and displays object diffs by pretty printing both objects and performing a line-by-line string comparison over the output.

For example, given the following structs:

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

the following output is generated:

<pre>
<span style="color:green">
        (struct { Name string; Age int; Addresses []string }) {
<span style="color:red">    <<   Name: (string) (len=6) "Joe B.",</span>
<span style="color:yellow">    >>   Name: (string) (len=6) "Joe C.",</span>
         Age: (int) 21,
         Addresses: ([]string) (len=2 cap=2) {
          (string) (len=20) "1 Hope St, Hopeville",
          (string) (len=19) "PO Box 1, Hopeville"
         }
        }
</span>
</pre>


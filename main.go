package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var htmlStr = `
<html>
  <body>
    <script src="ledger.js"></script>
    <div>This window will autoclose shortly...</div>
    <script>
      function callback(event) {
        if (JSON.stringify(event.response) == "{\"command\":\"has_session\",\"success\":true}") {
          close();
        }
      };
      Ledger.init({ callback: callback });
      Ledger.sendPayment('<ADDR>',<AMT>,'');
    </script>
  </body>
</html>
`

func main() {

	htmlStr = strings.Replace(htmlStr, "<ADDR>", "1FdawJAuUBMvEa4r4Dm3qNbBvUwZBiRy3Q", 1)
	htmlStr = strings.Replace(htmlStr, "<AMT>", "0.666", 1)

	tempPath := os.ExpandEnv("./temp.html")

	htmlBytes := []byte(htmlStr)
	err := ioutil.WriteFile(tempPath, htmlBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("google-chrome", tempPath)
	//cmd.Stdin = strings.NewReader("some input")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	//remove the temp file
	err = os.Remove(tempPath)
	if err != nil {
		log.Fatal(err)
	}
}

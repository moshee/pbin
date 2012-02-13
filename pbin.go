package main

import (
	"os"
	"fmt"
	"net/url"
	"net/http"
	"flag"
	"io/ioutil"
)


var (
	use_stdin		bool
	paste_private	bool
	paste_name		string
	paste_expire	string
	paste_format	string
	in_file			string
)

var expire_times = map[string]string{
	"never": "N",
	"minutes": "10M",
	"hours": "1H",
	"days": "1D",
	"months": "1M",
}

func main() {
	flags := new(flag.FlagSet)
	flags.Usage = func() {
		println("Usage:", os.Args[0], "[options] [filename]")
		flags.PrintDefaults()
		println(`Option ‘-x’ takes one of five arguments:
  never  Never expire (default)
minutes  Expire in 10 minutes
  hours  Expire in 1 hour
   days  Expire in 1 day
  month  Expire in 1 month`)
		os.Exit(0)
	}

	flags.BoolVar(&use_stdin, "s", false, "Accept input from STDIN")
	flags.BoolVar(&paste_private, "p", false, "Private paste")
	flags.StringVar(&paste_name, "n", "", "Paste name")
	flags.StringVar(&paste_expire, "x", "never", "Time before paste expires")
	flags.StringVar(&paste_format, "f", "text", "Paste language")

	flags.Parse(os.Args[1:])

	values := url.Values{
		"api_option":		{"paste"},
		"api_dev_key":		{pastebin_dev_key},
		"api_paste_name":	{paste_name},
		"api_paste_format":	{paste_format},
	}
	if expire_time, ok := expire_times[paste_expire]; ok {
		values.Add("api_paste_expire", expire_time)
	} else {
		println("Invalid paste expiration time (-help for details)")
		return
	}

	if paste_private {
		values.Add("api_paste_private", "1")
	} else {
		values.Add("api_paste_private", "0")
	}

	var (
		input []byte
		err error
	)
	if use_stdin {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			println(err.Error())
			return
		}
		if len(input) == 0 {
			println("There's no text!")
			return
		}
	} else {
		in_file = flags.Arg(0)
		if in_file == "" {
			println("No input file specified")
			return
		}
		input, err = ioutil.ReadFile(in_file)
		if err != nil {
			println(err.Error())
			return
		}
	}
	values.Add("api_paste_code", string(input))

	resp, err := http.PostForm("http://pastebin.com/api/api_post.php", values)
	if err != nil {
		println(err.Error())
		return
	}

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(string(resp_body))
}

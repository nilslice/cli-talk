func main() {
	flag.IntVar(&port, "port", 8080, "port for ponzu to bind its HTTP listener")
	flag.IntVar(&httpsport, "https-port", 443, "port for ponzu to bind its HTTPS listener")
	flag.BoolVar(&https, "https", false, "enable automatic TLS/SSL certificate management")
	flag.BoolVar(&devhttps, "dev-https", false, "[dev environment] enable automatic TLS/SSL certificate management")
	flag.BoolVar(&dev, "dev", false, "modify environment for Ponzu core development")
	flag.BoolVar(&cli, "cli", false, "specify that information should be returned about the CLI, not project")
	flag.StringVar(&fork, "fork", "", "modify repo source for Ponzu core development")
	flag.StringVar(&gocmd, "gocmd", "go", "custom go command if using beta or new release of Go")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println(usage)
		os.Exit(0)
	}

	switch args[0] {
	case "help", "h":
		if len(args) < 2 {
			fmt.Println(usageHelp)
			fmt.Println(usage)
			os.Exit(0)
		}

// ...and so on
}
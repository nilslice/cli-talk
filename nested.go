func main() {
	// ...
	args := flag.Args()
	switch args[0] {
	case "help", "h":
		if len(args) < 2 {
			fmt.Println(usageHelp)
			fmt.Println(usage)
			os.Exit(0)
		}

		switch args[1] {
		case "new":
			fmt.Println(usageNew)
			os.Exit(0)
		// ...and so on

	case "generate", "gen", "g":
		// check what we are asked to generate
		switch args[1] {
		case "content", "c":
			err := generateContentType(args[2:])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		default:
			msg := fmt.Sprintf("Generator '%s' is not implemented.", args[1])
			fmt.Println(msg)
			fmt.Println(usageGenerate)			
		}
	}
}
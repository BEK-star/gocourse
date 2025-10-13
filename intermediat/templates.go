package intermediat

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	// tmpl := template.Must(template.New("example").Parse("Hello, {{.Name}}!\n"))

	// data := struct {
	// 	Name string
	// }{
	// 	Name: "World",
	// }

	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"welcome":      "Hello, {{.name}}! Glad you joined\n",
		"notification": "{{.nm}}, you have a new notification {{.notification}}\n",
		"error":        "Oops, an error occurred: {{.em}}\n",
	}

	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Join")
		fmt.Println("2. Notification")
		fmt.Println("3. Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose rhe option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}

		case "2":
			fmt.Println("Enter your notification: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"nm": name, "notification": notification}

		case "3":
			fmt.Println("Enter your error massage: ")
			errorMassage, _ := reader.ReadString('\n')
			errorMassage = strings.TrimSpace(errorMassage)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"em": errorMassage}

		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose right option")
			continue
		}

		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error, ", err)
		}

	}
}

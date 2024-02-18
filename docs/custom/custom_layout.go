package custom

import "fmt"

var CustomJS = fmt.Sprintf(`
	// custom title
	document.title = 'Swagger Dark mode Golang';

	// dark mode
	const style = document.createElement('style');
	style.innerHTML = %s;
	document.head.appendChild(style);
	`, "`"+customCSS+"`")

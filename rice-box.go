package main

import (
	"github.com/nkovacs/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file3 := &embedded.EmbeddedFile{
		Filename:    "css/styles.css",
		FileModTime: time.Unix(1714733583, 0),

		Content: string(""),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "html/about.html",
		FileModTime: time.Unix(1714733583, 0),

		Content: string("<!DOCTYPE html>\r\n<html>\r\n    <head>\r\n        <script src=\"https://unpkg.com/htmx.org@1.9.11\"></script>\r\n        <title>about chat</title>\r\n    </head>\r\n    <body style=\"background-color: black;color: white;\">\r\n        <h1>\r\n            В нашем динамично меняющемся мире связь и обмен информацией \r\n            становятся все более важными. Онлайн-чат предоставляет возможность \r\n            людям из разных уголков мира общаться в реальном времени, делиться мнениями, \r\n            идеями и опытом. Онлайн-чат может стать платформой для создания сообщества единомышленников, \r\n            объединенных общими интересами. Это может быть как место для неформального общения, \r\n            так и площадка для профессионального роста и развития.\r\n        </h1><br>\r\n        <h2>Оставить отзыв</h2>\r\n        <form>\r\n            <input type=\"text\" name=\"user-feedback\" id=\"user-feedback\" required>\r\n            <input type=\"submit\" value=\"sendfeed\">\r\n        </form>\r\n    </body>\r\n</html>"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "html/index.html",
		FileModTime: time.Unix(1714745543, 0),

		Content: string("<!DOCTYPE html>\r\n<html>\r\n    <head>\r\n        <script src=\"https://unpkg.com/htmx.org@1.9.11\"></script>\r\n        <script type=\"text/javascript\" src=\"../js/websockets.js\"></script>\r\n        <title>online chat site</title>\r\n    </head>\r\n    <body style=\"background-color: black;color: white;\">\r\n        <h1>ЧАТ</h1><span style=\"text-align: right;\"><a href=\"http://localhost:8085/about\">about</a></span>\r\n        <form hx-post=\"/sendMessage\" hx-target=\"#chat-box\" hx-swap=\"beforeend\" >\r\n            <input type=\"text\" name=\"user-message\" id=\"user-message\" required>\r\n            <input type=\"submit\" value=\"send\" id=\"send-message\">\r\n        </form>\r\n        <div name=\"chat-box\" id=\"chat-box\" style=\"height: 1000px; width: 200px; border:3px;border-color: white; border-style: solid;\">\r\n            {{ range . }}\r\n               <p> {{.}} </p>\r\n            {{ end }}\r\n        </div>\r\n    </body>\r\n</html>"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    "js/websockets.js",
		FileModTime: time.Unix(1714734620, 0),

		Content: string("clientWS = WebSocket(\"ws://localhost:8085/\")\r\n\r\nvar send_msg_btn = document.getElementById(\"send-message\")\r\nsend_msg_btn.onclick = function() {\r\n    clientWS.send(document.getElementById(\"user-message\").value)\r\n}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1714734179, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dir2 := &embedded.EmbeddedDir{
		Filename:   "css",
		DirModTime: time.Unix(1714733583, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file3, // "css/styles.css"

		},
	}
	dir4 := &embedded.EmbeddedDir{
		Filename:   "html",
		DirModTime: time.Unix(1714733583, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file5, // "html/about.html"
			file6, // "html/index.html"

		},
	}
	dir7 := &embedded.EmbeddedDir{
		Filename:   "js",
		DirModTime: time.Unix(1714734296, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file8, // "js/websockets.js"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir2, // "css"
		dir4, // "html"
		dir7, // "js"

	}
	dir2.ChildDirs = []*embedded.EmbeddedDir{}
	dir4.ChildDirs = []*embedded.EmbeddedDir{}
	dir7.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`static`, &embedded.EmbeddedBox{
		Name: `static`,
		Time: time.Unix(1714734179, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":     dir1,
			"css":  dir2,
			"html": dir4,
			"js":   dir7,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"css/styles.css":   file3,
			"html/about.html":  file5,
			"html/index.html":  file6,
			"js/websockets.js": file8,
		},
	})
}

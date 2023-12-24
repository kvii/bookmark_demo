# bookmark_demo

> book.pdf download from [https://www.gopl.io](https://www.gopl.io) and used for demonstration purpose only.

Intention: Give a place to specify a global offset when importing bookmarks.

Steps:

1. Find a book which not have bookmark (book.pdf).
2. Create bookmarks file from the content page (bookmark.json). Note the "settings" field.
https://github.com/kvii/bookmark_demo/blob/12338d28ce98a80bcf232d967dfe2b7c459d6a55/bookmark.json#L27-L29
3. Execute `pdfcpu bookmarks import book.pdf bookmark.json out.pdf`.
4. Open "out.pdf".

Expect: Click the bookmark "1.1. Hello, World" will point to page 20 (which is semantic page 1).

Got: Click the bookmark "1.1. Hello, World" will point to the cover page (which is physical page 1).

In fact, I've solve it myself. I've written a cmd tool to solve my requirement. Here is the step:

1. Execute `go run . book.pdf bookmark.json out2.pdf`.
2. Open "out2.pdf".

Now click the bookmark "1.1. Hello, World" will point to page 20 correctly.

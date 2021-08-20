When I get an "a" tag, need to:

- Look at all children. "a".FirstChild, then NextSibling of each of those
- Follow all children of these children, and their siblings.
- For each Type: TextNode, add to the link text, in order.

  A tab precedes this.

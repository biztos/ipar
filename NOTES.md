# ipar notes

dev notes, for now.

## App vs Site Level Config

Though it is possible to serve many sites from one app, this is probably a
bad idea because:

* It adds a *lot* of complexity to the app.
* The multi-site use-case is better served with reverse proxies.
* Unclear how to do multi-site certs across one listener on one port.

## Templating

what to do about dir-style templates?  map them or not?

    default.html - template knows!  (also default.txt, don't need HTML!)
    page.html - only if it's a page
    index.html - only if it's a request to the index level

What about config for mapping index pages?

Default logic would be if you have a 

So what is the template logic?

    IF REQUEST HAS PAGE:
        FIND MOST DETAILED MATCH FOR TEMPLATE DIRECTORY
        WITH THAT TAKE ... WHAT?

Problem is **I REALLY WANT blog.html INSTEAD OF blog/index.html**. Because
editors, dammit.

    1. Match directory to directory, most detailed directory wins.
    2. Within the directory, if you have an "index" template

**ALSO** what do we do about specified content for list pages? Index I suppose
if available.

    foo/bar/index.md --> /foo/bar
    foo/bar.md       --> /foo/bar ALSO, so complain?

Yes, complain, because either way should work.  The first is if you have
`bar` as its own group.

## OK, and another problem, from kisipar: rendering assets in local dir.

Given a page with dir `somepage` we want to be able to render stuff that is
local to the dir.  However the page is not rendered as under a dir, the
page is rendered with a path of `/somepage` and then looks for (e.g.)
`/foo.jpg` instead of `/somepage/foo.jpg`.

The smart way around that is to filter it when creating the pages, but that's
potentially dangerous because we have to walk the HTML right?  OR we can deal
with resources at the server level and if it's `/foo.jpg` with a referer from
`/somepage` then we serve `/foopage/foo.jpg` if it's available?

This is... a bit of a problem isn't it?

```go
s,err := site.New(dir|config)
if err != nil {
    panic(err)
}
s.Serve()


// TYPE SWITCH FOR PATH MAP, panic on default, serve 404 first.
type Thing struct {
    Foo string
}

func checkThing(i interface{}) {

    switch v := i.(type) {
    case *Thing:
        fmt.Println("HAVE A THING")
    case string:
        fmt.Println("HAVE A STRING")
    default:
        fmt.Println("DUNNO:",v)
    }
}

func main() {
    p := "i am a path"
    checkThing(p)
    o := &Thing{"HELO"}
    checkThing(o)
    checkThing(123)

```